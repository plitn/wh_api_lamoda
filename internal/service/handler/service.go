package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/plitn/wh_api_lamoda/internal/logger"
	"github.com/plitn/wh_api_lamoda/internal/model"
	"github.com/plitn/wh_api_lamoda/internal/service/warehouse"
)

type service struct {
	warehouse warehouse.Service
}

func New(whService warehouse.Service) *service {
	return &service{warehouse: whService}
}

func (s *service) ReserveProduct(w http.ResponseWriter, r *http.Request) {
	var req model.ReserveReq

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logger.Logger.Printf("error unmarshalling body: %v", err)
		return
	}
	err = s.warehouse.ReserveProducts(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Logger.Printf("cannot reserve products: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *service) UnReserveProduct(w http.ResponseWriter, r *http.Request) {
	var req model.ReserveReq

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeResponseJson(w, model.Response{Status: http.StatusBadRequest,
			Result: fmt.Sprintf("Cannot decode body, err: %v", err)})
		logger.Logger.Printf("Error decoding body: %v", err)
		return
	}
	err = s.warehouse.UnReserveProducts(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeResponseJson(w, model.Response{Status: http.StatusInternalServerError,
			Result: fmt.Sprintf("Cannot unreserve products, err: %v", err)})
		logger.Logger.Printf("Cannot unreserve products: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	writeResponseJson(w, model.Response{Status: http.StatusOK, Result: "Products unreserved successfully"})

}

func (s *service) GetWhQtyData(w http.ResponseWriter, r *http.Request) {
	whIdParam := r.URL.Query().Get("whId")
	if whIdParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		writeResponseJson(w, model.Response{Status: http.StatusBadRequest, Result: "Missing 'whId' query parameter"})
		return
	}

	whId, err := strconv.ParseInt(whIdParam, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeResponseJson(w, model.Response{Status: http.StatusBadRequest,
			Result: fmt.Sprintf("Invalid 'whId' parameter: %v", err)})
		return
	}

	qtyData, err := s.warehouse.GetProductsQty(whId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeResponseJson(w, model.Response{Status: http.StatusInternalServerError,
			Result: fmt.Sprintf("Cannot get wh qty data: %v", err)})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(qtyData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
