package main

import (
	"flag"
	"fmt"
	"github.com/heyimblake/charbalancer-service/internal/service"
	"github.com/heyimblake/charbalancer-service/pkg/api/proto"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

var port int
const PREFIX string = "[CharBalancer-Server]"

func init() {
	flag.IntVar(&port, "port", 7777, "the port to run the charbalancer service on")
}

func main() {
	// Get port from flag
	flag.Parse()

	// Create a TCP listener on a given port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
		return
	}

	fmt.Printf("%s Listening on port %d.\n", PREFIX, port)

	// Create service instance
	balService := service.BalancerService{}

	// Create grpc server instance
	grpcServer := grpc.NewServer()

	// Register service with grpc server
	proto.RegisterCharBalancerServer(grpcServer, &balService)

	// Handler for os signals
	go func(grpcServer *grpc.Server) {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		<-c
		grpcServer.GracefulStop()
		fmt.Printf("%s Stopping server...\n", PREFIX)
	}(grpcServer)

	// Start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s\n", err)
	}

	fmt.Printf("%s Server stopped!\n", PREFIX)
}