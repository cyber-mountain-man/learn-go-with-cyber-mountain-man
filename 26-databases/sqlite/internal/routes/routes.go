package routes

import (
	"database/sql"
	"net/http"

	"sqlite/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Register sets up the application's routes and middlewares.
func Register(db *sql.DB) http.Handler {
	r := chi.NewRouter()

	// -----------------------------
	// 🔧 GLOBAL MIDDLEWARES
	// -----------------------------
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// -----------------------------
	// 👤 USER ROUTES
	// -----------------------------
	userHandler := handlers.NewUserHandler(db)
	r.Route("/users", userHandler.RegisterRoutes)

	// -----------------------------
	// 🏠 DEFAULT ROUTE
	// -----------------------------
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("🏁 Welcome to the SQLite API — Go + Chi + SQL!"))
	})

	return r
}
