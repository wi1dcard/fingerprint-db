#!/bin/sh

which yum && yum install -y curl || {
    apt-get update -y
    apt-get install -y curl
}

curl --version
curl -o fingerprint.json https://host.docker.internal:8443/json --insecure "$@"
