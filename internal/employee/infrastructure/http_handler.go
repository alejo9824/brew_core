package infrastructure

import (
	"encoding/json"
	"net/http"

	"github.com/alejo9824/brew_core/internal/employee/application"
)

// HTTPHandler maneja las peticiones HTTP para el módulo de empleados.
// Es un Adaptador de Entrada en la Arquitectura Hexagonal.
type HTTPHandler struct {
	service *application.Service
}

// NewHTTPHandler crea una nueva instancia de HTTPHandler.
func NewHTTPHandler(service *application.Service) *HTTPHandler {
	return &HTTPHandler{
		service: service,
	}
}

// CreateEmployeeRequest es el DTO (Data Transfer Object) que define
// la estructura esperada del JSON en las peticiones POST.
type CreateEmployeeRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	RoleID    string `json:"role_id"`
}

// CreateEmployee maneja las peticiones POST para crear un nuevo empleado.
func (h *HTTPHandler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	// Solo aceptamos método POST
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Decodificamos el cuerpo de la petición en nuestro DTO
	var req CreateEmployeeRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Error al decodificar la petición: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Validación básica
	if req.FirstName == "" || req.LastName == "" || req.Email == "" || req.Password == "" || req.RoleID == "" {
		http.Error(w, "Todos los campos son requeridos", http.StatusBadRequest)
		return
	}

	// Llamamos al servicio de aplicación
	employee, err := h.service.CreateEmployee(
		r.Context(),
		req.FirstName,
		req.LastName,
		req.Email,
		req.Password,
		req.RoleID,
	)
	if err != nil {
		// TODO: En el futuro, implementaremos un mejor sistema de manejo de errores
		// que distinguirá entre diferentes tipos de errores y devolverá
		// códigos de estado HTTP apropiados.
		http.Error(w, "Error al crear el empleado: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Preparamos la respuesta
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201 Created

	// Codificamos la respuesta a JSON
	if err := json.NewEncoder(w).Encode(employee); err != nil {
		// Si fallara la codificación (muy improbable), log del error
		// pero ya no podemos cambiar el código de estado porque ya escribimos los headers
		// TODO: Implementar un sistema de logging apropiado
		return
	}
}
