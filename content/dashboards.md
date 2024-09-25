# Grafana dashboards & Prometheus

Depending on how you are runnning Prometheus and Grafana, you might need to tweak the PromQL used in the various dashboards provided in [our chart](https://github.com/jetstack/finops-stack/tree/main/charts/opencost-config/dashboards). This page contains notes and troubleshooting tips based on our experiences.

## Prometheus

### Duplicated labels

The Opencost metrics use the `instance` and `namespace` labels to identify the _pod and namespace to which the metric relates_, i.e. _not_ the Opencost Prometheus Exporter pod. Depending on how your Prometheus service is configured or if you are using Google Managed Prometheus (GMP) the values of these labels could be overwritten, namely:

- GMP replaces the `instance` and `namespace` values with the name and namespace of the pod the metrics were scraped from. In our case this is always the Exporter pod in the `finops-stack` namespace.

- In standalone Prometheus the same behaviour as above will occur unless config value `honor_labels` has been set to true.

In both cases the original values of the namespace and instance labels will be copied to `exported_namespace` and `exported_instance` respectively.

__NOTE:__ Our dashboards use the `exported_*` labels as GMP does not support the `honor_labels` configuration.

### Query execution time

Several of the PromQL queries use group_left 'joins'; be aware these can get expensive and time consuming if you select very broad time ranges.

### Accessing Google Managed Prometheus APIs

**Only applicable if you are using GKE + Google Managed Prometheus**

GMP provides endpoints for almost all of Prometheus' API (as documented [here](https://cloud.google.com/stackdriver/docs/managed-prometheus/query-api-ui#http-api-details)); like all of GCP's APIs they require a valid Bearer token in the Authorization header. Thus, the Grafana GMP data source needs a valid token. Google do provide 2 suggestions for automating this process ([here](https://cloud.google.com/stackdriver/docs/managed-prometheus/query-api-ui)). Neither solution is simple and both have downsides, so we provide an alternative via the [`gmp-proxy` chart](https://github.com/jetstack/finops-stack/tree/main/charts/gmp-proxy) in the `finops-stack` repo. This chart installs an Envoy Proxy workload which is configured to add a valid Bearer token to each GMP API request before forwarding it to the GMP endpoint.

#### Usage

##### Pre-requisites

The GMP Proxy Pod uses GCP Workload Identity so you'll need to associate the SA used by the Pod with a GCP SA that has the following roles:

- monitor.viewer
- iam.serviceAccountTokenCreator

##### Envoy Proxy image

The standard Envoy Docker image can't be used as is because an additional Lua library is required, so we have provided a custom image which includes this library. If you prefer to use your own image, this is `Dockerfile` we used:

```
# Image from https://hub.docker.com/r/envoyproxy/envoy
FROM envoyproxy/envoy:v1.31-latest
RUN  apt update && apt install -y luarocks
RUN  luarocks install lua-cjson
```

## Dashboards

Currently the working and tested dashboards are:

### FinOps Stack: Cluster cost & efficiency

The purpose of this dashboard is to provide an overview of the cost of your cluster and namespaces using Opencost's calculations. The time range control is disabled and you can only select from 3 time periods using a drop down in the top left. This is because some of the queries are expensive and with too long a time period cause the query request to timeout.

### FinOps Stack: Used | Wasted Resources for cluster & namespace

The purpose of this dashboard is to provide details of resource usage and wastage at the cluster, namespace and pod level. You can select different time periods using the standard Grafana time range control in the top right.
