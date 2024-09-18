{{/*
Expand the name of the chart.
*/}}
{{- define "cert-manager-config.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}
{{- define "cert-manager-config.dashboardname" -}}
{{ base .path | trimSuffix (ext .path) | lower | replace " " "-" | trunc 63 | trimSuffix "-" }}
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