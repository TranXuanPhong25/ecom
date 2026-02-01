#!/bin/bash

#!/bin/bash
DOCKER_USERNAME=rengumin
SERVICE_NAME=jwt-svc
VERSION="1.0"

echo "Building $SERVICE_NAME..."

# Build the Docker image
docker build -t $DOCKER_USERNAME/$SERVICE_NAME:$VERSION . --target=release
docker save $DOCKER_USERNAME/$SERVICE_NAME:$VERSION | sudo k3s ctr images import -

# # Push to Docker Hub
# echo "Pushing $SERVICE_NAME to Docker Hub..."
# docker push $DOCKER_USERNAME/$SERVICE_NAME:$VERSION
kubectl rollout restart deployment jwt-service -n services

echo "$SERVICE_NAME built and pushed successfully!"
