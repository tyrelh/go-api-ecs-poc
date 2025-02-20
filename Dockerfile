FROM golang:1.24-alpine AS base

# Development stage
# =======================================================
# Sets up a development environment with hot-reloading
# and debugging tools.
FROM base AS development
RUN apk update
RUN apk add --no-cache git
RUN apk add --no-cache curl
WORKDIR /app
COPY . .
RUN go mod download
EXPOSE 8080
RUN curl -sSfL https://goblin.run/github.com/air-verse/air | sh
CMD ["air"]

# Build stage
# =======================================================
# Builds the Go binary in a container so that it's
# build environment is consistent and isolated. 
FROM base AS builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o app

# Production stage
# =======================================================
# Copies the Go binary from the build stage into a
# scratch container to keep the image size small.
FROM scratch AS production
WORKDIR /prod
ARG GO_API_VERSION=none
ARG GO_API_AWS_REGION=none
COPY --from=builder /build/app ./
ENV GO_API_VERSION=${GO_API_VERSION}
ENV GO_API_AWS_REGION=${GO_API_AWS_REGION}
EXPOSE 8080
CMD ["/prod/app"]
