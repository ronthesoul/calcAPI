# syntax=docker/dockerfile:1.6

## Build
FROM --platform=$BUILDPLATFORM golang:1.23-alpine AS build
WORKDIR /src
RUN apk add --no-cache ca-certificates git

# deps cache
COPY go.mod go.sum ./
RUN go mod download

# source
COPY . .

# build for target platform
ARG TARGETOS TARGETARCH
ENV CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    go build -trimpath -ldflags="-s -w" -o /out/trafficgen ./traffic-gen

## Runtime
FROM alpine:3.20
RUN adduser -D -H -s /sbin/nologin app && apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=build /out/trafficgen /app/trafficgen
RUN chmod 0555 /app/trafficgen
USER app
ENTRYPOINT ["/app/trafficgen"]
