#!/bin/bash
kubectl delete  configmap opa-server-policy  -n opa
kubectl create configmap opa-server-policy --from-file=../opa/policies/server.rego -n opa
kubectl rollout restart deployment opa-server -n opa
