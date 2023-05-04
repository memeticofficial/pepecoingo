#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# Directory above this script
PEPECOIN_PATH=$( cd "$( dirname "${BASH_SOURCE[0]}" )"; cd .. && pwd )

# Load the constants
source "$PEPECOIN_PATH"/scripts/constants.sh

# WARNING: this will use the most recent commit even if there are un-committed changes present
full_commit_hash="$(git --git-dir="$PEPECOIN_PATH/.git" rev-parse HEAD)"
commit_hash="${full_commit_hash::8}"

echo "Building Docker Image with tags: $pepecoingo_dockerhub_repo:$commit_hash , $pepecoingo_dockerhub_repo:$current_branch"
docker build -t "$pepecoingo_dockerhub_repo:$commit_hash" \
        -t "$pepecoingo_dockerhub_repo:$current_branch" "$PEPECOIN_PATH" -f "$PEPECOIN_PATH/Dockerfile"
