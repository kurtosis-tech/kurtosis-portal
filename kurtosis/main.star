TEST_ROOT_FOLDER = "github.com/kurtosis-tech/kurtosis-portal/kurtosis/integration_tests/"

PORTAL_IMAGE="kurtosistech/kurtosis-portal"

def run(plan, args):
    args_valid, test_name, portal_version = parse_args(args)
    if not args_valid:
        plan.verify(1, "==", 0)

    portal_image_with_version = "{image}:{version}".format(image=PORTAL_IMAGE, version=portal_version)
    test_to_run = import_module(TEST_ROOT_FOLDER + args["test"])
    test_to_run.test(plan, portal_image_with_version)


def parse_args(args):
    args_valid = False

    if "test" not in args or args["test"] == "":
        return args_valid, "", ""
    test_name = args["test"]

    if "version" not in args or args["version"] == "":
        return args_valid, "", ""
    version = args["version"]

    args_valid = True
    return args_valid, test_name, version
