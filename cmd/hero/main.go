package main

import (
	"log"
	"net"

	svc "hk01/go-rpc-service/internal"
)

func main() {
	var (
		grpcAddr = ":6002"
	)

	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := svc.NewServer()
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
