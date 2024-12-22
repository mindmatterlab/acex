#!/usr/bin/env bash

env_file="$1"
template_file="$2"

ACEX_ROOT=$(dirname "${BASH_SOURCE[0]}")/..

source "${ACEX_ROOT}/scripts/lib/init.sh"

if [ $# -ne 2 ];then
    acex::log::error "Usage: gen-config.sh manifests/env.local configs/acex.service.tmpl"
    exit 1
fi

source "${env_file}"

declare -A envs

set +u
for env in $(sed -n 's/^[^#].*${\(.*\)}.*/\1/p' ${template_file})
do
    if [ -z "$(eval echo \$${env})" ];then
        acex::log::error "environment variable '${env}' not set"
        missing=true
    fi
done

if [ "${missing}" ];then
    acex::log::error 'You may run `source manifests/env.local` to set these environment'
    exit 1
fi

eval "cat << EOF
$(cat ${template_file})
EOF"
