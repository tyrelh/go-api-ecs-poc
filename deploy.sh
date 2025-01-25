#! /usr/bin/env bash

VERSION=$(cat version)

echo "Building version ${VERSION}..."
docker build -t go-api-poc:${VERSION} . --target production --build-arg VERSION=${VERSION}
echo "🟢 Build complete."

echo "Logging in to ECR..."
aws ecr get-login-password --region us-west-2 --profile infrastructure-admin-dev | docker login --username AWS --password-stdin 784593521445.dkr.ecr.us-west-2.amazonaws.com
echo "🟢 Logged in to ECR."

echo "Tagging image..."
docker tag go-api-poc:${VERSION} 784593521445.dkr.ecr.us-west-2.amazonaws.com/go-api-poc-repository:${VERSION}
docker tag go-api-poc:${VERSION} 784593521445.dkr.ecr.us-west-2.amazonaws.com/go-api-poc-repository:latest
echo "🟢 Tagging complete."

echo "Pushing image to ECR..."
docker push 784593521445.dkr.ecr.us-west-2.amazonaws.com/go-api-poc-repository:${VERSION}
docker push 784593521445.dkr.ecr.us-west-2.amazonaws.com/go-api-poc-repository:latest
echo "🟢 Push complete."

echo ""; echo "✅ Done."