#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

echo "Building docker image based off of most recent local commits of pepecoingo and coreth"

PEPECOIN_REMOTE="git@github.com:memeticofficial/pepecoingo.git"
CORETH_REMOTE="git@github.com:memeticofficial/coreth.git"
DOCKERHUB_REPO="avaplatform/pepecoingo"

DOCKER="${DOCKER:-docker}"
SCRIPT_DIRPATH=$(cd $(dirname "${BASH_SOURCE[0]}") && pwd)

AVA_LABS_RELATIVE_PATH="src/github.com/memeticofficial"
EXISTING_GOPATH="$GOPATH"

export GOPATH="$SCRIPT_DIRPATH/.build_image_gopath"
WORKPREFIX="$GOPATH/src/github.com/memeticofficial"

# Clone the remotes and checkout the desired branch/commits
PEPECOIN_CLONE="$WORKPREFIX/pepecoingo"
CORETH_CLONE="$WORKPREFIX/coreth"

# Replace the WORKPREFIX directory
rm -rf "$WORKPREFIX"
mkdir -p "$WORKPREFIX"


PEPECOIN_COMMIT_HASH="$(git -C "$EXISTING_GOPATH/$AVA_LABS_RELATIVE_PATH/pepecoingo" rev-parse --short HEAD)"
CORETH_COMMIT_HASH="$(git -C "$EXISTING_GOPATH/$AVA_LABS_RELATIVE_PATH/coreth" rev-parse --short HEAD)"

git config --global credential.helper cache

git clone "$PEPECOIN_REMOTE" "$PEPECOIN_CLONE"
git -C "$PEPECOIN_CLONE" checkout "$PEPECOIN_COMMIT_HASH"

git clone "$CORETH_REMOTE" "$CORETH_CLONE"
git -C "$CORETH_CLONE" checkout "$CORETH_COMMIT_HASH"

CONCATENATED_HASHES="$PEPECOIN_COMMIT_HASH-$CORETH_COMMIT_HASH"

"$DOCKER" build -t "$DOCKERHUB_REPO:$CONCATENATED_HASHES" "$WORKPREFIX" -f "$SCRIPT_DIRPATH/local.Dockerfile"
