package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/thuonghidien/grpc-init/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type server struct{}

var (
	addr = flag.String("addr", ":50051", "Network host:port to listen on for gRPC connections.")
)

func main() {

	fmt.Println("Server init")
	listener, err := net.Listen("tcp", *addr)
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	if err = srv.Serve(listener); err != nil {
		panic(err)
	}
}

func (s *server) Add(ctx context.Context, request *proto.Request) (*proto.Response, error) {

	a, b := request.GetA(), request.GetB()
	result := a + b
	return &proto.Response{Result: result}, nil
}

func (s *server) Subtract(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()
	result := a - b
	return &proto.Response{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, request *proto.Request) (*proto.Response, error) {

	a, b := request.GetA(), request.GetB()
	result := a * b
	return &proto.Response{Result: result}, nil
}

func (s *server) Divide(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()
	if b == 0 {

		return &proto.Response{Result: -1}, nil
	}
	result := a / b
	return &proto.Response{Result: result}, nil
}
