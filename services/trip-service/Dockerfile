FROM golang:alpine AS builder
WORKDIR /app

COPY go.sum go.mod ./
RUN go mod download

COPY shared ./shared
COPY services/trip-service ./services/trip-service

# Tell Docker to change directory to where main.go is!
WORKDIR /app/services/trip-service/cmd
RUN go build -o trip-service .

FROM alpine:latest
WORKDIR /app
# Copy from the cmd folder where it was built
COPY --from=builder /app/services/trip-service/cmd/trip-service .

EXPOSE 8083
CMD ["./trip-service"]
