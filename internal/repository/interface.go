package repository

import (
	"context"

	"github.com/plitn/wh_api_lamoda/internal/model"
)

type Repository interface {
	UpdateReserveProducts(ctx context.Context, reserveData model.ReserveReq) error
	UpdateUnReserveProducts(ctx context.Context, reserveData model.ReserveReq) error
	GetProductsQtyInWh(ctx context.Context, whId int64) ([]model.WarehouseAllProductsQty, error)
}
