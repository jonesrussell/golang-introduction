package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jonesrussell/go-fundamentals-best-practices/internal/executor"
	"github.com/jonesrussell/go-fundamentals-best-practices/internal/parser"
	"github.com/jonesrussell/go-fundamentals-best-practices/internal/storage"
	"github.com/jonesrussell/go-fundamentals-best-practices/pkg/models"
)

// Handlers contains all HTTP handlers
type Handlers struct {
	parser   *parser.TutorialParser
	executor *executor.CodeExecutor
	storage  *storage.ProgressStorage
	tutorials []*models.Tutorial
}

// NewHandlers creates a new handlers instance
func NewHandlers(parser *parser.TutorialParser, executor *executor.CodeExecutor, storage *storage.ProgressStorage) (*Handlers, error) {
	// Load all tutorials
	tutorials, err := parser.LoadAllTutorials()
	if err != nil {
		return nil, fmt.Errorf("failed to load tutorials: %w", err)
	}

	return &Handlers{
		parser:   parser,
		executor: executor,
		storage:  storage,
		tutorials: tutorials,
	}, nil
}

// ListTutorials returns all tutorials with metadata
func (h *Handlers) ListTutorials(w http.ResponseWriter, r *http.Request) {
	var metadata []models.TutorialMetadata
	
	for _, tutorial := range h.tutorials {
		metadata = append(metadata, models.TutorialMetadata{
			ID:            tutorial.ID,
			Title:         tutorial.Title,
			Duration:      tutorial.Duration,
			Difficulty:    tutorial.Difficulty,
			Prerequisites: tutorial.Prerequisites,
			Level:         tutorial.Level,
			SectionCount:  len(tutorial.Sections),
		})
	}
	
	respondJSON(w, http.StatusOK, metadata)
}

// GetTutorial returns a full tutorial by ID
func (h *Handlers) GetTutorial(w http.ResponseWriter, r *http.Request) {
	tutorialID := r.URL.Query().Get("id")
	if tutorialID == "" {
		http.Error(w, "tutorial ID required", http.StatusBadRequest)
		return
	}
	
	for _, tutorial := range h.tutorials {
		if tutorial.ID == tutorialID {
			respondJSON(w, http.StatusOK, tutorial)
			return
		}
	}
	
	http.Error(w, "tutorial not found", http.StatusNotFound)
}

// GetTutorialSections returns sections for a tutorial
func (h *Handlers) GetTutorialSections(w http.ResponseWriter, r *http.Request) {
	tutorialID := r.URL.Query().Get("id")
	if tutorialID == "" {
		http.Error(w, "tutorial ID required", http.StatusBadRequest)
		return
	}
	
	for _, tutorial := range h.tutorials {
		if tutorial.ID == tutorialID {
			respondJSON(w, http.StatusOK, tutorial.Sections)
			return
		}
	}
	
	http.Error(w, "tutorial not found", http.StatusNotFound)
}

// ExecuteCode executes Go code and returns the result
func (h *Handlers) ExecuteCode(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	var req struct {
		Code string `json:"code"`
	}
	
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	
	if req.Code == "" {
		http.Error(w, "code is required", http.StatusBadRequest)
		return
	}
	
	ctx, cancel := context.WithTimeout(r.Context(), 15*time.Second)
	defer cancel()
	
	result, err := h.executor.Execute(ctx, req.Code)
	if err != nil {
		http.Error(w, fmt.Sprintf("execution error: %v", err), http.StatusInternalServerError)
		return
	}
	
	respondJSON(w, http.StatusOK, result)
}

// GetProgress returns user progress
func (h *Handlers) GetProgress(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("userId")
	if userID == "" {
		userID = "default" // Default user for now
	}
	
	progress := h.storage.GetProgress(userID)
	respondJSON(w, http.StatusOK, progress)
}

// UpdateProgress updates user progress
func (h *Handlers) UpdateProgress(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	var progress models.Progress
	if err := json.NewDecoder(r.Body).Decode(&progress); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	
	userID := r.URL.Query().Get("userId")
	if userID == "" {
		userID = "default"
	}
	
	if err := h.storage.UpdateProgress(userID, &progress); err != nil {
		http.Error(w, fmt.Sprintf("failed to update progress: %v", err), http.StatusInternalServerError)
		return
	}
	
	respondJSON(w, http.StatusOK, map[string]string{"status": "success"})
}

// MarkSectionComplete marks a section as completed
func (h *Handlers) MarkSectionComplete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	var req struct {
		TutorialID string `json:"tutorialId"`
		SectionID  string `json:"sectionId"`
	}
	
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	
	userID := r.URL.Query().Get("userId")
	if userID == "" {
		userID = "default"
	}
	
	if err := h.storage.MarkSectionComplete(userID, req.TutorialID, req.SectionID); err != nil {
		http.Error(w, fmt.Sprintf("failed to mark section complete: %v", err), http.StatusInternalServerError)
		return
	}
	
	respondJSON(w, http.StatusOK, map[string]string{"status": "success"})
}

// GetExercises returns exercises for a tutorial
func (h *Handlers) GetExercises(w http.ResponseWriter, r *http.Request) {
	tutorialID := r.URL.Query().Get("tutorialId")
	if tutorialID == "" {
		http.Error(w, "tutorial ID required", http.StatusBadRequest)
		return
	}
	
	// For now, return empty exercises
	// This can be enhanced to parse exercises from tutorial files
	respondJSON(w, http.StatusOK, []models.Exercise{})
}

// respondJSON sends a JSON response
func respondJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

