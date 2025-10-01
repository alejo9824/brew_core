package ierrors

import "errors"

// Definimos errores comunes de la aplicación.
// Esto nos permite comparar errores de forma estandarizada.
var (
	ErrNotFound     = errors.New("recurso no encontrado")
	ErrInvalidInput = errors.New("input inválido")
	ErrUnauthorized = errors.New("no autorizado")
	ErrForbidden    = errors.New("prohibido")
	ErrConflict     = errors.New("conflicto de datos") // Ej. violación de restricción UNIQUE
	// Puedes añadir más errores específicos de tu dominio aquí a medida que los necesites.
)

// Is es una función de conveniencia para comprobar si un error es de un tipo específico de ierrors.
// Es un wrapper para errors.Is, que es la forma idiomática de Go para comparar errores.
func Is(err, target error) bool {
	return errors.Is(err, target)
}
