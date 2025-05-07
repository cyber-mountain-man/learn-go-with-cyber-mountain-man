package router

import (
	"net/http"

	"github.com/cyber-mountain-man/learn-go-with-cyber-mountain-man/28-deployment/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Serve static files like index.html, styles, etc.
	fileServer(r, "/static", http.Dir("static"))

	// Redirect root to static index
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/static/index.html", http.StatusFound)
	})

	// Health check
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("✅ OK"))
	})

	// User management API for HTMX
	r.Route("/users", func(r chi.Router) {
		r.Get("/", handlers.ListUsersHTMX)           // GET /users
		r.Post("/", handlers.CreateUserHTMX)         // POST /users
		r.Get("/{id}/edit", handlers.EditUserFormHTMX) // GET /users/{id}/edit
		r.Put("/{id}", handlers.UpdateUserHTMX)      // PUT /users/{id}
		r.Delete("/{id}", handlers.DeleteUserHTMX)   // DELETE /users/{id}
	})

	return r
}

// fileServer serves static files from a directory at a given URL path.
func fileServer(r chi.Router, path string, root http.FileSystem) {
	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}

	fs := http.StripPrefix(path, http.FileServer(root))
	r.Get(path+"*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}
