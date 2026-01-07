# Ingress-NGINX Migration to Gateway API

<a href="https://youtu.be/7CaDvoKO9wo" title="nginx-ingress-migration"><img src="https://i.ytimg.com/vi/7CaDvoKO9wo/hqdefault.jpg" width="40%" alt="nginx-ingress-migration" /></a>


### Ingress Class

```shell
kubectl get ingress -A -o custom-columns="NAMESPACE:.metadata.namespace,NAME:.metadata.name,CLASS:.spec.ingressClassName"
```

### Ingress to Gateway Tool

```shell
wget -q "https://github.com/kubernetes-sigs/ingress2gateway/releases/download/v0.4.0/ingress2gateway_Linux_x86_64.tar.gz" -O /tmp/igtogw.tar.gz && \
  tar -C /tmp -xzf /tmp/igtogw.tar.gz && \
  mv /tmp/ingress2gateway /usr/local/bin/ingress2gateway && \
  chmod +x /usr/local/bin/ingress2gateway

./ingress2gateway print --providers=ingress-nginx
```

### Annotations Audit

To get a list of all NGINX Ingress Annotations we are using: 

```shell
kubectl get ingress -A -o json  | \
  jq -r '.items[].metadata.annotations | keys[]' | \
  grep "nginx.ingress" | \
  sort -u
```

Now we can generate our audit list by selecting the above annotated output:

```shell
apk add util-linux
audit_fields=(
  "NAMESPACE:.metadata.namespace"
  "NAME:.metadata.name"
  "HOSTS:.spec.rules[*].host"
  "MIGRATED:.metadata.labels.migrated"
  "REWRITE:.metadata.annotations.nginx\.ingress\.kubernetes\.io/rewrite-target"
  "SNIPPET:.metadata.annotations.nginx\.ingress\.kubernetes\.io/configuration-snippet"
  "AUTH:.metadata.annotations.nginx\.ingress\.kubernetes\.io/auth-type"
  "PathType:.spec.rules[*].http.paths[*].pathType"
  "TLS:.spec.tls[*].secretName"
)

IFS=$','
column_args="${audit_fields[*]}"

kubectl get ingress -A -o custom-columns="$column_args"
```

### Gateway API migration

| Feature | Ingress Nginx Annotation | Native Gateway API Resource / Filter | Example |
| :--- | :--- | :--- | :---|
| Basic Routing | `<none>` | `HTTPRoute` | [example: Route by Domain](#route-by-domain) |
| Path Rewrite | `rewrite-target` | `HTTPRoute` → `URLRewrite` filter | [example: URL Rewrite](#path-rewrite) |
| TLS termination | `tls.secretName` | `Gateway` → `Listeners` → `tls.certificateRefs` | [example: TLS termination](#tls-termination) |
| Redirects | `app-root`, `ssl-redirect` | `HTTPRoute` → `RequestRedirect` filter | [example: HTTP redirects](#http-redirects) |
| Header Mod | `add-header` | `HTTPRoute` → `RequestHeaderModifier` filter |
| Basic Auth | `auth-type: basic` | **Policy Attachment** (Targeting `HTTPRoute`) | [custom example: BasicAuth](#basic-auth-custom-feature) |
| Canary/Weight | `canary-weight` | `HTTPRoute` → `backendRefs.weight` |
| Mirroring | `mirror-target` | `HTTPRoute` → `RequestMirror` filter |
| Backend TLS | `backend-protocol: HTTPS` | `BackendTLSPolicy` (Standard Object) |
| Cross-Namespace | *(Standard Ingress)* | `ReferenceGrant` (Standard Object) |
| Request Timeout| `proxy-read-timeout` | `HTTPRoute` → `spec.rules.timeouts` |

### Examples

In these examples, we will run `port-forward` to test the old and new endpoints using `curl`. You want to ensure you get similar or same expected outcomes to ensure compatibility on your new endpoint. </br>

```shell
# port-forward to ingress
kubectl -n ingress-nginx port-forward svc/ingress-nginx-controller 80
# port-forward to gateway
kubectl -n traefik port-forward svc/traefik 80
```

#### Route by Domain

Before:
```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: python-route
spec:
  ingressClassName: nginx
  rules:
  - host: example-app-python.com
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: python-svc
            port:
              number: 5000
---
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: go-route
spec:
  ingressClassName: nginx
  rules:
  - host: example-app-go.com
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: go-svc
            port:
              number: 5000
```

This is what the [HTTPRoute:Route by Domain](../03-httproute-by-hostname.yaml) looks like. </br>

```shell 
kubectl apply -f kubernetes/gateway-api/03-httproute-by-hostname.yaml
```

Test:
```shell
curl example-app-python.com
Hello from Python!

curl example-app-go.com
Hello from Go
```

Mark as done
```
kubectl label ingress go-route python-route migrated=true
```

#### Path Rewrite

Before: 

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  annotations:
    nginx.ingress.kubernetes.io/rewrite-target: /$2
  name: python-route-bypath
  namespace: default
spec:
  ingressClassName: nginx
  rules:
  - host: example-app.com
    http:
      paths:
      - backend:
          service:
            name: python-svc
            port:
              number: 5000
        path: /api/python(/|$)(.*)
        pathType: ImplementationSpecific
---
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  annotations:
    nginx.ingress.kubernetes.io/rewrite-target: /$2
  name: go-route-bypath
  namespace: default
spec:
  ingressClassName: nginx
  rules:
  - host: example-app.com
    http:
      paths:
      - backend:
          service:
            name: go-svc
            port:
              number: 5000
        path: /api/go(/|$)(.*)
        pathType: ImplementationSpecific
```

After:
```yaml 
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: python-route-bypath
  namespace: default

spec:
  parentRefs:
  - name: gateway-api
    sectionName: http
    kind: Gateway

  hostnames:
  - example-app.com

  rules:
  - matches:
    - path: 
        type: PathPrefix
        value: /api/python
    filters:
    - type: URLRewrite
      urlRewrite:
        path:
          type: ReplacePrefixMatch
          replacePrefixMatch: /
    backendRefs:
    - name: python-svc
      port: 5000
      weight: 1
---

apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: go-route-bypath
  namespace: default

spec:
  parentRefs:
  - name: gateway-api
    sectionName: http
    kind: Gateway

  hostnames:
  - example-app.com

  rules:
  - matches:
    - path:
        type: PathPrefix
        value: /api/go 
    filters:
    - type: URLRewrite
      urlRewrite:
        path:
          type: ReplacePrefixMatch
          replacePrefixMatch: /
    backendRefs:
    - name: go-svc
      port: 5000
      weight: 1
```

Test:

```shell 
curl example-app.com/api/python
Hello from Python!

curl example-app.com/api/go
Hello from Go
```

Mark as done
```
kubectl label ingress go-route-bypath python-route-bypath  migrated=true
```

### TLS Termination

Before:
```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: go-route-tls
  namespace: default
  annotations:
    nginx.ingress.kubernetes.io/rewrite-target: /$2
spec:
  ingressClassName: nginx
  tls:
  - hosts:
    - example-app.com
    secretName: secret-tls
  rules:
  - host: example-app.com
    http:
      paths:
      - path: /api/go/tls(/|$)(.*)
        pathType: ImplementationSpecific
        backend:
          service:
            name: go-svc
            port:
              number: 5000
```

After:

```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: go-route-tls
  namespace: default

spec:
  parentRefs:
  - name: gateway-api
    sectionName: http
    kind: Gateway
  - name: gateway-api
    sectionName: https
    kind: Gateway

  hostnames:
  - example-app.com
  rules:
  - matches:
    - path:
        type: PathPrefix
        value: /api/go/tls
    filters:
    - type: URLRewrite
      urlRewrite:
        path:
          type: ReplacePrefixMatch
          replacePrefixMatch: /
    backendRefs:
    - name: go-svc
      port: 5000
      weight: 1
```

We can use our `port-forward` on port `443` to test the TLS in the browser 

Mark as done
```
kubectl label ingress go-route-tls python-route migrated=true
```

### HTTP Redirects

Before:
```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: python-route-tls
  namespace: default
  annotations:
    nginx.ingress.kubernetes.io/ssl-redirect: "true"
    nginx.ingress.kubernetes.io/rewrite-target: /$2
spec:
  ingressClassName: nginx
  tls:
  - hosts:
    - example-app.com
    secretName: secret-tls
  rules:
  - host: example-app.com
    http:
      paths:
      - path: /api/python/tls(/|$)(.*)
        pathType: ImplementationSpecific
        backend:
          service:
            name: python-svc
            port:
              number: 5000
```

After:
```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: python-route-tls-redirect
  namespace: default
spec:
  parentRefs:
  - name: gateway-api
    sectionName: http
  hostnames:
  - example-app.com
  rules:
  - matches:
    - path:
        type: PathPrefix
        value: /api/python/tls
    filters:
    - type: RequestRedirect
      requestRedirect:
        scheme: https
        statusCode: 301
---
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: python-route-tls
  namespace: default

spec:
  parentRefs:
  - name: gateway-api
    sectionName: https
    kind: Gateway

  hostnames:
  - example-app.com
  rules:
  - matches:
    - path:
        type: PathPrefix
        value: /api/python/tls
    filters:
    - type: URLRewrite
      urlRewrite:
        path:
          type: ReplacePrefixMatch
          replacePrefixMatch: /
    backendRefs:
    - name: python-svc
      port: 5000
      weight: 1
```

Test:

```shell
# before
curl http://example-app.com/api/python/tls/
<html>
<head><title>308 Permanent Redirect</title></head>
<body>
<center><h1>308 Permanent Redirect</h1></center>
<hr><center>nginx</center>
</body>
</html>

# after 
curl http://example-app.com/api/python/tls/ -v
< HTTP/1.1 301 Moved Permanently
< Location: https://example-app.com/api/python/tls/
```

Mark as done
```
kubectl label ingress python-route-tls migrated=true
```


### Basic Auth (Custom Feature)

Create a Basic Auth secret:

```shell
apk add apache2-utils
htpasswd -c auth python
kubectl create secret generic python-route-auth-secret --from-file=auth

```
Before: 

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: python-route-auth
  annotations:
    nginx.ingress.kubernetes.io/auth-type: basic
    nginx.ingress.kubernetes.io/auth-secret: python-route-auth-secret
    nginx.ingress.kubernetes.io/auth-secret-type: auth-file
    nginx.ingress.kubernetes.io/rewrite-target: /$2
spec:
  ingressClassName: nginx
  rules:
  - host: example-app.com
    http:
      paths:
      - path: /api/python/login(/|$)(.*)
        pathType: ImplementationSpecific
        backend:
          service:
            name: python-svc
            port:
              number: 5000
```

After:

```yaml
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: python-route-auth
  namespace: default
spec:
  parentRefs:
  - name: gateway-api
    sectionName: http
    kind: Gateway
  hostnames:
  - "example-app.com"
  rules:
  - matches:
    - path:
        type: PathPrefix
        value: /api/python/login
    filters:
    - type: ExtensionRef
      extensionRef:
        group: traefik.io
        kind: Middleware
        name: basic-auth
    - type: URLRewrite
      urlRewrite:
        path:
          type: ReplacePrefixMatch
          replacePrefixMatch: /
    backendRefs:
    - name: python-svc
      port: 5000
---
apiVersion: traefik.io/v1alpha1
kind: Middleware
metadata:
  name: basic-auth
spec:
  basicAuth:
    secret: python-route-auth-secret
```

Test:

```shell
#Before:
curl http://example-app.com/api/python/login
<html>
<head><title>401 Authorization Required</title></head>
<body>
<center><h1>401 Authorization Required</h1></center>
<hr><center>nginx</center>
</body>
</html>

#After:
curl http://public.my-services.com/path-a/login
401 Unauthorized
```

Mark as done
```
kubectl label ingress python-route-auth migrated=true
```