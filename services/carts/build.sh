#!/bin/bash
# Build and push Order Placement service to Docker Hub
SERVICE_NAME="carts-svc"
DOCKER_USERNAME="rengumin"
VERSION="1.0"

echo "Building $SERVICE_NAME..."

# Build the Docker image
docker build -t $DOCKER_USERNAME/$SERVICE_NAME:$VERSION . --target=release
docker save $DOCKER_USERNAME/$SERVICE_NAME:$VERSION | sudo k3s ctr images import -

# # Push to Docker Hub
# echo "Pushing $SERVICE_NAME to Docker Hub..."
# docker push $DOCKER_USERNAME/$SERVICE_NAME:$VERSION
kubectl rollout restart deployment carts-service -n services

echo "$SERVICE_NAME built and pushed successfully!"
