package v2

import (
	"context"
	"fmt"

	"github.com/imjoshholloway/grpc-greeter/api/v2"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Greet(ctx context.Context, req *greeter.GreetRequest) (*greeter.GreetResponse, error) {

	name := "friend"
	if req.GetName() != "" {
		name = req.GetName()
	}

	return &greeter.GreetResponse{
		Message: fmt.Sprintf("Nice to meet you, %v", name),
	}, nil
}

func Register(grpcServer *grpc.Server) {
	greeter.RegisterGreeterServer(grpcServer, &server{})
}
