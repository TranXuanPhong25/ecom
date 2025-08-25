```
kubectl apply -f k8s/shopiew.namespace.yaml
```
```
 kubectl config set-context --current --namespace=shopiew 
```
```
kubectl apply -f https://github.com/kubernetes-sigs/gateway-api/releases/download/v1.3.0/standard-install.yaml
```
```
kubectl apply -f k8s/gateway/gateway-api-setup.yaml
```
