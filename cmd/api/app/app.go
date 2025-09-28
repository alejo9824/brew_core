package app

import (
	"context"
	"net/http"

	"github.com/alejo9824/brew_core/internal/shared/config"
	"github.com/alejo9824/brew_core/internal/shared/database"
	"github.com/rs/zerolog/log"
)

type App struct {
	HTTPServer *http.Server
}

func New(ctx context.Context) (*App, func()) {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal().Err(err).Msg("Fallo al cargar la configuración")
	}

	dbpool, err := database.NewConnection(ctx, cfg.DB.ConnectionString)
	if err != nil {
		log.Fatal().Err(err).Msg("Fallo al conectar a la DB")
	}

	container, err := newContainer(ctx, dbpool)
	if err != nil {
		log.Fatal().Err(err).Msg("Fallo al construir el contenedor")
	}

	router := newRouter(container)

	server := &http.Server{
		Addr:    ":" + cfg.Server.Port,
		Handler: router,
	}

	cleanup := func() {
		log.Info().Msg("Ejecutando cleanup...")
		dbpool.Close()
	}

	return &App{HTTPServer: server}, cleanup
}
