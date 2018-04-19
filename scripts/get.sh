#!/bin/bash

curl $CI_PROJECT_URL/registry/repository/$CI_PROJECT_ID/tags?format=json -u gitlab-ci-token:$CI_JOB_TOKEN \
-H 'X-Requested-With: XMLHttpRequest' \
-H 'Connection: keep-alive' \
-H 'X-Requested-With: XMLHttpRequest' \
-H 'Accept: application/json, text/plain, */*' --compressed