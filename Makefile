PROJECT_NAME=go-api-poc
VERSION=$(shell cat version)
BUILD_DIR=bin
ECR_REGISTRY=784593521445.dkr.ecr.us-west-2.amazonaws.com
ECR_REPOSITORY=go-api-poc-repository
ECS_CLUSTER=go-api-poc-cluster
ECS_SERVICE=go-api-poc-service

dev:
	@echo "Local dev using Air"
	@air

api:
	@echo "Generating OpenAPI code..."
	oapi-codegen --config=api/oapi-codegen.yml api/api.yml
	@echo "🟢 OpenAPI code generated." && echo

build: api
	@echo "##### BUILD #####" && echo
	@echo "Building version ${VERSION} to ${BUILD_DIR}/${PROJECT_NAME}..."
	export GOOS=linux && \
	export CGO_ENABLED=0 && \
	go build -o ${BUILD_DIR}/${PROJECT_NAME} .
	@echo "🟢 Build complete." && echo

build-image: build
	@echo "##### BUILD IMAGE #####" && echo
	@echo "Building version ${VERSION}..."
	docker build -t ${PROJECT_NAME}:${VERSION} --target production --build-arg VERSION=${VERSION} .
	@echo "🟢 Build complete." && echo

push: build-image
	@echo "##### PUSH #####" && echo
	@echo "Logging in to ECR..."
	aws ecr get-login-password --region us-west-2 --profile infrastructure-admin-dev | docker login --username AWS --password-stdin ${ECR_REGISTRY}
	@echo "🟢 Logged in to ECR." && echo
	@echo "Tagging image..."
	docker tag ${PROJECT_NAME}:${VERSION} ${ECR_REGISTRY}/${ECR_REPOSITORY}:${VERSION}
	docker tag ${PROJECT_NAME}:${VERSION} ${ECR_REGISTRY}/${ECR_REPOSITORY}:latest
	@echo "🟢 Tagging complete." && echo
	@echo "Pushing image to ECR..."
	docker push ${ECR_REGISTRY}/${ECR_REPOSITORY}:${VERSION}
	docker push ${ECR_REGISTRY}/${ECR_REPOSITORY}:latest
	@echo "🟢 Push complete." && echo

deploy: push
	@echo "##### DEPLOY #####" && echo
	@echo "Deploying ${VERSION} to ECS service..."
	DEPLOY_RESPONSE=$$(aws ecs update-service --cluster ${ECS_CLUSTER} --service ${ECS_SERVICE} --force-new-deployment --profile infrastructure-admin-dev) && echo "$${DEPLOY_RESPONSE}"
	@echo "🟢 Deployment initiated." && echo