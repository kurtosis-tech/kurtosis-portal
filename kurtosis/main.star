TEST_ROOT_FOLDER = "github.com/kurtosis-tech/kurtosis-portal/kurtosis/integration_tests/"

def run(plan, args):
    if "test" not in args or args["test"] == "":
        return

    test_to_run = import_module(TEST_ROOT_FOLDER + args["test"])
    test_to_run.test(plan)
