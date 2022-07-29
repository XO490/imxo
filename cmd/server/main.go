package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"imxo/proto"
	"log"
	"net"
)

var (
	port = flag.Int("port", 4994, "the server port")
)

func main() {
	flag.Parse()
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Fatal to listen: %v", err)
	}

	s := grpc.NewServer()
	//reflection.Register(s)
	proto.RegisterSenderServer(s, &proto.ChatServer{})

	log.Printf("IMXO server start at %v", l.Addr())

	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
