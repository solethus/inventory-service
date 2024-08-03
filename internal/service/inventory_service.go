package service

import (
	"context"

	"github.com/google/uuid"
	pb "github.com/solethus/shared-proto/proto/inventory"

	"github.com/solethus/inventory-service/internal/model"
	"github.com/solethus/inventory-service/internal/repository"
)

type InventoryService interface {
	AddCar(ctx context.Context, car *pb.Car) (*string, error)
	GetCar(ctx context.Context, id uuid.UUID) (*model.Car, error)
	UpdateCarStock(ctx context.Context, id uuid.UUID, change int) error
	SearchCars(ctx context.Context, filter model.CarFilter, pagination model.PaginationParams, sort model.SortParams) ([]*model.Car, int, error)
	GenerateInventoryReport(ctx context.Context) (*model.InventoryStats, error)
}

type inventoryService struct {
	inventoryRepo repository.InventoryRepository
}

func NewInventoryService(inventoryRepo repository.InventoryRepository) InventoryService {
	return &inventoryService{
		inventoryRepo: inventoryRepo,
	}
}

func (i inventoryService) AddCar(ctx context.Context, car *pb.Car) (*string, error) {
	panic("implement me")
}

func (i inventoryService) GetCar(ctx context.Context, id uuid.UUID) (*model.Car, error) {
	panic("implement me")
}

func (i inventoryService) UpdateCarStock(ctx context.Context, id uuid.UUID, change int) error {
	panic("implement me")
}

func (i inventoryService) SearchCars(ctx context.Context, filter model.CarFilter, pagination model.PaginationParams, sort model.SortParams) ([]*model.Car, int, error) {
	panic("implement me")
}

func (i inventoryService) GenerateInventoryReport(ctx context.Context) (*model.InventoryStats, error) {
	panic("implement me")
}
