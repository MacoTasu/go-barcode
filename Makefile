.PHONY: setup install-tools install-lefthook init-lefthook help

setup: install-tools install-lefthook init-lefthook ## Setup all development tools and configurations

install-tools: ## Install required Go tools
	@echo "Installing development tools..."
	go install golang.org/x/tools/cmd/godoc@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/evilmartians/lefthook@latest

install-lefthook: install-tools ## Install lefthook
	@echo "Installing lefthook..."
	$(shell go env GOPATH)/bin/lefthook install

init-lefthook: install-lefthook ## Initialize lefthook configuration
	@echo "Initializing lefthook..."
	@if [ ! -f lefthook.yml ]; then \
		echo "Creating lefthook.yml..."; \
		cp lefthook.yml.example lefthook.yml 2>/dev/null || echo "lefthook.yml already exists"; \
	fi

help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' 