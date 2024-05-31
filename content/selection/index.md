# Selection

There are an increasing number of paid-for solutions which offer partial or complete FinOps tooling.
However, there is much less maturity and choice in the open source space (for example, review the FinOps Landscape filtered by open source tools).
In some cases, tooling may have a free tier, but not be open source.

This Project aims to bring together an opinionated, easy-to-install collection of open-source tools that automate FinOps best practices within Kubernetes.
This tooling is Kubernetes-specific, and cloud provider agnostic.

## Factors when selecting tooling

Where possible, we aim to adopt existing, well-maintained open source tools. We prefer choosing open-source tools, as opposed to free-tier, as these are more universally available to the community without restriction. Where these do not exist, or where we believe there is more benefit available, we will consider building custom tooling.

Because we aim to make this tooling available under an Apache license, we prefer tooling without more restrictive license that would affect future use of the stack.

### Scope

This tooling is Kubernetes-specific, and cloud provider agnostic. This means that Kubernetes-related resources (e.g. Persistent Disks, Virtual Machines) are in-scope, whereas wider Cloud resources, and tooling that is cloud provider specific, is out-of-scope.

As this project is Kubernetes-specific, all tooling must be able to be installed and run within a Kubernetes cluster. This means that tooling run at other stages of the SLDC/CICD processes is out of scope.

### Who this project is for

This project is aimed at engineering teams that are aiming to automate FinOps best practices within Kubernetes using open source tooling. In FinOps terms, we would call this the [Engineering persona](https://www.finops.org/framework/personas/).

### What are the FinOps challenges this persona is trying to solve?

We believe that the key areas that the Engineering persona is trying to solve are:

- **Monitoring**
  - Gathering metrics/cost data
  - Visualizing this data to teams and stakeholders
  - Multi-Cloud cost data
- **Alerting**
  - Anomaly management
- **Automation of best practices** such as:
  - Right-sizing/workload optimization
  - Providing sensible defaults
  - Forecasting (how much will x cost(if…))?
  - Rate reduction
  - Cost avoidance

## What open-source tooling is available to solve these problems

### Summary of tooling choices

**Monitoring**: [[Allocation](https://www.finops.org/framework/capabilities/allocation/), [Reporting and Analytics](https://www.finops.org/framework/capabilities/reporting-analytics/) ]

- [OpenCost](https://www.opencost.io/) (with Prometheus and Grafana)
- Gap: What about reporting?

**Alerting**: [[Anomaly Management](https://www.finops.org/framework/capabilities/anomaly-management/)]

- Prometheus/Alertmanager (with custom prometheus rules?)

**Automation and tooling**:

- Admission control and policy enforcement: Kyverno (and custom tooling to add LimitRange per namespace) [[Cloud Policy & Governance](https://www.finops.org/framework/capabilities/cloud-policy-governance/), [Workload Optimization](https://www.finops.org/framework/capabilities/workload-optimization/)]
- Rate optimization: Gap: Recommendations.  Custom required for spot. No recommendation tooling available.
- Workload optimization:
  - Cost avoidance:

### Whats Missing

#### Problems that we are not (yet) solving in this stack

- Recommendation engine for rate-reduction and cost avoidance
- Reporting
- Normalization of cost data across clouds (could explore [https://focus.finops.org/](https://focus.finops.org/) in the future)
- Calculation of unit economics
- Recommendations for rate optimization/cost avoidance
- Forecasting
