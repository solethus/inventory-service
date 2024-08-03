package server

import (
	"context"
	"github.com/solethus/inventory-service/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/solethus/shared-proto/proto/inventory"

	log "github.com/solethus/inventory-service/pkg/logger"
)

type InventoryServer struct {
	pb.UnimplementedInventoryServiceServer
	service service.InventoryService
}

func NewInventoryServer(service service.InventoryService) *InventoryServer {
	return &InventoryServer{
		service: service,
	}
}

func (s *InventoryServer) AddCar(ctx context.Context, req *pb.AddCarRequest) (*pb.AddCarResponse, error) {
	id, err := s.service.AddCar(ctx, req.Car)
	if err != nil {
		log.Logger.Error("error failed to add car", err)
		return nil, status.Errorf(codes.Internal, "failed to add car: %v", err)
	}
	if id == nil || *id == "" {
		log.Logger.Infof("invalid car id: %v\n", id)
		return nil, status.Error(codes.Internal, "invalid car id")
	}
	return &pb.AddCarResponse{Id: *id}, nil
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
