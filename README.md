# Go API POC

This is a simple API POC using Go, Docker, and ECS.

## Build

Build Go:
```bash
make build
```

Build Go and Docker Image:
```bash
make build-image
```

Build Docker Image and push to ECR:
```bash
# Will build image, then upload to ECR
make push
```

## Run Locally

Run just Go:
```bash
# install Air
curl -sSfL https://goblin.run/github.com/air-verse/air | sh
# run locally using Air
make dev
```

Or run the docker container:
```bash
docker run -p 8080:8080 go-api-poc
```

## Deploy to ECS

```bash
# Will build image, upload to ECR, then deploy to ECS
make deploy
```


## API

Create a new item:
```bash
curl -X POST -H "Content-Type: application/json" -d '{"id": "22", "name": "🔮"}' http://localhost:8080/item
```

Fetch an item by id:
```bash
curl http://localhost:8080/item/22
```

Fetch all items:
```bash
curl http://localhost:8080/item
```