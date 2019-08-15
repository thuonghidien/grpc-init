package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/thuonghidien/grpc-init/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type server struct{}

var (
	addr = flag.String("addr", ":50051", "Network host:port to listen on for gRPC connections.")
)

func main() {

	flag.Parse()
	fmt.Println("Server init")
	listener, err := net.Listen("tcp", *addr)
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	pb.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	if err = srv.Serve(listener); err != nil {
		panic(err)
	}
}

func (s *server) Add(ctx context.Context, request *pb.Request) (*pb.Response, error) {

	a, b := request.GetA(), request.GetB()
	result := a + b
	return &pb.Response{Result: result}, nil
}

func (s *server) Subtract(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	a, b := request.GetA(), request.GetB()
	result := a - b
	return &pb.Response{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, request *pb.Request) (*pb.Response, error) {

	a, b := request.GetA(), request.GetB()
	result := a * b
	return &pb.Response{Result: result}, nil
}

func (s *server) Divide(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	a, b := request.GetA(), request.GetB()
	if b == 0 {

		return &pb.Response{Result: -1}, nil
	}
	result := a / b
	return &pb.Response{Result: result}, nil
}
