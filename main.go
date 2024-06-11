package main

import (
	"log"
	"net"

	pbCalculator "github.com/andrewlawrence80/grpc-demo-proto/calculator"
	pbGreeter "github.com/andrewlawrence80/grpc-demo-proto/greeter"
	calculator "github.com/andrewlawrence80/grpc-demo-server/calculator"
	greeter "github.com/andrewlawrence80/grpc-demo-server/greeter"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greeterServer := &greeter.GreeterServerImpl{}
	calculatorServer := &calculator.CalculatorServerImpl{}
	pbGreeter.RegisterGreeterServer(s, greeterServer)
	pbCalculator.RegisterCalculatorServer(s, calculatorServer)

	reflection.Register(s)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
