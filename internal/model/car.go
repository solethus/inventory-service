package model

import (
	"time"

	"github.com/google/uuid"
)

type Car struct {
	ID           uuid.UUID `json:"id"`
	Make         string    `json:"make"`
	Model        string    `json:"model"`
	Year         int       `json:"year"`
	Color        string    `json:"color"`
	Price        float64   `json:"price"`
	Stock        int       `json:"stock"`
	VIN          string    `json:"vin"`
	FuelType     string    `json:"fuel_type"`
	Transmission string    `json:"transmission"`
	EngineSize   float64   `json:"engine_size"`
	Mileage      int       `json:"mileage"`
	Features     []string  `json:"features"`
	Status       string    `json:"status"`
	LastUpdated  time.Time `json:"last_updated"`
	DateAdded    time.Time `json:"date_added"`
}

type CarFilter struct {
	Make          string
	Model         string
	MinYear       int
	MaxYear       int
	MinPrice      float64
	MaxPrice      float64
	Colors        []string
	FuelTypes     []string
	Transmissions []string
	InStock       *bool
}

type PaginationParams struct {
	Page     int
	PageSize int
}

type SortParams struct {
	Field string
	Order string // "asc" or "desc"
}

type InventoryStats struct {
	TotalCars         int
	TotalValue        float64
	AveragePrice      float64
	MostCommonMake    string
	MostCommonModel   string
	LowStockThreshold int
	LowStockCount     int
}
