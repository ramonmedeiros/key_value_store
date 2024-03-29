package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/caarlos0/env"
	"github.com/ramonmedeiros/key_value_store/internal/app/server"
	"github.com/ramonmedeiros/key_value_store/internal/hash"
	"github.com/ramonmedeiros/key_value_store/internal/keystore"
)

type config struct {
	Port  string `env:"PORT" envDefault:"8080"`
	Nodes int    `env:"NODES" envDefault:"4"`
}

func main() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	logger.Enabled(context.Background(), slog.LevelError)

	keyStore, err := keystore.New(logger, hash.New(), cfg.Nodes)
	if err != nil {
		os.Exit(1)
	}

	httpServer := server.New(cfg.Port, logger, keyStore)
	httpServer.Serve()
}
