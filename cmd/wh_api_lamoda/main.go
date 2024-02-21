package main

import (
	"errors"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // <------------ here
	"github.com/plitn/wh_api_lamoda/internal/config"
	"github.com/plitn/wh_api_lamoda/internal/logger"
	"github.com/plitn/wh_api_lamoda/internal/repository"
	"github.com/plitn/wh_api_lamoda/internal/service/handler"
	"github.com/plitn/wh_api_lamoda/internal/service/warehouse"
	"net/http"
	"time"
)

func main() {
	cfg := config.LoadConfig()

	dbInst, err := sqlx.Open(cfg.DB.DriverName, cfg.DB.DSN)
	dbInst.SetConnMaxLifetime(10 * time.Minute)
	if err != nil {
		logger.Logger.Println("failed to connect database: %v", err)
	}
	err = dbInst.Ping()
	if err != nil {
		logger.Logger.Println(err)
		return
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

	logger.Logger.Printf("listening to http://0.0.0.0:%d/ for debug http", 8080)
	if err = httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Logger.Println("can't start server: %v", err)
	}
}
