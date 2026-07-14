package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/veeddduuuu/distributed-ride-booking-platform/services/trip-service/internal/infrastructure/repository"
	"github.com/veeddduuuu/distributed-ride-booking-platform/services/trip-service/internal/service"
	"github.com/veeddduuuu/distributed-ride-booking-platform/shared/types"
	grpcserver "google.golang.org/grpc"
)

type PreviewRequest struct {
	Pickup      types.Coordinates `json:"pickup"`
	Destination types.Coordinates `json:"destination"`
}

var grpcAddr = ":9093"

func main() {
	// inmemrepo := repository.NewInmemRepository()
	// svc := service.NewService(inmemrepo)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM) 
		<- sigCh
		cancel()
	}()

	lis, err := net.Listen("tcp", grpcAddr)
	if err!=nil{
		log.Fatalf("failed to listen: %v", err)
	}

	grpcserver := grpcserver.NewServer()
	
	log.Printf("Starting gRPC server Trip service on port %s", lis.Addr().String())

	go func() {
		if err:= grpcserver.Serve(lis); err!=nil{
			log.Printf("failed to serve: %v", err)
			cancel()
		}
	}()

	<-ctx.Done()
	log.Printf("shutting down the server....")
	grpcserver.GracefulStop()

}
