#!/bin/bash
minikube start && minikube dashboard &

(
  sleep 30
  kubectl port-forward -n envoy-gateway-system svc/envoy-default-gateway-b7f3e5b1 8000:80
) &

wait
