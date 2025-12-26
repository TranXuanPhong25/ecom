#!/bin/bash
kubectl apply -f shopiew.namespace.yaml
# helm install eg oci://docker.io/envoyproxy/gateway-helm --version v1.5.1 -n envoy-gateway-system --create-namespace
kubectl create secret generic app-secrets --from-env-file=.env -n services
kubectl create configmap opa-server-policy --from-file=../opa/policies/server.rego -n opa
kubectl apply -f storages/
kubectl apply -f services/
kubectl apply -f minio/
kubectl apply -f opa/
kubectl apply -f envoy-gateway/