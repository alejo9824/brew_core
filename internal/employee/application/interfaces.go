package application

import (
	"context"

	"github.com/alejo9824/brew_core/internal/employee/domain"
)

// EmployeeRepository es un Puerto (Port) en la Arquitectura Hexagonal.
// Define el contrato que cualquier adaptador de persistencia debe cumplir
// para interactuar con las entidades de Employee.
// La capa de aplicación depende de esta interfaz, no de una implementación concreta.
type EmployeeRepository interface {
	// Save guarda un empleado en el repositorio.
	// Puede ser una inserción (si es nuevo) o una actualización (si ya existe).
	Save(ctx context.Context, employee *domain.Employee) error

	// FindByID busca un empleado por su ID.
	FindByID(ctx context.Context, id string) (*domain.Employee, error)
}
