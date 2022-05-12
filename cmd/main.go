package main

import (
	"fmt"

	"github.com/anmimos/delivery/api"
	"github.com/anmimos/delivery/config"
	"github.com/anmimos/delivery/pkg/logger"
	"github.com/anmimos/delivery/storage"
	"github.com/jmoiron/sqlx"
	_ "github.com/jackc/pgx/v4/stdlib"
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
