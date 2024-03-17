package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/caarlos0/env"
	"github.com/ramonmedeiros/key_value_store/internal/app/server"
)

type config struct {
	Port string `env:"PORT" envDefault:"8080"`
}

func main() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	logger.Enabled(context.Background(), slog.LevelError)

	httpServer := server.New(cfg.Port, logger)
	httpServer.Serve()
}
