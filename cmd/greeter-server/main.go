package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/imjoshholloway/grpc-greeter/internal/v1"
	"github.com/imjoshholloway/grpc-greeter/internal/v2"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	name     = "greeter-server"
	version  = "dev"
	datetime = "unknown"

	// flags
	port = flag.Int("port", 8080, "The port to start the server on")
)

func main() {
	flag.Parse()

	log.Printf("Starting %s@%s listening on %v", name, version, *port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	// register both versions of the service
	v1.Register(server)
	v2.Register(server)

	reflection.Register(server)

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
