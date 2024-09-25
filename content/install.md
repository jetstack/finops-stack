---
title: Quickstart
---

## Installation

To simplify installation, the FinOps Stack is installed using a single Helmfile command. 

The following instructions are designed to work with GKE standard and GKE autopilot. For full instructions, prerequisites and customisations, please see the [installation README](https://github.com/jetstack/finops-stack/blob/main/installation/README.md). 

### Helmfile 

1. Clone the FinOps Stack GitHub repository: 
```shell
git clone https://github.com/jetstack/finops-stack.git
```

2. In the `/installation` directory, copy `./env.tmpl` to `./.env`. Replace the env var values accordingly. As a minimum, you will need to change the GCP_PROJECT, CSP_API_KEY,  GRAFANA_SA_ANNOTATION values.

3. Install using helmfile:

```shell
set -a; source .env; set +a; helmfile apply --interactive
```
