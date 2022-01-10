# Dynamic Secret Creation [PostgreSQL]

Deploy our test database


```
kubectl create ns postgres
kubectl -n postgres apply -f ./hashicorp/vault/example-apps/dynamic-postgresql/postgres.yaml
kubectl -n postgres apply -f ./hashicorp/vault/example-apps/dynamic-postgresql/pgadmin.yaml
kubectl -n postgres get pods

kubectl -n postgres exec -it <podname> bash
psql --username=postgresadmin postgresdb
```

Enable the database engine

```
kubectl -n vault-example exec -it vault-example-0 vault login
kubectl -n vault-example exec -it vault-example-0 vault secrets enable database
```

## Configure DB Credential creation

```
kubectl -n vault-example exec -it vault-example-0 sh 

vault write database/config/postgresdb \
    plugin_name=postgresql-database-plugin \
    allowed_roles="sql-role" \
    connection_url="postgresql://{{username}}:{{password}}@postgres.postgres:5432/postgresdb?sslmode=disable" \
    username="postgresadmin" \
    password="admin123"

 vault write database/roles/sql-role \
    db_name=postgresdb \
    creation_statements="CREATE ROLE \"{{name}}\" WITH LOGIN PASSWORD '{{password}}' VALID UNTIL '{{expiration}}'; \
        GRANT SELECT ON ALL TABLES IN SCHEMA public TO \"{{name}}\";" \
    default_ttl="1h" \
    max_ttl="24h"

#test 
vault read database/creds/sql-role

```

## Example Application

Create a policy to control access to secrets

```
kubectl -n vault-example exec -it vault-example-0 sh

cat <<EOF > /home/vault/postgres-app-policy.hcl
path "database/creds/sql-role" {
  capabilities = ["read"]
}
EOF

vault policy write postgres-app-policy /home/vault/postgres-app-policy.hcl

```


Bind our role to a service account for our application


```
kubectl -n vault-example exec -it vault-example-0 sh

vault write auth/kubernetes/role/sql-role \
   bound_service_account_names=dynamic-postgres \
   bound_service_account_namespaces=vault-example \
   policies=postgres-app-policy \
   ttl=1h

```

kubectl -n vault-example apply -f .\hashicorp\vault\example-apps\dynamic-postgresql\deployment.yaml