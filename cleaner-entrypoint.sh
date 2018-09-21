#!/bin/sh

set -e

printf 'before %s\n' "$@"

which sh
which cleaner

if [[ $# -eq 0 ]] ; then
    echo "exit"
    cleaner --help
    exit 0
fi

# first arg is `-f` or `--some-option`
if [ "${1#-}" != "$1" ]; then
    echo "-"
    set -- cleaner "$@"
fi

printf 'after %s\n' "$@"

exec "$@"