portal_launcher = import_module("../portal_launcher/portal_launcher.star")

def test(plan, portal_image):
    # Create a server tha represents a Portal server running in the cloud
    server_name = "portal-server"
    portal_server = portal_launcher.launch_portal(plan, portal_image, server_name, server_only=True)

    # Create a client that is connected to the above Portal server
    client_name = "portal-client"
    remote_context_id="remote_context"
    remote_contexts = {
        remote_context_id: portal_server["portal"].ip_address
    }
    portal_client = portal_launcher.launch_portal(plan, portal_image, client_name, server_only=False, current_context_id=remote_context_id, remote_contexts=remote_contexts)

    # test the client is alive. TODO: add an endpoint to retrieve the currently selected context, and check it matches the remote context in this case
    response = plan.request(
        service_name=portal_client["proxy"].name,
        recipe=PostHttpRequestRecipe(
            port_id=portal_launcher.PORTAL_PROXY_SINGLE_PORT_ID,
            endpoint="/client/grpc/ping",
            content_type="application/json",
            body="{}"
        )
    )

    plan.verify(response["code"], "==", 200)
    plan.verify(response["body"], "==", "{}\n")
