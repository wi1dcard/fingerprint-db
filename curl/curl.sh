#!/bin/sh

curl --version
curl -o fingerprint.json https://localhost:8443/json --insecure --http2
