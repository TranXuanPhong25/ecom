#!/bin/bash
DOCKER_USERNAME=rengumin
SERVICE_NAME=carts-svc
TAG=1.0
SERVICE_PATH=.
docker build -t $DOCKER_USERNAME/$SERVICE_NAME:$TAG -t $DOCKER_USERNAME/$SERVICE_NAME:latest -f $SERVICE_PATH/Dockerfile $SERVICE_PATH
kubectl rollout restart deployment carts-service -n services
