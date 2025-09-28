package application

import (
	"context"
	"time"

	"github.com/alejo9824/brew_core/internal/menu/domain"
	"github.com/google/uuid"
)

// MenuService define el contrato para el servicio de menú.
type MenuService interface {
	CreateMenuItem(ctx context.Context, name, description string, price float64) (domain.MenuItem, error)
}

// Service encapsula la lógica de negocio para el menú.
// Recibe el repositorio como una dependencia a través de su interfaz.
type Service struct {
	repo Repository
}

// NewService es el constructor para nuestro servicio.
// Así es como implementamos la Inyección de Dependencias.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// CreateMenuItem es el caso de uso para crear un nuevo item en el menú.
func (s *Service) CreateMenuItem(ctx context.Context, name, description string, price float64) (domain.MenuItem, error) {
	// Aquí es donde vivirían las validaciones de negocio complejas.
	// Por ejemplo: if price <= 0 { return domain.MenuItem{}, errors.New("el precio debe ser positivo") }

	// Creamos la entidad de dominio.
	item := domain.MenuItem{
		ID:          uuid.NewString(), // Generamos un ID único universal.
		Name:        name,
		Description: description,
		Price:       price,
		IsAvailable: true, // Por defecto, un nuevo item está disponible.
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Llamamos al método Save de nuestro repositorio (el puerto).
	// El servicio no sabe ni le importa cómo se implementa este Save.
	if err := s.repo.Save(ctx, item); err != nil {
		return domain.MenuItem{}, err // Simplemente propagamos el error si la persistencia falla.
	}

	// Devolvemos la entidad recién creada.
	return item, nil
}
