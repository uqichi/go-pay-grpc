package main

import (
	"log"
	"net"

	"github.com/uqichi/go-pay-grpc/proto"
	"github.com/uqichi/go-pay-grpc/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port string = ":50051"

func main() {
	m := service.NewPayService()

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	pb.RegisterPayServiceServer(server, m)
	reflection.Register(server)

	// run server
	log.Printf("gRPC server started on %s", port)
	err = server.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
