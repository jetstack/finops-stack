# Monitoring Selection

**Capability**: [Allocation](https://www.finops.org/framework/capabilities/allocation/), [Reporting and Analytics](https://www.finops.org/framework/capabilities/reporting-analytics/)

## **Aims to understand**: What are we spending where?

Should cover:

- Spend
- Cost-allocation
- Gathering and visualising this data

## Tooling considered

### [OpenCost](https://www.opencost.io/)

"A vendor-neutral open source project for measuring and allocating cloud infrastructure and container costs in real time".
Community recognition: A CNCF Sandbox Project, and part of FinOps Foundation's Landscape, this project has strong credibility in this space.
Requires Prometheus to be installed.

**License**: Apache
**Level of Maintenance**: actively maintained.

**Provides**: metrics and visibility.
Does not provide: recommendations/reports (these are in Kubecost free).
Requires user configuration to get started (can't install and get running without a GCP api key if using google cloud - autodetected).
**Installation**: first Prometheus, then Cloud provider configuration, then opencost.

### [KubeCost](https://www.kubecost.com/) (free tier)

KubeCost is the commercial version of OpenCost. It is an open core product, and the free version of KubeCost can be installed using their cost-analyzer helm-chart. This chart is open-source, [https://github.com/kubecost/cost-analyzer-helm-chart].
**License**: Apache
**Level of Maintenance**: actively maintained.
**Question**: do you need to sign up for an API key to keep using the free version?
Full comparison of KubeCost vs OpenCost is [here](https://docs.kubecost.com/architecture/opencost-product-comparison).
Grafana and Prometheus are bundled with the install.
Get data/running ui on install without any preconfiguration.

### [Crane](https://gocrane.io/)

"A FinOps Platform for Cloud Resource Analytics and Economics in Kubernetes clusters."
Community recognition: On CNCF Landscape, not a CNCF Project, not on FinOps Foundation's Landscape.
**License**: Apache
**Level of Maintenance**: last release Jul 2023. Does not appear actively maintained.

### [Komiser](https://www.komiser.io/)

"A cloud-agnostic resource mananger... integrates with multiple cloud providers … builds a cloud asset inventory, and helps you break down your cost at the resource level 💰"
Community recognition: Not on CNCF Landscape, not a CNCF Project, not on FinOps Foundation's Landscape.
Has community following: 4.8k stars on GitHub.
**License**: Elastic License (ELv2).
**Level of Maintenance**: actively maintained.

### Custom Tooling
Self-calculation of costs based on data from publicly available cloud provider costing APIs is an option. However, existing tools (OpenCost and Crane) are already able to calculate and visualise these costs.
So custom tooling would be an unacceptable/unnecessary overhead.
