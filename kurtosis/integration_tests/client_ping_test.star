portal_launcher = import_module("github.com/kurtosis-tech/kurtosis-portal/kurtosis/portal_launcher/portal_launcher.star")

def test(plan):
    service_name = "portal-client"
    portal_and_proxy = portal_launcher.launch_portal(plan, service_name)
    proxy = portal_and_proxy["proxy"]

    response = plan.request(
        service_name=proxy.name,
        recipe=PostHttpRequestRecipe(
            port_id=portal_launcher.PORTAL_PROXY_SINGLE_PORT_ID,
            endpoint="/client/grpc/ping",
            content_type="application/json",
            body="{}"
        )
    )

    plan.assert(response["code"], "==", 200)
    plan.assert(response["body"], "==", "{}\n")
