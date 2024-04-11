package main

import (
	"github.com/go-chi/chi/v5"
	"log/slog"
	v1 "project0/internal/api/v1"
	"project0/internal/cache"
	"project0/internal/config"
	"project0/internal/logger"
	"project0/internal/server"
	"project0/internal/storage"
	"project0/internal/stream"
	"project0/models"
)

func main() {

	//Init config and logger

	cfg := config.Mustload()
	log := logger.SetLogger()

	log.Info("init config & logger success")

	// Init storage

	stor, err := storage.New(cfg.StoragePath)
	if err != nil {
		log.Error("fail to connect postgres", err)
	}

	log.Info("init storage success")

	//Raising up the cache from DB
	cacheMap := map[string]models.Order{}
	Cache := cache.MakeCache(stor, cacheMap)

	log.Info("Raise cache success")

	//Read data from stream to DB and Cache
	err = stream.SavetoDBandCache(stor, Cache)
	if err != nil {
		log.Error("MAIN save error ", err)
	}
	log.Info("read from stream success")

	//Set Routers

	router := chi.NewRouter()
	router.Get("/", v1.OrderGetter(log, Cache))

	// Start server

	log.Info("starting server", slog.String("address", cfg.Address))

	server.ServerRun(log, cfg, router)

}
