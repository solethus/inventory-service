package main

import (
	"net"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"

	"github.com/solethus/inventory-service/internal/server"
	log "github.com/solethus/inventory-service/pkg/logger"
	pb "github.com/solethus/inventory-service/proto/inventory"
)

func main() {
	log.InitLogging()

	// Set up database connection
	db, err := sqlx.Open("postgres", "postgres://username:password@localhost/inventory?sslmode=disable")
	if err != nil {
		log.Logger.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Create gRPC server
	s := grpc.NewServer()
	inventoryServer := server.NewInventoryServer(db)
	pb.RegisterInventoryServiceServer(s, inventoryServer)

	// Start listening
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Logger.Fatalf("Failed to listen: %v", err)
	}

	log.Logger.Info("Server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Logger.Fatalf("Failed to serve: %v", err)
	}
}
