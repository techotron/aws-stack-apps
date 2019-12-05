# aws-stack-apps

## localstack-multi

### Deploy

```bash
helm upgrade -i -f ./localstack-multi/values.yaml localstack-multi ./localstack-multi --namespace ops
```

### Delete

```bash
helm uninstall localstack-multi
```