#!/usr/bin/env bash

set -o errexit

echo "Resolving modules in $(pwd)"

function find_modules() {
  find . -not \( \
    \( \
    -path './deploy' \
    -o -path './.git' \
    -o -path '*/third_party/*' \
    \) -prune \
    \) -name 'go.mod' -print0 | xargs -0 -I {} dirname {}
}

all_modules=$(find_modules)

PATHS=""

for mod in $all_modules; do
  PATHS+=$(printf '{"workdir":"%s"},' "${mod}")
done

echo "::set-output name=matrix::{\"include\":[${PATHS%?}]}"
