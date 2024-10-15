package main

import (
	"context"
	"letual/internal/modules/users"
	"log/slog"
	"os"

	"letual/internal/adapters/db/postgres"
	"letual/internal/app/http/letual"
	"letual/internal/config"
	"letual/internal/modules/product"
)

func main() {
	ctx := context.Background()

	cfg := config.MustLoad()

	log := slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)

	storage, err := postgres.NewStorage(ctx, log, &cfg.Postgres)
	if err != nil {
		log.Error("failed to init storage", err)
		os.Exit(1)
	}

	productModule := product.NewProduct(ctx, log, storage)
	userModule := users.NewUser(ctx, log, storage)

	srv := letual.NewServer(ctx, log, &cfg.HTTPServer, productModule, userModule)

	log.Info("starting server", slog.String("address", cfg.HTTPServer.Address))

	if err := srv.Run(); err != nil {
		log.Error("failed to start server")
	}

	log.Error("server stopped")
}
