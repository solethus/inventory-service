package main

import (
	"github.com/solethus/inventory-service/internal/repository"
	"github.com/solethus/inventory-service/internal/service"
	"net"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"

	"github.com/solethus/inventory-service/internal/server"
	log "github.com/solethus/inventory-service/pkg/logger"
	pb "github.com/solethus/shared-proto/proto/inventory"
)

func main() {
	log.InitLogging()

	// Set up database connection
	db, err := sqlx.Open("postgres", "postgres://username:password@localhost/inventory?sslmode=disable")
	if err != nil {
		log.Logger.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Create repository, service, and server
	repo := repository.NewInventoryRepository(db)
	svc := service.NewInventoryService(repo)
	srv := server.NewInventoryServer(svc)

	// Create gRPC server
	s := grpc.NewServer()
	pb.RegisterInventoryServiceServer(s, srv)

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
