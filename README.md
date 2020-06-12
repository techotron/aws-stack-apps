# aws-stack-apps

## localstack-multi

Localstack deployment with multiple services enabled

URL: SERVICENAME.NAMESPACE.svc.cluster.local

## Intention

Run a bunch of microservices in a k8s cluster so I can learn about k8s, code, service mesh etc.

### Deploy

```bash
helm upgrade -i -f ./localstack-multi/values.yaml localstack-multi ./localstack-multi --namespace awsstackapps
```

### Delete

```bash
helm uninstall localstack-multi
```

## fe-api

Front End API written in Go

### Deploy

```bash
helm upgrade -i -f ./fe-api/values.yaml fe-api ./fe-api --namespace awsstackapps
```

### Delete

```bash
helm uninstall fe-api
```