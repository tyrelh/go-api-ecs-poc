# Use Go 1.23 bookworm as base image
# FROM golang:latest AS base

# Development stage
# =======================================================
# Create a development stage based on the base image
# FROM base AS development

# # Change the working directory to /app
# WORKDIR /app

# # Install the air CLI for auto-reloading
# RUN go install github.com/airdb/air@latest
# # RUN curl -sSfL https://goblin.run/github.com/air-verse/air | sh

# # Copy the go.mod and go.sum files to the /app directory
# COPY go.mod ./

# # Install dependencies
# RUN go mod download

# # Start air for live reloading
# CMD ["air"]

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

ARG VERSION=none

# WORKDIR /prod

# Copy binary from builder stage
# COPY bin ./
ADD bin /

ENV VERSION=${VERSION}
# Document the port that may need to be published
EXPOSE 8080

# Start the application
CMD ["/go-api-poc"]

