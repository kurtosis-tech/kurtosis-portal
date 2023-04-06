#!/usr/bin/env bash
# 2021-07-08 WATERMARK, DO NOT REMOVE - This script was generated from the Kurtosis Bash script template

set -euo pipefail   # Bash "strict mode"
script_dirpath="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
root_dirpath="$(dirname "${script_dirpath}")"

# ==================================================================================================
#                                             Constants
# ==================================================================================================

source "${script_dirpath}/_constants.env"
DEFAULT_SKIP_DOCKER_IMAGE_BUILDING=false

# Daemon code module
build_directory="${root_dirpath}/build"
binary_name="kurtosis-portal"

docker_image_builder="docker-image-builder.sh"
get_docker_tag="get-docker-tag.sh"

# ==================================================================================================
#                                             Download dependencies
# ==================================================================================================
docker_image_builder_abs_path="${script_dirpath}/${docker_image_builder}"
if [ ! -f "${docker_image_builder_abs_path}" ]; then
  echo "${docker_image_builder} script not present. Downloading it to ${docker_image_builder_abs_path}"
  curl --fail -XGET -o "${docker_image_builder_abs_path}" https://raw.githubusercontent.com/kurtosis-tech/kurtosis/main/scripts/docker-image-builder.sh
  chmod +x "${docker_image_builder_abs_path}"
fi

get_docker_tag_abs_path="${script_dirpath}/${get_docker_tag}"
if [ ! -f "${get_docker_tag_abs_path}" ]; then
  echo "${get_docker_tag} script not present. Downloading it to ${get_docker_tag_abs_path}"
  curl --fail -XGET -o "${get_docker_tag_abs_path}" https://raw.githubusercontent.com/kurtosis-tech/kurtosis/main/scripts/get-docker-tag.sh
  chmod +x "${get_docker_tag_abs_path}"
fi

# ==================================================================================================
#                                             Main Logic
# ==================================================================================================
# Parse arguments
skip_docker_image_building="${1:-"${DEFAULT_SKIP_DOCKER_IMAGE_BUILDING}"}"
if [ "${skip_docker_image_building}" != "true" ] && [ "${skip_docker_image_building}" != "false" ]; then
    echo "Error: Invalid skip-docker-image-building arg '${skip_docker_image_building}'" >&2
fi


# Run tests in daemon golang module
cd "${root_dirpath}"
go test ./...

# Build the daemon
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o "${build_directory}/${binary_name}"

# Build docker image
if "${skip_docker_image_building}"; then
  echo "Not building docker image as requested"
  exit 0
fi

if ! docker_tag="$(./scripts/get-docker-tag.sh)"; then
    echo "Error: Couldn't get the Docker image tag" >&2
    exit 1
fi

dockerfile_filepath="${root_dirpath}/Dockerfile"
image_name="${IMAGE_ORG_AND_REPO}:${docker_tag}"
load_not_push_image=false
docker_build_script_cmd="${script_dirpath}/docker-image-builder.sh ${load_not_push_image} ${dockerfile_filepath} ${image_name}"
if ! eval "${docker_build_script_cmd}"; then
  echo "Error: Docker build failed" >&2
fi
