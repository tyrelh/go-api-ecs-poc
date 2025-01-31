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

### OpenAPI Code Generation
Install [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen):
```bash
go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
```
For a go binary like this to work on your machine as expected, you need to add your GOPATH to your PATH. On MacOS the default GOPATH is *~/go/bin*.

Add the following somewhere in your *~/.zshrc* file:
```
# go path
export GOPATH="$HOME/go"
export PATH=$PATH:$GOPATH/bin
```

Just run OpenAPI codegen for the *api/api.yml* API specification:
```bash
make api
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
make build-image
docker run -p 8080:8080 go-api-poc
```

## Deploy to ECS

```bash
# Will build image, upload to ECR, then deploy to ECS
make deploy
```
