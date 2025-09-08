```
kubectl apply -f k8s/shopiew.namespace.yaml
```
```
kubectl create secret generic app-secrets --from-env-file=k8s/.env -n services
```
```
kubectl apply -f https://github.com/kubernetes-sigs/gateway-api/releases/download/v1.3.0/standard-install.yaml
```
```
kubectl apply -f k8s/gateway/gateway-api-setup.yaml
```
```
kubectl create configmap opa-server-policy --from-file=opa/policies/server.rego -n opa
```
```
 kubectl port-forward -n envoy-gateway-system  svc/envoy-default-gateway-b7f3e5b1  8000:80
```