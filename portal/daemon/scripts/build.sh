#!/usr/bin/env bash
# 2021-07-08 WATERMARK, DO NOT REMOVE - This script was generated from the Kurtosis Bash script template

set -euo pipefail   # Bash "strict mode"
script_dirpath="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
root_dirpath="$(dirname "${script_dirpath}")"

# ==================================================================================================
#                                             Constants
# ==================================================================================================

# Daemon code module
build_directory="${root_dirpath}/build"
binary_name="kurtosis_portal_daemon"

# ==================================================================================================
#                                             Main Logic
# ==================================================================================================

# Run tests in daemon golang module
cd "${root_dirpath}"
go test ./...

# Build the daemon
CGO_ENABLED=0 go build -o "${build_directory}/${binary_name}"
