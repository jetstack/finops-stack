# finops-stack

![Version: 0.0.2](https://img.shields.io/badge/Version-0.0.2-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 0.0.2](https://img.shields.io/badge/AppVersion-0.0.2-informational?style=flat-square)

A FinOps Stack for Kubernetes

## Requirements

| Repository | Name | Version |
|------------|------|---------|
| file://../../autospot/chart/ | autospot | * |
| file://../../finops-policies/chart/ | finops-policies | * |
| file://../../limit-ranger/chart/ | limit-ranger | * |
| file://../../office-hours/chart/ | office-hours | * |
| file://../../vpatron/chart/ | vpatron | * |
| https://charts.fairwinds.com/stable | vpa | * |
| https://grafana.github.io/helm-charts | grafana | * |
| https://kyverno.github.io/kyverno/ | kyverno | * |

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| feature.autospot | bool | `false` | Enable / Disable the AutoSpot controller |
| feature.finops-policies | bool | `false` | Enable / Disable the installation of Kyverno FinOps Policies |
| feature.grafana | bool | `true` | Enable / Disable the installation of the Grafana |
| feature.kyverno | bool | `true` | Enable / Disable the installation of Kyverno |
| feature.limit-ranger | bool | `true` | Enable / Disable the limit-ranger controller |
| feature.office-hours | bool | `true` | Enable / Disable the office-hours controller |
| feature.vpa | bool | `false` | Enable / Disable the installation of the VPA Controller |
| feature.vpatron | bool | `true` | Enable / Disable the vpatron controller |

----------------------------------------------
Autogenerated from chart metadata using [helm-docs v1.13.1](https://github.com/norwoodj/helm-docs/releases/v1.13.1)