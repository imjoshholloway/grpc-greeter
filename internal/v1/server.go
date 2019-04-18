package v1

import (
	"github.com/imjoshholloway/grpc-greeter/api/v1"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Greet(ctx context.Context, req *greeter.GreetRequest) (*greeter.GreetResponse, error) {
	return &greeter.GreetResponse{
		Message: fmt.Sprintf("Thank you for the greeting \"%v\"", req.GetGreeting()),
	}, nil
}

func Register(grpcServer *grpc.Server) {
	greeter.RegisterGreeterServer(grpcServer, &server{})
}