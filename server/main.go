package main

import (
	"log"
	"net"

	pb "../pb"
	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:8090")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	lendMemory := &lendMemory{}

	pb.RegisterLendServiceServer(server, lendMemory)
	server.Serve(listen)
}
