# Go API POC

This is a simple API POC using Go, Docker, and ECS.

## Build

```bash
docker build -t go-api-poc:0.1.0 .
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

## Push to Amazon ECR

1. Authenticate Docker to your Amazon ECR registry:
```bash
aws ecr get-login-password --region <your-region> | docker login --username AWS --password-stdin <your-account-id>.dkr.ecr.<your-region>.amazonaws.com
```

2. Tag your image to match your ECR repository:
```bash
docker tag go-api-poc:latest <your-account-id>.dkr.ecr.<your-region>.amazonaws.com/go-api-poc:latest
```

3. Push the image to ECR:
```bash
docker push <your-account-id>.dkr.ecr.<your-region>.amazonaws.com/go-api-poc:latest
```

Note: Replace the following placeholders:
- `<your-account-id>`: Your AWS account ID
- `<your-region>`: Your AWS region (e.g., us-east-1)

Make sure you have:
1. AWS CLI installed and configured
2. Appropriate IAM permissions to push to ECR
3. Created an ECR repository named 'go-api-poc'

