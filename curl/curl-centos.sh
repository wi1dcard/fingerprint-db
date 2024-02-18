#!/bin/sh

yum install -y curl

curl --version
curl -o fingerprint.json https://host.docker.internal:8443/json --insecure "$@"
