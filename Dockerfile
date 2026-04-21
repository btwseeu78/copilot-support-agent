# Stage 1: Build the Go binary
FROM golang:1.26-alpine AS builder

RUN apk add --no-cache git
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN rm -f zcopilot_darwin_*.go zcopilot_*_darwin_*.zst zcopilot_*_darwin_*.license
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /bin/copilot-support-agent ./main.go

# Stage 2: Runtime with Copilot CLI
FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y --no-install-recommends curl ca-certificates && \
    curl -fsSL https://gh.io/copilot-install | bash && \
    apt-get purge -y curl && apt-get autoremove -y && rm -rf /var/lib/apt/lists/*

COPY --from=builder /bin/copilot-support-agent /usr/local/bin/copilot-support-agent
RUN chmod +x /usr/local/bin/copilot-support-agent

WORKDIR /workspace
CMD ["sh"]