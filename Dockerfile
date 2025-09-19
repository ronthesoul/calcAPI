# syntax=docker/dockerfile:1.6

FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/calcapi ./main.go

FROM alpine:3.19
WORKDIR /app
RUN apk add --no-cache ca-certificates curl
COPY --from=builder /out/calcapi /app/calcapi
EXPOSE 8080
ENTRYPOINT ["/app/calcapi"]

