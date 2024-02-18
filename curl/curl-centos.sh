#!/bin/sh

yum install -y curl

curl --version
curl -o fingerprint.json https://localhost:8443/json --insecure "$@"
