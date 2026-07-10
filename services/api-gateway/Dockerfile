FROM golang:alpine AS builder
WORKDIR /app

COPY go.sum go.mod ./
RUN go mod download

COPY services/api-gateway ./services/api-gateway

# Tell Docker to change directory before building!
WORKDIR /app/services/api-gateway
RUN go build -o api-gateway .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/services/api-gateway/api-gateway .

EXPOSE 8080
CMD ["./api-gateway"]
