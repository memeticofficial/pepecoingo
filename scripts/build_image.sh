#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# Pepecoingo root folder
PEPECOIN_PATH=$( cd "$( dirname "${BASH_SOURCE[0]}" )"; cd .. && pwd )
# Load the constants
source "$PEPECOIN_PATH"/scripts/constants.sh

# build_image_from_remote.sh is deprecated
source "$PEPECOIN_PATH"/scripts/build_local_image.sh
