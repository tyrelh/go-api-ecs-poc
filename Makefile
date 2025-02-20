PROJECT_NAME=go-api-poc
GO_API_VERSION=$(shell cat version)
BUILD_DIR=bin
AWS_REGION?=ca-west-1
# ECR_REGISTRY=784593521445.dkr.ecr.us-west-2.amazonaws.com
ECR_REGISTRY=784593521445.dkr.ecr.${AWS_REGION}.amazonaws.com
# ECR_REPOSITORY=go-api-poc-repository
ECR_REPOSITORY=go-api-poc-ecr-repository
ECS_CLUSTER=go-api-poc-cluster
ECS_SERVICE=go-api-poc-ecs-service

# docker-dev:
# 	docker build -t ${PROJECT_NAME}:${GO_API_VERSION} --target development --build-arg GO_API_VERSION=${GO_API_VERSION} --build-arg GO_API_AWS_REGION=${AWS_REGION} .
# 	docker run --rm --name go-api-poc-local-dev -p 8080:8080 ${PROJECT_NAME}:${GO_API_VERSION}

deps:
	@echo "Checking for global tool dependencies..."
	@if ! which brew &> /dev/null; then \
		echo "Installing Brew..."; \
		curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh | sh; \
		echo "✅ Brew installed"; \
	else \
		echo "🟢 Brew is already installed"; \
	fi
	@if ! which go &> /dev/null; then \
		echo "Installing Go..."; \
		brew install golang; \
		echo "✅ Go installed"; \
	else \
		echo "🟢 Go is already installed"; \
	fi
	@if ! which docker &> /dev/null; then \
		echo "Installing Docker..."; \
		brew install --cask docker; \
		echo "✅ Docker installed"; \
	else \
		echo "🟢 Docker is already installed"; \
	fi
	@if ! which oapi-codegen &> /dev/null; then \
		echo "Installing oapi-codegen..."; \
		go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest; \
		echo "✅ oapi-codegen installed"; \
	else \
		echo "🟢 oapi-codegen is already installed"; \
	fi
	@if ! which air &> /dev/null; then \
		echo "Installing Air..."; \
		go install github.com/air-verse/air@latest; \
		echo "✅ Air installed"; \
	else \
		echo "🟢 Air is already installed"; \
	fi
	@if ! which dlv &> /dev/null; then \
		echo "Installing Delve..."; \
		go install github.com/go-delve/delve/cmd/dlv@latest; \
		echo "✅ Delve installed"; \
	else \
		echo "🟢 Delve is already installed"; \
	fi
	@if ! which sqlc &> /dev/null; then \
		echo "Installing sqlc..."; \
		go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest; \
		echo "✅ sqlc installed"; \
	else \
		echo "🟢 sqlc is already installed"; \
	fi
	@echo "Downloading project dependencies..."
	go mod download
	@echo "✅ All dependencies are installed."

oapi:
	@echo "Generating OpenAPI code..."
	oapi-codegen --config=api/oapi-codegen.yml api/api.yml
	@echo "✅ OpenAPI code generated." && echo

dev: oapi
	@echo "Starting DB container..."
	@docker compose up -d
	@echo "Waiting for DB container to be healthy..."
	@while [ "$$(docker inspect --format='{{json .State.Health.Status}}' go-api-poc-db | jq -r .)" != "healthy" ]; do \
		echo "Waiting for go-api-poc-db to be healthy..."; \
		sleep 1; \
	done
	@sleep 0.5
	@echo "✅ DB container is runing and healthy." && echo
	@echo "Starting local dev using Air..."
	@echo "⚠️ Air won't hot-reload changes to the OpenAPI spec api/api.yml. You'll need to rerun make dev or make api"
	@echo "ℹ️ You can connect to the debugger at 127.0.0.1:2345"
	@air

build: oapi
	@echo "##### BUILD #####" && echo
	@echo "Building version ${GO_API_VERSION} to ${BUILD_DIR}/${PROJECT_NAME}..."
	go mod tidy
	export GOOS=linux && \
	export CGO_ENABLED=0 && \
	go build -o ${BUILD_DIR}/${PROJECT_NAME} .
	@echo "🟢 Build complete." && echo

build-image: build
	@echo "##### BUILD IMAGE #####" && echo
	@echo "Building version ${GO_API_VERSION}..."
	docker build -t ${PROJECT_NAME}:${GO_API_VERSION} --target production --build-arg GO_API_VERSION=${GO_API_VERSION} --build-arg GO_API_AWS_REGION=${AWS_REGION} .
	@echo "🟢 Build complete." && echo

push: build-image
	@echo "##### PUSH #####" && echo
	@echo "Logging in to ECR..."
	aws ecr get-login-password --region ${AWS_REGION} --profile infrastructure-admin-dev | docker login --username AWS --password-stdin ${ECR_REGISTRY}
	@echo "🟢 Logged in to ECR." && echo
	@echo "Tagging image..."
	docker tag ${PROJECT_NAME}:${GO_API_VERSION} ${ECR_REGISTRY}/${ECR_REPOSITORY}:${GO_API_VERSION}
	docker tag ${PROJECT_NAME}:${GO_API_VERSION} ${ECR_REGISTRY}/${ECR_REPOSITORY}:latest
	@echo "🟢 Tagging complete." && echo
	@echo "Pushing image to ECR..."
	docker push ${ECR_REGISTRY}/${ECR_REPOSITORY}:${GO_API_VERSION}
	docker push ${ECR_REGISTRY}/${ECR_REPOSITORY}:latest
	@echo "🟢 Push complete." && echo

deploy: push
	@echo "##### DEPLOY #####" && echo
	@echo "Deploying ${GO_API_VERSION} to ECS service..."
	DEPLOY_RESPONSE=$$(aws ecs update-service --cluster ${ECS_CLUSTER} --service ${ECS_SERVICE} --force-new-deployment --profile infrastructure-admin-dev --region ${AWS_REGION}) && echo "$${DEPLOY_RESPONSE}"
	@echo "🟢 Deployment initiated." && echo