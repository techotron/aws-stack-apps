# aws-stack-apps

## localstack-multi

Localstack deployment with multiple services enabled

URL: SERVICENAME.NAMESPACE.svc.cluster.local

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