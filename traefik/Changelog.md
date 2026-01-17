# Change Log

## 37.3.0  ![AppVersion: v3.6.0](https://img.shields.io/static/v1?label=AppVersion&message=v3.6.0&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2025-11-10

* tests: update certificatesResolvers with Proxy v3.6 options
* fix: add missing flag arg
* fix(doc): :books: comment grammar in values.yaml
* feat: knative provider
* feat(deps): update traefik docker tag to v3.6.0
* feat(deps): update traefik docker tag to v3.5.4
* feat(CRDs): update for Traefik Proxy v3.6 and Gateway API v1.4.0
* feat(CRDs): update Traefik Hub to v1.23.1
* docs: 📚 fix comment grammar in values
* chore(release): 🚀 publish traefik 37.3.0 and 1.12.0

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index c697f6f..e90d6b9 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -141,13 +141,15 @@ experimental:
   localPlugins: {}
   # -- Enable OTLP logging experimental feature.
   otlpLogs: false
+  # -- Enable Knative provider experimental feature.
+  knative: false
 
 gateway:
   # -- When providers.kubernetesGateway.enabled, deploy a default gateway
   enabled: true
   # -- Set a custom name to gateway
   name: ""
-  # -- By default, Gateway is created in the same `Namespace` than Traefik.
+  # -- By default, Gateway is created in the same `Namespace` as Traefik.
   namespace: ""
   # -- Additional gateway annotations (e.g. for cert-manager.io/issuer)
   annotations: {}
@@ -281,7 +283,7 @@ providers:  # @schema additionalProperties: false
     allowCrossNamespace: false
     # -- Allows to reference ExternalName services in IngressRoute
     allowExternalNameServices: false
-    # -- Allows to return 503 when there is no endpoints available
+    # -- Allows to return 503 when there are no endpoints available
     allowEmptyServices: true
     # -- When the parameter is set, only resources containing an annotation with the same value are processed. Otherwise, resources missing the annotation, having an empty value, or the value traefik are processed. It will also set required annotation on Dashboard and Healthcheck IngressRoute when enabled.
     ingressClass: ""
@@ -296,7 +298,7 @@ providers:  # @schema additionalProperties: false
     enabled: true
     # -- Allows to reference ExternalName services in Ingress
     allowExternalNameServices: false
-    # -- Allows to return 503 when there is no endpoints available
+    # -- Allows to return 503 when there are no endpoints available
     allowEmptyServices: true
     # -- When ingressClass is set, only Ingresses containing an annotation with the same value are processed. Otherwise, Ingresses missing the annotation, having an empty value, or the value traefik are processed.
     ingressClass:  # @schema type:[string, null]
@@ -346,6 +348,14 @@ providers:  # @schema additionalProperties: false
     # -- File content (YAML format, go template supported) (see https://doc.traefik.io/traefik/providers/file/)
     content: ""
 
+  knative:
+    # -- Enable Knative provider
+    enabled: false
+    # -- Array of namespaces to watch. If left empty, Traefik watches all namespaces. . When using `rbac.namespaced`, it will watch helm release namespace and namespaces listed in this array.
+    namespaces: []
+    # -- Allow filtering Knative Ingress objects
+    labelselector: ""
+
 # -- Add volumes to the traefik pod. The volume name will be passed to tpl.
 # This can be used to mount a cert pair or a configmap that holds a config.toml file.
 # After the volume has been mounted, add the configs into traefik by using the `additionalArguments` list below, eg:
@@ -982,7 +992,7 @@ autoscaling:  # @schema additionalProperties: false
   metrics: []
   # -- behavior configures the scaling behavior of the target in both Up and Down directions (scaleUp and scaleDown fields respectively).
   behavior: {}
-  # -- scaleTargetRef points to the target resource to scale, and is used to the pods for which metrics should be collected, as well as to actually change the replica count.
+  # -- scaleTargetRef points to the target resource to scale, and is used for the pods for which metrics should be collected, as well as to actually change the replica count.
   scaleTargetRef:
     apiVersion: apps/v1
     kind: Deployment
@@ -1100,14 +1110,14 @@ podSecurityContext:
 # See #595 for more details and traefik/tests/values/extra.yaml for example.
 extraObjects: []
 
-# -- This field override the default Release Namespace for Helm.
+# -- This field overrides the default Release Namespace for Helm.
 # It will not affect optional CRDs such as `ServiceMonitor` and `PrometheusRules`
 namespaceOverride: ""
 
-# -- This field override the default app.kubernetes.io/instance label for all Objects.
+# -- This field overrides the default app.kubernetes.io/instance label for all Objects.
 instanceLabelOverride: ""
 
-# -- This field override the default version extracted from image.tag
+# -- This field overrides the default version extracted from image.tag
 versionOverride: ""
 
 # Traefik Hub configuration. See https://doc.traefik.io/traefik-hub/
@@ -1197,7 +1207,7 @@ hub:
       partition: ""
       # -- Prefix for consul service tags.
       prefix: "traefik"
-      # -- Interval for check Consul API.
+      # -- Interval for checking Consul API.
       refreshInterval: 15
       # -- Forces the read to be fully consistent.
       requireConsistent: false
@@ -1265,7 +1275,7 @@ hub:
       key: ""
       # -- When insecureSkipVerify is set to true, the TLS connection accepts any certificate presented by the server. Default: false.
       insecureSkipVerify: false
-  # Enable export of errors logs to the platform. Default: true.
+  # Enable export of error logs to the platform. Default: true.
   sendlogs:  # @schema type:[boolean, null]
 
   tracing:
```

## 37.2.0  ![AppVersion: v3.5.3](https://img.shields.io/static/v1?label=AppVersion&message=v3.5.3&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2025-10-21

* feat: support API basePath
* feat: make dashboard toggleable
* feat: :package: support Traefik Hub v3.18 pluginRegistry feature
* feat(traefik-hub): add mcpgateway option
* feat(observability): :mag: add per entrypoint observability
* feat(metrics): :chart_with_upwards_trend: add OTLP resourceAttributes support
* feat(logs): :memo: add missing support of OTLP logs
* chore(release): 🚀 publish traefik 37.2.0
* chore(hub): :twisted_rightwards_arrows: update hub and proxy mapping for v3.18

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index cc48b7d..c697f6f 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -139,6 +139,8 @@ experimental:
   plugins: {}
   # -- Enable experimental local plugins
   localPlugins: {}
+  # -- Enable OTLP logging experimental feature.
+  otlpLogs: false
 
 gateway:
   # -- When providers.kubernetesGateway.enabled, deploy a default gateway
@@ -188,6 +190,12 @@ gatewayClass:  # @schema additionalProperties: false
   # -- Additional gatewayClass labels (e.g. for filtering gateway objects by custom labels)
   labels: {}
 
+api:
+  # -- Enable the dashboard
+  dashboard: true
+  # -- Configure API basePath
+  basePath: ""  # @schema type:[string, null]; default: "/"
+
 # -- Only dashboard & healthcheck IngressRoute are supported. It's recommended to create workloads CR outside of this Chart.
 ingressRoute:
   dashboard:
@@ -370,6 +378,48 @@ logs:
     filePath: ""
     # -- When set to true and format is common, it disables the colorized output.
     noColor: false
+    otlp:
+      # -- Set to true in order to enable OpenTelemetry on logs. Note that experimental.otlpLogs needs to be enabled.
+      enabled: false
+      # -- Service name used in OTLP backend. Default: traefik.
+      serviceName:  # @schema type:[string, null]
+      http:
+        # -- Set to true in order to send logs to the OpenTelemetry Collector using HTTP.
+        enabled: false
+        # -- Format: <scheme>://<host>:<port><path>. Default: https://localhost:4318/v1/logs
+        endpoint: ""
+        # -- Additional headers sent with logs by the reporter to the OpenTelemetry Collector.
+        headers: {}
+        ## Defines the TLS configuration used by the reporter to send logs to the OpenTelemetry Collector.
+        tls:
+          # -- The path to the certificate authority, it defaults to the system bundle.
+          ca: ""
+          # -- The path to the public certificate. When using this option, setting the key option is required.
+          cert: ""
+          # -- The path to the private key. When using this option, setting the cert option is required.
+          key: ""
+          # -- When set to true, the TLS connection accepts any certificate presented by the server regardless of the hostnames it covers.
+          insecureSkipVerify:  # @schema type:[boolean, null]
+      grpc:
+        # -- Set to true in order to send logs  to the OpenTelemetry Collector using gRPC
+        enabled: false
+        # -- Format: <host>:<port>. Default: "localhost:4317"
+        endpoint: ""
+        # -- Allows reporter to send logs to the OpenTelemetry Collector without using a secured protocol.
+        insecure: false
+        ## Defines the TLS configuration used by the reporter to send logs to the OpenTelemetry Collector.
+        tls:
+          # -- The path to the certificate authority, it defaults to the system bundle.
+          ca: ""
+          # -- The path to the public certificate. When using this option, setting the key option is required.
+          cert: ""
+          # -- The path to the private key. When using this option, setting the cert option is required.
+          key: ""
+          # -- When set to true, the TLS connection accepts any certificate presented by the server regardless of the hostnames it covers.
+          insecureSkipVerify: false
+      # -- Defines additional resource attributes to be sent to the collector.
+      resourceAttributes: {}
+
   access:
     # -- To enable access logs
     enabled: false
@@ -401,6 +451,47 @@ logs:
         # -- Set default mode for fields.headers
         defaultmode: drop  # @schema enum:[keep, drop, redact]; default: drop
         names: {}
+    otlp:
+      # -- Set to true in order to enable OpenTelemetry on access logs. Note that experimental.otlpLogs needs to be enabled.
+      enabled: false
+      # -- Service name used in OTLP backend. Default: traefik.
+      serviceName:  # @schema type:[string, null]
+      http:
+        # -- Set to true in order to send access logs to the OpenTelemetry Collector using HTTP.
+        enabled: false
+        # -- Format: <scheme>://<host>:<port><path>. Default: https://localhost:4318/v1/logs
+        endpoint: ""
+        # -- Additional headers sent with access logs by the reporter to the OpenTelemetry Collector.
+        headers: {}
+        ## Defines the TLS configuration used by the reporter to send access logs to the OpenTelemetry Collector.
+        tls:
+          # -- The path to the certificate authority, it defaults to the system bundle.
+          ca: ""
+          # -- The path to the public certificate. When using this option, setting the key option is required.
+          cert: ""
+          # -- The path to the private key. When using this option, setting the cert option is required.
+          key: ""
+          # -- When set to true, the TLS connection accepts any certificate presented by the server regardless of the hostnames it covers.
+          insecureSkipVerify:  # @schema type:[boolean, null]
+      grpc:
+        # -- Set to true in order to send access logs to the OpenTelemetry Collector using gRPC
+        enabled: false
+        # -- Format: <host>:<port>. Default: "localhost:4317"
+        endpoint: ""
+        # -- Allows reporter to send access logs to the OpenTelemetry Collector without using a secured protocol.
+        insecure: false
+        ## Defines the TLS configuration used by the reporter to send access logs to the OpenTelemetry Collector.
+        tls:
+          # -- The path to the certificate authority, it defaults to the system bundle.
+          ca: ""
+          # -- The path to the public certificate. When using this option, setting the key option is required.
+          cert: ""
+          # -- The path to the private key. When using this option, setting the cert option is required.
+          key: ""
+          # -- When set to true, the TLS connection accepts any certificate presented by the server regardless of the hostnames it covers.
+          insecureSkipVerify: false
+      # -- Defines additional resource attributes to be sent to the collector.
+      resourceAttributes: {}
 
 metrics:
   # -- Enable metrics for internal resources. Default: false
@@ -519,7 +610,7 @@ metrics:
     http:
       # -- Set to true in order to send metrics to the OpenTelemetry Collector using HTTP.
       enabled: false
-      # -- Format: <scheme>://<host>:<port><path>. Default: http://localhost:4318/v1/metrics
+      # -- Format: <scheme>://<host>:<port><path>. Default: https://localhost:4318/v1/metrics
       endpoint: ""
       # -- Additional headers sent with metrics by the reporter to the OpenTelemetry Collector.
       headers: {}
@@ -536,7 +627,7 @@ metrics:
     grpc:
       # -- Set to true in order to send metrics to the OpenTelemetry Collector using gRPC
       enabled: false
-      # -- Format: <scheme>://<host>:<port><path>. Default: http://localhost:4318/v1/metrics
+      # -- Format: <host>:<port>. Default: "localhost:4317"
       endpoint: ""
       # -- Allows reporter to send metrics to the OpenTelemetry Collector without using a secured protocol.
       insecure: false
@@ -550,6 +641,8 @@ metrics:
         key: ""
         # -- When set to true, the TLS connection accepts any certificate presented by the server regardless of the hostnames it covers.
         insecureSkipVerify: false
+    # -- Defines additional resource attributes to be sent to the collector.
+    resourceAttributes: {}
 
 ocsp:
   # -- Enable OCSP stapling support.
@@ -581,7 +674,7 @@ tracing:  # @schema additionalProperties: false
     http:
       # -- Set to true in order to send metrics to the OpenTelemetry Collector using HTTP.
       enabled: false
-      # -- Format: <scheme>://<host>:<port><path>. Default: http://localhost:4318/v1/metrics
+      # -- Format: <scheme>://<host>:<port><path>. Default: https://localhost:4318/v1/tracing
       endpoint: ""
       # -- Additional headers sent with metrics by the reporter to the OpenTelemetry Collector.
       headers: {}
@@ -598,7 +691,7 @@ tracing:  # @schema additionalProperties: false
     grpc:
       # -- Set to true in order to send metrics to the OpenTelemetry Collector using gRPC
       enabled: false
-      # -- Format: <scheme>://<host>:<port><path>. Default: http://localhost:4318/v1/metrics
+      # -- Format: <host>:<port>. Default: "localhost:4317"
       endpoint: ""
       # -- Allows reporter to send metrics to the OpenTelemetry Collector without using a secured protocol.
       insecure: false
@@ -670,6 +763,15 @@ ports:
     exposedPort: 8080
     # -- The port protocol (TCP/UDP)
     protocol: TCP
+    observability:  # @schema additionalProperties: false
+      # -- Defines whether a router attached to this EntryPoint produces metrics by default.
+      metrics:  # @schema type:[boolean, null]; default: true
+      # -- Defines whether a router attached to this EntryPoint produces access-logs by default.
+      accessLogs:  # @schema type:[boolean, null]; default: true
+      # -- Defines whether a router attached to this EntryPoint produces traces by default.
+      tracing:  # @schema type:[boolean, null]; default: true
+      # -- Defines the tracing verbosity level for routers attached to this EntryPoint.
+      traceVerbosity:  # @schema enum:[minimal, detailed, null]; type:[string, null]; default: minimal
   web:
     ## -- Enable this entrypoint as a default entrypoint. When a service doesn't explicitly set an entrypoint it will only use this entrypoint.
     # asDefault: true
@@ -711,6 +813,15 @@ ports:
         graceTimeOut:               # @schema type:[string, integer, null]
       keepAliveMaxRequests:         # @schema type:[integer, null]; minimum:0
       keepAliveMaxTime:             # @schema type:[string, integer, null]
+    observability:  # @schema additionalProperties: false
+      # -- Enables metrics for this entryPoint.
+      metrics:  # @schema type:[boolean, null]; default: true
+      # -- Enables access-logs for this entryPoint.
+      accessLogs:  # @schema type:[boolean, null]; default: true
+      # -- Enables tracing for this entryPoint.
+      tracing:  # @schema type:[boolean, null]; default: true
+      # -- Defines the tracing verbosity level for this entryPoint.
+      traceVerbosity:  # @schema enum:[minimal, detailed, null]; type:[string, null]; default: minimal
   websecure:
     ## -- Enable this entrypoint as a default entrypoint. When a service doesn't explicitly set an entrypoint it will only use this entrypoint.
     # asDefault: true
@@ -772,6 +883,15 @@ ports:
     #   - namespace-name1@kubernetescrd
     #   - namespace-name2@kubernetescrd
     middlewares: []
+    observability:  # @schema additionalProperties: false
+      # -- Enables metrics for this entryPoint.
+      metrics:  # @schema type:[boolean, null]; default: true
+      # -- Enables access-logs for this entryPoint.
+      accessLogs:  # @schema type:[boolean, null]; default: true
+      # -- Enables tracing for this entryPoint.
+      tracing:  # @schema type:[boolean, null]; default: true
+      # -- Defines the tracing verbosity level for this entryPoint.
+      traceVerbosity:  # @schema enum:[minimal, detailed, null]; type:[string, null]; default: minimal
   metrics:
     # -- When using hostNetwork, use another port to avoid conflict with node exporter:
     # https://github.com/prometheus/prometheus/wiki/Default-port-allocations
@@ -785,6 +905,15 @@ ports:
     exposedPort: 9100
     # -- The port protocol (TCP/UDP)
     protocol: TCP
+    observability:  # @schema additionalProperties: false
+      # -- Enables metrics for this entryPoint.
+      metrics:  # @schema type:[boolean, null]; default: true
+      # -- Enables access-logs for this entryPoint.
+      accessLogs:  # @schema type:[boolean, null]; default: true
+      # -- Enables tracing for this entryPoint.
+      tracing:  # @schema type:[boolean, null]; default: true
+      # -- Defines the tracing verbosity level for this entryPoint.
+      traceVerbosity:  # @schema enum:[minimal, detailed, null]; type:[string, null]; default: minimal
 
 # -- TLS Options are created as [TLSOption CRDs](https://doc.traefik.io/traefik/https/tls/#tls-options)
 # When using `labelSelector`, you'll need to set labels on tlsOption accordingly.
@@ -1010,6 +1139,12 @@ hub:
       # -- When set to true, it will only accept paths and methods that are explicitly defined in its OpenAPI specification
       validateRequestMethodAndPath: false
 
+  mcpgateway:
+    # -- Set to true in order to enable AI MCP Gateway. Requires a valid license token.
+    enabled: false
+    # -- Hard limit for the size of request bodies inspected by the gateway. Accepts a plain integer representing **bytes**. The default value is `1048576` (1 MiB).
+    maxRequestBodySize:  # @schema type:[integer, null]; minimum:0
+
   aigateway:
     # -- Set to true in order to enable AI Gateway. Requires a valid license token.
     enabled: false
@@ -1148,6 +1283,10 @@ hub:
         # -- Name of the header that will contain the tracestate copy.
         traceState: ""
 
+  # Define private plugin sources
+  pluginRegistry:
+    sources: {}
+
 # -- Required for OCI Marketplace integration.
 # See https://docs.public.content.oci.oraclecloud.com/en-us/iaas/Content/Marketplace/understanding-helm-charts.htm
 oci_meta:
```

## 37.1.2  ![AppVersion: v3.5.3](https://img.shields.io/static/v1?label=AppVersion&message=v3.5.3&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2025-10-03

* fix(observability): tracer creation warning with default security context
* fix(CRDs): ✨ update for Traefik Proxy v3.5.2
* feat(deps): update traefik docker tag to v3.5.3 + add plugin hash option
* feat(CRDs): update for Traefik to v3.5.3
* chore(release): :rocket: publish traefik 37.1.2 and crds 1.11.1

## 37.1.1  ![AppVersion: v3.5.2](https://img.shields.io/static/v1?label=AppVersion&message=v3.5.2&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2025-09-10

* feat(hub): allow to specify admission controller certificate from existing secret
* feat(deps): update traefik docker tag to v3.5.2
* feat(accesslog): ✨ add genericCLF format
* chore(release): 🚀 publish v37.1.1

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index b56a33b..cc48b7d 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -374,7 +374,7 @@ logs:
     # -- To enable access logs
     enabled: false
     # -- Set [access log format](https://doc.traefik.io/traefik/observability/access-logs/#format)
-    format:  # @schema enum:["common", "json", null]; type:[string, null]; default: "common"
+    format:  # @schema enum:["common", "genericCLF", "json", null]; type:[string, null]; default: "common"
     # filePath: "/var/log/traefik/access.log
     # -- Set [bufferingSize](https://doc.traefik.io/traefik/observability/access-logs/#bufferingsize)
     bufferingSize:  # @schema type:[integer, null]
@@ -998,6 +998,8 @@ hub:
       listenAddr: ""
       # -- Certificate name of the WebHook admission server. Default: "hub-agent-cert".
       secretName: "hub-agent-cert"
+      # -- By default, this chart handles directly the tls certificate required for the admission webhook. It's possible to disable this behavior and handle it outside of the chart. See EXAMPLES.md for more details.
+      selfManagedCertificate: false
       # -- Set custom certificate for the WebHook admission server. The certificate should be specified with _tls.crt_ and _tls.key_ in base64 encoding.
       customWebhookCertificate: {}
       # -- Set it to false if you need to disable Traefik Hub pod restart when mutating webhook certificate is updated. It's done with a label update.
```

## 37.1.0  ![AppVersion: v3.5.1](https://img.shields.io/static/v1?label=AppVersion&message=v3.5.1&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2025-09-03

* refactor: remove `$root` hacks in favor of using `$`
* refactor: only render `--global.checkNewVersion` when it differs from default
* fix: prevent blank lines in args
* fix(deployment): allow to disable checkNewVersion via values.yaml
* feat: support custom monitoring api
* feat: support Traefik v3.5 features
* feat(hub): add annotations for webhook admission
* feat(hooks): use now stable prestop command syntax
* feat(deps): update traefik docker tag to v3.5.1
* feat(deployment): add chart value timezone that auomatically configures access logs timezone
* feat(CRDs): update Traefik Hub to v1.21.1
* feat(CRDs): add gatewayAPI experimental channel option
* docs(plugins): Sync VALUES.md
* chore(release): :rocket: Publish 37.1.0 and 1.11.0

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 0c1274f..b56a33b 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -91,8 +91,8 @@ deployment:
   # -- Pod lifecycle actions
   lifecycle: {}
   # preStop:
-  #   exec:
-  #     command: ["/bin/sh", "-c", "sleep 40"]
+  #   sleep:
+  #     seconds: 20
   # postStart:
   #   httpGet:
   #     path: /ping
@@ -161,7 +161,7 @@ gateway:
       hostname: ""
       # Specify expected protocol on this listener. See [ProtocolType](https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.ProtocolType)
       protocol: HTTP
-      # -- Routes are restricted to namespace of the gateway [by default](https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.FromNamespaces
+      # -- (object) Routes are restricted to namespace of the gateway [by default](https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.FromNamespaces
       namespacePolicy:  # @schema type:[object, null]
     # websecure listener is disabled by default because certificateRefs needs to be added,
     # or you may specify TLS protocol with Passthrough mode and add "--providers.kubernetesGateway.experimentalChannel=true" in additionalArguments section.
@@ -378,6 +378,8 @@ logs:
     # filePath: "/var/log/traefik/access.log
     # -- Set [bufferingSize](https://doc.traefik.io/traefik/observability/access-logs/#bufferingsize)
     bufferingSize:  # @schema type:[integer, null]
+    # -- Set [timezone](https://doc.traefik.io/traefik/observability/access-logs/#time-zones)
+    timezone: ""
     # -- Set [filtering](https://docs.traefik.io/observability/access-logs/#filtering)
     filters:  # @schema additionalProperties: false
       # -- Set statusCodes, to limit the access logs to requests with a status codes in the specified range
@@ -432,6 +434,7 @@ metrics:
     serviceMonitor:
       # -- Enable optional CR for Prometheus Operator. See EXAMPLES.md for more details.
       enabled: false
+      apiVersion: "monitoring.coreos.com/v1"
       metricRelabelings: []
       relabelings: []
       jobLabel: ""
@@ -447,6 +450,7 @@ metrics:
     prometheusRule:
       # -- Enable optional CR for Prometheus Operator. See EXAMPLES.md for more details.
       enabled: false
+      apiVersion: "monitoring.coreos.com/v1"
       additionalLabels: {}
       namespace: ""
 
@@ -547,6 +551,13 @@ metrics:
         # -- When set to true, the TLS connection accepts any certificate presented by the server regardless of the hostnames it covers.
         insecureSkipVerify: false
 
+ocsp:
+  # -- Enable OCSP stapling support.
+  # See https://doc.traefik.io/traefik/https/ocsp/#overview
+  enabled: false
+  # -- Defines the OCSP responder URLs to use instead of the one provided by the certificate.
+  responderOverrides: {}
+
 ## Tracing
 # -- https://doc.traefik.io/traefik/observability/tracing/overview/
 tracing:  # @schema additionalProperties: false
@@ -991,6 +1002,8 @@ hub:
       customWebhookCertificate: {}
       # -- Set it to false if you need to disable Traefik Hub pod restart when mutating webhook certificate is updated. It's done with a label update.
       restartOnCertificateChange: true
+      # -- Set custom annotations.
+      annotations: {}
     openApi:
       # -- When set to true, it will only accept paths and methods that are explicitly defined in its OpenAPI specification
       validateRequestMethodAndPath: false
```

## 37.0.0  ![AppVersion: v3.5.0](https://img.shields.io/static/v1?label=AppVersion&message=v3.5.0&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2025-07-29

* fix(observability): allow `tracing.sampleRate` to be set to zero
* fix(entryPoint): allow scheme to be unset on redirect
* fix(Deployment): revision history should be disableable
* feat(podtemplate): add support for localPlugins
* feat(podtemplate): add capacity to set GOMEMLIMIT with default at 90% of user-set memory limit
* feat(hub): offline mode
* feat(gateway-api)!: support selector for namespace policy
* feat(deps): update traefik docker tag to v3.5.0
* feat(CRDs): update for Traefik Proxy v3.5 and Gateway API v1.3.0
* feat(CRDs): update Traefik Hub to v1.21.0
* docs(plugins): improve wording and sync with `VALUES.md`
* chore(release): :rocket: publish 37.0.0 and 1.10.0

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 25a47a9..0c1274f 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -101,6 +101,10 @@ deployment:
   #     scheme: HTTP
   # -- Set a runtimeClassName on pod
   runtimeClassName: ""
+  # -- Percentage of memory limit to set for GOMEMLIMIT
+  # -- set as decimal (0.9 = 90%, 0.95 = 95% etc)
+  # -- only takes effect when resources.limits.memory is set
+  goMemLimitPercentage: 0.9
 
 # -- [Pod Disruption Budget](https://kubernetes.io/docs/reference/kubernetes-api/policy-resources/pod-disruption-budget-v1/)
 podDisruptionBudget:  # @schema additionalProperties: false
@@ -131,11 +135,10 @@ experimental:
   kubernetesGateway:
     # -- Enable traefik experimental GatewayClass CRD
     enabled: false
-  # -- Enable traefik experimental plugins
+  # -- Enable experimental plugins
   plugins: {}
-  # demo:
-  #   moduleName: github.com/traefik/plugindemo
-  #   version: v0.2.1
+  # -- Enable experimental local plugins
+  localPlugins: {}
 
 gateway:
   # -- When providers.kubernetesGateway.enabled, deploy a default gateway
@@ -159,7 +162,7 @@ gateway:
       # Specify expected protocol on this listener. See [ProtocolType](https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.ProtocolType)
       protocol: HTTP
       # -- Routes are restricted to namespace of the gateway [by default](https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.FromNamespaces
-      namespacePolicy:  # @schema type:[string, null]
+      namespacePolicy:  # @schema type:[object, null]
     # websecure listener is disabled by default because certificateRefs needs to be added,
     # or you may specify TLS protocol with Passthrough mode and add "--providers.kubernetesGateway.experimentalChannel=true" in additionalArguments section.
     # websecure:
@@ -301,6 +304,8 @@ providers:  # @schema additionalProperties: false
       pathOverride: ""
     # -- Defines whether to use Native Kubernetes load-balancing mode by default.
     nativeLBByDefault: false
+    # -- Defines whether to make prefix matching strictly comply with the Kubernetes Ingress specification.
+    strictPrefixMatching: false
 
   kubernetesGateway:
     # -- Enable Traefik Gateway provider for Gateway API
@@ -971,7 +976,7 @@ hub:
   # It enables API Gateway.
   token: ""
   # -- Disables all external network connections.
-  offline: false
+  offline:  # @schema type:[boolean, null]
   # -- By default, Traefik Hub provider watches all namespaces. When using `rbac.namespaced`, it will watch helm release namespace and namespaces listed in this array.
   namespaces: []  # @schema required:true
   apimanagement:
```

## 36.3.0  ![AppVersion: v3.4.3](https://img.shields.io/static/v1?label=AppVersion&message=v3.4.3&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2025-07-01

* feat(deps): update traefik docker tag to v3.4.3
* feat(deployment): allow null and 0 replicas
* chore(release): 🚀 publish 36.3.0

## 36.2.0  ![AppVersion: v3.4.1](https://img.shields.io/static/v1?label=AppVersion&message=v3.4.1&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2025-06-23

* fix(CRDs): :bug: kustomization file for CRDs
* feat(hub): ✨ initial support for AI Gateway
* feat(hub): update version mapping with Proxy v3.4
* feat(hpa): ✨ customizable scaleTargetRef
* chore(schema): update linter
* chore(release): 🚀 publish v36.2.0 and CRDs v1.9.0

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index c6daa8d..25a47a9 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -825,10 +825,23 @@ service:
   #   # externalIPs: []
   #   # ipFamilies: [ "IPv4","IPv6" ]
 
-autoscaling:
+autoscaling:  # @schema additionalProperties: false
   # -- Create HorizontalPodAutoscaler object.
   # See EXAMPLES.md for more details.
   enabled: false
+  # -- minReplicas is the lower limit for the number of replicas to which the autoscaler can scale down. It defaults to 1 pod.
+  minReplicas:  # @schema type:[integer, null]; minimum:0
+  # -- maxReplicas is the upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas.
+  maxReplicas:  # @schema type:[integer, null]; minimum:0
+  # -- metrics contains the specifications for which to use to calculate the desired replica count (the maximum replica count across all metrics will be used).
+  metrics: []
+  # -- behavior configures the scaling behavior of the target in both Up and Down directions (scaleUp and scaleDown fields respectively).
+  behavior: {}
+  # -- scaleTargetRef points to the target resource to scale, and is used to the pods for which metrics should be collected, as well as to actually change the replica count.
+  scaleTargetRef:
+    apiVersion: apps/v1
+    kind: Deployment
+    name: "{{ template \"traefik.fullname\" . }}"
 
 persistence:
   # -- Enable persistence using Persistent Volume Claims
@@ -977,9 +990,11 @@ hub:
       # -- When set to true, it will only accept paths and methods that are explicitly defined in its OpenAPI specification
       validateRequestMethodAndPath: false
 
-  experimental:
+  aigateway:
     # -- Set to true in order to enable AI Gateway. Requires a valid license token.
-    aigateway: false
+    enabled: false
+    # -- Hard limit for the size of request bodies inspected by the gateway. Accepts a plain integer representing **bytes**. The default value is `1048576` (1 MiB).
+    maxRequestBodySize:  # @schema type:[integer, null]; minimum:0
   providers:
     consulCatalogEnterprise:
       # -- Enable Consul Catalog Enterprise backend with default settings.
```

## 36.2.0-rc1  ![AppVersion: v3.4.1](https://img.shields.io/static/v1?label=AppVersion&message=v3.4.1&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2025-06-11

* feat(CRDs): update Traefik Hub to v1.20.1
* chore: release traefik v36.2.0-rc1 and traefik-crds v1.9.0-rc1


## 36.1.0  ![AppVersion: v3.4.1](https://img.shields.io/static/v1?label=AppVersion&message=v3.4.1&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2025-06-10

* fix(schema): 🐛 allow additional properties on `global`
* fix(chart): update icon link to track upstream master branch
* fix(Traefik Hub): add strict check on admission cert
* feat(Traefik Hub): add v3.17 version mapping
* chore(release): 🚀 publish v36.1.0

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 9bcb400..c6daa8d 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -597,7 +597,7 @@ tracing:  # @schema additionalProperties: false
         # -- When set to true, the TLS connection accepts any certificate presented by the server regardless of the hostnames it covers.
         insecureSkipVerify: false
 
-global:  # @schema additionalProperties: false
+global:
   checkNewVersion: true
   # -- Please take time to consider whether or not you wish to share anonymous data with us
   # See https://doc.traefik.io/traefik/contributing/data-collection/
```

## 36.0.0  ![AppVersion: v3.4.1](https://img.shields.io/static/v1?label=AppVersion&message=v3.4.1&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2025-06-06

* fix(notes): update condition to display certificate persistence warning
* fix(Traefik Proxy): supported `ingressRoute.*.annotations` breaks templating
* fix(Traefik Proxy)!: strict opt-in on data collection
* feat(deps): update traefik docker tag to v3.4.1
* feat(Traefik Hub): add offline flag
* chore(schema): 🔧 update following latest upstream release
* chore(release): 🚀 publish v36.0.0

**Upgrade Notes**

There is a breaking change on `globalArguments` which has been replaced by `global.xx`, following upstream.
See PR [#1436](https://github.com/traefik/traefik-helm-chart/pull/1436) for details.

Following GDPR, anonymous stats usage has been disabled by default.
Please take time to consider whether or not you wish to share anonymous data to help TraefikLabs improve Traefik Proxy.

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index a9b74a3..9bcb400 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -185,6 +185,7 @@ gatewayClass:  # @schema additionalProperties: false
   # -- Additional gatewayClass labels (e.g. for filtering gateway objects by custom labels)
   labels: {}
 
+# -- Only dashboard & healthcheck IngressRoute are supported. It's recommended to create workloads CR outside of this Chart.
 ingressRoute:
   dashboard:
     # -- Create an IngressRoute for the dashboard
@@ -596,10 +597,25 @@ tracing:  # @schema additionalProperties: false
         # -- When set to true, the TLS connection accepts any certificate presented by the server regardless of the hostnames it covers.
         insecureSkipVerify: false
 
-# -- Global command arguments to be passed to all traefik's pods
-globalArguments:
-  - "--global.checknewversion"
-  - "--global.sendanonymoususage"
+global:  # @schema additionalProperties: false
+  checkNewVersion: true
+  # -- Please take time to consider whether or not you wish to share anonymous data with us
+  # See https://doc.traefik.io/traefik/contributing/data-collection/
+  sendAnonymousUsage: false
+  # -- Required for Azure Marketplace integration.
+  # See https://learn.microsoft.com/en-us/partner-center/marketplace-offers/azure-container-technical-assets-kubernetes?tabs=linux,linux2#update-the-helm-chart
+  azure:
+    # -- Enable specific values for Azure Marketplace
+    enabled: false
+    images:
+      proxy:
+        image: traefik
+        tag: latest
+        registry: docker.io/library
+      hub:
+        image: traefik-hub
+        tag: latest
+        registry: ghcr.io/traefik
 
 # -- Additional arguments to be passed at Traefik's binary
 # See [CLI Reference](https://docs.traefik.io/reference/static-configuration/cli/)
@@ -941,6 +957,8 @@ hub:
   # -- Name of `Secret` with key 'token' set to a valid license token.
   # It enables API Gateway.
   token: ""
+  # -- Disables all external network connections.
+  offline: false
   # -- By default, Traefik Hub provider watches all namespaces. When using `rbac.namespaced`, it will watch helm release namespace and namespaces listed in this array.
   namespaces: []  # @schema required:true
   apimanagement:
@@ -1109,19 +1127,3 @@ oci_meta:
     hub:
       image: traefik-hub
       tag: latest
-
-# -- Required for Azure Marketplace integration.
-# See https://learn.microsoft.com/en-us/partner-center/marketplace-offers/azure-container-technical-assets-kubernetes?tabs=linux,linux2#update-the-helm-chart
-global:
-  azure:
-    # -- Enable specific values for Azure Marketplace
-    enabled: false
-    images:
-      proxy:
-        image: traefik
-        tag: latest
-        registry: docker.io/library
-      hub:
-        image: traefik-hub
-        tag: latest
-        registry: ghcr.io/traefik
```

## 35.4.0  ![AppVersion: v3.4.0](https://img.shields.io/static/v1?label=AppVersion&message=v3.4.0&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2025-05-23

* fix(CRDs): validation check on RootCA for both servertransport
* feat(Traefik Hub): :sparkles: automatically restart API Management pods on admission certificate change
* chore(release): :rocket: publish v35.4.0 and CRDs v1.8.1

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index ff04c8b..a9b74a3 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -953,6 +953,8 @@ hub:
       secretName: "hub-agent-cert"
       # -- Set custom certificate for the WebHook admission server. The certificate should be specified with _tls.crt_ and _tls.key_ in base64 encoding.
       customWebhookCertificate: {}
+      # -- Set it to false if you need to disable Traefik Hub pod restart when mutating webhook certificate is updated. It's done with a label update.
+      restartOnCertificateChange: true
     openApi:
       # -- When set to true, it will only accept paths and methods that are explicitly defined in its OpenAPI specification
       validateRequestMethodAndPath: false
```

## 35.3.0  ![AppVersi   on: v3.4.0](https://img.shields.io/static/v1?label=AppVersion&message=v3.4.0&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2025-05-19

* fix: :bug: ingress route annotations should not be null
* fix(Traefik Hub): fail when upgrading with --reuse-values
* fix(Traefik Hub): custom certificate name for WebHook
* feat: azure marketplace integration
* feat: add serviceName for otlp metrics
* feat(deps): update traefik docker tag to v3.4.0
* feat(deps): update traefik docker tag to v3.3.7
* feat(Traefik Hub): improve UserXP on token
* feat(Traefik Hub): :sparkles: set custom certificate for Hub webhooks
* feat(CRDs): ✨ update CRDs for Traefik Proxy v3.4.x
* chore: update maintainers
* chore: remove K8s version check for unsupported version
* chore(release): :rocket: publish v35.3.0 and CRDs v1.8.0

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 04f4973..ff04c8b 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -115,7 +115,7 @@ ingressClass:  # @schema additionalProperties: false
   name: ""
 
 core:  # @schema additionalProperties: false
-  # -- Can be used to use globally v2 router syntax
+  # -- Can be used to use globally v2 router syntax. Deprecated since v3.4 /!\.
   # See https://doc.traefik.io/traefik/v3.0/migration/v2-to-v3/#new-v3-syntax-notable-changes
   defaultRuleSyntax: ""
 
@@ -504,6 +504,8 @@ metrics:
     explicitBoundaries: []
     # -- Interval at which metrics are sent to the OpenTelemetry Collector. Default: 10s
     pushInterval: ""
+    # -- Service name used in OTLP backend. Default: traefik.
+    serviceName:  # @schema type:[string, null]
     http:
       # -- Set to true in order to send metrics to the OpenTelemetry Collector using HTTP.
       enabled: false
@@ -596,8 +598,8 @@ tracing:  # @schema additionalProperties: false
 
 # -- Global command arguments to be passed to all traefik's pods
 globalArguments:
-- "--global.checknewversion"
-- "--global.sendanonymoususage"
+  - "--global.checknewversion"
+  - "--global.sendanonymoususage"
 
 # -- Additional arguments to be passed at Traefik's binary
 # See [CLI Reference](https://docs.traefik.io/reference/static-configuration/cli/)
@@ -940,15 +942,17 @@ hub:
   # It enables API Gateway.
   token: ""
   # -- By default, Traefik Hub provider watches all namespaces. When using `rbac.namespaced`, it will watch helm release namespace and namespaces listed in this array.
-  namespaces: []
+  namespaces: []  # @schema required:true
   apimanagement:
     # -- Set to true in order to enable API Management. Requires a valid license token.
     enabled: false
     admission:
       # -- WebHook admission server listen address. Default: "0.0.0.0:9943".
       listenAddr: ""
-      # -- Certificate of the WebHook admission server. Default: "hub-agent-cert".
-      secretName: ""
+      # -- Certificate name of the WebHook admission server. Default: "hub-agent-cert".
+      secretName: "hub-agent-cert"
+      # -- Set custom certificate for the WebHook admission server. The certificate should be specified with _tls.crt_ and _tls.key_ in base64 encoding.
+      customWebhookCertificate: {}
     openApi:
       # -- When set to true, it will only accept paths and methods that are explicitly defined in its OpenAPI specification
       validateRequestMethodAndPath: false
@@ -1103,3 +1107,19 @@ oci_meta:
     hub:
       image: traefik-hub
       tag: latest
+
+# -- Required for Azure Marketplace integration.
+# See https://learn.microsoft.com/en-us/partner-center/marketplace-offers/azure-container-technical-assets-kubernetes?tabs=linux,linux2#update-the-helm-chart
+global:
+  azure:
+    # -- Enable specific values for Azure Marketplace
+    enabled: false
+    images:
+      proxy:
+        image: traefik
+        tag: latest
+        registry: docker.io/library
+      hub:
+        image: traefik-hub
+        tag: latest
+        registry: ghcr.io/traefik
```

## 35.2.0  ![AppVersion: v3.3.6](https://img.shields.io/static/v1?label=AppVersion&message=v3.3.6&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2025-04-29

* fix(Traefik Hub): really disable sendlogs when set to false
* fix(Traefik Hub): prefix mutating webhook by release name
* feat(Traefik Hub): option to set token in values
* chore(release): 🚀 publish v35.2.0

## 35.1.0  ![AppVersion: v3.3.6](https://img.shields.io/static/v1?label=AppVersion&message=v3.3.6&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2025-04-25

* feat: ✨ versionOverride
* feat(Traefik Hub): namespaces
* feat(CRDs): remove APIAccess resource
* chore(release): :rocket: publish v35.1.0 and CRDs v1.7.0

**Upgrade Notes**

Traefik-Hub users should follow this procedure:

1. `APIAccess` resources needs to be converted to [`APICatalogItem`](https://doc.traefik.io/traefik-hub/api-management/api-catalogitem) ones
2. run the [usual upgrade procedure](https://github.com/traefik/traefik-helm-chart/blob/master/README.md#upgrading)
3. delete the `APIAccess` CRD by running:

```bash
kubectl delete customresourcedefinitions.apiextensions.k8s.io apiaccesses.hub.traefik.io
```

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index f2b90da..04f4973 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -7,7 +7,8 @@ image:  # @schema additionalProperties: false
   registry: docker.io
   # -- Traefik image repository
   repository: traefik
-  # -- defaults to appVersion
+  # -- defaults to appVersion. It's used for version checking, even prefixed with experimental- or latest-.
+  # When a digest is required, `versionOverride` can be used to set the version.
   tag:  # @schema type:[string, null]
   # -- Traefik image pull policy
   pullPolicy: IfNotPresent
@@ -273,7 +274,7 @@ providers:  # @schema additionalProperties: false
     # -- When the parameter is set, only resources containing an annotation with the same value are processed. Otherwise, resources missing the annotation, having an empty value, or the value traefik are processed. It will also set required annotation on Dashboard and Healthcheck IngressRoute when enabled.
     ingressClass: ""
     # labelSelector: environment=production,method=traefik
-    # -- Array of namespaces to watch. If left empty, Traefik watches all namespaces.
+    # -- Array of namespaces to watch. If left empty, Traefik watches all namespaces. . When using `rbac.namespaced`, it will watch helm release namespace and namespaces listed in this array.
     namespaces: []
     # -- Defines whether to use Native Kubernetes load-balancing mode by default.
     nativeLBByDefault: false
@@ -288,7 +289,7 @@ providers:  # @schema additionalProperties: false
     # -- When ingressClass is set, only Ingresses containing an annotation with the same value are processed. Otherwise, Ingresses missing the annotation, having an empty value, or the value traefik are processed.
     ingressClass:  # @schema type:[string, null]
     # labelSelector: environment=production,method=traefik
-    # -- Array of namespaces to watch. If left empty, Traefik watches all namespaces.
+    # -- Array of namespaces to watch. If left empty, Traefik watches all namespaces. . When using `rbac.namespaced`, it will watch helm release namespace and namespaces listed in this array.
     namespaces: []
     # IP used for Kubernetes Ingress endpoints
     publishedService:
@@ -306,7 +307,7 @@ providers:  # @schema additionalProperties: false
     # -- Toggles support for the Experimental Channel resources (Gateway API release channels documentation).
     # This option currently enables support for TCPRoute and TLSRoute.
     experimentalChannel: false
-    # -- Array of namespaces to watch. If left empty, Traefik watches all namespaces.
+    # -- Array of namespaces to watch. If left empty, Traefik watches all namespaces. . When using `rbac.namespaced`, it will watch helm release namespace and namespaces listed in this array.
     namespaces: []
     # -- A label selector can be defined to filter on specific GatewayClass objects only.
     labelselector: ""
@@ -927,14 +928,19 @@ extraObjects: []
 # It will not affect optional CRDs such as `ServiceMonitor` and `PrometheusRules`
 namespaceOverride: ""
 
-## -- This field override the default app.kubernetes.io/instance label for all Objects.
+# -- This field override the default app.kubernetes.io/instance label for all Objects.
 instanceLabelOverride: ""
 
+# -- This field override the default version extracted from image.tag
+versionOverride: ""
+
 # Traefik Hub configuration. See https://doc.traefik.io/traefik-hub/
 hub:
   # -- Name of `Secret` with key 'token' set to a valid license token.
   # It enables API Gateway.
   token: ""
+  # -- By default, Traefik Hub provider watches all namespaces. When using `rbac.namespaced`, it will watch helm release namespace and namespaces listed in this array.
+  namespaces: []
   apimanagement:
     # -- Set to true in order to enable API Management. Requires a valid license token.
     enabled: false

## 35.0.1  ![AppVersion: v3.3.6](https://img.shields.io/static/v1?label=AppVersion&message=v3.3.6&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2025-04-18

* feat(deps): update traefik docker tag to v3.3.6
* chore(release): :rocket: publish traefik 35.0.1

## 35.0.0  ![AppVersion: v3.3.5](https://img.shields.io/static/v1?label=AppVersion&message=v3.3.5&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2025-04-07

* fix: do not quote protocol on pod ports
* fix(tracing): 🐛 multiple response or request headers
* feat: ✨ Oracle Cloud marketplace integration
* feat(deps): update traefik docker tag to v3.3.5
* feat!: add port name template functions
* chore(release): :rocket: publish traefik 35.0.0
* chore(helpers): :bookmark: update hub proxy corresponding versions
* chore(ci): :bug: should fail on test error

**Upgrade Notes**

This release has been marked as major as it might [modify service and deployment port names](https://github.com/traefik/traefik-helm-chart/pull/1363) (if they use uppercase characters or are longer than 15 characters).
Nevertheless, even in these cases, it should not impact the availability of your endpoints.

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 956000d..f2b90da 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -1082,3 +1082,18 @@ hub:
         traceParent: ""
         # -- Name of the header that will contain the tracestate copy.
         traceState: ""
+
+# -- Required for OCI Marketplace integration.
+# See https://docs.public.content.oci.oraclecloud.com/en-us/iaas/Content/Marketplace/understanding-helm-charts.htm
+oci_meta:
+  # -- Enable specific values for Oracle Cloud Infrastructure
+  enabled: false
+    # -- It needs to be an ocir repo
+  repo: traefik
+  images:
+    proxy:
+      image: traefik
+      tag: latest
+    hub:
+      image: traefik-hub
+      tag: latest
```

## 34.5.0  ![AppVersion: v3.3.4](https://img.shields.io/static/v1?label=AppVersion&message=v3.3.4&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2025-03-31

* fix(gateway): `gateway.namespace` value is ignored
* feat: allow templating the additionalVolumeMounts configuration
* feat(CRDs): 🔧 update Traefik Hub CRDs to v1.17.2
* chore(release): publish crds 1.6.0 and traefik 34.5.0

## 34.4.1  ![AppVersion: v3.3.4](https://img.shields.io/static/v1?label=AppVersion&message=v3.3.4&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2025-02-28

* fix(Traefik Proxy): headerLabels does not exist for metrics.prometheus
* fix(Traefik Hub): add missing consulCatalogEnterprise provider
* feat(deps): update traefik docker tag to v3.3.4
* test(Traefik Proxy): fix metrics header labels
* fix(chart): reorder source urls annotations
* docs(Traefik Proxy): fix VALUES.md generation on prometheus values
* chore(release): 🚀 publish v34.4.1

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 2d8ac73..956000d 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -394,25 +394,27 @@ logs:
         names: {}
 
 metrics:
-  ## -- Enable metrics for internal resources. Default: false
+  # -- Enable metrics for internal resources. Default: false
   addInternals: false
 
-  ## -- Prometheus is enabled by default.
-  ## -- It can be disabled by setting "prometheus: null"
+  ## Prometheus is enabled by default.
+  ## It can be disabled by setting "prometheus: null"
   prometheus:
     # -- Entry point used to expose metrics.
     entryPoint: metrics
-    ## Enable metrics on entry points. Default: true
+    # -- Enable metrics on entry points. Default: true
     addEntryPointsLabels:  # @schema type:[boolean, null]
-    ## Enable metrics on routers. Default: false
+    # -- Enable metrics on routers. Default: false
     addRoutersLabels:  # @schema type:[boolean, null]
-    ## Enable metrics on services. Default: true
+    # -- Enable metrics on services. Default: true
     addServicesLabels:  # @schema type:[boolean, null]
-    ## Buckets for latency metrics. Default="0.1,0.3,1.2,5.0"
+    # -- Buckets for latency metrics. Default="0.1,0.3,1.2,5.0"
     buckets: ""
-    ## When manualRouting is true, it disables the default internal router in
+    # -- When manualRouting is true, it disables the default internal router in
     ## order to allow creating a custom router for prometheus@internal service.
     manualRouting: false
+    # -- Add HTTP header labels to metrics. See EXAMPLES.md or upstream doc for usage.
+    headerLabels: {}  # @schema type:[object, null]
     service:
       # -- Create a dedicated metrics service to use with ServiceMonitor
       enabled: false
@@ -949,6 +951,64 @@ hub:
     # -- Set to true in order to enable AI Gateway. Requires a valid license token.
     aigateway: false
   providers:
+    consulCatalogEnterprise:
+      # -- Enable Consul Catalog Enterprise backend with default settings.
+      enabled: false
+      # -- Use local agent caching for catalog reads.
+      cache: false
+      # -- Enable Consul Connect support.
+      connectAware: false
+      # -- Consider every service as Connect capable by default.
+      connectByDefault: false
+      # -- Constraints is an expression that Traefik matches against the container's labels
+      constraints: ""
+      # -- Default rule.
+      defaultRule: "Host(`{{ normalize .Name }}`)"
+      endpoint:
+        # -- The address of the Consul server
+        address: ""
+        # -- Data center to use. If not provided, the default agent data center is used
+        datacenter: ""
+        # -- WaitTime limits how long a Watch will block. If not provided, the agent default
+        endpointWaitTime: 0
+        httpauth:
+          # -- Basic Auth password
+          password: ""
+          # -- Basic Auth username
+          username: ""
+        # -- The URI scheme for the Consul server
+        scheme: ""
+        tls:
+          # -- TLS CA
+          ca: ""
+          # -- TLS cert
+          cert: ""
+          # -- TLS insecure skip verify
+          insecureSkipVerify: false
+          # -- TLS key
+          key: ""
+        # -- Token is used to provide a per-request ACL token which overrides the agent's
+        token: ""
+      # -- Expose containers by default.
+      exposedByDefault: true
+      # -- Sets the namespaces used to discover services (Consul Enterprise only).
+      namespaces: ""
+      # -- Sets the partition used to discover services (Consul Enterprise only).
+      partition: ""
+      # -- Prefix for consul service tags.
+      prefix: "traefik"
+      # -- Interval for check Consul API.
+      refreshInterval: 15
+      # -- Forces the read to be fully consistent.
+      requireConsistent: false
+      # -- Name of the Traefik service in Consul Catalog (needs to be registered via the
+      serviceName: "traefik"
+      # -- Use stale consistency for catalog reads.
+      stale: false
+      # -- A list of service health statuses to allow taking traffic.
+      strictChecks: "passing, warning"
+      # -- Watch Consul API events.
+      watch: false
     microcks:
       # -- Enable Microcks provider.
       enabled: false
@@ -1007,6 +1067,7 @@ hub:
       insecureSkipVerify: false
   # Enable export of errors logs to the platform. Default: true.
   sendlogs:  # @schema type:[boolean, null]
+
   tracing:
     # -- Tracing headers to duplicate.
     # To configure the following, tracing.otlp.enabled needs to be set to true.
```

## 34.4.0  ![AppVersion: v3.3.3](https://img.shields.io/static/v1?label=AppVersion&message=v3.3.3&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2025-02-19

* feat(CRDs): update Traefik Hub CRDs to v1.17.0
* chore(release): 🚀 publish v34.4.0 and CRDs v1.4.0

## 34.3.0  ![AppVersion: v3.3.3](https://img.shields.io/static/v1?label=AppVersion&message=v3.3.3&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* fix(Traefik Hub): AIServices are available in API Gateway
* fix(Traefik Hub): :bug: handle main and latest build
* feat: :sparkles: add missing microcks provider for Hub
* feat(deps): update traefik docker tag to v3.3.3
* chore: update CRDs to v1.14.1

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 4f93924..2d8ac73 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -948,6 +948,34 @@ hub:
   experimental:
     # -- Set to true in order to enable AI Gateway. Requires a valid license token.
     aigateway: false
+  providers:
+    microcks:
+      # -- Enable Microcks provider.
+      enabled: false
+      auth:
+        # -- Microcks API client ID.
+        clientId: ""
+        # -- Microcks API client secret.
+        clientSecret: ""
+        # -- Microcks API endpoint.
+        endpoint: ""
+        # -- Microcks API token.
+        token: ""
+      # -- Microcks API endpoint.
+      endpoint: ""
+      # -- Polling interval for Microcks API.
+      pollInterval: 30
+      # -- Polling timeout for Microcks API.
+      pollTimeout: 5
+      tls:
+        # -- TLS CA
+        ca: ""
+        # -- TLS cert
+        cert: ""
+        # -- TLS insecure skip verify
+        insecureSkipVerify: false
+        # -- TLS key
+        key: ""
   redis:
     # -- Enable Redis Cluster. Default: true.
     cluster:    # @schema type:[boolean, null]
```

## 34.2.0  ![AppVersion: v3.3.2](https://img.shields.io/static/v1?label=AppVersion&message=v3.3.2&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2025-01-28

* fix: :bug: permanent redirect should be disableable
* feat: :sparkles: add hub tracing capabilities
* docs: 📚️ fix typo in Guidelines.md
* chore(release): publish v34.2.0

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 3335e78..4f93924 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -979,3 +979,17 @@ hub:
       insecureSkipVerify: false
   # Enable export of errors logs to the platform. Default: true.
   sendlogs:  # @schema type:[boolean, null]
+  tracing:
+    # -- Tracing headers to duplicate.
+    # To configure the following, tracing.otlp.enabled needs to be set to true.
+    additionalTraceHeaders:
+      enabled: false
+      traceContext:
+        # -- Name of the header that will contain the parent-id header copy.
+        parentId: ""
+        # -- Name of the header that will contain the trace-id copy.
+        traceId: ""
+        # -- Name of the header that will contain the traceparent copy.
+        traceParent: ""
+        # -- Name of the header that will contain the tracestate copy.
+        traceState: ""
```

## 34.1.0  ![AppVersion: v3.3.2](https://img.shields.io/static/v1?label=AppVersion&message=v3.3.2&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2025-01-15

* fix(Traefik Proxy): support uppercase letters in entrypoint names
* feat(deps): update traefik docker tag to v3.3.2
* feat(Traefik Hub): add OAS validateRequestMethodAndPath - CRDs update
* chore(release): publish v34.1.0 and CRDs v1.2.0

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index f5922fe..3335e78 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -941,6 +941,10 @@ hub:
       listenAddr: ""
       # -- Certificate of the WebHook admission server. Default: "hub-agent-cert".
       secretName: ""
+    openApi:
+      # -- When set to true, it will only accept paths and methods that are explicitly defined in its OpenAPI specification
+      validateRequestMethodAndPath: false
+
   experimental:
     # -- Set to true in order to enable AI Gateway. Requires a valid license token.
     aigateway: false
```

## 34.0.0  ![AppVersion: v3.3.1](https://img.shields.io/static/v1?label=AppVersion&message=v3.3.1&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2025-01-13

* fix(Traefik Proxy)!: use namespaceOverride as expected
* fix(Traefik Proxy)!: move redirectTo => redirections
* fix(Gateway API): status should not use default service when it's disabled
* feat(deps): update traefik docker tag to v3.3.1
* feat(deps): update traefik docker tag to v3.2.3
* feat(Traefik Proxy): apply migration guide to v3.3
* feat(Traefik Proxy): add support for experimental FastProxy
* feat(Traefik Hub): add support for AI Gateway
* feat(Chart): :package: add optional separated chart for CRDs
* feat(CRDs): update CRDs for Traefik Proxy v3.3.x
* chore(release): publish v34.0.0
* chore(Gateway API): :recycle: remove template from values

**Upgrade Notes**

There are multiple breaking changes in this release:

1. When using namespaceOverride, the label selector will be changed. On a production environment, it's recommended to deploy a new instance with the new version, switch the traffic to it and delete the previous one. See PR #1290 for more information
2. `ports.x.redirectTo` has been refactored to be aligned with upstream syntax. See PR #1301 for a complete before / after example.


### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 78c8ea4..f5922fe 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -122,14 +122,19 @@ core:  # @schema additionalProperties: false
 experimental:
   # -- Defines whether all plugins must be loaded successfully for Traefik to start
   abortOnPluginFailure: false
+  fastProxy:
+    # -- Enables the FastProxy implementation.
+    enabled: false
+    # -- Enable debug mode for the FastProxy implementation.
+    debug: false
+  kubernetesGateway:
+    # -- Enable traefik experimental GatewayClass CRD
+    enabled: false
   # -- Enable traefik experimental plugins
   plugins: {}
   # demo:
   #   moduleName: github.com/traefik/plugindemo
   #   version: v0.2.1
-  kubernetesGateway:
-    # -- Enable traefik experimental GatewayClass CRD
-    enabled: false
 
 gateway:
   # -- When providers.kubernetesGateway.enabled, deploy a default gateway
@@ -314,8 +319,9 @@ providers:  # @schema additionalProperties: false
       hostname: ""
       # -- The Kubernetes service to copy status addresses from. When using third parties tools like External-DNS, this option can be used to copy the service loadbalancer.status (containing the service's endpoints IPs) to the gateways. Default to Service of this Chart.
       service:
-        name: "{{ (include \"traefik.fullname\" .) }}"
-        namespace: "{{ .Release.Namespace }}"
+        enabled: true
+        name: ""
+        namespace: ""
 
   file:
     # -- Create a file provider
@@ -537,8 +543,8 @@ tracing:  # @schema additionalProperties: false
   addInternals: false
   # -- Service name used in selected backend. Default: traefik.
   serviceName:  # @schema type:[string, null]
-  # -- Applies a list of shared key:value attributes on all spans.
-  globalAttributes: {}
+  # -- Defines additional resource attributes to be sent to the collector.
+  resourceAttributes: {}
   # -- Defines the list of request headers to add as attributes. It applies to client and server kind spans.
   capturedRequestHeaders: []
   # -- Defines the list of response headers to add as attributes. It applies to client and server kind spans.
@@ -642,10 +648,12 @@ ports:
     protocol: TCP
     # -- See [upstream documentation](https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport)
     nodePort:  # @schema type:[integer, null]; minimum:0
-    # Port Redirections
-    # Added in 2.2, you can make permanent redirects via entrypoints.
-    # https://docs.traefik.io/routing/entrypoints/#redirection
-    redirectTo: {}
+    redirections:
+      # -- Port Redirections
+      # Added in 2.2, one can make permanent redirects via entrypoints.
+      # Same sets of parameters: to, scheme, permanent and priority.
+      # https://docs.traefik.io/routing/entrypoints/#redirection
+      entryPoint: {}
     forwardedHeaders:
       # -- Trust forwarded headers information (X-Forwarded-*).
       trustedIPs: []
@@ -869,7 +877,7 @@ affinity: {}
 #      - labelSelector:
 #          matchLabels:
 #            app.kubernetes.io/name: '{{ template "traefik.name" . }}'
-#            app.kubernetes.io/instance: '{{ .Release.Name }}-{{ .Release.Namespace }}'
+#            app.kubernetes.io/instance: '{{ .Release.Name }}-{{ include "traefik.namespace" . }}'
 #        topologyKey: kubernetes.io/hostname
 
 # -- nodeSelector is the simplest recommended form of node selection constraint.
@@ -933,7 +941,9 @@ hub:
       listenAddr: ""
       # -- Certificate of the WebHook admission server. Default: "hub-agent-cert".
       secretName: ""
-
+  experimental:
+    # -- Set to true in order to enable AI Gateway. Requires a valid license token.
+    aigateway: false
   redis:
     # -- Enable Redis Cluster. Default: true.
     cluster:    # @schema type:[boolean, null]
```

## 33.2.1  ![AppVersion: v3.2.2](https://img.shields.io/static/v1?label=AppVersion&message=v3.2.2&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2024-12-13

* fix(Gateway API): CRDs should only be defined once
* chore(release): 🚀 publish v33.2.1

## 33.2.0  ![AppVersion: v3.2.2](https://img.shields.io/static/v1?label=AppVersion&message=v3.2.2&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2024-12-11

* fix(Traefik Proxy): :bug: abortOnPluginFailure not released yet
* feat(deps): update traefik docker tag to v3.2.2
* feat(Traefik Proxy): support NativeLB option in GatewayAPI provider
* feat(Traefik Proxy): add `tracing`parameters to helm chart values
* feat(Traefik Proxy): :art: harmonize semverCompare calls
* feat(Gateway API): update sigs.k8s.io/gateway-api to v1.2.1
* chore(release): 🚀 publish v33.2.0

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 9b4379c..78c8ea4 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -305,6 +305,8 @@ providers:  # @schema additionalProperties: false
     namespaces: []
     # -- A label selector can be defined to filter on specific GatewayClass objects only.
     labelselector: ""
+    # -- Defines whether to use Native Kubernetes load-balancing mode by default.
+    nativeLBByDefault: false
     statusAddress:
       # -- This IP will get copied to the Gateway status.addresses, and currently only supports one IP value (IPv4 or IPv6).
       ip: ""
@@ -533,6 +535,18 @@ metrics:
 tracing:  # @schema additionalProperties: false
   # -- Enables tracing for internal resources. Default: false.
   addInternals: false
+  # -- Service name used in selected backend. Default: traefik.
+  serviceName:  # @schema type:[string, null]
+  # -- Applies a list of shared key:value attributes on all spans.
+  globalAttributes: {}
+  # -- Defines the list of request headers to add as attributes. It applies to client and server kind spans.
+  capturedRequestHeaders: []
+  # -- Defines the list of response headers to add as attributes. It applies to client and server kind spans.
+  capturedResponseHeaders: []
+  # -- By default, all query parameters are redacted. Defines the list of query parameters to not redact.
+  safeQueryParams: []
+  # -- The proportion of requests to trace, specified between 0.0 and 1.0. Default: 1.0.
+  sampleRate:  # @schema type:[number, null]; minimum:0; maximum:1
   otlp:
     # -- See https://doc.traefik.io/traefik/v3.0/observability/tracing/opentelemetry/
     enabled: false
```

## 33.1.0  ![AppVersion: v3.2.1](https://img.shields.io/static/v1?label=AppVersion&message=v3.2.1&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2024-11-26

* fix: :bug: support specifying plugins storage
* fix(Traefik): support for entrypoint option on allowACMEByPass
* fix(Traefik Proxy): allowEmptyServices not disabled when set to false
* fix(Traefik Hub): compatibility with Traefik Proxy v3.2
* fix(KubernetesCRD): 🐛 IngressClass should be readable even when kubernetesIngress is disabled
* feat(deps): update traefik docker tag to v3.2.1
* feat(Traefik Proxy): add `abortOnPluginFailure` field
* feat(Traefik Hub): add APICatalogItem and ManagedSubscription support
* docs: 📚️ fix typos in values and readme
* chore(release): 🚀 publish v33.1.0-rc1

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index be89b00..9b4379c 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -120,6 +120,8 @@ core:  # @schema additionalProperties: false
 
 # Traefik experimental features
 experimental:
+  # -- Defines whether all plugins must be loaded successfully for Traefik to start
+  abortOnPluginFailure: false
   # -- Enable traefik experimental plugins
   plugins: {}
   # demo:
@@ -807,7 +809,7 @@ persistence:
 certificatesResolvers: {}
 
 # -- If hostNetwork is true, runs traefik in the host network namespace
-# To prevent unschedulabel pods due to port collisions, if hostNetwork=true
+# To prevent unschedulable pods due to port collisions, if hostNetwork=true
 # and replicas>1, a pod anti-affinity is recommended and will be set if the
 # affinity is left as default.
 hostNetwork: false
```


## 33.1.0-rc1  ![AppVersion: v3.2.1](https://img.shields.io/static/v1?label=AppVersion&message=v3.2.1&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2024-11-26

* fix: :bug: support specifying plugins storage
* fix(Traefik): support for entrypoint option on allowACMEByPass
* fix(Traefik Proxy): allowEmptyServices not disabled when set to false
* fix(Traefik Hub): compatibility with Traefik Proxy v3.2
* fix(KubernetesCRD): 🐛 IngressClass should be readable even when kubernetesIngress is disabled
* feat(deps): update traefik docker tag to v3.2.1
* feat(Traefik Proxy): add `abortOnPluginFailure` field
* feat(Traefik Hub): add APICatalogItem and ManagedSubscription support
* docs: 📚️ fix typos in values and readme
* chore(release): 🚀 publish v33.1.0-rc1

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index be89b00..9b4379c 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -120,6 +120,8 @@ core:  # @schema additionalProperties: false
 
 # Traefik experimental features
 experimental:
+  # -- Defines whether all plugins must be loaded successfully for Traefik to start
+  abortOnPluginFailure: false
   # -- Enable traefik experimental plugins
   plugins: {}
   # demo:
@@ -807,7 +809,7 @@ persistence:
 certificatesResolvers: {}
 
 # -- If hostNetwork is true, runs traefik in the host network namespace
-# To prevent unschedulabel pods due to port collisions, if hostNetwork=true
+# To prevent unschedulable pods due to port collisions, if hostNetwork=true
 # and replicas>1, a pod anti-affinity is recommended and will be set if the
 # affinity is left as default.
 hostNetwork: false
```

## 33.0.0  ![AppVersion: v3.2.0](https://img.shields.io/static/v1?label=AppVersion&message=v3.2.0&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2024-10-30

* fix: 🐛 http3 with internal service
* fix: use correct children indentation for logs.access.filters
* fix(schema): :bug: targetPort can also be a string
* fix(certificateResolvers)!: :boom: :bug: use same syntax in Chart and in Traefik
* fix(Traefik)!: :boom: set 8080 as default port for `traefik` entrypoint
* fix(Traefik Hub): RBAC for distributedAcme
* fix(Kubernetes Ingress)!: :boom: :sparkles: enable publishedService by default
* fix(Gateway API): :bug: add missing required RBAC for v3.2 with experimental Channel
* fix(Env Variables)!: allow extending env without overwrite
* feat(deps): update traefik docker tag to v3.2.0
* feat(deps): update traefik docker tag to v3.1.6
* feat(Traefik): ✨ support Gateway API statusAddress
* feat(Traefik Proxy): CRDs for v3.2+
* feat(Gateway API): :sparkles: standard install CRD v1.2.0
* feat(Gateway API): :sparkles: add infrastructure in the values
* chore: allow TRACE log level
* chore(release): 🚀 publish v33.0.0
* Update topology spread constraints comments

**Upgrade Notes**

There are multiple breaking changes in this release:

1. The default port of `traefik` entrypoint has changed from `9000` to `8080`, just like the Traefik Proxy default port
    * You _may_ have to update probes accordingly (or set this port back to 9000)
2. `publishedService` is enabled by default on Ingress provider
    * You _can_ disable it, if needed
3. The `POD_NAME` and `POD_NAMESPACE` environment variables are now set by default, without values.
    * It is no longer necessary to add them in values and so, it can be removed from user values.
4. In _values_, **certResolvers** specific syntax has been reworked to align with Traefik Proxy syntax.
    * PR [#1214](https://github.com/traefik/traefik-helm-chart/pull/1214) contains a complete before / after example on how to update _values_
5. Traefik Proxy 3.2 supports Gateway API v1.2 (standard channel)
    * It is recommended to check that other software using Gateway API on your cluster are compatible
    * There is a breaking change on CRD upgrade in this version, it _may_ do not work out of the box
    * See [release notes](https://github.com/kubernetes-sigs/gateway-api/releases/tag/v1.2.0) of gateway API v1.2 on how to upgrade their CRDs and avoid issues about invalid values on v1alpha2 version

The CRDs needs to be updated, as documented in the README.

:information_source: A separate helm chart, just for CRDs, is being considered for a future release. See PR [#1123](https://github.com/traefik/traefik-helm-chart/pull/1223)

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 73371f3..be89b00 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -95,7 +95,7 @@ deployment:
   # postStart:
   #   httpGet:
   #     path: /ping
-  #     port: 9000
+  #     port: 8080
   #     host: localhost
   #     scheme: HTTP
   # -- Set a runtimeClassName on pod
@@ -138,6 +138,8 @@ gateway:
   namespace: ""
   # -- Additional gateway annotations (e.g. for cert-manager.io/issuer)
   annotations: {}
+  # -- [Infrastructure](https://kubernetes.io/blog/2023/11/28/gateway-api-ga/#gateway-infrastructure-labels)
+  infrastructure: {}
   # -- Define listeners
   listeners:
     web:
@@ -283,10 +285,11 @@ providers:  # @schema additionalProperties: false
     namespaces: []
     # IP used for Kubernetes Ingress endpoints
     publishedService:
-      enabled: false
-      # Published Kubernetes Service to copy status from. Format: namespace/servicename
-      # By default this Traefik service
-      # pathOverride: ""
+      # -- Enable [publishedService](https://doc.traefik.io/traefik/providers/kubernetes-ingress/#publishedservice)
+      enabled: true
+      # -- Override path of Kubernetes Service used to copy status from. Format: namespace/servicename.
+      # Default to Service deployed with this Chart.
+      pathOverride: ""
     # -- Defines whether to use Native Kubernetes load-balancing mode by default.
     nativeLBByDefault: false
 
@@ -300,6 +303,15 @@ providers:  # @schema additionalProperties: false
     namespaces: []
     # -- A label selector can be defined to filter on specific GatewayClass objects only.
     labelselector: ""
+    statusAddress:
+      # -- This IP will get copied to the Gateway status.addresses, and currently only supports one IP value (IPv4 or IPv6).
+      ip: ""
+      # -- This Hostname will get copied to the Gateway status.addresses.
+      hostname: ""
+      # -- The Kubernetes service to copy status addresses from. When using third parties tools like External-DNS, this option can be used to copy the service loadbalancer.status (containing the service's endpoints IPs) to the gateways. Default to Service of this Chart.
+      service:
+        name: "{{ (include \"traefik.fullname\" .) }}"
+        namespace: "{{ .Release.Namespace }}"
 
   file:
     # -- Create a file provider
@@ -335,8 +347,8 @@ logs:
     # -- Set [logs format](https://doc.traefik.io/traefik/observability/logs/#format)
     format:  # @schema enum:["common", "json", null]; type:[string, null]; default: "common"
     # By default, the level is set to INFO.
-    # -- Alternative logging levels are DEBUG, PANIC, FATAL, ERROR, WARN, and INFO.
-    level: "INFO"  # @schema enum:[INFO,WARN,ERROR,FATAL,PANIC,DEBUG]; default: "INFO"
+    # -- Alternative logging levels are TRACE, DEBUG, INFO, WARN, ERROR, FATAL, and PANIC.
+    level: "INFO"  # @schema enum:[TRACE,DEBUG,INFO,WARN,ERROR,FATAL,PANIC]; default: "INFO"
     # -- To write the logs into a log file, use the filePath option.
     filePath: ""
     # -- When set to true and format is common, it disables the colorized output.
@@ -350,10 +362,13 @@ logs:
     # -- Set [bufferingSize](https://doc.traefik.io/traefik/observability/access-logs/#bufferingsize)
     bufferingSize:  # @schema type:[integer, null]
     # -- Set [filtering](https://docs.traefik.io/observability/access-logs/#filtering)
-    filters: {}
-    statuscodes: ""
-    retryattempts: false
-    minduration: ""
+    filters:  # @schema additionalProperties: false
+      # -- Set statusCodes, to limit the access logs to requests with a status codes in the specified range
+      statuscodes: ""
+      # -- Set retryAttempts, to keep the access logs when at least one retry has happened
+      retryattempts: false
+      # -- Set minDuration, to keep access logs when requests take longer than the specified duration
+      minduration: ""
     # -- Enables accessLogs for internal resources. Default: false.
     addInternals: false
     fields:
@@ -566,24 +581,16 @@ additionalArguments: []
 #  - "--providers.kubernetesingress.ingressclass=traefik-internal"
 #  - "--log.level=DEBUG"
 
-# -- Environment variables to be passed to Traefik's binary
+# -- Additional Environment variables to be passed to Traefik's binary
 # @default -- See _values.yaml_
-env:
-- name: POD_NAME
-  valueFrom:
-    fieldRef:
-      fieldPath: metadata.name
-- name: POD_NAMESPACE
-  valueFrom:
-    fieldRef:
-      fieldPath: metadata.namespace
+env: []
 
 # -- Environment variables to be passed to Traefik's binary from configMaps or secrets
 envFrom: []
 
 ports:
   traefik:
-    port: 9000
+    port: 8080
     # -- Use hostPort if set.
     hostPort:  # @schema type:[integer, null]; minimum:0
     # -- Use hostIP if set. If not set, Kubernetes will default to 0.0.0.0, which
@@ -601,7 +608,7 @@ ports:
     expose:
       default: false
     # -- The exposed port for this service
-    exposedPort: 9000
+    exposedPort: 8080
     # -- The port protocol (TCP/UDP)
     protocol: TCP
   web:
@@ -614,7 +621,7 @@ ports:
       default: true
     exposedPort: 80
     ## -- Different target traefik port on the cluster, useful for IP type LB
-    targetPort:  # @schema type:[integer, null]; minimum:0
+    targetPort:  # @schema type:[string, integer, null]; minimum:0
     # The port protocol (TCP/UDP)
     protocol: TCP
     # -- See [upstream documentation](https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport)
@@ -653,7 +660,7 @@ ports:
       default: true
     exposedPort: 443
     ## -- Different target traefik port on the cluster, useful for IP type LB
-    targetPort:  # @schema type:[integer, null]; minimum:0
+    targetPort:  # @schema type:[string, integer, null]; minimum:0
     ## -- The port protocol (TCP/UDP)
     protocol: TCP
     # -- See [upstream documentation](https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport)
@@ -780,8 +787,8 @@ autoscaling:
 
 persistence:
   # -- Enable persistence using Persistent Volume Claims
-  # ref: http://kubernetes.io/docs/user-guide/persistent-volumes/
-  # It can be used to store TLS certificates, see `storage` in certResolvers
+  # ref: http://kubernetes.io/docs/user-guide/persistent-volumes/.
+  # It can be used to store TLS certificates along with `certificatesResolvers.<name>.acme.storage`  option
   enabled: false
   name: data
   existingClaim: ""
@@ -797,7 +804,7 @@ persistence:
 # -- Certificates resolvers configuration.
 # Ref: https://doc.traefik.io/traefik/https/acme/#certificate-resolvers
 # See EXAMPLES.md for more details.
-certResolvers: {}
+certificatesResolvers: {}
 
 # -- If hostNetwork is true, runs traefik in the host network namespace
 # To prevent unschedulabel pods due to port collisions, if hostNetwork=true
@@ -860,7 +867,7 @@ topologySpreadConstraints: []
 # on nodes where no other traefik pods are scheduled.
 #  - labelSelector:
 #      matchLabels:
-#        app: '{{ template "traefik.name" . }}'
+#        app.kubernetes.io/name: '{{ template "traefik.name" . }}'
 #    maxSkew: 1
 #    topologyKey: kubernetes.io/hostname
 #    whenUnsatisfiable: DoNotSchedule
```

## 32.1.0  ![AppVersion: v3.1.5](https://img.shields.io/static/v1?label=AppVersion&message=v3.1.5&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2024-10-04

* fix: :bug: set disableIngressClassLookup until 3.1.4
* feat(deps): update traefik docker tag to v3.1.5
* feat(Traefik Proxy): update rbac following v3.2 migration guide
* chore(release): 🚀 publish v32.1.0

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index f36a9dd..73371f3 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -809,9 +809,9 @@ hostNetwork: false
 rbac:  # @schema additionalProperties: false
   enabled: true
   # When set to true:
-  # 1. Use `Role` and `RoleBinding` instead of `ClusterRole` and `ClusterRoleBinding`.
-  # 2. Set `disableIngressClassLookup` on Kubernetes Ingress providers with Traefik Proxy v3 until v3.1.1
-  # 3. Set `disableClusterScopeResources` on Kubernetes Ingress and CRD providers with Traefik Proxy v3.1.2+
+  # 1. It switches respectively the use of `ClusterRole` and `ClusterRoleBinding` to `Role` and `RoleBinding`.
+  # 2. It adds `disableIngressClassLookup` on Kubernetes Ingress with Traefik Proxy v3 until v3.1.4
+  # 3. It adds `disableClusterScopeResources` on Ingress and CRD (Kubernetes) providers with Traefik Proxy v3.1.2+
   # **NOTE**: `IngressClass`, `NodePortLB` and **Gateway** provider cannot be used with namespaced RBAC.
   # See [upstream documentation](https://doc.traefik.io/traefik/providers/kubernetes-ingress/#disableclusterscoperesources) for more details.
   namespaced: false

## 32.0.0  ![AppVersion: v3.1.4](https://img.shields.io/static/v1?label=AppVersion&message=v3.1.4&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2024-09-27

* chore(release): :rocket: publish 32.0.0
* fix: replace `CLF` with `common` in `values.yaml`
* feat(Traefik Hub): add APIPlans and APIBundles CRDs

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 51dec67..f36a9dd 100644
index d5173dc..f36a9dd 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -345,7 +345,7 @@ logs:
     # -- To enable access logs
     enabled: false
     # -- Set [access log format](https://doc.traefik.io/traefik/observability/access-logs/#format)
-    format:  # @schema enum:["CLF", "json", null]; type:[string, null]; default: "CLF"
+    format:  # @schema enum:["common", "json", null]; type:[string, null]; default: "common"
     # filePath: "/var/log/traefik/access.log
     # -- Set [bufferingSize](https://doc.traefik.io/traefik/observability/access-logs/#bufferingsize)
     bufferingSize:  # @schema type:[integer, null]
@@ -911,35 +911,34 @@ hub:
       # -- Certificate of the WebHook admission server. Default: "hub-agent-cert".
       secretName: ""
 
-  ratelimit:
-    redis:
-      # -- Enable Redis Cluster. Default: true.
-      cluster:    # @schema type:[boolean, null]
-      # -- Database used to store information. Default: "0".
-      database:   # @schema type:[string, null]
-      # -- Endpoints of the Redis instances to connect to. Default: "".
-      endpoints: ""
-      # -- The username to use when connecting to Redis endpoints. Default: "".
+  redis:
+    # -- Enable Redis Cluster. Default: true.
+    cluster:    # @schema type:[boolean, null]
+    # -- Database used to store information. Default: "0".
+    database:   # @schema type:[string, null]
+    # -- Endpoints of the Redis instances to connect to. Default: "".
+    endpoints: ""
+    # -- The username to use when connecting to Redis endpoints. Default: "".
+    username: ""
+    # -- The password to use when connecting to Redis endpoints. Default: "".
+    password: ""
+    sentinel:
+      # -- Name of the set of main nodes to use for main selection. Required when using Sentinel. Default: "".
+      masterset: ""
+      # -- Username to use for sentinel authentication (can be different from endpoint username). Default: "".
       username: ""
-      # -- The password to use when connecting to Redis endpoints. Default: "".
+      # -- Password to use for sentinel authentication (can be different from endpoint password). Default: "".
       password: ""
-      sentinel:
-        # -- Name of the set of main nodes to use for main selection. Required when using Sentinel. Default: "".
-        masterset: ""
-        # -- Username to use for sentinel authentication (can be different from endpoint username). Default: "".
-        username: ""
-        # -- Password to use for sentinel authentication (can be different from endpoint password). Default: "".
-        password: ""
-      # -- Timeout applied on connection with redis. Default: "0s".
-      timeout: ""
-      tls:
-        # -- Path to the certificate authority used for the secured connection.
-        ca: ""
-        # -- Path to the public certificate used for the secure connection.
-        cert: ""
-        # -- Path to the private key used for the secure connection.
-        key: ""
-        # -- When insecureSkipVerify is set to true, the TLS connection accepts any certificate presented by the server. Default: false.
-        insecureSkipVerify: false
+    # -- Timeout applied on connection with redis. Default: "0s".
+    timeout: ""
+    tls:
+      # -- Path to the certificate authority used for the secured connection.
+      ca: ""
+      # -- Path to the public certificate used for the secure connection.
+      cert: ""
+      # -- Path to the private key used for the secure connection.
+      key: ""
+      # -- When insecureSkipVerify is set to true, the TLS connection accepts any certificate presented by the server. Default: false.
+      insecureSkipVerify: false
   # Enable export of errors logs to the platform. Default: true.
   sendlogs:  # @schema type:[boolean, null]
```

## 32.0.0-rc1  ![AppVersion: v3.1.4](https://img.shields.io/static/v1?label=AppVersion&message=v3.1.4&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2024-09-20

* feat(Traefik Hub): add APIPlans and APIBundles CRDs
* chore(release): 🚀 publish 32.0.0-rc1

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index d5173dc..51dec67 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -911,35 +911,34 @@ hub:
       # -- Certificate of the WebHook admission server. Default: "hub-agent-cert".
       secretName: ""
 
-  ratelimit:
-    redis:
-      # -- Enable Redis Cluster. Default: true.
-      cluster:    # @schema type:[boolean, null]
-      # -- Database used to store information. Default: "0".
-      database:   # @schema type:[string, null]
-      # -- Endpoints of the Redis instances to connect to. Default: "".
-      endpoints: ""
-      # -- The username to use when connecting to Redis endpoints. Default: "".
+  redis:
+    # -- Enable Redis Cluster. Default: true.
+    cluster:    # @schema type:[boolean, null]
+    # -- Database used to store information. Default: "0".
+    database:   # @schema type:[string, null]
+    # -- Endpoints of the Redis instances to connect to. Default: "".
+    endpoints: ""
+    # -- The username to use when connecting to Redis endpoints. Default: "".
+    username: ""
+    # -- The password to use when connecting to Redis endpoints. Default: "".
+    password: ""
+    sentinel:
+      # -- Name of the set of main nodes to use for main selection. Required when using Sentinel. Default: "".
+      masterset: ""
+      # -- Username to use for sentinel authentication (can be different from endpoint username). Default: "".
       username: ""
-      # -- The password to use when connecting to Redis endpoints. Default: "".
+      # -- Password to use for sentinel authentication (can be different from endpoint password). Default: "".
       password: ""
-      sentinel:
-        # -- Name of the set of main nodes to use for main selection. Required when using Sentinel. Default: "".
-        masterset: ""
-        # -- Username to use for sentinel authentication (can be different from endpoint username). Default: "".
-        username: ""
-        # -- Password to use for sentinel authentication (can be different from endpoint password). Default: "".
-        password: ""
-      # -- Timeout applied on connection with redis. Default: "0s".
-      timeout: ""
-      tls:
-        # -- Path to the certificate authority used for the secured connection.
-        ca: ""
-        # -- Path to the public certificate used for the secure connection.
-        cert: ""
-        # -- Path to the private key used for the secure connection.
-        key: ""
-        # -- When insecureSkipVerify is set to true, the TLS connection accepts any certificate presented by the server. Default: false.
-        insecureSkipVerify: false
+    # -- Timeout applied on connection with redis. Default: "0s".
+    timeout: ""
+    tls:
+      # -- Path to the certificate authority used for the secured connection.
+      ca: ""
+      # -- Path to the public certificate used for the secure connection.
+      cert: ""
+      # -- Path to the private key used for the secure connection.
+      key: ""
+      # -- When insecureSkipVerify is set to true, the TLS connection accepts any certificate presented by the server. Default: false.
+      insecureSkipVerify: false
   # Enable export of errors logs to the platform. Default: true.
   sendlogs:  # @schema type:[boolean, null]
```

## 31.1.1  ![AppVersion: v3.1.4](https://img.shields.io/static/v1?label=AppVersion&message=v3.1.4&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2024-09-20

* fix: 🐛 updateStrategy behavior
* feat(deps): update traefik docker tag to v3.1.4
* chore(release): 🚀 publish v31.1.1

## 31.1.0  ![AppVersion: v3.1.3](https://img.shields.io/static/v1?label=AppVersion&message=v3.1.3&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2024-09-18

* fix: 🐛 update CRD to v3.1
* feat: ✨ input validation using schema
* feat: ✨ add AllowACMEByPass and improve schema/doc on ports values
* feat: add new webhooks and removes unnecessary ones
* feat(deps): update traefik docker tag to v3.1.3
* chore(release): 🚀 publish v31.1.0

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 2232d9e..1b9d0fd 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -2,13 +2,13 @@
 # This is a YAML-formatted file.
 # Declare variables to be passed into templates
 
-image:
+image:  # @schema additionalProperties: false
   # -- Traefik image host registry
   registry: docker.io
   # -- Traefik image repository
   repository: traefik
   # -- defaults to appVersion
-  tag:
+  tag:  # @schema type:[string, null]
   # -- Traefik image pull policy
   pullPolicy: IfNotPresent
 
@@ -23,27 +23,27 @@ deployment:
   # -- Number of pods of the deployment (only applies when kind == Deployment)
   replicas: 1
   # -- Number of old history to retain to allow rollback (If not set, default Kubernetes value is set to 10)
-  # revisionHistoryLimit: 1
+  revisionHistoryLimit:  # @schema type:[integer, null];minimum:0
   # -- Amount of time (in seconds) before Kubernetes will send the SIGKILL signal if Traefik does not shut down
   terminationGracePeriodSeconds: 60
   # -- The minimum number of seconds Traefik needs to be up and running before the DaemonSet/Deployment controller considers it available
   minReadySeconds: 0
-  ## Override the liveness/readiness port. This is useful to integrate traefik
+  ## -- Override the liveness/readiness port. This is useful to integrate traefik
   ## with an external Load Balancer that performs healthchecks.
   ## Default: ports.traefik.port
-  # healthchecksPort: 9000
-  ## Override the liveness/readiness host. Useful for getting ping to respond on non-default entryPoint.
+  healthchecksPort:  # @schema type:[integer, null];minimum:0
+  ## -- Override the liveness/readiness host. Useful for getting ping to respond on non-default entryPoint.
   ## Default: ports.traefik.hostIP if set, otherwise Pod IP
-  # healthchecksHost: localhost
-  ## Override the liveness/readiness scheme. Useful for getting ping to
+  healthchecksHost: ""
+  ## -- Override the liveness/readiness scheme. Useful for getting ping to
   ## respond on websecure entryPoint.
-  # healthchecksScheme: HTTPS
-  ## Override the readiness path.
+  healthchecksScheme:   # @schema enum:[HTTP, HTTPS, null]; type:[string, null]; default: HTTP
+  ## -- Override the readiness path.
   ## Default: /ping
-  # readinessPath: /ping
-  # Override the liveness path.
+  readinessPath: ""
+  # -- Override the liveness path.
   # Default: /ping
-  # livenessPath: /ping
+  livenessPath: ""
   # -- Additional deployment annotations (e.g. for jaeger-operator sidecar injection)
   annotations: {}
   # -- Additional deployment labels (e.g. for filtering deployment by custom labels)
@@ -80,7 +80,7 @@ deployment:
   # -- Use process namespace sharing
   shareProcessNamespace: false
   # -- Custom pod DNS policy. Apply if `hostNetwork: true`
-  # dnsPolicy: ClusterFirstWithHostNet
+  dnsPolicy: ""
   # -- Custom pod [DNS config](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#poddnsconfig-v1-core)
   dnsConfig: {}
   # -- Custom [host aliases](https://kubernetes.io/docs/tasks/network/customize-hosts-file-for-pods/)
@@ -99,24 +99,24 @@ deployment:
   #     host: localhost
   #     scheme: HTTP
   # -- Set a runtimeClassName on pod
-  runtimeClassName:
+  runtimeClassName: ""
 
 # -- [Pod Disruption Budget](https://kubernetes.io/docs/reference/kubernetes-api/policy-resources/pod-disruption-budget-v1/)
-podDisruptionBudget:
-  enabled:
-  maxUnavailable:
-  minAvailable:
+podDisruptionBudget:  # @schema additionalProperties: false
+  enabled: false
+  maxUnavailable:  # @schema type:[string, integer, null];minimum:0
+  minAvailable:    # @schema type:[string, integer, null];minimum:0
 
 # -- Create a default IngressClass for Traefik
-ingressClass:
+ingressClass:  # @schema additionalProperties: false
   enabled: true
   isDefaultClass: true
-  # name: my-custom-class
+  name: ""
 
-core:
+core:  # @schema additionalProperties: false
   # -- Can be used to use globally v2 router syntax
   # See https://doc.traefik.io/traefik/v3.0/migration/v2-to-v3/#new-v3-syntax-notable-changes
-  defaultRuleSyntax:
+  defaultRuleSyntax: ""
 
 # Traefik experimental features
 experimental:
@@ -133,11 +133,11 @@ gateway:
   # -- When providers.kubernetesGateway.enabled, deploy a default gateway
   enabled: true
   # -- Set a custom name to gateway
-  name:
+  name: ""
   # -- By default, Gateway is created in the same `Namespace` than Traefik.
-  namespace:
+  namespace: ""
   # -- Additional gateway annotations (e.g. for cert-manager.io/issuer)
-  annotations:
+  annotations: {}
   # -- Define listeners
   listeners:
     web:
@@ -145,11 +145,11 @@ gateway:
       # The port must match a port declared in ports section.
       port: 8000
       # -- Optional hostname. See [Hostname](https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.Hostname)
-      hostname:
+      hostname: ""
       # Specify expected protocol on this listener. See [ProtocolType](https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.ProtocolType)
       protocol: HTTP
       # -- Routes are restricted to namespace of the gateway [by default](https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.FromNamespaces
-      namespacePolicy:
+      namespacePolicy:  # @schema type:[string, null]
     # websecure listener is disabled by default because certificateRefs needs to be added,
     # or you may specify TLS protocol with Passthrough mode and add "--providers.kubernetesGateway.experimentalChannel=true" in additionalArguments section.
     # websecure:
@@ -167,13 +167,13 @@ gateway:
     #   # -- TLS behavior for the TLS session initiated by the client. See [TLSModeType](https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.TLSModeType).
     #   mode:
 
-gatewayClass:
+gatewayClass:  # @schema additionalProperties: false
   # -- When providers.kubernetesGateway.enabled and gateway.enabled, deploy a default gatewayClass
   enabled: true
   # -- Set a custom name to GatewayClass
-  name:
+  name: ""
   # -- Additional gatewayClass labels (e.g. for filtering gateway objects by custom labels)
-  labels:
+  labels: {}
 
 ingressRoute:
   dashboard:
@@ -218,14 +218,14 @@ ingressRoute:
     # -- TLS options (e.g. secret containing certificate)
     tls: {}
 
-updateStrategy:
+updateStrategy:  # @schema additionalProperties: false
   # -- Customize updateStrategy: RollingUpdate or OnDelete
   type: RollingUpdate
   rollingUpdate:
-    maxUnavailable: 0
-    maxSurge: 1
+    maxUnavailable: 0  # @schema type:[integer, string, null]
+    maxSurge: 1        # @schema type:[integer, string, null]
 
-readinessProbe:
+readinessProbe:  # @schema additionalProperties: false
   # -- The number of consecutive failures allowed before considering the probe as failed.
   failureThreshold: 1
   # -- The number of seconds to wait before starting the first probe.
@@ -236,7 +236,7 @@ readinessProbe:
   successThreshold: 1
   # -- The number of seconds to wait for a probe response before considering it as failed.
   timeoutSeconds: 2
-livenessProbe:
+livenessProbe:  # @schema additionalProperties: false
   # -- The number of consecutive failures allowed before considering the probe as failed.
   failureThreshold: 3
   # -- The number of seconds to wait before starting the first probe.
@@ -249,9 +249,9 @@ livenessProbe:
   timeoutSeconds: 2
 
 # -- Define [Startup Probe](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/#define-startup-probes)
-startupProbe:
+startupProbe: {}
 
-providers:
+providers:  # @schema additionalProperties: false
   kubernetesCRD:
     # -- Load Kubernetes IngressRoute provider
     enabled: true
@@ -262,12 +262,12 @@ providers:
     # -- Allows to return 503 when there is no endpoints available
     allowEmptyServices: true
     # -- When the parameter is set, only resources containing an annotation with the same value are processed. Otherwise, resources missing the annotation, having an empty value, or the value traefik are processed. It will also set required annotation on Dashboard and Healthcheck IngressRoute when enabled.
-    ingressClass:
+    ingressClass: ""
     # labelSelector: environment=production,method=traefik
     # -- Array of namespaces to watch. If left empty, Traefik watches all namespaces.
     namespaces: []
     # -- Defines whether to use Native Kubernetes load-balancing mode by default.
-    nativeLBByDefault:
+    nativeLBByDefault: false
 
   kubernetesIngress:
     # -- Load Kubernetes Ingress provider
@@ -277,7 +277,7 @@ providers:
     # -- Allows to return 503 when there is no endpoints available
     allowEmptyServices: true
     # -- When ingressClass is set, only Ingresses containing an annotation with the same value are processed. Otherwise, Ingresses missing the annotation, having an empty value, or the value traefik are processed.
-    ingressClass:
+    ingressClass:  # @schema type:[string, null]
     # labelSelector: environment=production,method=traefik
     # -- Array of namespaces to watch. If left empty, Traefik watches all namespaces.
     namespaces: []
@@ -288,7 +288,7 @@ providers:
       # By default this Traefik service
       # pathOverride: ""
     # -- Defines whether to use Native Kubernetes load-balancing mode by default.
-    nativeLBByDefault:
+    nativeLBByDefault: false
 
   kubernetesGateway:
     # -- Enable Traefik Gateway provider for Gateway API
@@ -299,7 +299,7 @@ providers:
     # -- Array of namespaces to watch. If left empty, Traefik watches all namespaces.
     namespaces: []
     # -- A label selector can be defined to filter on specific GatewayClass objects only.
-    labelselector:
+    labelselector: ""
 
   file:
     # -- Create a file provider
@@ -307,7 +307,7 @@ providers:
     # -- Allows Traefik to automatically watch for file changes
     watch: true
     # -- File content (YAML format, go template supported) (see https://doc.traefik.io/traefik/providers/file/)
-    content:
+    content: ""
 
 # -- Add volumes to the traefik pod. The volume name will be passed to tpl.
 # This can be used to mount a cert pair or a configmap that holds a config.toml file.
@@ -333,90 +333,88 @@ additionalVolumeMounts: []
 logs:
   general:
     # -- Set [logs format](https://doc.traefik.io/traefik/observability/logs/#format)
-    # @default common
-    format:
+    format:  # @schema enum:["common", "json", null]; type:[string, null]; default: "common"
     # By default, the level is set to INFO.
     # -- Alternative logging levels are DEBUG, PANIC, FATAL, ERROR, WARN, and INFO.
-    level: INFO
-    #
-    # filePath: "/var/log/traefik/traefik.log
-    # noColor: true
+    level: "INFO"  # @schema enum:[INFO,WARN,ERROR,FATAL,PANIC,DEBUG]; default: "INFO"
+    # -- To write the logs into a log file, use the filePath option.
+    filePath: ""
+    # -- When set to true and format is common, it disables the colorized output.
+    noColor: false
   access:
     # -- To enable access logs
     enabled: false
     # -- Set [access log format](https://doc.traefik.io/traefik/observability/access-logs/#format)
-    format:
+    format:  # @schema enum:["CLF", "json", null]; type:[string, null]; default: "CLF"
     # filePath: "/var/log/traefik/access.log
     # -- Set [bufferingSize](https://doc.traefik.io/traefik/observability/access-logs/#bufferingsize)
-    bufferingSize:
+    bufferingSize:  # @schema type:[integer, null]
     # -- Set [filtering](https://docs.traefik.io/observability/access-logs/#filtering)
     filters: {}
-    # statuscodes: "200,300-302"
-    # retryattempts: true
-    # minduration: 10ms
+    statuscodes: ""
+    retryattempts: false
+    minduration: ""
     # -- Enables accessLogs for internal resources. Default: false.
-    addInternals:
+    addInternals: false
     fields:
       general:
-        # -- Available modes: keep, drop, redact.
-        defaultmode: keep
+        # -- Set default mode for fields.names
+        defaultmode: keep  # @schema enum:[keep, drop, redact]; default: keep
         # -- Names of the fields to limit.
         names: {}
-        ## Examples:
-        # ClientUsername: drop
       # -- [Limit logged fields or headers](https://doc.traefik.io/traefik/observability/access-logs/#limiting-the-fieldsincluding-headers)
       headers:
-        # -- Available modes: keep, drop, redact.
-        defaultmode: drop
+        # -- Set default mode for fields.headers
+        defaultmode: drop  # @schema enum:[keep, drop, redact]; default: drop
         names: {}
 
 metrics:
   ## -- Enable metrics for internal resources. Default: false
-  addInternals:
+  addInternals: false
 
   ## -- Prometheus is enabled by default.
   ## -- It can be disabled by setting "prometheus: null"
   prometheus:
     # -- Entry point used to expose metrics.
     entryPoint: metrics
-    ## Enable metrics on entry points. Default=true
-    # addEntryPointsLabels: false
-    ## Enable metrics on routers. Default=false
-    # addRoutersLabels: true
-    ## Enable metrics on services. Default=true
-    # addServicesLabels: false
+    ## Enable metrics on entry points. Default: true
+    addEntryPointsLabels:  # @schema type:[boolean, null]
+    ## Enable metrics on routers. Default: false
+    addRoutersLabels:  # @schema type:[boolean, null]
+    ## Enable metrics on services. Default: true
+    addServicesLabels:  # @schema type:[boolean, null]
     ## Buckets for latency metrics. Default="0.1,0.3,1.2,5.0"
-    # buckets: "0.5,1.0,2.5"
+    buckets: ""
     ## When manualRouting is true, it disables the default internal router in
     ## order to allow creating a custom router for prometheus@internal service.
-    # manualRouting: true
+    manualRouting: false
     service:
       # -- Create a dedicated metrics service to use with ServiceMonitor
-      enabled:
-      labels:
-      annotations:
+      enabled: false
+      labels: {}
+      annotations: {}
     # -- When set to true, it won't check if Prometheus Operator CRDs are deployed
-    disableAPICheck:
+    disableAPICheck:  # @schema type:[boolean, null]
     serviceMonitor:
       # -- Enable optional CR for Prometheus Operator. See EXAMPLES.md for more details.
       enabled: false
-      metricRelabelings:
-      relabelings:
-      jobLabel:
-      interval:
-      honorLabels:
-      scrapeTimeout:
-      honorTimestamps:
-      enableHttp2:
-      followRedirects:
-      additionalLabels:
-      namespace:
-      namespaceSelector:
+      metricRelabelings: []
+      relabelings: []
+      jobLabel: ""
+      interval: ""
+      honorLabels: false
+      scrapeTimeout: ""
+      honorTimestamps: false
+      enableHttp2: false
+      followRedirects: false
+      additionalLabels: {}
+      namespace: ""
+      namespaceSelector: {}
     prometheusRule:
       # -- Enable optional CR for Prometheus Operator. See EXAMPLES.md for more details.
       enabled: false
-      additionalLabels:
-      namespace:
+      additionalLabels: {}
+      namespace: ""
 
   #  datadog:
   #    ## Address instructs exporter to send metrics to datadog-agent at this address.
@@ -469,55 +467,55 @@ metrics:
     # -- Set to true in order to enable the OpenTelemetry metrics
     enabled: false
     # -- Enable metrics on entry points. Default: true
-    addEntryPointsLabels:
+    addEntryPointsLabels:  # @schema type:[boolean, null]
     # -- Enable metrics on routers. Default: false
-    addRoutersLabels:
+    addRoutersLabels:  # @schema type:[boolean, null]
     # -- Enable metrics on services. Default: true
-    addServicesLabels:
+    addServicesLabels:  # @schema type:[boolean, null]
     # -- Explicit boundaries for Histogram data points. Default: [.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10]
-    explicitBoundaries:
+    explicitBoundaries: []
     # -- Interval at which metrics are sent to the OpenTelemetry Collector. Default: 10s
-    pushInterval:
+    pushInterval: ""
     http:
       # -- Set to true in order to send metrics to the OpenTelemetry Collector using HTTP.
       enabled: false
       # -- Format: <scheme>://<host>:<port><path>. Default: http://localhost:4318/v1/metrics
-      endpoint:
+      endpoint: ""
       # -- Additional headers sent with metrics by the reporter to the OpenTelemetry Collector.
-      headers:
+      headers: {}
       ## Defines the TLS configuration used by the reporter to send metrics to the OpenTelemetry Collector.
       tls:
         # -- The path to the certificate authority, it defaults to the system bundle.
-        ca:
+        ca: ""
         # -- The path to the public certificate. When using this option, setting the key option is required.
-        cert:
+        cert: ""
         # -- The path to the private key. When using this option, setting the cert option is required.
-        key:
+        key: ""
         # -- When set to true, the TLS connection accepts any certificate presented by the server regardless of the hostnames it covers.
-        insecureSkipVerify:
+        insecureSkipVerify:  # @schema type:[boolean, null]
     grpc:
       # -- Set to true in order to send metrics to the OpenTelemetry Collector using gRPC
       enabled: false
       # -- Format: <scheme>://<host>:<port><path>. Default: http://localhost:4318/v1/metrics
-      endpoint:
+      endpoint: ""
       # -- Allows reporter to send metrics to the OpenTelemetry Collector without using a secured protocol.
-      insecure:
+      insecure: false
       ## Defines the TLS configuration used by the reporter to send metrics to the OpenTelemetry Collector.
       tls:
         # -- The path to the certificate authority, it defaults to the system bundle.
-        ca:
+        ca: ""
         # -- The path to the public certificate. When using this option, setting the key option is required.
-        cert:
+        cert: ""
         # -- The path to the private key. When using this option, setting the cert option is required.
-        key:
+        key: ""
         # -- When set to true, the TLS connection accepts any certificate presented by the server regardless of the hostnames it covers.
-        insecureSkipVerify:
+        insecureSkipVerify: false
 
 ## Tracing
 # -- https://doc.traefik.io/traefik/observability/tracing/overview/
-tracing:
+tracing:  # @schema additionalProperties: false
   # -- Enables tracing for internal resources. Default: false.
-  addInternals:
+  addInternals: false
   otlp:
     # -- See https://doc.traefik.io/traefik/v3.0/observability/tracing/opentelemetry/
     enabled: false
@@ -525,36 +523,36 @@ tracing:
       # -- Set to true in order to send metrics to the OpenTelemetry Collector using HTTP.
       enabled: false
       # -- Format: <scheme>://<host>:<port><path>. Default: http://localhost:4318/v1/metrics
-      endpoint:
+      endpoint: ""
       # -- Additional headers sent with metrics by the reporter to the OpenTelemetry Collector.
-      headers:
+      headers: {}
       ## Defines the TLS configuration used by the reporter to send metrics to the OpenTelemetry Collector.
       tls:
         # -- The path to the certificate authority, it defaults to the system bundle.
-        ca:
+        ca: ""
         # -- The path to the public certificate. When using this option, setting the key option is required.
-        cert:
+        cert: ""
         # -- The path to the private key. When using this option, setting the cert option is required.
-        key:
+        key: ""
         # -- When set to true, the TLS connection accepts any certificate presented by the server regardless of the hostnames it covers.
-        insecureSkipVerify:
+        insecureSkipVerify: false
     grpc:
       # -- Set to true in order to send metrics to the OpenTelemetry Collector using gRPC
       enabled: false
       # -- Format: <scheme>://<host>:<port><path>. Default: http://localhost:4318/v1/metrics
-      endpoint:
+      endpoint: ""
       # -- Allows reporter to send metrics to the OpenTelemetry Collector without using a secured protocol.
-      insecure:
+      insecure: false
       ## Defines the TLS configuration used by the reporter to send metrics to the OpenTelemetry Collector.
       tls:
         # -- The path to the certificate authority, it defaults to the system bundle.
-        ca:
+        ca: ""
         # -- The path to the public certificate. When using this option, setting the key option is required.
-        cert:
+        cert: ""
         # -- The path to the private key. When using this option, setting the cert option is required.
-        key:
+        key: ""
         # -- When set to true, the TLS connection accepts any certificate presented by the server regardless of the hostnames it covers.
-        insecureSkipVerify:
+        insecureSkipVerify: false
 
 # -- Global command arguments to be passed to all traefik's pods
 globalArguments:
@@ -587,13 +585,12 @@ ports:
   traefik:
     port: 9000
     # -- Use hostPort if set.
-    # hostPort: 9000
-    #
+    hostPort:  # @schema type:[integer, null]; minimum:0
     # -- Use hostIP if set. If not set, Kubernetes will default to 0.0.0.0, which
     # means it's listening on all your interfaces and all your IPs. You may want
     # to set this value if you need traefik to listen on specific interface
     # only.
-    # hostIP: 192.168.100.10
+    hostIP:  # @schema type:[string, null]
 
     # Defines whether the port is exposed if service.type is LoadBalancer or
     # NodePort.
@@ -617,112 +614,93 @@ ports:
       default: true
     exposedPort: 80
     ## -- Different target traefik port on the cluster, useful for IP type LB
-    # targetPort: 80
+    targetPort:  # @schema type:[integer, null]; minimum:0
     # The port protocol (TCP/UDP)
     protocol: TCP
-    # -- Use nodeport if set. This is useful if you have configured Traefik in a
-    # LoadBalancer.
-    # nodePort: 32080
+    # -- See [upstream documentation](https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport)
+    nodePort:  # @schema type:[integer, null]; minimum:0
     # Port Redirections
     # Added in 2.2, you can make permanent redirects via entrypoints.
     # https://docs.traefik.io/routing/entrypoints/#redirection
-    # redirectTo:
-    #   port: websecure
-    #   (Optional)
-    #   priority: 10
-    #   permanent: true
-    #
-    # -- Trust forwarded headers information (X-Forwarded-*).
-    # forwardedHeaders:
-    #   trustedIPs: []
-    #   insecure: false
-    #
-    # -- Enable the Proxy Protocol header parsing for the entry point
-    # proxyProtocol:
-    #   trustedIPs: []
-    #   insecure: false
-    #
+    redirectTo: {}
+    forwardedHeaders:
+      # -- Trust forwarded headers information (X-Forwarded-*).
+      trustedIPs: []
+      insecure: false
+    proxyProtocol:
+      # -- Enable the Proxy Protocol header parsing for the entry point
+      trustedIPs: []
+      insecure: false
     # -- Set transport settings for the entrypoint; see also
     # https://doc.traefik.io/traefik/routing/entrypoints/#transport
     transport:
       respondingTimeouts:
-        readTimeout:
-        writeTimeout:
-        idleTimeout:
+        readTimeout:   # @schema type:[string, integer, null]
+        writeTimeout:  # @schema type:[string, integer, null]
+        idleTimeout:   # @schema type:[string, integer, null]
       lifeCycle:
-        requestAcceptGraceTimeout:
-        graceTimeOut:
-      keepAliveMaxRequests:
-      keepAliveMaxTime:
+        requestAcceptGraceTimeout:  # @schema type:[string, integer, null]
+        graceTimeOut:               # @schema type:[string, integer, null]
+      keepAliveMaxRequests:         # @schema type:[integer, null]; minimum:0
+      keepAliveMaxTime:             # @schema type:[string, integer, null]
   websecure:
     ## -- Enable this entrypoint as a default entrypoint. When a service doesn't explicitly set an entrypoint it will only use this entrypoint.
     # asDefault: true
     port: 8443
-    # hostPort: 8443
-    # containerPort: 8443
+    hostPort:  # @schema type:[integer, null]; minimum:0
+    containerPort:  # @schema type:[integer, null]; minimum:0
     expose:
       default: true
     exposedPort: 443
     ## -- Different target traefik port on the cluster, useful for IP type LB
-    # targetPort: 80
+    targetPort:  # @schema type:[integer, null]; minimum:0
     ## -- The port protocol (TCP/UDP)
     protocol: TCP
-    # nodePort: 32443
-    ## -- Specify an application protocol. This may be used as a hint for a Layer 7 load balancer.
-    # appProtocol: https
-    #
-    ## -- Enable HTTP/3 on the entrypoint
-    ## Enabling it will also enable http3 experimental feature
-    ## https://doc.traefik.io/traefik/routing/entrypoints/#http3
-    ## There are known limitations when trying to listen on same ports for
-    ## TCP & UDP (Http3). There is a workaround in this chart using dual Service.
-    ## https://github.com/kubernetes/kubernetes/issues/47249#issuecomment-587960741
+    # -- See [upstream documentation](https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport)
+    nodePort:  # @schema type:[integer, null]; minimum:0
+    # -- See [upstream documentation](https://kubernetes.io/docs/concepts/services-networking/service/#application-protocol)
+    appProtocol:  # @schema type:[string, null]
+    # -- See [upstream documentation](https://doc.traefik.io/traefik/routing/entrypoints/#allowacmebypass)
+    allowACMEByPass: false
     http3:
+      ## -- Enable HTTP/3 on the entrypoint
+      ## Enabling it will also enable http3 experimental feature
+      ## https://doc.traefik.io/traefik/routing/entrypoints/#http3
+      ## There are known limitations when trying to listen on same ports for
+      ## TCP & UDP (Http3). There is a workaround in this chart using dual Service.
+      ## https://github.com/kubernetes/kubernetes/issues/47249#issuecomment-587960741
       enabled: false
-    # advertisedPort: 4443
-    #
-    # -- Trust forwarded headers information (X-Forwarded-*).
-    # forwardedHeaders:
-    #   trustedIPs: []
-    #   insecure: false
-    #
-    # -- Enable the Proxy Protocol header parsing for the entry point
-    # proxyProtocol:
-    #   trustedIPs: []
-    #   insecure: false
-    #
-    # -- Set transport settings for the entrypoint; see also
-    # https://doc.traefik.io/traefik/routing/entrypoints/#transport
+      advertisedPort:  # @schema type:[integer, null]; minimum:0
+    forwardedHeaders:
+        # -- Trust forwarded headers information (X-Forwarded-*).
+      trustedIPs: []
+      insecure: false
+    proxyProtocol:
+      # -- Enable the Proxy Protocol header parsing for the entry point
+      trustedIPs: []
+      insecure: false
+    # -- See [upstream documentation](https://doc.traefik.io/traefik/routing/entrypoints/#transport)
     transport:
       respondingTimeouts:
-        readTimeout:
-        writeTimeout:
-        idleTimeout:
+        readTimeout:   # @schema type:[string, integer, null]
+        writeTimeout:  # @schema type:[string, integer, null]
+        idleTimeout:   # @schema type:[string, integer, null]
       lifeCycle:
-        requestAcceptGraceTimeout:
-        graceTimeOut:
-      keepAliveMaxRequests:
-      keepAliveMaxTime:
-    #
-    ## Set TLS at the entrypoint
-    ## https://doc.traefik.io/traefik/routing/entrypoints/#tls
+        requestAcceptGraceTimeout:  # @schema type:[string, integer, null]
+        graceTimeOut:               # @schema type:[string, integer, null]
+      keepAliveMaxRequests:         # @schema type:[integer, null]; minimum:0
+      keepAliveMaxTime:             # @schema type:[string, integer, null]
+    # --  See [upstream documentation](https://doc.traefik.io/traefik/routing/entrypoints/#tls)
     tls:
       enabled: true
-      # this is the name of a TLSOption definition
       options: ""
       certResolver: ""
       domains: []
-      # - main: example.com
-      #   sans:
-      #     - foo.example.com
-      #     - bar.example.com
-    #
     # -- One can apply Middlewares on an entrypoint
     # https://doc.traefik.io/traefik/middlewares/overview/
     # https://doc.traefik.io/traefik/routing/entrypoints/#middlewares
     # -- /!\ It introduces here a link between your static configuration and your dynamic configuration /!\
     # It follows the provider naming convention: https://doc.traefik.io/traefik/providers/overview/#provider-namespace
-    # middlewares:
     #   - namespace-name1@kubernetescrd
     #   - namespace-name2@kubernetescrd
     middlewares: []
@@ -730,10 +708,6 @@ ports:
     # -- When using hostNetwork, use another port to avoid conflict with node exporter:
     # https://github.com/prometheus/prometheus/wiki/Default-port-allocations
     port: 9100
-    # hostPort: 9100
-    # Defines whether the port is exposed if service.type is LoadBalancer or
-    # NodePort.
-    #
     # -- You may not want to expose the metrics port on production deployments.
     # If you want to access it from outside your cluster,
     # use `kubectl port-forward` or create a secure ingress
@@ -810,15 +784,15 @@ persistence:
   # It can be used to store TLS certificates, see `storage` in certResolvers
   enabled: false
   name: data
-  #  existingClaim: ""
+  existingClaim: ""
   accessMode: ReadWriteOnce
   size: 128Mi
-  # storageClass: ""
-  # volumeName: ""
+  storageClass: ""
+  volumeName: ""
   path: /data
   annotations: {}
   # -- Only mount a subpath of the Volume into the pod
-  # subPath: ""
+  subPath: ""
 
 # -- Certificates resolvers configuration.
 # Ref: https://doc.traefik.io/traefik/https/acme/#certificate-resolvers
@@ -832,7 +806,7 @@ certResolvers: {}
 hostNetwork: false
 
 # -- Whether Role Based Access Control objects like roles and rolebindings should be created
-rbac:
+rbac:  # @schema additionalProperties: false
   enabled: true
   # When set to true:
   # 1. Use `Role` and `RoleBinding` instead of `ClusterRole` and `ClusterRoleBinding`.
@@ -843,7 +817,7 @@ rbac:
   namespaced: false
   # Enable user-facing roles
   # https://kubernetes.io/docs/reference/access-authn-authz/rbac/#user-facing-roles
-  # aggregateTo: [ "admin" ]
+  aggregateTo: []
   # List of Kubernetes secrets that are accessible for Traefik. If empty, then access is granted to every secret.
   secretResourceNames: []
 
@@ -852,7 +826,7 @@ podSecurityPolicy:
   enabled: false
 
 # -- The service account the pods will use to interact with the Kubernetes API
-serviceAccount:
+serviceAccount:  # @schema additionalProperties: false
   # If set, an existing service account is used
   # If not set, a service account is created automatically using the fullname template
   name: ""
@@ -918,54 +892,54 @@ extraObjects: []
 
 # -- This field override the default Release Namespace for Helm.
 # It will not affect optional CRDs such as `ServiceMonitor` and `PrometheusRules`
-namespaceOverride:
+namespaceOverride: ""
 
 ## -- This field override the default app.kubernetes.io/instance label for all Objects.
-instanceLabelOverride:
+instanceLabelOverride: ""
 
 # Traefik Hub configuration. See https://doc.traefik.io/traefik-hub/
 hub:
   # -- Name of `Secret` with key 'token' set to a valid license token.
   # It enables API Gateway.
-  token:
+  token: ""
   apimanagement:
     # -- Set to true in order to enable API Management. Requires a valid license token.
-    enabled:
+    enabled: false
     admission:
       # -- WebHook admission server listen address. Default: "0.0.0.0:9943".
-      listenAddr:
+      listenAddr: ""
       # -- Certificate of the WebHook admission server. Default: "hub-agent-cert".
-      secretName:
+      secretName: ""
 
   ratelimit:
     redis:
       # -- Enable Redis Cluster. Default: true.
-      cluster:
+      cluster:    # @schema type:[boolean, null]
       # -- Database used to store information. Default: "0".
-      database:
+      database:   # @schema type:[string, null]
       # -- Endpoints of the Redis instances to connect to. Default: "".
-      endpoints:
+      endpoints: ""
       # -- The username to use when connecting to Redis endpoints. Default: "".
-      username:
+      username: ""
       # -- The password to use when connecting to Redis endpoints. Default: "".
-      password:
+      password: ""
       sentinel:
         # -- Name of the set of main nodes to use for main selection. Required when using Sentinel. Default: "".
-        masterset:
+        masterset: ""
         # -- Username to use for sentinel authentication (can be different from endpoint username). Default: "".
-        username:
+        username: ""
         # -- Password to use for sentinel authentication (can be different from endpoint password). Default: "".
-        password:
+        password: ""
       # -- Timeout applied on connection with redis. Default: "0s".
-      timeout:
+      timeout: ""
       tls:
         # -- Path to the certificate authority used for the secured connection.
-        ca:
+        ca: ""
         # -- Path to the public certificate used for the secure connection.
-        cert:
+        cert: ""
         # -- Path to the private key used for the secure connection.
-        key:
+        key: ""
         # -- When insecureSkipVerify is set to true, the TLS connection accepts any certificate presented by the server. Default: false.
-        insecureSkipVerify:
+        insecureSkipVerify: false
   # Enable export of errors logs to the platform. Default: true.
-  sendlogs:
+  sendlogs:  # @schema type:[boolean, null]
```

## 31.0.0  ![AppVersion: v3.1.2](https://img.shields.io/static/v1?label=AppVersion&message=v3.1.2&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2024-09-03

* fix(Traefik Hub): update CRDs to v1.5.0
* fix(HTTP3): split udp and tcp Service when service.single is false
* fix!: 🐛 set allowEmptyServices to true by default
* feat(Traefik Hub): update CRDs to v1.7.0
* chore(release): 🚀 publish v31.0.0

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 78eeacf..2232d9e 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -260,7 +260,7 @@ providers:
     # -- Allows to reference ExternalName services in IngressRoute
     allowExternalNameServices: false
     # -- Allows to return 503 when there is no endpoints available
-    allowEmptyServices: false
+    allowEmptyServices: true
     # -- When the parameter is set, only resources containing an annotation with the same value are processed. Otherwise, resources missing the annotation, having an empty value, or the value traefik are processed. It will also set required annotation on Dashboard and Healthcheck IngressRoute when enabled.
     ingressClass:
     # labelSelector: environment=production,method=traefik
@@ -275,7 +275,7 @@ providers:
     # -- Allows to reference ExternalName services in Ingress
     allowExternalNameServices: false
     # -- Allows to return 503 when there is no endpoints available
-    allowEmptyServices: false
+    allowEmptyServices: true
     # -- When ingressClass is set, only Ingresses containing an annotation with the same value are processed. Otherwise, Ingresses missing the annotation, having an empty value, or the value traefik are processed.
     ingressClass:
     # labelSelector: environment=production,method=traefik
```

## 30.1.0  ![AppVersion: v3.1.2](https://img.shields.io/static/v1?label=AppVersion&message=v3.1.2&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2024-08-14

* fix: disable default HTTPS listener for gateway
* fix(Gateway API): wildcard support in hostname
* fix(Gateway API): use Standard channel by default
* feat: ✨ rework namespaced RBAC with `disableClusterScopeResources`
* chore(release): 🚀 publish v30.1.0
* chore(deps): update traefik docker tag to v3.1.2
* chore(deps): update traefik docker tag to v3.1.1

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 83b6d98..78eeacf 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -150,20 +150,22 @@ gateway:
       protocol: HTTP
       # -- Routes are restricted to namespace of the gateway [by default](https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.FromNamespaces
       namespacePolicy:
-    websecure:
-      # -- Port is the network port. Multiple listeners may use the same port, subject to the Listener compatibility rules.
-      # The port must match a port declared in ports section.
-      port: 8443
-      # -- Optional hostname. See [Hostname](https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.Hostname)
-      hostname:
-      # Specify expected protocol on this listener See [ProtocolType](https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.ProtocolType)
-      protocol: HTTPS
-      # -- Routes are restricted to namespace of the gateway [by default](https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.FromNamespaces)
-      namespacePolicy:
-      # -- Add certificates for TLS or HTTPS protocols. See [GatewayTLSConfig](https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io%2fv1.GatewayTLSConfig)
-      certificateRefs:
-      # -- TLS behavior for the TLS session initiated by the client. See [TLSModeType](https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.TLSModeType).
-      mode:
+    # websecure listener is disabled by default because certificateRefs needs to be added,
+    # or you may specify TLS protocol with Passthrough mode and add "--providers.kubernetesGateway.experimentalChannel=true" in additionalArguments section.
+    # websecure:
+    #   # -- Port is the network port. Multiple listeners may use the same port, subject to the Listener compatibility rules.
+    #   # The port must match a port declared in ports section.
+    #   port: 8443
+    #   # -- Optional hostname. See [Hostname](https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.Hostname)
+    #   hostname:
+    #   # Specify expected protocol on this listener See [ProtocolType](https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.ProtocolType)
+    #   protocol: HTTPS
+    #   # -- Routes are restricted to namespace of the gateway [by default](https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.FromNamespaces)
+    #   namespacePolicy:
+    #   # -- Add certificates for TLS or HTTPS protocols. See [GatewayTLSConfig](https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io%2fv1.GatewayTLSConfig)
+    #   certificateRefs:
+    #   # -- TLS behavior for the TLS session initiated by the client. See [TLSModeType](https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.TLSModeType).
+    #   mode:
 
 gatewayClass:
   # -- When providers.kubernetesGateway.enabled and gateway.enabled, deploy a default gatewayClass
@@ -279,10 +281,6 @@ providers:
     # labelSelector: environment=production,method=traefik
     # -- Array of namespaces to watch. If left empty, Traefik watches all namespaces.
     namespaces: []
-    # - "default"
-    # Disable cluster IngressClass Lookup - Requires Traefik V3.
-    # When combined with rbac.namespaced: true, ClusterRole will not be created and ingresses must use kubernetes.io/ingress.class annotation instead of spec.ingressClassName.
-    disableIngressClassLookup: false
     # IP used for Kubernetes Ingress endpoints
     publishedService:
       enabled: false
@@ -836,9 +834,12 @@ hostNetwork: false
 # -- Whether Role Based Access Control objects like roles and rolebindings should be created
 rbac:
   enabled: true
-  # If set to false, installs ClusterRole and ClusterRoleBinding so Traefik can be used across namespaces.
-  # If set to true, installs Role and RoleBinding instead of ClusterRole/ClusterRoleBinding. Providers will only watch target namespace.
-  # When combined with providers.kubernetesIngress.disableIngressClassLookup: true and Traefik V3, ClusterRole to watch IngressClass is also disabled.
+  # When set to true:
+  # 1. Use `Role` and `RoleBinding` instead of `ClusterRole` and `ClusterRoleBinding`.
+  # 2. Set `disableIngressClassLookup` on Kubernetes Ingress providers with Traefik Proxy v3 until v3.1.1
+  # 3. Set `disableClusterScopeResources` on Kubernetes Ingress and CRD providers with Traefik Proxy v3.1.2+
+  # **NOTE**: `IngressClass`, `NodePortLB` and **Gateway** provider cannot be used with namespaced RBAC.
+  # See [upstream documentation](https://doc.traefik.io/traefik/providers/kubernetes-ingress/#disableclusterscoperesources) for more details.
   namespaced: false
   # Enable user-facing roles
   # https://kubernetes.io/docs/reference/access-authn-authz/rbac/#user-facing-roles
```

## 30.0.2  ![AppVersion: v3.1.0](https://img.shields.io/static/v1?label=AppVersion&message=v3.1.0&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2024-07-30

* fix(Traefik Hub): missing RBACs for Traefik Hub
* chore(release): 🚀 publish v30.0.2

## 30.0.1  ![AppVersion: v3.1.0](https://img.shields.io/static/v1?label=AppVersion&message=v3.1.0&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2024-07-29

* fix(Traefik Hub): support new RBACs for upcoming traefik hub release
* fix(Traefik Hub): RBACs missing with API Gateway
* feat: :release: v30.0.1

## 30.0.0  ![AppVersion: v3.1.0](https://img.shields.io/static/v1?label=AppVersion&message=v3.1.0&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2024-07-24

* fix: 🐛 ingressroute default name
* fix: namespaced RBACs hub api gateway
* fix: can't set gateway name
* fix(Gateway API): provide expected roles when using namespaced RBAC
* fix(Gateway API)!: revamp Gateway implementation
* feat: ✨ display release name and image full path in installation notes
* feat: use single ingressRoute template
* feat: handle log filePath and noColor
* chore(release): 🚀 publish v30.0.0
* chore(deps): update traefik docker tag to v3.1.0

**Upgrade Notes**

There is a breaking upgrade on how to configure Gateway with _values_.
This release supports Traefik Proxy v3.0 **and** v3.1.

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index c8bfd5b..83b6d98 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -134,14 +134,36 @@ gateway:
   enabled: true
   # -- Set a custom name to gateway
   name:
-  # -- Routes are restricted to namespace of the gateway [by default](https://gateway-api.sigs.k8s.io/references/spec/#gateway.networking.k8s.io/v1beta1.FromNamespaces)
-  namespacePolicy:
-  # -- See [GatewayTLSConfig](https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io%2fv1.GatewayTLSConfig)
-  certificateRefs:
   # -- By default, Gateway is created in the same `Namespace` than Traefik.
   namespace:
   # -- Additional gateway annotations (e.g. for cert-manager.io/issuer)
   annotations:
+  # -- Define listeners
+  listeners:
+    web:
+      # -- Port is the network port. Multiple listeners may use the same port, subject to the Listener compatibility rules.
+      # The port must match a port declared in ports section.
+      port: 8000
+      # -- Optional hostname. See [Hostname](https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.Hostname)
+      hostname:
+      # Specify expected protocol on this listener. See [ProtocolType](https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.ProtocolType)
+      protocol: HTTP
+      # -- Routes are restricted to namespace of the gateway [by default](https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.FromNamespaces
+      namespacePolicy:
+    websecure:
+      # -- Port is the network port. Multiple listeners may use the same port, subject to the Listener compatibility rules.
+      # The port must match a port declared in ports section.
+      port: 8443
+      # -- Optional hostname. See [Hostname](https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.Hostname)
+      hostname:
+      # Specify expected protocol on this listener See [ProtocolType](https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.ProtocolType)
+      protocol: HTTPS
+      # -- Routes are restricted to namespace of the gateway [by default](https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.FromNamespaces)
+      namespacePolicy:
+      # -- Add certificates for TLS or HTTPS protocols. See [GatewayTLSConfig](https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io%2fv1.GatewayTLSConfig)
+      certificateRefs:
+      # -- TLS behavior for the TLS session initiated by the client. See [TLSModeType](https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.TLSModeType).
+      mode:
 
 gatewayClass:
   # -- When providers.kubernetesGateway.enabled and gateway.enabled, deploy a default gatewayClass
@@ -161,6 +183,10 @@ ingressRoute:
     labels: {}
     # -- The router match rule used for the dashboard ingressRoute
     matchRule: PathPrefix(`/dashboard`) || PathPrefix(`/api`)
+    # -- The internal service used for the dashboard ingressRoute
+    services:
+      - name: api@internal
+        kind: TraefikService
     # -- Specify the allowed entrypoints to use for the dashboard ingress route, (e.g. traefik, web, websecure).
     # By default, it's using traefik entrypoint, which is not exposed.
     # /!\ Do not expose your dashboard without any protection over the internet /!\
@@ -178,6 +204,10 @@ ingressRoute:
     labels: {}
     # -- The router match rule used for the healthcheck ingressRoute
     matchRule: PathPrefix(`/ping`)
+    # -- The internal service used for the healthcheck ingressRoute
+    services:
+      - name: ping@internal
+        kind: TraefikService
     # -- Specify the allowed entrypoints to use for the healthcheck ingress route, (e.g. traefik, web, websecure).
     # By default, it's using traefik entrypoint, which is not exposed.
     entryPoints: ["traefik"]
@@ -307,9 +337,12 @@ logs:
     # -- Set [logs format](https://doc.traefik.io/traefik/observability/logs/#format)
     # @default common
     format:
-    # By default, the level is set to ERROR.
+    # By default, the level is set to INFO.
     # -- Alternative logging levels are DEBUG, PANIC, FATAL, ERROR, WARN, and INFO.
     level: INFO
+    #
+    # filePath: "/var/log/traefik/traefik.log
+    # noColor: true
   access:
     # -- To enable access logs
     enabled: false
```


## 29.0.1  ![AppVersion: v3.0.4](https://img.shields.io/static/v1?label=AppVersion&message=v3.0.4&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2024-07-09

* fix: semverCompare failing on some legitimate tags
* fix: RBACs for hub and disabled namespaced RBACs
* chore(release): 🚀 publish v29.0.1
* chore(deps): update jnorwood/helm-docs docker tag to v1.14.0

## 29.0.0  ![AppVersion: v3.0.4](https://img.shields.io/static/v1?label=AppVersion&message=v3.0.4&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Upgrade Notes**

This is a major breaking upgrade. [Migration guide](https://doc.traefik.io/traefik/v3.1/migration/v3/#v30-to-v31) from v3.0 to v3.1rc has been applied on this chart.

This release supports both Traefik Proxy v3.0.x and v3.1rc.

It comes with those breaking changes:

- Far better support on Gateway API v1.1: Gateway, GatewayClass, CRDs & RBAC (#1107)
- Many changes on CRDs & RBAC (#1072 & #1108)
- Refactor on Prometheus Operator support. Values has changed (#1114)
- Dashboard `IngressRoute` is now disabled by default (#1111)

CRDs needs to be upgraded: `kubectl apply --server-side --force-conflicts -k https://github.com/traefik/traefik-helm-chart/traefik/crds/`

**Release date:** 2024-07-05

* fix: 🐛 improve error message on additional service without ports
* fix:  allow multiples values in the `secretResourceNames` slice
* fix(rbac)!: nodes API permissions for Traefik v3.1+
* fix(dashboard): Only set ingressClass annotation when kubernetesCRD provider is listening for it
* fix!: prometheus operator settings
* feat: ✨ update CRDs & RBAC for Traefik Proxy
* feat: ✨ migrate to endpointslices rbac
* feat: allow to set hostAliases for traefik pod
* feat(providers): add nativeLBByDefault support
* feat(providers)!: improve kubernetesGateway and Gateway API support
* feat(dashboard)!: dashboard `IngressRoute` should be disabled by default
* docs: fix typos and broken link
* chore: update CRDs to v1.5.0
* chore: update CRDs to v1.4.0
* chore(release): publish v29.0.0
* chore(deps): update traefik docker tag to v3.0.4
* chore(deps): update traefik docker tag to v3.0.3

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index e440dcf..c8bfd5b 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -8,7 +8,7 @@ image:
   # -- Traefik image repository
   repository: traefik
   # -- defaults to appVersion
-  tag: ""
+  tag:
   # -- Traefik image pull policy
   pullPolicy: IfNotPresent

@@ -81,19 +81,12 @@ deployment:
   shareProcessNamespace: false
   # -- Custom pod DNS policy. Apply if `hostNetwork: true`
   # dnsPolicy: ClusterFirstWithHostNet
+  # -- Custom pod [DNS config](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#poddnsconfig-v1-core)
   dnsConfig: {}
-  # nameservers:
-  #   - 192.0.2.1 # this is an example
-  # searches:
-  #   - ns1.svc.cluster-domain.example
-  #   - my.dns.search.suffix
-  # options:
-  #   - name: ndots
-  #     value: "2"
-  #   - name: edns0
-  # -- Additional imagePullSecrets
+  # -- Custom [host aliases](https://kubernetes.io/docs/tasks/network/customize-hosts-file-for-pods/)
+  hostAliases: []
+  # -- Pull secret for fetching traefik container image
   imagePullSecrets: []
-  # - name: myRegistryKeySecretName
   # -- Pod lifecycle actions
   lifecycle: {}
   # preStop:
@@ -135,24 +128,33 @@ experimental:
   kubernetesGateway:
     # -- Enable traefik experimental GatewayClass CRD
     enabled: false
-    ## Routes are restricted to namespace of the gateway by default.
-    ## https://gateway-api.sigs.k8s.io/references/spec/#gateway.networking.k8s.io/v1beta1.FromNamespaces
-    # namespacePolicy: All
-    # certificate:
-    #   group: "core"
-    #   kind: "Secret"
-    #   name: "mysecret"
-    # -- By default, Gateway would be created to the Namespace you are deploying Traefik to.
-    # You may create that Gateway in another namespace, setting its name below:
-    # namespace: default
-    # Additional gateway annotations (e.g. for cert-manager.io/issuer)
-    # annotations:
-    #   cert-manager.io/issuer: letsencrypt
+
+gateway:
+  # -- When providers.kubernetesGateway.enabled, deploy a default gateway
+  enabled: true
+  # -- Set a custom name to gateway
+  name:
+  # -- Routes are restricted to namespace of the gateway [by default](https://gateway-api.sigs.k8s.io/references/spec/#gateway.networking.k8s.io/v1beta1.FromNamespaces)
+  namespacePolicy:
+  # -- See [GatewayTLSConfig](https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io%2fv1.GatewayTLSConfig)
+  certificateRefs:
+  # -- By default, Gateway is created in the same `Namespace` than Traefik.
+  namespace:
+  # -- Additional gateway annotations (e.g. for cert-manager.io/issuer)
+  annotations:
+
+gatewayClass:
+  # -- When providers.kubernetesGateway.enabled and gateway.enabled, deploy a default gatewayClass
+  enabled: true
+  # -- Set a custom name to GatewayClass
+  name:
+  # -- Additional gatewayClass labels (e.g. for filtering gateway objects by custom labels)
+  labels:

 ingressRoute:
   dashboard:
     # -- Create an IngressRoute for the dashboard
-    enabled: true
+    enabled: false
     # -- Additional ingressRoute annotations (e.g. for kubernetes.io/ingress.class)
     annotations: {}
     # -- Additional ingressRoute labels (e.g. for filtering IngressRoute by custom labels)
@@ -227,11 +229,13 @@ providers:
     allowExternalNameServices: false
     # -- Allows to return 503 when there is no endpoints available
     allowEmptyServices: false
-    # ingressClass: traefik-internal
+    # -- When the parameter is set, only resources containing an annotation with the same value are processed. Otherwise, resources missing the annotation, having an empty value, or the value traefik are processed. It will also set required annotation on Dashboard and Healthcheck IngressRoute when enabled.
+    ingressClass:
     # labelSelector: environment=production,method=traefik
     # -- Array of namespaces to watch. If left empty, Traefik watches all namespaces.
     namespaces: []
-    # - "default"
+    # -- Defines whether to use Native Kubernetes load-balancing mode by default.
+    nativeLBByDefault:

   kubernetesIngress:
     # -- Load Kubernetes Ingress provider
@@ -240,7 +244,8 @@ providers:
     allowExternalNameServices: false
     # -- Allows to return 503 when there is no endpoints available
     allowEmptyServices: false
-    # ingressClass: traefik-internal
+    # -- When ingressClass is set, only Ingresses containing an annotation with the same value are processed. Otherwise, Ingresses missing the annotation, having an empty value, or the value traefik are processed.
+    ingressClass:
     # labelSelector: environment=production,method=traefik
     # -- Array of namespaces to watch. If left empty, Traefik watches all namespaces.
     namespaces: []
@@ -254,6 +259,19 @@ providers:
       # Published Kubernetes Service to copy status from. Format: namespace/servicename
       # By default this Traefik service
       # pathOverride: ""
+    # -- Defines whether to use Native Kubernetes load-balancing mode by default.
+    nativeLBByDefault:
+
+  kubernetesGateway:
+    # -- Enable Traefik Gateway provider for Gateway API
+    enabled: false
+    # -- Toggles support for the Experimental Channel resources (Gateway API release channels documentation).
+    # This option currently enables support for TCPRoute and TLSRoute.
+    experimentalChannel: false
+    # -- Array of namespaces to watch. If left empty, Traefik watches all namespaces.
+    namespaces: []
+    # -- A label selector can be defined to filter on specific GatewayClass objects only.
+    labelselector:

   file:
     # -- Create a file provider
@@ -341,6 +359,34 @@ metrics:
     ## When manualRouting is true, it disables the default internal router in
     ## order to allow creating a custom router for prometheus@internal service.
     # manualRouting: true
+    service:
+      # -- Create a dedicated metrics service to use with ServiceMonitor
+      enabled:
+      labels:
+      annotations:
+    # -- When set to true, it won't check if Prometheus Operator CRDs are deployed
+    disableAPICheck:
+    serviceMonitor:
+      # -- Enable optional CR for Prometheus Operator. See EXAMPLES.md for more details.
+      enabled: false
+      metricRelabelings:
+      relabelings:
+      jobLabel:
+      interval:
+      honorLabels:
+      scrapeTimeout:
+      honorTimestamps:
+      enableHttp2:
+      followRedirects:
+      additionalLabels:
+      namespace:
+      namespaceSelector:
+    prometheusRule:
+      # -- Enable optional CR for Prometheus Operator. See EXAMPLES.md for more details.
+      enabled: false
+      additionalLabels:
+      namespace:
+
   #  datadog:
   #    ## Address instructs exporter to send metrics to datadog-agent at this address.
   #    address: "127.0.0.1:8125"
@@ -436,55 +482,6 @@ metrics:
         # -- When set to true, the TLS connection accepts any certificate presented by the server regardless of the hostnames it covers.
         insecureSkipVerify:

-  ## -- enable optional CRDs for Prometheus Operator
-  ##
-  ## Create a dedicated metrics service for use with ServiceMonitor
-  #  service:
-  #    enabled: false
-  #    labels: {}
-  #    annotations: {}
-  ## When set to true, it won't check if Prometheus Operator CRDs are deployed
-  #  disableAPICheck: false
-  #  serviceMonitor:
-  #    metricRelabelings: []
-  #      - sourceLabels: [__name__]
-  #        separator: ;
-  #        regex: ^fluentd_output_status_buffer_(oldest|newest)_.+
-  #        replacement: $1
-  #        action: drop
-  #    relabelings: []
-  #      - sourceLabels: [__meta_kubernetes_pod_node_name]
-  #        separator: ;
-  #        regex: ^(.*)$
-  #        targetLabel: nodename
-  #        replacement: $1
-  #        action: replace
-  #    jobLabel: traefik
-  #    interval: 30s
-  #    honorLabels: true
-  #    # (Optional)
-  #    # scrapeTimeout: 5s
-  #    # honorTimestamps: true
-  #    # enableHttp2: true
-  #    # followRedirects: true
-  #    # additionalLabels:
-  #    #   foo: bar
-  #    # namespace: "another-namespace"
-  #    # namespaceSelector: {}
-  #  prometheusRule:
-  #    additionalLabels: {}
-  #    namespace: "another-namespace"
-  #    rules:
-  #      - alert: TraefikDown
-  #        expr: up{job="traefik"} == 0
-  #        for: 5m
-  #        labels:
-  #          context: traefik
-  #          severity: warning
-  #        annotations:
-  #          summary: "Traefik Down"
-  #          description: "{{ $labels.pod }} on {{ $labels.nodename }} is down"
-
 ## Tracing
 # -- https://doc.traefik.io/traefik/observability/tracing/overview/
 tracing:
```

## 28.3.0  ![AppVersion: v3.0.2](https://img.shields.io/static/v1?label=AppVersion&message=v3.0.2&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2024-06-14

* fix: 🐛 namespaced rbac when kubernetesIngress provider is disabled
* fix: 🐛  add divisor: '1' to GOMAXPROCS and GOMEMLIMIT
* fix(security): 🐛 🔒️ mount service account token on pod level
* fix(Traefik Hub): remove obsolete CRD
* fix(Traefik Hub): remove namespace in mutating webhook
* feat: allow setting permanent on redirectTo
* chore(release): publish v28.3.0
* chore(deps): update traefik docker tag to v3.0.2

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index c558c78..e440dcf 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -602,6 +602,7 @@ ports:
     #   port: websecure
     #   (Optional)
     #   priority: 10
+    #   permanent: true
     #
     # -- Trust forwarded headers information (X-Forwarded-*).
     # forwardedHeaders:
```

## 28.2.0  ![AppVersion: v3.0.1](https://img.shields.io/static/v1?label=AppVersion&message=v3.0.1&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2024-05-28

* fix(IngressClass): provides annotation on IngressRoutes when it's enabled
* feat: ✨ simplify values and provide more examples
* feat: add deletecollection right on secrets
* chore(release): 🚀 publish v28.2.0
* chore(deps): update traefik docker tag to v3.0.1

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 2fd9282..c558c78 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -1,4 +1,7 @@
 # Default values for Traefik
+# This is a YAML-formatted file.
+# Declare variables to be passed into templates
+
 image:
   # -- Traefik image host registry
   registry: docker.io
@@ -12,9 +15,6 @@ image:
 # -- Add additional label to all resources
 commonLabels: {}

-#
-# Configure the deployment
-#
 deployment:
   # -- Enable deployment
   enabled: true
@@ -74,10 +74,6 @@ deployment:
   # - name: volume-permissions
   #   image: busybox:latest
   #   command: ["sh", "-c", "touch /data/acme.json; chmod -v 600 /data/acme.json"]
-  #   securityContext:
-  #     runAsNonRoot: true
-  #     runAsGroup: 65532
-  #     runAsUser: 65532
   #   volumeMounts:
   #     - name: data
   #       mountPath: /data
@@ -112,13 +108,11 @@ deployment:
   # -- Set a runtimeClassName on pod
   runtimeClassName:

-# -- Pod disruption budget
+# -- [Pod Disruption Budget](https://kubernetes.io/docs/reference/kubernetes-api/policy-resources/pod-disruption-budget-v1/)
 podDisruptionBudget:
-  enabled: false
-  # maxUnavailable: 1
-  # maxUnavailable: 33%
-  # minAvailable: 0
-  # minAvailable: 25%
+  enabled:
+  maxUnavailable:
+  minAvailable:

 # -- Create a default IngressClass for Traefik
 ingressClass:
@@ -155,7 +149,6 @@ experimental:
     # annotations:
     #   cert-manager.io/issuer: letsencrypt

-## Create an IngressRoute for the dashboard
 ingressRoute:
   dashboard:
     # -- Create an IngressRoute for the dashboard
@@ -221,15 +214,7 @@ livenessProbe:
   # -- The number of seconds to wait for a probe response before considering it as failed.
   timeoutSeconds: 2

-# -- Define Startup Probe for container: https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/#define-startup-probes
-# eg.
-# `startupProbe:
-#   exec:
-#     command:
-#       - mycommand
-#       - foo
-#   initialDelaySeconds: 5
-#   periodSeconds: 5`
+# -- Define [Startup Probe](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/#define-startup-probes)
 startupProbe:

 providers:
@@ -276,18 +261,8 @@ providers:
     # -- Allows Traefik to automatically watch for file changes
     watch: true
     # -- File content (YAML format, go template supported) (see https://doc.traefik.io/traefik/providers/file/)
-    content: ""
-      # http:
-      #   routers:
-      #     router0:
-      #       entryPoints:
-      #       - web
-      #       middlewares:
-      #       - my-basic-auth
-      #       service: service-foo
-      #       rule: Path(`/foo`)
+    content:

-#
 # -- Add volumes to the traefik pod. The volume name will be passed to tpl.
 # This can be used to mount a cert pair or a configmap that holds a config.toml file.
 # After the volume has been mounted, add the configs into traefik by using the `additionalArguments` list below, eg:
@@ -311,26 +286,21 @@ additionalVolumeMounts: []

 logs:
   general:
-    # -- By default, the logs use a text format (common), but you can
-    # also ask for the json format in the format option
-    # format: json
+    # -- Set [logs format](https://doc.traefik.io/traefik/observability/logs/#format)
+    # @default common
+    format:
     # By default, the level is set to ERROR.
     # -- Alternative logging levels are DEBUG, PANIC, FATAL, ERROR, WARN, and INFO.
     level: INFO
   access:
     # -- To enable access logs
     enabled: false
-    ## By default, logs are written using the Common Log Format (CLF) on stdout.
-    ## To write logs in JSON, use json in the format option.
-    ## If the given format is unsupported, the default (CLF) is used instead.
-    # format: json
+    # -- Set [access log format](https://doc.traefik.io/traefik/observability/access-logs/#format)
+    format:
     # filePath: "/var/log/traefik/access.log
-    ## To write the logs in an asynchronous fashion, specify a bufferingSize option.
-    ## This option represents the number of log lines Traefik will keep in memory before writing
-    ## them to the selected output. In some cases, this option can greatly help performances.
-    # bufferingSize: 100
-    ## Filtering
-    # -- https://docs.traefik.io/observability/access-logs/#filtering
+    # -- Set [bufferingSize](https://doc.traefik.io/traefik/observability/access-logs/#bufferingsize)
+    bufferingSize:
+    # -- Set [filtering](https://docs.traefik.io/observability/access-logs/#filtering)
     filters: {}
     # statuscodes: "200,300-302"
     # retryattempts: true
@@ -345,15 +315,11 @@ logs:
         names: {}
         ## Examples:
         # ClientUsername: drop
+      # -- [Limit logged fields or headers](https://doc.traefik.io/traefik/observability/access-logs/#limiting-the-fieldsincluding-headers)
       headers:
         # -- Available modes: keep, drop, redact.
         defaultmode: drop
-        # -- Names of the headers to limit.
         names: {}
-        ## Examples:
-        # User-Agent: redact
-        # Authorization: drop
-        # Content-Type: keep

 metrics:
   ## -- Enable metrics for internal resources. Default: false
@@ -567,16 +533,15 @@ globalArguments:
 - "--global.checknewversion"
 - "--global.sendanonymoususage"

-#
-# Configure Traefik static configuration
 # -- Additional arguments to be passed at Traefik's binary
-# All available options available on https://docs.traefik.io/reference/static-configuration/cli/
-## Use curly braces to pass values: `helm install --set="additionalArguments={--providers.kubernetesingress.ingressclass=traefik-internal,--log.level=DEBUG}"`
+# See [CLI Reference](https://docs.traefik.io/reference/static-configuration/cli/)
+# Use curly braces to pass values: `helm install --set="additionalArguments={--providers.kubernetesingress.ingressclass=traefik-internal,--log.level=DEBUG}"`
 additionalArguments: []
 #  - "--providers.kubernetesingress.ingressclass=traefik-internal"
 #  - "--log.level=DEBUG"

 # -- Environment variables to be passed to Traefik's binary
+# @default -- See _values.yaml_
 env:
 - name: POD_NAME
   valueFrom:
@@ -586,25 +551,9 @@ env:
   valueFrom:
     fieldRef:
       fieldPath: metadata.namespace
-# - name: SOME_VAR
-#   value: some-var-value
-# - name: SOME_VAR_FROM_CONFIG_MAP
-#   valueFrom:
-#     configMapRef:
-#       name: configmap-name
-#       key: config-key
-# - name: SOME_SECRET
-#   valueFrom:
-#     secretKeyRef:
-#       name: secret-name
-#       key: secret-key

 # -- Environment variables to be passed to Traefik's binary from configMaps or secrets
 envFrom: []
-# - configMapRef:
-#     name: config-map-name
-# - secretRef:
-#     name: secret-name

 ports:
   traefik:
@@ -766,28 +715,12 @@ ports:
     # -- The port protocol (TCP/UDP)
     protocol: TCP

-# -- TLS Options are created as TLSOption CRDs
-# https://doc.traefik.io/traefik/https/tls/#tls-options
+# -- TLS Options are created as [TLSOption CRDs](https://doc.traefik.io/traefik/https/tls/#tls-options)
 # When using `labelSelector`, you'll need to set labels on tlsOption accordingly.
-# Example:
-# tlsOptions:
-#   default:
-#     labels: {}
-#     sniStrict: true
-#   custom-options:
-#     labels: {}
-#     curvePreferences:
-#       - CurveP521
-#       - CurveP384
+# See EXAMPLE.md for details.
 tlsOptions: {}

-# -- TLS Store are created as TLSStore CRDs. This is useful if you want to set a default certificate
-# https://doc.traefik.io/traefik/https/tls/#default-certificate
-# Example:
-# tlsStore:
-#   default:
-#     defaultCertificate:
-#       secretName: tls-cert
+# -- TLS Store are created as [TLSStore CRDs](https://doc.traefik.io/traefik/https/tls/#default-certificate). This is useful if you want to set a default certificate. See EXAMPLE.md for details.
 tlsStore: {}

 service:
@@ -839,29 +772,8 @@ service:

 autoscaling:
   # -- Create HorizontalPodAutoscaler object.
+  # See EXAMPLES.md for more details.
   enabled: false
-#   minReplicas: 1
-#   maxReplicas: 10
-#   metrics:
-#   - type: Resource
-#     resource:
-#       name: cpu
-#       target:
-#         type: Utilization
-#         averageUtilization: 60
-#   - type: Resource
-#     resource:
-#       name: memory
-#       target:
-#         type: Utilization
-#         averageUtilization: 60
-#   behavior:
-#     scaleDown:
-#       stabilizationWindowSeconds: 300
-#       policies:
-#       - type: Pods
-#         value: 1
-#         periodSeconds: 60

 persistence:
   # -- Enable persistence using Persistent Volume Claims
@@ -879,27 +791,10 @@ persistence:
   # -- Only mount a subpath of the Volume into the pod
   # subPath: ""

-# -- Certificates resolvers configuration
+# -- Certificates resolvers configuration.
+# Ref: https://doc.traefik.io/traefik/https/acme/#certificate-resolvers
+# See EXAMPLES.md for more details.
 certResolvers: {}
-#   letsencrypt:
-#     # for challenge options cf. https://doc.traefik.io/traefik/https/acme/
-#     email: email@example.com
-#     dnsChallenge:
-#       # also add the provider's required configuration under env
-#       # or expand then from secrets/configmaps with envfrom
-#       # cf. https://doc.traefik.io/traefik/https/acme/#providers
-#       provider: digitalocean
-#       # add futher options for the dns challenge as needed
-#       # cf. https://doc.traefik.io/traefik/https/acme/#dnschallenge
-#       delayBeforeCheck: 30
-#       resolvers:
-#         - 1.1.1.1
-#         - 8.8.8.8
-#     tlsChallenge: true
-#     httpChallenge:
-#       entryPoint: "web"
-#     # It has to match the path with a persistent volume
-#     storage: /data/acme.json

 # -- If hostNetwork is true, runs traefik in the host network namespace
 # To prevent unschedulabel pods due to port collisions, if hostNetwork=true
@@ -933,14 +828,8 @@ serviceAccount:
 # -- Additional serviceAccount annotations (e.g. for oidc authentication)
 serviceAccountAnnotations: {}

-# -- The resources parameter defines CPU and memory requirements and limits for Traefik's containers.
+# -- [Resources](https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/) for `traefik` container.
 resources: {}
-# requests:
-#   cpu: "100m"
-#   memory: "50Mi"
-# limits:
-#   cpu: "300m"
-#   memory: "150Mi"

 # -- This example pod anti-affinity forces the scheduler to put traefik pods
 # -- on nodes where no other traefik pods are scheduled.
@@ -970,30 +859,22 @@ topologySpreadConstraints: []
 #    topologyKey: kubernetes.io/hostname
 #    whenUnsatisfiable: DoNotSchedule

-# -- Pods can have priority.
-# -- Priority indicates the importance of a Pod relative to other Pods.
+# -- [Pod Priority and Preemption](https://kubernetes.io/docs/concepts/scheduling-eviction/pod-priority-preemption/)
 priorityClassName: ""

-# -- Set the container security context
-# -- To run the container with ports below 1024 this will need to be adjusted to run as root
+# -- [SecurityContext](https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/pod-v1/#security-context-1)
+# @default -- See _values.yaml_
 securityContext:
+  allowPrivilegeEscalation: false
   capabilities:
     drop: [ALL]
   readOnlyRootFilesystem: true
-  allowPrivilegeEscalation: false

+# -- [Pod Security Context](https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/pod-v1/#security-context)
+# @default -- See _values.yaml_
 podSecurityContext:
-  # /!\ When setting fsGroup, Kubernetes will recursively change ownership and
-  # permissions for the contents of each volume to match the fsGroup. This can
-  # be an issue when storing sensitive content like TLS Certificates /!\
-  # fsGroup: 65532
-  # -- Specifies the policy for changing ownership and permissions of volume contents to match the fsGroup.
-  fsGroupChangePolicy: "OnRootMismatch"
-  # -- The ID of the group for all containers in the pod to run as.
   runAsGroup: 65532
-  # -- Specifies whether the containers should run as a non-root user.
   runAsNonRoot: true
-  # -- The ID of the user for all containers in the pod to run as.
   runAsUser: 65532

 #
@@ -1003,16 +884,16 @@ podSecurityContext:
 # See #595 for more details and traefik/tests/values/extra.yaml for example.
 extraObjects: []

-# This will override the default Release Namespace for Helm.
+# -- This field override the default Release Namespace for Helm.
 # It will not affect optional CRDs such as `ServiceMonitor` and `PrometheusRules`
-# namespaceOverride: traefik
-#
-## -- This will override the default app.kubernetes.io/instance label for all Objects.
-# instanceLabelOverride: traefik
+namespaceOverride:
+
+## -- This field override the default app.kubernetes.io/instance label for all Objects.
+instanceLabelOverride:

-# -- Traefik Hub configuration. See https://doc.traefik.io/traefik-hub/
+# Traefik Hub configuration. See https://doc.traefik.io/traefik-hub/
 hub:
-  # Name of Secret with key 'token' set to a valid license token.
+  # -- Name of `Secret` with key 'token' set to a valid license token.
   # It enables API Gateway.
   token:
   apimanagement:
```

## 28.1.0 ![AppVersion: v3.0.0](https://img.shields.io/static/v1?label=AppVersion&message=v3.0.0&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

* fix(Traefik Hub): do not deploy mutating webhook when enabling only API Gateway
* feat(Traefik Hub): use Traefik Proxy otlp config
* chore: 🔧 update Traefik Hub CRD to v1.3.3

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 70297f6..2fd9282 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -1010,3 +1010,49 @@
 ## -- This will override the default app.kubernetes.io/instance label for all Objects.
 # instanceLabelOverride: traefik

+# -- Traefik Hub configuration. See https://doc.traefik.io/traefik-hub/
+hub:
+  # Name of Secret with key 'token' set to a valid license token.
+  # It enables API Gateway.
+  token:
+  apimanagement:
+    # -- Set to true in order to enable API Management. Requires a valid license token.
+    enabled:
+    admission:
+      # -- WebHook admission server listen address. Default: "0.0.0.0:9943".
+      listenAddr:
+      # -- Certificate of the WebHook admission server. Default: "hub-agent-cert".
+      secretName:
+
+  ratelimit:
+    redis:
+      # -- Enable Redis Cluster. Default: true.
+      cluster:
+      # -- Database used to store information. Default: "0".
+      database:
+      # -- Endpoints of the Redis instances to connect to. Default: "".
+      endpoints:
+      # -- The username to use when connecting to Redis endpoints. Default: "".
+      username:
+      # -- The password to use when connecting to Redis endpoints. Default: "".
+      password:
+      sentinel:
+        # -- Name of the set of main nodes to use for main selection. Required when using Sentinel. Default: "".
+        masterset:
+        # -- Username to use for sentinel authentication (can be different from endpoint username). Default: "".
+        username:
+        # -- Password to use for sentinel authentication (can be different from endpoint password). Default: "".
+        password:
+      # -- Timeout applied on connection with redis. Default: "0s".
+      timeout:
+      tls:
+        # -- Path to the certificate authority used for the secured connection.
+        ca:
+        # -- Path to the public certificate used for the secure connection.
+        cert:
+        # -- Path to the private key used for the secure connection.
+        key:
+        # -- When insecureSkipVerify is set to true, the TLS connection accepts any certificate presented by the server. Default: false.
+        insecureSkipVerify:
+  # Enable export of errors logs to the platform. Default: true.
+  sendlogs:
```

## 28.1.0-beta.3  ![AppVersion: v3.0.0](https://img.shields.io/static/v1?label=AppVersion&message=v3.0.0&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2024-05-03

* chore: 🔧 update Traefik Hub CRD to v1.3.2
* chore(release): 🚀 publish v28.1.0-beta.3

## 28.1.0-beta.2  ![AppVersion: v3.0.0](https://img.shields.io/static/v1?label=AppVersion&message=v3.0.0&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2024-05-02

* fix: 🐛 refine Traefik Hub support
* chore(release): 🚀 publish v28.1.0-beta.2

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index ce0a7a3..70297f6 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -1015,13 +1015,15 @@ hub:
   # Name of Secret with key 'token' set to a valid license token.
   # It enables API Gateway.
   token:
-  admission:
-    # -- WebHook admission server listen address. Default: "0.0.0.0:9943".
-    listenAddr:
-    # -- Certificate of the WebHook admission server. Default: "hub-agent-cert".
-    secretName:
-  # -- Set to true in order to enable API Management. Requires a valid license token.
   apimanagement:
+    # -- Set to true in order to enable API Management. Requires a valid license token.
+    enabled:
+    admission:
+      # -- WebHook admission server listen address. Default: "0.0.0.0:9943".
+      listenAddr:
+      # -- Certificate of the WebHook admission server. Default: "hub-agent-cert".
+      secretName:
+
   metrics:
     opentelemetry:
       # -- Set to true to enable OpenTelemetry metrics exporter of Traefik Hub.
```

## 28.1.0-beta.1  ![AppVersion: v3.0.0](https://img.shields.io/static/v1?label=AppVersion&message=v3.0.0&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2024-04-30

* feat: :rocket: add initial support for Traefik Hub Api Gateway
* chore(release): 🚀 publish v28.1.0-beta.1

## 28.0.0  ![AppVersion: v3.0.0](https://img.shields.io/static/v1?label=AppVersion&message=v3.0.0&color=success&logo=) ![Kubernetes: >=1.22.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.22.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2024-04-30

* style: 🎨 consistent capitalization on `--entryPoints` CLI flag
* fix: 🐛 only expose http3 port on service when TCP variant is exposed
* fix: 🐛 logs filters on status codes
* feat: ✨ add support of `experimental-v3.0` unstable version
* feat: ability to override liveness and readiness probe paths
* feat(ports): add transport options
* chore(release): publish v28.0.0

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index c0d72d8..2bff10d 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -38,6 +38,12 @@ deployment:
   ## Override the liveness/readiness scheme. Useful for getting ping to
   ## respond on websecure entryPoint.
   # healthchecksScheme: HTTPS
+  ## Override the readiness path.
+  ## Default: /ping
+  # readinessPath: /ping
+  # Override the liveness path.
+  # Default: /ping
+  # livenessPath: /ping
   # -- Additional deployment annotations (e.g. for jaeger-operator sidecar injection)
   annotations: {}
   # -- Additional deployment labels (e.g. for filtering deployment by custom labels)
@@ -648,15 +654,28 @@ ports:
     #   (Optional)
     #   priority: 10
     #
-    # Trust forwarded  headers information (X-Forwarded-*).
+    # -- Trust forwarded headers information (X-Forwarded-*).
     # forwardedHeaders:
     #   trustedIPs: []
     #   insecure: false
     #
-    # Enable the Proxy Protocol header parsing for the entry point
+    # -- Enable the Proxy Protocol header parsing for the entry point
     # proxyProtocol:
     #   trustedIPs: []
     #   insecure: false
+    #
+    # -- Set transport settings for the entrypoint; see also
+    # https://doc.traefik.io/traefik/routing/entrypoints/#transport
+    transport:
+      respondingTimeouts:
+        readTimeout:
+        writeTimeout:
+        idleTimeout:
+      lifeCycle:
+        requestAcceptGraceTimeout:
+        graceTimeOut:
+      keepAliveMaxRequests:
+      keepAliveMaxTime:
   websecure:
     ## -- Enable this entrypoint as a default entrypoint. When a service doesn't explicitly set an entrypoint it will only use this entrypoint.
     # asDefault: true
@@ -684,16 +703,29 @@ ports:
       enabled: false
     # advertisedPort: 4443
     #
-    ## -- Trust forwarded  headers information (X-Forwarded-*).
+    # -- Trust forwarded headers information (X-Forwarded-*).
     # forwardedHeaders:
     #   trustedIPs: []
     #   insecure: false
     #
-    ## -- Enable the Proxy Protocol header parsing for the entry point
+    # -- Enable the Proxy Protocol header parsing for the entry point
     # proxyProtocol:
     #   trustedIPs: []
     #   insecure: false
     #
+    # -- Set transport settings for the entrypoint; see also
+    # https://doc.traefik.io/traefik/routing/entrypoints/#transport
+    transport:
+      respondingTimeouts:
+        readTimeout:
+        writeTimeout:
+        idleTimeout:
+      lifeCycle:
+        requestAcceptGraceTimeout:
+        graceTimeOut:
+      keepAliveMaxRequests:
+      keepAliveMaxTime:
+    #
     ## Set TLS at the entrypoint
     ## https://doc.traefik.io/traefik/routing/entrypoints/#tls
     tls:
```

## 28.0.0-rc1  ![AppVersion: v3.0.0-rc5](https://img.shields.io/static/v1?label=AppVersion&message=v3.0.0-rc5&color=success&logo=) ![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.16.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2024-04-17

**Upgrade Notes**

This is a major breaking upgrade. [Migration guide](https://doc.traefik.io/traefik/v3.0/migration/v2-to-v3/) have been applied on the chart.

It needs a Kubernetes v1.22 or higher.
All CRDs using _API Group_ `traefik.containo.us` are not supported anymore in Traefik Proxy v3

CRDs needs to be upgraded: `kubectl apply --server-side --force-conflicts -k https://github.com/traefik/traefik-helm-chart/traefik/crds/`

After upgrade, CRDs with _API Group_ `traefik.containo.us` can be removed:

```shell
kubectl delete crds \
  ingressroutes.traefik.containo.us \
  ingressroutetcps.traefik.containo.us \
  ingressrouteudps.traefik.containo.us \
  middlewares.traefik.containo.us \
  middlewaretcps.traefik.containo.us \
  serverstransports.traefik.containo.us \
  tlsoptions.traefik.containo.us \
  tlsstores.traefik.containo.us \
  traefikservices.traefik.containo.us
```

**Changes**

* feat(podtemplate): set GOMEMLIMIT, GOMAXPROCS when limits are defined
* feat: ✨ fail gracefully when required port number is not set
* feat!: :boom: initial support of Traefik Proxy v3
* docs: 📚️ improve EXAMPLES on acme resolver
* chore(release): 🚀 publish v28 rc1

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index cd9fb6e..c0d72d8 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -120,12 +120,13 @@ ingressClass:
   isDefaultClass: true
   # name: my-custom-class

+core:
+  # -- Can be used to use globally v2 router syntax
+  # See https://doc.traefik.io/traefik/v3.0/migration/v2-to-v3/#new-v3-syntax-notable-changes
+  defaultRuleSyntax:
+
 # Traefik experimental features
 experimental:
-  # This value is no longer used, set the image.tag to a semver higher than 3.0, e.g. "v3.0.0-beta3"
-  # v3:
-  # -- Enable traefik version 3
-
   # -- Enable traefik experimental plugins
   plugins: {}
   # demo:
@@ -309,7 +310,7 @@ logs:
     # format: json
     # By default, the level is set to ERROR.
     # -- Alternative logging levels are DEBUG, PANIC, FATAL, ERROR, WARN, and INFO.
-    level: ERROR
+    level: INFO
   access:
     # -- To enable access logs
     enabled: false
@@ -328,6 +329,8 @@ logs:
     # statuscodes: "200,300-302"
     # retryattempts: true
     # minduration: 10ms
+    # -- Enables accessLogs for internal resources. Default: false.
+    addInternals:
     fields:
       general:
         # -- Available modes: keep, drop, redact.
@@ -347,6 +350,9 @@ logs:
         # Content-Type: keep

 metrics:
+  ## -- Enable metrics for internal resources. Default: false
+  addInternals:
+
   ## -- Prometheus is enabled by default.
   ## -- It can be disabled by setting "prometheus: null"
   prometheus:
@@ -376,31 +382,6 @@ metrics:
   #    # addRoutersLabels: true
   #    ## Enable metrics on services. Default=true
   #    # addServicesLabels: false
-  #  influxdb:
-  #    ## Address instructs exporter to send metrics to influxdb at this address.
-  #    address: localhost:8089
-  #    ## InfluxDB's address protocol (udp or http). Default="udp"
-  #    protocol: udp
-  #    ## InfluxDB database used when protocol is http. Default=""
-  #    # database: ""
-  #    ## InfluxDB retention policy used when protocol is http. Default=""
-  #    # retentionPolicy: ""
-  #    ## InfluxDB username (only with http). Default=""
-  #    # username: ""
-  #    ## InfluxDB password (only with http). Default=""
-  #    # password: ""
-  #    ## The interval used by the exporter to push metrics to influxdb. Default=10s
-  #    # pushInterval: 30s
-  #    ## Additional labels (influxdb tags) on all metrics.
-  #    # additionalLabels:
-  #    #   env: production
-  #    #   foo: bar
-  #    ## Enable metrics on entry points. Default=true
-  #    # addEntryPointsLabels: false
-  #    ## Enable metrics on routers. Default=false
-  #    # addRoutersLabels: true
-  #    ## Enable metrics on services. Default=true
-  #    # addServicesLabels: false
   #  influxdb2:
   #    ## Address instructs exporter to send metrics to influxdb v2 at this address.
   #    address: localhost:8086
@@ -435,43 +416,53 @@ metrics:
   #    # addRoutersLabels: true
   #    ## Enable metrics on services. Default=true
   #    # addServicesLabels: false
-  #  openTelemetry:
-  #    ## Address of the OpenTelemetry Collector to send metrics to.
-  #    address: "localhost:4318"
-  #    ## Enable metrics on entry points.
-  #    addEntryPointsLabels: true
-  #    ## Enable metrics on routers.
-  #    addRoutersLabels: true
-  #    ## Enable metrics on services.
-  #    addServicesLabels: true
-  #    ## Explicit boundaries for Histogram data points.
-  #    explicitBoundaries:
-  #      - "0.1"
-  #      - "0.3"
-  #      - "1.2"
-  #      - "5.0"
-  #    ## Additional headers sent with metrics by the reporter to the OpenTelemetry Collector.
-  #    headers:
-  #      foo: bar
-  #      test: test
-  #    ## Allows reporter to send metrics to the OpenTelemetry Collector without using a secured protocol.
-  #    insecure: true
-  #    ## Interval at which metrics are sent to the OpenTelemetry Collector.
-  #    pushInterval: 10s
-  #    ## Allows to override the default URL path used for sending metrics. This option has no effect when using gRPC transport.
-  #    path: /foo/v1/traces
-  #    ## Defines the TLS configuration used by the reporter to send metrics to the OpenTelemetry Collector.
-  #    tls:
-  #      ## The path to the certificate authority, it defaults to the system bundle.
-  #      ca: path/to/ca.crt
-  #      ## The path to the public certificate. When using this option, setting the key option is required.
-  #      cert: path/to/foo.cert
-  #      ## The path to the private key. When using this option, setting the cert option is required.
-  #      key: path/to/key.key
-  #      ## If set to true, the TLS connection accepts any certificate presented by the server regardless of the hostnames it covers.
-  #      insecureSkipVerify: true
-  #    ## This instructs the reporter to send metrics to the OpenTelemetry Collector using gRPC.
-  #    grpc: true
+  otlp:
+    # -- Set to true in order to enable the OpenTelemetry metrics
+    enabled: false
+    # -- Enable metrics on entry points. Default: true
+    addEntryPointsLabels:
+    # -- Enable metrics on routers. Default: false
+    addRoutersLabels:
+    # -- Enable metrics on services. Default: true
+    addServicesLabels:
+    # -- Explicit boundaries for Histogram data points. Default: [.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10]
+    explicitBoundaries:
+    # -- Interval at which metrics are sent to the OpenTelemetry Collector. Default: 10s
+    pushInterval:
+    http:
+      # -- Set to true in order to send metrics to the OpenTelemetry Collector using HTTP.
+      enabled: false
+      # -- Format: <scheme>://<host>:<port><path>. Default: http://localhost:4318/v1/metrics
+      endpoint:
+      # -- Additional headers sent with metrics by the reporter to the OpenTelemetry Collector.
+      headers:
+      ## Defines the TLS configuration used by the reporter to send metrics to the OpenTelemetry Collector.
+      tls:
+        # -- The path to the certificate authority, it defaults to the system bundle.
+        ca:
+        # -- The path to the public certificate. When using this option, setting the key option is required.
+        cert:
+        # -- The path to the private key. When using this option, setting the cert option is required.
+        key:
+        # -- When set to true, the TLS connection accepts any certificate presented by the server regardless of the hostnames it covers.
+        insecureSkipVerify:
+    grpc:
+      # -- Set to true in order to send metrics to the OpenTelemetry Collector using gRPC
+      enabled: false
+      # -- Format: <scheme>://<host>:<port><path>. Default: http://localhost:4318/v1/metrics
+      endpoint:
+      # -- Allows reporter to send metrics to the OpenTelemetry Collector without using a secured protocol.
+      insecure:
+      ## Defines the TLS configuration used by the reporter to send metrics to the OpenTelemetry Collector.
+      tls:
+        # -- The path to the certificate authority, it defaults to the system bundle.
+        ca:
+        # -- The path to the public certificate. When using this option, setting the key option is required.
+        cert:
+        # -- The path to the private key. When using this option, setting the cert option is required.
+        key:
+        # -- When set to true, the TLS connection accepts any certificate presented by the server regardless of the hostnames it covers.
+        insecureSkipVerify:

   ## -- enable optional CRDs for Prometheus Operator
   ##
@@ -524,51 +515,46 @@ metrics:

 ## Tracing
 # -- https://doc.traefik.io/traefik/observability/tracing/overview/
-tracing: {}
-#  openTelemetry: # traefik v3+ only
-#    grpc: true
-#    insecure: true
-#    address: localhost:4317
-# instana:
-#   localAgentHost: 127.0.0.1
-#   localAgentPort: 42699
-#   logLevel: info
-#   enableAutoProfile: true
-# datadog:
-#   localAgentHostPort: 127.0.0.1:8126
-#   debug: false
-#   globalTag: ""
-#   prioritySampling: false
-# jaeger:
-#   samplingServerURL: http://localhost:5778/sampling
-#   samplingType: const
-#   samplingParam: 1.0
-#   localAgentHostPort: 127.0.0.1:6831
-#   gen128Bit: false
-#   propagation: jaeger
-#   traceContextHeaderName: uber-trace-id
-#   disableAttemptReconnecting: true
-#   collector:
-#      endpoint: ""
-#      user: ""
-#      password: ""
-# zipkin:
-#   httpEndpoint: http://localhost:9411/api/v2/spans
-#   sameSpan: false
-#   id128Bit: true
-#   sampleRate: 1.0
-# haystack:
-#   localAgentHost: 127.0.0.1
-#   localAgentPort: 35000
-#   globalTag: ""
-#   traceIDHeaderName: ""
-#   parentIDHeaderName: ""
-#   spanIDHeaderName: ""
-#   baggagePrefixHeaderName: ""
-# elastic:
-#   serverURL: http://localhost:8200
-#   secretToken: ""
-#   serviceEnvironment: ""
+tracing:
+  # -- Enables tracing for internal resources. Default: false.
+  addInternals:
+  otlp:
+    # -- See https://doc.traefik.io/traefik/v3.0/observability/tracing/opentelemetry/
+    enabled: false
+    http:
+      # -- Set to true in order to send metrics to the OpenTelemetry Collector using HTTP.
+      enabled: false
+      # -- Format: <scheme>://<host>:<port><path>. Default: http://localhost:4318/v1/metrics
+      endpoint:
+      # -- Additional headers sent with metrics by the reporter to the OpenTelemetry Collector.
+      headers:
+      ## Defines the TLS configuration used by the reporter to send metrics to the OpenTelemetry Collector.
+      tls:
+        # -- The path to the certificate authority, it defaults to the system bundle.
+        ca:
+        # -- The path to the public certificate. When using this option, setting the key option is required.
+        cert:
+        # -- The path to the private key. When using this option, setting the cert option is required.
+        key:
+        # -- When set to true, the TLS connection accepts any certificate presented by the server regardless of the hostnames it covers.
+        insecureSkipVerify:
+    grpc:
+      # -- Set to true in order to send metrics to the OpenTelemetry Collector using gRPC
+      enabled: false
+      # -- Format: <scheme>://<host>:<port><path>. Default: http://localhost:4318/v1/metrics
+      endpoint:
+      # -- Allows reporter to send metrics to the OpenTelemetry Collector without using a secured protocol.
+      insecure:
+      ## Defines the TLS configuration used by the reporter to send metrics to the OpenTelemetry Collector.
+      tls:
+        # -- The path to the certificate authority, it defaults to the system bundle.
+        ca:
+        # -- The path to the public certificate. When using this option, setting the key option is required.
+        cert:
+        # -- The path to the private key. When using this option, setting the cert option is required.
+        key:
+        # -- When set to true, the TLS connection accepts any certificate presented by the server regardless of the hostnames it covers.
+        insecureSkipVerify:

 # -- Global command arguments to be passed to all traefik's pods
 globalArguments:
@@ -756,7 +742,6 @@ ports:
 #   default:
 #     labels: {}
 #     sniStrict: true
-#     preferServerCipherSuites: true
 #   custom-options:
 #     labels: {}
 #     curvePreferences:
```

## 27.0.0  ![AppVersion: v2.11.0](https://img.shields.io/static/v1?label=AppVersion&message=v2.11.0&color=success&logo=) ![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.16.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2024-04-02

**Upgrade notes**

Custom services and port exposure have been redesigned, requiring the following changes:
- if you were overriding port exposure behavior using the `expose` or `exposeInternal` flags, you should replace them with a service name to boolean mapping, i.e. replace this:

```yaml
ports:
   web:
      expose: false
      exposeInternal: true
```

with this:

```yaml
ports:
   web:
      expose:
         default: false
         internal: true
```

- if you were previously using the `service.internal` value, you should migrate the values to the `service.additionalServices.internal` value instead; this should yield the same results, but make sure to carefully check for any changes!

**Changes**

* fix: remove null annotations on dashboard `IngressRoute`
* fix(rbac): do not create clusterrole for namespace deployment on Traefik v3
* feat: restrict access to secrets
* feat!: :boom: refactor custom services and port exposure
* chore(release): 🚀 publish v27.0.0

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index dbd078f..363871d 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -250,6 +250,9 @@ providers:
     # -- Array of namespaces to watch. If left empty, Traefik watches all namespaces.
     namespaces: []
     # - "default"
+    # Disable cluster IngressClass Lookup - Requires Traefik V3.
+    # When combined with rbac.namespaced: true, ClusterRole will not be created and ingresses must use kubernetes.io/ingress.class annotation instead of spec.ingressClassName.
+    disableIngressClassLookup: false
     # IP used for Kubernetes Ingress endpoints
     publishedService:
       enabled: false
@@ -626,22 +629,20 @@ ports:
     # -- You SHOULD NOT expose the traefik port on production deployments.
     # If you want to access it from outside your cluster,
     # use `kubectl port-forward` or create a secure ingress
-    expose: false
+    expose:
+      default: false
     # -- The exposed port for this service
     exposedPort: 9000
     # -- The port protocol (TCP/UDP)
     protocol: TCP
-    # -- Defines whether the port is exposed on the internal service;
-    # note that ports exposed on the default service are exposed on the internal
-    # service by default as well.
-    exposeInternal: false
   web:
     ## -- Enable this entrypoint as a default entrypoint. When a service doesn't explicitly set an entrypoint it will only use this entrypoint.
     # asDefault: true
     port: 8000
     # hostPort: 8000
     # containerPort: 8000
-    expose: true
+    expose:
+      default: true
     exposedPort: 80
     ## -- Different target traefik port on the cluster, useful for IP type LB
     # targetPort: 80
@@ -650,10 +651,6 @@ ports:
     # -- Use nodeport if set. This is useful if you have configured Traefik in a
     # LoadBalancer.
     # nodePort: 32080
-    # -- Defines whether the port is exposed on the internal service;
-    # note that ports exposed on the default service are exposed on the internal
-    # service by default as well.
-    exposeInternal: false
     # Port Redirections
     # Added in 2.2, you can make permanent redirects via entrypoints.
     # https://docs.traefik.io/routing/entrypoints/#redirection
@@ -677,17 +674,14 @@ ports:
     port: 8443
     # hostPort: 8443
     # containerPort: 8443
-    expose: true
+    expose:
+      default: true
     exposedPort: 443
     ## -- Different target traefik port on the cluster, useful for IP type LB
     # targetPort: 80
     ## -- The port protocol (TCP/UDP)
     protocol: TCP
     # nodePort: 32443
-    # -- Defines whether the port is exposed on the internal service;
-    # note that ports exposed on the default service are exposed on the internal
-    # service by default as well.
-    exposeInternal: false
     ## -- Specify an application protocol. This may be used as a hint for a Layer 7 load balancer.
     # appProtocol: https
     #
@@ -744,15 +738,12 @@ ports:
     # -- You may not want to expose the metrics port on production deployments.
     # If you want to access it from outside your cluster,
     # use `kubectl port-forward` or create a secure ingress
-    expose: false
+    expose:
+      default: false
     # -- The exposed port for this service
     exposedPort: 9100
     # -- The port protocol (TCP/UDP)
     protocol: TCP
-    # -- Defines whether the port is exposed on the internal service;
-    # note that ports exposed on the default service are exposed on the internal
-    # service by default as well.
-    exposeInternal: false

 # -- TLS Options are created as TLSOption CRDs
 # https://doc.traefik.io/traefik/https/tls/#tls-options
@@ -814,6 +805,7 @@ service:
   #   - IPv4
   #   - IPv6
   ##
+  additionalServices: {}
   ## -- An additional and optional internal Service.
   ## Same parameters as external Service
   # internal:
@@ -899,11 +891,14 @@ hostNetwork: false
 rbac:
   enabled: true
   # If set to false, installs ClusterRole and ClusterRoleBinding so Traefik can be used across namespaces.
-  # If set to true, installs Role and RoleBinding. Providers will only watch target namespace.
+  # If set to true, installs Role and RoleBinding instead of ClusterRole/ClusterRoleBinding. Providers will only watch target namespace.
+  # When combined with providers.kubernetesIngress.disableIngressClassLookup: true and Traefik V3, ClusterRole to watch IngressClass is also disabled.
   namespaced: false
   # Enable user-facing roles
   # https://kubernetes.io/docs/reference/access-authn-authz/rbac/#user-facing-roles
   # aggregateTo: [ "admin" ]
+  # List of Kubernetes secrets that are accessible for Traefik. If empty, then access is granted to every secret.
+  secretResourceNames: []

 # -- Enable to create a PodSecurityPolicy and assign it to the Service Account via RoleBinding or ClusterRoleBinding
 podSecurityPolicy:
```

## 26.1.0  ![AppVersion: v2.11.0](https://img.shields.io/static/v1?label=AppVersion&message=v2.11.0&color=success&logo=) ![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.16.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2024-02-19

* fix: 🐛 set runtimeClassName at pod level
* fix: 🐛 missing quote on experimental plugin args
* fix: update traefik v3 serverstransporttcps CRD
* feat: set runtimeClassName on pod spec
* feat: create v1 Gateway and GatewayClass Version for Traefik v3
* feat: allow exposure of ports on internal service only
* doc: fix invalid suggestion on TLSOption (#996)
* chore: 🔧 update maintainers
* chore: 🔧 promote jnoordsij to Traefik Helm Chart maintainer
* chore(release): 🚀 publish v26.1.0
* chore(deps): update traefik docker tag to v2.11.0
* chore(deps): update traefik docker tag to v2.10.7
* chore(crds): update definitions for traefik v2.11

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index f9dac91..dbd078f 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -100,6 +100,8 @@ deployment:
   #     port: 9000
   #     host: localhost
   #     scheme: HTTP
+  # -- Set a runtimeClassName on pod
+  runtimeClassName:

 # -- Pod disruption budget
 podDisruptionBudget:
@@ -629,6 +631,10 @@ ports:
     exposedPort: 9000
     # -- The port protocol (TCP/UDP)
     protocol: TCP
+    # -- Defines whether the port is exposed on the internal service;
+    # note that ports exposed on the default service are exposed on the internal
+    # service by default as well.
+    exposeInternal: false
   web:
     ## -- Enable this entrypoint as a default entrypoint. When a service doesn't explicitly set an entrypoint it will only use this entrypoint.
     # asDefault: true
@@ -644,6 +650,10 @@ ports:
     # -- Use nodeport if set. This is useful if you have configured Traefik in a
     # LoadBalancer.
     # nodePort: 32080
+    # -- Defines whether the port is exposed on the internal service;
+    # note that ports exposed on the default service are exposed on the internal
+    # service by default as well.
+    exposeInternal: false
     # Port Redirections
     # Added in 2.2, you can make permanent redirects via entrypoints.
     # https://docs.traefik.io/routing/entrypoints/#redirection
@@ -674,6 +684,10 @@ ports:
     ## -- The port protocol (TCP/UDP)
     protocol: TCP
     # nodePort: 32443
+    # -- Defines whether the port is exposed on the internal service;
+    # note that ports exposed on the default service are exposed on the internal
+    # service by default as well.
+    exposeInternal: false
     ## -- Specify an application protocol. This may be used as a hint for a Layer 7 load balancer.
     # appProtocol: https
     #
@@ -735,6 +749,10 @@ ports:
     exposedPort: 9100
     # -- The port protocol (TCP/UDP)
     protocol: TCP
+    # -- Defines whether the port is exposed on the internal service;
+    # note that ports exposed on the default service are exposed on the internal
+    # service by default as well.
+    exposeInternal: false

 # -- TLS Options are created as TLSOption CRDs
 # https://doc.traefik.io/traefik/https/tls/#tls-options
@@ -745,7 +763,7 @@ ports:
 #     labels: {}
 #     sniStrict: true
 #     preferServerCipherSuites: true
-#   customOptions:
+#   custom-options:
 #     labels: {}
 #     curvePreferences:
 #       - CurveP521
@@ -796,7 +814,7 @@ service:
   #   - IPv4
   #   - IPv6
   ##
-  ## -- An additionnal and optional internal Service.
+  ## -- An additional and optional internal Service.
   ## Same parameters as external Service
   # internal:
   #   type: ClusterIP
```

## 26.0.0  ![AppVersion: v2.10.6](https://img.shields.io/static/v1?label=AppVersion&message=v2.10.6&color=success&logo=) ![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.16.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2023-12-05

* fix: 🐛 improve confusing suggested value on openTelemetry.grpc
* fix: 🐛 declare http3 udp port, with or without hostport
* feat: 💥 deployment.podannotations support interpolation with tpl
* feat: allow update of namespace policy for websecure listener
* feat: allow defining startupProbe
* feat: add file provider
* feat: :boom: unify plugin import between traefik and this chart
* chore(release): 🚀 publish v26
* chore(deps): update traefik docker tag to v2.10.6
* Release namespace for Prometheus Operator resources

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 71e377e..f9dac91 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -40,6 +40,7 @@ deployment:
   # -- Additional deployment labels (e.g. for filtering deployment by custom labels)
   labels: {}
   # -- Additional pod annotations (e.g. for mesh injection or prometheus scraping)
+  # It supports templating. One can set it with values like traefik/name: '{{ template "traefik.name" . }}'
   podAnnotations: {}
   # -- Additional Pod labels (e.g. for filtering Pod by custom labels)
   podLabels: {}
@@ -119,10 +120,12 @@ experimental:
   # This value is no longer used, set the image.tag to a semver higher than 3.0, e.g. "v3.0.0-beta3"
   # v3:
   # -- Enable traefik version 3
-  #  enabled: false
-  plugins:
-    # -- Enable traefik experimental plugins
-    enabled: false
+
+  # -- Enable traefik experimental plugins
+  plugins: {}
+  # demo:
+  #   moduleName: github.com/traefik/plugindemo
+  #   version: v0.2.1
   kubernetesGateway:
     # -- Enable traefik experimental GatewayClass CRD
     enabled: false
@@ -206,6 +209,17 @@ livenessProbe:
   # -- The number of seconds to wait for a probe response before considering it as failed.
   timeoutSeconds: 2

+# -- Define Startup Probe for container: https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/#define-startup-probes
+# eg.
+# `startupProbe:
+#   exec:
+#     command:
+#       - mycommand
+#       - foo
+#   initialDelaySeconds: 5
+#   periodSeconds: 5`
+startupProbe:
+
 providers:
   kubernetesCRD:
     # -- Load Kubernetes IngressRoute provider
@@ -241,6 +255,23 @@ providers:
       # By default this Traefik service
       # pathOverride: ""

+  file:
+    # -- Create a file provider
+    enabled: false
+    # -- Allows Traefik to automatically watch for file changes
+    watch: true
+    # -- File content (YAML format, go template supported) (see https://doc.traefik.io/traefik/providers/file/)
+    content: ""
+      # http:
+      #   routers:
+      #     router0:
+      #       entryPoints:
+      #       - web
+      #       middlewares:
+      #       - my-basic-auth
+      #       service: service-foo
+      #       rule: Path(`/foo`)
+
 #
 # -- Add volumes to the traefik pod. The volume name will be passed to tpl.
 # This can be used to mount a cert pair or a configmap that holds a config.toml file.
@@ -487,7 +518,7 @@ metrics:
 # -- https://doc.traefik.io/traefik/observability/tracing/overview/
 tracing: {}
 #  openTelemetry: # traefik v3+ only
-#    grpc: {}
+#    grpc: true
 #    insecure: true
 #    address: localhost:4317
 # instana:
```

## 25.0.0  ![AppVersion: v2.10.5](https://img.shields.io/static/v1?label=AppVersion&message=v2.10.5&color=success&logo=) ![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.16.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2023-10-23

* revert: "fix: 🐛 remove old CRDs using traefik.containo.us"
* fix: 🐛 remove old CRDs using traefik.containo.us
* fix: disable ClusterRole and ClusterRoleBinding when not needed
* fix: detect correctly v3 version when using sha in `image.tag`
* fix: allow updateStrategy.rollingUpdate.maxUnavailable to be passed in as an int or string
* fix: add missing separator in crds
* fix: add Prometheus scraping annotations only if serviceMonitor not created
* feat: ✨ add healthcheck ingressRoute
* feat: :boom: support http redirections and http challenges with cert-manager
* feat: :boom: rework and allow update of namespace policy for Gateway
* docs: Fix typo in the default values file
* chore: remove label whitespace at TLSOption
* chore(release): publish v25.0.0
* chore(deps): update traefik docker tag to v2.10.5
* chore(deps): update docker.io/helmunittest/helm-unittest docker tag to v3.12.3
* chore(ci): 🔧 👷 add e2e test when releasing

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index aeec85c..71e377e 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -45,60 +45,60 @@ deployment:
   podLabels: {}
   # -- Additional containers (e.g. for metric offloading sidecars)
   additionalContainers: []
-    # https://docs.datadoghq.com/developers/dogstatsd/unix_socket/?tab=host
-    # - name: socat-proxy
-    #   image: alpine/socat:1.0.5
-    #   args: ["-s", "-u", "udp-recv:8125", "unix-sendto:/socket/socket"]
-    #   volumeMounts:
-    #     - name: dsdsocket
-    #       mountPath: /socket
+  # https://docs.datadoghq.com/developers/dogstatsd/unix_socket/?tab=host
+  # - name: socat-proxy
+  #   image: alpine/socat:1.0.5
+  #   args: ["-s", "-u", "udp-recv:8125", "unix-sendto:/socket/socket"]
+  #   volumeMounts:
+  #     - name: dsdsocket
+  #       mountPath: /socket
   # -- Additional volumes available for use with initContainers and additionalContainers
   additionalVolumes: []
-    # - name: dsdsocket
-    #   hostPath:
-    #     path: /var/run/statsd-exporter
+  # - name: dsdsocket
+  #   hostPath:
+  #     path: /var/run/statsd-exporter
   # -- Additional initContainers (e.g. for setting file permission as shown below)
   initContainers: []
-    # The "volume-permissions" init container is required if you run into permission issues.
-    # Related issue: https://github.com/traefik/traefik-helm-chart/issues/396
-    # - name: volume-permissions
-    #   image: busybox:latest
-    #   command: ["sh", "-c", "touch /data/acme.json; chmod -v 600 /data/acme.json"]
-    #   securityContext:
-    #     runAsNonRoot: true
-    #     runAsGroup: 65532
-    #     runAsUser: 65532
-    #   volumeMounts:
-    #     - name: data
-    #       mountPath: /data
+  # The "volume-permissions" init container is required if you run into permission issues.
+  # Related issue: https://github.com/traefik/traefik-helm-chart/issues/396
+  # - name: volume-permissions
+  #   image: busybox:latest
+  #   command: ["sh", "-c", "touch /data/acme.json; chmod -v 600 /data/acme.json"]
+  #   securityContext:
+  #     runAsNonRoot: true
+  #     runAsGroup: 65532
+  #     runAsUser: 65532
+  #   volumeMounts:
+  #     - name: data
+  #       mountPath: /data
   # -- Use process namespace sharing
   shareProcessNamespace: false
   # -- Custom pod DNS policy. Apply if `hostNetwork: true`
   # dnsPolicy: ClusterFirstWithHostNet
   dnsConfig: {}
-    # nameservers:
-    #   - 192.0.2.1 # this is an example
-    # searches:
-    #   - ns1.svc.cluster-domain.example
-    #   - my.dns.search.suffix
-    # options:
-    #   - name: ndots
-    #     value: "2"
-    #   - name: edns0
+  # nameservers:
+  #   - 192.0.2.1 # this is an example
+  # searches:
+  #   - ns1.svc.cluster-domain.example
+  #   - my.dns.search.suffix
+  # options:
+  #   - name: ndots
+  #     value: "2"
+  #   - name: edns0
   # -- Additional imagePullSecrets
   imagePullSecrets: []
-    # - name: myRegistryKeySecretName
+  # - name: myRegistryKeySecretName
   # -- Pod lifecycle actions
   lifecycle: {}
-    # preStop:
-    #   exec:
-    #     command: ["/bin/sh", "-c", "sleep 40"]
-    # postStart:
-    #   httpGet:
-    #     path: /ping
-    #     port: 9000
-    #     host: localhost
-    #     scheme: HTTP
+  # preStop:
+  #   exec:
+  #     command: ["/bin/sh", "-c", "sleep 40"]
+  # postStart:
+  #   httpGet:
+  #     path: /ping
+  #     port: 9000
+  #     host: localhost
+  #     scheme: HTTP

 # -- Pod disruption budget
 podDisruptionBudget:
@@ -116,9 +116,9 @@ ingressClass:

 # Traefik experimental features
 experimental:
-  #This value is no longer used, set the image.tag to a semver higher than 3.0, e.g. "v3.0.0-beta3"
-  #v3:
-    # -- Enable traefik version 3
+  # This value is no longer used, set the image.tag to a semver higher than 3.0, e.g. "v3.0.0-beta3"
+  # v3:
+  # -- Enable traefik version 3
   #  enabled: false
   plugins:
     # -- Enable traefik experimental plugins
@@ -126,9 +126,9 @@ experimental:
   kubernetesGateway:
     # -- Enable traefik experimental GatewayClass CRD
     enabled: false
-    gateway:
-      # -- Enable traefik regular kubernetes gateway
-      enabled: true
+    ## Routes are restricted to namespace of the gateway by default.
+    ## https://gateway-api.sigs.k8s.io/references/spec/#gateway.networking.k8s.io/v1beta1.FromNamespaces
+    # namespacePolicy: All
     # certificate:
     #   group: "core"
     #   kind: "Secret"
@@ -159,6 +159,22 @@ ingressRoute:
     middlewares: []
     # -- TLS options (e.g. secret containing certificate)
     tls: {}
+  healthcheck:
+    # -- Create an IngressRoute for the healthcheck probe
+    enabled: false
+    # -- Additional ingressRoute annotations (e.g. for kubernetes.io/ingress.class)
+    annotations: {}
+    # -- Additional ingressRoute labels (e.g. for filtering IngressRoute by custom labels)
+    labels: {}
+    # -- The router match rule used for the healthcheck ingressRoute
+    matchRule: PathPrefix(`/ping`)
+    # -- Specify the allowed entrypoints to use for the healthcheck ingress route, (e.g. traefik, web, websecure).
+    # By default, it's using traefik entrypoint, which is not exposed.
+    entryPoints: ["traefik"]
+    # -- Additional ingressRoute middlewares (e.g. for authentication)
+    middlewares: []
+    # -- TLS options (e.g. secret containing certificate)
+    tls: {}

 updateStrategy:
   # -- Customize updateStrategy: RollingUpdate or OnDelete
@@ -204,10 +220,10 @@ providers:
     # labelSelector: environment=production,method=traefik
     # -- Array of namespaces to watch. If left empty, Traefik watches all namespaces.
     namespaces: []
-      # - "default"
+    # - "default"

   kubernetesIngress:
-    # -- Load Kubernetes IngressRoute provider
+    # -- Load Kubernetes Ingress provider
     enabled: true
     # -- Allows to reference ExternalName services in Ingress
     allowExternalNameServices: false
@@ -217,7 +233,7 @@ providers:
     # labelSelector: environment=production,method=traefik
     # -- Array of namespaces to watch. If left empty, Traefik watches all namespaces.
     namespaces: []
-      # - "default"
+    # - "default"
     # IP used for Kubernetes Ingress endpoints
     publishedService:
       enabled: false
@@ -243,9 +259,9 @@ volumes: []

 # -- Additional volumeMounts to add to the Traefik container
 additionalVolumeMounts: []
-  # -- For instance when using a logshipper for access logs
-  # - name: traefik-logs
-  #   mountPath: /var/log/traefik
+# -- For instance when using a logshipper for access logs
+# - name: traefik-logs
+#   mountPath: /var/log/traefik

 logs:
   general:
@@ -270,26 +286,26 @@ logs:
     ## Filtering
     # -- https://docs.traefik.io/observability/access-logs/#filtering
     filters: {}
-      # statuscodes: "200,300-302"
-      # retryattempts: true
-      # minduration: 10ms
+    # statuscodes: "200,300-302"
+    # retryattempts: true
+    # minduration: 10ms
     fields:
       general:
         # -- Available modes: keep, drop, redact.
         defaultmode: keep
         # -- Names of the fields to limit.
         names: {}
-          ## Examples:
-          # ClientUsername: drop
+        ## Examples:
+        # ClientUsername: drop
       headers:
         # -- Available modes: keep, drop, redact.
         defaultmode: drop
         # -- Names of the headers to limit.
         names: {}
-          ## Examples:
-          # User-Agent: redact
-          # Authorization: drop
-          # Content-Type: keep
+        ## Examples:
+        # User-Agent: redact
+        # Authorization: drop
+        # Content-Type: keep

 metrics:
   ## -- Prometheus is enabled by default.
@@ -308,118 +324,118 @@ metrics:
     ## When manualRouting is true, it disables the default internal router in
     ## order to allow creating a custom router for prometheus@internal service.
     # manualRouting: true
-#  datadog:
-#    ## Address instructs exporter to send metrics to datadog-agent at this address.
-#    address: "127.0.0.1:8125"
-#    ## The interval used by the exporter to push metrics to datadog-agent. Default=10s
-#    # pushInterval: 30s
-#    ## The prefix to use for metrics collection. Default="traefik"
-#    # prefix: traefik
-#    ## Enable metrics on entry points. Default=true
-#    # addEntryPointsLabels: false
-#    ## Enable metrics on routers. Default=false
-#    # addRoutersLabels: true
-#    ## Enable metrics on services. Default=true
-#    # addServicesLabels: false
-#  influxdb:
-#    ## Address instructs exporter to send metrics to influxdb at this address.
-#    address: localhost:8089
-#    ## InfluxDB's address protocol (udp or http). Default="udp"
-#    protocol: udp
-#    ## InfluxDB database used when protocol is http. Default=""
-#    # database: ""
-#    ## InfluxDB retention policy used when protocol is http. Default=""
-#    # retentionPolicy: ""
-#    ## InfluxDB username (only with http). Default=""
-#    # username: ""
-#    ## InfluxDB password (only with http). Default=""
-#    # password: ""
-#    ## The interval used by the exporter to push metrics to influxdb. Default=10s
-#    # pushInterval: 30s
-#    ## Additional labels (influxdb tags) on all metrics.
-#    # additionalLabels:
-#    #   env: production
-#    #   foo: bar
-#    ## Enable metrics on entry points. Default=true
-#    # addEntryPointsLabels: false
-#    ## Enable metrics on routers. Default=false
-#    # addRoutersLabels: true
-#    ## Enable metrics on services. Default=true
-#    # addServicesLabels: false
-#  influxdb2:
-#    ## Address instructs exporter to send metrics to influxdb v2 at this address.
-#    address: localhost:8086
-#    ## Token with which to connect to InfluxDB v2.
-#    token: xxx
-#    ## Organisation where metrics will be stored.
-#    org: ""
-#    ## Bucket where metrics will be stored.
-#    bucket: ""
-#    ## The interval used by the exporter to push metrics to influxdb. Default=10s
-#    # pushInterval: 30s
-#    ## Additional labels (influxdb tags) on all metrics.
-#    # additionalLabels:
-#    #   env: production
-#    #   foo: bar
-#    ## Enable metrics on entry points. Default=true
-#    # addEntryPointsLabels: false
-#    ## Enable metrics on routers. Default=false
-#    # addRoutersLabels: true
-#    ## Enable metrics on services. Default=true
-#    # addServicesLabels: false
-#  statsd:
-#    ## Address instructs exporter to send metrics to statsd at this address.
-#    address: localhost:8125
-#    ## The interval used by the exporter to push metrics to influxdb. Default=10s
-#    # pushInterval: 30s
-#    ## The prefix to use for metrics collection. Default="traefik"
-#    # prefix: traefik
-#    ## Enable metrics on entry points. Default=true
-#    # addEntryPointsLabels: false
-#    ## Enable metrics on routers. Default=false
-#    # addRoutersLabels: true
-#    ## Enable metrics on services. Default=true
-#    # addServicesLabels: false
-#  openTelemetry:
-#    ## Address of the OpenTelemetry Collector to send metrics to.
-#    address: "localhost:4318"
-#    ## Enable metrics on entry points.
-#    addEntryPointsLabels: true
-#    ## Enable metrics on routers.
-#    addRoutersLabels: true
-#    ## Enable metrics on services.
-#    addServicesLabels: true
-#    ## Explicit boundaries for Histogram data points.
-#    explicitBoundaries:
-#      - "0.1"
-#      - "0.3"
-#      - "1.2"
-#      - "5.0"
-#    ## Additional headers sent with metrics by the reporter to the OpenTelemetry Collector.
-#    headers:
-#      foo: bar
-#      test: test
-#    ## Allows reporter to send metrics to the OpenTelemetry Collector without using a secured protocol.
-#    insecure: true
-#    ## Interval at which metrics are sent to the OpenTelemetry Collector.
-#    pushInterval: 10s
-#    ## Allows to override the default URL path used for sending metrics. This option has no effect when using gRPC transport.
-#    path: /foo/v1/traces
-#    ## Defines the TLS configuration used by the reporter to send metrics to the OpenTelemetry Collector.
-#    tls:
-#      ## The path to the certificate authority, it defaults to the system bundle.
-#      ca: path/to/ca.crt
-#      ## The path to the public certificate. When using this option, setting the key option is required.
-#      cert: path/to/foo.cert
-#      ## The path to the private key. When using this option, setting the cert option is required.
-#      key: path/to/key.key
-#      ## If set to true, the TLS connection accepts any certificate presented by the server regardless of the hostnames it covers.
-#      insecureSkipVerify: true
-#    ## This instructs the reporter to send metrics to the OpenTelemetry Collector using gRPC.
-#    grpc: true
-
-## -- enable optional CRDs for Prometheus Operator
-##
+  #  datadog:
+  #    ## Address instructs exporter to send metrics to datadog-agent at this address.
+  #    address: "127.0.0.1:8125"
+  #    ## The interval used by the exporter to push metrics to datadog-agent. Default=10s
+  #    # pushInterval: 30s
+  #    ## The prefix to use for metrics collection. Default="traefik"
+  #    # prefix: traefik
+  #    ## Enable metrics on entry points. Default=true
+  #    # addEntryPointsLabels: false
+  #    ## Enable metrics on routers. Default=false
+  #    # addRoutersLabels: true
+  #    ## Enable metrics on services. Default=true
+  #    # addServicesLabels: false
+  #  influxdb:
+  #    ## Address instructs exporter to send metrics to influxdb at this address.
+  #    address: localhost:8089
+  #    ## InfluxDB's address protocol (udp or http). Default="udp"
+  #    protocol: udp
+  #    ## InfluxDB database used when protocol is http. Default=""
+  #    # database: ""
+  #    ## InfluxDB retention policy used when protocol is http. Default=""
+  #    # retentionPolicy: ""
+  #    ## InfluxDB username (only with http). Default=""
+  #    # username: ""
+  #    ## InfluxDB password (only with http). Default=""
+  #    # password: ""
+  #    ## The interval used by the exporter to push metrics to influxdb. Default=10s
+  #    # pushInterval: 30s
+  #    ## Additional labels (influxdb tags) on all metrics.
+  #    # additionalLabels:
+  #    #   env: production
+  #    #   foo: bar
+  #    ## Enable metrics on entry points. Default=true
+  #    # addEntryPointsLabels: false
+  #    ## Enable metrics on routers. Default=false
+  #    # addRoutersLabels: true
+  #    ## Enable metrics on services. Default=true
+  #    # addServicesLabels: false
+  #  influxdb2:
+  #    ## Address instructs exporter to send metrics to influxdb v2 at this address.
+  #    address: localhost:8086
+  #    ## Token with which to connect to InfluxDB v2.
+  #    token: xxx
+  #    ## Organisation where metrics will be stored.
+  #    org: ""
+  #    ## Bucket where metrics will be stored.
+  #    bucket: ""
+  #    ## The interval used by the exporter to push metrics to influxdb. Default=10s
+  #    # pushInterval: 30s
+  #    ## Additional labels (influxdb tags) on all metrics.
+  #    # additionalLabels:
+  #    #   env: production
+  #    #   foo: bar
+  #    ## Enable metrics on entry points. Default=true
+  #    # addEntryPointsLabels: false
+  #    ## Enable metrics on routers. Default=false
+  #    # addRoutersLabels: true
+  #    ## Enable metrics on services. Default=true
+  #    # addServicesLabels: false
+  #  statsd:
+  #    ## Address instructs exporter to send metrics to statsd at this address.
+  #    address: localhost:8125
+  #    ## The interval used by the exporter to push metrics to influxdb. Default=10s
+  #    # pushInterval: 30s
+  #    ## The prefix to use for metrics collection. Default="traefik"
+  #    # prefix: traefik
+  #    ## Enable metrics on entry points. Default=true
+  #    # addEntryPointsLabels: false
+  #    ## Enable metrics on routers. Default=false
+  #    # addRoutersLabels: true
+  #    ## Enable metrics on services. Default=true
+  #    # addServicesLabels: false
+  #  openTelemetry:
+  #    ## Address of the OpenTelemetry Collector to send metrics to.
+  #    address: "localhost:4318"
+  #    ## Enable metrics on entry points.
+  #    addEntryPointsLabels: true
+  #    ## Enable metrics on routers.
+  #    addRoutersLabels: true
+  #    ## Enable metrics on services.
+  #    addServicesLabels: true
+  #    ## Explicit boundaries for Histogram data points.
+  #    explicitBoundaries:
+  #      - "0.1"
+  #      - "0.3"
+  #      - "1.2"
+  #      - "5.0"
+  #    ## Additional headers sent with metrics by the reporter to the OpenTelemetry Collector.
+  #    headers:
+  #      foo: bar
+  #      test: test
+  #    ## Allows reporter to send metrics to the OpenTelemetry Collector without using a secured protocol.
+  #    insecure: true
+  #    ## Interval at which metrics are sent to the OpenTelemetry Collector.
+  #    pushInterval: 10s
+  #    ## Allows to override the default URL path used for sending metrics. This option has no effect when using gRPC transport.
+  #    path: /foo/v1/traces
+  #    ## Defines the TLS configuration used by the reporter to send metrics to the OpenTelemetry Collector.
+  #    tls:
+  #      ## The path to the certificate authority, it defaults to the system bundle.
+  #      ca: path/to/ca.crt
+  #      ## The path to the public certificate. When using this option, setting the key option is required.
+  #      cert: path/to/foo.cert
+  #      ## The path to the private key. When using this option, setting the cert option is required.
+  #      key: path/to/key.key
+  #      ## If set to true, the TLS connection accepts any certificate presented by the server regardless of the hostnames it covers.
+  #      insecureSkipVerify: true
+  #    ## This instructs the reporter to send metrics to the OpenTelemetry Collector using gRPC.
+  #    grpc: true
+
+  ## -- enable optional CRDs for Prometheus Operator
+  ##
   ## Create a dedicated metrics service for use with ServiceMonitor
   #  service:
   #    enabled: false
@@ -470,55 +486,55 @@ metrics:
 ## Tracing
 # -- https://doc.traefik.io/traefik/observability/tracing/overview/
 tracing: {}
-  #  openTelemetry: # traefik v3+ only
-  #    grpc: {}
-  #    insecure: true
-  #    address: localhost:4317
-  # instana:
-  #   localAgentHost: 127.0.0.1
-  #   localAgentPort: 42699
-  #   logLevel: info
-  #   enableAutoProfile: true
-  # datadog:
-  #   localAgentHostPort: 127.0.0.1:8126
-  #   debug: false
-  #   globalTag: ""
-  #   prioritySampling: false
-  # jaeger:
-  #   samplingServerURL: http://localhost:5778/sampling
-  #   samplingType: const
-  #   samplingParam: 1.0
-  #   localAgentHostPort: 127.0.0.1:6831
-  #   gen128Bit: false
-  #   propagation: jaeger
-  #   traceContextHeaderName: uber-trace-id
-  #   disableAttemptReconnecting: true
-  #   collector:
-  #      endpoint: ""
-  #      user: ""
-  #      password: ""
-  # zipkin:
-  #   httpEndpoint: http://localhost:9411/api/v2/spans
-  #   sameSpan: false
-  #   id128Bit: true
-  #   sampleRate: 1.0
-  # haystack:
-  #   localAgentHost: 127.0.0.1
-  #   localAgentPort: 35000
-  #   globalTag: ""
-  #   traceIDHeaderName: ""
-  #   parentIDHeaderName: ""
-  #   spanIDHeaderName: ""
-  #   baggagePrefixHeaderName: ""
-  # elastic:
-  #   serverURL: http://localhost:8200
-  #   secretToken: ""
-  #   serviceEnvironment: ""
+#  openTelemetry: # traefik v3+ only
+#    grpc: {}
+#    insecure: true
+#    address: localhost:4317
+# instana:
+#   localAgentHost: 127.0.0.1
+#   localAgentPort: 42699
+#   logLevel: info
+#   enableAutoProfile: true
+# datadog:
+#   localAgentHostPort: 127.0.0.1:8126
+#   debug: false
+#   globalTag: ""
+#   prioritySampling: false
+# jaeger:
+#   samplingServerURL: http://localhost:5778/sampling
+#   samplingType: const
+#   samplingParam: 1.0
+#   localAgentHostPort: 127.0.0.1:6831
+#   gen128Bit: false
+#   propagation: jaeger
+#   traceContextHeaderName: uber-trace-id
+#   disableAttemptReconnecting: true
+#   collector:
+#      endpoint: ""
+#      user: ""
+#      password: ""
+# zipkin:
+#   httpEndpoint: http://localhost:9411/api/v2/spans
+#   sameSpan: false
+#   id128Bit: true
+#   sampleRate: 1.0
+# haystack:
+#   localAgentHost: 127.0.0.1
+#   localAgentPort: 35000
+#   globalTag: ""
+#   traceIDHeaderName: ""
+#   parentIDHeaderName: ""
+#   spanIDHeaderName: ""
+#   baggagePrefixHeaderName: ""
+# elastic:
+#   serverURL: http://localhost:8200
+#   secretToken: ""
+#   serviceEnvironment: ""

 # -- Global command arguments to be passed to all traefik's pods
 globalArguments:
-  - "--global.checknewversion"
-  - "--global.sendanonymoususage"
+- "--global.checknewversion"
+- "--global.sendanonymoususage"

 #
 # Configure Traefik static configuration
@@ -531,14 +547,14 @@ additionalArguments: []

 # -- Environment variables to be passed to Traefik's binary
 env:
-  - name: POD_NAME
-    valueFrom:
-      fieldRef:
-        fieldPath: metadata.name
-  - name: POD_NAMESPACE
-    valueFrom:
-      fieldRef:
-        fieldPath: metadata.namespace
+- name: POD_NAME
+  valueFrom:
+    fieldRef:
+      fieldPath: metadata.name
+- name: POD_NAMESPACE
+  valueFrom:
+    fieldRef:
+      fieldPath: metadata.namespace
 # - name: SOME_VAR
 #   value: some-var-value
 # - name: SOME_VAR_FROM_CONFIG_MAP
@@ -600,7 +616,10 @@ ports:
     # Port Redirections
     # Added in 2.2, you can make permanent redirects via entrypoints.
     # https://docs.traefik.io/routing/entrypoints/#redirection
-    # redirectTo: websecure
+    # redirectTo:
+    #   port: websecure
+    #   (Optional)
+    #   priority: 10
     #
     # Trust forwarded  headers information (X-Forwarded-*).
     # forwardedHeaders:
@@ -638,14 +657,14 @@ ports:
     # advertisedPort: 4443
     #
     ## -- Trust forwarded  headers information (X-Forwarded-*).
-    #forwardedHeaders:
-    #  trustedIPs: []
-    #  insecure: false
+    # forwardedHeaders:
+    #   trustedIPs: []
+    #   insecure: false
     #
     ## -- Enable the Proxy Protocol header parsing for the entry point
-    #proxyProtocol:
-    #  trustedIPs: []
-    #  insecure: false
+    # proxyProtocol:
+    #   trustedIPs: []
+    #   insecure: false
     #
     ## Set TLS at the entrypoint
     ## https://doc.traefik.io/traefik/routing/entrypoints/#tls
@@ -728,16 +747,16 @@ service:
   # -- Additional entries here will be added to the service spec.
   # -- Cannot contain type, selector or ports entries.
   spec: {}
-    # externalTrafficPolicy: Cluster
-    # loadBalancerIP: "1.2.3.4"
-    # clusterIP: "2.3.4.5"
+  # externalTrafficPolicy: Cluster
+  # loadBalancerIP: "1.2.3.4"
+  # clusterIP: "2.3.4.5"
   loadBalancerSourceRanges: []
-    # - 192.168.0.1/32
-    # - 172.16.0.0/16
+  # - 192.168.0.1/32
+  # - 172.16.0.0/16
   ## -- Class of the load balancer implementation
   # loadBalancerClass: service.k8s.aws/nlb
   externalIPs: []
-    # - 1.2.3.4
+  # - 1.2.3.4
   ## One of SingleStack, PreferDualStack, or RequireDualStack.
   # ipFamilyPolicy: SingleStack
   ## List of IP families (e.g. IPv4 and/or IPv6).
@@ -789,7 +808,7 @@ persistence:
   # It can be used to store TLS certificates, see `storage` in certResolvers
   enabled: false
   name: data
-#  existingClaim: ""
+  #  existingClaim: ""
   accessMode: ReadWriteOnce
   size: 128Mi
   # storageClass: ""
@@ -852,12 +871,12 @@ serviceAccountAnnotations: {}

 # -- The resources parameter defines CPU and memory requirements and limits for Traefik's containers.
 resources: {}
-  # requests:
-  #   cpu: "100m"
-  #   memory: "50Mi"
-  # limits:
-  #   cpu: "300m"
-  #   memory: "150Mi"
+# requests:
+#   cpu: "100m"
+#   memory: "50Mi"
+# limits:
+#   cpu: "300m"
+#   memory: "150Mi"

 # -- This example pod anti-affinity forces the scheduler to put traefik pods
 # -- on nodes where no other traefik pods are scheduled.
```

## 24.0.0  ![AppVersion: v2.10.4](https://img.shields.io/static/v1?label=AppVersion&message=v2.10.4&color=success&logo=) ![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.16.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2023-08-10

* fix: 💥 BREAKING CHANGE on healthchecks and traefik port
* fix: tracing.opentelemetry.tls is optional for all values
* fix: http3 support broken when advertisedPort set
* feat: multi namespace RBAC manifests
* chore(tests): 🔧 fix typo on tracing test
* chore(release): 🚀 publish v24.0.0
* chore(deps): update docker.io/helmunittest/helm-unittest docker tag to v3.12.2

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 947ba56..aeec85c 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -28,6 +28,13 @@ deployment:
   terminationGracePeriodSeconds: 60
   # -- The minimum number of seconds Traefik needs to be up and running before the DaemonSet/Deployment controller considers it available
   minReadySeconds: 0
+  ## Override the liveness/readiness port. This is useful to integrate traefik
+  ## with an external Load Balancer that performs healthchecks.
+  ## Default: ports.traefik.port
+  # healthchecksPort: 9000
+  ## Override the liveness/readiness scheme. Useful for getting ping to
+  ## respond on websecure entryPoint.
+  # healthchecksScheme: HTTPS
   # -- Additional deployment annotations (e.g. for jaeger-operator sidecar injection)
   annotations: {}
   # -- Additional deployment labels (e.g. for filtering deployment by custom labels)
@@ -112,7 +119,7 @@ experimental:
   #This value is no longer used, set the image.tag to a semver higher than 3.0, e.g. "v3.0.0-beta3"
   #v3:
     # -- Enable traefik version 3
-  #  enabled: false
+  #  enabled: false
   plugins:
     # -- Enable traefik experimental plugins
     enabled: false
@@ -564,15 +571,6 @@ ports:
     # only.
     # hostIP: 192.168.100.10

-    # Override the liveness/readiness port. This is useful to integrate traefik
-    # with an external Load Balancer that performs healthchecks.
-    # Default: ports.traefik.port
-    # healthchecksPort: 9000
-
-    # Override the liveness/readiness scheme. Useful for getting ping to
-    # respond on websecure entryPoint.
-    # healthchecksScheme: HTTPS
-
     # Defines whether the port is exposed if service.type is LoadBalancer or
     # NodePort.
     #
@@ -877,7 +875,7 @@ affinity: {}
 nodeSelector: {}
 # -- Tolerations allow the scheduler to schedule pods with matching taints.
 tolerations: []
-# -- You can use topology spread constraints to control
+# -- You can use topology spread constraints to control
 # how Pods are spread across your cluster among failure-domains.
 topologySpreadConstraints: []
 # This example topologySpreadConstraints forces the scheduler to put traefik pods
```

## 23.2.0  ![AppVersion: v2.10.4](https://img.shields.io/static/v1?label=AppVersion&message=v2.10.4&color=success&logo=) ![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.16.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2023-07-27

* ⬆️ Upgrade traefik Docker tag to v2.10.3
* release: :rocket: publish v23.2.0
* fix: 🐛 update traefik.containo.us CRDs to v2.10
* fix: 🐛 traefik or metrics port can be disabled
* fix: ingressclass name should be customizable (#864)
* feat: ✨ add support for traefik v3.0.0-beta3 and openTelemetry
* feat: disable allowPrivilegeEscalation
* feat: add pod_name as default in values.yaml
* chore(tests): 🔧 use more accurate asserts on refactor'd isNull test
* chore(deps): update traefik docker tag to v2.10.4
* chore(deps): update docker.io/helmunittest/helm-unittest docker tag to v3.11.3

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 345bbd8..947ba56 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -105,12 +105,14 @@ podDisruptionBudget:
 ingressClass:
   enabled: true
   isDefaultClass: true
+  # name: my-custom-class

 # Traefik experimental features
 experimental:
-  v3:
+  #This value is no longer used, set the image.tag to a semver higher than 3.0, e.g. "v3.0.0-beta3"
+  #v3:
     # -- Enable traefik version 3
-    enabled: false
+  #  enabled: false
   plugins:
     # -- Enable traefik experimental plugins
     enabled: false
@@ -461,6 +463,10 @@ metrics:
 ## Tracing
 # -- https://doc.traefik.io/traefik/observability/tracing/overview/
 tracing: {}
+  #  openTelemetry: # traefik v3+ only
+  #    grpc: {}
+  #    insecure: true
+  #    address: localhost:4317
   # instana:
   #   localAgentHost: 127.0.0.1
   #   localAgentPort: 42699
@@ -517,7 +523,15 @@ additionalArguments: []
 #  - "--log.level=DEBUG"

 # -- Environment variables to be passed to Traefik's binary
-env: []
+env:
+  - name: POD_NAME
+    valueFrom:
+      fieldRef:
+        fieldPath: metadata.name
+  - name: POD_NAMESPACE
+    valueFrom:
+      fieldRef:
+        fieldPath: metadata.namespace
 # - name: SOME_VAR
 #   value: some-var-value
 # - name: SOME_VAR_FROM_CONFIG_MAP
@@ -563,7 +577,7 @@ ports:
     # NodePort.
     #
     # -- You SHOULD NOT expose the traefik port on production deployments.
-    # If you want to access it from outside of your cluster,
+    # If you want to access it from outside your cluster,
     # use `kubectl port-forward` or create a secure ingress
     expose: false
     # -- The exposed port for this service
@@ -571,7 +585,7 @@ ports:
     # -- The port protocol (TCP/UDP)
     protocol: TCP
   web:
-    ## -- Enable this entrypoint as a default entrypoint. When a service doesn't explicity set an entrypoint it will only use this entrypoint.
+    ## -- Enable this entrypoint as a default entrypoint. When a service doesn't explicitly set an entrypoint it will only use this entrypoint.
     # asDefault: true
     port: 8000
     # hostPort: 8000
@@ -600,7 +614,7 @@ ports:
     #   trustedIPs: []
     #   insecure: false
   websecure:
-    ## -- Enable this entrypoint as a default entrypoint. When a service doesn't explicity set an entrypoint it will only use this entrypoint.
+    ## -- Enable this entrypoint as a default entrypoint. When a service doesn't explicitly set an entrypoint it will only use this entrypoint.
     # asDefault: true
     port: 8443
     # hostPort: 8443
@@ -666,7 +680,7 @@ ports:
     # NodePort.
     #
     # -- You may not want to expose the metrics port on production deployments.
-    # If you want to access it from outside of your cluster,
+    # If you want to access it from outside your cluster,
     # use `kubectl port-forward` or create a secure ingress
     expose: false
     # -- The exposed port for this service
@@ -880,14 +894,15 @@ topologySpreadConstraints: []
 priorityClassName: ""

 # -- Set the container security context
-# -- To run the container with ports below 1024 this will need to be adjust to run as root
+# -- To run the container with ports below 1024 this will need to be adjusted to run as root
 securityContext:
   capabilities:
     drop: [ALL]
   readOnlyRootFilesystem: true
+  allowPrivilegeEscalation: false

 podSecurityContext:
-  # /!\ When setting fsGroup, Kubernetes will recursively changes ownership and
+  # /!\ When setting fsGroup, Kubernetes will recursively change ownership and
   # permissions for the contents of each volume to match the fsGroup. This can
   # be an issue when storing sensitive content like TLS Certificates /!\
   # fsGroup: 65532
```

## 23.1.0  ![AppVersion: v2.10.1](https://img.shields.io/static/v1?label=AppVersion&message=v2.10.1&color=success&logo=) ![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.16.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2023-06-06

* release: 🚀 publish v23.1.0
* fix: 🐛 use k8s version for hpa api version
* fix: 🐛 http3 support on traefik v3
* fix: use `targetPort` instead of `port` on ServiceMonitor
* feat: ➖ remove Traefik Hub v1 integration
* feat: ✨ add a warning when labelSelector don't match
* feat: common labels for all resources
* feat: allow specifying service loadBalancerClass
* feat: add optional `appProtocol` field on Service ports
* doc: added values README via helm-docs cli

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 71273cc..345bbd8 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -1,70 +1,56 @@
 # Default values for Traefik
 image:
+  # -- Traefik image host registry
   registry: docker.io
+  # -- Traefik image repository
   repository: traefik
-  # defaults to appVersion
+  # -- defaults to appVersion
   tag: ""
+  # -- Traefik image pull policy
   pullPolicy: IfNotPresent

-#
-# Configure integration with Traefik Hub
-#
-hub:
-  ## Enabling Hub will:
-  # * enable Traefik Hub integration on Traefik
-  # * add `traefikhub-tunl` endpoint
-  # * enable Prometheus metrics with addRoutersLabels
-  # * enable allowExternalNameServices on KubernetesIngress provider
-  # * enable allowCrossNamespace on KubernetesCRD provider
-  # * add an internal (ClusterIP) Service, dedicated for Traefik Hub
-  enabled: false
-  ## Default port can be changed
-  # tunnelPort: 9901
-  ## TLS is optional. Insecure is mutually exclusive with any other options
-  # tls:
-  #   insecure: false
-  #   ca: "/path/to/ca.pem"
-  #   cert: "/path/to/cert.pem"
-  #   key: "/path/to/key.pem"
+# -- Add additional label to all resources
+commonLabels: {}

 #
 # Configure the deployment
 #
 deployment:
+  # -- Enable deployment
   enabled: true
-  # Can be either Deployment or DaemonSet
+  # -- Deployment or DaemonSet
   kind: Deployment
-  # Number of pods of the deployment (only applies when kind == Deployment)
+  # -- Number of pods of the deployment (only applies when kind == Deployment)
   replicas: 1
-  # Number of old history to retain to allow rollback (If not set, default Kubernetes value is set to 10)
+  # -- Number of old history to retain to allow rollback (If not set, default Kubernetes value is set to 10)
   # revisionHistoryLimit: 1
-  # Amount of time (in seconds) before Kubernetes will send the SIGKILL signal if Traefik does not shut down
+  # -- Amount of time (in seconds) before Kubernetes will send the SIGKILL signal if Traefik does not shut down
   terminationGracePeriodSeconds: 60
-  # The minimum number of seconds Traefik needs to be up and running before the DaemonSet/Deployment controller considers it available
+  # -- The minimum number of seconds Traefik needs to be up and running before the DaemonSet/Deployment controller considers it available
   minReadySeconds: 0
-  # Additional deployment annotations (e.g. for jaeger-operator sidecar injection)
+  # -- Additional deployment annotations (e.g. for jaeger-operator sidecar injection)
   annotations: {}
-  # Additional deployment labels (e.g. for filtering deployment by custom labels)
+  # -- Additional deployment labels (e.g. for filtering deployment by custom labels)
   labels: {}
-  # Additional pod annotations (e.g. for mesh injection or prometheus scraping)
+  # -- Additional pod annotations (e.g. for mesh injection or prometheus scraping)
   podAnnotations: {}
-  # Additional Pod labels (e.g. for filtering Pod by custom labels)
+  # -- Additional Pod labels (e.g. for filtering Pod by custom labels)
   podLabels: {}
-  # Additional containers (e.g. for metric offloading sidecars)
+  # -- Additional containers (e.g. for metric offloading sidecars)
   additionalContainers: []
     # https://docs.datadoghq.com/developers/dogstatsd/unix_socket/?tab=host
     # - name: socat-proxy
-    # image: alpine/socat:1.0.5
-    # args: ["-s", "-u", "udp-recv:8125", "unix-sendto:/socket/socket"]
-    # volumeMounts:
-    #   - name: dsdsocket
-    #     mountPath: /socket
-  # Additional volumes available for use with initContainers and additionalContainers
+    #   image: alpine/socat:1.0.5
+    #   args: ["-s", "-u", "udp-recv:8125", "unix-sendto:/socket/socket"]
+    #   volumeMounts:
+    #     - name: dsdsocket
+    #       mountPath: /socket
+  # -- Additional volumes available for use with initContainers and additionalContainers
   additionalVolumes: []
     # - name: dsdsocket
     #   hostPath:
     #     path: /var/run/statsd-exporter
-  # Additional initContainers (e.g. for setting file permission as shown below)
+  # -- Additional initContainers (e.g. for setting file permission as shown below)
   initContainers: []
     # The "volume-permissions" init container is required if you run into permission issues.
     # Related issue: https://github.com/traefik/traefik-helm-chart/issues/396
@@ -78,9 +64,9 @@ deployment:
     #   volumeMounts:
     #     - name: data
     #       mountPath: /data
-  # Use process namespace sharing
+  # -- Use process namespace sharing
   shareProcessNamespace: false
-  # Custom pod DNS policy. Apply if `hostNetwork: true`
+  # -- Custom pod DNS policy. Apply if `hostNetwork: true`
   # dnsPolicy: ClusterFirstWithHostNet
   dnsConfig: {}
     # nameservers:
@@ -92,10 +78,10 @@ deployment:
     #   - name: ndots
     #     value: "2"
     #   - name: edns0
-  # Additional imagePullSecrets
+  # -- Additional imagePullSecrets
   imagePullSecrets: []
     # - name: myRegistryKeySecretName
-  # Pod lifecycle actions
+  # -- Pod lifecycle actions
   lifecycle: {}
     # preStop:
     #   exec:
@@ -107,7 +93,7 @@ deployment:
     #     host: localhost
     #     scheme: HTTP

-# Pod disruption budget
+# -- Pod disruption budget
 podDisruptionBudget:
   enabled: false
   # maxUnavailable: 1
@@ -115,93 +101,112 @@ podDisruptionBudget:
   # minAvailable: 0
   # minAvailable: 25%

-# Create a default IngressClass for Traefik
+# -- Create a default IngressClass for Traefik
 ingressClass:
   enabled: true
   isDefaultClass: true

-# Enable experimental features
+# Traefik experimental features
 experimental:
   v3:
+    # -- Enable traefik version 3
     enabled: false
   plugins:
+    # -- Enable traefik experimental plugins
     enabled: false
   kubernetesGateway:
+    # -- Enable traefik experimental GatewayClass CRD
     enabled: false
     gateway:
+      # -- Enable traefik regular kubernetes gateway
       enabled: true
     # certificate:
     #   group: "core"
     #   kind: "Secret"
     #   name: "mysecret"
-    # By default, Gateway would be created to the Namespace you are deploying Traefik to.
+    # -- By default, Gateway would be created to the Namespace you are deploying Traefik to.
     # You may create that Gateway in another namespace, setting its name below:
     # namespace: default
     # Additional gateway annotations (e.g. for cert-manager.io/issuer)
     # annotations:
     #   cert-manager.io/issuer: letsencrypt

-# Create an IngressRoute for the dashboard
+## Create an IngressRoute for the dashboard
 ingressRoute:
   dashboard:
+    # -- Create an IngressRoute for the dashboard
     enabled: true
-    # Additional ingressRoute annotations (e.g. for kubernetes.io/ingress.class)
+    # -- Additional ingressRoute annotations (e.g. for kubernetes.io/ingress.class)
     annotations: {}
-    # Additional ingressRoute labels (e.g. for filtering IngressRoute by custom labels)
+    # -- Additional ingressRoute labels (e.g. for filtering IngressRoute by custom labels)
     labels: {}
-    # The router match rule used for the dashboard ingressRoute
+    # -- The router match rule used for the dashboard ingressRoute
     matchRule: PathPrefix(`/dashboard`) || PathPrefix(`/api`)
-    # Specify the allowed entrypoints to use for the dashboard ingress route, (e.g. traefik, web, websecure).
+    # -- Specify the allowed entrypoints to use for the dashboard ingress route, (e.g. traefik, web, websecure).
     # By default, it's using traefik entrypoint, which is not exposed.
     # /!\ Do not expose your dashboard without any protection over the internet /!\
     entryPoints: ["traefik"]
-    # Additional ingressRoute middlewares (e.g. for authentication)
+    # -- Additional ingressRoute middlewares (e.g. for authentication)
     middlewares: []
-    # TLS options (e.g. secret containing certificate)
+    # -- TLS options (e.g. secret containing certificate)
     tls: {}

-# Customize updateStrategy of traefik pods
 updateStrategy:
+  # -- Customize updateStrategy: RollingUpdate or OnDelete
   type: RollingUpdate
   rollingUpdate:
     maxUnavailable: 0
     maxSurge: 1

-# Customize liveness and readiness probe values.
 readinessProbe:
+  # -- The number of consecutive failures allowed before considering the probe as failed.
   failureThreshold: 1
+  # -- The number of seconds to wait before starting the first probe.
   initialDelaySeconds: 2
+  # -- The number of seconds to wait between consecutive probes.
   periodSeconds: 10
+  # -- The minimum consecutive successes required to consider the probe successful.
   successThreshold: 1
+  # -- The number of seconds to wait for a probe response before considering it as failed.
   timeoutSeconds: 2
-
 livenessProbe:
+  # -- The number of consecutive failures allowed before considering the probe as failed.
   failureThreshold: 3
+  # -- The number of seconds to wait before starting the first probe.
   initialDelaySeconds: 2
+  # -- The number of seconds to wait between consecutive probes.
   periodSeconds: 10
+  # -- The minimum consecutive successes required to consider the probe successful.
   successThreshold: 1
+  # -- The number of seconds to wait for a probe response before considering it as failed.
   timeoutSeconds: 2

-#
-# Configure providers
-#
 providers:
   kubernetesCRD:
+    # -- Load Kubernetes IngressRoute provider
     enabled: true
+    # -- Allows IngressRoute to reference resources in namespace other than theirs
     allowCrossNamespace: false
+    # -- Allows to reference ExternalName services in IngressRoute
     allowExternalNameServices: false
+    # -- Allows to return 503 when there is no endpoints available
     allowEmptyServices: false
     # ingressClass: traefik-internal
     # labelSelector: environment=production,method=traefik
+    # -- Array of namespaces to watch. If left empty, Traefik watches all namespaces.
     namespaces: []
       # - "default"

   kubernetesIngress:
+    # -- Load Kubernetes IngressRoute provider
     enabled: true
+    # -- Allows to reference ExternalName services in Ingress
     allowExternalNameServices: false
+    # -- Allows to return 503 when there is no endpoints available
     allowEmptyServices: false
     # ingressClass: traefik-internal
     # labelSelector: environment=production,method=traefik
+    # -- Array of namespaces to watch. If left empty, Traefik watches all namespaces.
     namespaces: []
       # - "default"
     # IP used for Kubernetes Ingress endpoints
@@ -212,13 +217,13 @@ providers:
       # pathOverride: ""

 #
-# Add volumes to the traefik pod. The volume name will be passed to tpl.
+# -- Add volumes to the traefik pod. The volume name will be passed to tpl.
 # This can be used to mount a cert pair or a configmap that holds a config.toml file.
 # After the volume has been mounted, add the configs into traefik by using the `additionalArguments` list below, eg:
-# additionalArguments:
+# `additionalArguments:
 # - "--providers.file.filename=/config/dynamic.toml"
 # - "--ping"
-# - "--ping.entrypoint=web"
+# - "--ping.entrypoint=web"`
 volumes: []
 # - name: public-cert
 #   mountPath: "/certs"
@@ -227,25 +232,22 @@ volumes: []
 #   mountPath: "/config"
 #   type: configMap

-# Additional volumeMounts to add to the Traefik container
+# -- Additional volumeMounts to add to the Traefik container
 additionalVolumeMounts: []
-  # For instance when using a logshipper for access logs
+  # -- For instance when using a logshipper for access logs
   # - name: traefik-logs
   #   mountPath: /var/log/traefik

-## Logs
-## https://docs.traefik.io/observability/logs/
 logs:
-  ## Traefik logs concern everything that happens to Traefik itself (startup, configuration, events, shutdown, and so on).
   general:
-    # By default, the logs use a text format (common), but you can
+    # -- By default, the logs use a text format (common), but you can
     # also ask for the json format in the format option
     # format: json
     # By default, the level is set to ERROR.
-    # Alternative logging levels are DEBUG, PANIC, FATAL, ERROR, WARN, and INFO.
+    # -- Alternative logging levels are DEBUG, PANIC, FATAL, ERROR, WARN, and INFO.
     level: ERROR
   access:
-    # To enable access logs
+    # -- To enable access logs
     enabled: false
     ## By default, logs are written using the Common Log Format (CLF) on stdout.
     ## To write logs in JSON, use json in the format option.
@@ -256,21 +258,24 @@ logs:
     ## This option represents the number of log lines Traefik will keep in memory before writing
     ## them to the selected output. In some cases, this option can greatly help performances.
     # bufferingSize: 100
-    ## Filtering https://docs.traefik.io/observability/access-logs/#filtering
+    ## Filtering
+    # -- https://docs.traefik.io/observability/access-logs/#filtering
     filters: {}
       # statuscodes: "200,300-302"
       # retryattempts: true
       # minduration: 10ms
-    ## Fields
-    ## https://docs.traefik.io/observability/access-logs/#limiting-the-fieldsincluding-headers
     fields:
       general:
+        # -- Available modes: keep, drop, redact.
         defaultmode: keep
+        # -- Names of the fields to limit.
         names: {}
           ## Examples:
           # ClientUsername: drop
       headers:
+        # -- Available modes: keep, drop, redact.
         defaultmode: drop
+        # -- Names of the headers to limit.
         names: {}
           ## Examples:
           # User-Agent: redact
@@ -278,10 +283,10 @@ logs:
           # Content-Type: keep

 metrics:
-  ## Prometheus is enabled by default.
-  ## It can be disabled by setting "prometheus: null"
+  ## -- Prometheus is enabled by default.
+  ## -- It can be disabled by setting "prometheus: null"
   prometheus:
-    ## Entry point used to expose metrics.
+    # -- Entry point used to expose metrics.
     entryPoint: metrics
     ## Enable metrics on entry points. Default=true
     # addEntryPointsLabels: false
@@ -404,11 +409,9 @@ metrics:
 #    ## This instructs the reporter to send metrics to the OpenTelemetry Collector using gRPC.
 #    grpc: true

-##
-##  enable optional CRDs for Prometheus Operator
+## -- enable optional CRDs for Prometheus Operator
 ##
   ## Create a dedicated metrics service for use with ServiceMonitor
-  ## When hub.enabled is set to true, it's not needed: it will use hub service.
   #  service:
   #    enabled: false
   #    labels: {}
@@ -455,6 +458,8 @@ metrics:
   #          summary: "Traefik Down"
   #          description: "{{ $labels.pod }} on {{ $labels.nodename }} is down"

+## Tracing
+# -- https://doc.traefik.io/traefik/observability/tracing/overview/
 tracing: {}
   # instana:
   #   localAgentHost: 127.0.0.1
@@ -497,20 +502,21 @@ tracing: {}
   #   secretToken: ""
   #   serviceEnvironment: ""

+# -- Global command arguments to be passed to all traefik's pods
 globalArguments:
   - "--global.checknewversion"
   - "--global.sendanonymoususage"

 #
 # Configure Traefik static configuration
-# Additional arguments to be passed at Traefik's binary
+# -- Additional arguments to be passed at Traefik's binary
 # All available options available on https://docs.traefik.io/reference/static-configuration/cli/
 ## Use curly braces to pass values: `helm install --set="additionalArguments={--providers.kubernetesingress.ingressclass=traefik-internal,--log.level=DEBUG}"`
 additionalArguments: []
 #  - "--providers.kubernetesingress.ingressclass=traefik-internal"
 #  - "--log.level=DEBUG"

-# Environment variables to be passed to Traefik's binary
+# -- Environment variables to be passed to Traefik's binary
 env: []
 # - name: SOME_VAR
 #   value: some-var-value
@@ -525,22 +531,20 @@ env: []
 #       name: secret-name
 #       key: secret-key

+# -- Environment variables to be passed to Traefik's binary from configMaps or secrets
 envFrom: []
 # - configMapRef:
 #     name: config-map-name
 # - secretRef:
 #     name: secret-name

-# Configure ports
 ports:
-  # The name of this one can't be changed as it is used for the readiness and
-  # liveness probes, but you can adjust its config to your liking
   traefik:
     port: 9000
-    # Use hostPort if set.
+    # -- Use hostPort if set.
     # hostPort: 9000
     #
-    # Use hostIP if set. If not set, Kubernetes will default to 0.0.0.0, which
+    # -- Use hostIP if set. If not set, Kubernetes will default to 0.0.0.0, which
     # means it's listening on all your interfaces and all your IPs. You may want
     # to set this value if you need traefik to listen on specific interface
     # only.
@@ -558,27 +562,27 @@ ports:
     # Defines whether the port is exposed if service.type is LoadBalancer or
     # NodePort.
     #
-    # You SHOULD NOT expose the traefik port on production deployments.
+    # -- You SHOULD NOT expose the traefik port on production deployments.
     # If you want to access it from outside of your cluster,
     # use `kubectl port-forward` or create a secure ingress
     expose: false
-    # The exposed port for this service
+    # -- The exposed port for this service
     exposedPort: 9000
-    # The port protocol (TCP/UDP)
+    # -- The port protocol (TCP/UDP)
     protocol: TCP
   web:
-    ## Enable this entrypoint as a default entrypoint. When a service doesn't explicity set an entrypoint it will only use this entrypoint.
+    ## -- Enable this entrypoint as a default entrypoint. When a service doesn't explicity set an entrypoint it will only use this entrypoint.
     # asDefault: true
     port: 8000
     # hostPort: 8000
     # containerPort: 8000
     expose: true
     exposedPort: 80
-    ## Different target traefik port on the cluster, useful for IP type LB
+    ## -- Different target traefik port on the cluster, useful for IP type LB
     # targetPort: 80
     # The port protocol (TCP/UDP)
     protocol: TCP
-    # Use nodeport if set. This is useful if you have configured Traefik in a
+    # -- Use nodeport if set. This is useful if you have configured Traefik in a
     # LoadBalancer.
     # nodePort: 32080
     # Port Redirections
@@ -596,20 +600,22 @@ ports:
     #   trustedIPs: []
     #   insecure: false
   websecure:
-    ## Enable this entrypoint as a default entrypoint. When a service doesn't explicity set an entrypoint it will only use this entrypoint.
+    ## -- Enable this entrypoint as a default entrypoint. When a service doesn't explicity set an entrypoint it will only use this entrypoint.
     # asDefault: true
     port: 8443
     # hostPort: 8443
     # containerPort: 8443
     expose: true
     exposedPort: 443
-    ## Different target traefik port on the cluster, useful for IP type LB
+    ## -- Different target traefik port on the cluster, useful for IP type LB
     # targetPort: 80
-    ## The port protocol (TCP/UDP)
+    ## -- The port protocol (TCP/UDP)
     protocol: TCP
     # nodePort: 32443
+    ## -- Specify an application protocol. This may be used as a hint for a Layer 7 load balancer.
+    # appProtocol: https
     #
-    ## Enable HTTP/3 on the entrypoint
+    ## -- Enable HTTP/3 on the entrypoint
     ## Enabling it will also enable http3 experimental feature
     ## https://doc.traefik.io/traefik/routing/entrypoints/#http3
     ## There are known limitations when trying to listen on same ports for
@@ -619,12 +625,12 @@ ports:
       enabled: false
     # advertisedPort: 4443
     #
-    ## Trust forwarded  headers information (X-Forwarded-*).
+    ## -- Trust forwarded  headers information (X-Forwarded-*).
     #forwardedHeaders:
     #  trustedIPs: []
     #  insecure: false
     #
-    ## Enable the Proxy Protocol header parsing for the entry point
+    ## -- Enable the Proxy Protocol header parsing for the entry point
     #proxyProtocol:
     #  trustedIPs: []
     #  insecure: false
@@ -642,33 +648,33 @@ ports:
       #     - foo.example.com
       #     - bar.example.com
     #
-    # One can apply Middlewares on an entrypoint
+    # -- One can apply Middlewares on an entrypoint
     # https://doc.traefik.io/traefik/middlewares/overview/
     # https://doc.traefik.io/traefik/routing/entrypoints/#middlewares
-    # /!\ It introduces here a link between your static configuration and your dynamic configuration /!\
+    # -- /!\ It introduces here a link between your static configuration and your dynamic configuration /!\
     # It follows the provider naming convention: https://doc.traefik.io/traefik/providers/overview/#provider-namespace
     # middlewares:
     #   - namespace-name1@kubernetescrd
     #   - namespace-name2@kubernetescrd
     middlewares: []
   metrics:
-    # When using hostNetwork, use another port to avoid conflict with node exporter:
+    # -- When using hostNetwork, use another port to avoid conflict with node exporter:
     # https://github.com/prometheus/prometheus/wiki/Default-port-allocations
     port: 9100
     # hostPort: 9100
     # Defines whether the port is exposed if service.type is LoadBalancer or
     # NodePort.
     #
-    # You may not want to expose the metrics port on production deployments.
+    # -- You may not want to expose the metrics port on production deployments.
     # If you want to access it from outside of your cluster,
     # use `kubectl port-forward` or create a secure ingress
     expose: false
-    # The exposed port for this service
+    # -- The exposed port for this service
     exposedPort: 9100
-    # The port protocol (TCP/UDP)
+    # -- The port protocol (TCP/UDP)
     protocol: TCP

-# TLS Options are created as TLSOption CRDs
+# -- TLS Options are created as TLSOption CRDs
 # https://doc.traefik.io/traefik/https/tls/#tls-options
 # When using `labelSelector`, you'll need to set labels on tlsOption accordingly.
 # Example:
@@ -684,7 +690,7 @@ ports:
 #       - CurveP384
 tlsOptions: {}

-# TLS Store are created as TLSStore CRDs. This is useful if you want to set a default certificate
+# -- TLS Store are created as TLSStore CRDs. This is useful if you want to set a default certificate
 # https://doc.traefik.io/traefik/https/tls/#default-certificate
 # Example:
 # tlsStore:
@@ -693,24 +699,22 @@ tlsOptions: {}
 #       secretName: tls-cert
 tlsStore: {}

-# Options for the main traefik service, where the entrypoints traffic comes
-# from.
 service:
   enabled: true
-  ## Single service is using `MixedProtocolLBService` feature gate.
-  ## When set to false, it will create two Service, one for TCP and one for UDP.
+  ## -- Single service is using `MixedProtocolLBService` feature gate.
+  ## -- When set to false, it will create two Service, one for TCP and one for UDP.
   single: true
   type: LoadBalancer
-  # Additional annotations applied to both TCP and UDP services (e.g. for cloud provider specific config)
+  # -- Additional annotations applied to both TCP and UDP services (e.g. for cloud provider specific config)
   annotations: {}
-  # Additional annotations for TCP service only
+  # -- Additional annotations for TCP service only
   annotationsTCP: {}
-  # Additional annotations for UDP service only
+  # -- Additional annotations for UDP service only
   annotationsUDP: {}
-  # Additional service labels (e.g. for filtering Service by custom labels)
+  # -- Additional service labels (e.g. for filtering Service by custom labels)
   labels: {}
-  # Additional entries here will be added to the service spec.
-  # Cannot contain type, selector or ports entries.
+  # -- Additional entries here will be added to the service spec.
+  # -- Cannot contain type, selector or ports entries.
   spec: {}
     # externalTrafficPolicy: Cluster
     # loadBalancerIP: "1.2.3.4"
@@ -718,6 +722,8 @@ service:
   loadBalancerSourceRanges: []
     # - 192.168.0.1/32
     # - 172.16.0.0/16
+  ## -- Class of the load balancer implementation
+  # loadBalancerClass: service.k8s.aws/nlb
   externalIPs: []
     # - 1.2.3.4
   ## One of SingleStack, PreferDualStack, or RequireDualStack.
@@ -728,7 +734,7 @@ service:
   #   - IPv4
   #   - IPv6
   ##
-  ## An additionnal and optional internal Service.
+  ## -- An additionnal and optional internal Service.
   ## Same parameters as external Service
   # internal:
   #   type: ClusterIP
@@ -739,9 +745,8 @@ service:
   #   # externalIPs: []
   #   # ipFamilies: [ "IPv4","IPv6" ]

-## Create HorizontalPodAutoscaler object.
-##
 autoscaling:
+  # -- Create HorizontalPodAutoscaler object.
   enabled: false
 #   minReplicas: 1
 #   maxReplicas: 10
@@ -766,10 +771,10 @@ autoscaling:
 #         value: 1
 #         periodSeconds: 60

-# Enable persistence using Persistent Volume Claims
-# ref: http://kubernetes.io/docs/user-guide/persistent-volumes/
-# It can be used to store TLS certificates, see `storage` in certResolvers
 persistence:
+  # -- Enable persistence using Persistent Volume Claims
+  # ref: http://kubernetes.io/docs/user-guide/persistent-volumes/
+  # It can be used to store TLS certificates, see `storage` in certResolvers
   enabled: false
   name: data
 #  existingClaim: ""
@@ -779,8 +784,10 @@ persistence:
   # volumeName: ""
   path: /data
   annotations: {}
-  # subPath: "" # only mount a subpath of the Volume into the pod
+  # -- Only mount a subpath of the Volume into the pod
+  # subPath: ""

+# -- Certificates resolvers configuration
 certResolvers: {}
 #   letsencrypt:
 #     # for challenge options cf. https://doc.traefik.io/traefik/https/acme/
@@ -802,13 +809,13 @@ certResolvers: {}
 #     # It has to match the path with a persistent volume
 #     storage: /data/acme.json

-# If hostNetwork is true, runs traefik in the host network namespace
+# -- If hostNetwork is true, runs traefik in the host network namespace
 # To prevent unschedulabel pods due to port collisions, if hostNetwork=true
 # and replicas>1, a pod anti-affinity is recommended and will be set if the
 # affinity is left as default.
 hostNetwork: false

-# Whether Role Based Access Control objects like roles and rolebindings should be created
+# -- Whether Role Based Access Control objects like roles and rolebindings should be created
 rbac:
   enabled: true
   # If set to false, installs ClusterRole and ClusterRoleBinding so Traefik can be used across namespaces.
@@ -818,19 +825,20 @@ rbac:
   # https://kubernetes.io/docs/reference/access-authn-authz/rbac/#user-facing-roles
   # aggregateTo: [ "admin" ]

-# Enable to create a PodSecurityPolicy and assign it to the Service Account via RoleBinding or ClusterRoleBinding
+# -- Enable to create a PodSecurityPolicy and assign it to the Service Account via RoleBinding or ClusterRoleBinding
 podSecurityPolicy:
   enabled: false

-# The service account the pods will use to interact with the Kubernetes API
+# -- The service account the pods will use to interact with the Kubernetes API
 serviceAccount:
   # If set, an existing service account is used
   # If not set, a service account is created automatically using the fullname template
   name: ""

-# Additional serviceAccount annotations (e.g. for oidc authentication)
+# -- Additional serviceAccount annotations (e.g. for oidc authentication)
 serviceAccountAnnotations: {}

+# -- The resources parameter defines CPU and memory requirements and limits for Traefik's containers.
 resources: {}
   # requests:
   #   cpu: "100m"
@@ -839,8 +847,8 @@ resources: {}
   #   cpu: "300m"
   #   memory: "150Mi"

-# This example pod anti-affinity forces the scheduler to put traefik pods
-# on nodes where no other traefik pods are scheduled.
+# -- This example pod anti-affinity forces the scheduler to put traefik pods
+# -- on nodes where no other traefik pods are scheduled.
 # It should be used when hostNetwork: true to prevent port conflicts
 affinity: {}
 #  podAntiAffinity:
@@ -851,11 +859,15 @@ affinity: {}
 #            app.kubernetes.io/instance: '{{ .Release.Name }}-{{ .Release.Namespace }}'
 #        topologyKey: kubernetes.io/hostname

+# -- nodeSelector is the simplest recommended form of node selection constraint.
 nodeSelector: {}
+# -- Tolerations allow the scheduler to schedule pods with matching taints.
 tolerations: []
+# -- You can use topology spread constraints to control
+# how Pods are spread across your cluster among failure-domains.
 topologySpreadConstraints: []
-# # This example topologySpreadConstraints forces the scheduler to put traefik pods
-# # on nodes where no other traefik pods are scheduled.
+# This example topologySpreadConstraints forces the scheduler to put traefik pods
+# on nodes where no other traefik pods are scheduled.
 #  - labelSelector:
 #      matchLabels:
 #        app: '{{ template "traefik.name" . }}'
@@ -863,29 +875,33 @@ topologySpreadConstraints: []
 #    topologyKey: kubernetes.io/hostname
 #    whenUnsatisfiable: DoNotSchedule

-# Pods can have priority.
-# Priority indicates the importance of a Pod relative to other Pods.
+# -- Pods can have priority.
+# -- Priority indicates the importance of a Pod relative to other Pods.
 priorityClassName: ""

-# Set the container security context
-# To run the container with ports below 1024 this will need to be adjust to run as root
+# -- Set the container security context
+# -- To run the container with ports below 1024 this will need to be adjust to run as root
 securityContext:
   capabilities:
     drop: [ALL]
   readOnlyRootFilesystem: true

 podSecurityContext:
-#  # /!\ When setting fsGroup, Kubernetes will recursively changes ownership and
-#  # permissions for the contents of each volume to match the fsGroup. This can
-#  # be an issue when storing sensitive content like TLS Certificates /!\
-#  fsGroup: 65532
+  # /!\ When setting fsGroup, Kubernetes will recursively changes ownership and
+  # permissions for the contents of each volume to match the fsGroup. This can
+  # be an issue when storing sensitive content like TLS Certificates /!\
+  # fsGroup: 65532
+  # -- Specifies the policy for changing ownership and permissions of volume contents to match the fsGroup.
   fsGroupChangePolicy: "OnRootMismatch"
+  # -- The ID of the group for all containers in the pod to run as.
   runAsGroup: 65532
+  # -- Specifies whether the containers should run as a non-root user.
   runAsNonRoot: true
+  # -- The ID of the user for all containers in the pod to run as.
   runAsUser: 65532

 #
-# Extra objects to deploy (value evaluated as a template)
+# -- Extra objects to deploy (value evaluated as a template)
 #
 # In some cases, it can avoid the need for additional, extended or adhoc deployments.
 # See #595 for more details and traefik/tests/values/extra.yaml for example.
@@ -895,5 +911,5 @@ extraObjects: []
 # It will not affect optional CRDs such as `ServiceMonitor` and `PrometheusRules`
 # namespaceOverride: traefik
 #
-## This will override the default app.kubernetes.io/instance label for all Objects.
+## -- This will override the default app.kubernetes.io/instance label for all Objects.
 # instanceLabelOverride: traefik
```

## 23.0.1  ![AppVersion: v2.10.1](https://img.shields.io/static/v1?label=AppVersion&message=v2.10.1&color=success&logo=) ![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.16.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2023-04-28

* fix: ⬆️ Upgrade traefik Docker tag to v2.10.1


## 23.0.0  ![AppVersion: v2.10.0](https://img.shields.io/static/v1?label=AppVersion&message=v2.10.0&color=success&logo=) ![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.16.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2023-04-26

* BREAKING CHANGE: Traefik 2.10 comes with CRDs update on API Group


## 22.3.0  ![AppVersion: v2.10.0](https://img.shields.io/static/v1?label=AppVersion&message=v2.10.0&color=success&logo=) ![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.16.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2023-04-25

* ⬆️ Upgrade traefik Docker tag to v2.10.0
* fix: 🐛 update rbac for both traefik.io and containo.us apigroups (#836)
* breaking: 💥 update CRDs needed for Traefik v2.10


## 22.2.0  ![AppVersion: v2.9.10](https://img.shields.io/static/v1?label=AppVersion&message=v2.9.10&color=success&logo=) ![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.16.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2023-04-24

* test: 👷 Update unit tests tooling
* fix: 🐛 annotations leaking between aliased subcharts
* fix: indentation on `TLSOption`
* feat: override container port
* feat: allow to set dnsConfig on pod template
* chore: 🔧 new release
* added targetPort support

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 9ece303..71273cc 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -82,6 +82,16 @@ deployment:
   shareProcessNamespace: false
   # Custom pod DNS policy. Apply if `hostNetwork: true`
   # dnsPolicy: ClusterFirstWithHostNet
+  dnsConfig: {}
+    # nameservers:
+    #   - 192.0.2.1 # this is an example
+    # searches:
+    #   - ns1.svc.cluster-domain.example
+    #   - my.dns.search.suffix
+    # options:
+    #   - name: ndots
+    #     value: "2"
+    #   - name: edns0
   # Additional imagePullSecrets
   imagePullSecrets: []
     # - name: myRegistryKeySecretName
@@ -561,8 +571,11 @@ ports:
     # asDefault: true
     port: 8000
     # hostPort: 8000
+    # containerPort: 8000
     expose: true
     exposedPort: 80
+    ## Different target traefik port on the cluster, useful for IP type LB
+    # targetPort: 80
     # The port protocol (TCP/UDP)
     protocol: TCP
     # Use nodeport if set. This is useful if you have configured Traefik in a
@@ -587,8 +600,11 @@ ports:
     # asDefault: true
     port: 8443
     # hostPort: 8443
+    # containerPort: 8443
     expose: true
     exposedPort: 443
+    ## Different target traefik port on the cluster, useful for IP type LB
+    # targetPort: 80
     ## The port protocol (TCP/UDP)
     protocol: TCP
     # nodePort: 32443
```

## 22.1.0  ![AppVersion: v2.9.10](https://img.shields.io/static/v1?label=AppVersion&message=v2.9.10&color=success&logo=) ![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.16.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2023-04-07

* ⬆️ Upgrade traefik Docker tag to v2.9.10
* feat: add additional labels to tlsoption

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 4762b77..9ece303 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -654,12 +654,15 @@ ports:

 # TLS Options are created as TLSOption CRDs
 # https://doc.traefik.io/traefik/https/tls/#tls-options
+# When using `labelSelector`, you'll need to set labels on tlsOption accordingly.
 # Example:
 # tlsOptions:
 #   default:
+#     labels: {}
 #     sniStrict: true
 #     preferServerCipherSuites: true
-#   foobar:
+#   customOptions:
+#     labels: {}
 #     curvePreferences:
 #       - CurveP521
 #       - CurveP384
```

## 22.0.0  ![AppVersion: v2.9.9](https://img.shields.io/static/v1?label=AppVersion&message=v2.9.9&color=success&logo=) ![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.16.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2023-03-29

* BREAKING CHANGE: `image.repository` introduction may break during the upgrade. See PR #802.


## 21.2.1  ![AppVersion: v2.9.9](https://img.shields.io/static/v1?label=AppVersion&message=v2.9.9&color=success&logo=) ![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.16.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2023-03-28

* 🎨 Introduce `image.registry` and add explicit default (it may impact custom `image.repository`)
* ⬆️ Upgrade traefik Docker tag to v2.9.9
* :memo: Clarify the need of an initContainer when enabling persistence for TLS Certificates

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index cadc7a6..4762b77 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -1,5 +1,6 @@
 # Default values for Traefik
 image:
+  registry: docker.io
   repository: traefik
   # defaults to appVersion
   tag: ""
@@ -66,10 +67,14 @@ deployment:
   # Additional initContainers (e.g. for setting file permission as shown below)
   initContainers: []
     # The "volume-permissions" init container is required if you run into permission issues.
-    # Related issue: https://github.com/traefik/traefik/issues/6825
+    # Related issue: https://github.com/traefik/traefik-helm-chart/issues/396
     # - name: volume-permissions
-    #   image: busybox:1.35
-    #   command: ["sh", "-c", "touch /data/acme.json && chmod -Rv 600 /data/* && chown 65532:65532 /data/acme.json"]
+    #   image: busybox:latest
+    #   command: ["sh", "-c", "touch /data/acme.json; chmod -v 600 /data/acme.json"]
+    #   securityContext:
+    #     runAsNonRoot: true
+    #     runAsGroup: 65532
+    #     runAsUser: 65532
     #   volumeMounts:
     #     - name: data
     #       mountPath: /data
@@ -849,13 +854,17 @@ securityContext:
   capabilities:
     drop: [ALL]
   readOnlyRootFilesystem: true
+
+podSecurityContext:
+#  # /!\ When setting fsGroup, Kubernetes will recursively changes ownership and
+#  # permissions for the contents of each volume to match the fsGroup. This can
+#  # be an issue when storing sensitive content like TLS Certificates /!\
+#  fsGroup: 65532
+  fsGroupChangePolicy: "OnRootMismatch"
   runAsGroup: 65532
   runAsNonRoot: true
   runAsUser: 65532

-podSecurityContext:
-  fsGroup: 65532
-
 #
 # Extra objects to deploy (value evaluated as a template)
 #
```

## 21.2.0  ![AppVersion: v2.9.8](https://img.shields.io/static/v1?label=AppVersion&message=v2.9.8&color=success&logo=) ![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.16.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2023-03-08

* 🚨 Fail when enabling PSP on Kubernetes v1.25+ (#801)
* ⬆️ Upgrade traefik Docker tag to v2.9.8
* Separate UDP hostPort for HTTP/3
* :sparkles: release 21.2.0 (#805)


## 21.1.0  ![AppVersion: v2.9.7](https://img.shields.io/static/v1?label=AppVersion&message=v2.9.7&color=success&logo=) ![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.16.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2023-02-15

* ⬆️ Upgrade traefik Docker tag to v2.9.7
* ✨ release 21.1.0
* fix: traefik image name for renovate
* feat: Add volumeName to PersistentVolumeClaim (#792)
* Allow setting TLS options on dashboard IngressRoute

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 780b04b..cadc7a6 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -142,6 +142,8 @@ ingressRoute:
     entryPoints: ["traefik"]
     # Additional ingressRoute middlewares (e.g. for authentication)
     middlewares: []
+    # TLS options (e.g. secret containing certificate)
+    tls: {}

 # Customize updateStrategy of traefik pods
 updateStrategy:
@@ -750,6 +752,7 @@ persistence:
   accessMode: ReadWriteOnce
   size: 128Mi
   # storageClass: ""
+  # volumeName: ""
   path: /data
   annotations: {}
   # subPath: "" # only mount a subpath of the Volume into the pod
```

## 21.0.0  ![AppVersion: v2.9.6](https://img.shields.io/static/v1?label=AppVersion&message=v2.9.6&color=success&logo=) ![Kubernetes: >=1.16.0-0](https://img.shields.io/static/v1?label=Kubernetes&message=%3E%3D1.16.0-0&color=informational&logo=kubernetes) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2023-02-10

* 🙈 Add a setting disable API check on Prometheus Operator (#769)
* 📝 Improve documentation on entrypoint options
* 💥 New release with BREAKING changes (#786)
* ✨ Chart.yaml - add kubeVersion: ">=1.16.0-0"
* fix: allowExternalNameServices for kubernetes ingress when hub enabled (#772)
* fix(service-metrics): invert prometheus svc & fullname length checking
* Configure Renovate (#783)
* :necktie: Improve labels settings behavior on metrics providers (#774)
* :bug: Disabling dashboard ingressroute should delete it (#785)
* :boom: Rename image.name => image.repository (#784)

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 42a27f9..780b04b 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -1,6 +1,6 @@
 # Default values for Traefik
 image:
-  name: traefik
+  repository: traefik
   # defaults to appVersion
   tag: ""
   pullPolicy: IfNotPresent
@@ -396,6 +396,8 @@ metrics:
   #    enabled: false
   #    labels: {}
   #    annotations: {}
+  ## When set to true, it won't check if Prometheus Operator CRDs are deployed
+  #  disableAPICheck: false
   #  serviceMonitor:
   #    metricRelabelings: []
   #      - sourceLabels: [__name__]
@@ -580,7 +582,7 @@ ports:
     # hostPort: 8443
     expose: true
     exposedPort: 443
-    # The port protocol (TCP/UDP)
+    ## The port protocol (TCP/UDP)
     protocol: TCP
     # nodePort: 32443
     #
@@ -594,6 +596,16 @@ ports:
       enabled: false
     # advertisedPort: 4443
     #
+    ## Trust forwarded  headers information (X-Forwarded-*).
+    #forwardedHeaders:
+    #  trustedIPs: []
+    #  insecure: false
+    #
+    ## Enable the Proxy Protocol header parsing for the entry point
+    #proxyProtocol:
+    #  trustedIPs: []
+    #  insecure: false
+    #
     ## Set TLS at the entrypoint
     ## https://doc.traefik.io/traefik/routing/entrypoints/#tls
     tls:
@@ -607,16 +619,6 @@ ports:
       #     - foo.example.com
       #     - bar.example.com
     #
-    # Trust forwarded  headers information (X-Forwarded-*).
-    # forwardedHeaders:
-    #   trustedIPs: []
-    #   insecure: false
-    #
-    # Enable the Proxy Protocol header parsing for the entry point
-    # proxyProtocol:
-    #   trustedIPs: []
-    #   insecure: false
-    #
     # One can apply Middlewares on an entrypoint
     # https://doc.traefik.io/traefik/middlewares/overview/
     # https://doc.traefik.io/traefik/routing/entrypoints/#middlewares
```

## 20.8.0  ![AppVersion: v2.9.6](https://img.shields.io/static/v1?label=AppVersion&message=v2.9.6&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-12-09

* ✨ update chart to version 20.8.0
* ✨ add support for default entrypoints
* ✨ add support for OpenTelemetry and Traefik v3

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index b77539d..42a27f9 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -107,6 +107,8 @@ ingressClass:

 # Enable experimental features
 experimental:
+  v3:
+    enabled: false
   plugins:
     enabled: false
   kubernetesGateway:
@@ -347,7 +349,43 @@ metrics:
 #    # addRoutersLabels: true
 #    ## Enable metrics on services. Default=true
 #    # addServicesLabels: false
-
+#  openTelemetry:
+#    ## Address of the OpenTelemetry Collector to send metrics to.
+#    address: "localhost:4318"
+#    ## Enable metrics on entry points.
+#    addEntryPointsLabels: true
+#    ## Enable metrics on routers.
+#    addRoutersLabels: true
+#    ## Enable metrics on services.
+#    addServicesLabels: true
+#    ## Explicit boundaries for Histogram data points.
+#    explicitBoundaries:
+#      - "0.1"
+#      - "0.3"
+#      - "1.2"
+#      - "5.0"
+#    ## Additional headers sent with metrics by the reporter to the OpenTelemetry Collector.
+#    headers:
+#      foo: bar
+#      test: test
+#    ## Allows reporter to send metrics to the OpenTelemetry Collector without using a secured protocol.
+#    insecure: true
+#    ## Interval at which metrics are sent to the OpenTelemetry Collector.
+#    pushInterval: 10s
+#    ## Allows to override the default URL path used for sending metrics. This option has no effect when using gRPC transport.
+#    path: /foo/v1/traces
+#    ## Defines the TLS configuration used by the reporter to send metrics to the OpenTelemetry Collector.
+#    tls:
+#      ## The path to the certificate authority, it defaults to the system bundle.
+#      ca: path/to/ca.crt
+#      ## The path to the public certificate. When using this option, setting the key option is required.
+#      cert: path/to/foo.cert
+#      ## The path to the private key. When using this option, setting the cert option is required.
+#      key: path/to/key.key
+#      ## If set to true, the TLS connection accepts any certificate presented by the server regardless of the hostnames it covers.
+#      insecureSkipVerify: true
+#    ## This instructs the reporter to send metrics to the OpenTelemetry Collector using gRPC.
+#    grpc: true

 ##
 ##  enable optional CRDs for Prometheus Operator
@@ -510,6 +548,8 @@ ports:
     # The port protocol (TCP/UDP)
     protocol: TCP
   web:
+    ## Enable this entrypoint as a default entrypoint. When a service doesn't explicity set an entrypoint it will only use this entrypoint.
+    # asDefault: true
     port: 8000
     # hostPort: 8000
     expose: true
@@ -534,6 +574,8 @@ ports:
     #   trustedIPs: []
     #   insecure: false
   websecure:
+    ## Enable this entrypoint as a default entrypoint. When a service doesn't explicity set an entrypoint it will only use this entrypoint.
+    # asDefault: true
     port: 8443
     # hostPort: 8443
     expose: true
```

## 20.7.0  ![AppVersion: v2.9.6](https://img.shields.io/static/v1?label=AppVersion&message=v2.9.6&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-12-08

* 🐛 Don't fail when prometheus is disabled (#756)
* ⬆️  Update default Traefik release to v2.9.6 (#758)
* ✨ support for Gateway annotations
* add keywords [networking], for artifacthub category quering
* :bug: Fix typo on bufferingSize for access logs (#753)
* :adhesive_bandage: Add quotes for artifacthub changelog parsing (#748)

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 4f2fb2a..b77539d 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -120,6 +120,9 @@ experimental:
     # By default, Gateway would be created to the Namespace you are deploying Traefik to.
     # You may create that Gateway in another namespace, setting its name below:
     # namespace: default
+    # Additional gateway annotations (e.g. for cert-manager.io/issuer)
+    # annotations:
+    #   cert-manager.io/issuer: letsencrypt

 # Create an IngressRoute for the dashboard
 ingressRoute:
@@ -219,7 +222,8 @@ logs:
     # By default, the logs use a text format (common), but you can
     # also ask for the json format in the format option
     # format: json
-    # By default, the level is set to ERROR. Alternative logging levels are DEBUG, PANIC, FATAL, ERROR, WARN, and INFO.
+    # By default, the level is set to ERROR.
+    # Alternative logging levels are DEBUG, PANIC, FATAL, ERROR, WARN, and INFO.
     level: ERROR
   access:
     # To enable access logs
```

## 20.6.0  ![AppVersion: v2.9.5](https://img.shields.io/static/v1?label=AppVersion&message=v2.9.5&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-11-30

* 🔍️ Add filePath support on access logs (#747)
* :memo: Improve documentation on using PVC with TLS certificates
* :bug: Add missing scheme in help on Traefik Hub integration (#746)

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 15f1682..4f2fb2a 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -211,10 +211,10 @@ additionalVolumeMounts: []
   # - name: traefik-logs
   #   mountPath: /var/log/traefik

-# Logs
-# https://docs.traefik.io/observability/logs/
+## Logs
+## https://docs.traefik.io/observability/logs/
 logs:
-  # Traefik logs concern everything that happens to Traefik itself (startup, configuration, events, shutdown, and so on).
+  ## Traefik logs concern everything that happens to Traefik itself (startup, configuration, events, shutdown, and so on).
   general:
     # By default, the logs use a text format (common), but you can
     # also ask for the json format in the format option
@@ -224,31 +224,32 @@ logs:
   access:
     # To enable access logs
     enabled: false
-    # By default, logs are written using the Common Log Format (CLF).
-    # To write logs in JSON, use json in the format option.
-    # If the given format is unsupported, the default (CLF) is used instead.
+    ## By default, logs are written using the Common Log Format (CLF) on stdout.
+    ## To write logs in JSON, use json in the format option.
+    ## If the given format is unsupported, the default (CLF) is used instead.
     # format: json
-    # To write the logs in an asynchronous fashion, specify a bufferingSize option.
-    # This option represents the number of log lines Traefik will keep in memory before writing
-    # them to the selected output. In some cases, this option can greatly help performances.
+    # filePath: "/var/log/traefik/access.log
+    ## To write the logs in an asynchronous fashion, specify a bufferingSize option.
+    ## This option represents the number of log lines Traefik will keep in memory before writing
+    ## them to the selected output. In some cases, this option can greatly help performances.
     # bufferingSize: 100
-    # Filtering https://docs.traefik.io/observability/access-logs/#filtering
+    ## Filtering https://docs.traefik.io/observability/access-logs/#filtering
     filters: {}
       # statuscodes: "200,300-302"
       # retryattempts: true
       # minduration: 10ms
-    # Fields
-    # https://docs.traefik.io/observability/access-logs/#limiting-the-fieldsincluding-headers
+    ## Fields
+    ## https://docs.traefik.io/observability/access-logs/#limiting-the-fieldsincluding-headers
     fields:
       general:
         defaultmode: keep
         names: {}
-          # Examples:
+          ## Examples:
           # ClientUsername: drop
       headers:
         defaultmode: drop
         names: {}
-          # Examples:
+          ## Examples:
           # User-Agent: redact
           # Authorization: drop
           # Content-Type: keep
@@ -693,10 +694,7 @@ autoscaling:

 # Enable persistence using Persistent Volume Claims
 # ref: http://kubernetes.io/docs/user-guide/persistent-volumes/
-# After the pvc has been mounted, add the configs into traefik by using the `additionalArguments` list below, eg:
-# additionalArguments:
-# - "--certificatesresolvers.le.acme.storage=/data/acme.json"
-# It will persist TLS certificates.
+# It can be used to store TLS certificates, see `storage` in certResolvers
 persistence:
   enabled: false
   name: data
@@ -726,7 +724,7 @@ certResolvers: {}
 #     tlsChallenge: true
 #     httpChallenge:
 #       entryPoint: "web"
-#     # match the path to persistence
+#     # It has to match the path with a persistent volume
 #     storage: /data/acme.json

 # If hostNetwork is true, runs traefik in the host network namespace
```

## 20.5.3  ![AppVersion: v2.9.5](https://img.shields.io/static/v1?label=AppVersion&message=v2.9.5&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-11-25

* 🐛 Fix template issue with obsolete helm version + add helm version requirement (#743)


## 20.5.2  ![AppVersion: v2.9.5](https://img.shields.io/static/v1?label=AppVersion&message=v2.9.5&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-11-24

* ⬆️Update Traefik to v2.9.5 (#740)


## 20.5.1  ![AppVersion: v2.9.4](https://img.shields.io/static/v1?label=AppVersion&message=v2.9.4&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-11-23

* 🐛 Fix namespaceSelector on ServiceMonitor (#737)


## 20.5.0  ![AppVersion: v2.9.4](https://img.shields.io/static/v1?label=AppVersion&message=v2.9.4&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-11-23

* 🚀 Add complete support on metrics options (#735)
* 🐛 make tests use fixed version

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index e49d02d..15f1682 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -12,7 +12,7 @@ hub:
   ## Enabling Hub will:
   # * enable Traefik Hub integration on Traefik
   # * add `traefikhub-tunl` endpoint
-  # * enable addRoutersLabels on prometheus metrics
+  # * enable Prometheus metrics with addRoutersLabels
   # * enable allowExternalNameServices on KubernetesIngress provider
   # * enable allowCrossNamespace on KubernetesCRD provider
   # * add an internal (ClusterIP) Service, dedicated for Traefik Hub
@@ -254,16 +254,96 @@ logs:
           # Content-Type: keep

 metrics:
-  # datadog:
-  #   address: 127.0.0.1:8125
-  # influxdb:
-  #   address: localhost:8089
-  #   protocol: udp
+  ## Prometheus is enabled by default.
+  ## It can be disabled by setting "prometheus: null"
   prometheus:
+    ## Entry point used to expose metrics.
     entryPoint: metrics
-  #  addRoutersLabels: true
-  #  statsd:
-  #    address: localhost:8125
+    ## Enable metrics on entry points. Default=true
+    # addEntryPointsLabels: false
+    ## Enable metrics on routers. Default=false
+    # addRoutersLabels: true
+    ## Enable metrics on services. Default=true
+    # addServicesLabels: false
+    ## Buckets for latency metrics. Default="0.1,0.3,1.2,5.0"
+    # buckets: "0.5,1.0,2.5"
+    ## When manualRouting is true, it disables the default internal router in
+    ## order to allow creating a custom router for prometheus@internal service.
+    # manualRouting: true
+#  datadog:
+#    ## Address instructs exporter to send metrics to datadog-agent at this address.
+#    address: "127.0.0.1:8125"
+#    ## The interval used by the exporter to push metrics to datadog-agent. Default=10s
+#    # pushInterval: 30s
+#    ## The prefix to use for metrics collection. Default="traefik"
+#    # prefix: traefik
+#    ## Enable metrics on entry points. Default=true
+#    # addEntryPointsLabels: false
+#    ## Enable metrics on routers. Default=false
+#    # addRoutersLabels: true
+#    ## Enable metrics on services. Default=true
+#    # addServicesLabels: false
+#  influxdb:
+#    ## Address instructs exporter to send metrics to influxdb at this address.
+#    address: localhost:8089
+#    ## InfluxDB's address protocol (udp or http). Default="udp"
+#    protocol: udp
+#    ## InfluxDB database used when protocol is http. Default=""
+#    # database: ""
+#    ## InfluxDB retention policy used when protocol is http. Default=""
+#    # retentionPolicy: ""
+#    ## InfluxDB username (only with http). Default=""
+#    # username: ""
+#    ## InfluxDB password (only with http). Default=""
+#    # password: ""
+#    ## The interval used by the exporter to push metrics to influxdb. Default=10s
+#    # pushInterval: 30s
+#    ## Additional labels (influxdb tags) on all metrics.
+#    # additionalLabels:
+#    #   env: production
+#    #   foo: bar
+#    ## Enable metrics on entry points. Default=true
+#    # addEntryPointsLabels: false
+#    ## Enable metrics on routers. Default=false
+#    # addRoutersLabels: true
+#    ## Enable metrics on services. Default=true
+#    # addServicesLabels: false
+#  influxdb2:
+#    ## Address instructs exporter to send metrics to influxdb v2 at this address.
+#    address: localhost:8086
+#    ## Token with which to connect to InfluxDB v2.
+#    token: xxx
+#    ## Organisation where metrics will be stored.
+#    org: ""
+#    ## Bucket where metrics will be stored.
+#    bucket: ""
+#    ## The interval used by the exporter to push metrics to influxdb. Default=10s
+#    # pushInterval: 30s
+#    ## Additional labels (influxdb tags) on all metrics.
+#    # additionalLabels:
+#    #   env: production
+#    #   foo: bar
+#    ## Enable metrics on entry points. Default=true
+#    # addEntryPointsLabels: false
+#    ## Enable metrics on routers. Default=false
+#    # addRoutersLabels: true
+#    ## Enable metrics on services. Default=true
+#    # addServicesLabels: false
+#  statsd:
+#    ## Address instructs exporter to send metrics to statsd at this address.
+#    address: localhost:8125
+#    ## The interval used by the exporter to push metrics to influxdb. Default=10s
+#    # pushInterval: 30s
+#    ## The prefix to use for metrics collection. Default="traefik"
+#    # prefix: traefik
+#    ## Enable metrics on entry points. Default=true
+#    # addEntryPointsLabels: false
+#    ## Enable metrics on routers. Default=false
+#    # addRoutersLabels: true
+#    ## Enable metrics on services. Default=true
+#    # addServicesLabels: false
+
+
 ##
 ##  enable optional CRDs for Prometheus Operator
 ##
```

## 20.4.1  ![AppVersion: v2.9.4](https://img.shields.io/static/v1?label=AppVersion&message=v2.9.4&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-11-21

* 🐛 fix namespace references to support namespaceOverride


## 20.4.0  ![AppVersion: v2.9.4](https://img.shields.io/static/v1?label=AppVersion&message=v2.9.4&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-11-21

* Add (optional) dedicated metrics service (#727)

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index ca15f6a..e49d02d 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -267,6 +267,12 @@ metrics:
 ##
 ##  enable optional CRDs for Prometheus Operator
 ##
+  ## Create a dedicated metrics service for use with ServiceMonitor
+  ## When hub.enabled is set to true, it's not needed: it will use hub service.
+  #  service:
+  #    enabled: false
+  #    labels: {}
+  #    annotations: {}
   #  serviceMonitor:
   #    metricRelabelings: []
   #      - sourceLabels: [__name__]
```

## 20.3.1  ![AppVersion: v2.9.4](https://img.shields.io/static/v1?label=AppVersion&message=v2.9.4&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-11-21

* 🐛 Fix namespace override which was missing on `ServiceAccount` (#731)


## 20.3.0  ![AppVersion: v2.9.4](https://img.shields.io/static/v1?label=AppVersion&message=v2.9.4&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-11-17

* Add overwrite option for instance label value (#725)

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index c7f84a7..ca15f6a 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -731,3 +731,6 @@ extraObjects: []
 # This will override the default Release Namespace for Helm.
 # It will not affect optional CRDs such as `ServiceMonitor` and `PrometheusRules`
 # namespaceOverride: traefik
+#
+## This will override the default app.kubernetes.io/instance label for all Objects.
+# instanceLabelOverride: traefik
```

## 20.2.1  ![AppVersion: v2.9.4](https://img.shields.io/static/v1?label=AppVersion&message=v2.9.4&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-11-17

* 🙈 do not namespace ingress class (#723)
* ✨ copy LICENSE and README.md on release


## 20.2.0  ![AppVersion: v2.9.4](https://img.shields.io/static/v1?label=AppVersion&message=v2.9.4&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-11-15

* ✨ add support for namespace overrides (#718)
* Document recent changes in the README (#717)

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 97a1b71..c7f84a7 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -725,5 +725,9 @@ podSecurityContext:
 # Extra objects to deploy (value evaluated as a template)
 #
 # In some cases, it can avoid the need for additional, extended or adhoc deployments.
-# See #595 for more details and traefik/tests/extra.yaml for example.
+# See #595 for more details and traefik/tests/values/extra.yaml for example.
 extraObjects: []
+
+# This will override the default Release Namespace for Helm.
+# It will not affect optional CRDs such as `ServiceMonitor` and `PrometheusRules`
+# namespaceOverride: traefik
```

## 20.1.1  ![AppVersion: v2.9.4](https://img.shields.io/static/v1?label=AppVersion&message=v2.9.4&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-11-10

* fix: use consistent appVersion with Traefik Proxy


## 20.1.0  ![AppVersion: 2.9.4](https://img.shields.io/static/v1?label=AppVersion&message=2.9.4&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-11-09

* 🔧 Adds more settings for dashboard ingressRoute (#710)
* 🐛 fix chart releases

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 2ec3736..97a1b71 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -129,10 +129,14 @@ ingressRoute:
     annotations: {}
     # Additional ingressRoute labels (e.g. for filtering IngressRoute by custom labels)
     labels: {}
+    # The router match rule used for the dashboard ingressRoute
+    matchRule: PathPrefix(`/dashboard`) || PathPrefix(`/api`)
     # Specify the allowed entrypoints to use for the dashboard ingress route, (e.g. traefik, web, websecure).
     # By default, it's using traefik entrypoint, which is not exposed.
     # /!\ Do not expose your dashboard without any protection over the internet /!\
     entryPoints: ["traefik"]
+    # Additional ingressRoute middlewares (e.g. for authentication)
+    middlewares: []

 # Customize updateStrategy of traefik pods
 updateStrategy:
```

## 20.0.0  ![AppVersion: 2.9.4](https://img.shields.io/static/v1?label=AppVersion&message=2.9.4&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-11-08

* 🐛 remove old deployment workflow
* ✨ migrate to centralised helm repository
* Allow updateStrategy to be configurable

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 413aa88..2ec3736 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -134,9 +134,12 @@ ingressRoute:
     # /!\ Do not expose your dashboard without any protection over the internet /!\
     entryPoints: ["traefik"]

-rollingUpdate:
-  maxUnavailable: 0
-  maxSurge: 1
+# Customize updateStrategy of traefik pods
+updateStrategy:
+  type: RollingUpdate
+  rollingUpdate:
+    maxUnavailable: 0
+    maxSurge: 1

 # Customize liveness and readiness probe values.
 readinessProbe:
```

## 19.0.4  ![AppVersion: 2.9.4](https://img.shields.io/static/v1?label=AppVersion&message=2.9.4&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-11-08

* 🔧 Adds more settings & rename (wrong) scrapeInterval to (valid) interval on ServiceMonitor (#703)

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index b24c1cb..413aa88 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -261,10 +261,6 @@ metrics:
 ##  enable optional CRDs for Prometheus Operator
 ##
   #  serviceMonitor:
-  #    additionalLabels:
-  #      foo: bar
-  #    namespace: "another-namespace"
-  #    namespaceSelector: {}
   #    metricRelabelings: []
   #      - sourceLabels: [__name__]
   #        separator: ;
@@ -279,9 +275,17 @@ metrics:
   #        replacement: $1
   #        action: replace
   #    jobLabel: traefik
-  #    scrapeInterval: 30s
-  #    scrapeTimeout: 5s
+  #    interval: 30s
   #    honorLabels: true
+  #    # (Optional)
+  #    # scrapeTimeout: 5s
+  #    # honorTimestamps: true
+  #    # enableHttp2: true
+  #    # followRedirects: true
+  #    # additionalLabels:
+  #    #   foo: bar
+  #    # namespace: "another-namespace"
+  #    # namespaceSelector: {}
   #  prometheusRule:
   #    additionalLabels: {}
   #    namespace: "another-namespace"
```

## 19.0.3  ![AppVersion: 2.9.4](https://img.shields.io/static/v1?label=AppVersion&message=2.9.4&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-11-03

* 🎨 Don't require exposed Ports when enabling Hub (#700)


## 19.0.2  ![AppVersion: 2.9.4](https://img.shields.io/static/v1?label=AppVersion&message=2.9.4&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-11-03

* :speech_balloon: Support volume secrets with '.' in name (#695)


## 19.0.1  ![AppVersion: 2.9.4](https://img.shields.io/static/v1?label=AppVersion&message=2.9.4&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-11-03

* 🐛 Fix IngressClass install on EKS (#699)


## 19.0.0  ![AppVersion: 2.9.4](https://img.shields.io/static/v1?label=AppVersion&message=2.9.4&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-11-02

* ✨ Provides Default IngressClass for Traefik by default (#693)

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 69190f1..b24c1cb 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -100,11 +100,10 @@ podDisruptionBudget:
   # minAvailable: 0
   # minAvailable: 25%

-# Use ingressClass. Ignored if Traefik version < 2.3 / kubernetes < 1.18.x
+# Create a default IngressClass for Traefik
 ingressClass:
-  # true is not unit-testable yet, pending https://github.com/rancher/helm-unittest/pull/12
-  enabled: false
-  isDefaultClass: false
+  enabled: true
+  isDefaultClass: true

 # Enable experimental features
 experimental:
```

## 18.3.0  ![AppVersion: 2.9.4](https://img.shields.io/static/v1?label=AppVersion&message=2.9.4&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-10-31

* ⬆️  Update Traefik appVersion to 2.9.4 (#696)


## 18.2.0  ![AppVersion: 2.9.1](https://img.shields.io/static/v1?label=AppVersion&message=2.9.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-10-31

* 🚩 Add an optional "internal" service (#683)

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 8033a87..69190f1 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -416,7 +416,7 @@ ports:
     # The port protocol (TCP/UDP)
     protocol: TCP
     # Use nodeport if set. This is useful if you have configured Traefik in a
-    # LoadBalancer
+    # LoadBalancer.
     # nodePort: 32080
     # Port Redirections
     # Added in 2.2, you can make permanent redirects via entrypoints.
@@ -549,13 +549,24 @@ service:
     # - 172.16.0.0/16
   externalIPs: []
     # - 1.2.3.4
-  # One of SingleStack, PreferDualStack, or RequireDualStack.
+  ## One of SingleStack, PreferDualStack, or RequireDualStack.
   # ipFamilyPolicy: SingleStack
-  # List of IP families (e.g. IPv4 and/or IPv6).
-  # ref: https://kubernetes.io/docs/concepts/services-networking/dual-stack/#services
+  ## List of IP families (e.g. IPv4 and/or IPv6).
+  ## ref: https://kubernetes.io/docs/concepts/services-networking/dual-stack/#services
   # ipFamilies:
   #   - IPv4
   #   - IPv6
+  ##
+  ## An additionnal and optional internal Service.
+  ## Same parameters as external Service
+  # internal:
+  #   type: ClusterIP
+  #   # labels: {}
+  #   # annotations: {}
+  #   # spec: {}
+  #   # loadBalancerSourceRanges: []
+  #   # externalIPs: []
+  #   # ipFamilies: [ "IPv4","IPv6" ]

 ## Create HorizontalPodAutoscaler object.
 ##
```

## 18.1.0  ![AppVersion: 2.9.1](https://img.shields.io/static/v1?label=AppVersion&message=2.9.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-10-27

* 🚀 Add native support for Traefik Hub (#676)

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index acce704..8033a87 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -5,6 +5,27 @@ image:
   tag: ""
   pullPolicy: IfNotPresent

+#
+# Configure integration with Traefik Hub
+#
+hub:
+  ## Enabling Hub will:
+  # * enable Traefik Hub integration on Traefik
+  # * add `traefikhub-tunl` endpoint
+  # * enable addRoutersLabels on prometheus metrics
+  # * enable allowExternalNameServices on KubernetesIngress provider
+  # * enable allowCrossNamespace on KubernetesCRD provider
+  # * add an internal (ClusterIP) Service, dedicated for Traefik Hub
+  enabled: false
+  ## Default port can be changed
+  # tunnelPort: 9901
+  ## TLS is optional. Insecure is mutually exclusive with any other options
+  # tls:
+  #   insecure: false
+  #   ca: "/path/to/ca.pem"
+  #   cert: "/path/to/cert.pem"
+  #   key: "/path/to/key.pem"
+
 #
 # Configure the deployment
 #
@@ -505,6 +526,8 @@ tlsStore: {}
 # from.
 service:
   enabled: true
+  ## Single service is using `MixedProtocolLBService` feature gate.
+  ## When set to false, it will create two Service, one for TCP and one for UDP.
   single: true
   type: LoadBalancer
   # Additional annotations applied to both TCP and UDP services (e.g. for cloud provider specific config)
```

## 18.0.0  ![AppVersion: 2.9.1](https://img.shields.io/static/v1?label=AppVersion&message=2.9.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-10-26

* Refactor http3 and merge TCP with UDP ports into a single service (#656)

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 807bd09..acce704 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -87,8 +87,6 @@ ingressClass:

 # Enable experimental features
 experimental:
-  http3:
-    enabled: false
   plugins:
     enabled: false
   kubernetesGateway:
@@ -421,12 +419,19 @@ ports:
     # The port protocol (TCP/UDP)
     protocol: TCP
     # nodePort: 32443
-    # Enable HTTP/3.
-    # Requires enabling experimental http3 feature and tls.
-    # Note that you cannot have a UDP entrypoint with the same port.
-    # http3: true
-    # Set TLS at the entrypoint
-    # https://doc.traefik.io/traefik/routing/entrypoints/#tls
+    #
+    ## Enable HTTP/3 on the entrypoint
+    ## Enabling it will also enable http3 experimental feature
+    ## https://doc.traefik.io/traefik/routing/entrypoints/#http3
+    ## There are known limitations when trying to listen on same ports for
+    ## TCP & UDP (Http3). There is a workaround in this chart using dual Service.
+    ## https://github.com/kubernetes/kubernetes/issues/47249#issuecomment-587960741
+    http3:
+      enabled: false
+    # advertisedPort: 4443
+    #
+    ## Set TLS at the entrypoint
+    ## https://doc.traefik.io/traefik/routing/entrypoints/#tls
     tls:
       enabled: true
       # this is the name of a TLSOption definition
@@ -500,6 +505,7 @@ tlsStore: {}
 # from.
 service:
   enabled: true
+  single: true
   type: LoadBalancer
   # Additional annotations applied to both TCP and UDP services (e.g. for cloud provider specific config)
   annotations: {}
```

## 17.0.5  ![AppVersion: 2.9.1](https://img.shields.io/static/v1?label=AppVersion&message=2.9.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-10-21

* 📝 Add annotations changelog for artifacthub.io & update Maintainers


## 17.0.4  ![AppVersion: 2.9.1](https://img.shields.io/static/v1?label=AppVersion&message=2.9.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-10-21

* :art: Add helper function for label selector


## 17.0.3  ![AppVersion: 2.9.1](https://img.shields.io/static/v1?label=AppVersion&message=2.9.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-10-20

* 🐛 fix changing label selectors


## 17.0.2  ![AppVersion: 2.9.1](https://img.shields.io/static/v1?label=AppVersion&message=2.9.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-10-20

* fix: setting ports.web.proxyProtocol.insecure=true


## 17.0.1  ![AppVersion: 2.9.1](https://img.shields.io/static/v1?label=AppVersion&message=2.9.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-10-20

* :bug: Unify all labels selector with traefik chart labels (#681)

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 6a90bc6..807bd09 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -639,7 +639,7 @@ affinity: {}
 #      - labelSelector:
 #          matchLabels:
 #            app.kubernetes.io/name: '{{ template "traefik.name" . }}'
-#            app.kubernetes.io/instance: '{{ .Release.Name }}'
+#            app.kubernetes.io/instance: '{{ .Release.Name }}-{{ .Release.Namespace }}'
 #        topologyKey: kubernetes.io/hostname

 nodeSelector: {}
```

## 17.0.0  ![AppVersion: 2.9.1](https://img.shields.io/static/v1?label=AppVersion&message=2.9.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-10-20

* :bug: Fix `ClusterRole`, `ClusterRoleBinding` names and `app.kubernetes.io/instance` label (#662)


## 16.2.0  ![AppVersion: 2.9.1](https://img.shields.io/static/v1?label=AppVersion&message=2.9.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-10-20

* Add forwardedHeaders and proxyProtocol config (#673)

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 9b5afc4..6a90bc6 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -403,6 +403,16 @@ ports:
     # Added in 2.2, you can make permanent redirects via entrypoints.
     # https://docs.traefik.io/routing/entrypoints/#redirection
     # redirectTo: websecure
+    #
+    # Trust forwarded  headers information (X-Forwarded-*).
+    # forwardedHeaders:
+    #   trustedIPs: []
+    #   insecure: false
+    #
+    # Enable the Proxy Protocol header parsing for the entry point
+    # proxyProtocol:
+    #   trustedIPs: []
+    #   insecure: false
   websecure:
     port: 8443
     # hostPort: 8443
@@ -428,6 +438,16 @@ ports:
       #     - foo.example.com
       #     - bar.example.com
     #
+    # Trust forwarded  headers information (X-Forwarded-*).
+    # forwardedHeaders:
+    #   trustedIPs: []
+    #   insecure: false
+    #
+    # Enable the Proxy Protocol header parsing for the entry point
+    # proxyProtocol:
+    #   trustedIPs: []
+    #   insecure: false
+    #
     # One can apply Middlewares on an entrypoint
     # https://doc.traefik.io/traefik/middlewares/overview/
     # https://doc.traefik.io/traefik/routing/entrypoints/#middlewares
```

## 16.1.0  ![AppVersion: 2.9.1](https://img.shields.io/static/v1?label=AppVersion&message=2.9.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-10-19

* ✨ add optional ServiceMonitor & PrometheusRules CRDs (#425)

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 7e335b5..9b5afc4 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -237,8 +237,46 @@ metrics:
   prometheus:
     entryPoint: metrics
   #  addRoutersLabels: true
-  # statsd:
-  #   address: localhost:8125
+  #  statsd:
+  #    address: localhost:8125
+##
+##  enable optional CRDs for Prometheus Operator
+##
+  #  serviceMonitor:
+  #    additionalLabels:
+  #      foo: bar
+  #    namespace: "another-namespace"
+  #    namespaceSelector: {}
+  #    metricRelabelings: []
+  #      - sourceLabels: [__name__]
+  #        separator: ;
+  #        regex: ^fluentd_output_status_buffer_(oldest|newest)_.+
+  #        replacement: $1
+  #        action: drop
+  #    relabelings: []
+  #      - sourceLabels: [__meta_kubernetes_pod_node_name]
+  #        separator: ;
+  #        regex: ^(.*)$
+  #        targetLabel: nodename
+  #        replacement: $1
+  #        action: replace
+  #    jobLabel: traefik
+  #    scrapeInterval: 30s
+  #    scrapeTimeout: 5s
+  #    honorLabels: true
+  #  prometheusRule:
+  #    additionalLabels: {}
+  #    namespace: "another-namespace"
+  #    rules:
+  #      - alert: TraefikDown
+  #        expr: up{job="traefik"} == 0
+  #        for: 5m
+  #        labels:
+  #          context: traefik
+  #          severity: warning
+  #        annotations:
+  #          summary: "Traefik Down"
+  #          description: "{{ $labels.pod }} on {{ $labels.nodename }} is down"

 tracing: {}
   # instana:
```

## 16.0.0  ![AppVersion: 2.9.1](https://img.shields.io/static/v1?label=AppVersion&message=2.9.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-10-19

* :fire: Remove `Pilot` and `fallbackApiVersion` (#665)

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 03fdaed..7e335b5 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -84,15 +84,6 @@ ingressClass:
   # true is not unit-testable yet, pending https://github.com/rancher/helm-unittest/pull/12
   enabled: false
   isDefaultClass: false
-  # Use to force a networking.k8s.io API Version for certain CI/CD applications. E.g. "v1beta1"
-  fallbackApiVersion: ""
-
-# Activate Pilot integration
-pilot:
-  enabled: false
-  token: ""
-  # Toggle Pilot Dashboard
-  # dashboard: false

 # Enable experimental features
 experimental:
```

## 15.3.1  ![AppVersion: 2.9.1](https://img.shields.io/static/v1?label=AppVersion&message=2.9.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-10-18

* :art: Improve `IngressRoute` structure (#674)


## 15.3.0  ![AppVersion: 2.9.1](https://img.shields.io/static/v1?label=AppVersion&message=2.9.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-10-18

* 📌 Add capacity to enable User-facing role

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 76aac93..03fdaed 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -553,10 +553,12 @@ hostNetwork: false
 # Whether Role Based Access Control objects like roles and rolebindings should be created
 rbac:
   enabled: true
-
   # If set to false, installs ClusterRole and ClusterRoleBinding so Traefik can be used across namespaces.
   # If set to true, installs Role and RoleBinding. Providers will only watch target namespace.
   namespaced: false
+  # Enable user-facing roles
+  # https://kubernetes.io/docs/reference/access-authn-authz/rbac/#user-facing-roles
+  # aggregateTo: [ "admin" ]

 # Enable to create a PodSecurityPolicy and assign it to the Service Account via RoleBinding or ClusterRoleBinding
 podSecurityPolicy:
```

## 15.2.2  ![AppVersion: 2.9.1](https://img.shields.io/static/v1?label=AppVersion&message=2.9.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-10-17

* Fix provider namespace changes


## 15.2.1  ![AppVersion: 2.9.1](https://img.shields.io/static/v1?label=AppVersion&message=2.9.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-10-17

* 🐛 fix provider namespace changes


## 15.2.0  ![AppVersion: 2.9.1](https://img.shields.io/static/v1?label=AppVersion&message=2.9.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-10-17

* :bug: Allow to watch on specific namespaces without using rbac.namespaced (#666)

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 781ac15..76aac93 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -555,7 +555,7 @@ rbac:
   enabled: true

   # If set to false, installs ClusterRole and ClusterRoleBinding so Traefik can be used across namespaces.
-  # If set to true, installs namespace-specific Role and RoleBinding and requires provider configuration be set to that same namespace
+  # If set to true, installs Role and RoleBinding. Providers will only watch target namespace.
   namespaced: false

 # Enable to create a PodSecurityPolicy and assign it to the Service Account via RoleBinding or ClusterRoleBinding
```

## 15.1.1  ![AppVersion: 2.9.1](https://img.shields.io/static/v1?label=AppVersion&message=2.9.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-10-17

* :goal_net: Fail gracefully when http3 is not enabled correctly (#667)


## 15.1.0  ![AppVersion: 2.9.1](https://img.shields.io/static/v1?label=AppVersion&message=2.9.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-10-14

* :sparkles: add optional topologySpreadConstraints (#663)

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index fc2c371..781ac15 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -593,6 +593,15 @@ affinity: {}

 nodeSelector: {}
 tolerations: []
+topologySpreadConstraints: []
+# # This example topologySpreadConstraints forces the scheduler to put traefik pods
+# # on nodes where no other traefik pods are scheduled.
+#  - labelSelector:
+#      matchLabels:
+#        app: '{{ template "traefik.name" . }}'
+#    maxSkew: 1
+#    topologyKey: kubernetes.io/hostname
+#    whenUnsatisfiable: DoNotSchedule

 # Pods can have priority.
 # Priority indicates the importance of a Pod relative to other Pods.
```

## 15.0.0  ![AppVersion: 2.9.1](https://img.shields.io/static/v1?label=AppVersion&message=2.9.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-10-13

* :rocket: Enable TLS by default on `websecure` port (#657)

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 400a29a..fc2c371 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -389,7 +389,7 @@ ports:
     # Set TLS at the entrypoint
     # https://doc.traefik.io/traefik/routing/entrypoints/#tls
     tls:
-      enabled: false
+      enabled: true
       # this is the name of a TLSOption definition
       options: ""
       certResolver: ""
```

## 14.0.2  ![AppVersion: 2.9.1](https://img.shields.io/static/v1?label=AppVersion&message=2.9.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-10-13

* :memo: Add Changelog (#661)


## 14.0.1  ![AppVersion: 2.9.1](https://img.shields.io/static/v1?label=AppVersion&message=2.9.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-10-11

* :memo: Update workaround for permissions 660 on acme.json

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index a4e4ff2..400a29a 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -45,10 +45,10 @@ deployment:
   # Additional initContainers (e.g. for setting file permission as shown below)
   initContainers: []
     # The "volume-permissions" init container is required if you run into permission issues.
-    # Related issue: https://github.com/traefik/traefik/issues/6972
+    # Related issue: https://github.com/traefik/traefik/issues/6825
     # - name: volume-permissions
-    #   image: busybox:1.31.1
-    #   command: ["sh", "-c", "chmod -Rv 600 /data/*"]
+    #   image: busybox:1.35
+    #   command: ["sh", "-c", "touch /data/acme.json && chmod -Rv 600 /data/* && chown 65532:65532 /data/acme.json"]
     #   volumeMounts:
     #     - name: data
     #       mountPath: /data
```

## 14.0.0  ![AppVersion: 2.9.1](https://img.shields.io/static/v1?label=AppVersion&message=2.9.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-10-11

* Limit rbac to only required resources for Ingress and CRD providers


## 13.0.1  ![AppVersion: 2.9.1](https://img.shields.io/static/v1?label=AppVersion&message=2.9.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-10-11

* Add helper function for common labels


## 13.0.0  ![AppVersion: 2.9.1](https://img.shields.io/static/v1?label=AppVersion&message=2.9.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-10-11

* Moved list object to individual objects


## 12.0.7  ![AppVersion: 2.9.1](https://img.shields.io/static/v1?label=AppVersion&message=2.9.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-10-10

* :lipstick: Affinity templating and example (#557)

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 4431c36..a4e4ff2 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -578,19 +578,19 @@ resources: {}
   # limits:
   #   cpu: "300m"
   #   memory: "150Mi"
+
+# This example pod anti-affinity forces the scheduler to put traefik pods
+# on nodes where no other traefik pods are scheduled.
+# It should be used when hostNetwork: true to prevent port conflicts
 affinity: {}
-# # This example pod anti-affinity forces the scheduler to put traefik pods
-# # on nodes where no other traefik pods are scheduled.
-# # It should be used when hostNetwork: true to prevent port conflicts
-#   podAntiAffinity:
-#     requiredDuringSchedulingIgnoredDuringExecution:
-#       - labelSelector:
-#           matchExpressions:
-#             - key: app.kubernetes.io/name
-#               operator: In
-#               values:
-#                 - {{ template "traefik.name" . }}
-#         topologyKey: kubernetes.io/hostname
+#  podAntiAffinity:
+#    requiredDuringSchedulingIgnoredDuringExecution:
+#      - labelSelector:
+#          matchLabels:
+#            app.kubernetes.io/name: '{{ template "traefik.name" . }}'
+#            app.kubernetes.io/instance: '{{ .Release.Name }}'
+#        topologyKey: kubernetes.io/hostname
+
 nodeSelector: {}
 tolerations: []

```

## 12.0.6  ![AppVersion: 2.9.1](https://img.shields.io/static/v1?label=AppVersion&message=2.9.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-10-10

* :bug: Ignore kustomization file used for CRDs update (#653)


## 12.0.5  ![AppVersion: 2.9.1](https://img.shields.io/static/v1?label=AppVersion&message=2.9.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-10-10

* :memo: Establish Traefik & CRD update process

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 3526729..4431c36 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -342,6 +342,7 @@ ports:

     # Override the liveness/readiness port. This is useful to integrate traefik
     # with an external Load Balancer that performs healthchecks.
+    # Default: ports.traefik.port
     # healthchecksPort: 9000

     # Override the liveness/readiness scheme. Useful for getting ping to
```

## 12.0.4  ![AppVersion: 2.9.1](https://img.shields.io/static/v1?label=AppVersion&message=2.9.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-10-10

* Allows ingressClass to be used without semver-compatible image tag


## 12.0.3  ![AppVersion: 2.9.1](https://img.shields.io/static/v1?label=AppVersion&message=2.9.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-10-10

* :bug: Should check hostNetwork when hostPort != containerPort


## 12.0.2  ![AppVersion: 2.9.1](https://img.shields.io/static/v1?label=AppVersion&message=2.9.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-10-07

* :goal_net: Fail gracefully when hostNetwork is enabled and hostPort != containerPort


## 12.0.1  ![AppVersion: 2.9.1](https://img.shields.io/static/v1?label=AppVersion&message=2.9.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-10-07

* :bug: Fix a typo on `behavior` for HPA v2


## 12.0.0  ![AppVersion: 2.9.1](https://img.shields.io/static/v1?label=AppVersion&message=2.9.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-10-06

* Update default HPA API Version to `v2` and add support for behavior (#518)

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 2bd51f8..3526729 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -488,11 +488,22 @@ autoscaling:
 #   - type: Resource
 #     resource:
 #       name: cpu
-#       targetAverageUtilization: 60
+#       target:
+#         type: Utilization
+#         averageUtilization: 60
 #   - type: Resource
 #     resource:
 #       name: memory
-#       targetAverageUtilization: 60
+#       target:
+#         type: Utilization
+#         averageUtilization: 60
+#   behavior:
+#     scaleDown:
+#       stabilizationWindowSeconds: 300
+#       policies:
+#       - type: Pods
+#         value: 1
+#         periodSeconds: 60

 # Enable persistence using Persistent Volume Claims
 # ref: http://kubernetes.io/docs/user-guide/persistent-volumes/
```

## 11.1.1  ![AppVersion: 2.9.1](https://img.shields.io/static/v1?label=AppVersion&message=2.9.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-10-05

* 🔊 add failure message when using maxUnavailable 0 and hostNetwork


## 11.1.0  ![AppVersion: 2.9.1](https://img.shields.io/static/v1?label=AppVersion&message=2.9.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-10-04

* Update Traefik to v2.9.1


## 11.0.0  ![AppVersion: 2.8.7](https://img.shields.io/static/v1?label=AppVersion&message=2.8.7&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-10-04

* tweak default values to avoid downtime when updating

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 844cadc..2bd51f8 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -126,20 +126,20 @@ ingressRoute:
     entryPoints: ["traefik"]

 rollingUpdate:
-  maxUnavailable: 1
+  maxUnavailable: 0
   maxSurge: 1

 # Customize liveness and readiness probe values.
 readinessProbe:
   failureThreshold: 1
-  initialDelaySeconds: 10
+  initialDelaySeconds: 2
   periodSeconds: 10
   successThreshold: 1
   timeoutSeconds: 2

 livenessProbe:
   failureThreshold: 3
-  initialDelaySeconds: 10
+  initialDelaySeconds: 2
   periodSeconds: 10
   successThreshold: 1
   timeoutSeconds: 2
```

## 10.33.0  ![AppVersion: 2.8.7](https://img.shields.io/static/v1?label=AppVersion&message=2.8.7&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-10-04

* :rocket: Add `extraObjects` value that allows creating adhoc resources

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index c926bd9..844cadc 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -598,3 +598,10 @@ securityContext:

 podSecurityContext:
   fsGroup: 65532
+
+#
+# Extra objects to deploy (value evaluated as a template)
+#
+# In some cases, it can avoid the need for additional, extended or adhoc deployments.
+# See #595 for more details and traefik/tests/extra.yaml for example.
+extraObjects: []
```

## 10.32.0  ![AppVersion: 2.8.7](https://img.shields.io/static/v1?label=AppVersion&message=2.8.7&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-10-03

* Add support setting middleware on entrypoint

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 3957448..c926bd9 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -397,6 +397,16 @@ ports:
       #   sans:
       #     - foo.example.com
       #     - bar.example.com
+    #
+    # One can apply Middlewares on an entrypoint
+    # https://doc.traefik.io/traefik/middlewares/overview/
+    # https://doc.traefik.io/traefik/routing/entrypoints/#middlewares
+    # /!\ It introduces here a link between your static configuration and your dynamic configuration /!\
+    # It follows the provider naming convention: https://doc.traefik.io/traefik/providers/overview/#provider-namespace
+    # middlewares:
+    #   - namespace-name1@kubernetescrd
+    #   - namespace-name2@kubernetescrd
+    middlewares: []
   metrics:
     # When using hostNetwork, use another port to avoid conflict with node exporter:
     # https://github.com/prometheus/prometheus/wiki/Default-port-allocations
```

## 10.31.0  ![AppVersion: 2.8.7](https://img.shields.io/static/v1?label=AppVersion&message=2.8.7&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-10-03

* Support setting dashboard entryPoints for ingressRoute resource

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index c9feb76..3957448 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -120,6 +120,10 @@ ingressRoute:
     annotations: {}
     # Additional ingressRoute labels (e.g. for filtering IngressRoute by custom labels)
     labels: {}
+    # Specify the allowed entrypoints to use for the dashboard ingress route, (e.g. traefik, web, websecure).
+    # By default, it's using traefik entrypoint, which is not exposed.
+    # /!\ Do not expose your dashboard without any protection over the internet /!\
+    entryPoints: ["traefik"]

 rollingUpdate:
   maxUnavailable: 1
```

## 10.30.2  ![AppVersion: 2.8.7](https://img.shields.io/static/v1?label=AppVersion&message=2.8.7&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-10-03

* :test_tube: Fail gracefully when asked to provide a service without ports


## 10.30.1  ![AppVersion: 2.8.7](https://img.shields.io/static/v1?label=AppVersion&message=2.8.7&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-09-30

* :arrow_up: Upgrade helm, ct & unittest (#638)


## 10.30.0  ![AppVersion: 2.8.7](https://img.shields.io/static/v1?label=AppVersion&message=2.8.7&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-09-30

* Add support HTTPS scheme for healthcheks

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index fed4a8a..c9feb76 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -340,6 +340,10 @@ ports:
     # with an external Load Balancer that performs healthchecks.
     # healthchecksPort: 9000

+    # Override the liveness/readiness scheme. Useful for getting ping to
+    # respond on websecure entryPoint.
+    # healthchecksScheme: HTTPS
+
     # Defines whether the port is exposed if service.type is LoadBalancer or
     # NodePort.
     #
```

## 10.29.0  ![AppVersion: 2.8.7](https://img.shields.io/static/v1?label=AppVersion&message=2.8.7&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-09-29

* Add missing tracing options

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index d1708cc..fed4a8a 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -247,12 +247,45 @@ metrics:

 tracing: {}
   # instana:
-  #   enabled: true
+  #   localAgentHost: 127.0.0.1
+  #   localAgentPort: 42699
+  #   logLevel: info
+  #   enableAutoProfile: true
   # datadog:
   #   localAgentHostPort: 127.0.0.1:8126
   #   debug: false
   #   globalTag: ""
   #   prioritySampling: false
+  # jaeger:
+  #   samplingServerURL: http://localhost:5778/sampling
+  #   samplingType: const
+  #   samplingParam: 1.0
+  #   localAgentHostPort: 127.0.0.1:6831
+  #   gen128Bit: false
+  #   propagation: jaeger
+  #   traceContextHeaderName: uber-trace-id
+  #   disableAttemptReconnecting: true
+  #   collector:
+  #      endpoint: ""
+  #      user: ""
+  #      password: ""
+  # zipkin:
+  #   httpEndpoint: http://localhost:9411/api/v2/spans
+  #   sameSpan: false
+  #   id128Bit: true
+  #   sampleRate: 1.0
+  # haystack:
+  #   localAgentHost: 127.0.0.1
+  #   localAgentPort: 35000
+  #   globalTag: ""
+  #   traceIDHeaderName: ""
+  #   parentIDHeaderName: ""
+  #   spanIDHeaderName: ""
+  #   baggagePrefixHeaderName: ""
+  # elastic:
+  #   serverURL: http://localhost:8200
+  #   secretToken: ""
+  #   serviceEnvironment: ""

 globalArguments:
   - "--global.checknewversion"
```

## 10.28.0  ![AppVersion: 2.8.7](https://img.shields.io/static/v1?label=AppVersion&message=2.8.7&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-09-29

* feat: add lifecycle for prestop and poststart

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 19a133c..d1708cc 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -59,6 +59,17 @@ deployment:
   # Additional imagePullSecrets
   imagePullSecrets: []
     # - name: myRegistryKeySecretName
+  # Pod lifecycle actions
+  lifecycle: {}
+    # preStop:
+    #   exec:
+    #     command: ["/bin/sh", "-c", "sleep 40"]
+    # postStart:
+    #   httpGet:
+    #     path: /ping
+    #     port: 9000
+    #     host: localhost
+    #     scheme: HTTP

 # Pod disruption budget
 podDisruptionBudget:
```

## 10.27.0  ![AppVersion: 2.8.7](https://img.shields.io/static/v1?label=AppVersion&message=2.8.7&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-09-29

* feat: add create gateway option

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index d9c745e..19a133c 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -91,6 +91,8 @@ experimental:
     enabled: false
   kubernetesGateway:
     enabled: false
+    gateway:
+      enabled: true
     # certificate:
     #   group: "core"
     #   kind: "Secret"
```

## 10.26.1  ![AppVersion: 2.8.7](https://img.shields.io/static/v1?label=AppVersion&message=2.8.7&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-09-28

* 🐛 fix rbac templating (#636)


## 10.26.0  ![AppVersion: 2.8.7](https://img.shields.io/static/v1?label=AppVersion&message=2.8.7&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-09-28

* :bug: Fix ingressClass support when rbac.namespaced=true (#499)


## 10.25.1  ![AppVersion: 2.8.7](https://img.shields.io/static/v1?label=AppVersion&message=2.8.7&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-09-28

* Add ingressclasses to traefik role


## 10.25.0  ![AppVersion: 2.8.7](https://img.shields.io/static/v1?label=AppVersion&message=2.8.7&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-09-27

* Add TLSStore resource to chart

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index d4011c3..d9c745e 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -373,6 +373,15 @@ ports:
 #       - CurveP384
 tlsOptions: {}

+# TLS Store are created as TLSStore CRDs. This is useful if you want to set a default certificate
+# https://doc.traefik.io/traefik/https/tls/#default-certificate
+# Example:
+# tlsStore:
+#   default:
+#     defaultCertificate:
+#       secretName: tls-cert
+tlsStore: {}
+
 # Options for the main traefik service, where the entrypoints traffic comes
 # from.
 service:
```

## 10.24.5  ![AppVersion: 2.8.7](https://img.shields.io/static/v1?label=AppVersion&message=2.8.7&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-09-27

* Suggest an alternative port for metrics

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 81f2e85..d4011c3 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -344,6 +344,8 @@ ports:
       #     - foo.example.com
       #     - bar.example.com
   metrics:
+    # When using hostNetwork, use another port to avoid conflict with node exporter:
+    # https://github.com/prometheus/prometheus/wiki/Default-port-allocations
     port: 9100
     # hostPort: 9100
     # Defines whether the port is exposed if service.type is LoadBalancer or
```

## 10.24.4  ![AppVersion: 2.8.7](https://img.shields.io/static/v1?label=AppVersion&message=2.8.7&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-09-26

* Update Traefik to v2.8.7


## 10.24.3  ![AppVersion: 2.8.5](https://img.shields.io/static/v1?label=AppVersion&message=2.8.5&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-09-14

* Update Traefik version to v2.8.5


## 10.24.2  ![AppVersion: 2.8.4](https://img.shields.io/static/v1?label=AppVersion&message=2.8.4&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-09-05

* Update Traefik version to v2.8.4


## 10.24.1  ![AppVersion: 2.8.0](https://img.shields.io/static/v1?label=AppVersion&message=2.8.0&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-08-29

* Update PodDisruptionBudget apiVersion to policy/v1


## 10.24.0  ![AppVersion: 2.8.0](https://img.shields.io/static/v1?label=AppVersion&message=2.8.0&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-06-30

* Update Traefik version to v2.8.0


## 10.23.0  ![AppVersion: 2.7.1](https://img.shields.io/static/v1?label=AppVersion&message=2.7.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-06-27

* Support environment variable usage for Datadog


## 10.22.0  ![AppVersion: 2.7.1](https://img.shields.io/static/v1?label=AppVersion&message=2.7.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-06-22

* Allow setting revisionHistoryLimit for Deployment and DaemonSet

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index d5785ab..81f2e85 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -14,6 +14,8 @@ deployment:
   kind: Deployment
   # Number of pods of the deployment (only applies when kind == Deployment)
   replicas: 1
+  # Number of old history to retain to allow rollback (If not set, default Kubernetes value is set to 10)
+  # revisionHistoryLimit: 1
   # Amount of time (in seconds) before Kubernetes will send the SIGKILL signal if Traefik does not shut down
   terminationGracePeriodSeconds: 60
   # The minimum number of seconds Traefik needs to be up and running before the DaemonSet/Deployment controller considers it available
```

## 10.21.1  ![AppVersion: 2.7.1](https://img.shields.io/static/v1?label=AppVersion&message=2.7.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-06-15

* Update Traefik version to 2.7.1


## 10.21.0  ![AppVersion: 2.7.0](https://img.shields.io/static/v1?label=AppVersion&message=2.7.0&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-06-15

* Support allowEmptyServices config for KubernetesCRD

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index e141e29..d5785ab 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -133,6 +133,7 @@ providers:
     enabled: true
     allowCrossNamespace: false
     allowExternalNameServices: false
+    allowEmptyServices: false
     # ingressClass: traefik-internal
     # labelSelector: environment=production,method=traefik
     namespaces: []
```

## 10.20.1  ![AppVersion: 2.7.0](https://img.shields.io/static/v1?label=AppVersion&message=2.7.0&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-06-01

* Add Acme certificate resolver configuration

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index a16b107..e141e29 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -433,6 +433,27 @@ persistence:
   annotations: {}
   # subPath: "" # only mount a subpath of the Volume into the pod

+certResolvers: {}
+#   letsencrypt:
+#     # for challenge options cf. https://doc.traefik.io/traefik/https/acme/
+#     email: email@example.com
+#     dnsChallenge:
+#       # also add the provider's required configuration under env
+#       # or expand then from secrets/configmaps with envfrom
+#       # cf. https://doc.traefik.io/traefik/https/acme/#providers
+#       provider: digitalocean
+#       # add futher options for the dns challenge as needed
+#       # cf. https://doc.traefik.io/traefik/https/acme/#dnschallenge
+#       delayBeforeCheck: 30
+#       resolvers:
+#         - 1.1.1.1
+#         - 8.8.8.8
+#     tlsChallenge: true
+#     httpChallenge:
+#       entryPoint: "web"
+#     # match the path to persistence
+#     storage: /data/acme.json
+
 # If hostNetwork is true, runs traefik in the host network namespace
 # To prevent unschedulabel pods due to port collisions, if hostNetwork=true
 # and replicas>1, a pod anti-affinity is recommended and will be set if the
```

## 10.20.0  ![AppVersion: 2.7.0](https://img.shields.io/static/v1?label=AppVersion&message=2.7.0&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-05-25

* Update Traefik Proxy to v2.7.0


## 10.19.5  ![AppVersion: 2.6.6](https://img.shields.io/static/v1?label=AppVersion&message=2.6.6&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-05-04

* Upgrade Traefik to 2.6.6


## 10.19.4  ![AppVersion: 2.6.3](https://img.shields.io/static/v1?label=AppVersion&message=2.6.3&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-03-31

* Update Traefik dependency version to 2.6.3


## 10.19.3  ![AppVersion: 2.6.2](https://img.shields.io/static/v1?label=AppVersion&message=2.6.2&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-03-30

* Update CRDs to match the ones defined in the reference documentation


## 10.19.2  ![AppVersion: 2.6.2](https://img.shields.io/static/v1?label=AppVersion&message=2.6.2&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-03-30

* Revert Traefik version to 2.6.2


## 10.19.1  ![AppVersion: 2.6.3](https://img.shields.io/static/v1?label=AppVersion&message=2.6.3&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-03-30

* Update Traefik version to 2.6.3


## 10.19.0  ![AppVersion: 2.6.2](https://img.shields.io/static/v1?label=AppVersion&message=2.6.2&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-03-28

* Support ingressClass option for KubernetesIngress provider

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 02ab704..a16b107 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -142,6 +142,7 @@ providers:
     enabled: true
     allowExternalNameServices: false
     allowEmptyServices: false
+    # ingressClass: traefik-internal
     # labelSelector: environment=production,method=traefik
     namespaces: []
       # - "default"
```

## 10.18.0  ![AppVersion: 2.6.2](https://img.shields.io/static/v1?label=AppVersion&message=2.6.2&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-03-28

* Support liveness and readyness probes customization

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 15f1103..02ab704 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -110,6 +110,20 @@ rollingUpdate:
   maxUnavailable: 1
   maxSurge: 1

+# Customize liveness and readiness probe values.
+readinessProbe:
+  failureThreshold: 1
+  initialDelaySeconds: 10
+  periodSeconds: 10
+  successThreshold: 1
+  timeoutSeconds: 2
+
+livenessProbe:
+  failureThreshold: 3
+  initialDelaySeconds: 10
+  periodSeconds: 10
+  successThreshold: 1
+  timeoutSeconds: 2

 #
 # Configure providers
```

## 10.17.0  ![AppVersion: 2.6.2](https://img.shields.io/static/v1?label=AppVersion&message=2.6.2&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-03-28

* Support Datadog tracing

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 4dccd1a..15f1103 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -217,6 +217,11 @@ metrics:
 tracing: {}
   # instana:
   #   enabled: true
+  # datadog:
+  #   localAgentHostPort: 127.0.0.1:8126
+  #   debug: false
+  #   globalTag: ""
+  #   prioritySampling: false

 globalArguments:
   - "--global.checknewversion"
```

## 10.16.1  ![AppVersion: 2.6.2](https://img.shields.io/static/v1?label=AppVersion&message=2.6.2&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-03-28

* Update Traefik version to 2.6.2


## 10.16.0  ![AppVersion: 2.6.1](https://img.shields.io/static/v1?label=AppVersion&message=2.6.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-03-28

* Support allowEmptyServices for KubernetesIngress provider

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 1f9dbbe..4dccd1a 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -127,6 +127,7 @@ providers:
   kubernetesIngress:
     enabled: true
     allowExternalNameServices: false
+    allowEmptyServices: false
     # labelSelector: environment=production,method=traefik
     namespaces: []
       # - "default"
```

## 10.15.0  ![AppVersion: 2.6.1](https://img.shields.io/static/v1?label=AppVersion&message=2.6.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-03-08

* Add metrics.prometheus.addRoutersLabels option

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index cd4d49b..1f9dbbe 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -209,6 +209,7 @@ metrics:
   #   protocol: udp
   prometheus:
     entryPoint: metrics
+  #  addRoutersLabels: true
   # statsd:
   #   address: localhost:8125

```

## 10.14.2  ![AppVersion: 2.6.1](https://img.shields.io/static/v1?label=AppVersion&message=2.6.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-02-18

* Update Traefik to v2.6.1


## 10.14.1  ![AppVersion: 2.6.0](https://img.shields.io/static/v1?label=AppVersion&message=2.6.0&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-02-09

* Add missing inFlightConn TCP middleware CRD


## 10.14.0  ![AppVersion: 2.6.0](https://img.shields.io/static/v1?label=AppVersion&message=2.6.0&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-02-03

* Add experimental HTTP/3 support

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index d49122f..cd4d49b 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -83,6 +83,8 @@ pilot:

 # Enable experimental features
 experimental:
+  http3:
+    enabled: false
   plugins:
     enabled: false
   kubernetesGateway:
@@ -300,6 +302,10 @@ ports:
     # The port protocol (TCP/UDP)
     protocol: TCP
     # nodePort: 32443
+    # Enable HTTP/3.
+    # Requires enabling experimental http3 feature and tls.
+    # Note that you cannot have a UDP entrypoint with the same port.
+    # http3: true
     # Set TLS at the entrypoint
     # https://doc.traefik.io/traefik/routing/entrypoints/#tls
     tls:
```

## 10.13.0  ![AppVersion: 2.6.0](https://img.shields.io/static/v1?label=AppVersion&message=2.6.0&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-02-01

* Add support for ipFamilies

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 32fce6f..d49122f 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -366,6 +366,11 @@ service:
     # - 1.2.3.4
   # One of SingleStack, PreferDualStack, or RequireDualStack.
   # ipFamilyPolicy: SingleStack
+  # List of IP families (e.g. IPv4 and/or IPv6).
+  # ref: https://kubernetes.io/docs/concepts/services-networking/dual-stack/#services
+  # ipFamilies:
+  #   - IPv4
+  #   - IPv6

 ## Create HorizontalPodAutoscaler object.
 ##
```

## 10.12.0  ![AppVersion: 2.6.0](https://img.shields.io/static/v1?label=AppVersion&message=2.6.0&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-02-01

* Add shareProcessNamespace option to podtemplate

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index ab25456..32fce6f 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -50,6 +50,8 @@ deployment:
     #   volumeMounts:
     #     - name: data
     #       mountPath: /data
+  # Use process namespace sharing
+  shareProcessNamespace: false
   # Custom pod DNS policy. Apply if `hostNetwork: true`
   # dnsPolicy: ClusterFirstWithHostNet
   # Additional imagePullSecrets
```

## 10.11.1  ![AppVersion: 2.6.0](https://img.shields.io/static/v1?label=AppVersion&message=2.6.0&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-01-31

* Fix anti-affinity example

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 8c72905..ab25456 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -438,13 +438,13 @@ affinity: {}
 # # It should be used when hostNetwork: true to prevent port conflicts
 #   podAntiAffinity:
 #     requiredDuringSchedulingIgnoredDuringExecution:
-#     - labelSelector:
-#         matchExpressions:
-#         - key: app
-#           operator: In
-#           values:
-#           - {{ template "traefik.name" . }}
-#       topologyKey: failure-domain.beta.kubernetes.io/zone
+#       - labelSelector:
+#           matchExpressions:
+#             - key: app.kubernetes.io/name
+#               operator: In
+#               values:
+#                 - {{ template "traefik.name" . }}
+#         topologyKey: kubernetes.io/hostname
 nodeSelector: {}
 tolerations: []

```

## 10.11.0  ![AppVersion: 2.6.0](https://img.shields.io/static/v1?label=AppVersion&message=2.6.0&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-01-31

* Add setting to enable Instana tracing

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 7fe4a2c..8c72905 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -208,6 +208,10 @@ metrics:
   # statsd:
   #   address: localhost:8125

+tracing: {}
+  # instana:
+  #   enabled: true
+
 globalArguments:
   - "--global.checknewversion"
   - "--global.sendanonymoususage"
```

## 10.10.0  ![AppVersion: 2.6.0](https://img.shields.io/static/v1?label=AppVersion&message=2.6.0&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2022-01-31

* Update Traefik to v2.6

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 8ae4bd8..7fe4a2c 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -85,9 +85,8 @@ experimental:
     enabled: false
   kubernetesGateway:
     enabled: false
-    appLabelSelector: "traefik"
-    certificates: []
-    # - group: "core"
+    # certificate:
+    #   group: "core"
     #   kind: "Secret"
     #   name: "mysecret"
     # By default, Gateway would be created to the Namespace you are deploying Traefik to.
```

## 10.9.1  ![AppVersion: 2.5.6](https://img.shields.io/static/v1?label=AppVersion&message=2.5.6&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-12-24

* Bump traefik version to 2.5.6


## 10.9.0  ![AppVersion: 2.5.4](https://img.shields.io/static/v1?label=AppVersion&message=2.5.4&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-12-20

* feat: add allowExternalNameServices to KubernetesIngress provider

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 79df205..8ae4bd8 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -123,6 +123,7 @@ providers:

   kubernetesIngress:
     enabled: true
+    allowExternalNameServices: false
     # labelSelector: environment=production,method=traefik
     namespaces: []
       # - "default"
```

## 10.8.0  ![AppVersion: 2.5.4](https://img.shields.io/static/v1?label=AppVersion&message=2.5.4&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-12-20

* Add support to specify minReadySeconds on Deployment/DaemonSet

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 7e9186b..79df205 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -16,6 +16,8 @@ deployment:
   replicas: 1
   # Amount of time (in seconds) before Kubernetes will send the SIGKILL signal if Traefik does not shut down
   terminationGracePeriodSeconds: 60
+  # The minimum number of seconds Traefik needs to be up and running before the DaemonSet/Deployment controller considers it available
+  minReadySeconds: 0
   # Additional deployment annotations (e.g. for jaeger-operator sidecar injection)
   annotations: {}
   # Additional deployment labels (e.g. for filtering deployment by custom labels)
```

## 10.7.1  ![AppVersion: 2.5.4](https://img.shields.io/static/v1?label=AppVersion&message=2.5.4&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-12-06

* Fix pod disruption when using percentages

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index e0655c8..7e9186b 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -52,13 +52,15 @@ deployment:
   # dnsPolicy: ClusterFirstWithHostNet
   # Additional imagePullSecrets
   imagePullSecrets: []
-   # - name: myRegistryKeySecretName
+    # - name: myRegistryKeySecretName

 # Pod disruption budget
 podDisruptionBudget:
   enabled: false
   # maxUnavailable: 1
+  # maxUnavailable: 33%
   # minAvailable: 0
+  # minAvailable: 25%

 # Use ingressClass. Ignored if Traefik version < 2.3 / kubernetes < 1.18.x
 ingressClass:
```

## 10.7.0  ![AppVersion: 2.5.4](https://img.shields.io/static/v1?label=AppVersion&message=2.5.4&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-12-06

* Add support for ipFamilyPolicy

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 3ec7105..e0655c8 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -343,8 +343,8 @@ service:
   annotationsUDP: {}
   # Additional service labels (e.g. for filtering Service by custom labels)
   labels: {}
-  # Additional entries here will be added to the service spec. Cannot contains
-  # type, selector or ports entries.
+  # Additional entries here will be added to the service spec.
+  # Cannot contain type, selector or ports entries.
   spec: {}
     # externalTrafficPolicy: Cluster
     # loadBalancerIP: "1.2.3.4"
@@ -354,6 +354,8 @@ service:
     # - 172.16.0.0/16
   externalIPs: []
     # - 1.2.3.4
+  # One of SingleStack, PreferDualStack, or RequireDualStack.
+  # ipFamilyPolicy: SingleStack

 ## Create HorizontalPodAutoscaler object.
 ##
```

## 10.6.2  ![AppVersion: 2.5.4](https://img.shields.io/static/v1?label=AppVersion&message=2.5.4&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-11-15

* Bump Traefik version to 2.5.4


## 10.6.1  ![AppVersion: 2.5.3](https://img.shields.io/static/v1?label=AppVersion&message=2.5.3&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-11-05

* Add missing Gateway API resources to ClusterRole


## 10.6.0  ![AppVersion: 2.5.3](https://img.shields.io/static/v1?label=AppVersion&message=2.5.3&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-10-13

* feat: allow termination grace period to be configurable

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index f06ebc6..3ec7105 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -14,6 +14,8 @@ deployment:
   kind: Deployment
   # Number of pods of the deployment (only applies when kind == Deployment)
   replicas: 1
+  # Amount of time (in seconds) before Kubernetes will send the SIGKILL signal if Traefik does not shut down
+  terminationGracePeriodSeconds: 60
   # Additional deployment annotations (e.g. for jaeger-operator sidecar injection)
   annotations: {}
   # Additional deployment labels (e.g. for filtering deployment by custom labels)
```

## 10.5.0  ![AppVersion: 2.5.3](https://img.shields.io/static/v1?label=AppVersion&message=2.5.3&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-10-13

* feat: add allowExternalNameServices to Kubernetes CRD provider

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 3bcb350..f06ebc6 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -109,6 +109,7 @@ providers:
   kubernetesCRD:
     enabled: true
     allowCrossNamespace: false
+    allowExternalNameServices: false
     # ingressClass: traefik-internal
     # labelSelector: environment=production,method=traefik
     namespaces: []
```

## 10.4.2  ![AppVersion: 2.5.3](https://img.shields.io/static/v1?label=AppVersion&message=2.5.3&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-10-13

* fix(crd): add permissionsPolicy to headers middleware


## 10.4.1  ![AppVersion: 2.5.3](https://img.shields.io/static/v1?label=AppVersion&message=2.5.3&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-10-13

* fix(crd): add peerCertURI option to ServersTransport


## 10.4.0  ![AppVersion: 2.5.3](https://img.shields.io/static/v1?label=AppVersion&message=2.5.3&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-10-12

* Add Kubernetes CRD labelSelector and ingressClass options

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index f54f5fe..3bcb350 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -109,8 +109,11 @@ providers:
   kubernetesCRD:
     enabled: true
     allowCrossNamespace: false
+    # ingressClass: traefik-internal
+    # labelSelector: environment=production,method=traefik
     namespaces: []
       # - "default"
+
   kubernetesIngress:
     enabled: true
     # labelSelector: environment=production,method=traefik
```

## 10.3.6  ![AppVersion: 2.5.3](https://img.shields.io/static/v1?label=AppVersion&message=2.5.3&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-09-24

* Fix missing RequireAnyClientCert value to TLSOption CRD


## 10.3.5  ![AppVersion: 2.5.3](https://img.shields.io/static/v1?label=AppVersion&message=2.5.3&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-09-23

* Bump Traefik version to 2.5.3


## 10.3.4  ![AppVersion: 2.5.1](https://img.shields.io/static/v1?label=AppVersion&message=2.5.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-09-17

* Add allowCrossNamespace option on kubernetesCRD provider

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 7e3a579..f54f5fe 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -108,6 +108,7 @@ rollingUpdate:
 providers:
   kubernetesCRD:
     enabled: true
+    allowCrossNamespace: false
     namespaces: []
       # - "default"
   kubernetesIngress:
```

## 10.3.3  ![AppVersion: 2.5.1](https://img.shields.io/static/v1?label=AppVersion&message=2.5.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-09-17

* fix(crd): missing alpnProtocols in TLSOption


## 10.3.2  ![AppVersion: 2.5.1](https://img.shields.io/static/v1?label=AppVersion&message=2.5.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-08-23

* Releasing 2.5.1


## 10.3.1  ![AppVersion: 2.5.0](https://img.shields.io/static/v1?label=AppVersion&message=2.5.0&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-08-20

* Fix Ingress RBAC for namespaced scoped deployment


## 10.3.0  ![AppVersion: 2.5.0](https://img.shields.io/static/v1?label=AppVersion&message=2.5.0&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-08-18

* Releasing Traefik 2.5.0


## 10.2.0  ![AppVersion: 2.4.13](https://img.shields.io/static/v1?label=AppVersion&message=2.4.13&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-08-18

* Allow setting TCP and UDP service annotations separately

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 72a01ea..7e3a579 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -328,8 +328,12 @@ tlsOptions: {}
 service:
   enabled: true
   type: LoadBalancer
-  # Additional annotations (e.g. for cloud provider specific config)
+  # Additional annotations applied to both TCP and UDP services (e.g. for cloud provider specific config)
   annotations: {}
+  # Additional annotations for TCP service only
+  annotationsTCP: {}
+  # Additional annotations for UDP service only
+  annotationsUDP: {}
   # Additional service labels (e.g. for filtering Service by custom labels)
   labels: {}
   # Additional entries here will be added to the service spec. Cannot contains
```

## 10.1.6  ![AppVersion: 2.4.13](https://img.shields.io/static/v1?label=AppVersion&message=2.4.13&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-08-17

* fix: missing service labels


## 10.1.5  ![AppVersion: 2.4.13](https://img.shields.io/static/v1?label=AppVersion&message=2.4.13&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-08-17

* fix(pvc-annotaions): see traefik/traefik-helm-chart#471


## 10.1.4  ![AppVersion: 2.4.13](https://img.shields.io/static/v1?label=AppVersion&message=2.4.13&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-08-17

* fix(ingressclass): fallbackApiVersion default shouldn't be `nil`

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 04d336c..72a01ea 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -64,7 +64,7 @@ ingressClass:
   enabled: false
   isDefaultClass: false
   # Use to force a networking.k8s.io API Version for certain CI/CD applications. E.g. "v1beta1"
-  fallbackApiVersion:
+  fallbackApiVersion: ""

 # Activate Pilot integration
 pilot:
```

## 10.1.3  ![AppVersion: 2.4.13](https://img.shields.io/static/v1?label=AppVersion&message=2.4.13&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-08-16

* Move Prometheus annotations to Pods


## 10.1.2  ![AppVersion: 2.4.13](https://img.shields.io/static/v1?label=AppVersion&message=2.4.13&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-08-10

* Version bumped 2.4.13


## 10.1.1  ![AppVersion: 2.4.9](https://img.shields.io/static/v1?label=AppVersion&message=2.4.9&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-07-20

* Fixing Prometheus.io/port annotation


## 10.1.0  ![AppVersion: 2.4.9](https://img.shields.io/static/v1?label=AppVersion&message=2.4.9&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-07-20

* Add metrics framework, and prom annotations

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index f6e370a..04d336c 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -186,6 +186,17 @@ logs:
           # Authorization: drop
           # Content-Type: keep

+metrics:
+  # datadog:
+  #   address: 127.0.0.1:8125
+  # influxdb:
+  #   address: localhost:8089
+  #   protocol: udp
+  prometheus:
+    entryPoint: metrics
+  # statsd:
+  #   address: localhost:8125
+
 globalArguments:
   - "--global.checknewversion"
   - "--global.sendanonymoususage"
@@ -284,6 +295,20 @@ ports:
       #   sans:
       #     - foo.example.com
       #     - bar.example.com
+  metrics:
+    port: 9100
+    # hostPort: 9100
+    # Defines whether the port is exposed if service.type is LoadBalancer or
+    # NodePort.
+    #
+    # You may not want to expose the metrics port on production deployments.
+    # If you want to access it from outside of your cluster,
+    # use `kubectl port-forward` or create a secure ingress
+    expose: false
+    # The exposed port for this service
+    exposedPort: 9100
+    # The port protocol (TCP/UDP)
+    protocol: TCP

 # TLS Options are created as TLSOption CRDs
 # https://doc.traefik.io/traefik/https/tls/#tls-options
```

## 10.0.2  ![AppVersion: 2.4.9](https://img.shields.io/static/v1?label=AppVersion&message=2.4.9&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-07-14

* feat(gateway): introduces param / pick Namespace installing Gateway

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 9bf90ea..f6e370a 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -84,6 +84,9 @@ experimental:
     # - group: "core"
     #   kind: "Secret"
     #   name: "mysecret"
+    # By default, Gateway would be created to the Namespace you are deploying Traefik to.
+    # You may create that Gateway in another namespace, setting its name below:
+    # namespace: default

 # Create an IngressRoute for the dashboard
 ingressRoute:
```

## 10.0.1  ![AppVersion: 2.4.9](https://img.shields.io/static/v1?label=AppVersion&message=2.4.9&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-07-14

* Add RBAC for middlewaretcps


## 10.0.0  ![AppVersion: 2.4.9](https://img.shields.io/static/v1?label=AppVersion&message=2.4.9&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-07-07

* Update CRD versions


## 9.20.1  ![AppVersion: 2.4.8](https://img.shields.io/static/v1?label=AppVersion&message=2.4.8&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-07-05

* Revert CRD templating


## 9.20.0  ![AppVersion: 2.4.8](https://img.shields.io/static/v1?label=AppVersion&message=2.4.8&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-07-05

* Add support for apiextensions v1 CRDs


## 9.19.2  ![AppVersion: 2.4.8](https://img.shields.io/static/v1?label=AppVersion&message=2.4.8&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-06-16

* Add name-metadata for service "List" object


## 9.19.1  ![AppVersion: 2.4.8](https://img.shields.io/static/v1?label=AppVersion&message=2.4.8&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-05-13

* fix simple typo

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index b30afac..9bf90ea 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -363,7 +363,7 @@ rbac:
   # If set to true, installs namespace-specific Role and RoleBinding and requires provider configuration be set to that same namespace
   namespaced: false

-# Enable to create a PodSecurityPolicy and assign it to the Service Account via RoleBindin or ClusterRoleBinding
+# Enable to create a PodSecurityPolicy and assign it to the Service Account via RoleBinding or ClusterRoleBinding
 podSecurityPolicy:
   enabled: false

```

## 9.19.0  ![AppVersion: 2.4.8](https://img.shields.io/static/v1?label=AppVersion&message=2.4.8&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-04-29

* Fix IngressClass api version

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 0aa2d6b..b30afac 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -63,6 +63,8 @@ ingressClass:
   # true is not unit-testable yet, pending https://github.com/rancher/helm-unittest/pull/12
   enabled: false
   isDefaultClass: false
+  # Use to force a networking.k8s.io API Version for certain CI/CD applications. E.g. "v1beta1"
+  fallbackApiVersion:

 # Activate Pilot integration
 pilot:
```

## 9.18.3  ![AppVersion: 2.4.8](https://img.shields.io/static/v1?label=AppVersion&message=2.4.8&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-04-26

* Fix: ignore provider namespace args on disabled


## 9.18.2  ![AppVersion: 2.4.8](https://img.shields.io/static/v1?label=AppVersion&message=2.4.8&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-04-02

* Fix pilot dashboard deactivation


## 9.18.1  ![AppVersion: 2.4.8](https://img.shields.io/static/v1?label=AppVersion&message=2.4.8&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-03-29

* Do not disable Traefik Pilot in the dashboard by default


## 9.18.0  ![AppVersion: 2.4.8](https://img.shields.io/static/v1?label=AppVersion&message=2.4.8&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-03-24

* Add an option to toggle the pilot dashboard

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 017f771..0aa2d6b 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -68,6 +68,8 @@ ingressClass:
 pilot:
   enabled: false
   token: ""
+  # Toggle Pilot Dashboard
+  # dashboard: false

 # Enable experimental features
 experimental:
```

## 9.17.6  ![AppVersion: 2.4.8](https://img.shields.io/static/v1?label=AppVersion&message=2.4.8&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-03-24

* Bump Traefik to 2.4.8


## 9.17.5  ![AppVersion: 2.4.7](https://img.shields.io/static/v1?label=AppVersion&message=2.4.7&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-03-17

* feat(labelSelector): option matching Ingresses based on labelSelectors

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 868a985..017f771 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -105,6 +105,7 @@ providers:
       # - "default"
   kubernetesIngress:
     enabled: true
+    # labelSelector: environment=production,method=traefik
     namespaces: []
       # - "default"
     # IP used for Kubernetes Ingress endpoints
```

## 9.17.4  ![AppVersion: 2.4.7](https://img.shields.io/static/v1?label=AppVersion&message=2.4.7&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-03-17

* Add helm resource-policy annotation on PVC


## 9.17.3  ![AppVersion: 2.4.7](https://img.shields.io/static/v1?label=AppVersion&message=2.4.7&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-03-17

* Throw error with explicit latest tag


## 9.17.2  ![AppVersion: 2.4.7](https://img.shields.io/static/v1?label=AppVersion&message=2.4.7&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-03-10

* fix(keywords): removed by mistake


## 9.17.1  ![AppVersion: 2.4.7](https://img.shields.io/static/v1?label=AppVersion&message=2.4.7&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-03-10

* feat(healthchecksPort): Support for overriding the liveness/readiness probes port

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 56abb93..868a985 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -120,6 +120,8 @@ providers:
 # After the volume has been mounted, add the configs into traefik by using the `additionalArguments` list below, eg:
 # additionalArguments:
 # - "--providers.file.filename=/config/dynamic.toml"
+# - "--ping"
+# - "--ping.entrypoint=web"
 volumes: []
 # - name: public-cert
 #   mountPath: "/certs"
@@ -225,6 +227,10 @@ ports:
     # only.
     # hostIP: 192.168.100.10

+    # Override the liveness/readiness port. This is useful to integrate traefik
+    # with an external Load Balancer that performs healthchecks.
+    # healthchecksPort: 9000
+
     # Defines whether the port is exposed if service.type is LoadBalancer or
     # NodePort.
     #
```

## 9.16.2  ![AppVersion: 2.4.7](https://img.shields.io/static/v1?label=AppVersion&message=2.4.7&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-03-09

* Bump Traefik to 2.4.7


## 9.16.1  ![AppVersion: 2.4.6](https://img.shields.io/static/v1?label=AppVersion&message=2.4.6&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-03-09

* Adding custom labels to deployment

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index ba24be7..56abb93 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -16,6 +16,8 @@ deployment:
   replicas: 1
   # Additional deployment annotations (e.g. for jaeger-operator sidecar injection)
   annotations: {}
+  # Additional deployment labels (e.g. for filtering deployment by custom labels)
+  labels: {}
   # Additional pod annotations (e.g. for mesh injection or prometheus scraping)
   podAnnotations: {}
   # Additional Pod labels (e.g. for filtering Pod by custom labels)
```

## 9.15.2  ![AppVersion: 2.4.6](https://img.shields.io/static/v1?label=AppVersion&message=2.4.6&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-03-02

* Upgrade Traefik to 2.4.6


## 9.15.1  ![AppVersion: 2.4.5](https://img.shields.io/static/v1?label=AppVersion&message=2.4.5&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-03-02

* Configurable PVC name

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 1e0e5a9..ba24be7 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -327,6 +327,7 @@ autoscaling:
 # It will persist TLS certificates.
 persistence:
   enabled: false
+  name: data
 #  existingClaim: ""
   accessMode: ReadWriteOnce
   size: 128Mi
```

## 9.14.4  ![AppVersion: 2.4.5](https://img.shields.io/static/v1?label=AppVersion&message=2.4.5&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-03-02

* fix typo


## 9.14.3  ![AppVersion: 2.4.5](https://img.shields.io/static/v1?label=AppVersion&message=2.4.5&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-02-19

* Bump Traefik to 2.4.5


## 9.14.2  ![AppVersion: 2.4.2](https://img.shields.io/static/v1?label=AppVersion&message=2.4.2&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-02-03

* docs: indent nit for dsdsocket example

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 56485ad..1e0e5a9 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -33,7 +33,7 @@ deployment:
   additionalVolumes: []
     # - name: dsdsocket
     #   hostPath:
-    #   path: /var/run/statsd-exporter
+    #     path: /var/run/statsd-exporter
   # Additional initContainers (e.g. for setting file permission as shown below)
   initContainers: []
     # The "volume-permissions" init container is required if you run into permission issues.
```

## 9.14.1  ![AppVersion: 2.4.2](https://img.shields.io/static/v1?label=AppVersion&message=2.4.2&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-02-03

* Update Traefik to 2.4.2


## 9.14.0  ![AppVersion: 2.4.0](https://img.shields.io/static/v1?label=AppVersion&message=2.4.0&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-02-01

* Enable Kubernetes Gateway provider with an experimental flag

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 50cab94..56485ad 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -71,6 +71,13 @@ pilot:
 experimental:
   plugins:
     enabled: false
+  kubernetesGateway:
+    enabled: false
+    appLabelSelector: "traefik"
+    certificates: []
+    # - group: "core"
+    #   kind: "Secret"
+    #   name: "mysecret"

 # Create an IngressRoute for the dashboard
 ingressRoute:
```

## 9.13.0  ![AppVersion: 2.4.0](https://img.shields.io/static/v1?label=AppVersion&message=2.4.0&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2021-01-22

* Update Traefik to 2.4 and add resources


## 9.12.3  ![AppVersion: 2.3.6](https://img.shields.io/static/v1?label=AppVersion&message=2.3.6&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-12-31

* Revert API Upgrade


## 9.12.2  ![AppVersion: 2.3.6](https://img.shields.io/static/v1?label=AppVersion&message=2.3.6&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-12-31

* Bump Traefik to 2.3.6


## 9.12.1  ![AppVersion: 2.3.3](https://img.shields.io/static/v1?label=AppVersion&message=2.3.3&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-12-30

* Resolve #303, change CRD version from v1beta1 to v1


## 9.12.0  ![AppVersion: 2.3.3](https://img.shields.io/static/v1?label=AppVersion&message=2.3.3&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-12-30

* Implement support for DaemonSet

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 60a721d..50cab94 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -10,7 +10,9 @@ image:
 #
 deployment:
   enabled: true
-  # Number of pods of the deployment
+  # Can be either Deployment or DaemonSet
+  kind: Deployment
+  # Number of pods of the deployment (only applies when kind == Deployment)
   replicas: 1
   # Additional deployment annotations (e.g. for jaeger-operator sidecar injection)
   annotations: {}
```

## 9.11.0  ![AppVersion: 2.3.3](https://img.shields.io/static/v1?label=AppVersion&message=2.3.3&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-11-20

* add podLabels - custom labels

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index a187df7..60a721d 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -16,6 +16,8 @@ deployment:
   annotations: {}
   # Additional pod annotations (e.g. for mesh injection or prometheus scraping)
   podAnnotations: {}
+  # Additional Pod labels (e.g. for filtering Pod by custom labels)
+  podLabels: {}
   # Additional containers (e.g. for metric offloading sidecars)
   additionalContainers: []
     # https://docs.datadoghq.com/developers/dogstatsd/unix_socket/?tab=host
```

## 9.10.2  ![AppVersion: 2.3.3](https://img.shields.io/static/v1?label=AppVersion&message=2.3.3&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-11-20

* Bump Traefik to 2.3.3


## 9.10.1  ![AppVersion: 2.3.1](https://img.shields.io/static/v1?label=AppVersion&message=2.3.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-11-04

* Specify IngressClass resource when checking for cluster capability


## 9.10.0  ![AppVersion: 2.3.1](https://img.shields.io/static/v1?label=AppVersion&message=2.3.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-11-03

* Add list of watched provider namespaces

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index e6b85ca..a187df7 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -88,8 +88,12 @@ rollingUpdate:
 providers:
   kubernetesCRD:
     enabled: true
+    namespaces: []
+      # - "default"
   kubernetesIngress:
     enabled: true
+    namespaces: []
+      # - "default"
     # IP used for Kubernetes Ingress endpoints
     publishedService:
       enabled: false
```

## 9.9.0  ![AppVersion: 2.3.1](https://img.shields.io/static/v1?label=AppVersion&message=2.3.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-11-03

* Add additionalVolumeMounts for traefik container

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 37dd151..e6b85ca 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -111,6 +111,12 @@ volumes: []
 #   mountPath: "/config"
 #   type: configMap

+# Additional volumeMounts to add to the Traefik container
+additionalVolumeMounts: []
+  # For instance when using a logshipper for access logs
+  # - name: traefik-logs
+  #   mountPath: /var/log/traefik
+
 # Logs
 # https://docs.traefik.io/observability/logs/
 logs:
```

## 9.8.4  ![AppVersion: 2.3.1](https://img.shields.io/static/v1?label=AppVersion&message=2.3.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-11-03

* fix: multiple ImagePullSecrets


## 9.8.3  ![AppVersion: 2.3.1](https://img.shields.io/static/v1?label=AppVersion&message=2.3.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-10-30

* Add imagePullSecrets

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 87f60c0..37dd151 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -42,6 +42,9 @@ deployment:
     #       mountPath: /data
   # Custom pod DNS policy. Apply if `hostNetwork: true`
   # dnsPolicy: ClusterFirstWithHostNet
+  # Additional imagePullSecrets
+  imagePullSecrets: []
+   # - name: myRegistryKeySecretName

 # Pod disruption budget
 podDisruptionBudget:
```

## 9.8.2  ![AppVersion: 2.3.1](https://img.shields.io/static/v1?label=AppVersion&message=2.3.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-10-28

* Add chart repo to source


## 9.8.1  ![AppVersion: 2.3.1](https://img.shields.io/static/v1?label=AppVersion&message=2.3.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-10-23

* fix semver compare

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 4ca1f8f..87f60c0 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -1,7 +1,8 @@
 # Default values for Traefik
 image:
   name: traefik
-  tag: 2.3.1
+  # defaults to appVersion
+  tag: ""
   pullPolicy: IfNotPresent

 #
```

## 9.8.0  ![AppVersion: 2.3.1](https://img.shields.io/static/v1?label=AppVersion&message=2.3.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-10-20

* feat: Enable entrypoint tls config + TLSOption

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index eee3622..4ca1f8f 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -231,6 +231,31 @@ ports:
     # The port protocol (TCP/UDP)
     protocol: TCP
     # nodePort: 32443
+    # Set TLS at the entrypoint
+    # https://doc.traefik.io/traefik/routing/entrypoints/#tls
+    tls:
+      enabled: false
+      # this is the name of a TLSOption definition
+      options: ""
+      certResolver: ""
+      domains: []
+      # - main: example.com
+      #   sans:
+      #     - foo.example.com
+      #     - bar.example.com
+
+# TLS Options are created as TLSOption CRDs
+# https://doc.traefik.io/traefik/https/tls/#tls-options
+# Example:
+# tlsOptions:
+#   default:
+#     sniStrict: true
+#     preferServerCipherSuites: true
+#   foobar:
+#     curvePreferences:
+#       - CurveP521
+#       - CurveP384
+tlsOptions: {}

 # Options for the main traefik service, where the entrypoints traffic comes
 # from.
```

## 9.7.0  ![AppVersion: 2.3.1](https://img.shields.io/static/v1?label=AppVersion&message=2.3.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-10-15

* Add a configuration option for an emptyDir as plugin storage

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index b7153a1..eee3622 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -54,10 +54,16 @@ ingressClass:
   enabled: false
   isDefaultClass: false

+# Activate Pilot integration
 pilot:
   enabled: false
   token: ""

+# Enable experimental features
+experimental:
+  plugins:
+    enabled: false
+
 # Create an IngressRoute for the dashboard
 ingressRoute:
   dashboard:
```

## 9.6.0  ![AppVersion: 2.3.1](https://img.shields.io/static/v1?label=AppVersion&message=2.3.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-10-15

* Add additional volumes for init and additional containers

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 9bac45e..b7153a1 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -17,6 +17,18 @@ deployment:
   podAnnotations: {}
   # Additional containers (e.g. for metric offloading sidecars)
   additionalContainers: []
+    # https://docs.datadoghq.com/developers/dogstatsd/unix_socket/?tab=host
+    # - name: socat-proxy
+    # image: alpine/socat:1.0.5
+    # args: ["-s", "-u", "udp-recv:8125", "unix-sendto:/socket/socket"]
+    # volumeMounts:
+    #   - name: dsdsocket
+    #     mountPath: /socket
+  # Additional volumes available for use with initContainers and additionalContainers
+  additionalVolumes: []
+    # - name: dsdsocket
+    #   hostPath:
+    #   path: /var/run/statsd-exporter
   # Additional initContainers (e.g. for setting file permission as shown below)
   initContainers: []
     # The "volume-permissions" init container is required if you run into permission issues.
```

## 9.5.2  ![AppVersion: 2.3.1](https://img.shields.io/static/v1?label=AppVersion&message=2.3.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-10-15

* Replace extensions with policy because of deprecation


## 9.5.1  ![AppVersion: 2.3.1](https://img.shields.io/static/v1?label=AppVersion&message=2.3.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-10-14

* Template custom volume name

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 5a8d8ea..9bac45e 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -76,7 +76,7 @@ providers:
       # pathOverride: ""

 #
-# Add volumes to the traefik pod.
+# Add volumes to the traefik pod. The volume name will be passed to tpl.
 # This can be used to mount a cert pair or a configmap that holds a config.toml file.
 # After the volume has been mounted, add the configs into traefik by using the `additionalArguments` list below, eg:
 # additionalArguments:
@@ -85,7 +85,7 @@ volumes: []
 # - name: public-cert
 #   mountPath: "/certs"
 #   type: secret
-# - name: configs
+# - name: '{{ printf "%s-configs" .Release.Name }}'
 #   mountPath: "/config"
 #   type: configMap

```

## 9.5.0  ![AppVersion: 2.3.1](https://img.shields.io/static/v1?label=AppVersion&message=2.3.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-10-02

* Create PodSecurityPolicy and RBAC when needed.

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 8c4d866..5a8d8ea 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -281,6 +281,10 @@ rbac:
   # If set to true, installs namespace-specific Role and RoleBinding and requires provider configuration be set to that same namespace
   namespaced: false

+# Enable to create a PodSecurityPolicy and assign it to the Service Account via RoleBindin or ClusterRoleBinding
+podSecurityPolicy:
+  enabled: false
+
 # The service account the pods will use to interact with the Kubernetes API
 serviceAccount:
   # If set, an existing service account is used
```

## 9.4.3  ![AppVersion: 2.3.1](https://img.shields.io/static/v1?label=AppVersion&message=2.3.1&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-10-02

* Update traefik to v2.3.1

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 3df75a4..8c4d866 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -1,7 +1,7 @@
 # Default values for Traefik
 image:
   name: traefik
-  tag: 2.3.0
+  tag: 2.3.1
   pullPolicy: IfNotPresent

 #
```

## 9.4.2  ![AppVersion: 2.3.0](https://img.shields.io/static/v1?label=AppVersion&message=2.3.0&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-10-02

* Add Artifact Hub repository metadata file


## 9.4.1  ![AppVersion: 2.3.0](https://img.shields.io/static/v1?label=AppVersion&message=2.3.0&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-10-01

* Fix broken chart icon url


## 9.4.0  ![AppVersion: 2.3.0](https://img.shields.io/static/v1?label=AppVersion&message=2.3.0&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-10-01

* Allow to specify custom labels on Service

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index a6175ff..3df75a4 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -221,6 +221,8 @@ service:
   type: LoadBalancer
   # Additional annotations (e.g. for cloud provider specific config)
   annotations: {}
+  # Additional service labels (e.g. for filtering Service by custom labels)
+  labels: {}
   # Additional entries here will be added to the service spec. Cannot contains
   # type, selector or ports entries.
   spec: {}
```

## 9.3.0  ![AppVersion: 2.3.0](https://img.shields.io/static/v1?label=AppVersion&message=2.3.0&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-09-24

* Release Traefik 2.3

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index fba955d..a6175ff 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -1,7 +1,7 @@
 # Default values for Traefik
 image:
   name: traefik
-  tag: 2.2.8
+  tag: 2.3.0
   pullPolicy: IfNotPresent

 #
@@ -36,6 +36,16 @@ podDisruptionBudget:
   # maxUnavailable: 1
   # minAvailable: 0

+# Use ingressClass. Ignored if Traefik version < 2.3 / kubernetes < 1.18.x
+ingressClass:
+  # true is not unit-testable yet, pending https://github.com/rancher/helm-unittest/pull/12
+  enabled: false
+  isDefaultClass: false
+
+pilot:
+  enabled: false
+  token: ""
+
 # Create an IngressRoute for the dashboard
 ingressRoute:
   dashboard:
```

## 9.2.1  ![AppVersion: 2.2.8](https://img.shields.io/static/v1?label=AppVersion&message=2.2.8&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-09-18

* Add new helm url


## 9.2.0  ![AppVersion: 2.2.8](https://img.shields.io/static/v1?label=AppVersion&message=2.2.8&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-09-16

* chore: move to new organization.

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 9f52c39..fba955d 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -20,7 +20,7 @@ deployment:
   # Additional initContainers (e.g. for setting file permission as shown below)
   initContainers: []
     # The "volume-permissions" init container is required if you run into permission issues.
-    # Related issue: https://github.com/containous/traefik/issues/6972
+    # Related issue: https://github.com/traefik/traefik/issues/6972
     # - name: volume-permissions
     #   image: busybox:1.31.1
     #   command: ["sh", "-c", "chmod -Rv 600 /data/*"]
```

## 9.1.1  ![AppVersion: 2.2.8](https://img.shields.io/static/v1?label=AppVersion&message=2.2.8&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-09-04

* Update reference to using kubectl proxy to kubectl port-forward

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 7b74a39..9f52c39 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -175,7 +175,7 @@ ports:
     #
     # You SHOULD NOT expose the traefik port on production deployments.
     # If you want to access it from outside of your cluster,
-    # use `kubectl proxy` or create a secure ingress
+    # use `kubectl port-forward` or create a secure ingress
     expose: false
     # The exposed port for this service
     exposedPort: 9000
```

## 9.1.0  ![AppVersion: 2.2.8](https://img.shields.io/static/v1?label=AppVersion&message=2.2.8&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-08-24

* PublishedService option

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index e161a14..7b74a39 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -58,6 +58,12 @@ providers:
     enabled: true
   kubernetesIngress:
     enabled: true
+    # IP used for Kubernetes Ingress endpoints
+    publishedService:
+      enabled: false
+      # Published Kubernetes Service to copy status from. Format: namespace/servicename
+      # By default this Traefik service
+      # pathOverride: ""

 #
 # Add volumes to the traefik pod.
```

## 9.0.0  ![AppVersion: 2.2.8](https://img.shields.io/static/v1?label=AppVersion&message=2.2.8&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-08-21

* feat: Move Chart apiVersion: v2


## 8.13.3  ![AppVersion: 2.2.8](https://img.shields.io/static/v1?label=AppVersion&message=2.2.8&color=success&logo=) ![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-08-21

* bug: Check for port config


## 8.13.2  ![AppVersion: 2.2.8](https://img.shields.io/static/v1?label=AppVersion&message=2.2.8&color=success&logo=) ![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-08-19

* Fix log level configuration


## 8.13.1  ![AppVersion: 2.2.8](https://img.shields.io/static/v1?label=AppVersion&message=2.2.8&color=success&logo=) ![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-08-18

* Dont redirect to websecure by default

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 67276f7..e161a14 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -188,7 +188,7 @@ ports:
     # Port Redirections
     # Added in 2.2, you can make permanent redirects via entrypoints.
     # https://docs.traefik.io/routing/entrypoints/#redirection
-    redirectTo: websecure
+    # redirectTo: websecure
   websecure:
     port: 8443
     # hostPort: 8443
```

## 8.13.0  ![AppVersion: 2.2.8](https://img.shields.io/static/v1?label=AppVersion&message=2.2.8&color=success&logo=) ![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-08-18

* Add logging, and http redirect config

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 6f79580..67276f7 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -73,6 +73,48 @@ volumes: []
 #   mountPath: "/config"
 #   type: configMap

+# Logs
+# https://docs.traefik.io/observability/logs/
+logs:
+  # Traefik logs concern everything that happens to Traefik itself (startup, configuration, events, shutdown, and so on).
+  general:
+    # By default, the logs use a text format (common), but you can
+    # also ask for the json format in the format option
+    # format: json
+    # By default, the level is set to ERROR. Alternative logging levels are DEBUG, PANIC, FATAL, ERROR, WARN, and INFO.
+    level: ERROR
+  access:
+    # To enable access logs
+    enabled: false
+    # By default, logs are written using the Common Log Format (CLF).
+    # To write logs in JSON, use json in the format option.
+    # If the given format is unsupported, the default (CLF) is used instead.
+    # format: json
+    # To write the logs in an asynchronous fashion, specify a bufferingSize option.
+    # This option represents the number of log lines Traefik will keep in memory before writing
+    # them to the selected output. In some cases, this option can greatly help performances.
+    # bufferingSize: 100
+    # Filtering https://docs.traefik.io/observability/access-logs/#filtering
+    filters: {}
+      # statuscodes: "200,300-302"
+      # retryattempts: true
+      # minduration: 10ms
+    # Fields
+    # https://docs.traefik.io/observability/access-logs/#limiting-the-fieldsincluding-headers
+    fields:
+      general:
+        defaultmode: keep
+        names: {}
+          # Examples:
+          # ClientUsername: drop
+      headers:
+        defaultmode: drop
+        names: {}
+          # Examples:
+          # User-Agent: redact
+          # Authorization: drop
+          # Content-Type: keep
+
 globalArguments:
   - "--global.checknewversion"
   - "--global.sendanonymoususage"
@@ -143,6 +185,10 @@ ports:
     # Use nodeport if set. This is useful if you have configured Traefik in a
     # LoadBalancer
     # nodePort: 32080
+    # Port Redirections
+    # Added in 2.2, you can make permanent redirects via entrypoints.
+    # https://docs.traefik.io/routing/entrypoints/#redirection
+    redirectTo: websecure
   websecure:
     port: 8443
     # hostPort: 8443
```

## 8.12.0  ![AppVersion: 2.2.8](https://img.shields.io/static/v1?label=AppVersion&message=2.2.8&color=success&logo=) ![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-08-14

* Add image pull policy

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 10b3949..6f79580 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -2,6 +2,7 @@
 image:
   name: traefik
   tag: 2.2.8
+  pullPolicy: IfNotPresent

 #
 # Configure the deployment
```

## 8.11.0  ![AppVersion: 2.2.8](https://img.shields.io/static/v1?label=AppVersion&message=2.2.8&color=success&logo=) ![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-08-12

* Add dns policy option

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 80ddaaa..10b3949 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -26,6 +26,8 @@ deployment:
     #   volumeMounts:
     #     - name: data
     #       mountPath: /data
+  # Custom pod DNS policy. Apply if `hostNetwork: true`
+  # dnsPolicy: ClusterFirstWithHostNet

 # Pod disruption budget
 podDisruptionBudget:
```

## 8.10.0  ![AppVersion: 2.2.8](https://img.shields.io/static/v1?label=AppVersion&message=2.2.8&color=success&logo=) ![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-08-11

* Add hostIp to port configuration

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 936ab92..80ddaaa 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -112,6 +112,12 @@ ports:
     port: 9000
     # Use hostPort if set.
     # hostPort: 9000
+    #
+    # Use hostIP if set. If not set, Kubernetes will default to 0.0.0.0, which
+    # means it's listening on all your interfaces and all your IPs. You may want
+    # to set this value if you need traefik to listen on specific interface
+    # only.
+    # hostIP: 192.168.100.10

     # Defines whether the port is exposed if service.type is LoadBalancer or
     # NodePort.
```

## 8.9.2  ![AppVersion: 2.2.8](https://img.shields.io/static/v1?label=AppVersion&message=2.2.8&color=success&logo=) ![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-08-10

* Bump Traefik to 2.2.8

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 42ee893..936ab92 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -1,7 +1,7 @@
 # Default values for Traefik
 image:
   name: traefik
-  tag: 2.2.5
+  tag: 2.2.8

 #
 # Configure the deployment
```

## 8.9.1  ![AppVersion: 2.2.5](https://img.shields.io/static/v1?label=AppVersion&message=2.2.5&color=success&logo=) ![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-07-15

* Upgrade traefik version

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index a7fb668..42ee893 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -1,7 +1,7 @@
 # Default values for Traefik
 image:
   name: traefik
-  tag: 2.2.1
+  tag: 2.2.5

 #
 # Configure the deployment
```

## 8.9.0  ![AppVersion: 2.2.1](https://img.shields.io/static/v1?label=AppVersion&message=2.2.1&color=success&logo=) ![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-07-08

* run init container to set proper permissions on volume

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 62e3a77..a7fb668 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -16,6 +16,16 @@ deployment:
   podAnnotations: {}
   # Additional containers (e.g. for metric offloading sidecars)
   additionalContainers: []
+  # Additional initContainers (e.g. for setting file permission as shown below)
+  initContainers: []
+    # The "volume-permissions" init container is required if you run into permission issues.
+    # Related issue: https://github.com/containous/traefik/issues/6972
+    # - name: volume-permissions
+    #   image: busybox:1.31.1
+    #   command: ["sh", "-c", "chmod -Rv 600 /data/*"]
+    #   volumeMounts:
+    #     - name: data
+    #       mountPath: /data

 # Pod disruption budget
 podDisruptionBudget:
```

## 8.8.1  ![AppVersion: 2.2.1](https://img.shields.io/static/v1?label=AppVersion&message=2.2.1&color=success&logo=) ![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-07-02

* Additional container fix

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 85df29c..62e3a77 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -15,7 +15,7 @@ deployment:
   # Additional pod annotations (e.g. for mesh injection or prometheus scraping)
   podAnnotations: {}
   # Additional containers (e.g. for metric offloading sidecars)
-  additionalContainers: {}
+  additionalContainers: []

 # Pod disruption budget
 podDisruptionBudget:
```

## 8.8.0  ![AppVersion: 2.2.1](https://img.shields.io/static/v1?label=AppVersion&message=2.2.1&color=success&logo=) ![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-07-01

* added additionalContainers option to chart

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 6a9dfd8..85df29c 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -14,6 +14,8 @@ deployment:
   annotations: {}
   # Additional pod annotations (e.g. for mesh injection or prometheus scraping)
   podAnnotations: {}
+  # Additional containers (e.g. for metric offloading sidecars)
+  additionalContainers: {}

 # Pod disruption budget
 podDisruptionBudget:
```

## 8.7.2  ![AppVersion: 2.2.1](https://img.shields.io/static/v1?label=AppVersion&message=2.2.1&color=success&logo=) ![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-06-30

* Update image


## 8.7.1  ![AppVersion: 2.2.1](https://img.shields.io/static/v1?label=AppVersion&message=2.2.1&color=success&logo=) ![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-06-26

* Update values.yaml

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 05f9eab..6a9dfd8 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -196,7 +196,7 @@ rbac:
   # If set to true, installs namespace-specific Role and RoleBinding and requires provider configuration be set to that same namespace
   namespaced: false

-# The service account the pods will use to interact with the Kubernates API
+# The service account the pods will use to interact with the Kubernetes API
 serviceAccount:
   # If set, an existing service account is used
   # If not set, a service account is created automatically using the fullname template
```

## 8.7.0  ![AppVersion: 2.2.1](https://img.shields.io/static/v1?label=AppVersion&message=2.2.1&color=success&logo=) ![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-06-23

* Add option to disable providers

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 102ae00..05f9eab 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -34,6 +34,16 @@ rollingUpdate:
   maxUnavailable: 1
   maxSurge: 1

+
+#
+# Configure providers
+#
+providers:
+  kubernetesCRD:
+    enabled: true
+  kubernetesIngress:
+    enabled: true
+
 #
 # Add volumes to the traefik pod.
 # This can be used to mount a cert pair or a configmap that holds a config.toml file.
```

## 8.6.1  ![AppVersion: 2.2.1](https://img.shields.io/static/v1?label=AppVersion&message=2.2.1&color=success&logo=) ![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-06-18

* Fix read-only /tmp


## 8.6.0  ![AppVersion: 2.2.1](https://img.shields.io/static/v1?label=AppVersion&message=2.2.1&color=success&logo=) ![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-06-17

* Add existing PVC support(#158)

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index b2f4fc3..102ae00 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -164,6 +164,7 @@ autoscaling:
 # It will persist TLS certificates.
 persistence:
   enabled: false
+#  existingClaim: ""
   accessMode: ReadWriteOnce
   size: 128Mi
   # storageClass: ""
```

## 8.5.0  ![AppVersion: 2.2.1](https://img.shields.io/static/v1?label=AppVersion&message=2.2.1&color=success&logo=) ![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-06-16

* UDP support

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 9a9b668..b2f4fc3 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -100,11 +100,15 @@ ports:
     expose: false
     # The exposed port for this service
     exposedPort: 9000
+    # The port protocol (TCP/UDP)
+    protocol: TCP
   web:
     port: 8000
     # hostPort: 8000
     expose: true
     exposedPort: 80
+    # The port protocol (TCP/UDP)
+    protocol: TCP
     # Use nodeport if set. This is useful if you have configured Traefik in a
     # LoadBalancer
     # nodePort: 32080
@@ -113,6 +117,8 @@ ports:
     # hostPort: 8443
     expose: true
     exposedPort: 443
+    # The port protocol (TCP/UDP)
+    protocol: TCP
     # nodePort: 32443

 # Options for the main traefik service, where the entrypoints traffic comes
```

## 8.4.1  ![AppVersion: 2.2.1](https://img.shields.io/static/v1?label=AppVersion&message=2.2.1&color=success&logo=) ![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-06-10

* Fix PDB with minAvailable set

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index e812b98..9a9b668 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -18,7 +18,7 @@ deployment:
 # Pod disruption budget
 podDisruptionBudget:
   enabled: false
-  maxUnavailable: 1
+  # maxUnavailable: 1
   # minAvailable: 0

 # Create an IngressRoute for the dashboard
```

## 8.4.0  ![AppVersion: 2.2.1](https://img.shields.io/static/v1?label=AppVersion&message=2.2.1&color=success&logo=) ![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-06-09

* Add pod disruption budget (#192)

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 5f44e5c..e812b98 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -15,6 +15,12 @@ deployment:
   # Additional pod annotations (e.g. for mesh injection or prometheus scraping)
   podAnnotations: {}

+# Pod disruption budget
+podDisruptionBudget:
+  enabled: false
+  maxUnavailable: 1
+  # minAvailable: 0
+
 # Create an IngressRoute for the dashboard
 ingressRoute:
   dashboard:
```

## 8.3.0  ![AppVersion: 2.2.1](https://img.shields.io/static/v1?label=AppVersion&message=2.2.1&color=success&logo=) ![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-06-08

* Add option to disable RBAC and ServiceAccount

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 96bba18..5f44e5c 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -165,6 +165,20 @@ persistence:
 # affinity is left as default.
 hostNetwork: false

+# Whether Role Based Access Control objects like roles and rolebindings should be created
+rbac:
+  enabled: true
+
+  # If set to false, installs ClusterRole and ClusterRoleBinding so Traefik can be used across namespaces.
+  # If set to true, installs namespace-specific Role and RoleBinding and requires provider configuration be set to that same namespace
+  namespaced: false
+
+# The service account the pods will use to interact with the Kubernates API
+serviceAccount:
+  # If set, an existing service account is used
+  # If not set, a service account is created automatically using the fullname template
+  name: ""
+
 # Additional serviceAccount annotations (e.g. for oidc authentication)
 serviceAccountAnnotations: {}

```

## 8.2.1  ![AppVersion: 2.2.1](https://img.shields.io/static/v1?label=AppVersion&message=2.2.1&color=success&logo=) ![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-05-25

* Remove suggested providers.kubernetesingress value

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index e35bdf9..96bba18 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -50,9 +50,9 @@ globalArguments:
 # Configure Traefik static configuration
 # Additional arguments to be passed at Traefik's binary
 # All available options available on https://docs.traefik.io/reference/static-configuration/cli/
-## Use curly braces to pass values: `helm install --set="additionalArguments={--providers.kubernetesingress,--log.level=DEBUG}"`
+## Use curly braces to pass values: `helm install --set="additionalArguments={--providers.kubernetesingress.ingressclass=traefik-internal,--log.level=DEBUG}"`
 additionalArguments: []
-#  - "--providers.kubernetesingress"
+#  - "--providers.kubernetesingress.ingressclass=traefik-internal"
 #  - "--log.level=DEBUG"

 # Environment variables to be passed to Traefik's binary
```

## 8.2.0  ![AppVersion: 2.2.1](https://img.shields.io/static/v1?label=AppVersion&message=2.2.1&color=success&logo=) ![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-05-18

* Add kubernetes ingress by default


## 8.1.5  ![AppVersion: 2.2.1](https://img.shields.io/static/v1?label=AppVersion&message=2.2.1&color=success&logo=) ![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-05-18

* Fix example log params in values.yml

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index abe2334..e35bdf9 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -50,10 +50,10 @@ globalArguments:
 # Configure Traefik static configuration
 # Additional arguments to be passed at Traefik's binary
 # All available options available on https://docs.traefik.io/reference/static-configuration/cli/
-## Use curly braces to pass values: `helm install --set="additionalArguments={--providers.kubernetesingress,--logs.level=DEBUG}"`
+## Use curly braces to pass values: `helm install --set="additionalArguments={--providers.kubernetesingress,--log.level=DEBUG}"`
 additionalArguments: []
 #  - "--providers.kubernetesingress"
-#  - "--logs.level=DEBUG"
+#  - "--log.level=DEBUG"

 # Environment variables to be passed to Traefik's binary
 env: []
```

## 8.1.4  ![AppVersion: 2.2.1](https://img.shields.io/static/v1?label=AppVersion&message=2.2.1&color=success&logo=) ![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-04-30

* Update Traefik to v2.2.1

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 57cc7e1..abe2334 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -1,7 +1,7 @@
 # Default values for Traefik
 image:
   name: traefik
-  tag: 2.2.0
+  tag: 2.2.1

 #
 # Configure the deployment
```

## 8.1.3  ![AppVersion: 2.2.0](https://img.shields.io/static/v1?label=AppVersion&message=2.2.0&color=success&logo=) ![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-04-29

* Clarify additionnal arguments log

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index d639f72..57cc7e1 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -50,9 +50,10 @@ globalArguments:
 # Configure Traefik static configuration
 # Additional arguments to be passed at Traefik's binary
 # All available options available on https://docs.traefik.io/reference/static-configuration/cli/
-## Use curly braces to pass values: `helm install --set="additionalArguments={--providers.kubernetesingress,--global.checknewversion=true}"`
+## Use curly braces to pass values: `helm install --set="additionalArguments={--providers.kubernetesingress,--logs.level=DEBUG}"`
 additionalArguments: []
 #  - "--providers.kubernetesingress"
+#  - "--logs.level=DEBUG"

 # Environment variables to be passed to Traefik's binary
 env: []
```

## 8.1.2  ![AppVersion: 2.2.0](https://img.shields.io/static/v1?label=AppVersion&message=2.2.0&color=success&logo=) ![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-04-23

* Remove invalid flags. (#161)

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 0e7aaef..d639f72 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -34,8 +34,6 @@ rollingUpdate:
 # After the volume has been mounted, add the configs into traefik by using the `additionalArguments` list below, eg:
 # additionalArguments:
 # - "--providers.file.filename=/config/dynamic.toml"
-# - "--tls.certificates.certFile=/certs/tls.crt"
-# - "--tls.certificates.keyFile=/certs/tls.key"
 volumes: []
 # - name: public-cert
 #   mountPath: "/certs"
```

## 8.1.1  ![AppVersion: 2.2.0](https://img.shields.io/static/v1?label=AppVersion&message=2.2.0&color=success&logo=) ![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-04-23

* clarify project philosophy and guidelines


## 8.1.0  ![AppVersion: 2.2.0](https://img.shields.io/static/v1?label=AppVersion&message=2.2.0&color=success&logo=) ![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-04-22

* Add priorityClassName & securityContext

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index d55a40a..0e7aaef 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -191,3 +191,20 @@ affinity: {}
 #       topologyKey: failure-domain.beta.kubernetes.io/zone
 nodeSelector: {}
 tolerations: []
+
+# Pods can have priority.
+# Priority indicates the importance of a Pod relative to other Pods.
+priorityClassName: ""
+
+# Set the container security context
+# To run the container with ports below 1024 this will need to be adjust to run as root
+securityContext:
+  capabilities:
+    drop: [ALL]
+  readOnlyRootFilesystem: true
+  runAsGroup: 65532
+  runAsNonRoot: true
+  runAsUser: 65532
+
+podSecurityContext:
+  fsGroup: 65532
```

## 8.0.4  ![AppVersion: 2.2.0](https://img.shields.io/static/v1?label=AppVersion&message=2.2.0&color=success&logo=) ![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-04-20

* Possibility to bind environment variables via envFrom

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 7f8092e..d55a40a 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -71,6 +71,12 @@ env: []
 #       name: secret-name
 #       key: secret-key

+envFrom: []
+# - configMapRef:
+#     name: config-map-name
+# - secretRef:
+#     name: secret-name
+
 # Configure ports
 ports:
   # The name of this one can't be changed as it is used for the readiness and
```

## 8.0.3  ![AppVersion: 2.2.0](https://img.shields.io/static/v1?label=AppVersion&message=2.2.0&color=success&logo=) ![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-04-15

* Add support for data volume subPath. (#147)

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 152339b..7f8092e 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -152,6 +152,7 @@ persistence:
   # storageClass: ""
   path: /data
   annotations: {}
+  # subPath: "" # only mount a subpath of the Volume into the pod

 # If hostNetwork is true, runs traefik in the host network namespace
 # To prevent unschedulabel pods due to port collisions, if hostNetwork=true
```

## 8.0.2  ![AppVersion: 2.2.0](https://img.shields.io/static/v1?label=AppVersion&message=2.2.0&color=success&logo=) ![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-04-10

* Ability to add custom labels to dashboard's IngressRoute

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 5d294b7..152339b 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -21,6 +21,8 @@ ingressRoute:
     enabled: true
     # Additional ingressRoute annotations (e.g. for kubernetes.io/ingress.class)
     annotations: {}
+    # Additional ingressRoute labels (e.g. for filtering IngressRoute by custom labels)
+    labels: {}

 rollingUpdate:
   maxUnavailable: 1
```

## 8.0.1  ![AppVersion: 2.2.0](https://img.shields.io/static/v1?label=AppVersion&message=2.2.0&color=success&logo=) ![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-04-10

* rbac does not need "pods" per documentation


## 8.0.0  ![AppVersion: 2.2.0](https://img.shields.io/static/v1?label=AppVersion&message=2.2.0&color=success&logo=) ![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-04-07

* follow helm best practices

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index e61a9fd..5d294b7 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -10,7 +10,7 @@ deployment:
   enabled: true
   # Number of pods of the deployment
   replicas: 1
-  # Addtional deployment annotations (e.g. for jaeger-operator sidecar injection)
+  # Additional deployment annotations (e.g. for jaeger-operator sidecar injection)
   annotations: {}
   # Additional pod annotations (e.g. for mesh injection or prometheus scraping)
   podAnnotations: {}
@@ -19,7 +19,7 @@ deployment:
 ingressRoute:
   dashboard:
     enabled: true
-    # Addtional ingressRoute annotations (e.g. for kubernetes.io/ingress.class)
+    # Additional ingressRoute annotations (e.g. for kubernetes.io/ingress.class)
     annotations: {}

 rollingUpdate:
```

## 7.2.1  ![AppVersion: 2.2.0](https://img.shields.io/static/v1?label=AppVersion&message=2.2.0&color=success&logo=) ![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-04-07

* add annotations to ingressRoute

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 15d1c25..e61a9fd 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -19,6 +19,8 @@ deployment:
 ingressRoute:
   dashboard:
     enabled: true
+    # Addtional ingressRoute annotations (e.g. for kubernetes.io/ingress.class)
+    annotations: {}

 rollingUpdate:
   maxUnavailable: 1
```

## 7.2.0  ![AppVersion: 2.2.0](https://img.shields.io/static/v1?label=AppVersion&message=2.2.0&color=success&logo=) ![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-04-03

* Add support for helm 2


## 7.1.0  ![AppVersion: 2.2.0](https://img.shields.io/static/v1?label=AppVersion&message=2.2.0&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-03-31

* Add support for externalIPs

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 6d6d13f..15d1c25 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -116,6 +116,8 @@ service:
   loadBalancerSourceRanges: []
     # - 192.168.0.1/32
     # - 172.16.0.0/16
+  externalIPs: []
+    # - 1.2.3.4

 ## Create HorizontalPodAutoscaler object.
 ##
```

## 7.0.0  ![AppVersion: 2.2.0](https://img.shields.io/static/v1?label=AppVersion&message=2.2.0&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-03-27

* Remove secretsEnv value key

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 1ac720d..6d6d13f 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -52,18 +52,20 @@ globalArguments:
 additionalArguments: []
 #  - "--providers.kubernetesingress"

-# Secret to be set as environment variables to be passed to Traefik's binary
-secretEnv: []
-  # - name: SOME_VAR
-  #   secretName: my-secret-name
-  #   secretKey: my-secret-key
-
 # Environment variables to be passed to Traefik's binary
 env: []
-  # - name: SOME_VAR
-  #   value: some-var-value
-  # - name: SOME_OTHER_VAR
-  #   value: some-other-var-value
+# - name: SOME_VAR
+#   value: some-var-value
+# - name: SOME_VAR_FROM_CONFIG_MAP
+#   valueFrom:
+#     configMapRef:
+#       name: configmap-name
+#       key: config-key
+# - name: SOME_SECRET
+#   valueFrom:
+#     secretKeyRef:
+#       name: secret-name
+#       key: secret-key

 # Configure ports
 ports:
```

## 6.4.0  ![AppVersion: 2.2.0](https://img.shields.io/static/v1?label=AppVersion&message=2.2.0&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-03-27

* Add ability to set serviceAccount annotations

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 85abe42..1ac720d 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -151,6 +151,9 @@ persistence:
 # affinity is left as default.
 hostNetwork: false

+# Additional serviceAccount annotations (e.g. for oidc authentication)
+serviceAccountAnnotations: {}
+
 resources: {}
   # requests:
   #   cpu: "100m"
```

## 6.3.0  ![AppVersion: 2.2.0](https://img.shields.io/static/v1?label=AppVersion&message=2.2.0&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-03-27

* hpa

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 2f5d132..85abe42 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -115,6 +115,22 @@ service:
     # - 192.168.0.1/32
     # - 172.16.0.0/16

+## Create HorizontalPodAutoscaler object.
+##
+autoscaling:
+  enabled: false
+#   minReplicas: 1
+#   maxReplicas: 10
+#   metrics:
+#   - type: Resource
+#     resource:
+#       name: cpu
+#       targetAverageUtilization: 60
+#   - type: Resource
+#     resource:
+#       name: memory
+#       targetAverageUtilization: 60
+
 # Enable persistence using Persistent Volume Claims
 # ref: http://kubernetes.io/docs/user-guide/persistent-volumes/
 # After the pvc has been mounted, add the configs into traefik by using the `additionalArguments` list below, eg:
```

## 6.2.0  ![AppVersion: 2.2.0](https://img.shields.io/static/v1?label=AppVersion&message=2.2.0&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-03-26

* Update to v2.2 (#96)

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index ebd2fde..2f5d132 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -1,7 +1,7 @@
 # Default values for Traefik
 image:
   name: traefik
-  tag: 2.1.8
+  tag: 2.2.0

 #
 # Configure the deployment
```

## 6.1.2  ![AppVersion: 2.1.8](https://img.shields.io/static/v1?label=AppVersion&message=2.1.8&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-03-20

* Upgrade traefik version

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 65c7665..ebd2fde 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -1,7 +1,7 @@
 # Default values for Traefik
 image:
   name: traefik
-  tag: 2.1.4
+  tag: 2.1.8

 #
 # Configure the deployment
```

## 6.1.1  ![AppVersion: 2.1.4](https://img.shields.io/static/v1?label=AppVersion&message=2.1.4&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-03-20

* Upgrade traefik version

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 89c7ac1..65c7665 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -1,7 +1,7 @@
 # Default values for Traefik
 image:
   name: traefik
-  tag: 2.1.3
+  tag: 2.1.4

 #
 # Configure the deployment
```

## 6.1.0  ![AppVersion: 2.1.3](https://img.shields.io/static/v1?label=AppVersion&message=2.1.3&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-03-20

* Add ability to add annotations to deployment

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 8d66111..89c7ac1 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -10,6 +10,8 @@ deployment:
   enabled: true
   # Number of pods of the deployment
   replicas: 1
+  # Addtional deployment annotations (e.g. for jaeger-operator sidecar injection)
+  annotations: {}
   # Additional pod annotations (e.g. for mesh injection or prometheus scraping)
   podAnnotations: {}

```

## 6.0.2  ![AppVersion: 2.1.3](https://img.shields.io/static/v1?label=AppVersion&message=2.1.3&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-03-16

* Correct storage class key name


## 6.0.1  ![AppVersion: 2.1.3](https://img.shields.io/static/v1?label=AppVersion&message=2.1.3&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-03-16

* Change default values of arrays from objects to actual arrays

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 490b2b6..8d66111 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -51,13 +51,13 @@ additionalArguments: []
 #  - "--providers.kubernetesingress"

 # Secret to be set as environment variables to be passed to Traefik's binary
-secretEnv: {}
+secretEnv: []
   # - name: SOME_VAR
   #   secretName: my-secret-name
   #   secretKey: my-secret-key

 # Environment variables to be passed to Traefik's binary
-env: {}
+env: []
   # - name: SOME_VAR
   #   value: some-var-value
   # - name: SOME_OTHER_VAR
@@ -109,7 +109,7 @@ service:
     # externalTrafficPolicy: Cluster
     # loadBalancerIP: "1.2.3.4"
     # clusterIP: "2.3.4.5"
-  loadBalancerSourceRanges: {}
+  loadBalancerSourceRanges: []
     # - 192.168.0.1/32
     # - 172.16.0.0/16

```

## 6.0.0  ![AppVersion: 2.1.3](https://img.shields.io/static/v1?label=AppVersion&message=2.1.3&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-03-15

* Cleanup

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 7aebefe..490b2b6 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -18,15 +18,10 @@ ingressRoute:
   dashboard:
     enabled: true

-additional:
-  checkNewVersion: true
-  sendAnonymousUsage: true
-
 rollingUpdate:
   maxUnavailable: 1
   maxSurge: 1

-
 #
 # Add volumes to the traefik pod.
 # This can be used to mount a cert pair or a configmap that holds a config.toml file.
@@ -43,9 +38,14 @@ volumes: []
 #   mountPath: "/config"
 #   type: configMap

+globalArguments:
+  - "--global.checknewversion"
+  - "--global.sendanonymoususage"
+
 #
-# Configure Traefik entry points
+# Configure Traefik static configuration
 # Additional arguments to be passed at Traefik's binary
+# All available options available on https://docs.traefik.io/reference/static-configuration/cli/
 ## Use curly braces to pass values: `helm install --set="additionalArguments={--providers.kubernetesingress,--global.checknewversion=true}"`
 additionalArguments: []
 #  - "--providers.kubernetesingress"
@@ -63,7 +63,7 @@ env: {}
   # - name: SOME_OTHER_VAR
   #   value: some-other-var-value

-#
+# Configure ports
 ports:
   # The name of this one can't be changed as it is used for the readiness and
   # liveness probes, but you can adjust its config to your liking
@@ -94,7 +94,7 @@ ports:
     # hostPort: 8443
     expose: true
     exposedPort: 443
-  # nodePort: 32443
+    # nodePort: 32443

 # Options for the main traefik service, where the entrypoints traffic comes
 # from.
@@ -113,9 +113,6 @@ service:
     # - 192.168.0.1/32
     # - 172.16.0.0/16

-logs:
-  loglevel: WARN
-
 # Enable persistence using Persistent Volume Claims
 # ref: http://kubernetes.io/docs/user-guide/persistent-volumes/
 # After the pvc has been mounted, add the configs into traefik by using the `additionalArguments` list below, eg:
```

## 5.6.0  ![AppVersion: 2.1.3](https://img.shields.io/static/v1?label=AppVersion&message=2.1.3&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-03-12

* Add field enabled for resources

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 38bb263..7aebefe 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -7,11 +7,17 @@ image:
 # Configure the deployment
 #
 deployment:
+  enabled: true
   # Number of pods of the deployment
   replicas: 1
   # Additional pod annotations (e.g. for mesh injection or prometheus scraping)
   podAnnotations: {}

+# Create an IngressRoute for the dashboard
+ingressRoute:
+  dashboard:
+    enabled: true
+
 additional:
   checkNewVersion: true
   sendAnonymousUsage: true
@@ -93,6 +99,7 @@ ports:
 # Options for the main traefik service, where the entrypoints traffic comes
 # from.
 service:
+  enabled: true
   type: LoadBalancer
   # Additional annotations (e.g. for cloud provider specific config)
   annotations: {}
```

## 5.5.0  ![AppVersion: 2.1.3](https://img.shields.io/static/v1?label=AppVersion&message=2.1.3&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-03-12

* expose hostnetwork option

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index ecb2833..38bb263 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -123,6 +123,12 @@ persistence:
   path: /data
   annotations: {}

+# If hostNetwork is true, runs traefik in the host network namespace
+# To prevent unschedulabel pods due to port collisions, if hostNetwork=true
+# and replicas>1, a pod anti-affinity is recommended and will be set if the
+# affinity is left as default.
+hostNetwork: false
+
 resources: {}
   # requests:
   #   cpu: "100m"
@@ -131,5 +137,17 @@ resources: {}
   #   cpu: "300m"
   #   memory: "150Mi"
 affinity: {}
+# # This example pod anti-affinity forces the scheduler to put traefik pods
+# # on nodes where no other traefik pods are scheduled.
+# # It should be used when hostNetwork: true to prevent port conflicts
+#   podAntiAffinity:
+#     requiredDuringSchedulingIgnoredDuringExecution:
+#     - labelSelector:
+#         matchExpressions:
+#         - key: app
+#           operator: In
+#           values:
+#           - {{ template "traefik.name" . }}
+#       topologyKey: failure-domain.beta.kubernetes.io/zone
 nodeSelector: {}
 tolerations: []
```

## 5.4.0  ![AppVersion: 2.1.3](https://img.shields.io/static/v1?label=AppVersion&message=2.1.3&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-03-12

* Add support for hostport

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index ec1d619..ecb2833 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -63,6 +63,9 @@ ports:
   # liveness probes, but you can adjust its config to your liking
   traefik:
     port: 9000
+    # Use hostPort if set.
+    # hostPort: 9000
+
     # Defines whether the port is exposed if service.type is LoadBalancer or
     # NodePort.
     #
@@ -74,6 +77,7 @@ ports:
     exposedPort: 9000
   web:
     port: 8000
+    # hostPort: 8000
     expose: true
     exposedPort: 80
     # Use nodeport if set. This is useful if you have configured Traefik in a
@@ -81,6 +85,7 @@ ports:
     # nodePort: 32080
   websecure:
     port: 8443
+    # hostPort: 8443
     expose: true
     exposedPort: 443
   # nodePort: 32443
```

## 5.3.3  ![AppVersion: 2.1.3](https://img.shields.io/static/v1?label=AppVersion&message=2.1.3&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-03-12

* Fix replica check

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 7f31548..ec1d619 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -40,7 +40,7 @@ volumes: []
 #
 # Configure Traefik entry points
 # Additional arguments to be passed at Traefik's binary
-## Use curly braces to pass values: `helm install --set="{--providers.kubernetesingress,--global.checknewversion=true}" ."
+## Use curly braces to pass values: `helm install --set="additionalArguments={--providers.kubernetesingress,--global.checknewversion=true}"`
 additionalArguments: []
 #  - "--providers.kubernetesingress"

```

## 5.3.2  ![AppVersion: 2.1.3](https://img.shields.io/static/v1?label=AppVersion&message=2.1.3&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-03-11

* Fixed typo in README


## 5.3.1  ![AppVersion: 2.1.3](https://img.shields.io/static/v1?label=AppVersion&message=2.1.3&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-03-11

* Production ready


## 5.3.0  ![AppVersion: 2.1.3](https://img.shields.io/static/v1?label=AppVersion&message=2.1.3&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-03-11

* Not authorise acme if replica > 1


## 5.2.1  ![AppVersion: 2.1.3](https://img.shields.io/static/v1?label=AppVersion&message=2.1.3&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-03-11

* Fix volume mount


## 5.2.0  ![AppVersion: 2.1.3](https://img.shields.io/static/v1?label=AppVersion&message=2.1.3&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-03-11

* Add secret as env var

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index ccea845..7f31548 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -44,12 +44,18 @@ volumes: []
 additionalArguments: []
 #  - "--providers.kubernetesingress"

+# Secret to be set as environment variables to be passed to Traefik's binary
+secretEnv: {}
+  # - name: SOME_VAR
+  #   secretName: my-secret-name
+  #   secretKey: my-secret-key
+
 # Environment variables to be passed to Traefik's binary
 env: {}
-#  - name: SOME_VAR
-#    value: some-var-value
-#  - name: SOME_OTHER_VAR
-#    value: some-other-var-value
+  # - name: SOME_VAR
+  #   value: some-var-value
+  # - name: SOME_OTHER_VAR
+  #   value: some-other-var-value

 #
 ports:
```

## 5.1.0  ![AppVersion: 2.1.3](https://img.shields.io/static/v1?label=AppVersion&message=2.1.3&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-03-10

* Enhance security by add loadBalancerSourceRanges to lockdown ip address.

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 78bbee0..ccea845 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -91,6 +91,9 @@ service:
     # externalTrafficPolicy: Cluster
     # loadBalancerIP: "1.2.3.4"
     # clusterIP: "2.3.4.5"
+  loadBalancerSourceRanges: {}
+    # - 192.168.0.1/32
+    # - 172.16.0.0/16

 logs:
   loglevel: WARN
```

## 5.0.0  ![AppVersion: 2.1.3](https://img.shields.io/static/v1?label=AppVersion&message=2.1.3&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-03-10

* Expose dashboard by default but only on traefik entrypoint

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index a442fca..78bbee0 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -92,15 +92,6 @@ service:
     # loadBalancerIP: "1.2.3.4"
     # clusterIP: "2.3.4.5"

-dashboard:
-  # Enable the dashboard on Traefik
-  enable: true
-
-  # Expose the dashboard and api through an ingress route at /dashboard
-  # and /api This is not secure and SHOULD NOT be enabled on production
-  # deployments
-  ingressRoute: false
-
 logs:
   loglevel: WARN

```

## 4.1.3  ![AppVersion: 2.1.3](https://img.shields.io/static/v1?label=AppVersion&message=2.1.3&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-03-10

* Add annotations for PVC (#98)

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 8b2f4db..a442fca 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -116,6 +116,7 @@ persistence:
   size: 128Mi
   # storageClass: ""
   path: /data
+  annotations: {}

 resources: {}
   # requests:
```

## 4.1.2  ![AppVersion: 2.1.3](https://img.shields.io/static/v1?label=AppVersion&message=2.1.3&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-03-10

* Added persistent volume support. (#86)

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 2a2554f..8b2f4db 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -103,7 +103,20 @@ dashboard:

 logs:
   loglevel: WARN
-#
+
+# Enable persistence using Persistent Volume Claims
+# ref: http://kubernetes.io/docs/user-guide/persistent-volumes/
+# After the pvc has been mounted, add the configs into traefik by using the `additionalArguments` list below, eg:
+# additionalArguments:
+# - "--certificatesresolvers.le.acme.storage=/data/acme.json"
+# It will persist TLS certificates.
+persistence:
+  enabled: false
+  accessMode: ReadWriteOnce
+  size: 128Mi
+  # storageClass: ""
+  path: /data
+
 resources: {}
   # requests:
   #   cpu: "100m"
```

## 4.1.1  ![AppVersion: 2.1.3](https://img.shields.io/static/v1?label=AppVersion&message=2.1.3&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-03-10

* Add values to mount secrets or configmaps as volumes to the traefik pod (#84)

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 5401832..2a2554f 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -20,6 +20,23 @@ rollingUpdate:
   maxUnavailable: 1
   maxSurge: 1

+
+#
+# Add volumes to the traefik pod.
+# This can be used to mount a cert pair or a configmap that holds a config.toml file.
+# After the volume has been mounted, add the configs into traefik by using the `additionalArguments` list below, eg:
+# additionalArguments:
+# - "--providers.file.filename=/config/dynamic.toml"
+# - "--tls.certificates.certFile=/certs/tls.crt"
+# - "--tls.certificates.keyFile=/certs/tls.key"
+volumes: []
+# - name: public-cert
+#   mountPath: "/certs"
+#   type: secret
+# - name: configs
+#   mountPath: "/config"
+#   type: configMap
+
 #
 # Configure Traefik entry points
 # Additional arguments to be passed at Traefik's binary
```

## 4.1.0  ![AppVersion: 2.1.3](https://img.shields.io/static/v1?label=AppVersion&message=2.1.3&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-03-10

* Add podAnnotations to the deployment (#83)

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 5eab74b..5401832 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -9,6 +9,8 @@ image:
 deployment:
   # Number of pods of the deployment
   replicas: 1
+  # Additional pod annotations (e.g. for mesh injection or prometheus scraping)
+  podAnnotations: {}

 additional:
   checkNewVersion: true
```

## 4.0.0  ![AppVersion: 2.1.3](https://img.shields.io/static/v1?label=AppVersion&message=2.1.3&color=success&logo=) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-03-06

* Migrate to helm v3 (#94)


## 3.5.0  ![AppVersion: 2.1.3](https://img.shields.io/static/v1?label=AppVersion&message=2.1.3&color=success&logo=) ![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-02-18

* Publish helm chart (#81)


## 3.4.0  ![AppVersion: 2.1.3](https://img.shields.io/static/v1?label=AppVersion&message=2.1.3&color=success&logo=) ![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-02-13

* fix: tests.
* feat: bump traefik to v2.1.3
* Enable configuration of global checknewversion and sendanonymoususage (#80)

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index bcc42f8..5eab74b 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -1,7 +1,7 @@
 # Default values for Traefik
 image:
   name: traefik
-  tag: 2.1.1
+  tag: 2.1.3

 #
 # Configure the deployment
@@ -10,6 +10,10 @@ deployment:
   # Number of pods of the deployment
   replicas: 1

+additional:
+  checkNewVersion: true
+  sendAnonymousUsage: true
+
 rollingUpdate:
   maxUnavailable: 1
   maxSurge: 1
```

## 3.3.3  ![AppVersion: 2.1.1](https://img.shields.io/static/v1?label=AppVersion&message=2.1.1&color=success&logo=) ![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-02-05

* fix: deployment environment variables.
* fix: chart version.


## 3.3.2  ![AppVersion: 2.1.1](https://img.shields.io/static/v1?label=AppVersion&message=2.1.1&color=success&logo=) ![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-02-03

* ix: deployment environment variables.


## 3.3.1  ![AppVersion: 2.1.1](https://img.shields.io/static/v1?label=AppVersion&message=2.1.1&color=success&logo=) ![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-01-27

* fix: deployment environment variables.


## 3.3.0  ![AppVersion: 2.1.1](https://img.shields.io/static/v1?label=AppVersion&message=2.1.1&color=success&logo=) ![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-01-24

* Enable configuration of environment variables in traefik deployment (#71)

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index 4462359..bcc42f8 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -21,6 +21,13 @@ rollingUpdate:
 additionalArguments: []
 #  - "--providers.kubernetesingress"

+# Environment variables to be passed to Traefik's binary
+env: {}
+#  - name: SOME_VAR
+#    value: some-var-value
+#  - name: SOME_OTHER_VAR
+#    value: some-other-var-value
+
 #
 ports:
   # The name of this one can't be changed as it is used for the readiness and
```

## 3.2.1  ![AppVersion: 2.1.1](https://img.shields.io/static/v1?label=AppVersion&message=2.1.1&color=success&logo=) ![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-01-22

* Add Unit Tests for the chart (#60)


## 3.2.0  ![AppVersion: 2.1.1](https://img.shields.io/static/v1?label=AppVersion&message=2.1.1&color=success&logo=) ![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-01-22

* Make NodePort configurable (#67)

### Default value changes

```diff
diff --git a/traefik/values.yaml b/traefik/values.yaml
index b1fe42a..4462359 100644
--- a/traefik/values.yaml
+++ b/traefik/values.yaml
@@ -40,10 +40,14 @@ ports:
     port: 8000
     expose: true
     exposedPort: 80
+    # Use nodeport if set. This is useful if you have configured Traefik in a
+    # LoadBalancer
+    # nodePort: 32080
   websecure:
     port: 8443
     expose: true
     exposedPort: 443
+  # nodePort: 32443

 # Options for the main traefik service, where the entrypoints traffic comes
 # from.
```

## 3.1.0  ![AppVersion: 2.1.1](https://img.shields.io/static/v1?label=AppVersion&message=2.1.1&color=success&logo=) ![Helm: v2](https://img.shields.io/static/v1?label=Helm&message=v2&color=inactive&logo=helm) ![Helm: v3](https://img.shields.io/static/v1?label=Helm&message=v3&color=informational&logo=helm)

**Release date:** 2020-01-20

* Switch Chart linting to ct (#59)

### Default value changes

```diff
# Default values for Traefik
image:
  name: traefik
  tag: 2.1.1

#
# Configure the deployment
#
deployment:
  # Number of pods of the deployment
  replicas: 1

rollingUpdate:
  maxUnavailable: 1
  maxSurge: 1

#
# Configure Traefik entry points
# Additional arguments to be passed at Traefik's binary
## Use curly braces to pass values: `helm install --set="{--providers.kubernetesingress,--global.checknewversion=true}" ."
additionalArguments: []
#  - "--providers.kubernetesingress"

#
ports:
  # The name of this one can't be changed as it is used for the readiness and
  # liveness probes, but you can adjust its config to your liking
  traefik:
    port: 9000
    # Defines whether the port is exposed if service.type is LoadBalancer or
    # NodePort.
    #
    # You SHOULD NOT expose the traefik port on production deployments.
    # If you want to access it from outside of your cluster,
    # use `kubectl proxy` or create a secure ingress
    expose: false
    # The exposed port for this service
    exposedPort: 9000
  web:
    port: 8000
    expose: true
    exposedPort: 80
  websecure:
    port: 8443
    expose: true
    exposedPort: 443

# Options for the main traefik service, where the entrypoints traffic comes
# from.
service:
  type: LoadBalancer
  # Additional annotations (e.g. for cloud provider specific config)
  annotations: {}
  # Additional entries here will be added to the service spec. Cannot contains
  # type, selector or ports entries.
  spec: {}
    # externalTrafficPolicy: Cluster
    # loadBalancerIP: "1.2.3.4"
    # clusterIP: "2.3.4.5"

dashboard:
  # Enable the dashboard on Traefik
  enable: true

  # Expose the dashboard and api through an ingress route at /dashboard
  # and /api This is not secure and SHOULD NOT be enabled on production
  # deployments
  ingressRoute: false

logs:
  loglevel: WARN
#
resources: {}
  # requests:
  #   cpu: "100m"
  #   memory: "50Mi"
  # limits:
  #   cpu: "300m"
  #   memory: "150Mi"
affinity: {}
nodeSelector: {}
tolerations: []
```

---
Autogenerated from Helm Chart and git history using [helm-changelog](https://github.com/mogensen/helm-changelog)
