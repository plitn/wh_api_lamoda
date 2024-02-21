package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/plitn/wh_api_lamoda/internal/model"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) UpdateReserveProducts(ctx context.Context, reserveData model.ReserveReq) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	var isActiveWh bool
	err = tx.QueryRowContext(ctx, "SELECT is_active FROM warehouses WHERE wh_id = $1",
		reserveData.WhId).Scan(&isActiveWh)
	if err != nil {
		return err
	}

	for _, product := range reserveData.ProductsData {
		var currentQuantity int64
		err = tx.QueryRowContext(ctx,
			"SELECT product_qty FROM wh_product WHERE product_id = $1 and wh_id = $2 FOR UPDATE",
			product.ProductId, reserveData.WhId).Scan(&currentQuantity)
		if err != nil {
			return err
		}

		if currentQuantity < product.Qty {
			return fmt.Errorf("not enough quantity available for product_id = %d", product.ProductId)
		}

		_, err = tx.ExecContext(ctx,
			"UPDATE wh_product SET product_qty = product_qty - $1, reserved_qty = reserved_qty + $1 "+
				"WHERE product_id = $2 and wh_id = $3",
			product.Qty, product.ProductId, reserveData.WhId)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) UpdateUnReserveProducts(ctx context.Context, reserveData model.ReserveReq) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	for _, product := range reserveData.ProductsData {
		var currentQuantity int64
		err = tx.QueryRowContext(ctx,
			"SELECT reserved_qty FROM wh_product WHERE product_id = $1 and wh_id = $2 FOR UPDATE",
			product.ProductId, reserveData.WhId).Scan(&currentQuantity)
		if err != nil {
			return err
		}

		if currentQuantity < product.Qty {
			return fmt.Errorf("not enough reserved quantity available for product_id = %d", product.ProductId)
		}

		_, err = tx.ExecContext(ctx,
			"UPDATE wh_product SET product_qty = product_qty + $1, reserved_qty = reserved_qty - $1 "+
				"WHERE product_id = $2 and wh_id = $3",
			product.Qty, product.ProductId, reserveData.WhId)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetProductsQtyInWh(ctx context.Context, whId int64) ([]model.WarehouseAllProductsQty, error) {
	var qtyData []model.WarehouseAllProductsQty
	//tx, err := r.db.Begin()
	//if err != nil {
	//	return nil, err
	//}
	//defer func() {
	//	_ = tx.Rollback()
	//}()
	query := "SELECT product_id, product_qty, reserved_qty FROM wh_product WHERE wh_id = $1"
	err := r.db.SelectContext(ctx, &qtyData, query, whId)
	if err != nil {
		return nil, err
	}
	return qtyData, nil
	//_, err = tx.ExecContext(ctx, "LOCK TABLE wh_product IN SHARE MODE")
	//if err != nil {
	//	return nil, err
	//}
	//
	//rows, err := tx.QueryContext(ctx,
	//	"SELECT product_id, product_qty, reserved_qty FROM wh_product WHERE wh_id = $1", whId)
	//if err != nil {
	//	return nil, err
	//}
	//defer func() {
	//	_ = rows.Close()
	//}()
	//
	//for rows.Next() {
	//	var productQuantityRow model.WarehouseAllProductsQty
	//	if err := rows.Scan(&productQuantityRow); err != nil {
	//		return nil, err
	//	}
	//	qtyData = append(qtyData, productQuantityRow)
	//}
	//
	//if err := rows.Err(); err != nil {
	//	return nil, err
	//}
	//
	//err = tx.Commit()
	//if err != nil {
	//	return nil, err
	//}
	//
	//return qtyData, nil
}
