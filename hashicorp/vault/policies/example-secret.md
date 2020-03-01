# Create example secret

```
kubectl -n vault-example exec -it vault-example-0 sh

vault login

vault secrets enable -path=secret/ kv
vault kv put secret/helloworld username=foobaruser password=foobarbazpass

```