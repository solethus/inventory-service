package repository

import (
	"context"

	"github.com/jmoiron/sqlx"

	pb "github.com/solethus/shared-proto/proto/inventory"
)

type InventoryRepository interface {
	AddCar(ctx context.Context, car *pb.Car) (string, error)
	GetCar(ctx context.Context, id string) (*pb.Car, error)
	ListCars(ctx context.Context, pageSize, pageNumber int32) ([]*pb.Car, int32, error)
	UpdateCarStock(ctx context.Context, id string, stockChange int32) (int32, error)
}

type inventoryRepository struct {
	db *sqlx.DB
}

func NewInventoryRepository(db *sqlx.DB) InventoryRepository {
	return &inventoryRepository{db: db}
}

func (i inventoryRepository) AddCar(ctx context.Context, car *pb.Car) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (i inventoryRepository) GetCar(ctx context.Context, id string) (*pb.Car, error) {
	//TODO implement me
	panic("implement me")
}

func (i inventoryRepository) ListCars(ctx context.Context, pageSize, pageNumber int32) ([]*pb.Car, int32, error) {
	//TODO implement me
	panic("implement me")
}

func (i inventoryRepository) UpdateCarStock(ctx context.Context, id string, stockChange int32) (int32, error) {
	//TODO implement me
	panic("implement me")
}
