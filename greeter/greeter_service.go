package greeter

import (
	"context"
	"fmt"

	pb "github.com/andrewlawrence80/grpc-demo-proto/greeter"
)

type GreeterServerImpl struct {
	pb.UnimplementedGreeterServer
}

func (s *GreeterServerImpl) SayHello(cxt context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Name),
	}, nil
}
