FROM golang:1.22-alpine as builder
WORKDIR /app
COPY go.sum go.mod ./
RUN go mod download

COPY services/trip-service ./services/trip-service

RUN go build -o trip-service .

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/services/trip-service/trip-service .

EXPOSE 8080
CMD ["./trip-service"]
