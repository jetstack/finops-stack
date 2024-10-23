{{/*
Expand the name of the chart.
*/}}
{{- define "cert-manager-config.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}
{{- define "cert-manager-config.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Common labels
*/}}
{{- define "cert-manager-config.labels" -}}
helm.sh/chart: {{ include "cert-manager-config.chart" . }}
{{ include "cert-manager-config.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "cert-manager-config.selectorLabels" -}}
app.kubernetes.io/name: {{ include "cert-manager-config.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}