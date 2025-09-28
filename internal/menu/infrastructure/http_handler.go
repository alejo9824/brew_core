package infrastructure

import (
	"encoding/json"
	"net/http"

	"github.com/alejo9824/brew_core/internal/menu/application"
)

// CreateMenuItemRequest es un DTO para decodificar la petición JSON de creación.
// Usamos etiquetas `json` para mapear los campos del JSON a los campos de la struct.
type CreateMenuItemRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// Handler contiene el servicio de aplicación como dependencia.
type Handler struct {
	service application.MenuService
}

func NewHandler(service application.MenuService) *Handler {
	return &Handler{service: service}
}

// CreateMenuItem es el método del handler que maneja la petición HTTP.
func (h *Handler) CreateMenuItem(w http.ResponseWriter, r *http.Request) {
	// 1. Decodificar la petición JSON en nuestro DTO.
	var req CreateMenuItemRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "cuerpo de la petición inválido", http.StatusBadRequest)
		return
	}

	// 2. Llamar al servicio de aplicación con los datos validados/extraídos.
	item, err := h.service.CreateMenuItem(r.Context(), req.Name, req.Description, req.Price)
	if err != nil {
		// Aquí es donde nuestro middleware de errores (Fase 0) actuará en el futuro.
		// Por ahora, una respuesta de error simple.
		http.Error(w, "error interno del servidor", http.StatusInternalServerError)
		return
	}

	// 3. Codificar la respuesta a JSON y enviarla al cliente.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201 Created
	json.NewEncoder(w).Encode(item)
}
