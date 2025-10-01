package domain

import (
	"time"

	"github.com/google/uuid"
)

type Employee struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Password  string // ¡Importante! Este será el hash, nunca el texto plano.
	RoleID    string // Referencia al ID del rol que ocupa.
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewEmployee(firstName, lastName, email, password, roleID string) (*Employee, error) {
	// Aquí podríamos añadir validaciones en el futuro.
	// Por ejemplo, verificar que el email tenga un formato válido.
	// Si algo no es válido, devolveríamos un error.

	id := uuid.New().String()
	now := time.Now().UTC()

	return &Employee{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password, // Recordatorio: esto será un hash.
		RoleID:    roleID,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
