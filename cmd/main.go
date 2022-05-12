package main

import (
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/werbaev/deliveryegy/api"
	"github.com/werbaev/deliveryegy/config"
	"github.com/werbaev/deliveryegy/pkg/logger"
	"github.com/werbaev/deliveryegy/storage"
)

func main() {

	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "delivery")

	conStr := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=%s", cfg.PostgresHost,
		cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB, "disable")
	db, err := sqlx.Connect("pgx", conStr)
	if err != nil {
		panic(err)
	}

	storageI := storage.NewStoragePg(db)

	server := api.New(&api.RouterOptions{
		Log:     log,
		Cfg:     &cfg,
		Storage: storageI,
	})

	server.Run(cfg.ServicePort)

	log.Info("delivery running...")
}
