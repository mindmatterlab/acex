#!/usr/bin/env bash

env_file="$1"
template_file="$2"

GOPRO_ROOT=$(dirname "${BASH_SOURCE[0]}")/..

source "${GOPRO_ROOT}/scripts/lib/init.sh"

if [ $# -ne 2 ];then
    gopro::log::error "Usage: gen-config.sh manifests/env.local configs/gopro.service.tmpl"
    exit 1
fi

source "${env_file}"

declare -A envs

set +u
for env in $(sed -n 's/^[^#].*${\(.*\)}.*/\1/p' ${template_file})
do
    if [ -z "$(eval echo \$${env})" ];then
        gopro::log::error "environment variable '${env}' not set"
        missing=true
    fi
done

if [ "${missing}" ];then
    gopro::log::error 'You may run `source manifests/env.local` to set these environment'
    exit 1
fi

eval "cat << EOF
$(cat ${template_file})
EOF"
