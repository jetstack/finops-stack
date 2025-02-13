# FinOps Stack deployment

This documentation provides instructions for installing the FinOps Stack in Kind cluster for a quick setup.

For deployment on a GKE cluster, refer to the [GKE docs](./gke.md) and deployment on a EKS cluster refer to the [EKS docs](./eks.md).

## Using Helmfile

Installing Helm charts with lots of dependencies and CRDs is challenging; these instructions use Helmfile to mitigate issues with Helm.

## Pre-requisites

- [Kind](https://kind.sigs.k8s.io/docs/user/quick-start/#installation) installed on your local machine
- [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)
- [Helmfile](https://helmfile.readthedocs.io/en/latest/#installation) installed on your local machine

## Installation

## Create a kind cluster

```bash
make cluster
```

### Configuration changes for your cluster environment

1. To control which Finops Stack components to install, edit the [enabled.yaml](./installation/config/common/enabled.yaml) file
1. Copy env.tmpl file and replace the env var values accordingly (`GRAFANA_FQDN` for example).

```sh
cp ./env.tmpl ./.env
```

### Install everything using Helmfile

For the first run:

```bash
make finops-stack
# FinOps stack is install using Helmfile:
# set -a; source .env; set +a; helmfile apply --file helmfile_kind.yaml --interactive
```

NOTE: it will take several minutes for all workloads to install and start running. Helmfile does display its progress in the terminal. All workloads get installed into the `finops-stack` namespace so you can also view progress using `kubectl`.

To speed up subsequent runs:

```bash
set -a; source .env; set +a; helmfile apply --interactive --skip-deps
```

## Optional: Configure ingress for Grafana

### Pre-requisites

Already have an FQDN setup and registered with a public IP, e.g. grafana.example.com

### Grafana Helm values

These are specified in `config/common/grafana-values.yaml`, `config/gke/grafana-values.yaml` and under the Grafana release in `helmfile.yaml`. Probably all the changes you will want to make can be done by changing the values in `helmfile.yaml`, e.g. the admin user and what type of ingress you require.

General guidance when configuring ingress:
- Update the `.env` file with the FQDN and public IP for you domain.
- If you wish to enable tls, then ensure that cert-manager.enabled is set to true and update the values in `.env` accordingly.

## Enable Goldilocks for namespaces

For Goldilocks to analyse namespaces and add then to its dashboard you need to add this label to the namespace resource: `goldilocks.fairwinds.com/enabled=true`, e.g:

```bash
kubectl label ns finops-stack goldilocks.fairwinds.com/enabled=true
```

## Useful commands

To port forward to Grafana:

```bash
kubectl --namespace finops-stack port-forward service/grafana 3000:80
```

Access via http://localhost:3000

To port forward to the metrics endpoint of the Opencost Prometheus exporter (to examine what metrics are being scraped):

```bash
kubectl --namespace finops-stack port-forward service/prometheus-opencost-exporter 9003:9003
```

To access the Goldilocks dashboard (assuming you've enabled it):

```bash
kubectl -n finops-stack port-forward svc/goldilocks-dashboard 8080:80
```

Then goto http://localhost:8080