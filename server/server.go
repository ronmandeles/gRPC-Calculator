package main

import (
	"context"
	"log"
	"net"

	pb "github.com/ronmandeles/gRPC-Calculator/pkg/pb/v1"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCalculatorServer
}

func (s *server) Add(ctx context.Context, req *pb.OpRequest) (*pb.OpResponse, error) {
	res := req.Num1 + req.Num2
	return &pb.OpResponse{Result: res}, nil
}

func (s *server) Div(ctx context.Context, req *pb.OpRequest) (*pb.OpResponse, error) {
	res := req.Num1 / req.Num2
	return &pb.OpResponse{Result: res}, nil
}

func (s *server) Mul(ctx context.Context, req *pb.OpRequest) (*pb.OpResponse, error) {
	res := req.Num1 * req.Num2
	return &pb.OpResponse{Result: res}, nil
}

func (s *server) Sub(ctx context.Context, req *pb.OpRequest) (*pb.OpResponse, error) {
	res := req.Num1 - req.Num2
	return &pb.OpResponse{Result: res}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Printf("Failed to listen, %v", err)
	}

	grpc_server := grpc.NewServer()
	pb.RegisterCalculatorServer(grpc_server, &server{})
	if err := grpc_server.Serve(lis); err != nil {
		log.Printf("failed to server, %v", err)
	}

}
