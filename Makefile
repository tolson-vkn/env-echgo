# Get the git tag from the current commit...
TAG=$(shell git describe --abbrev=0 --tags)
IMAGE=timmyolson/env-echgo
MAKEFILE_DIR=$(PWD)

.PHONY: help
help:
	@echo "+------------------+"
	@echo "| Makefile Targets |"
	@echo "+------------------+"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-12s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## Build Web.
	@echo "+---------------------+"
	@echo "| Building Containers |"
	@echo "+---------------------+"
	docker build -t $(IMAGE)

.PHONY: up 
up: ## Start development environment.
	@echo "Starting docker environment";
	docker run --rm --name env-echgo -i -t -p 8080:8080 $(IMAGE)

.PHONY: down
down: ## Bring down development environment.
	@echo "Stopping docker environment";
	docker stop env-echgo
	docker rm env-echgo

.PHONY: login
login: ## Login to Docker Hub
	@echo "Login to DockerHub."
	docker login

.PHONY: version
version: ## Make a release tag
	@echo "Tagging version."
	./scripts/semantic_version.sh

.PHONY: publish
publish: ## Publish to ECR
	@echo "Build and Publish"

	make build
	docker tag $(IMAGE) $(IMAGE):$(TAG)
	docker tag $(IMAGE) $(IMAGE):latest
	docker push $(IMAGE):$(TAG)
	docker push $(IMAGE):latest
