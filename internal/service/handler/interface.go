package handler

import "net/http"

type Handler interface {
	ReserveProduct(w http.ResponseWriter, r *http.Request)
	UnReserveProduct(w http.ResponseWriter, r *http.Request)
	GetWhQtyData(w http.ResponseWriter, r *http.Request)
}
