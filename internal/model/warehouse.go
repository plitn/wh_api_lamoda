package model

import "time"

type Warehouse struct {
	WhId       int64     `json:"wh_id" db:"wh_id"`
	Name       string    `json:"name" db:"name"`
	IsActive   bool      `json:"is_active" db:"is_active"`
	IsActiveDt time.Time `json:"is_active_dt" db:"is_active_dt"`
}

type WarehouseAllProductsQty struct {
	ProductId     int64 `json:"product_id" db:"product_id"`
	ReservedQty   int64 `json:"reserved_qty" db:"reserved_qty"`
	UnReservedQty int64 `json:"product_qty" db:"product_qty"`
}
