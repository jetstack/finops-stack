# Alerting Selection

**Capability**: [Anomaly Management](https://www.finops.org/framework/capabilities/anomaly-management/)

There appears to be little to no open-source tooling that specifically helps with monitoring and alerting. In this case, options are to either make use of existing monitoring and alerting tools, or to build custom alerting.

## Tooling considered

### Prometheus / Alertmanager

This is a logical choice as a leading open-source tool for Kubernetes metrics and alerting, and is also a prerequisite for OpenCost installation.

Custom tooling would be unnecessarily complicated in this case. Instead, we should consider offering some useful metrics/prometheus rules to alert on.
