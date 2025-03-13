# The OpenStack Architecture

FinOps Stack makes use of a wide range of software to provide a highly efficient Stack:

- **[Kyverno](https://kyverno.io/)**: A Kubernetes-native policy engine that helps enforce, validate, and mutate resource configurations, allowing users to define and manage security, operational, and compliance policies directly as Kubernetes resources. In the FinOps Stack, Kyverno is used to audit or enforce FinOps best practices.
- **[OpenCost](https://www.opencost.io/)**: An open-source cost monitoring and management tool for Kubernetes, designed to provide real-time visibility into the costs associated with workloads, resources, and cloud infrastructure usage in a Kubernetes cluster. In the Stack, we use the OpenCost Prometheus exporter to gather cost metrics, and send these to Prometheus or Google Managed Prometheus when using a GKE cluster. 
- **[Goldilocks](https://www.fairwinds.com/goldilocks)**: Fairwinds Goldilocks is designed to provide 'just-right' recommendations for container requests and limits. It does this by using the Vertical Pod Autoscaler in recommendation mode. Additional guidance on interpresting and setting resource requests and limits can be found in our [GKE Autopilot guide](./distribution-gke-autopilot.md).
- **[Vertical Pod Autoscaler](https://github.com/kubernetes/autoscaler/blob/master/vertical-pod-autoscaler/README.md)**: frees users from the necessity of setting up-to-date resource limits and requests for the containers in their pods. When configured, it will set the requests automatically based on usage and thus allow proper scheduling onto nodes so that appropriate resource amount is available for each pod.
- **[Grafana](https://grafana.com/grafana/)**: An open-source platform used for monitoring and visualising time-series data from various sources, allowing users to create interactive and real-time dashboards for metrics, logs, and analytics. The FinOps Stack uses Grafana to provide visualisations of cost and efficiency metrics.
- **gmp-proxy**: A proxy designed to simplify the integration between open-source Grafana and Google Managed Prometheus (if using a GKE cluster).
- **[Cert-Manager](https://cert-manager.io/)**: An open source tool that provides certificate management for Kubernetes. In the FinOps Stack, cert-manager can be optionally installed to create a tls certificate if using ingress for Grafana. 


