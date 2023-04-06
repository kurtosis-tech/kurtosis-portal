DEFAULT_VERSION = "latest"

PORTAL_IMAGE="kurtosistech/kurtosis-portal"

def run(plan, args):
    version = parse_args(args)

    # Very simple for now, just spin up a Portal container. Will be extended in a FLUP
    plan.add_service(
        name="portal",
        config=ServiceConfig(
            image="{image}:{version}".format(image=PORTAL_IMAGE, version=version),
        )
    )

def parse_args(args):
    if "version" not in args or args["version"] == "":
        version = "latest"
    else:
        version = args["version"]
    return version
