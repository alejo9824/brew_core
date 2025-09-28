package application

import (
	"context"

	"github.com/alejo9824/brew_core/internal/menu/domain"
)

// Repository define el contrato que la capa de infraestructura debe implementar
// para la persistencia de datos de MenuItem. Es nuestro "puerto" de salida.
type Repository interface {
	Save(ctx context.Context, item domain.MenuItem) error
}
