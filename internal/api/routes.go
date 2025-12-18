package api

import (
	"net/http"
)

// SetupRoutes configures all API routes
func (h *Handlers) SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	
	// Tutorial endpoints
	mux.HandleFunc("/api/tutorials", h.ListTutorials)
	mux.HandleFunc("/api/tutorial", h.GetTutorial)
	mux.HandleFunc("/api/tutorial/sections", h.GetTutorialSections)
	
	// Code execution
	mux.HandleFunc("/api/execute", h.ExecuteCode)
	
	// Progress tracking
	mux.HandleFunc("/api/progress", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			h.GetProgress(w, r)
		} else if r.Method == http.MethodPost {
			h.UpdateProgress(w, r)
		} else {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/api/progress/section", h.MarkSectionComplete)
	
	// Exercises
	mux.HandleFunc("/api/exercises", h.GetExercises)
	
	return mux
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

