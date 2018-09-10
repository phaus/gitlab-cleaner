#!/bin/bash

echo "accessing:"

#curl $CI_PROJECT_URL/registry/repository/$CI_PROJECT_ID/tags?format=json \
#-H "Private-Token: $PRIVATE_ACCESS_TOKEN" \
#-H "X-Requested-With: XMLHttpRequest" \
#-H "Connection: keep-alive" \
#-H "X-Requested-With: XMLHttpRequest" \
#-H "Accept: application/json, text/plain, */*" --compressed

echo ""
echo "via CI_REGISTRY_USER"
curl -s $CI_PROJECT_URL/container_registry.json \
-u $CI_REGISTRY_USER:$CI_REGISTRY_PASSWORD \
-H 'accept: application/json, text/plain, */*' | jq

echo ""
echo "via PRIVATE_ACCESS_TOKEN"
curl -s $CI_PROJECT_URL/container_registry.json \
-H "Private-Token: $PRIVATE_ACCESS_TOKEN" \
-H 'accept: application/json, text/plain, */*' | jq

echo ""
echo "via PRIVATE_ACCESS_TOKEN"
curl -s $CI_PROJECT_URL/registry/repository/$CI_PROJECT_ID/tags?format=json \
-D /tmp/head \
-H "Private-Token: $PRIVATE_ACCESS_TOKEN" \
-H 'accept: application/json, text/plain, */*' | jq

echo "headers: $(cat /tmp/head)"