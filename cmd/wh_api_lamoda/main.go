package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"

	"github.com/plitn/wh_api_lamoda/internal/config"
	"github.com/plitn/wh_api_lamoda/internal/logger"
	"github.com/plitn/wh_api_lamoda/internal/repository"
	"github.com/plitn/wh_api_lamoda/internal/service/handler"
	"github.com/plitn/wh_api_lamoda/internal/service/warehouse"
)

func main() {
	cfg := config.LoadConfig()

	dbInst, err := sqlx.Open(cfg.DB.DriverName, cfg.DB.DSN)
	if err != nil {
		logger.Logger.Fatalf("failed to connect database: %v", err)
	}
	defer func() {
		err = dbInst.Close()
		if err != nil {
			logger.Logger.Println("can not close database client: %v", err)
		}
	}()

	repo := repository.NewRepository(dbInst)
	whService := warehouse.New(repo)
	handlerService := handler.New(whService)

	mux := chi.NewRouter()
	mux.Post("/api/reserve-products", handlerService.ReserveProduct)
	mux.Post("/api/unreserve-products", handlerService.UnReserveProduct)
	mux.Get("/api/get-wh-qty", handlerService.GetWhQtyData)

	httpServer := http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", 8080),
		Handler: mux,
	}

	logger.Logger.Println("listening to http://0.0.0.0:%d/ for debug http", 8080)
	if err = httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Logger.Println("can't start server: %v", err)
	}
}
