{{/*
Expand the name of the chart.
*/}}
{{- define "ingress-config.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "ingress-config.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "ingress-config.labels" -}}
helm.sh/chart: {{ include "ingress-config.chart" . }}
{{ include "ingress-config.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "ingress-config.selectorLabels" -}}
app.kubernetes.io/name: {{ include "ingress-config.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}