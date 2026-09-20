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

echo "Deploying web-app:$TAG"

helm upgrade --install web-app ./helm/web-app \
    --namespace konstantin20092026 \
    --set image.tag="$TAG"

echo "Deployment completed"