package main

import (
	"log"
	"net"
	"net/http"

	"github.com/wari/grpc-trace-test/protos"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port  = ":50051"
	trace = ":8888"
)

type server struct{}

func (s *server) ThisWorks(ctx context.Context, _ *grpctest.EmptyRequest) (*grpctest.EmptyResponse, error) {
	return &grpctest.EmptyResponse{}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	grpctest.RegisterWorkingServer(s, &server{})
	go s.Serve(lis)
	log.Fatal(http.ListenAndServe(trace, nil)) // For net/trace
}
