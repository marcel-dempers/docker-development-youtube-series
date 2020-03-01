# Create an App policy

```
kubectl -n vault-example exec -it vault-example-0 sh

cat <<EOF > /home/vault/app-policy.hcl
path "secret*" {
  capabilities = ["read"]
}
EOF

vault login
vault policy write app /home/vault/app-policy.hcl

```