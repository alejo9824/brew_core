package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func newRouter(container *Container) *chi.Mux {
	r := chi.NewRouter()

	// Middlewares globales
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Endpoint de Health Check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// --- Grupo de Rutas para la API v1 ---
	// Agrupamos todas nuestras rutas de la API bajo /api/v1
	r.Route("/api/v1", func(r chi.Router) {
		// Grupo de rutas para el recurso "menu"
		r.Route("/menu", func(r chi.Router) {
			// Mapeamos el método POST a nuestro handler
			r.Post("/", container.MenuHandler.CreateMenuItem)
			// Aquí irían otras rutas como GET, PUT, DELETE...
			// r.Get("/{menuID}", container.MenuHandler.GetMenuItem)
		})

		// Aquí irían otros recursos como /orders, /tables, etc.
	})

	return r
}
