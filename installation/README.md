# Installation using Helmfile

Installing Helm charts with lots of dependencies and CRDs is challenging; these instructions use Helmfile to mitigate issues with Helm.

## Pre-requisites

- A GKE cluster and kubectl access to it with cluster-admin access.
- [Helmfile](https://helmfile.readthedocs.io/en/latest/#installation) installed on your local machine
- A GSA with roles/monitor.viewer, roles/iam.serviceAccountTokenCreator, workload identity for KSA: `[finops-stack/grafana]`

## Distribution support

### GKE Autopilot

- Enable cost allocation

## Installation

### Configuration changes for your cluster environment

1. To control which Finops Stack components to install edit the `./installation/config/common/enabled.yaml` file
1. Copy `./env.tmpl` to `./.env` and replace the env var values accordingly (at the very least you'll need to change CSP_API_KEY and GRAFANA_SA_ANNOTATION values)

### Install everything using Helmfile

For the first run:

```bash
set -a; source .env; set +a; helmfile apply --interactive
```

To speed up subsequent runs:

```bash
set -a; source .env; set +a; helmfile apply --interactive --skip-deps
```

### Manually install Envoy Proxy

Envoy Proxy is used as a proxy for accessing GMP (it injects a valid Bearer auth token into the request). Currently this needs to be installed by `kubectl` as there's not a Helm chart yet. The manifest is `./gmp-proxy.yaml`; you'll need to edit it to change the GCP Project - search for this line: `substitution: "/v1/projects/jetstack-steve-judd/location/global/prometheus/api/v1\\1"`