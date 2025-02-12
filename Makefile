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
	@echo "##### DEPS #####" && echo
	@echo "Installing dependencies..."
	go mod download
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
	go install github.com/air-verse/air@latest
	go install github.com/go-delve/delve/cmd/dlv@latest
	@echo "🟢 Dependencies installed." && echo

oapi:
	@echo "Generating OpenAPI code..."
	oapi-codegen --config=api/oapi-codegen.yml api/api.yml
	@echo "🟢 OpenAPI code generated." && echo

dev: oapi
	@docker compose up -d
	@sleep 8
	@echo "Local dev using Air"
	@echo "⚠️ Air won't hot-reload changes to the OpenAPI spec api/api.yml. You'll need to rerun make dev or make api."
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