package router

import (
	"net/http"

	"github.com/cyber-mountain-man/learn-go-with-cyber-mountain-man/28-deployment/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// SetupRouter defines all routes for the application and returns the configured router.
func SetupRouter() http.Handler {
	r := chi.NewRouter()

	// Log each request to the console for debugging
	r.Use(middleware.Logger)

	// Serve static files (index.html, styles, etc.) from the "static" folder at "/static"
	fileServer(r, "/static", http.Dir("static"))

	// Redirect root "/" to the main index.html page
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/static/index.html", http.StatusFound)
	})

	// Simple health check route (e.g. for uptime monitoring)
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("✅ OK"))
	})

	// Define routes under the "/users" group
	r.Route("/users", func(r chi.Router) {
		r.Get("/", handlers.ListUsersHTMX)              // Load user list (HTML)
		r.Post("/", handlers.CreateUserHTMX)            // Create new user (HTMX form POST)
		r.Get("/{id}/edit", handlers.EditUserFormHTMX)  // Load user edit form (HTMX)
		r.Put("/{id}", handlers.UpdateUserHTMX)         // Update user (HTMX form PUT)
		r.Delete("/{id}", handlers.DeleteUserHTMX)      // Delete user
	})

	return r
}

// fileServer serves static files from the given root folder.
// It safely handles trailing slashes and serves files like CSS, JS, HTML from the /static URL path.
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

/*
 Blurb: Purpose of router.go
This file is responsible for configuring all the HTTP routes of your Go application using the Chi router. It:

Registers middleware for logging requests.

Serves static assets like HTML, CSS, and JS.

Redirects the root path to your frontend's index.html.

Defines RESTful endpoints for managing users via HTMX (GET, POST, PUT, DELETE).

Supports health monitoring through a lightweight /health endpoint.

This modular routing setup makes your app scalable and easy to debug or extend.
*/