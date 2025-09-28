package infrastructure

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alejo9824/brew_core/internal/menu/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockService es nuestro doble de prueba para la interfaz del servicio.
// Nota: Nuestro servicio es una struct, no una interfaz. Para testear, podemos
// definir una interfaz que nuestro servicio cumpla, o mockear la struct directamente.
// Por simplicidad aquí, crearemos una interfaz localmente para el mock.
type MockMenuService struct {
	mock.Mock
}

func (m *MockMenuService) CreateMenuItem(ctx context.Context, name, description string, price float64) (domain.MenuItem, error) {
	args := m.Called(ctx, name, description, price)
	return args.Get(0).(domain.MenuItem), args.Error(1)
}

func TestCreateMenuItemHandler_Success(t *testing.T) {
	// 1. Arrange
	mockService := new(MockMenuService)
	// El handler real necesita un application.Service, no nuestro mock.
	// Para este test, creamos un handler con el mock. En un proyecto real,
	// el servicio definiría una interfaz que el mock implementaría.
	// Por ahora, vamos a adaptar nuestro handler para que acepte una interfaz.
	// (Este es un punto de refactorización que el testing revela!)
	// ---
	// Vamos a asumir que hemos refactorizado el handler para aceptar una interfaz:
	// func NewHandler(service MenuServiceInterface) *Handler
	// ---
	handler := NewHandler(mockService) // Asumiendo el refactor

	// Configuramos el mock del servicio
	expectedItem := domain.MenuItem{ID: "test-uuid", Name: "Té Verde"}
	mockService.On("CreateMenuItem", mock.Anything, "Té Verde", "Té japonés", 3.0).Return(expectedItem, nil)

	// Creamos el cuerpo de la petición JSON
	requestBody, _ := json.Marshal(map[string]interface{}{
		"name":        "Té Verde",
		"description": "Té japonés",
		"price":       3.0,
	})

	// Creamos una petición HTTP simulada
	req := httptest.NewRequest("POST", "/api/v1/menu", bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	// Creamos un "Response Recorder" para grabar la respuesta del handler
	rr := httptest.NewRecorder()

	// 2. Act
	// Llamamos directamente al método del handler
	handler.CreateMenuItem(rr, req)

	// 3. Assert
	assert.Equal(t, http.StatusCreated, rr.Code) // Verificamos el código de estado

	// Verificamos el cuerpo de la respuesta
	var responseItem domain.MenuItem
	err := json.Unmarshal(rr.Body.Bytes(), &responseItem)
	assert.NoError(t, err)
	assert.Equal(t, expectedItem.ID, responseItem.ID)
	assert.Equal(t, expectedItem.Name, responseItem.Name)

	mockService.AssertExpectations(t)
}
