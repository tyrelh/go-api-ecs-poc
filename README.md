# Go API POC

This is a simple API POC using Go, Docker, and ECS.

## Build

Development build:
```bash
docker build -t go-api-poc:0.1.0 . --target development
```

Deployable build:
```bash
docker build -t go-api-poc:0.1.0 . --target production
```

## Run Locally

Run just Go:
```bash
go run .
```

Or run the docker container:
```bash
docker run -p 8080:8080 go-api-poc
```

Or using Air locally for hot-reloading:
```bash
# install Air
curl -sSfL https://goblin.run/github.com/air-verse/air | sh
# run Air
air
```

## API

Create a new item:
```bash
curl -X POST -H "Content-Type: application/json" -d '{"id": "22", "name": "🔮"}' http://localhost:8080/items
```

Fetch an item by id:
```bash
curl http://localhost:8080/items/22
```

Fetch all items:
```bash
curl http://localhost:8080/items
```

## Push to Amazon ECR (Elastic Container Registry)

1. Authenticate Docker to your Amazon ECR registry:
```bash
aws ecr get-login-password --region us-west-2 --profile infrastructure-admin-dev | docker login --username AWS --password-stdin 784593521445.dkr.ecr.us-west-2.amazonaws.com
```

2. Tag your image to match your ECR repository:
```bash
VERSION="0.2.1"
docker tag go-api-poc:${VERSION} 784593521445.dkr.ecr.us-west-2.amazonaws.com/go-api-poc-repository:${VERSION}
docker tag go-api-poc:${VERSION}784593521445.dkr.ecr.us-west-2.amazonaws.com/go-api-poc-repository:latest
```

3. Push the image to ECR:
```bash
docker push 784593521445.dkr.ecr.us-west-2.amazonaws.com/go-api-poc-repository:${VERSION}
docker push 784593521445.dkr.ecr.us-west-2.amazonaws.com/go-api-poc-repository:latest
```

You can also use the Docker cli to pull images from ECR:
```bash
docker image pull 784593521445.dkr.ecr.us-west-2.amazonaws.com/go-api-poc-repository:0.1.1
```

