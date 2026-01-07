```bash
kubectl apply -f k8s/shopiew.namespace.yaml
```

```bash
kubectl create secret generic app-secrets --from-env-file=k8s/.env -n services
```

```bash
helm install eg oci://docker.io/envoyproxy/gateway-helm --version v1.5.1 -n envoy-gateway-system --create-namespace
```

```bash
kubectl create -f 'https://strimzi.io/install/latest?namespace=kafka' -n kafka
kubectl create -f https://download.elastic.co/downloads/eck/3.2.0/crds.yaml
kubectl apply -f https://download.elastic.co/downloads/eck/3.2.0/operator.yaml


```

```bash
kubectl create configmap opa-server-policy --from-file=opa/policies/server.rego -n opa
```

```bash
 kubectl delete  configmap opa-server-policy  -n opa
```

```bash
 kubectl port-forward -n envoy-gateway-system  svc/envoy-default-gateway-b7f3e5b1  8000:80
```

```bash
eval $(minikube docker-env)
```
