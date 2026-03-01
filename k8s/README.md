```bash
curl -sfL https://get.k3s.io | sh -s - --disable=traefik
```

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
helm install redpanda redpanda/redpanda \
  --version 25.3.2 \
  --namespace redpanda \
  --create-namespace \
  --set external.domain=customredpandadomain.local \
  --set statefulset.initContainers.setDataDirOwnership.enabled=true \
  --set statefulset.replicas=1 \
  --set statefulset.podAntiAffinity.type=soft \

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
docker save product-categories:dev | sudo k3s ctr images import -
firewalld to expose
```

```bash
kubectl create secret generic kafka-truststore \
  --from-file=truststore.jks \
  -n services


kubectl create secret generic kafka-ca \
--from-file=ca.crt \
-n services
```
