# EXAMPLES

## Install as a DaemonSet

Default install is using a `Deployment` but it's possible to use `DaemonSet`

```yaml
deployment:
  kind: DaemonSet
```

## Configure Traefik Pod parameters

### Extending /etc/hosts records

In some specific cases, you'll need to add extra records to the `/etc/hosts` file for the Traefik containers.
You can configure it using [hostAliases](https://kubernetes.io/docs/tasks/network/customize-hosts-file-for-pods/):

```yaml
deployment:
  hostAliases:
  - ip: "127.0.0.1" # this is an example
    hostnames:
     - "foo.local"
     - "bar.local"
```

## Extending DNS config

In order to configure additional DNS servers for your traefik pod, you can use `dnsConfig` option:

```yaml
deployment:
  dnsConfig:
    nameservers:
      - 192.0.2.1 # this is an example
    searches:
      - ns1.svc.cluster-domain.example
      - my.dns.search.suffix
    options:
      - name: ndots
        value: "2"
      - name: edns0
```

## Install in a dedicated namespace, with limited RBAC

Default install is using Cluster-wide RBAC but it can be restricted to target namespace.

```yaml
rbac:
  namespaced: true
```

## Install with auto-scaling

When enabling [HPA](https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/)
to adjust replicas count according to CPU Usage, you'll need to set resources and nullify replicas.

```yaml
deployment:
  replicas: null
resources:
  requests:
    cpu: "100m"
    memory: "50Mi"
  limits:
    cpu: "300m"
    memory: "150Mi"
autoscaling:
  enabled: true
  maxReplicas: 2
  metrics:
  - type: Resource
    resource:
      name: cpu
      target:
        type: Utilization
        averageUtilization: 80
```

## Install with Argo Rollouts

When using [ArgoCD Rollouts](https://argoproj.github.io/rollouts/), one can delegate replica management to a `Rollout` resource, enabling progressive delivery strategies like canary and blue-green deployments.
In order to delegate replica management, `deployment.replicas` should be set to `0` and the `Rollout` resource can be defined in a separate YAML or in `extraObjects`.

```yaml
deployment:
  replicas: 0
autoscaling:
  enabled: true
  minReplicas: 5
  maxReplicas: 50
  metrics:
    - type: Resource
      resource:
        name: cpu
        target:
          type: Utilization
          averageUtilization: 80
  scaleTargetRef:
    apiVersion: argoproj.io/v1alpha1
    kind: Rollout
extraObjects:
  - apiVersion: argoproj.io/v1alpha1
    kind: Rollout
    metadata:
      name: "{{ template \"traefik.fullname\" . }}"
    spec:
      workloadRef:
        apiVersion: apps/v1
        kind: Deployment
        name: "{{ template \"traefik.fullname\" . }}"
      strategy:
        canary:
          steps:
          - setWeight: 10
          - pause:
              duration: 5m
```

## Access Traefik dashboard without exposing it

This Chart does not expose the Traefik local dashboard by default. It's explained in upstream [documentation](https://doc.traefik.io/traefik/operations/api/) why:

> Enabling the API in production is not recommended, because it will expose all configuration elements, including sensitive data.

It says also:

> In production, it should be at least secured by authentication and authorizations.

Thus, there are multiple ways to expose the dashboard. For instance, after enabling the creation of dashboard `IngressRoute` in the values:

```yaml
ingressRoute:
  dashboard:
    enabled: true
```

The traefik admin port can be forwarded locally. Assuming the default `traefik` namespace is used:

```bash
NAMESPACE=traefik
kubectl port-forward $(kubectl get pods --selector "app.kubernetes.io/name=traefik" --output=name -n $NAMESPACE) 8080:8080 -n $NAMESPACE
```

This command makes the dashboard accessible locally on [127.0.0.1:8080/dashboard/](http://127.0.0.1:8080/dashboard/)

> [!IMPORTANT]
> Note that the slash is required.

## Redirect permanently traffic from http to https

It's possible to redirect all incoming requests on an entrypoint to an other entrypoint.

```yaml
ports:
  web:
    redirections:
      entryPoint:
        to: websecure
        scheme: https
        permanent: true
```

## Publish and protect Traefik Dashboard with basic Auth

To expose the dashboard in a secure way as [recommended](https://doc.traefik.io/traefik/operations/dashboard/#dashboard-router-rule)
in the documentation, it may be useful to override the router rule to specify
a domain to match, or accept requests on the root path (/) in order to redirect
them to /dashboard/.

```yaml
# Create an IngressRoute for the dashboard
ingressRoute:
  dashboard:
    enabled: true
    # Custom match rule with host domain
    matchRule: Host(`traefik-dashboard.example.com`)
    entryPoints: ["websecure"]
    # Add custom middlewares : authentication and redirection
    middlewares:
      - name: traefik-dashboard-auth

# Create the custom middlewares used by the IngressRoute dashboard (can also be created in another way).
# /!\ Yes, you need to replace "changeme" password with a better one. /!\
extraObjects:
  - apiVersion: v1
    kind: Secret
    metadata:
      name: traefik-dashboard-auth-secret
    type: kubernetes.io/basic-auth
    stringData:
      username: admin
      password: changeme

  - apiVersion: traefik.io/v1alpha1
    kind: Middleware
    metadata:
      name: traefik-dashboard-auth
    spec:
      basicAuth:
        secret: traefik-dashboard-auth-secret
```

## Publish and protect Traefik Dashboard with an Ingress

To expose the dashboard without IngressRoute, it's more complicated and less
secure. You'll need to create an internal Service exposing Traefik API with
special _traefik_ entrypoint. This internal Service can be created from an other tool, with the `extraObjects` section or using [custom services](#add-custom-internal-services).

You'll need to double check:

1. Service selector with your setup.
2. Middleware annotation on the ingress, _default_ should be replaced with traefik's namespace

```yaml
ingressRoute:
  dashboard:
    enabled: false
additionalArguments:
- "--api.insecure=true"
# Create the service, middleware and Ingress used to expose the dashboard (can also be created in another way).
# /!\ Yes, you need to replace "changeme" password with a better one. /!\
extraObjects:
  - apiVersion: v1
    kind: Service
    metadata:
      name: traefik-api
    spec:
      type: ClusterIP
      selector:
        app.kubernetes.io/name: traefik
        app.kubernetes.io/instance: traefik-default
      ports:
      - port: 8080
        name: traefik
        targetPort: 8080
        protocol: TCP

  - apiVersion: v1
    kind: Secret
    metadata:
      name: traefik-dashboard-auth-secret
    type: kubernetes.io/basic-auth
    stringData:
      username: admin
      password: changeme

  - apiVersion: traefik.io/v1alpha1
    kind: Middleware
    metadata:
      name: traefik-dashboard-auth
    spec:
      basicAuth:
        secret: traefik-dashboard-auth-secret

  - apiVersion: networking.k8s.io/v1
    kind: Ingress
    metadata:
      name: traefik-dashboard
      annotations:
        traefik.ingress.kubernetes.io/router.entrypoints: websecure
        traefik.ingress.kubernetes.io/router.middlewares: default-traefik-dashboard-auth@kubernetescrd
    spec:
      rules:
      - host: traefik-dashboard.example.com
        http:
          paths:
          - path: /
            pathType: Prefix
            backend:
              service:
                name: traefik-api
                port:
                  name: traefik
```

## Publish Traefik Dashboard in Rancher UI

To expose the dashboard with rancher UI some paths modifications are required.
`basePath` needs to be changed and a `Middleware` needs to be used to URL rewriting.

```yaml
# Configure the basePath
api:
  basePath: "/api/v1/namespaces/traefik/services/https:traefik:443/proxy/"

# Create an IngressRoute for the dashboard
ingressRoute:
  dashboard:
    enabled: true
    # Custom match rule with host domain
    matchRule: PathPrefix(`/dashboard`) || PathPrefix(`/api`)
    entryPoints: ["websecure"]
    # Add custom middleware : this makes the path matching the internal Go router 
    middlewares:
      - name: traefik-dashboard-basepath

# Create the custom middlewares used by the IngressRoute dashboard (can also be created from an other source).
extraObjects:
  - apiVersion: traefik.io/v1alpha1
    kind: Middleware
    metadata:
      name: traefik-dashboard-basepath
    spec:
      addPrefix:
        prefix: "/api/v1/namespaces/traefik/services/https:traefik:443/proxy"
```

## Install on AWS

It can use [native AWS support](https://kubernetes.io/docs/concepts/services-networking/service/#aws-nlb-support) on Kubernetes

```yaml
service:
  annotations:
    service.beta.kubernetes.io/aws-load-balancer-type: nlb
```

Or if [AWS LB controller](https://kubernetes-sigs.github.io/aws-load-balancer-controller/v2.2/guide/service/annotations/#legacy-cloud-provider) is installed :

```yaml
service:
  annotations:
    service.beta.kubernetes.io/aws-load-balancer-type: nlb-ip
```

## Install on GCP

A [regional IP with a Service](https://cloud.google.com/kubernetes-engine/docs/tutorials/configuring-domain-name-static-ip#use_a_service) can be used

```yaml
service:
  spec:
    loadBalancerIP: "1.2.3.4"
```

Or a [global IP on Ingress](https://cloud.google.com/kubernetes-engine/docs/tutorials/configuring-domain-name-static-ip#use_an_ingress)

```yaml
service:
  type: NodePort
extraObjects:
  - apiVersion: networking.k8s.io/v1
    kind: Ingress
    metadata:
      name: traefik
      annotations:
        kubernetes.io/ingress.global-static-ip-name: "myGlobalIpName"
    spec:
      defaultBackend:
        service:
          name: traefik
          port:
            number: 80
```

Or a [global IP on a Gateway](https://cloud.google.com/kubernetes-engine/docs/how-to/deploying-gateways) with continuous HTTPS encryption.

```yaml
ports:
  websecure:
    appProtocol: HTTPS # Hint for Google L7 load balancer
service:
  type: ClusterIP
extraObjects:
- apiVersion: gateway.networking.k8s.io/v1beta1
  kind: Gateway
  metadata:
    name: traefik
    annotations:
      networking.gke.io/certmap: "myCertificateMap"
  spec:
    gatewayClassName: gke-l7-global-external-managed
    addresses:
    - type: NamedAddress
      value: "myGlobalIPName"
    listeners:
    - name: https
      protocol: HTTPS
      port: 443
- apiVersion: gateway.networking.k8s.io/v1beta1
  kind: HTTPRoute
  metadata:
    name: traefik
  spec:
    parentRefs:
    - kind: Gateway
      name: traefik
    rules:
    - backendRefs:
      - name: traefik
        port: 443
- apiVersion: networking.gke.io/v1
  kind: HealthCheckPolicy
  metadata:
    name: traefik
  spec:
    default:
      config:
        type: HTTP
        httpHealthCheck:
          port: 8080
          requestPath: /ping
    targetRef:
      group: ""
      kind: Service
      name: traefik
```

## Install on Azure

A [static IP on a resource group](https://learn.microsoft.com/en-us/azure/aks/static-ip) can be used:

```yaml
service:
  spec:
    loadBalancerIP: "1.2.3.4"
  annotations:
    service.beta.kubernetes.io/azure-load-balancer-resource-group: myResourceGroup
```

Here is a more complete example, using also native Let's encrypt feature of Traefik Proxy with Azure DNS:

```yaml
persistence:
  enabled: true
  size: 128Mi
certificatesResolvers:
  letsencrypt:
    acme:
      email: "{{ letsencrypt_email }}"
      #caServer: https://acme-v02.api.letsencrypt.org/directory # Production server
      caServer: https://acme-staging-v02.api.letsencrypt.org/directory # Staging server
      dnsChallenge:
        provider: azuredns
      storage: /data/acme.json
env:
  - name: AZURE_CLIENT_ID
    value: "{{ azure_dns_challenge_application_id }}"
  - name: AZURE_CLIENT_SECRET
    valueFrom:
      secretKeyRef:
        name: azuredns-secret
        key: client-secret
  - name: AZURE_SUBSCRIPTION_ID
    value: "{{ azure_subscription_id }}"
  - name: AZURE_TENANT_ID
    value: "{{ azure_tenant_id }}"
  - name: AZURE_RESOURCE_GROUP
    value: "{{ azure_resource_group }}"
deployment:
  initContainers:
    - name: volume-permissions
      image: busybox:latest
      command: ["sh", "-c", "ls -la /; touch /data/acme.json; chmod -v 600 /data/acme.json"]
      volumeMounts:
      - mountPath: /data
        name: data
podSecurityContext:
  fsGroup: 65532
  fsGroupChangePolicy: "OnRootMismatch"
service:
  spec:
    type: LoadBalancer
  annotations:
    service.beta.kubernetes.io/azure-load-balancer-resource-group: "{{ azure_node_resource_group }}"
    service.beta.kubernetes.io/azure-pip-name: "{{ azure_resource_group }}"
    service.beta.kubernetes.io/azure-dns-label-name: "{{ azure_resource_group }}"
    service.beta.kubernetes.io/azure-allowed-ip-ranges: "{{ ip_range | join(',') }}"
extraObjects:
  - apiVersion: v1
    kind: Secret
    metadata:
      name: azuredns-secret
      namespace: traefik
    type: Opaque
    stringData:
      client-secret: "{{ azure_dns_challenge_application_secret }}"
```

## Use ServiceMonitor on AKS (Azure Monitor / managed Prometheus)

Enable the optional ServiceMonitor so managed Prometheus can scrape Traefik metrics on AKS. You may override the CRD apiVersion if your environment requires it.

```yaml
metrics:
  prometheus:
    service:
      enabled: true
    # Set to true when using Azure Monitor to skip the CRD check (monitoring.coreos.com/v1)
    disableAPICheck: true
    serviceMonitor:
      enabled: true
      # Defaults to monitoring.coreos.com/v1
      apiVersion: "azmonitoring.coreos.com/v1"
    prometheusRule:
      # Defaults to monitoring.coreos.com/v1
      apiVersion: "azmonitoring.coreos.com/v1"
```

## Use an IngressClass

Default install comes with an `IngressClass` resource that can be enabled on providers.

Here's how one can enable it on CRD & Ingress Kubernetes provider:

```yaml
ingressClass:
  name: traefik
providers:
  kubernetesCRD:
    ingressClass: traefik
  kubernetesIngress:
    ingressClass: traefik
```

## Use HTTP3

By default, it will use a Load balancers with mixed protocols on `websecure`
entrypoint. They are available since v1.20 and in beta as of Kubernetes v1.24.
Availability may depend on your Kubernetes provider.

When using TCP and UDP with a single service, you may encounter [this issue](https://github.com/kubernetes/kubernetes/issues/47249#issuecomment-587960741) from Kubernetes.
If you want to avoid this issue, you can set `ports.websecure.http3.advertisedPort`
to an other value than 443

```yaml
ports:
  websecure:
    http3:
      enabled: true
```

You can also create two `Service`, one for TCP and one for UDP:

```yaml
ports:
  websecure:
    http3:
      enabled: true
service:
  single: false
```

## Use PROXY protocol on Digital Ocean

PROXY protocol is a protocol for sending client connection information, such as origin IP addresses and port numbers, to the final backend server, rather than discarding it at the load balancer.

```yaml
.DOTrustedIPs: &DOTrustedIPs
  - 127.0.0.1/32
  # IP range Load Balancer is on
  - 10.0.0.0/8
  # IP range of private (VPC) interface - CHANGE THIS TO YOUR NETWORK SETTINGS
  # This is needed when "externalTrafficPolicy: Cluster" is specified, as inbound traffic from the load balancer to a Traefik instance could be redirected from another cluster node on the way through.
  - 172.16.0.0/12

service:
  enabled: true
  type: LoadBalancer
  annotations:
    # This will tell DigitalOcean to enable the proxy protocol.
    service.beta.kubernetes.io/do-loadbalancer-enable-proxy-protocol: "true"
  spec:
    # This is the default and should stay as cluster to keep the DO health checks working.
    externalTrafficPolicy: Cluster

ports:
  web:
    forwardedHeaders:
      trustedIPs: *DOTrustedIPs
    proxyProtocol:
      trustedIPs: *DOTrustedIPs
  websecure:
    forwardedHeaders:
      trustedIPs: *DOTrustedIPs
    proxyProtocol:
      trustedIPs: *DOTrustedIPs
```

## Using plugins

This chart follows common security practices: it runs as non-root with a readonly root filesystem.
When enabling a plugin, this Chart provides by default an `emptyDir` for plugin storage.

Here is an example with [crowdsec](https://github.com/maxlerebourg/crowdsec-bouncer-traefik-plugin/blob/main/examples/kubernetes/README.md) plugin:

```yaml
experimental:
  plugins:
    demo:
      moduleName: github.com/maxlerebourg/crowdsec-bouncer-traefik-plugin
      version: v1.3.5
```

When persistence is needed, this `emptyDir` can be replaced with a PVC by adding:

```yaml
deployment:
  additionalVolumes:
  - name: plugins
    persistentVolumeClaim:
      claimName: my-plugins-vol
additionalVolumeMounts:
- name: plugins
  mountPath: /plugins-storage
extraObjects:
  - kind: PersistentVolumeClaim
    apiVersion: v1
    metadata:
      name: my-plugins-vol
    spec:
      accessModes:
        - ReadWriteOnce
      resources:
        requests:
          storage: 1Gi
```

## Local Plugins

To develop or test plugins without pushing them to a public registry, you can load plugin source code directly from your local filesystem.

>[!NOTE]
> The ``hostPath`` must point to a directory containing the plugin source code and a valid ``go.mod`` file. The ``moduleName`` must match the module name specified in the ``go.mod`` file.

```yaml
experimental:
  localPlugins:
    local-demo:
      moduleName: github.com/traefik/localplugindemo
      mountPath: /plugins-local/src/github.com/traefik/localplugindemo
      hostPath: /path/to/plugin-source
```

>[!IMPORTANT]
> When using ``hostPath`` volumes, the plugin source code must be available on every node where Traefik pods might be scheduled.

## Using Traefik-Hub with private plugin registries

With Traefik Hub, it's possible to use plugins deployed on both public or private registries.
Each registry source requires a base module name (domain) and authentication credentials.
This can be achieved this way:

```yaml
hub:
  token: traefik-hub-license
  pluginRegistry:
    sources:
      noop:
        baseModuleName: "github.com"
        github:
          token: "ghp_XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"

image:
  registry: ghcr.io
  repository: traefik/traefik-hub
  tag: v3.18.0

experimental:
  plugins:
    noop:
      moduleName: github.com/traefik-contrib/noop
      version: v0.1.0

extraObjects:
  - apiVersion: traefik.io/v1alpha1
    kind: Middleware
    metadata:
      name: noop
    spec:
      plugin:
        noop:
          responseCode: 204
  - apiVersion: traefik.io/v1alpha1
    kind: IngressRoute
    metadata:
      name: demo
    spec:
      entryPoints:
        - web
      routes:
        - kind: Rule
          match: Host(`demo.localhost`)
          services:
            - name: noop@internal
              kind: TraefikService
          middlewares:
            - name: noop
```

> [!NOTE]  
> This code is only written for demonstration purpose.
> The prefered way of configuration either Github or Gitlab credentials is to use an URN like `urn:k8s:secret:github-token:access-token`.

## Use Traefik native Let's Encrypt integration, without cert-manager

In Traefik Proxy, ACME certificates are stored in a JSON file.

This file needs to have 0600 permissions, meaning, only the owner of the file has full read and write access to it.
By default, Kubernetes recursively changes ownership and permissions for the content of each volume.

=> An initContainer can be used to avoid an issue on this sensitive file.
See [#396](https://github.com/traefik/traefik-helm-chart/issues/396) for more details.

Once the provider is ready, it can be used in an `IngressRoute`:

```yaml
apiVersion: traefik.io/v1alpha1
kind: IngressRoute
metadata:
  name: [...]
spec:
  entryPoints: [...]
  routes: [...]
  tls:
    certResolver: letsencrypt
```

:information_source: Change `apiVersion` to `traefik.containo.us/v1alpha1` for charts prior to v28.0.0

See [the list of supported providers](https://doc.traefik.io/traefik/https/acme/#providers) for others.

## Example with CloudFlare

This example needs a CloudFlare token in a Kubernetes `Secret` and a working `StorageClass`.

**Step 1**: Create `Secret` with CloudFlare token:

```yaml
---
apiVersion: v1
kind: Secret
metadata:
  name: cloudflare
type: Opaque
stringData:
  token: {{ SET_A_VALID_TOKEN_HERE }}
```

**Step 2**:

```yaml
persistence:
  enabled: true
  storageClass: xxx
certificatesResolvers:
  letsencrypt:
    acme:
      dnsChallenge:
        provider: cloudflare
      storage: /data/acme.json
env:
  - name: CF_DNS_API_TOKEN
    valueFrom:
      secretKeyRef:
        name: cloudflare
        key: token
deployment:
  initContainers:
    - name: volume-permissions
      image: busybox:latest
      command: ["sh", "-c", "touch /data/acme.json; chmod -v 600 /data/acme.json"]
      volumeMounts:
      - mountPath: /data
        name: data
podSecurityContext:
  fsGroup: 65532
  fsGroupChangePolicy: "OnRootMismatch"
```

>[!NOTE]
> With [Traefik Hub](https://traefik.io/traefik-hub/), certificates can be stored as a `Secret` on Kubernetes with `distributedAcme` resolver.

## Provide default certificate with cert-manager and CloudFlare DNS

Setup:

* cert-manager installed in `cert-manager` namespace
* A cloudflare account on a DNS Zone

**Step 1**: Create `Secret` and `Issuer` needed by `cert-manager` with your API Token.
See [cert-manager documentation](https://cert-manager.io/docs/configuration/acme/dns01/cloudflare/)
for creating this token with needed rights:

```yaml
---
apiVersion: v1
kind: Secret
metadata:
  name: cloudflare
  namespace: traefik
type: Opaque
stringData:
  api-token: XXX
---
apiVersion: cert-manager.io/v1
kind: Issuer
metadata:
  name: cloudflare
  namespace: traefik
spec:
  acme:
    server: https://acme-v02.api.letsencrypt.org/directory
    email: email@example.com
    privateKeySecretRef:
      name: cloudflare-key
    solvers:
      - dns01:
          cloudflare:
            apiTokenSecretRef:
              name: cloudflare
              key: api-token
```

**Step 2**: Create `Certificate` in traefik namespace

```yaml
apiVersion: cert-manager.io/v1
kind: Certificate
metadata:
  name: wildcard-example-com
  namespace: traefik
spec:
  secretName: wildcard-example-com-tls
  dnsNames:
    - "example.com"
    - "*.example.com"
  issuerRef:
    name: cloudflare
    kind: Issuer
```

**Step 3**: Check that it's ready

```bash
kubectl get certificate -n traefik
```

If needed, logs of cert-manager pod can give you more information

**Step 4**: Use it on the TLS Store in **values.yaml** file for this Helm Chart

```yaml
tlsStore:
  default:
    defaultCertificate:
      secretName: wildcard-example-com-tls
```

**Step 5**: Enjoy. All your `IngressRoute` use this certificate by default now.

They should use websecure entrypoint like this:

```yaml
apiVersion: traefik.io/v1alpha1
kind: IngressRoute
metadata:
  name: example-com-tls
spec:
  entryPoints:
    - websecure
  routes:
  - match: Host(`test.example.com`)
    kind: Rule
    services:
    - name: XXXX
      port: 80
```

## Add custom (internal) services

In some cases you might want to have more than one Traefik service within your cluster,
e.g. a default (external) one and a service that is only exposed internally to pods within your cluster.

The `service.additionalServices` allows you to add an arbitrary amount of services,
provided as a name to service details mapping; for example you can use the following values:

```yaml
service:
  additionalServices:
    internal:
      type: ClusterIP
      labels:
        traefik-service-label: internal
```

Ports can then be exposed on this service by using the port name to boolean mapping `expose` on the respective port;
e.g. to expose the `traefik` API port on your internal service so pods within your cluster can use it, you can do:

```yaml
ports:
  traefik:
    expose:
      # Sensitive data should not be exposed on the internet
      # => Keep this disabled !
      default: false
      internal: true
```

This will then provide an additional Service manifest, looking like this:

```yaml
---
# Source: traefik/templates/service.yaml
apiVersion: v1
kind: Service
metadata:
  name: traefik-internal
  namespace: traefik
[...]
spec:
  type: ClusterIP
  selector:
    app.kubernetes.io/name: traefik
    app.kubernetes.io/instance: traefik-traefik
  ports:
  - port: 8080
    name: traefik
    targetPort: traefik
    protocol: TCP
```

## Use this Chart as a dependency of your own chart

First, let's create a default Helm Chart, with Traefik as a dependency.

```bash
helm create foo
cd foo
echo "
dependencies:
  - name: traefik
    version: "24.0.0"
    repository: "https://traefik.github.io/charts"
" >> Chart.yaml
```

Second, let's tune some values like enabling HPA:

```bash
cat <<-EOF >> values.yaml
traefik:
  autoscaling:
    enabled: true
    maxReplicas: 3
EOF
```

Third, one can see if it works as expected:

```bash
helm dependency update
helm dependency build
helm template . | grep -A 14 -B 3 Horizontal
```

It should produce this output:

```yaml
---
# Source: foo/charts/traefik/templates/hpa.yaml
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: release-name-traefik
  namespace: flux-system
  labels:
    app.kubernetes.io/name: traefik
    app.kubernetes.io/instance: release-name-flux-system
    helm.sh/chart: traefik-24.0.0
    app.kubernetes.io/managed-by: Helm
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: release-name-traefik
  maxReplicas: 3
```

## Configure TLS

The [TLS options](https://doc.traefik.io/traefik/https/tls/#tls-options) allow one to configure some parameters of the TLS connection.

```yaml
tlsOptions:
  default:
    labels: {}
    sniStrict: true
  custom-options:
    labels: {}
    curvePreferences:
      - CurveP521
      - CurveP384
```

## Use latest build of Traefik v3 from master

An experimental build of Traefik Proxy is available on a specific community repository: `traefik/traefik`.

The tag does not follow semver, so it requires a _versionOverride_:

```yaml
image:
  repository: traefik/traefik
  tag: experimental-v3.4
versionOverride: v3.4
```

## Use Prometheus Operator

An optional support of this operator is included in this Chart. See documentation of this operator for more details.

It can be used with those _values_:

```yaml
metrics:
  prometheus:
    service:
      enabled: true
    disableAPICheck: false
    serviceMonitor:
      enabled: true
      metricRelabelings:
        - sourceLabels: [__name__]
          separator: ;
          regex: ^fluentd_output_status_buffer_(oldest|newest)_.+
          replacement: $1
          action: drop
      relabelings:
        - sourceLabels: [__meta_kubernetes_pod_node_name]
          separator: ;
          regex: ^(.*)$
          targetLabel: nodename
          replacement: $1
          action: replace
      jobLabel: traefik
      interval: 30s
      honorLabels: true
    headerLabels:
      user_id: X-User-Id
      tenant: X-Tenant
    prometheusRule:
      enabled: true
      rules:
        - alert: TraefikDown
          expr: up{job="traefik"} == 0
          for: 5m
          labels:
            context: traefik
            severity: warning
          annotations:
            summary: "Traefik Down"
            description: "{{ $labels.pod }} on {{ $labels.nodename }} is down"
```

## Use kubernetes Gateway API

One can use the new stable kubernetes gateway API provider by setting the following _values_:

```yaml
providers:
  kubernetesGateway:
    enabled: true
```

<details>

<summary>With those values, a whoami service can be exposed with an HTTPRoute</summary>

```yaml
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: whoami
spec:
  replicas: 2
  selector:
    matchLabels:
      app: whoami
  template:
    metadata:
      labels:
        app: whoami
    spec:
      containers:
        - name: whoami
          image: traefik/whoami

---
apiVersion: v1
kind: Service
metadata:
  name: whoami
spec:
  selector:
    app: whoami
  ports:
    - protocol: TCP
      port: 80

---
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: whoami
spec:
  parentRefs:
    - name: traefik-gateway
  hostnames:
    - whoami.docker.localhost
  rules:
    - matches:
        - path:
            type: Exact
            value: /

      backendRefs:
        - name: whoami
          port: 80
          weight: 1
```

Once it's applied, whoami should be accessible on [whoami.docker.localhost](http://whoami.docker.localhost/)

</details>

:information_source: In this example, `Deployment` and `HTTPRoute` should be deployed in the same namespace as the Traefik Gateway: Chart namespace.

## Use Kubernetes Gateway API with cert-manager

One can use the new stable kubernetes gateway API provider with automatic TLS certificates delivery (with cert-manager) by setting the following _values_:

```yaml
providers:
  kubernetesGateway:
    enabled: true
gateway:
  enabled: true
  annotations:
    cert-manager.io/issuer: selfsigned-issuer
  listeners:
    websecure:
      hostname: whoami.docker.localhost
      port: 8443
      protocol: HTTPS
      certificateRefs:
        - name: whoami-tls
```

Install cert-manager:

```bash
helm repo add jetstack https://charts.jetstack.io --force-update
helm upgrade --install \
cert-manager jetstack/cert-manager \
--namespace cert-manager \
--create-namespace \
--version v1.15.1 \
--set crds.enabled=true \
--set "extraArgs={--enable-gateway-api}"
```

<details>

<summary>With those values, a whoami service can be exposed with HTTPRoute on both HTTP and HTTPS</summary>

```yaml
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: whoami
spec:
  replicas: 2
  selector:
    matchLabels:
      app: whoami
  template:
    metadata:
      labels:
        app: whoami
    spec:
      containers:
        - name: whoami
          image: traefik/whoami

---
apiVersion: v1
kind: Service
metadata:
  name: whoami
spec:
  selector:
    app: whoami
  ports:
    - protocol: TCP
      port: 80

---
apiVersion: gateway.networking.k8s.io/v1
kind: HTTPRoute
metadata:
  name: whoami
spec:
  parentRefs:
    - name: traefik-gateway
  hostnames:
    - whoami.docker.localhost
  rules:
    - matches:
        - path:
            type: Exact
            value: /

      backendRefs:
        - name: whoami
          port: 80
          weight: 1

---
apiVersion: cert-manager.io/v1
kind: Issuer
metadata:
  name: selfsigned-issuer
spec:
  selfSigned: {}
```

Once it's applied, whoami should be accessible on https://whoami.docker.localhost

</details>

## Use Knative Provider

Starting with Traefik Proxy v3.6, one can use the Knative provider (_experimental_) by setting the following _values_:

```yaml
experimental:
  knative: true
providers:
  knative:
    enabled: true
```

> [!WARNING]
> You must first have Knative deployed. With Proxy v3.6, v1.19 of Knative is supported.
> Knative 1.19 requires Kubernetes v1.32+

> [!TIP]
> If you want to test it using k3d, you'll need to set the image accordingly, for instance: `--image rancher/k3s:v1.34.1-k3s1`

Finish configuring Knative:

```shell
# 1. Install/update the Knative CRDs
kubectl apply -f https://github.com/knative/serving/releases/download/knative-v1.19.0/serving-crds.yaml
# 2. Install the Knative Serving core components
kubectl apply -f https://github.com/knative/serving/releases/download/knative-v1.19.0/serving-core.yaml
# 3. Update the config-network configuration to use the Traefik ingress class
kubectl patch configmap/config-network -n knative-serving --type merge \
  -p '{"data":{"ingress.class":"traefik.ingress.networking.knative.dev"}}'
# Add a custom domain to Knative configuration (in this example, docker.localhost)
kubectl patch configmap config-domain -n knative-serving --type='merge' \
  -p='{"data":{"docker.localhost":""}}'
```

With that done and the specified values set, a Knative Service can now be deployed:

```yaml
---
apiVersion: serving.knative.dev/v1
kind: Service
metadata:
  name: whoami
spec:
  template:
    spec:
      containers:
        - name: whoami
          image: traefik/whoami
          ports:
            - containerPort: 80
```

Once it's applied, we can check the URLs:

```shell
# 1. List Knative services
kubectl get ksvc
# 2. Test URLs
curl http://whoami.default.docker.localhost
curl -k -H "Host: whoami.default.docker.localhost" https://localhost/
```

## Use templating for additionalVolumeMounts

This example demonstrates how to use templating for the `additionalVolumeMounts` configuration to dynamically set the `subPath` parameter based on a variable.

```yaml
additionalVolumeMounts:
  - name: plugin-volume
    mountPath: /plugins
    subPath: "{{ .Values.pluginVersion }}"
```

In your `values.yaml` file, you can specify the `pluginVersion` variable:

```yaml
pluginVersion: "v1.2.3"
```

This configuration will mount the `plugin-volume` at `/plugins` with the `subPath` set to `v1.2.3`.

## Use a custom certificate for Traefik Hub webhooks

Some CD tools may regenerate Traefik Hub mutating webhooks continuously, when using helm template.
This example demonstrates how to generate and use a custom certificate for Hub admission webhooks.

First, generate a self-signed certificate:

```bash
# this generates a self-signed certificate with a 2048 bits key, valid for 10 years, on admission.traefik.svc DNS name
openssl req -x509 -newkey rsa:2048 -sha256 -days 3650 -nodes -keyout /tmp/hub.key -out /tmp/hub.crt \
            -subj '/CN=admission.traefik.svc' -addext "subjectAltName=DNS:admission.traefik.svc" \
            -addext basicConstraints=critical,CA:FALSE -addext "keyUsage = digitalSignature, keyEncipherment" -addext "extendedKeyUsage = serverAuth, clientAuth"
cat /tmp/hub.crt | base64 -w0 > /tmp/hub.crt.b64
cat /tmp/hub.key | base64 -w0 > /tmp/hub.key.b64
```

Now, it can be set in the `values.yaml`:

```yaml
hub:
  token: traefik-hub-license
  apimanagement:
    enabled: true
    admission:
      customWebhookCertificate:
        tls.crt: xxxx # content of /tmp/hub.crt.b64
        tls.key: xxxx # content of /tmp/hub.key.b64
```

> [!TIP]
> When using the CLI, those parameters need to be escaped like this:
>
>```bash
> --set 'hub.apimanagement.admission.customWebhookCertificate.tls\.crt'=$(cat /tmp/hub.crt.b64)
> --set 'hub.apimanagement.admission.customWebhookCertificate.tls\.key'=$(cat /tmp/hub.key.b64)
>```

## Injecting CA data from a Certificate resource

It is also possible to use [CA injector](https://cert-manager.io/docs/concepts/ca-injector/) of cert-manager with annotations on the webhook.

First, you can create the certificate with a self-signed issuer:

```yaml
---
apiVersion: cert-manager.io/v1
kind: Certificate
metadata:
  name: admission
  namespace: traefik
spec:
  secretName: admission-tls
  dnsNames:
  - admission.traefik.svc
  issuerRef:
    name: selfsigned

---
apiVersion: cert-manager.io/v1
kind: Issuer
metadata:
  name: selfsigned
  namespace: traefik
spec:
  selfSigned: {}
```

Once the `Certificate` is ready, it can be set in the `values.yaml` like this:

```yaml
hub:
  token: traefik-hub-license
  apimanagement:
    enabled: true
    admission:
      selfManagedCertificate: true
      secretName: admission-tls
      annotations:
        cert-manager.io/inject-ca-from: traefik/admission-tls
```

## Use a custom certificate for Traefik Hub webhooks from an existing secret

Some CD tools may regenerate Traefik Hub mutating webhooks continuously, when using helm template.
This example demonstrates how to generate and use a custom certificate stored in a managed secret for Hub admission webhooks.

First, generate a self-signed certificate:

```bash
# this generates a self-signed certificate with a 2048 bits key, valid for 10 years, on admission.traefik.svc DNS name
openssl req -x509 -newkey rsa:2048 -sha256 -days 3650 -nodes -keyout tls.key -out tls.crt \
            -subj '/CN=admission.traefik.svc' -addext "subjectAltName=DNS:admission.traefik.svc" \
            -addext basicConstraints=critical,CA:FALSE -addext "keyUsage = digitalSignature, keyEncipherment" -addext "extendedKeyUsage = serverAuth, clientAuth"
```

Create secret from generated certificate files:

```bash
kubectl create secret tls hub-admission-cert --namespace traefik --cert=tls.crt --key=tls.key
```

Now, it can be set in the `values.yaml`:

```yaml
hub:
  apimanagement:
    admission:
      selfManagedCertificate: true
```

## Mount datadog DSD socket directly into traefik container (i.e. no more socat sidecar)

This example demonstrates how to directly mount datadog apm socket into traefik container, thus avoiding the need of socat sidecar container.

```yaml
metrics:
  datadog:
    address: unix:///var/run/datadog/dsd.socket # https://doc.traefik.io/traefik/observability/metrics/datadog/#address
additionalVolumeMounts:
  - name: ddsocketdir
    mountPath: /var/run/datadog
    readOnly: false
deployment:
  additionalVolumes:
    - hostPath:
        path: /var/run/datadog/
      name: ddsocketdir
```

## Use Traefik Hub AI Gateway

This example demonstrates how to enable AI Gateway in Traefik Hub and set a maxRequestBodySize of 10 MiB.

```yaml
hub:
  token: # <=== Set your token here
  aigateway:
    enabled: true
    maxRequestBodySize: 10485760 # optional, default to 1MiB
```
