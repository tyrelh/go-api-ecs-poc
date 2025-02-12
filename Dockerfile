# Pull base image
FROM golang:1.23-alpine AS base

# Development stage
# =======================================================
# Create a development stage based on the base image
FROM base AS development

# Install git
RUN apk update
RUN apk add --no-cache git
RUN apk add --no-cache curl

# Where our files will be in the docker container 
WORKDIR /app

# Copy the source from the current directory to the working Directory inside the container 
# Source also contains go.mod and go.sum which are dependency files
COPY . .

# Install dependencies
RUN go mod download

# RUN CGO_ENABLED=0 go build -o app

EXPOSE 8080

# Install the air CLI for auto-reloading
# RUN go install github.com/airdb/air@latest
RUN curl -sSfL https://goblin.run/github.com/air-verse/air | sh

# Start air for live reloading
CMD ["air"]

# Build stage
# =======================================================
# Create a build stage based on the base image
# FROM base AS builder

# # Move to working directory /build
# WORKDIR /build

# # Copy the go.mod and go.sum files to the /build directory
# COPY go.mod ./

# # Install dependencies
# RUN go mod download

# # Copy the entire source code into the container
# COPY . .

# # Build the application
# RUN CGO_ENABLED=0 go build -o app

# Production stage
# =======================================================
# Create a production stage to run the application binary
FROM scratch AS production

ARG GO_API_VERSION=none
ARG GO_API_AWS_REGION=none

# WORKDIR /prod

# Copy binary from builder stage
# COPY bin ./
ADD bin /

ENV GO_API_VERSION=${GO_API_VERSION}
ENV GO_API_AWS_REGION=${GO_API_AWS_REGION}
# Document the port that may need to be published
EXPOSE 8080

# Start the application
CMD ["/go-api-poc"]

