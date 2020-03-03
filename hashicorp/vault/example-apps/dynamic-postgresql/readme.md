# Dynamic Secret Creation [PostgreSQL]


Enable the database engine

```
kubectl -n vault-example exec -it vault-example-0 vault login
kubectl -n vault-example exec -it vault-example-0 vault secrets enable database

#map connection details to a role
kubectl -n vault-example exec -it vault-example-0 vault write database/config/my-postgresql-database \
    plugin_name=postgresql-database-plugin \
    allowed_roles="my-role" \
    connection_url="postgresql://{{username}}:{{password}}@localhost:5432/" \
    username="root" \
    password="root"

#create the role
kubectl -n vault-example exec -it vault-example-0  vault write database/roles/my-role \
    db_name=my-postgresql-database \
    creation_statements="CREATE ROLE \"{{name}}\" WITH LOGIN PASSWORD '{{password}}' VALID UNTIL '{{expiration}}'; \
        GRANT SELECT ON ALL TABLES IN SCHEMA public TO \"{{name}}\";" \
    default_ttl="1h" \
    max_ttl="24h"

```

```
#Create a role for our app

kubectl -n vault-example exec -it vault-example-0 vault write auth/kubernetes/role/basic-secret-role \
   bound_service_account_names=basic-secret \
   bound_service_account_namespaces=vault-example \
   policies=basic-secret-policy \
   ttl=1h
```

The above maps our Kubernetes service account, used by our pod, to a policy.
Now lets create the policy to map our service account to a bunch of secrets


```
kubectl -n vault-example exec -it vault-example-0 sh
cat <<EOF > /home/vault/app-policy.hcl
path "secret/basic-secret/*" {
  capabilities = ["read"]
}
EOF
vault policy write basic-secret-policy /home/vault/app-policy.hcl
exit
```

Now our service account for our pod can access all secrets under `secret/basic-secret/*`
Lets create some secrets.


```
kubectl -n vault-example exec -it vault-example-0 sh 
vault secrets enable -path=secret/ kv
vault kv put secret/basic-secret/helloworld username=dbuser password=sUp3rS3cUr3P@ssw0rd
exit
```

Lets deploy our app and see if it works:

```
kubectl -n vault-example apply -f ./hashicorp/vault/example-apps/basic-secret/deployment.yaml
```