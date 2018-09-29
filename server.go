package main

import (
	"github.com/uqichi/payjp/proto"
	"github.com/uqichi/payjp/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const port string = ":50051"

func main() {
	m := service.NewPayManager()

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	proto.RegisterPayManagerServer(server, m)
	reflection.Register(server)

	// run server
	log.Printf("gRPC server started on %s", port)
	err = server.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}
