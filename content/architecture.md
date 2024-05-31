# The OpenStack Architecture

FinOps Stack makes use of a wide range of software to provide a highly efficient Stack:

- **[VPAtron](https://github.com/jetstack/finops-toolkit/tree/main/vpatron)**: Deploys a [VPA](https://github.com/kubernetes/autoscaler/tree/master/vertical-pod-autoscaler) for each Deployment, Statefulset and Cron Workloads
- **[Limit-Ranger](https://github.com/jetstack/finops-toolkit/tree/main/limit-ranger)**: Deploys a [LimitRange](https://kubernetes.io/docs/concepts/policy/limit-range/) into all (subscribed) namespaces with a customised set of Default CPU/Memory Requests.
- **[AutoSpot](https://github.com/jetstack/finops-toolkit/tree/main/autospot)**: Ensures workloads are deployed onto Spot Instances (GKE Only)
- **[Office Hours](https://github.com/jetstack/finops-toolkit/tree/main/office-hours)**: Scales down all workloads out of office hours
- **[Kyverno](https://kyverno.io/)**: For FinOps policy Enforcement and Recommendations to engineers
- **[OpenCost](https://www.opencost.io/)**: Provides a Cloud(and On-Prem) Agnostic costing and utilisation metrics
- **[Prometheus](https://prometheus.io/)**: Provides a highly scaleable metrics and observation platform
- **[Grafana](https://grafana.com/grafana/)**: Proved a first in class Observability platform

## What does this look like:

This Diagram describes what the FinOps Stack looks like within your cluster:
![Architecture Diagram of FinOpsStack](assets/architecture.png)
