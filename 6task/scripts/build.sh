#!/bin/bash

set -e

TAG="local"

while getopts "t:" opt; do
    case "$opt" in
        t)
            TAG="$OPTARG"
            ;;
        *)
            echo "Usage: $0 [-t tag]"
            exit 1
            ;;
    esac
done

echo "Building web-app:$TAG"

docker build -t "web-app:$TAG" .

echo "Loading image into Minikube"

minikube image load "web-app:$TAG"

echo "Build completed"