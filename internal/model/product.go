package model

type Product struct {
	ProductId int64   `json:"product_id" db:"product_id"`
	Width     float64 `json:"width" db:"width"`
	Height    float64 `json:"height" db:"height"`
	Depth     float64 `json:"depth" db:"depth"`
	Volume    float64 `json:"volume" db:"volume"`
	QtyTotal  int64   `json:"qty_total" db:"qty_total"`
}
