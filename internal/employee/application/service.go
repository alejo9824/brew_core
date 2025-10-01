package application

import (
	"context"

	"github.com/alejo9824/brew_core/internal/employee/domain"
)

// Service es el servicio de aplicación para el módulo de empleados.
// Orquesta los casos de uso y depende de los puertos definidos.
type Service struct {
	employeeRepo EmployeeRepository
}

// NewService es el constructor para el servicio de aplicación.
// Recibe las dependencias (implementaciones de los puertos) y devuelve
// una nueva instancia del servicio.
func NewService(employeeRepo EmployeeRepository) *Service {
	return &Service{
		employeeRepo: employeeRepo,
	}
}

// CreateEmployee es el caso de uso para crear un nuevo empleado.
func (s *Service) CreateEmployee(ctx context.Context, firstName, lastName, email, password, roleID string) (*domain.Employee, error) {
	// 1. Usamos el constructor del dominio para crear la entidad.
	// Esto asegura que todas las reglas de negocio del dominio se cumplan.
	employee, err := domain.NewEmployee(firstName, lastName, email, password, roleID)
	if err != nil {
		return nil, err // Si el constructor fallara, propagamos el error.
	}

	// TODO: Aquí irá la lógica para hashear la contraseña del empleado
	// antes de guardarla. Lo haremos en una fase posterior.

	// 2. Usamos el puerto del repositorio para guardar la entidad.
	// El servicio no sabe cómo se guarda, solo delega esa responsabilidad.
	err = s.employeeRepo.Save(ctx, employee)
	if err != nil {
		return nil, err // Podría ser un error de base de datos, de duplicado, etc.
	}

	// 3. Devolvemos la entidad creada.
	return employee, nil
}
