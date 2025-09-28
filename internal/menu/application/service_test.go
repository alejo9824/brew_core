package application

import (
	"context"
	"errors"
	"testing"

	"github.com/alejo9824/brew_core/internal/menu/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockRepository es nuestro doble de prueba para la interfaz Repository.
// Embebemos mock.Mock de testify para obtener toda su funcionalidad.
type MockRepository struct {
	mock.Mock
}

// Implementamos el método Save de la interfaz Repository.
func (m *MockRepository) Save(ctx context.Context, item domain.MenuItem) error {
	// Le decimos a testify/mock que registre esta llamada y sus argumentos.
	args := m.Called(ctx, item)
	// Devolvemos lo que le hayamos configurado en el test (nil para éxito, un error para fallo).
	return args.Error(0)
}

func TestCreateMenuItem_Success(t *testing.T) {
	// 1. Arrange (Preparar)
	mockRepo := new(MockRepository)
	// Creamos una instancia de nuestro servicio, inyectando el mock.
	service := NewService(mockRepo)

	// Configuramos la expectativa para nuestro mock.
	// Le decimos: "Espero que el método 'Save' sea llamado con cualquier contexto
	// y cualquier MenuItem. Cuando eso ocurra, no devuelvas ningún error (nil)".
	mockRepo.On("Save", mock.Anything, mock.AnythingOfType("domain.MenuItem")).Return(nil)

	// 2. Act (Actuar)
	// Llamamos al método que estamos probando.
	createdItem, err := service.CreateMenuItem(context.Background(), "Café Americano", "Café de grano", 2.50)

	// 3. Assert (Verificar)
	// Usamos testify/assert para verificaciones limpias.
	assert.NoError(t, err)                              // Verificamos que no hubo error.
	assert.NotNil(t, createdItem)                       // Verificamos que el item devuelto no es nulo.
	assert.Equal(t, "Café Americano", createdItem.Name) // Verificamos que el nombre es correcto.
	assert.NotEmpty(t, createdItem.ID)                  // Verificamos que se generó un ID.

	// Verificamos que todas las expectativas que configuramos en el mock se cumplieron.
	mockRepo.AssertExpectations(t)
}

func TestCreateMenuItem_RepositoryError(t *testing.T) {
	// 1. Arrange
	mockRepo := new(MockRepository)
	service := NewService(mockRepo)

	// Configuramos el mock para que devuelva un error cuando se llame a Save.
	expectedError := errors.New("error de base de datos simulado")
	mockRepo.On("Save", mock.Anything, mock.AnythingOfType("domain.MenuItem")).Return(expectedError)

	// 2. Act
	_, err := service.CreateMenuItem(context.Background(), "Tarta de Queso", "Tarta casera", 4.50)

	// 3. Assert
	assert.Error(t, err)                // Verificamos que sí hubo un error.
	assert.Equal(t, expectedError, err) // Verificamos que el error es el que esperábamos.
	mockRepo.AssertExpectations(t)
}
