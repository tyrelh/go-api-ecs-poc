VERSION=$(shell cat version)

dev:
	@echo "Local dev using Air"
	@air

build:
	@echo "### Build ###" && echo
	@echo "Building version ${VERSION}..."
	docker build -t go-api-poc:${VERSION} . --target production --build-arg VERSION=${VERSION}
	@echo "🟢 Build complete." && echo

push: build
	@echo && echo "### Push ###" && echo
	@echo "Logging in to ECR..."
	aws ecr get-login-password --region us-west-2 --profile infrastructure-admin-dev | docker login --username AWS --password-stdin 784593521445.dkr.ecr.us-west-2.amazonaws.com
	@echo "🟢 Logged in to ECR." && echo
	@echo "Tagging image..."
	docker tag go-api-poc:${VERSION} 784593521445.dkr.ecr.us-west-2.amazonaws.com/go-api-poc-repository:${VERSION}
	docker tag go-api-poc:${VERSION} 784593521445.dkr.ecr.us-west-2.amazonaws.com/go-api-poc-repository:latest
	@echo "🟢 Tagging complete." && echo
	@echo "Pushing image to ECR..."
	docker push 784593521445.dkr.ecr.us-west-2.amazonaws.com/go-api-poc-repository:${VERSION}
	docker push 784593521445.dkr.ecr.us-west-2.amazonaws.com/go-api-poc-repository:latest
	@echo "🟢 Push complete." && echo

deploy: push
	@echo "### Deploy ###" && echo
	@echo "Deploying ${VERSION} to ECS service..."
	DEPLOY_RESPONSE=$$(aws ecs update-service --cluster go-api-poc-cluster --service go-api-poc-service --force-new-deployment --profile infrastructure-admin-dev) && echo "$${DEPLOY_RESPONSE}"
	@echo "🟢 Deployment initiated." && echo