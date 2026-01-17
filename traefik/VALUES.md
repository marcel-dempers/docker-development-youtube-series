# traefik

![Version: 37.3.0](https://img.shields.io/badge/Version-37.3.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: v3.6.0](https://img.shields.io/badge/AppVersion-v3.6.0-informational?style=flat-square)

A Traefik based Kubernetes ingress controller

**Homepage:** <https://traefik.io/>

## Maintainers

| Name | Email | Url |
| ---- | ------ | --- |
| mloiseleur | <michel.loiseleur@traefik.io> |  |
| darkweaver87 | <remi.buisson@traefik.io> |  |
| jnoordsij |  |  |

## Source Code

* <https://github.com/traefik/traefik-helm-chart>
* <https://github.com/traefik/traefik>

## Requirements

Kubernetes: `>=1.22.0-0`

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| additionalArguments | list | `[]` | Additional arguments to be passed at Traefik's binary See [CLI Reference](https://docs.traefik.io/reference/static-configuration/cli/) Use curly braces to pass values: `helm install --set="additionalArguments={--providers.kubernetesingress.ingressclass=traefik-internal,--log.level=DEBUG}"` |
| additionalVolumeMounts | list | `[]` | Additional volumeMounts to add to the Traefik container |
| affinity | object | `{}` | on nodes where no other traefik pods are scheduled. It should be used when hostNetwork: true to prevent port conflicts |
| api.basePath | string | `""` | Configure API basePath |
| api.dashboard | bool | `true` | Enable the dashboard |
| autoscaling.behavior | object | `{}` | behavior configures the scaling behavior of the target in both Up and Down directions (scaleUp and scaleDown fields respectively). |
| autoscaling.enabled | bool | `false` | Create HorizontalPodAutoscaler object. See EXAMPLES.md for more details. |
| autoscaling.maxReplicas | string | `nil` | maxReplicas is the upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas. |
| autoscaling.metrics | list | `[]` | metrics contains the specifications for which to use to calculate the desired replica count (the maximum replica count across all metrics will be used). |
| autoscaling.minReplicas | string | `nil` | minReplicas is the lower limit for the number of replicas to which the autoscaler can scale down. It defaults to 1 pod. |
| autoscaling.scaleTargetRef | object | `{"apiVersion":"apps/v1","kind":"Deployment","name":"{{ template \"traefik.fullname\" . }}"}` | scaleTargetRef points to the target resource to scale, and is used for the pods for which metrics should be collected, as well as to actually change the replica count. |
| certificatesResolvers | object | `{}` | Certificates resolvers configuration. Ref: https://doc.traefik.io/traefik/https/acme/#certificate-resolvers See EXAMPLES.md for more details. |
| commonLabels | object | `{}` | Add additional label to all resources |
| core.defaultRuleSyntax | string | `""` | Can be used to use globally v2 router syntax. Deprecated since v3.4 /!\. See https://doc.traefik.io/traefik/v3.0/migration/v2-to-v3/#new-v3-syntax-notable-changes |
| deployment.additionalContainers | list | `[]` | Additional containers (e.g. for metric offloading sidecars) |
| deployment.additionalVolumes | list | `[]` | Additional volumes available for use with initContainers and additionalContainers |
| deployment.annotations | object | `{}` | Additional deployment annotations (e.g. for jaeger-operator sidecar injection) |
| deployment.dnsConfig | object | `{}` | Custom pod [DNS config](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.30/#poddnsconfig-v1-core) |
| deployment.dnsPolicy | string | `""` | Custom pod DNS policy. Apply if `hostNetwork: true` |
| deployment.enabled | bool | `true` | Enable deployment |
| deployment.goMemLimitPercentage | float | `0.9` | only takes effect when resources.limits.memory is set |
| deployment.healthchecksHost | string | `""` |  |
| deployment.healthchecksPort | string | `nil` |  |
| deployment.healthchecksScheme | string | `nil` |  |
| deployment.hostAliases | list | `[]` | Custom [host aliases](https://kubernetes.io/docs/tasks/network/customize-hosts-file-for-pods/) |
| deployment.imagePullSecrets | list | `[]` | Pull secret for fetching traefik container image |
| deployment.initContainers | list | `[]` | Additional initContainers (e.g. for setting file permission as shown below) |
| deployment.kind | string | `"Deployment"` | Deployment or DaemonSet |
| deployment.labels | object | `{}` | Additional deployment labels (e.g. for filtering deployment by custom labels) |
| deployment.lifecycle | object | `{}` | Pod lifecycle actions |
| deployment.livenessPath | string | `""` | Override the liveness path. Default: /ping |
| deployment.minReadySeconds | int | `0` | The minimum number of seconds Traefik needs to be up and running before the DaemonSet/Deployment controller considers it available |
| deployment.podAnnotations | object | `{}` | Additional pod annotations (e.g. for mesh injection or prometheus scraping) It supports templating. One can set it with values like traefik/name: '{{ template "traefik.name" . }}' |
| deployment.podLabels | object | `{}` | Additional Pod labels (e.g. for filtering Pod by custom labels) |
| deployment.readinessPath | string | `""` |  |
| deployment.replicas | int | `1` | Number of pods of the deployment (only applies when kind == Deployment) |
| deployment.revisionHistoryLimit | string | `nil` | Number of old history to retain to allow rollback (If not set, default Kubernetes value is set to 10) |
| deployment.runtimeClassName | string | `""` | Set a runtimeClassName on pod |
| deployment.shareProcessNamespace | bool | `false` | Use process namespace sharing |
| deployment.terminationGracePeriodSeconds | int | `60` | Amount of time (in seconds) before Kubernetes will send the SIGKILL signal if Traefik does not shut down |
| env | list | See _values.yaml_ | Additional Environment variables to be passed to Traefik's binary |
| envFrom | list | `[]` | Environment variables to be passed to Traefik's binary from configMaps or secrets |
| experimental.abortOnPluginFailure | bool | `false` | Defines whether all plugins must be loaded successfully for Traefik to start |
| experimental.fastProxy.debug | bool | `false` | Enable debug mode for the FastProxy implementation. |
| experimental.fastProxy.enabled | bool | `false` | Enables the FastProxy implementation. |
| experimental.knative | bool | `false` | Enable Knative provider experimental feature. |
| experimental.kubernetesGateway.enabled | bool | `false` | Enable traefik experimental GatewayClass CRD |
| experimental.localPlugins | object | `{}` | Enable experimental local plugins |
| experimental.otlpLogs | bool | `false` | Enable OTLP logging experimental feature. |
| experimental.plugins | object | `{}` | Enable experimental plugins |
| extraObjects | list | `[]` | Extra objects to deploy (value evaluated as a template)  In some cases, it can avoid the need for additional, extended or adhoc deployments. See #595 for more details and traefik/tests/values/extra.yaml for example. |
| gateway.annotations | object | `{}` | Additional gateway annotations (e.g. for cert-manager.io/issuer) |
| gateway.enabled | bool | `true` | When providers.kubernetesGateway.enabled, deploy a default gateway |
| gateway.infrastructure | object | `{}` | [Infrastructure](https://kubernetes.io/blog/2023/11/28/gateway-api-ga/#gateway-infrastructure-labels) |
| gateway.listeners | object | `{"web":{"hostname":"","namespacePolicy":null,"port":8000,"protocol":"HTTP"}}` | Define listeners |
| gateway.listeners.web.hostname | string | `""` | Optional hostname. See [Hostname](https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.Hostname) |
| gateway.listeners.web.namespacePolicy | object | `nil` | Routes are restricted to namespace of the gateway [by default](https://gateway-api.sigs.k8s.io/reference/spec/#gateway.networking.k8s.io/v1.FromNamespaces |
| gateway.listeners.web.port | int | `8000` | Port is the network port. Multiple listeners may use the same port, subject to the Listener compatibility rules. The port must match a port declared in ports section. |
| gateway.name | string | `""` | Set a custom name to gateway |
| gateway.namespace | string | `""` | By default, Gateway is created in the same `Namespace` as Traefik. |
| gatewayClass.enabled | bool | `true` | When providers.kubernetesGateway.enabled and gateway.enabled, deploy a default gatewayClass |
| gatewayClass.labels | object | `{}` | Additional gatewayClass labels (e.g. for filtering gateway objects by custom labels) |
| gatewayClass.name | string | `""` | Set a custom name to GatewayClass |
| global.azure | object | `{"enabled":false,"images":{"hub":{"image":"traefik-hub","registry":"ghcr.io/traefik","tag":"latest"},"proxy":{"image":"traefik","registry":"docker.io/library","tag":"latest"}}}` | Required for Azure Marketplace integration. See https://learn.microsoft.com/en-us/partner-center/marketplace-offers/azure-container-technical-assets-kubernetes?tabs=linux,linux2#update-the-helm-chart |
| global.azure.enabled | bool | `false` | Enable specific values for Azure Marketplace |
| global.checkNewVersion | bool | `true` |  |
| global.sendAnonymousUsage | bool | `false` | Please take time to consider whether or not you wish to share anonymous data with us See https://doc.traefik.io/traefik/contributing/data-collection/ |
| hostNetwork | bool | `false` | If hostNetwork is true, runs traefik in the host network namespace To prevent unschedulable pods due to port collisions, if hostNetwork=true and replicas>1, a pod anti-affinity is recommended and will be set if the affinity is left as default. |
| hub.aigateway.enabled | bool | `false` | Set to true in order to enable AI Gateway. Requires a valid license token. |
| hub.aigateway.maxRequestBodySize | string | `nil` | Hard limit for the size of request bodies inspected by the gateway. Accepts a plain integer representing **bytes**. The default value is `1048576` (1 MiB). |
| hub.apimanagement.admission.annotations | object | `{}` | Set custom annotations. |
| hub.apimanagement.admission.customWebhookCertificate | object | `{}` | Set custom certificate for the WebHook admission server. The certificate should be specified with _tls.crt_ and _tls.key_ in base64 encoding. |
| hub.apimanagement.admission.listenAddr | string | `""` | WebHook admission server listen address. Default: "0.0.0.0:9943". |
| hub.apimanagement.admission.restartOnCertificateChange | bool | `true` | Set it to false if you need to disable Traefik Hub pod restart when mutating webhook certificate is updated. It's done with a label update. |
| hub.apimanagement.admission.secretName | string | `"hub-agent-cert"` | Certificate name of the WebHook admission server. Default: "hub-agent-cert". |
| hub.apimanagement.admission.selfManagedCertificate | bool | `false` | By default, this chart handles directly the tls certificate required for the admission webhook. It's possible to disable this behavior and handle it outside of the chart. See EXAMPLES.md for more details. |
| hub.apimanagement.enabled | bool | `false` | Set to true in order to enable API Management. Requires a valid license token. |
| hub.apimanagement.openApi.validateRequestMethodAndPath | bool | `false` | When set to true, it will only accept paths and methods that are explicitly defined in its OpenAPI specification |
| hub.mcpgateway.enabled | bool | `false` | Set to true in order to enable AI MCP Gateway. Requires a valid license token. |
| hub.mcpgateway.maxRequestBodySize | string | `nil` | Hard limit for the size of request bodies inspected by the gateway. Accepts a plain integer representing **bytes**. The default value is `1048576` (1 MiB). |
| hub.namespaces | list | `[]` | By default, Traefik Hub provider watches all namespaces. When using `rbac.namespaced`, it will watch helm release namespace and namespaces listed in this array. |
| hub.offline | string | `nil` | Disables all external network connections. |
| hub.pluginRegistry.sources | object | `{}` |  |
| hub.providers.consulCatalogEnterprise.cache | bool | `false` | Use local agent caching for catalog reads. |
| hub.providers.consulCatalogEnterprise.connectAware | bool | `false` | Enable Consul Connect support. |
| hub.providers.consulCatalogEnterprise.connectByDefault | bool | `false` | Consider every service as Connect capable by default. |
| hub.providers.consulCatalogEnterprise.constraints | string | `""` | Constraints is an expression that Traefik matches against the container's labels |
| hub.providers.consulCatalogEnterprise.defaultRule | string | `"Host(`{{ normalize .Name }}`)"` | Default rule. |
| hub.providers.consulCatalogEnterprise.enabled | bool | `false` | Enable Consul Catalog Enterprise backend with default settings. |
| hub.providers.consulCatalogEnterprise.endpoint.address | string | `""` | The address of the Consul server |
| hub.providers.consulCatalogEnterprise.endpoint.datacenter | string | `""` | Data center to use. If not provided, the default agent data center is used |
| hub.providers.consulCatalogEnterprise.endpoint.endpointWaitTime | int | `0` | WaitTime limits how long a Watch will block. If not provided, the agent default |
| hub.providers.consulCatalogEnterprise.endpoint.httpauth.password | string | `""` | Basic Auth password |
| hub.providers.consulCatalogEnterprise.endpoint.httpauth.username | string | `""` | Basic Auth username |
| hub.providers.consulCatalogEnterprise.endpoint.scheme | string | `""` | The URI scheme for the Consul server |
| hub.providers.consulCatalogEnterprise.endpoint.tls.ca | string | `""` | TLS CA |
| hub.providers.consulCatalogEnterprise.endpoint.tls.cert | string | `""` | TLS cert |
| hub.providers.consulCatalogEnterprise.endpoint.tls.insecureSkipVerify | bool | `false` | TLS insecure skip verify |
| hub.providers.consulCatalogEnterprise.endpoint.tls.key | string | `""` | TLS key |
| hub.providers.consulCatalogEnterprise.endpoint.token | string | `""` | Token is used to provide a per-request ACL token which overrides the agent's |
| hub.providers.consulCatalogEnterprise.exposedByDefault | bool | `true` | Expose containers by default. |
| hub.providers.consulCatalogEnterprise.namespaces | string | `""` | Sets the namespaces used to discover services (Consul Enterprise only). |
| hub.providers.consulCatalogEnterprise.partition | string | `""` | Sets the partition used to discover services (Consul Enterprise only). |
| hub.providers.consulCatalogEnterprise.prefix | string | `"traefik"` | Prefix for consul service tags. |
| hub.providers.consulCatalogEnterprise.refreshInterval | int | `15` | Interval for checking Consul API. |
| hub.providers.consulCatalogEnterprise.requireConsistent | bool | `false` | Forces the read to be fully consistent. |
| hub.providers.consulCatalogEnterprise.serviceName | string | `"traefik"` | Name of the Traefik service in Consul Catalog (needs to be registered via the |
| hub.providers.consulCatalogEnterprise.stale | bool | `false` | Use stale consistency for catalog reads. |
| hub.providers.consulCatalogEnterprise.strictChecks | string | `"passing, warning"` | A list of service health statuses to allow taking traffic. |
| hub.providers.consulCatalogEnterprise.watch | bool | `false` | Watch Consul API events. |
| hub.providers.microcks.auth.clientId | string | `""` | Microcks API client ID. |
| hub.providers.microcks.auth.clientSecret | string | `""` | Microcks API client secret. |
| hub.providers.microcks.auth.endpoint | string | `""` | Microcks API endpoint. |
| hub.providers.microcks.auth.token | string | `""` | Microcks API token. |
| hub.providers.microcks.enabled | bool | `false` | Enable Microcks provider. |
| hub.providers.microcks.endpoint | string | `""` | Microcks API endpoint. |
| hub.providers.microcks.pollInterval | int | `30` | Polling interval for Microcks API. |
| hub.providers.microcks.pollTimeout | int | `5` | Polling timeout for Microcks API. |
| hub.providers.microcks.tls.ca | string | `""` | TLS CA |
| hub.providers.microcks.tls.cert | string | `""` | TLS cert |
| hub.providers.microcks.tls.insecureSkipVerify | bool | `false` | TLS insecure skip verify |
| hub.providers.microcks.tls.key | string | `""` | TLS key |
| hub.redis.cluster | string | `nil` | Enable Redis Cluster. Default: true. |
| hub.redis.database | string | `nil` | Database used to store information. Default: "0". |
| hub.redis.endpoints | string | `""` | Endpoints of the Redis instances to connect to. Default: "". |
| hub.redis.password | string | `""` | The password to use when connecting to Redis endpoints. Default: "". |
| hub.redis.sentinel.masterset | string | `""` | Name of the set of main nodes to use for main selection. Required when using Sentinel. Default: "". |
| hub.redis.sentinel.password | string | `""` | Password to use for sentinel authentication (can be different from endpoint password). Default: "". |
| hub.redis.sentinel.username | string | `""` | Username to use for sentinel authentication (can be different from endpoint username). Default: "". |
| hub.redis.timeout | string | `""` | Timeout applied on connection with redis. Default: "0s". |
| hub.redis.tls.ca | string | `""` | Path to the certificate authority used for the secured connection. |
| hub.redis.tls.cert | string | `""` | Path to the public certificate used for the secure connection. |
| hub.redis.tls.insecureSkipVerify | bool | `false` | When insecureSkipVerify is set to true, the TLS connection accepts any certificate presented by the server. Default: false. |
| hub.redis.tls.key | string | `""` | Path to the private key used for the secure connection. |
| hub.redis.username | string | `""` | The username to use when connecting to Redis endpoints. Default: "". |
| hub.sendlogs | string | `nil` |  |
| hub.token | string | `""` | Name of `Secret` with key 'token' set to a valid license token. It enables API Gateway. |
| hub.tracing.additionalTraceHeaders | object | `{"enabled":false,"traceContext":{"parentId":"","traceId":"","traceParent":"","traceState":""}}` | Tracing headers to duplicate. To configure the following, tracing.otlp.enabled needs to be set to true. |
| hub.tracing.additionalTraceHeaders.traceContext.parentId | string | `""` | Name of the header that will contain the parent-id header copy. |
| hub.tracing.additionalTraceHeaders.traceContext.traceId | string | `""` | Name of the header that will contain the trace-id copy. |
| hub.tracing.additionalTraceHeaders.traceContext.traceParent | string | `""` | Name of the header that will contain the traceparent copy. |
| hub.tracing.additionalTraceHeaders.traceContext.traceState | string | `""` | Name of the header that will contain the tracestate copy. |
| image.pullPolicy | string | `"IfNotPresent"` | Traefik image pull policy |
| image.registry | string | `"docker.io"` | Traefik image host registry |
| image.repository | string | `"traefik"` | Traefik image repository |
| image.tag | string | `nil` | defaults to appVersion. It's used for version checking, even prefixed with experimental- or latest-. When a digest is required, `versionOverride` can be used to set the version. |
| ingressClass | object | `{"enabled":true,"isDefaultClass":true,"name":""}` | Create a default IngressClass for Traefik |
| ingressRoute | object | `{"dashboard":{"annotations":{},"enabled":false,"entryPoints":["traefik"],"labels":{},"matchRule":"PathPrefix(`/dashboard`) || PathPrefix(`/api`)","middlewares":[],"services":[{"kind":"TraefikService","name":"api@internal"}],"tls":{}},"healthcheck":{"annotations":{},"enabled":false,"entryPoints":["traefik"],"labels":{},"matchRule":"PathPrefix(`/ping`)","middlewares":[],"services":[{"kind":"TraefikService","name":"ping@internal"}],"tls":{}}}` | Only dashboard & healthcheck IngressRoute are supported. It's recommended to create workloads CR outside of this Chart. |
| ingressRoute.dashboard.annotations | object | `{}` | Additional ingressRoute annotations (e.g. for kubernetes.io/ingress.class) |
| ingressRoute.dashboard.enabled | bool | `false` | Create an IngressRoute for the dashboard |
| ingressRoute.dashboard.entryPoints | list | `["traefik"]` | Specify the allowed entrypoints to use for the dashboard ingress route, (e.g. traefik, web, websecure). By default, it's using traefik entrypoint, which is not exposed. /!\ Do not expose your dashboard without any protection over the internet /!\ |
| ingressRoute.dashboard.labels | object | `{}` | Additional ingressRoute labels (e.g. for filtering IngressRoute by custom labels) |
| ingressRoute.dashboard.matchRule | string | `"PathPrefix(`/dashboard`) || PathPrefix(`/api`)"` | The router match rule used for the dashboard ingressRoute |
| ingressRoute.dashboard.middlewares | list | `[]` | Additional ingressRoute middlewares (e.g. for authentication) |
| ingressRoute.dashboard.services | list | `[{"kind":"TraefikService","name":"api@internal"}]` | The internal service used for the dashboard ingressRoute |
| ingressRoute.dashboard.tls | object | `{}` | TLS options (e.g. secret containing certificate) |
| ingressRoute.healthcheck.annotations | object | `{}` | Additional ingressRoute annotations (e.g. for kubernetes.io/ingress.class) |
| ingressRoute.healthcheck.enabled | bool | `false` | Create an IngressRoute for the healthcheck probe |
| ingressRoute.healthcheck.entryPoints | list | `["traefik"]` | Specify the allowed entrypoints to use for the healthcheck ingress route, (e.g. traefik, web, websecure). By default, it's using traefik entrypoint, which is not exposed. |
| ingressRoute.healthcheck.labels | object | `{}` | Additional ingressRoute labels (e.g. for filtering IngressRoute by custom labels) |
| ingressRoute.healthcheck.matchRule | string | `"PathPrefix(`/ping`)"` | The router match rule used for the healthcheck ingressRoute |
| ingressRoute.healthcheck.middlewares | list | `[]` | Additional ingressRoute middlewares (e.g. for authentication) |
| ingressRoute.healthcheck.services | list | `[{"kind":"TraefikService","name":"ping@internal"}]` | The internal service used for the healthcheck ingressRoute |
| ingressRoute.healthcheck.tls | object | `{}` | TLS options (e.g. secret containing certificate) |
| instanceLabelOverride | string | `""` | This field overrides the default app.kubernetes.io/instance label for all Objects. |
| livenessProbe.failureThreshold | int | `3` | The number of consecutive failures allowed before considering the probe as failed. |
| livenessProbe.initialDelaySeconds | int | `2` | The number of seconds to wait before starting the first probe. |
| livenessProbe.periodSeconds | int | `10` | The number of seconds to wait between consecutive probes. |
| livenessProbe.successThreshold | int | `1` | The minimum consecutive successes required to consider the probe successful. |
| livenessProbe.timeoutSeconds | int | `2` | The number of seconds to wait for a probe response before considering it as failed. |
| logs.access.addInternals | bool | `false` | Enables accessLogs for internal resources. Default: false. |
| logs.access.bufferingSize | string | `nil` | Set [bufferingSize](https://doc.traefik.io/traefik/observability/access-logs/#bufferingsize) |
| logs.access.enabled | bool | `false` | To enable access logs |
| logs.access.fields.general.defaultmode | string | `"keep"` | Set default mode for fields.names |
| logs.access.fields.general.names | object | `{}` | Names of the fields to limit. |
| logs.access.fields.headers | object | `{"defaultmode":"drop","names":{}}` | [Limit logged fields or headers](https://doc.traefik.io/traefik/observability/access-logs/#limiting-the-fieldsincluding-headers) |
| logs.access.fields.headers.defaultmode | string | `"drop"` | Set default mode for fields.headers |
| logs.access.filters | object | `{"minduration":"","retryattempts":false,"statuscodes":""}` | Set [filtering](https://docs.traefik.io/observability/access-logs/#filtering) |
| logs.access.filters.minduration | string | `""` | Set minDuration, to keep access logs when requests take longer than the specified duration |
| logs.access.filters.retryattempts | bool | `false` | Set retryAttempts, to keep the access logs when at least one retry has happened |
| logs.access.filters.statuscodes | string | `""` | Set statusCodes, to limit the access logs to requests with a status codes in the specified range |
| logs.access.format | string | `nil` | Set [access log format](https://doc.traefik.io/traefik/observability/access-logs/#format) |
| logs.access.otlp.enabled | bool | `false` | Set to true in order to enable OpenTelemetry on access logs. Note that experimental.otlpLogs needs to be enabled. |
| logs.access.otlp.grpc.enabled | bool | `false` | Set to true in order to send access logs to the OpenTelemetry Collector using gRPC |
| logs.access.otlp.grpc.endpoint | string | `""` | Format: <host>:<port>. Default: "localhost:4317" |
| logs.access.otlp.grpc.insecure | bool | `false` | Allows reporter to send access logs to the OpenTelemetry Collector without using a secured protocol. |
| logs.access.otlp.grpc.tls.ca | string | `""` | The path to the certificate authority, it defaults to the system bundle. |
| logs.access.otlp.grpc.tls.cert | string | `""` | The path to the public certificate. When using this option, setting the key option is required. |
| logs.access.otlp.grpc.tls.insecureSkipVerify | bool | `false` | When set to true, the TLS connection accepts any certificate presented by the server regardless of the hostnames it covers. |
| logs.access.otlp.grpc.tls.key | string | `""` | The path to the private key. When using this option, setting the cert option is required. |
| logs.access.otlp.http.enabled | bool | `false` | Set to true in order to send access logs to the OpenTelemetry Collector using HTTP. |
| logs.access.otlp.http.endpoint | string | `""` | Format: <scheme>://<host>:<port><path>. Default: https://localhost:4318/v1/logs |
| logs.access.otlp.http.headers | object | `{}` | Additional headers sent with access logs by the reporter to the OpenTelemetry Collector. |
| logs.access.otlp.http.tls.ca | string | `""` | The path to the certificate authority, it defaults to the system bundle. |
| logs.access.otlp.http.tls.cert | string | `""` | The path to the public certificate. When using this option, setting the key option is required. |
| logs.access.otlp.http.tls.insecureSkipVerify | string | `nil` | When set to true, the TLS connection accepts any certificate presented by the server regardless of the hostnames it covers. |
| logs.access.otlp.http.tls.key | string | `""` | The path to the private key. When using this option, setting the cert option is required. |
| logs.access.otlp.resourceAttributes | object | `{}` | Defines additional resource attributes to be sent to the collector. |
| logs.access.otlp.serviceName | string | `nil` | Service name used in OTLP backend. Default: traefik. |
| logs.access.timezone | string | `""` | Set [timezone](https://doc.traefik.io/traefik/observability/access-logs/#time-zones) |
| logs.general.filePath | string | `""` | To write the logs into a log file, use the filePath option. |
| logs.general.format | string | `nil` | Set [logs format](https://doc.traefik.io/traefik/observability/logs/#format) |
| logs.general.level | string | `"INFO"` | Alternative logging levels are TRACE, DEBUG, INFO, WARN, ERROR, FATAL, and PANIC. |
| logs.general.noColor | bool | `false` | When set to true and format is common, it disables the colorized output. |
| logs.general.otlp.enabled | bool | `false` | Set to true in order to enable OpenTelemetry on logs. Note that experimental.otlpLogs needs to be enabled. |
| logs.general.otlp.grpc.enabled | bool | `false` | Set to true in order to send logs  to the OpenTelemetry Collector using gRPC |
| logs.general.otlp.grpc.endpoint | string | `""` | Format: <host>:<port>. Default: "localhost:4317" |
| logs.general.otlp.grpc.insecure | bool | `false` | Allows reporter to send logs to the OpenTelemetry Collector without using a secured protocol. |
| logs.general.otlp.grpc.tls.ca | string | `""` | The path to the certificate authority, it defaults to the system bundle. |
| logs.general.otlp.grpc.tls.cert | string | `""` | The path to the public certificate. When using this option, setting the key option is required. |
| logs.general.otlp.grpc.tls.insecureSkipVerify | bool | `false` | When set to true, the TLS connection accepts any certificate presented by the server regardless of the hostnames it covers. |
| logs.general.otlp.grpc.tls.key | string | `""` | The path to the private key. When using this option, setting the cert option is required. |
| logs.general.otlp.http.enabled | bool | `false` | Set to true in order to send logs to the OpenTelemetry Collector using HTTP. |
| logs.general.otlp.http.endpoint | string | `""` | Format: <scheme>://<host>:<port><path>. Default: https://localhost:4318/v1/logs |
| logs.general.otlp.http.headers | object | `{}` | Additional headers sent with logs by the reporter to the OpenTelemetry Collector. |
| logs.general.otlp.http.tls.ca | string | `""` | The path to the certificate authority, it defaults to the system bundle. |
| logs.general.otlp.http.tls.cert | string | `""` | The path to the public certificate. When using this option, setting the key option is required. |
| logs.general.otlp.http.tls.insecureSkipVerify | string | `nil` | When set to true, the TLS connection accepts any certificate presented by the server regardless of the hostnames it covers. |
| logs.general.otlp.http.tls.key | string | `""` | The path to the private key. When using this option, setting the cert option is required. |
| logs.general.otlp.resourceAttributes | object | `{}` | Defines additional resource attributes to be sent to the collector. |
| logs.general.otlp.serviceName | string | `nil` | Service name used in OTLP backend. Default: traefik. |
| metrics.addInternals | bool | `false` | Enable metrics for internal resources. Default: false |
| metrics.otlp.addEntryPointsLabels | string | `nil` | Enable metrics on entry points. Default: true |
| metrics.otlp.addRoutersLabels | string | `nil` | Enable metrics on routers. Default: false |
| metrics.otlp.addServicesLabels | string | `nil` | Enable metrics on services. Default: true |
| metrics.otlp.enabled | bool | `false` | Set to true in order to enable the OpenTelemetry metrics |
| metrics.otlp.explicitBoundaries | list | `[]` | Explicit boundaries for Histogram data points. Default: [.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10] |
| metrics.otlp.grpc.enabled | bool | `false` | Set to true in order to send metrics to the OpenTelemetry Collector using gRPC |
| metrics.otlp.grpc.endpoint | string | `""` | Format: <host>:<port>. Default: "localhost:4317" |
| metrics.otlp.grpc.insecure | bool | `false` | Allows reporter to send metrics to the OpenTelemetry Collector without using a secured protocol. |
| metrics.otlp.grpc.tls.ca | string | `""` | The path to the certificate authority, it defaults to the system bundle. |
| metrics.otlp.grpc.tls.cert | string | `""` | The path to the public certificate. When using this option, setting the key option is required. |
| metrics.otlp.grpc.tls.insecureSkipVerify | bool | `false` | When set to true, the TLS connection accepts any certificate presented by the server regardless of the hostnames it covers. |
| metrics.otlp.grpc.tls.key | string | `""` | The path to the private key. When using this option, setting the cert option is required. |
| metrics.otlp.http.enabled | bool | `false` | Set to true in order to send metrics to the OpenTelemetry Collector using HTTP. |
| metrics.otlp.http.endpoint | string | `""` | Format: <scheme>://<host>:<port><path>. Default: https://localhost:4318/v1/metrics |
| metrics.otlp.http.headers | object | `{}` | Additional headers sent with metrics by the reporter to the OpenTelemetry Collector. |
| metrics.otlp.http.tls.ca | string | `""` | The path to the certificate authority, it defaults to the system bundle. |
| metrics.otlp.http.tls.cert | string | `""` | The path to the public certificate. When using this option, setting the key option is required. |
| metrics.otlp.http.tls.insecureSkipVerify | string | `nil` | When set to true, the TLS connection accepts any certificate presented by the server regardless of the hostnames it covers. |
| metrics.otlp.http.tls.key | string | `""` | The path to the private key. When using this option, setting the cert option is required. |
| metrics.otlp.pushInterval | string | `""` | Interval at which metrics are sent to the OpenTelemetry Collector. Default: 10s |
| metrics.otlp.resourceAttributes | object | `{}` | Defines additional resource attributes to be sent to the collector. |
| metrics.otlp.serviceName | string | `nil` | Service name used in OTLP backend. Default: traefik. |
| metrics.prometheus.addEntryPointsLabels | string | `nil` | Enable metrics on entry points. Default: true |
| metrics.prometheus.addRoutersLabels | string | `nil` | Enable metrics on routers. Default: false |
| metrics.prometheus.addServicesLabels | string | `nil` | Enable metrics on services. Default: true |
| metrics.prometheus.buckets | string | `""` | Buckets for latency metrics. Default="0.1,0.3,1.2,5.0" |
| metrics.prometheus.disableAPICheck | string | `nil` | When set to true, it won't check if Prometheus Operator CRDs are deployed |
| metrics.prometheus.entryPoint | string | `"metrics"` | Entry point used to expose metrics. |
| metrics.prometheus.headerLabels | object | `{}` | Add HTTP header labels to metrics. See EXAMPLES.md or upstream doc for usage. |
| metrics.prometheus.manualRouting | bool | `false` | When manualRouting is true, it disables the default internal router in # order to allow creating a custom router for prometheus@internal service. |
| metrics.prometheus.prometheusRule.additionalLabels | object | `{}` |  |
| metrics.prometheus.prometheusRule.apiVersion | string | `"monitoring.coreos.com/v1"` |  |
| metrics.prometheus.prometheusRule.enabled | bool | `false` | Enable optional CR for Prometheus Operator. See EXAMPLES.md for more details. |
| metrics.prometheus.prometheusRule.namespace | string | `""` |  |
| metrics.prometheus.service.annotations | object | `{}` |  |
| metrics.prometheus.service.enabled | bool | `false` | Create a dedicated metrics service to use with ServiceMonitor |
| metrics.prometheus.service.labels | object | `{}` |  |
| metrics.prometheus.serviceMonitor.additionalLabels | object | `{}` |  |
| metrics.prometheus.serviceMonitor.apiVersion | string | `"monitoring.coreos.com/v1"` |  |
| metrics.prometheus.serviceMonitor.enableHttp2 | bool | `false` |  |
| metrics.prometheus.serviceMonitor.enabled | bool | `false` | Enable optional CR for Prometheus Operator. See EXAMPLES.md for more details. |
| metrics.prometheus.serviceMonitor.followRedirects | bool | `false` |  |
| metrics.prometheus.serviceMonitor.honorLabels | bool | `false` |  |
| metrics.prometheus.serviceMonitor.honorTimestamps | bool | `false` |  |
| metrics.prometheus.serviceMonitor.interval | string | `""` |  |
| metrics.prometheus.serviceMonitor.jobLabel | string | `""` |  |
| metrics.prometheus.serviceMonitor.metricRelabelings | list | `[]` |  |
| metrics.prometheus.serviceMonitor.namespace | string | `""` |  |
| metrics.prometheus.serviceMonitor.namespaceSelector | object | `{}` |  |
| metrics.prometheus.serviceMonitor.relabelings | list | `[]` |  |
| metrics.prometheus.serviceMonitor.scrapeTimeout | string | `""` |  |
| namespaceOverride | string | `""` | This field overrides the default Release Namespace for Helm. It will not affect optional CRDs such as `ServiceMonitor` and `PrometheusRules` |
| nodeSelector | object | `{}` | nodeSelector is the simplest recommended form of node selection constraint. |
| oci_meta | object | `{"enabled":false,"images":{"hub":{"image":"traefik-hub","tag":"latest"},"proxy":{"image":"traefik","tag":"latest"}},"repo":"traefik"}` | Required for OCI Marketplace integration. See https://docs.public.content.oci.oraclecloud.com/en-us/iaas/Content/Marketplace/understanding-helm-charts.htm |
| oci_meta.enabled | bool | `false` | Enable specific values for Oracle Cloud Infrastructure |
| oci_meta.repo | string | `"traefik"` | It needs to be an ocir repo |
| ocsp.enabled | bool | `false` | Enable OCSP stapling support. See https://doc.traefik.io/traefik/https/ocsp/#overview |
| ocsp.responderOverrides | object | `{}` | Defines the OCSP responder URLs to use instead of the one provided by the certificate. |
| persistence.accessMode | string | `"ReadWriteOnce"` |  |
| persistence.annotations | object | `{}` |  |
| persistence.enabled | bool | `false` | Enable persistence using Persistent Volume Claims ref: http://kubernetes.io/docs/user-guide/persistent-volumes/. It can be used to store TLS certificates along with `certificatesResolvers.<name>.acme.storage`  option |
| persistence.existingClaim | string | `""` |  |
| persistence.name | string | `"data"` |  |
| persistence.path | string | `"/data"` |  |
| persistence.size | string | `"128Mi"` |  |
| persistence.storageClass | string | `""` |  |
| persistence.subPath | string | `""` | Only mount a subpath of the Volume into the pod |
| persistence.volumeName | string | `""` |  |
| podDisruptionBudget | object | `{"enabled":false,"maxUnavailable":null,"minAvailable":null}` | [Pod Disruption Budget](https://kubernetes.io/docs/reference/kubernetes-api/policy-resources/pod-disruption-budget-v1/) |
| podSecurityContext | object | See _values.yaml_ | [Pod Security Context](https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/pod-v1/#security-context) |
| podSecurityPolicy | object | `{"enabled":false}` | Enable to create a PodSecurityPolicy and assign it to the Service Account via RoleBinding or ClusterRoleBinding |
| ports.metrics.expose | object | `{"default":false}` | You may not want to expose the metrics port on production deployments. If you want to access it from outside your cluster, use `kubectl port-forward` or create a secure ingress |
| ports.metrics.exposedPort | int | `9100` | The exposed port for this service |
| ports.metrics.observability.accessLogs | string | `nil` | Enables access-logs for this entryPoint. |
| ports.metrics.observability.metrics | string | `nil` | Enables metrics for this entryPoint. |
| ports.metrics.observability.traceVerbosity | string | `nil` | Defines the tracing verbosity level for this entryPoint. |
| ports.metrics.observability.tracing | string | `nil` | Enables tracing for this entryPoint. |
| ports.metrics.port | int | `9100` | When using hostNetwork, use another port to avoid conflict with node exporter: https://github.com/prometheus/prometheus/wiki/Default-port-allocations |
| ports.metrics.protocol | string | `"TCP"` | The port protocol (TCP/UDP) |
| ports.traefik.expose | object | `{"default":false}` | You SHOULD NOT expose the traefik port on production deployments. If you want to access it from outside your cluster, use `kubectl port-forward` or create a secure ingress |
| ports.traefik.exposedPort | int | `8080` | The exposed port for this service |
| ports.traefik.hostIP | string | `nil` | Use hostIP if set. If not set, Kubernetes will default to 0.0.0.0, which means it's listening on all your interfaces and all your IPs. You may want to set this value if you need traefik to listen on specific interface only. |
| ports.traefik.hostPort | string | `nil` | Use hostPort if set. |
| ports.traefik.observability.accessLogs | string | `nil` | Defines whether a router attached to this EntryPoint produces access-logs by default. |
| ports.traefik.observability.metrics | string | `nil` | Defines whether a router attached to this EntryPoint produces metrics by default. |
| ports.traefik.observability.traceVerbosity | string | `nil` | Defines the tracing verbosity level for routers attached to this EntryPoint. |
| ports.traefik.observability.tracing | string | `nil` | Defines whether a router attached to this EntryPoint produces traces by default. |
| ports.traefik.port | int | `8080` |  |
| ports.traefik.protocol | string | `"TCP"` | The port protocol (TCP/UDP) |
| ports.web.expose.default | bool | `true` |  |
| ports.web.exposedPort | int | `80` |  |
| ports.web.forwardedHeaders.insecure | bool | `false` |  |
| ports.web.forwardedHeaders.trustedIPs | list | `[]` | Trust forwarded headers information (X-Forwarded-*). |
| ports.web.nodePort | string | `nil` | See [upstream documentation](https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport) |
| ports.web.observability.accessLogs | string | `nil` | Enables access-logs for this entryPoint. |
| ports.web.observability.metrics | string | `nil` | Enables metrics for this entryPoint. |
| ports.web.observability.traceVerbosity | string | `nil` | Defines the tracing verbosity level for this entryPoint. |
| ports.web.observability.tracing | string | `nil` | Enables tracing for this entryPoint. |
| ports.web.port | int | `8000` |  |
| ports.web.protocol | string | `"TCP"` |  |
| ports.web.proxyProtocol.insecure | bool | `false` |  |
| ports.web.proxyProtocol.trustedIPs | list | `[]` | Enable the Proxy Protocol header parsing for the entry point |
| ports.web.redirections.entryPoint | object | `{}` | Port Redirections Added in 2.2, one can make permanent redirects via entrypoints. Same sets of parameters: to, scheme, permanent and priority. https://docs.traefik.io/routing/entrypoints/#redirection |
| ports.web.targetPort | string | `nil` |  |
| ports.web.transport | object | `{"keepAliveMaxRequests":null,"keepAliveMaxTime":null,"lifeCycle":{"graceTimeOut":null,"requestAcceptGraceTimeout":null},"respondingTimeouts":{"idleTimeout":null,"readTimeout":null,"writeTimeout":null}}` | Set transport settings for the entrypoint; see also https://doc.traefik.io/traefik/routing/entrypoints/#transport |
| ports.websecure.allowACMEByPass | bool | `false` | See [upstream documentation](https://doc.traefik.io/traefik/routing/entrypoints/#allowacmebypass) |
| ports.websecure.appProtocol | string | `nil` | See [upstream documentation](https://kubernetes.io/docs/concepts/services-networking/service/#application-protocol) |
| ports.websecure.containerPort | string | `nil` |  |
| ports.websecure.expose.default | bool | `true` |  |
| ports.websecure.exposedPort | int | `443` |  |
| ports.websecure.forwardedHeaders.insecure | bool | `false` |  |
| ports.websecure.forwardedHeaders.trustedIPs | list | `[]` | Trust forwarded headers information (X-Forwarded-*). |
| ports.websecure.hostPort | string | `nil` |  |
| ports.websecure.http3.advertisedPort | string | `nil` |  |
| ports.websecure.http3.enabled | bool | `false` |  |
| ports.websecure.middlewares | list | `[]` | /!\ It introduces here a link between your static configuration and your dynamic configuration /!\ It follows the provider naming convention: https://doc.traefik.io/traefik/providers/overview/#provider-namespace   - namespace-name1@kubernetescrd   - namespace-name2@kubernetescrd |
| ports.websecure.nodePort | string | `nil` | See [upstream documentation](https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport) |
| ports.websecure.observability.accessLogs | string | `nil` | Enables access-logs for this entryPoint. |
| ports.websecure.observability.metrics | string | `nil` | Enables metrics for this entryPoint. |
| ports.websecure.observability.traceVerbosity | string | `nil` | Defines the tracing verbosity level for this entryPoint. |
| ports.websecure.observability.tracing | string | `nil` | Enables tracing for this entryPoint. |
| ports.websecure.port | int | `8443` |  |
| ports.websecure.protocol | string | `"TCP"` |  |
| ports.websecure.proxyProtocol.insecure | bool | `false` |  |
| ports.websecure.proxyProtocol.trustedIPs | list | `[]` | Enable the Proxy Protocol header parsing for the entry point |
| ports.websecure.targetPort | string | `nil` |  |
| ports.websecure.tls | object | `{"certResolver":"","domains":[],"enabled":true,"options":""}` | See [upstream documentation](https://doc.traefik.io/traefik/routing/entrypoints/#tls) |
| ports.websecure.transport | object | `{"keepAliveMaxRequests":null,"keepAliveMaxTime":null,"lifeCycle":{"graceTimeOut":null,"requestAcceptGraceTimeout":null},"respondingTimeouts":{"idleTimeout":null,"readTimeout":null,"writeTimeout":null}}` | See [upstream documentation](https://doc.traefik.io/traefik/routing/entrypoints/#transport) |
| priorityClassName | string | `""` | [Pod Priority and Preemption](https://kubernetes.io/docs/concepts/scheduling-eviction/pod-priority-preemption/) |
| providers.file.content | string | `""` | File content (YAML format, go template supported) (see https://doc.traefik.io/traefik/providers/file/) |
| providers.file.enabled | bool | `false` | Create a file provider |
| providers.file.watch | bool | `true` | Allows Traefik to automatically watch for file changes |
| providers.knative.enabled | bool | `false` | Enable Knative provider |
| providers.knative.labelselector | string | `""` | Allow filtering Knative Ingress objects |
| providers.knative.namespaces | list | `[]` | Array of namespaces to watch. If left empty, Traefik watches all namespaces. . When using `rbac.namespaced`, it will watch helm release namespace and namespaces listed in this array. |
| providers.kubernetesCRD.allowCrossNamespace | bool | `false` | Allows IngressRoute to reference resources in namespace other than theirs |
| providers.kubernetesCRD.allowEmptyServices | bool | `true` | Allows to return 503 when there are no endpoints available |
| providers.kubernetesCRD.allowExternalNameServices | bool | `false` | Allows to reference ExternalName services in IngressRoute |
| providers.kubernetesCRD.enabled | bool | `true` | Load Kubernetes IngressRoute provider |
| providers.kubernetesCRD.ingressClass | string | `""` | When the parameter is set, only resources containing an annotation with the same value are processed. Otherwise, resources missing the annotation, having an empty value, or the value traefik are processed. It will also set required annotation on Dashboard and Healthcheck IngressRoute when enabled. |
| providers.kubernetesCRD.namespaces | list | `[]` | Array of namespaces to watch. If left empty, Traefik watches all namespaces. . When using `rbac.namespaced`, it will watch helm release namespace and namespaces listed in this array. |
| providers.kubernetesCRD.nativeLBByDefault | bool | `false` | Defines whether to use Native Kubernetes load-balancing mode by default. |
| providers.kubernetesGateway.enabled | bool | `false` | Enable Traefik Gateway provider for Gateway API |
| providers.kubernetesGateway.experimentalChannel | bool | `false` | Toggles support for the Experimental Channel resources (Gateway API release channels documentation). This option currently enables support for TCPRoute and TLSRoute. |
| providers.kubernetesGateway.labelselector | string | `""` | A label selector can be defined to filter on specific GatewayClass objects only. |
| providers.kubernetesGateway.namespaces | list | `[]` | Array of namespaces to watch. If left empty, Traefik watches all namespaces. . When using `rbac.namespaced`, it will watch helm release namespace and namespaces listed in this array. |
| providers.kubernetesGateway.nativeLBByDefault | bool | `false` | Defines whether to use Native Kubernetes load-balancing mode by default. |
| providers.kubernetesGateway.statusAddress.hostname | string | `""` | This Hostname will get copied to the Gateway status.addresses. |
| providers.kubernetesGateway.statusAddress.ip | string | `""` | This IP will get copied to the Gateway status.addresses, and currently only supports one IP value (IPv4 or IPv6). |
| providers.kubernetesGateway.statusAddress.service | object | `{"enabled":true,"name":"","namespace":""}` | The Kubernetes service to copy status addresses from. When using third parties tools like External-DNS, this option can be used to copy the service loadbalancer.status (containing the service's endpoints IPs) to the gateways. Default to Service of this Chart. |
| providers.kubernetesIngress.allowEmptyServices | bool | `true` | Allows to return 503 when there are no endpoints available |
| providers.kubernetesIngress.allowExternalNameServices | bool | `false` | Allows to reference ExternalName services in Ingress |
| providers.kubernetesIngress.enabled | bool | `true` | Load Kubernetes Ingress provider |
| providers.kubernetesIngress.ingressClass | string | `nil` | When ingressClass is set, only Ingresses containing an annotation with the same value are processed. Otherwise, Ingresses missing the annotation, having an empty value, or the value traefik are processed. |
| providers.kubernetesIngress.namespaces | list | `[]` | Array of namespaces to watch. If left empty, Traefik watches all namespaces. . When using `rbac.namespaced`, it will watch helm release namespace and namespaces listed in this array. |
| providers.kubernetesIngress.nativeLBByDefault | bool | `false` | Defines whether to use Native Kubernetes load-balancing mode by default. |
| providers.kubernetesIngress.publishedService.enabled | bool | `true` | Enable [publishedService](https://doc.traefik.io/traefik/providers/kubernetes-ingress/#publishedservice) |
| providers.kubernetesIngress.publishedService.pathOverride | string | `""` | Override path of Kubernetes Service used to copy status from. Format: namespace/servicename. Default to Service deployed with this Chart. |
| providers.kubernetesIngress.strictPrefixMatching | bool | `false` | Defines whether to make prefix matching strictly comply with the Kubernetes Ingress specification. |
| rbac | object | `{"aggregateTo":[],"enabled":true,"namespaced":false,"secretResourceNames":[]}` | Whether Role Based Access Control objects like roles and rolebindings should be created |
| readinessProbe.failureThreshold | int | `1` | The number of consecutive failures allowed before considering the probe as failed. |
| readinessProbe.initialDelaySeconds | int | `2` | The number of seconds to wait before starting the first probe. |
| readinessProbe.periodSeconds | int | `10` | The number of seconds to wait between consecutive probes. |
| readinessProbe.successThreshold | int | `1` | The minimum consecutive successes required to consider the probe successful. |
| readinessProbe.timeoutSeconds | int | `2` | The number of seconds to wait for a probe response before considering it as failed. |
| resources | object | `{}` | [Resources](https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/) for `traefik` container. |
| securityContext | object | See _values.yaml_ | [SecurityContext](https://kubernetes.io/docs/reference/kubernetes-api/workload-resources/pod-v1/#security-context-1) |
| service.additionalServices | object | `{}` |  |
| service.annotations | object | `{}` | Additional annotations applied to both TCP and UDP services (e.g. for cloud provider specific config) |
| service.annotationsTCP | object | `{}` | Additional annotations for TCP service only |
| service.annotationsUDP | object | `{}` | Additional annotations for UDP service only |
| service.enabled | bool | `true` |  |
| service.externalIPs | list | `[]` |  |
| service.labels | object | `{}` | Additional service labels (e.g. for filtering Service by custom labels) |
| service.loadBalancerSourceRanges | list | `[]` |  |
| service.single | bool | `true` |  |
| service.spec | object | `{}` | Cannot contain type, selector or ports entries. |
| service.type | string | `"LoadBalancer"` |  |
| serviceAccount | object | `{"name":""}` | The service account the pods will use to interact with the Kubernetes API |
| serviceAccountAnnotations | object | `{}` | Additional serviceAccount annotations (e.g. for oidc authentication) |
| startupProbe | object | `{}` | Define [Startup Probe](https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/#define-startup-probes) |
| tlsOptions | object | `{}` | TLS Options are created as [TLSOption CRDs](https://doc.traefik.io/traefik/https/tls/#tls-options) When using `labelSelector`, you'll need to set labels on tlsOption accordingly. See EXAMPLE.md for details. |
| tlsStore | object | `{}` | TLS Store are created as [TLSStore CRDs](https://doc.traefik.io/traefik/https/tls/#default-certificate). This is useful if you want to set a default certificate. See EXAMPLE.md for details. |
| tolerations | list | `[]` | Tolerations allow the scheduler to schedule pods with matching taints. |
| topologySpreadConstraints | list | `[]` | You can use topology spread constraints to control how Pods are spread across your cluster among failure-domains. |
| tracing | object | `{"addInternals":false,"capturedRequestHeaders":[],"capturedResponseHeaders":[],"otlp":{"enabled":false,"grpc":{"enabled":false,"endpoint":"","insecure":false,"tls":{"ca":"","cert":"","insecureSkipVerify":false,"key":""}},"http":{"enabled":false,"endpoint":"","headers":{},"tls":{"ca":"","cert":"","insecureSkipVerify":false,"key":""}}},"resourceAttributes":{},"safeQueryParams":[],"sampleRate":null,"serviceName":null}` | https://doc.traefik.io/traefik/observability/tracing/overview/ |
| tracing.addInternals | bool | `false` | Enables tracing for internal resources. Default: false. |
| tracing.capturedRequestHeaders | list | `[]` | Defines the list of request headers to add as attributes. It applies to client and server kind spans. |
| tracing.capturedResponseHeaders | list | `[]` | Defines the list of response headers to add as attributes. It applies to client and server kind spans. |
| tracing.otlp.enabled | bool | `false` | See https://doc.traefik.io/traefik/v3.0/observability/tracing/opentelemetry/ |
| tracing.otlp.grpc.enabled | bool | `false` | Set to true in order to send metrics to the OpenTelemetry Collector using gRPC |
| tracing.otlp.grpc.endpoint | string | `""` | Format: <host>:<port>. Default: "localhost:4317" |
| tracing.otlp.grpc.insecure | bool | `false` | Allows reporter to send metrics to the OpenTelemetry Collector without using a secured protocol. |
| tracing.otlp.grpc.tls.ca | string | `""` | The path to the certificate authority, it defaults to the system bundle. |
| tracing.otlp.grpc.tls.cert | string | `""` | The path to the public certificate. When using this option, setting the key option is required. |
| tracing.otlp.grpc.tls.insecureSkipVerify | bool | `false` | When set to true, the TLS connection accepts any certificate presented by the server regardless of the hostnames it covers. |
| tracing.otlp.grpc.tls.key | string | `""` | The path to the private key. When using this option, setting the cert option is required. |
| tracing.otlp.http.enabled | bool | `false` | Set to true in order to send metrics to the OpenTelemetry Collector using HTTP. |
| tracing.otlp.http.endpoint | string | `""` | Format: <scheme>://<host>:<port><path>. Default: https://localhost:4318/v1/tracing |
| tracing.otlp.http.headers | object | `{}` | Additional headers sent with metrics by the reporter to the OpenTelemetry Collector. |
| tracing.otlp.http.tls.ca | string | `""` | The path to the certificate authority, it defaults to the system bundle. |
| tracing.otlp.http.tls.cert | string | `""` | The path to the public certificate. When using this option, setting the key option is required. |
| tracing.otlp.http.tls.insecureSkipVerify | bool | `false` | When set to true, the TLS connection accepts any certificate presented by the server regardless of the hostnames it covers. |
| tracing.otlp.http.tls.key | string | `""` | The path to the private key. When using this option, setting the cert option is required. |
| tracing.resourceAttributes | object | `{}` | Defines additional resource attributes to be sent to the collector. |
| tracing.safeQueryParams | list | `[]` | By default, all query parameters are redacted. Defines the list of query parameters to not redact. |
| tracing.sampleRate | string | `nil` | The proportion of requests to trace, specified between 0.0 and 1.0. Default: 1.0. |
| tracing.serviceName | string | `nil` | Service name used in selected backend. Default: traefik. |
| updateStrategy.rollingUpdate.maxSurge | int | `1` |  |
| updateStrategy.rollingUpdate.maxUnavailable | int | `0` |  |
| updateStrategy.type | string | `"RollingUpdate"` | Customize updateStrategy of Deployment or DaemonSet |
| versionOverride | string | `""` | This field overrides the default version extracted from image.tag |
| volumes | list | `[]` | Add volumes to the traefik pod. The volume name will be passed to tpl. This can be used to mount a cert pair or a configmap that holds a config.toml file. After the volume has been mounted, add the configs into traefik by using the `additionalArguments` list below, eg: `additionalArguments: - "--providers.file.filename=/config/dynamic.toml" - "--ping" - "--ping.entrypoint=web"` |

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.14.2](https://github.com/norwoodj/helm-docs/releases/v1.14.2)
