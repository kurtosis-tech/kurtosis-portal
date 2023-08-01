portal_launcher = import_module("github.com/kurtosis-tech/kurtosis-portal/kurtosis/portal_launcher/portal_launcher.star")

def test(plan, portal_image):
    service_name = "portal-server"
    portal_and_proxy = portal_launcher.launch_portal(plan, portal_image, service_name, server_only=True)
    proxy = portal_and_proxy["proxy"]

    response = plan.request(
        service_name=proxy.name,
        recipe=PostHttpRequestRecipe(
            port_id=portal_launcher.PORTAL_PROXY_SINGLE_PORT_ID,
            endpoint="/server/grpc/get-remote-endpoints",
            content_type="application/json",
            body="{}"
        )
    )

    plan.assert(response["code"], "==", 200)
    expected_body = "{\n \"remoteEndpoints\": []\n}"
    plan.assert(response["body"], "==", expected_body)