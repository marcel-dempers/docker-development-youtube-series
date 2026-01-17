{{- define "traefik.metrics-service-metadata" }}
  labels:
  {{- include "traefik.metricsservicelabels" . | nindent 4 -}}
  {{- with .Values.metrics.prometheus.service.labels }}
  {{- toYaml . | nindent 4 }}
  {{- end }}
{{- end }}

{{/* Labels used for metrics-relevant selector*/}}
{{/* This is an immutable field: this should not change between upgrade */}}
{{- define "traefik.metricslabelselector" -}}
{{- include "traefik.labelselector" . }}
app.kubernetes.io/component: metrics
{{- end }}

{{/* Shared labels used in metadata of metrics-service and servicemonitor */}}
{{- define "traefik.metricsservicelabels" -}}
{{ include "traefik.metricslabelselector" . }}
helm.sh/chart: {{ template "traefik.chart" . }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- with .Values.commonLabels }}
{{ toYaml . }}
{{- end }}
{{- end }}

