# Overview

## The FinOps Stack

The FinOps Stack is the blueprint for a solution to automate FinOps best practices. It integrates a suite of open-source tools into a unified, easy-to-install platform.

Our goal is to empower organisations with the tools they need to manage, visualise, and optimise their cloud resources in complex, ever-changing environments.

The FinOps Stack is designed to work out-of-the-box seamlessly with GKE standard/autopilot clusters using Google Managed Prometheus, and can be customised for an organisation’s business requirements and/or Kubernetes distribution.

## Key Features and Capabilities

### Financial Dashboards

The FinOps Stack provides Grafana dashboards that visualise a cluster's financial health and efficiency, including:
- Cluster-level overviews
- Drill-down panels with detailed metrics on cost, efficiency, utilisation and waste.

### FinOps Policies

Kyverno policies in the Stack can be used in either audit or enforce mode, and are based on FinOps best practices. Included policy examples cover large container images, cost-allocation/cost-center labels, restricting the scale of deployments and statefulsets, automating generation of HPA resources, disallowing service loadbalancers and preventing naked pods. Policies can be used for reporting or policy enforcement.

More detail on each policy can be found in the dedicated Policies page. 

### Cluster Efficiency

Overall cluster efficiency, and especially right-sizing container requests and limits, is essential to running an optimised Kubernetes cluster. 

Goldilocks is included in the FinOps Stack to provide guidance on right-sized container requests and limits, and to complement the visualisations in the Stack's Grafana dashboard. 

In addition, see the distribution-gke-autopilot page of this website to see a detailed breakdown of considerations for achieving greater cluster efficiency in GKE Autopilot.