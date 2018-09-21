#!/bin/sh

set -e

echo "Test!"
which cleaner

if [[ $# -eq 0 ]] ; then
    cleaner --help
    exit 0
fi

# first arg is `-f` or `--some-option`
if [ "${1#-}" != "$1" ]; then
    set -- cleaner "$@"
fi

printf '%s\n' "$@"

exec "$@"