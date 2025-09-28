package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alejo9824/brew_core/cmd/api/app"
	"github.com/rs/zerolog/log"
)

func main() {
	application, cleanup := app.New(context.Background())
	defer cleanup()

	go func() {
		log.Info().Msgf("Servidor iniciado en %s", application.HTTPServer.Addr)
		if err := application.HTTPServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("El servidor HTTP falló")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Warn().Msg("Señal de apagado recibida...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := application.HTTPServer.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("Forzando el apagado del servidor")
	}
	log.Info().Msg("Servidor apagado correctamente.")
}
