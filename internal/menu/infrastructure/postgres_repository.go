package infrastructure

import (
	"context"
	"fmt"

	"github.com/alejo9824/brew_core/internal/menu/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

// PostgresRepository es la implementación del repositorio de menú para PostgreSQL.
// Contiene una dependencia a nuestro pool de conexiones de la base de datos.
type PostgresRepository struct {
	db *pgxpool.Pool
}

// NewPostgresRepository crea una nueva instancia del repositorio de Postgres.
func NewPostgresRepository(db *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{db: db}
}

// Save cumple con el contrato de la interfaz application.Repository.
func (r *PostgresRepository) Save(ctx context.Context, item domain.MenuItem) error {
	// Definimos la query SQL para insertar un nuevo registro.
	query := `INSERT INTO menu_items (id, name, description, price, is_available, created_at, updated_at)
                          VALUES ($1, $2, $3, $4, $5, $6, $7)`

	// Ejecutamos la query usando el pool de conexiones.
	// Usamos Exec para queries que no devuelven filas.
	_, err := r.db.Exec(ctx, query,
		item.ID,
		item.Name,
		item.Description,
		item.Price,
		item.IsAvailable,
		item.CreatedAt,
		item.UpdatedAt,
	)

	// Si hay un error, lo envolvemos con más contexto para facilitar la depuración.
	if err != nil {
		return fmt.Errorf("error al guardar el item del menú: %w", err)
	}

	return nil
}
