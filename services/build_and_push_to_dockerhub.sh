#!/bin/bash
set -e

# Configuration
DOCKER_USERNAME="rengumin"
TAG="1.0"

# List of services to build
# Format: name:type:path
SERVICES=(
    "shops-svc:go:./shops"
    "users-svc:go:./users"
    "upload-svc:go:./upload-service"
    "jwt-service:go:./jwt-service"
    "promotions-svc:go:./promotions"

#    "chatbots-svc:spring:./orders"
    "carts-svc:go:./carts"
    "auth-svc:go:./auth"
    "products-svc:spring:./products"
    "product-categories-svc:spring:./product-categories"
)
# minikube image load rengumin/shops-svc:1.0 rengumin/users-svc:1.0 rengumin/upload-svc:1.0 rengumin/jwt-service:1.0 rengumin/carts-svc:1.0 rengumin/auth-svc:1.0 rengumin/products-svc:1.0 rengumin/product-categories-svc:1.0
echo "Building and pushing Docker images with tag: $TAG"

# Login to Docker Hub
echo "Please login to Docker Hub:"
docker login

# Loop through each service and build/push
for service in "${SERVICES[@]}"; do
    # Parse service info
    IFS=':' read -r SERVICE_NAME SERVICE_TYPE SERVICE_PATH <<< "$service"

    echo "===== Processing $SERVICE_NAME ($SERVICE_TYPE) ====="
    if [ "$SERVICE_TYPE" = "go" ]; then
        echo "Building Go service: $SERVICE_NAME from $SERVICE_PATH"

        # Build and push Go service using existing Dockerfile
        docker build -t $DOCKER_USERNAME/$SERVICE_NAME:$TAG -t $DOCKER_USERNAME/$SERVICE_NAME:latest -f $SERVICE_PATH/Dockerfile $SERVICE_PATH
        docker push $DOCKER_USERNAME/$SERVICE_NAME:$TAG
        docker push $DOCKER_USERNAME/$SERVICE_NAME:latest

    elif [ "$SERVICE_TYPE" = "spring" ]; then
        echo "Building Spring service: $SERVICE_NAME from $SERVICE_PATH"

        # Build the Spring application
        cd $SERVICE_PATH

        # Build and push Spring service using existing Dockerfile
        docker build -t $DOCKER_USERNAME/$SERVICE_NAME:$TAG -t $DOCKER_USERNAME/$SERVICE_NAME:latest . --target=release
        docker push $DOCKER_USERNAME/$SERVICE_NAME:$TAG
        docker push $DOCKER_USERNAME/$SERVICE_NAME:latest

        cd - > /dev/null
    else
        echo "Unknown service type: $SERVICE_TYPE for $SERVICE_NAME. Skipping."
    fi

    echo "===== Completed $SERVICE_NAME ====="
    echo
done

echo "Successfully built and pushed all Docker images."
