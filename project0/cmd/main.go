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
	"project0/models"
)

func main() {

	//Init config and logger

	cfg := config.Mustload()
	_ = cfg
	log := logger.SetLogger()
	log.Info("start app", slog.String("env", "local"))

	// Init storage
	//path := "user=gros password=grosBerg22 dbname=wb_db sslmode=disable"
	stor, err := storage.New(cfg.StoragePath)
	if err != nil {
		log.Error("fail to connect postgres", err)
	}
	log.Info("init storage success")
	_ = stor

	//Raising up the cache from DB
	cache_map := map[string]models.Order{}
	Cache := cache.MakeCache(stor, cache_map)
	_ = Cache

	//Read data from stream to DB and Cache
	//err = stream.SavetoDBandCache(stor, Cache)
	//if err != nil {
	//	log.Error("MAIN save error ", err)
	//}

	router := chi.NewRouter()
	router.Get("/", v1.OrderGetter(log, Cache))

	log.Info("starting server", slog.String("address", cfg.Address))

	server.ServerRun(log, cfg, router)

}
