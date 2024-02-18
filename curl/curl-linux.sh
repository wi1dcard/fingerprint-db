#!/bin/sh

apt-get update -y
apt-get install -y curl

curl --version
curl -o fingerprint.json https://localhost:8443/json --insecure "$@"
