#!/bin/bash

# Build and push Order Placement service to Docker Hub
SERVICE_NAME="order-placement"
DOCKER_USERNAME="your-dockerhub-username"
VERSION="latest"

echo "Building $SERVICE_NAME..."
cd "$(dirname "$0")/$SERVICE_NAME" || exit

# Build the Docker image
docker build -t $DOCKER_USERNAME/$SERVICE_NAME:$VERSION .

# Push to Docker Hub
echo "Pushing $SERVICE_NAME to Docker Hub..."
docker push $DOCKER_USERNAME/$SERVICE_NAME:$VERSION

echo "$SERVICE_NAME built and pushed successfully!"
