package app

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	// 1. Importamos los paquetes del módulo de menú que necesitamos
	employeeApp "github.com/alejo9824/brew_core/internal/employee/application"
	employeeInfra "github.com/alejo9824/brew_core/internal/employee/infrastructure"
	menuApp "github.com/alejo9824/brew_core/internal/menu/application"
	menuInfra "github.com/alejo9824/brew_core/internal/menu/infrastructure"
)

// Container aloja todas las dependencias de la aplicación.
// Específicamente, expone los handlers que el router necesitará.
type Container struct {
	// 2. Añadimos el handler del menú al contenedor
	MenuHandler     *menuInfra.Handler
	EmployeeHandler *employeeInfra.HTTPHandler
}

// newContainer es privado y se encarga de construir y cablear las dependencias.
func newContainer(ctx context.Context, db *pgxpool.Pool) (*Container, error) {
	// --- Línea de Ensamblaje del Módulo de Menú ---

	// 3. Creamos la instancia del Repositorio (el adaptador de base de datos)
	// Le inyectamos la dependencia del pool de la base de datos.
	menuRepository := menuInfra.NewPostgresRepository(db)

	// 4. Creamos la instancia del Servicio (el núcleo de la lógica)
	// Le inyectamos la dependencia del repositorio (como una interfaz).
	menuService := menuApp.NewService(menuRepository)

	// 5. Creamos la instancia del Handler (el adaptador de API)
	// Le inyectamos la dependencia del servicio.
	menuHandler := menuInfra.NewHandler(menuService)

	// --- Módulo Employee ---
	employeeRepository := employeeInfra.NewPostgresRepository(db)
	employeeService := employeeApp.NewService(employeeRepository)
	employeeHandler := employeeInfra.NewHTTPHandler(employeeService)

	// 6. Devolvemos el contenedor con el handler listo para ser usado.
	return &Container{
		MenuHandler:     menuHandler,
		EmployeeHandler: employeeHandler,
	}, nil
}
