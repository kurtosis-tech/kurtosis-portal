PORTAL_SERVICE_SUFFIX = "-service"
PORTAL_PROXY_SUFFIX = "-proxy"

# Portal server constants
SERVER_GRPC_PORT_ID = "server_grpc"
SERVER_GRPC_PORT = 9720
SERVER_TUNNEL_PORT_ID = "server_tunnel"
SERVER_TUNNEL_PORT = 9721
SERVER_GRPC_SERVICE_NAME = "kurtosis_portal_daemon.KurtosisPortalServer"

# Portal client constants
CLIENT_GRPC_PORT_ID = "client_grpc"
CLIENT_GRPC_PORT = 9731
CLIENT_GRPC_SERVICE_NAME = "kurtosis_portal_daemon.KurtosisPortalClient"

# Portal Proxy constants
PORTAL_PROXY_SINGLE_PORT = 9719
PORTAL_PROXY_SINGLE_PORT_ID = "proxy"

# MISC constants
ENVOY_IMAGE_NAME = "envoyproxy/envoy:v1.25-latest"
RESOURCE_FOLDER = "github.com/kurtosis-tech/kurtosis-portal/kurtosis/portal_launcher/resources/"
DEFAULT_CONTEXT_ID_AND_NAME = "default"


def launch_portal(plan, portal_image, service_name, server_only=False, current_context_id=DEFAULT_CONTEXT_ID_AND_NAME, remote_contexts={}, remote_host=None):
    portal_service_name = service_name + PORTAL_SERVICE_SUFFIX
    proxy_service_name = service_name + PORTAL_PROXY_SUFFIX

    envoy_yaml_template = read_file(src=RESOURCE_FOLDER + "envoy.yaml.tmpl")
    api_file_descriptor = plan.upload_files(src=RESOURCE_FOLDER + "portal_api_descriptor_file.pb")
    contexts_config_artifact = render_contexts_config(plan, current_context_id, remote_contexts)

    portal_ports = {
        SERVER_GRPC_PORT_ID: PortSpec(number=SERVER_GRPC_PORT),
        SERVER_TUNNEL_PORT_ID: PortSpec(number=SERVER_TUNNEL_PORT),
        # backend port is unused right now
    }
    if not server_only:
        portal_ports[CLIENT_GRPC_PORT_ID] = PortSpec(number=CLIENT_GRPC_PORT)

    cmd = []
    if server_only:
        cmd.append("--server-only")
    if remote_host:
        cmd.extend(["--remote-host", remote_host])
    portal_service = plan.add_service(
        name=portal_service_name,
        config=ServiceConfig(
            image=portal_image,
            ports=portal_ports,
            cmd=cmd,
            files={
                "/root/.config/kurtosis/": contexts_config_artifact,
            }
        ),
    )

    envoy_yaml = plan.render_templates(
        config={
            "envoy.yaml": struct(
                template=envoy_yaml_template,
                data={
                    "SERVER_GRPC_SERVICE_NAME": SERVER_GRPC_SERVICE_NAME,
                    "CLIENT_GRPC_SERVICE_NAME": CLIENT_GRPC_SERVICE_NAME,
                    "PORTAL_SERVICE_IP": portal_service.ip_address,
                    "PORTAL_SERVER_GRPC_PORT": SERVER_GRPC_PORT,
                    "PORTAL_CLIENT_GRPC_PORT": CLIENT_GRPC_PORT,
                    "PORTAL_PROXY_SINGLE_PORT": PORTAL_PROXY_SINGLE_PORT,
                }
            ),
        },
    )

    envoy_proxy = plan.add_service(
        name=proxy_service_name,
        config=ServiceConfig(
            image=ENVOY_IMAGE_NAME,
            ports={
                PORTAL_PROXY_SINGLE_PORT_ID: PortSpec(number=PORTAL_PROXY_SINGLE_PORT),
            },
            files={
                "/home/envoy/protos/": api_file_descriptor,
                "/etc/envoy/": envoy_yaml,
            }
        ),
    )

    # wait for both client and server availability, if relevant
    plan.wait(
        recipe=PostHttpRequestRecipe(
            port_id=PORTAL_PROXY_SINGLE_PORT_ID,
            endpoint="/server/grpc/ping",
            content_type="application/json",
            body="{}"
        ),
        service_name=proxy_service_name,
        field="code",
        assertion="==",
        target_value=200,
    )
    expected_code_for_client = 503 if server_only else 200 # if server only, client endpoint should be unavailable
    plan.wait(
        recipe=PostHttpRequestRecipe(
            port_id=PORTAL_PROXY_SINGLE_PORT_ID,
            endpoint="/client/grpc/ping",
            content_type="application/json",
            body="{}"
        ),
        service_name=proxy_service_name,
        field="code",
        assertion="==",
        target_value=expected_code_for_client,
    )

    return {
        "proxy": envoy_proxy,
        "portal": portal_service
    }


def remove_portal(plan, service_name):
    plan.remove_service(service_name + PORTAL_PROXY_SUFFIX)
    plan.remove_service(service_name + PORTAL_SERVICE_SUFFIX)


def render_contexts_config(plan, current_context_id, contexts_to_add):
    # add only the default context for now as we don't need any other
    default_context = new_local_context(DEFAULT_CONTEXT_ID_AND_NAME)
    configured_context = [default_context]

    for context_name in contexts_to_add:
        additional_context = new_remote_context(context_name, contexts_to_add[context_name])
        configured_context.append(additional_context)

    # render the list of configured contexts
    all_context_configs_json = json.encode(configured_context)

    # render the contexts config from the template
    contexts_config_template = read_file(src=RESOURCE_FOLDER + "contexts_config/contexts-config.json.tmpl")
    contexts_config_file = plan.render_templates(
        config={
            "contexts-config.json": struct(
                template=contexts_config_template,
                data={
                    "CURRENT_CONTEXT_ID": current_context_id,
                    "CONTEXT_CONFIGS": all_context_configs_json,
                }
            ),
        },
    )
    return contexts_config_file


def new_local_context(context_name):
    context_config_json_template = read_file(RESOURCE_FOLDER + "contexts_config/empty_local_kurtosis_context.json")
    context_config_template = json.decode(context_config_json_template)

    context_config_template["name"] = context_name
    context_config_template["uuid"]["value"] = context_name # for testing purposes UUID can be equal to name

    return context_config_template


def new_remote_context(context_name, remote_server_host):
    context_config_json_template = read_file(RESOURCE_FOLDER + "contexts_config/empty_remote_kurtosis_context.json")
    context_config_template = json.decode(context_config_json_template)

    context_config_template["name"] = context_name
    context_config_template["uuid"]["value"] = context_name # for testing purposes UUID can be equal to name
    context_config_template["remoteContextV0"]["host"] = remote_server_host

    return context_config_template
