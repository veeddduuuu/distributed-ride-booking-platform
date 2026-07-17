package main

import (
	"context"
	// "encoding/json"
	// "fmt"
	"log"
	"net"

	// "net/http"
	"os"
	"os/signal"
	"syscall"

	// "time"

	"github.com/veeddduuuu/distributed-ride-booking-platform/shared/types"
	pb "github.com/veeddduuuu/distributed-ride-booking-platform/shared/proto/trip"
	grpcserver "google.golang.org/grpc"
)

type PreviewRequest struct {
	Pickup      types.Coordinates `json:"pickup"`
	Destination types.Coordinates `json:"destination"`
}

var grpcAddr = ":8083"

func main() {
	inmemrepo := NewInmemRepository()
	svc := NewService(inmemrepo)

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

	grpcServer := grpcserver.NewServer()
	handler := NewTripServiceHandler(svc)       // 1. create the handler
	pb.RegisterTripServiceServer(grpcServer, handler) // 2. register it so the server can route calls to it
	
	log.Printf("Starting gRPC server Trip service on port %s", lis.Addr().String())

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Printf("failed to serve: %v", err)
			cancel()
		}
	}()

	<-ctx.Done()
	log.Printf("shutting down the server....")
	grpcServer.GracefulStop()

}
