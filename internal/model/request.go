package model

type ReserveReq struct {
	WhId         int64        `json:"wh_id" db:"wh_id"`
	ProductsData []ProductQty `json:"products_data"`
}

type ProductQty struct {
	ProductId int64 `json:"product_id" db:"product_id"`
	Qty       int64 `json:"qty"`
}
