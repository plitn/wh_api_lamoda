package warehouse

import (
	"context"
	"fmt"

	"github.com/plitn/wh_api_lamoda/internal/logger"
	"github.com/plitn/wh_api_lamoda/internal/model"
	"github.com/plitn/wh_api_lamoda/internal/repository"
)

type service struct {
	repository repository.Repository
}

func New(repo repository.Repository) *service {
	return &service{
		repository: repo,
	}
}

func (s *service) ReserveProducts(req model.ReserveReq) error {
	ctx := context.Background()
	err := s.repository.UpdateReserveProducts(ctx, req)
	if err != nil {
		logger.Logger.Printf("cannot reserve %d products in warehouse=%d, err: %v",
			len(req.ProductsData), req.WhId, err)
		return fmt.Errorf("cannot reserve %d products in warehouse=%d, err: %v",
			len(req.ProductsData), req.WhId, err)
	}
	return err
}

func (s *service) UnReserveProducts(req model.ReserveReq) error {
	ctx := context.Background()
	err := s.repository.UpdateUnReserveProducts(ctx, req)
	if err != nil {
		logger.Logger.Printf("cannot unreserve %d products in warehouse=%d, err: %v",
			len(req.ProductsData), req.WhId, err)
		return fmt.Errorf("cannot unreserve %d products in warehouse=%d, err: %v",
			len(req.ProductsData), req.WhId, err)
	}
	return err
}

func (s *service) GetProductsQty(whId int64) (model.WhQtyResponse, error) {
	ctx := context.Background()
	qtyData, err := s.repository.GetProductsQtyInWh(ctx, whId)
	if err != nil {
		logger.Logger.Printf("cannot get products qty in warehouse=%d, err: %v",
			whId, err)
		return model.WhQtyResponse{}, fmt.Errorf("cannot get products qty in warehouse=%d, err: %v",
			whId, err)
	}
	var whQtyData model.WhQtyResponse
	for _, qtyRow := range qtyData {
		whQtyData.ReservedQty += qtyRow.ReservedQty
		whQtyData.UnReservedQty += qtyRow.UnReservedQty
	}
	return whQtyData, err
}
