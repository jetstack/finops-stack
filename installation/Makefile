.DEFAULT_GOAL := help

cluster: ## Setup cluster
	kind create cluster

finops-stack: ## Deploy FinOps Stack using Helmfile
	set -a; source .env; set +a; helmfile apply --file helmfile_kind.yaml --interactive

cleanup: ## Cleaup
	kind delete cluster

help: ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m \t%s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
