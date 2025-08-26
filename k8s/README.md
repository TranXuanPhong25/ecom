```
kubectl apply -f k8s/shopiew.namespace.yaml
```
```
 kubectl config set-context --current --namespace=shopiew 
```
```
helm repo add kong https://charts.konghq.com
helm repo update
helm install kong kong/ingress -n kong --create-namespace
```
