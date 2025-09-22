```
kubectl apply -f k8s/shopiew.namespace.yaml
```
```
kubectl create secret generic app-secrets --from-env-file=k8s/.env -n services
```
```
helm install eg oci://docker.io/envoyproxy/gateway-helm --version v1.5.1 -n envoy-gateway-system --create-namespace
```
```
kubectl create configmap opa-server-policy --from-file=opa/policies/server.rego -n opa
```
```
 kubectl port-forward -n envoy-gateway-system  svc/envoy-default-gateway-b7f3e5b1  8000:80
```