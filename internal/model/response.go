package model

type WhQtyResponse struct {
	ReservedQty   int64 `json:"reserved_qty"`
	UnReservedQty int64 `json:"unreserved_qty"`
}

type Response struct {
	Status int    `json:"status"`
	Result string `json:"result"`
}
