package api

import (
	"net/http"
	"strings"
)

// SetupRoutes configures all API routes
func (h *Handlers) SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// Tutorial endpoints
	mux.HandleFunc("/api/tutorials", h.ListTutorials)
	mux.HandleFunc("/api/tutorials/", h.handleTutorialRoutes)

	// Code execution
	mux.HandleFunc("/api/execute", h.ExecuteCode)

	// Progress tracking
	mux.HandleFunc("/api/progress", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h.GetProgress(w, r)
		case http.MethodPost:
			h.UpdateProgress(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/api/progress/section", h.MarkSectionComplete)

	// Exercises
	mux.HandleFunc("/api/exercises/", h.handleExerciseRoutes)

	return mux
}

// handleTutorialRoutes routes tutorial-specific endpoints with path parameters
func (h *Handlers) handleTutorialRoutes(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/tutorials/")
	parts := strings.Split(path, "/")

	if len(parts) == 0 || parts[0] == "" {
		http.Error(w, "tutorial ID required", http.StatusBadRequest)
		return
	}

	tutorialID := parts[0]

	if len(parts) > 1 && parts[1] == "sections" {
		h.GetTutorialSectionsByID(w, r, tutorialID)
		return
	}

	// Default: get full tutorial
	h.GetTutorialByID(w, r, tutorialID)
}

// handleExerciseRoutes routes exercise endpoints with path parameters
func (h *Handlers) handleExerciseRoutes(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/exercises/")
	parts := strings.Split(path, "/")

	if len(parts) == 0 || parts[0] == "" {
		http.Error(w, "tutorial ID required", http.StatusBadRequest)
		return
	}

	tutorialID := parts[0]
	h.GetExercisesByTutorialID(w, r, tutorialID)
}

// CORSMiddleware is a CORS middleware
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
