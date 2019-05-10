#!/usr/bin/env bash

set -ex
export TAG=v0.2.11

docker build -t quay.io/kubermatic/addons:${TAG} .
docker push quay.io/kubermatic/addons:${TAG}
