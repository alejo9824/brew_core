package domain

import "time"

// MenuItem es la representación de un producto en el menú.
// Esta es una entidad de dominio pura. No contiene etiquetas de JSON o DB.
type MenuItem struct {
	ID          string
	Name        string
	Description string
	Price       float64
	IsAvailable bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
