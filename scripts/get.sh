#!/bin/bash

curl $CI_REGISTRY -u gitlab-ci-token:$CI_JOB_TOKEN -H 'X-Requested-With: XMLHttpRequest' -H 'Connection: keep-alive' 