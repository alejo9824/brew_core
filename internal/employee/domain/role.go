package domain

import "github.com/google/uuid"

// Role representa el rol de un empleado dentro del sistema.
// Define qué permisos y accesos tiene un empleado.
type Role struct {
	ID   string
	Name string
}

// NewRole es el constructor para la entidad Role.
func NewRole(name string) (*Role, error) {
	id := uuid.New().String()
	return &Role{
		ID:   id,
		Name: name,
	}, nil
}
