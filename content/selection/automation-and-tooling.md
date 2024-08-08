# Automation and tooling Selection

## Policy enforcement/guard rails

**Capability**: [Cloud Policy & Governance](https://www.finops.org/framework/capabilities/cloud-policy-governance/), [Workload Optimisation](https://www.finops.org/framework/capabilities/workload-optimization/)

Admission control can provide enforcement of sensible defaults and guard rails through enforcement of sensible defaults, and enforcing policies that prevent expensive/anomalous misconfigurations.

There are several admission controllers that could be used for this purpose - so here this decision is based on existing/available reference policies as well as our opinion on which of these tools are most established from observed customer adoption.

Policies should cover:

- CPU and Memory requests and limits
- Cost allocation labels
- Possible: resource quotas
- LimitRange per namespace

### Tooling considered

#### [Kyverno](https://kyverno.io/)

[Nirmata](https://nirmata.com/) provide a collection of 'out-of-the-box' FinOps and Cost-Optimisation policies ([https://github.com/nirmata/finops-policies], [https://nirmata.com/2023/03/22/kubernetes-finops-policies-with-kyverno/]). These May require some additional policies - but this is a good starting point.
**Community recognition**: Not on FinOps landscape. CNCF Incubating Project.
**License**: Apache
**Level of Maintenance**: Actively maintained

#### [Gatekeeper](https://github.com/open-policy-agent/gatekeeper)

No cost/FinOps specific collections of policies - but library contains some useful policies (CPU requests and limits, container resource ratios etc). Additional policies would need writing to gain full benefit.
**Community recognition**: Not on FinOps landscape. On CNCF Landscape.
**License**: Apache
**Level of Maintenance**: Actively maintained

#### [Polaris](https://www.fairwinds.com/polaris)

Has a few out of the box policies related to efficiency - CPU and Memory requests and limits. Additional policies would need writing to gain full benefit.
**Community recognition**: On FinOps landscape. Not a CNCF Project.
**License**: Apache
**Level of Maintenance**: Actively maintained

## Rate optimisation

In other words - cloud discounts. Pay less for what you need.

Ideally tooling would:
Identify potential for discounts (e.g. Spot Instances or Commitments vs on-demand pricing)
Automate use of a cheaper machine type or machine discount model (i.e. Spot) where available

**Capability**: [Rate Optimisation](https://www.finops.org/framework/capabilities/rate-optimization/)

### Tooling considered

There does not appear to be open-source tooling that we can use to build recommendations. This therefore requires custom work.
For any custom work, we should think about preferred direction:

- Providing advice and recommendations. If we want to provide recommendations, we would need to build custom tooling to either make our own recommendations OR hook into cloud provider recommendation APIS (if /where available).
- Provide automation for rate optimisation as 'opt-in'. In reality, this would be spot instances only.

## Workload optimisation
**Capability**: [Workload Optimisation](https://www.finops.org/framework/capabilities/workload-optimization/)

Consider:
- Horizontal scaling
- Vertical scaling
- Out of hours

Tooling considered:
- [Goldilocks](https://www.fairwinds.com/goldilocks)
- VPA
- **Custom tooling with VPA**
- Out of hours
- HPA

Out of scope: Cloud Provider-specific (e.g. MDA)

## Cost avoidance

Identifying idle resources.
Deleting idle resources.
