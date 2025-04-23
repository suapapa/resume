#!/bin/bash

set -e

CR=icn.vultrcr.com/homincr1
IMAGE_TAG=$CR/resume:$1 
docker buildx build --platform linux/amd64 -t $IMAGE_TAG .
docker push $IMAGE_TAG

if [ "$1" != "latest" ]; then
    IMAGE_TAG_LATEST=$CR/resume:latest 
    docker rmi $IMAGE_TAG_LATEST || true
    docker tag $IMAGE_TAG $IMAGE_TAG_LATEST
    docker push $IMAGE_TAG_LATEST
fi

# git tag -a $1 -m "add tag for $1"
# git push --tags

