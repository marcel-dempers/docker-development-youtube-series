```
docker run -it --rm -v ${PWD}:/work -w /work debian:buster bash

apt-get update && apt-get install -y curl &&
curl https://pkg.cfssl.org/R1.2/cfssl_linux-amd64 -o /usr/local/bin/cfssl && \
curl https://pkg.cfssl.org/R1.2/cfssljson_linux-amd64 -o /usr/local/bin/cfssljson && \
chmod +x /usr/local/bin/cfssl && \
chmod +x /usr/local/bin/cfssljson

#generate ca in /tmp
cfssl gencert -initca ./tls/ca-csr.json | cfssljson -bare /tmp/ca

#generate certificate in /tmp
cfssl gencert \
  -ca=/tmp/ca.pem \
  -ca-key=/tmp/ca-key.pem \
  -config=./tls/ca-config.json \
  -hostname="example-webhook,example-webhook.default.svc.cluster.local,example-webhook.default.svc,localhost,127.0.0.1" \
  -profile=default \
  ./tls/ca-csr.json | cfssljson -bare /tmp/webhook-server

#make a secret
cat <<EOF > ./tls/webhook-server-tls.yaml
apiVersion: v1
kind: Secret
metadata:
  name: example-webhook-tls
type: Opaque
data:
  tls.crt: $(cat /tmp/webhook-server.pem | base64 | tr -d '\n')
  tls.key: $(cat /tmp/webhook-server-key.pem | base64 | tr -d '\n') 
EOF

#generate CA Bundle + inject into template
ca_pem_b64="$(openssl base64 -A <"/tmp/ca.pem")"

sed -e 's@${CA_PEM_B64}@'"$ca_pem_b64"'@g' <"webhook-template.yaml" \
    > webhook.yaml
```