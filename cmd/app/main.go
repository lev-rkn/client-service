package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"

	"CarFix/internal/config"
	"CarFix/internal/database"
	"CarFix/internal/transport"
)

func main() {
	// init configs
	os.Setenv("CONFIG_PATH", "configs/local.yaml")
	cfg := config.MustLoad()

	// connect to pgx
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, cfg.PostgresUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(ctx)
	clientRepo := &db.Database{
		Conn: conn,
		Ctx: ctx,
	}

	// run server
	transport.StartServer(clientRepo, cfg)
}

