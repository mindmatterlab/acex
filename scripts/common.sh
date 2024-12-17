#!/usr/bin/env bash

# Common utilities, variables and checks for all build scripts.
set -eEuo pipefail

# Unset CDPATH, having it set messes up with script import paths
unset CDPATH

# This will canonicalize the path
GOPRO_ROOT=$(cd "$(dirname "${BASH_SOURCE[0]}")"/.. && pwd -P)

source "${GOPRO_ROOT}/scripts/lib/init.sh"
