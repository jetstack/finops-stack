---
title: Quickstart
---

## Installation

To simplify installation, the FinOps Stack is installed using a single Helmfile command. 

The following instructions are designed to work with a Kind cluster for quick setup. For full instructions, prerequisites and customisations, please see the [installation README](https://github.com/jetstack/finops-stack/blob/main/installation/README.md).

To work with GKE standard and GKE autopilot see the [ GKE installation guide]((https://github.com/jetstack/finops-stack/blob/main/installation/README.md))

### Helmfile 

1. Clone the FinOps Stack GitHub repository: 
```shell
git clone https://github.com/jetstack/finops-stack.git
```

2. In the [installation](https://github.com/jetstack/finops-stack/blob/main/installation/) directory, Copy env.tmpl file and replace the env var values accordingly (`GRAFANA_FQDN` for example).
```shell
cp ./env.tmpl ./.env
```

3. Install using helmfile:

```shell
set -a; source .env; set +a; helmfile apply --interactive
```
