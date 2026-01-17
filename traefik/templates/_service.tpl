{{- define "traefik.service-name" -}}
{{- $fullname := printf "%s-%s" (include "traefik.fullname" .root) .name -}}
{{- if eq .name "default" -}}
{{- $fullname = include "traefik.fullname" .root -}}
{{- end -}}

{{- if ge (len $fullname) 60 -}} # 64 - 4 (udp-postfix) = 60
  {{- fail "ERROR: Cannot create a service whose full name contains more than 60 characters" -}}
{{- end -}}

{{- $fullname -}}
{{- end -}}

{{- define "traefik.service-metadata" }}
  labels:
  {{- include "traefik.labels" .root | nindent 4 -}}
  {{- with .service.labels }}
  {{- toYaml . | nindent 4 }}
  {{- end }}
{{- end }}

{{- define "traefik.service-spec" -}}
  {{- $type := default "LoadBalancer" .service.type }}
  type: {{ $type }}
  {{- with .service.loadBalancerClass }}
  loadBalancerClass: {{ . }}
  {{- end}}
  {{- with .service.spec }}
  {{- toYaml . | nindent 2 }}
  {{- end }}
  selector:
    {{- include "traefik.labelselector" .root | nindent 4 }}
  {{- if eq $type "LoadBalancer" }}
  {{- with .service.loadBalancerSourceRanges }}
  loadBalancerSourceRanges:
  {{- toYaml . | nindent 2 }}
  {{- end -}}
  {{- end -}}
  {{- with .service.externalIPs }}
  externalIPs:
  {{- toYaml . | nindent 2 }}
  {{- end -}}
  {{- with .service.ipFamilyPolicy }}
  ipFamilyPolicy: {{ . }}
  {{- end }}
  {{- with .service.ipFamilies }}
  ipFamilies:
  {{- toYaml . | nindent 2 }}
  {{- end -}}
{{- end }}

{{- define "traefik.service-ports" }}
 {{- range $portName, $config := .ports }}
  {{- $name := $portName | lower -}}
  {{- if (index (default dict $config.expose) $.serviceName) }}
  {{- $port := default $config.port $config.exposedPort }}
  {{- if empty $port }}
    {{- fail (print "ERROR: Cannot create " (trim $name) " port on Service without .port or .exposedPort") }}
  {{- end }}
  - port: {{ $port }}
    name: {{ include "traefik.portname" $name }}
    targetPort: {{ default $name $config.targetPort | include "traefik.portreference" }}
    protocol: {{ default "TCP" $config.protocol }}
    {{- if $config.nodePort }}
    nodePort: {{ $config.nodePort }}
    {{- end }}
    {{- if $config.appProtocol }}
    appProtocol: {{ $config.appProtocol }}
    {{- end }}
  {{- if and ($config.http3).enabled ($config.single) }}
  {{- $http3Port := default $config.exposedPort $config.http3.advertisedPort }}
  - port: {{ $http3Port }}
    name: {{ printf "%s-http3" $name | include "traefik.portname" }}
    targetPort: {{ default $name $config.targetPort | include "traefik.portreference" }}
    protocol: UDP
    {{- if $config.nodePort }}
    nodePort: {{ $config.nodePort }}
    {{- end }}
    {{- if $config.appProtocol }}
    appProtocol: {{ $config.appProtocol }}
    {{- end }}
   {{- end }}
  {{- end }}
 {{- end }}
{{- end }}
