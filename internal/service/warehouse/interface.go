package warehouse

import "github.com/plitn/wh_api_lamoda/internal/model"

type Service interface {
	ReserveProducts(req model.ReserveReq) error
	UnReserveProducts(req model.ReserveReq) error
	GetProductsQty(whId int64) (model.WhQtyResponse, error)
}
