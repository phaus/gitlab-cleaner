#!/bin/sh

set -e

if [[ $# -eq 0 ]] ; then
    cleaner --help
    exit 0
fi

printf '%s\n' "$@"

exec "$@"