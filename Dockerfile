# -------------------------------------------------------------
# Stage 1: Build the Go Application
# -------------------------------------------------------------
FROM golang:1.26-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the analyzer tool
# CGO_ENABLED=0 ensures a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /bin/copilot-support-agent ./main.go

# -------------------------------------------------------------
# Stage 2: Final Runtime Image
# -------------------------------------------------------------
# Switching to debian-slim as the runtime. Pre-compiled binaries downloaded 
# via shell scripts (like the official copilot CLI) often dynamically link 
# against glibc, which fails on Alpine's musl libc.
FROM debian:bookworm-slim

# Install essential tools and runtimes:
# - git: needed by helm fetching if using git repos
# - bash/curl: standard utilities for CI pipelines and installation scripts
# - ca-certificates: Required for secure curl downloads
RUN apt-get update && apt-get install -y --no-install-recommends \
    git \
    bash \
    curl \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Install the official GitHub Copilot CLI
RUN curl -fsSL https://gh.io/copilot-install | bash

# Copy the built custom tool from the builder stage
# Pipeline expects this at /usr/local/bin/copilot-support-agent
COPY --from=builder /bin/copilot-support-agent /usr/local/bin/copilot-support-agent

# Ensure the binary is executable
RUN chmod +x /usr/local/bin/copilot-support-agent

# Set up runtime environment for the CI
WORKDIR /workspace
CMD ["bash"]
