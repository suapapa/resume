#!/bin/bash

set -e

IMAGE_TAG=gcr.io/homin-dev/resume:$1 
docker buildx build --platform linux/amd64 -t $IMAGE_TAG .
docker push $IMAGE_TAG

IMAGE_TAG_LATEST=gcr.io/homin-dev/resume:latest 
docker rmi $IMAGE_TAG_LATEST || true
docker tag $IMAGE_TAG $IMAGE_TAG_LATEST
docker push $IMAGE_TAG_LATEST

git tag -a $1 -m "add tag for $1"
git push --tags

