package infrastructure

import (
	"context"
	"database/sql"

	"github.com/alejo9824/brew_core/internal/employee/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

// PostgresRepository es la implementación del EmployeeRepository para PostgreSQL.
// Es un Adaptador de Salida en la Arquitectura Hexagonal.
type PostgresRepository struct {
	db *pgxpool.Pool
}

// NewPostgresRepository es el constructor para PostgresRepository.
func NewPostgresRepository(db *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

// Save guarda un empleado en el repositorio.
// Por ahora, solo implementa la inserción. En el futuro, podría manejar una lógica de "upsert".
func (r *PostgresRepository) Save(ctx context.Context, employee *domain.Employee) error {
	query := `
        INSERT INTO employees (id, first_name, last_name, email, password, role_id, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    `

	_, err := r.db.Exec(ctx, query,
		employee.ID,
		employee.FirstName,
		employee.LastName,
		employee.Email,
		employee.Password,
		employee.RoleID,
		employee.CreatedAt,
		employee.UpdatedAt,
	)

	return err
}

// FindByID busca un empleado por su ID.
func (r *PostgresRepository) FindByID(ctx context.Context, id string) (*domain.Employee, error) {
	query := `
        SELECT id, first_name, last_name, email, password, role_id, created_at, updated_at
        FROM employees
        WHERE id = $1
    `

	var employee domain.Employee
	row := r.db.QueryRow(ctx, query, id)

	err := row.Scan(
		&employee.ID,
		&employee.FirstName,
		&employee.LastName,
		&employee.Email,
		&employee.Password,
		&employee.RoleID,
		&employee.CreatedAt,
		&employee.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			// Es una buena práctica devolver un error específico cuando no se encuentra el recurso.
			// Por ahora, devolvemos nil y el error original, pero en el futuro podríamos crear errores de dominio.
			return nil, err
		}
		// Otro tipo de error (ej. de conexión)
		return nil, err
	}

	return &employee, nil
}
