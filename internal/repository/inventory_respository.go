package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/solethus/inventory-service/internal/model"

	"github.com/jmoiron/sqlx"
)

type InventoryRepository interface {
	AddCar(ctx context.Context, car *model.Car) (*string, error)
	GetCar(ctx context.Context, id uuid.UUID) (*model.Car, error)
	UpdateCar(ctx context.Context, car *model.Car) error
	DeleteCar(ctx context.Context, id uuid.UUID) error
	ListCars(ctx context.Context, filter model.CarFilter, pagination model.PaginationParams, sort model.SortParams) ([]*model.Car, int, error)
	GetInventoryStats(ctx context.Context) (*model.InventoryStats, error)
}

type inventoryRepository struct {
	db *sqlx.DB
}

func NewInventoryRepository(db *sqlx.DB) InventoryRepository {
	return &inventoryRepository{db: db}
}

func (r *inventoryRepository) AddCar(ctx context.Context, car *model.Car) (*string, error) {
	//TODO implement me
	panic("implement me")
}

func (r *inventoryRepository) GetCar(ctx context.Context, id uuid.UUID) (*model.Car, error) {
	//TODO implement me
	panic("implement me")
}

func (r *inventoryRepository) UpdateCar(ctx context.Context, car *model.Car) error {
	//TODO implement me
	panic("implement me")
}

func (r *inventoryRepository) DeleteCar(ctx context.Context, id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (r *inventoryRepository) ListCars(ctx context.Context, filter model.CarFilter, pagination model.PaginationParams, sort model.SortParams) ([]*model.Car, int, error) {
	//TODO implement me
	panic("implement me")
}

func (r *inventoryRepository) GetInventoryStats(ctx context.Context) (*model.InventoryStats, error) {
	//TODO implement me
	panic("implement me")
}
