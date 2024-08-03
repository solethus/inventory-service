package server

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/solethus/inventory-service/internal/repository"
	pb "github.com/solethus/inventory-service/proto/inventory"
)

//type InventoryServer interface {
//	AddCar(ctx context.Context, req *pb.AddCarRequest) (*pb.AddCarResponse, error)
//	GetCar(ctx context.Context, req *pb.GetCarRequest) (*pb.Car, error)
//	ListCars(ctx context.Context, req *pb.ListCarsRequest) (*pb.ListCarsResponse, error)
//	UpdateCarStock(ctx context.Context, req *pb.UpdateCarStockRequest) (*pb.UpdateCarStockResponse, error)
//}

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
	// Implementation
	ud

	panic("implement me")
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
