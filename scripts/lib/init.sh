#!/usr/bin/env bash

set -o errexit
set +o nounset
set -o pipefail

# Short-circuit if init.sh has already been sourced
[[ $(type -t gopro::init::loaded) == function ]] && return 0

# Unset CDPATH so that path interpolation can work correctly
# https://github.com/minerrnetes/minerrnetes/issues/52255
unset CDPATH

# The root of the build/dist directory
GOPRO_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd -P)"

source "${GOPRO_ROOT}/scripts/lib/logging.sh"

# Marker function to indicate init.sh has been fully sourced
gopro::init::loaded() {
  return 0
}
