#!/bin/bash

docker buildx create --name myarm64
docker buildx use myarm64
docker buildx ls
docker buildx inspect --bootstrap
docker buildx build --platform linux/arm64,linux/amd64 -t mario21ic/goruntime:multi -f Dockerfile.multi ./ --push

