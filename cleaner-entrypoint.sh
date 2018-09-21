#!/bin/sh

set -e

printf '%s\n' "$@"

if [[ $# -eq 0 ]] ; then
    cleaner --help
    exit 0
fi

# first arg is `-f` or `--some-option`
if [ "${1#-}" != "$1" ]; then
    set -- cleaner "$@"
fi

for arg 
do
    shift
    [ "$arg" = "sh" -o "$arg" = "-c" ] && continue
    set -- "$@" "$arg"
done

printf '%s\n' "$@"

exec "$@"