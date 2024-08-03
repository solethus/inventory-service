package server

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/jmoiron/sqlx"
	pb "github.com/solethus/shared-proto/proto/inventory"

	"github.com/solethus/inventory-service/internal/repository"
	log "github.com/solethus/inventory-service/pkg/logger"
)

type InventoryServer struct {
	pb.UnimplementedInventoryServiceServer
	repo repository.InventoryRepository
}

func NewInventoryServer(db *sqlx.DB) *InventoryServer {
	return &InventoryServer{
		repo: repository.NewInventoryRepository(db),
	}
}

func (s *InventoryServer) AddCar(ctx context.Context, req *pb.AddCarRequest) (*pb.AddCarResponse, error) {
	id, err := s.repo.AddCar(ctx, req.Car)
	if err != nil {
		log.Logger.Error("error failed to add car", err)
		return nil, status.Errorf(codes.Internal, "failed to add car: %v", err)
	}
	return &pb.AddCarResponse{Id: id}, nil
}

func (s *InventoryServer) GetCar(ctx context.Context, req *pb.GetCarRequest) (*pb.Car, error) {
	// Implementation
	panic("implement me")
}

func (s *InventoryServer) ListCars(ctx context.Context, req *pb.ListCarsRequest) (*pb.ListCarsResponse, error) {
	// Implementation
	panic("implement me")
}

func (s *InventoryServer) UpdateCarStock(ctx context.Context, req *pb.UpdateCarStockRequest) (*pb.UpdateCarStockResponse, error) {
	// Implementation
	panic("implement me")
}
