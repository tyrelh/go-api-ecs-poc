# Go API POC

This is a simple API POC using Go, Docker, and ECS.

## Dependencies

### Go
If you don't already have Go installed, you can install with Brew:
```bash
brew install golang
# test
go version
```
### Docker
I recommend just installing Docker Desktop. It comes bundled with the Docker & Compose CLIs:
```bash
brew install --cask docker
# test
docker -v
```

### Air
[Air](https://github.com/air-verse/air) is a Go app used for hot-reloading Go projects. Can be installed via Goblin:
```bash
curl -sSfL https://goblin.run/github.com/air-verse/air | sh
# add GOPATH to PATH if you don't already have this
echo '\nexport GOPATH="$HOME/go"' >> ~/.zshrc
echo '\nexport PATH=$PATH:$GOPATH/bin' >> ~/.zshrc
# test
air -v
```

### oapi-codegen
[oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) is a Go app used to generate the backend scaffolding code from the OpenAPI specification file. Can be installed globally with Go:
```bash
go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
# add GOPATH to PATH if you don't already have this
echo '\nexport GOPATH="$HOME/go"' >> ~/.zshrc
echo '\nexport PATH=$PATH:$GOPATH/bin' >> ~/.zshrc
# test
oapi-codegen -version
```

### AWS CLI
For pushing images to ECR and triggering deployments to ECS you'll need the AWS CLI installed and appropriate permissions on your credentials.
```bash
brew install awscli
# test
aws --version
```

## Run Locally

```bash
make dev
```

Or run the docker container:
```bash
make build-image
docker run -p 8080:8080 go-api-poc
```

## Build

Build Go binary:
```bash
make build
```

Build Go binary and Docker Image:
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

## Deploy to ECS

```bash
# Will build image, upload to ECR, then deploy to ECS
make deploy
```


# API

Create new Reward
```bash
curl -X POST -H "Content-Type: application/json" -d '{"brand": "Viridian City Pokemart", "currency": "PMD", "denomination": 1000}' http://localhost:8080/go/reward | jq
```

Get all Rewards
```bash
curl http://localhost:8080/go/reward | jq
```

Get specific Reward by ID
```bash
curl http://localhost:8080/go/reward/1 | jq
```

Update Reward
```bash
curl -X PUT -H "Content-Type: application/json" -d '{"brand": "Viridian City Pokemart", "currency": "PMD", "denomination": 1000}' http://localhost:8080/go/reward/1 | jq
```

Delete Reward
```bash
curl -X DELETE http://localhost:8080/go/reward/1 | jq
```

Get Health
```bash
curl http://localhost:8080/go/system/health | jq
```

Get Version
```bash
curl http://localhost:8080/go/system/version | jq
```