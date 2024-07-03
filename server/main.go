package main

import (
	"discovery_servcie/config"
	pb "discovery_servcie/genproto"
	"discovery_servcie/service"
	"discovery_servcie/storage/postgres"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	cnf := config.Config{}
	db, err := postgres.ConnectionDb(&cnf)
	if err != nil {
		log.Fatalf("error:->%s", err.Error())
	}
	listen, err := net.Listen("tcp", ":8084")
	if err != nil {
		log.Fatalf("error:->%s", err.Error())
	}
	grpcServer := grpc.NewServer()
	pb.RegisterDiscoveryServiceServer(grpcServer, service.NewDiscoveryService(postgres.NewCompositionRepository(db)))
	log.Printf("server listening at %v", listen.Addr())

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
