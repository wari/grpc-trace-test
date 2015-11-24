package main

import (
	"log"

	"github.com/wari/grpc-trace-test/protos"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	m := grpctest.NewNotWorkingClient(conn)
	x, err := m.Unimplemented(context.Background(), &grpctest.EmptyRequest{})
	log.Println("Expected error", x, err) // But shows up on net/trace as Active

	n := grpctest.NewWorkingClient(conn)
	y, err := n.ThisWorks(context.Background(), &grpctest.EmptyRequest{})
	log.Println("No errors here", y, err)
}
