package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"letual/internal/config"
	"log/slog"
	"strings"
)

type Storage struct {
	client *pgx.Conn
	ctx    context.Context
	logger *slog.Logger
}

func NewStorage(ctx context.Context, log *slog.Logger, cfg *config.Postgres) (*Storage, error) {
	const fn = "NewStorage"

	dataBaseURL := connectionString(cfg)

	conn, err := pgx.Connect(ctx, dataBaseURL)
	if err != nil {
		return nil, fmt.Errorf("%s : %w", fn, err)
	}

	return &Storage{
		client: conn,
		ctx:    ctx,
		logger: log,
	}, nil
}

func connectionString(cfg *config.Postgres) string {
	const baseURL = "postgresql://"
	const colonSep = ":"
	const slashSep = "/"
	const sep = "@"

	var builder strings.Builder

	if cfg.Username == "" && cfg.Password == "" {
		builder.Grow(len(baseURL) + len(cfg.Host) + len(slashSep) + len(cfg.Dbname))
		builder.WriteString(baseURL)
		builder.WriteString(cfg.Host)
		builder.WriteString(slashSep)
		builder.WriteString(cfg.Dbname)

		return builder.String()
	}

	builder.Grow(len(baseURL) +
		len(cfg.Username) +
		len(colonSep) +
		len(cfg.Password) +
		len(sep) +
		len(cfg.Host) +
		len(slashSep) +
		len(cfg.Dbname))
	builder.WriteString(baseURL)
	builder.WriteString(cfg.Username)
	builder.WriteString(colonSep)
	builder.WriteString(cfg.Password)
	builder.WriteString(sep)
	builder.WriteString(cfg.Host)
	builder.WriteString(slashSep)
	builder.WriteString(cfg.Dbname)

	return builder.String()
}
