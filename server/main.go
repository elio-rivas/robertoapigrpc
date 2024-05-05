package server

import (
	pb "github.com/elio-rivas/robertoapigrpc/proto"
	"github.com/elio-rivas/robertoapigrpc/service"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	serverPort = ":8081" // Change port as needed
)

func main() {
	//Initialize the gRPC server
	listen, err := net.Listen("tcp", serverPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()

	catalogService := service.NewCatalogService()
	pb.RegisterCatalogServiceServer(s, catalogService)

	log.Printf("Starting server on port %s", serverPort)
	if err := s.Serve(listen); err != nil {

		log.Fatalf("Failed to serve: %v", err)
	}
}
