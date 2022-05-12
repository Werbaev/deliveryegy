package postgres

import (
	"fmt"
	"os"
	"testing"

	"github.com/anmimos/delivery/config"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
)

func TestMain(m *testing.M) {
	var err error
	cfg := config.Load()
	conStr := fmt.Sprintf("host=%s port=%v user=%s password=%s dbname=%s sslmode=%s", cfg.PostgresHost,
		cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB, "disable")

	db, err = sqlx.Connect("pgx", conStr)
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())

}
