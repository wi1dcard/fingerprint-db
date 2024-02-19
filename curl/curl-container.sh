#!/bin/sh

which yum && yum install -y curl || {
    apt-get update -y
    apt-get install -y curl
}

$(dirname $0)/curl.sh "$@"
