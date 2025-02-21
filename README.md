# Go API POC

This is a simple API POC using Go, Docker, and ECS.

## Dependencies

You should be able to run `make deps` in the root of the project and it will install all necessary dependencies for you, including:
- Brew
- Go
- Docker
- Air
- Delve
- oapi-codegen
- sqlc

There are two things you may need to manually do.

### Go PATH
After running `make deps`, try runnig `which air`. If you **DON'T** see a path to the `air` binary, then you'll need to do the following:
```bash
# add GOPATH to PATH if you don't already have this
echo '\n# go path' >> ~/.zshrc
echo 'export GOPATH="$HOME/go"' >> ~/.zshrc
echo 'export GOBIN=$GOPATH/bin' >> ~/.zshrc
echo 'export PATH=$PATH:$GOBIN' >> ~/.zshrc
```

### AWS CLI
For pushing images to ECR and triggering deployments to ECS you'll need the AWS CLI installed and appropriate permissions on your credentials.
```bash
brew install awscli
# test
aws --version
```
For Giftbit devs, use your development stick credentials you already have. You'll need access to _InfrastructureAdminRole_ in the dev account and the _infrastructure-admin-dev_ profile configured on your laptop. Reach out to Tyrel if you need assistance setting that up.

## Run Locally

```bash
make dev
```

Or experimentially run the app in a docker container locally:
```bash
make build
docker run -p 8080:8080 go-api-poc
```

## Build

Build Go binary:
```bash
make build-binary
```

Build Go binary and Docker Image:
```bash
make build
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
make oapi
```

## Deploy to ECS

```bash
# Will build image, upload to ECR, then deploy to ECS
make deploy
```

## Deploy whole stack of infrastructure to an AWS Region

Navigate to the _junk-drawer_ directory and run:
```bash
./deploy-all.sh --aws-account dev --aws-region ca-west-1
```

You can also terminate all resources in a given region
```bash
./terminate-all.sh --aws-account dev --aws-region ca-west-1
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