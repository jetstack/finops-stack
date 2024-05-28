# FinOps Stack

[Visit finops-stack.jetstack.io](https://finops-stack.jetstack.io) for the full site and comprehensive details.

## Overview

The FinOps Stack is the blueprint for a solution to automate FinOps best practices. It integrates a suite of open-source tools into a unified, easy-to-install platform.

Our goal is to empower organisations with the tools they need to manage, visualise, and optimise their cloud resources in complex, ever-changing environments.

The FinOps Stack is designed to work out-of-the-box seamlessly with GKE standard/autopilot clusters using Google Managed Prometheus, and can be customised for an organisation’s business requirements and/or Kubernetes distribution.

This repository contains the core components of FinOps Stack, including:

- **Static Site**: The source code for the FinOps Stack website, providing documentation, guides, and resources for users.
- **Helm Charts** (`charts` directory): Helm charts included as part of the FinOps Stack. Charts are included in this directory so they can be used as blueprints and customised/extended where required.
- **Installation**: Configuration and documentation for installing the FinOps Stack.

## Installation and Usage

For ease of use, the FinOps Stack is installed using a single Helmfile command. For detailed installation instructions, review the README under the `/installation` directory of this repository. 

## Contributing

We welcome contributions to FinOps Stack! If you have an idea for a new feature, an improvement, or a bug fix, please consider raising an issue first to discuss your plans. This helps avoid duplicate work and ensures your contribution aligns with the project's goals.

When contributing, please follow these guidelines:

1. **Raise an Issue**: Before starting work, check for existing issues. If none exist, raise a new one to discuss your idea.
2. **Follow Templates**: Use the provided issue and pull request templates to ensure your contribution is well-documented and easy to review.
3. **Suffix "Chart" for Chart Updates**: If your contribution involves changes to the Helm charts, please suffix the title of your issue or pull request with "Chart" to help us identify it quickly.

## License

This project is licensed under the Apache 2.0 License. See the [LICENSE](LICENSE) file for more details.
