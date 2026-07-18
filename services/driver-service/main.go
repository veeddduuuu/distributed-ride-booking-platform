package driverservice

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/veeddduuuu/distributed-ride-booking-platform/shared/proto/driver"
	"google.golang.org/grpc"
)

func main()  {
	inmem := NewInmemRepository()
	svc:= NewService(inmem)

	ctx, cancel := context.WithCancel(context.Background())

	go func(){
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM) 
		<-sigCh
		cancel()
	}()
	
	lis, err:= net.Listen("tcp", ":8084")
	if err!=nil{
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer:= grpc.NewServer()
	handler:=NewDriverServiceHandler(svc)
	pb.RegisterDriverServiceServer(grpcServer, handler)

	log.Printf("Starting gRPC server Driver service on port %s", lis.Addr().String())

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Printf("failed to serve: %v", err)
		}
		cancel()
	}()
		
	<-ctx.Done()
	log.Printf("shutting down the server....")
	grpcServer.GracefulStop()
}


