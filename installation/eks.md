# Installation using Helmfile

Installing Helm charts with lots of dependencies and CRDs is challenging; these instructions use Helmfile to mitigate issues with Helm. 

This documentation focuses on installing the FinOps Stack in EKS clusters.

## Pre-requisites

- A EKS cluster with:
   - kubectl access 
   - (Optional) If your cluster has Spot Instances, EKS Pod Identities need to be configured. See [documentation](https://www.opencost.io/docs/configuration/aws#eks-pod-identities).
- [Helmfile](https://helmfile.readthedocs.io/en/latest/#installation) installed on your local machine
- Unless you want to access the Grafana dashboard via `kubectl port-forward` you'll need a domain name or external public  IP.

## Installation

### Configuration changes for your cluster environment

1. To control which Finops Stack components to install, edit the [enabled.yaml](./installation/config/common/enabled.yaml) file
1. Copy `./env_eks.tmpl` to `./.env` and replace the env var values accordingly.

### Install everything using Helmfile

For the first run:

```bash
set -a; source .env; set +a; helmfile apply --file Helmfile_eks.yaml --interactive
```

NOTE: it will take several minutes for all workloads to install and start running. Helmfile does display its progress in the terminal. All workloads get installed into the `finops-stack` namespace so you can also view progress using `kubectl`.

To speed up subsequent runs:

```bash
set -a; source .env; set +a; helmfile apply --file Helmfile_eks.yaml --interactive --skip-deps
```

## Optional: Making Grafana accessible via DNS

### Pre-requisites

Already have an FQDN setup and registered with a public IP, e.g. grafana.example.com

### Grafana Helm values

These are specified in `config/common/grafana-values.yaml`, `config/gke/grafana-values.yaml` and under the Grafana release in `helmfile.yaml`. Probably all the changes you will want to make can be done by changing the values in `helmfile.yaml`, e.g. the admin user and what type of ingress you require.

General guidance when configuring ingress:
- Update the `.env` file with the FQDN and public IP for you domain.
- If you wish to enable tls, then ensure that cert-manager.enabled is set to true and update the values in `.env` accordingly.

## Enable Goldilocks for namespaces

For Goldilocks to analyse namespaces and add then to its dashboard you need to add this label to the namespace resource: `goldilocks.fairwinds.com/enabled=true`, e.g.  
`kubectl label ns finops-stack goldilocks.fairwinds.com/enabled=true`

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