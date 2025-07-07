# syntax=docker/dockerfile:1

# Build stage
FROM --platform=${BUILDPLATFORM} golang:1.24.4-bookworm AS builder

WORKDIR /app

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=bind,source=go.mod,target=go.mod \
    --mount=type=bind,source=go.sum,target=go.sum \
    go mod download

ENV CGO_ENABLED=0
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=bind,source=.,target=. \
    go build -ldflags="-w -s" -o /usr/bin/rsss .

# Deployment stage
FROM gcr.io/distroless/base-debian11 AS deploy

COPY --from=builder /usr/bin/rsss /usr/bin/rsss

USER 65532:65532

CMD ["/usr/bin/rsss"]
