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

# if our command is a valid Docker subcommand, let's invoke it through Docker instead
# (this allows for "docker run docker ps", etc)
if cleaner help "$1" > /dev/null 2>&1; then
    set -- cleaner "$@"
fi

printf '%s\n' "$@"

exec "$@"