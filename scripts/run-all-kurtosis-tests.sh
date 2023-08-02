#!/usr/bin/env bash
# 2021-07-08 WATERMARK, DO NOT REMOVE - This script was generated from the Kurtosis Bash script template

set -euo pipefail   # Bash "strict mode"
script_dirpath="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
root_dirpath="$(dirname "${script_dirpath}")"



# ==================================================================================================
#                                             Constants
# ==================================================================================================
KURTOSIS_PKG_ROOT="kurtosis/"
TEST_SOURCE_ROOT="${KURTOSIS_PKG_ROOT}integration_tests/"
TEST_FILE_PATTERN="*_test.star"

# ==================================================================================================
#                                             Main Logic
# ==================================================================================================
version="$(./daemon/scripts/get-docker-tag.sh)"
for test_file in ${TEST_SOURCE_ROOT}${TEST_FILE_PATTERN}; do
    test_file_name=$(basename ${test_file})
    enclave_name="${test_file_name%.*}"
    echo "Running test for test file ${test_file_name}"
    kurtosis run "${KURTOSIS_PKG_ROOT}" --enclave "${enclave_name}" "{\"version\": \"${version}\", \"test\": \"${test_file_name}\"}"
    if [ $? -eq 0 ] 
    then
        kurtosis enclave rm -f "${enclave_name}"
    fi
done
