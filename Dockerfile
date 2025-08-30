FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o calcapi ./main.go

FROM alpine:3.19

WORKDIR /app
COPY --from=builder /app/calcapi .

EXPOSE 8080

CMD ["./calcapi"]
