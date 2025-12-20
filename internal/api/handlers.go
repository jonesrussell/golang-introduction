package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/jonesrussell/go-fundamentals-best-practices/internal/executor"
	"github.com/jonesrussell/go-fundamentals-best-practices/internal/parser"
	"github.com/jonesrussell/go-fundamentals-best-practices/internal/storage"
	"github.com/jonesrussell/go-fundamentals-best-practices/pkg/models"
)

// Handlers contains all HTTP handlers
type Handlers struct {
	parser    *parser.TutorialParser
	executor  *executor.CodeExecutor
	storage   *storage.ProgressStorage
	tutorials []*models.Tutorial
	logger    *slog.Logger
}

// NewHandlers creates a new handlers instance
func NewHandlers(
	tutorialParser *parser.TutorialParser,
	codeExecutor *executor.CodeExecutor,
	progressStorage *storage.ProgressStorage,
) (*Handlers, error) {
	// Load all tutorials
	tutorials, err := tutorialParser.LoadAllTutorials()
	if err != nil {
		return nil, fmt.Errorf("failed to load tutorials: %w", err)
	}

	return &Handlers{
		parser:    tutorialParser,
		executor:  codeExecutor,
		storage:   progressStorage,
		tutorials: tutorials,
		logger:    slog.Default(),
	}, nil
}

// extractUserID extracts the user ID from the request query params, falling back to default.
func extractUserID(r *http.Request) string {
	if userID := r.URL.Query().Get("userId"); userID != "" {
		return userID
	}
	return DefaultUserID
}

// findTutorial looks up a tutorial by ID in the cached tutorials.
func (h *Handlers) findTutorial(tutorialID string) *models.Tutorial {
	for _, t := range h.tutorials {
		if t.ID == tutorialID {
			return t
		}
	}
	return nil
}

// respondNotFound sends a 404 Not Found response.
func respondNotFound(w http.ResponseWriter) {
	http.Error(w, "tutorial not found", http.StatusNotFound)
}

// respondBadRequest sends a 400 Bad Request response.
func respondBadRequest(w http.ResponseWriter, message string) {
	http.Error(w, message, http.StatusBadRequest)
}

// respondMethodNotAllowed sends a 405 Method Not Allowed response.
func respondMethodNotAllowed(w http.ResponseWriter) {
	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}

// respondInternalError sends a 500 Internal Server Error response.
func respondInternalError(w http.ResponseWriter, message string) {
	http.Error(w, message, http.StatusInternalServerError)
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

	respondJSON(w, h.logger, metadata)
}

// GetTutorialByID returns a full tutorial by ID (path parameter version)
func (h *Handlers) GetTutorialByID(w http.ResponseWriter, r *http.Request, tutorialID string) {
	// Check if instructor mode is requested
	instructorMode := r.URL.Query().Get("instructor") == "true"

	// If instructor mode, load fresh with instructor notes
	if instructorMode {
		tutorial, err := h.parser.GetTutorial(tutorialID, true)
		if err != nil {
			respondNotFound(w)
			return
		}
		respondJSON(w, h.logger, tutorial)
		return
	}

	// Otherwise use cached tutorials
	if tutorial := h.findTutorial(tutorialID); tutorial != nil {
		respondJSON(w, h.logger, tutorial)
		return
	}

	respondNotFound(w)
}

// GetTutorial returns a full tutorial by ID (query parameter version - for backward compatibility)
func (h *Handlers) GetTutorial(w http.ResponseWriter, r *http.Request) {
	tutorialID := r.URL.Query().Get("id")
	if tutorialID == "" {
		respondBadRequest(w, "tutorial ID required")
		return
	}
	h.GetTutorialByID(w, r, tutorialID)
}

// GetTutorialSectionsByID returns sections for a tutorial (path parameter version)
func (h *Handlers) GetTutorialSectionsByID(w http.ResponseWriter, r *http.Request, tutorialID string) {
	if tutorial := h.findTutorial(tutorialID); tutorial != nil {
		respondJSON(w, h.logger, tutorial.Sections)
		return
	}

	respondNotFound(w)
}

// GetTutorialSections returns sections for a tutorial (query parameter version - for backward compatibility)
func (h *Handlers) GetTutorialSections(w http.ResponseWriter, r *http.Request) {
	tutorialID := r.URL.Query().Get("id")
	if tutorialID == "" {
		respondBadRequest(w, "tutorial ID required")
		return
	}
	h.GetTutorialSectionsByID(w, r, tutorialID)
}

// ExecuteCode executes Go code and returns the result
func (h *Handlers) ExecuteCode(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondMethodNotAllowed(w)
		return
	}

	var req struct {
		Code    string `json:"code"`
		Snippet bool   `json:"snippet,omitempty"` // If true, code will be auto-wrapped
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondBadRequest(w, "invalid request body")
		return
	}

	if req.Code == "" {
		respondBadRequest(w, "code is required")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), ExecuteTimeout)
	defer cancel()

	// Use appropriate execution method based on snippet flag
	var result *executor.ExecutionResult
	var err error
	if req.Snippet {
		result, err = h.executor.ExecuteSnippet(ctx, req.Code)
	} else {
		result, err = h.executor.Execute(ctx, req.Code)
	}

	if err != nil {
		respondInternalError(w, fmt.Sprintf("execution error: %v", err))
		return
	}

	respondJSON(w, h.logger, result)
}

// GetProgress returns user progress
func (h *Handlers) GetProgress(w http.ResponseWriter, r *http.Request) {
	userID := extractUserID(r)
	progress := h.storage.GetProgress(userID)
	respondJSON(w, h.logger, progress)
}

// UpdateProgress updates user progress
func (h *Handlers) UpdateProgress(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondMethodNotAllowed(w)
		return
	}

	var progress models.Progress
	if err := json.NewDecoder(r.Body).Decode(&progress); err != nil {
		respondBadRequest(w, "invalid request body")
		return
	}

	userID := extractUserID(r)

	if err := h.storage.UpdateProgress(userID, &progress); err != nil {
		respondInternalError(w, fmt.Sprintf("failed to update progress: %v", err))
		return
	}

	respondJSON(w, h.logger, map[string]string{"status": "success"})
}

// MarkSectionComplete marks a section as completed
func (h *Handlers) MarkSectionComplete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondMethodNotAllowed(w)
		return
	}

	var req struct {
		TutorialID string `json:"tutorialId"`
		SectionID  string `json:"sectionId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondBadRequest(w, "invalid request body")
		return
	}

	userID := extractUserID(r)

	if err := h.storage.MarkSectionComplete(userID, req.TutorialID, req.SectionID); err != nil {
		respondInternalError(w, fmt.Sprintf("failed to mark section complete: %v", err))
		return
	}

	respondJSON(w, h.logger, map[string]string{"status": "success"})
}

// GetExercisesByTutorialID returns exercises for a tutorial (path parameter version)
func (h *Handlers) GetExercisesByTutorialID(w http.ResponseWriter, r *http.Request, tutorialID string) {
	// Find the tutorial
	if h.findTutorial(tutorialID) == nil {
		respondNotFound(w)
		return
	}

	// Find the matching file and parse exercises from it
	tutorialFiles, err := h.parser.ListTutorials()
	if err != nil {
		respondJSON(w, h.logger, []models.Exercise{})
		return
	}

	var exercises []models.Exercise
	for _, f := range tutorialFiles {
		if parser.ExtractTutorialID(f) == tutorialID {
			exercises = h.parser.ParseExercisesFromFile(tutorialID, f)
			break
		}
	}
	respondJSON(w, h.logger, exercises)
}

// GetExercises returns exercises for a tutorial (query parameter version - for backward compatibility)
func (h *Handlers) GetExercises(w http.ResponseWriter, r *http.Request) {
	tutorialID := r.URL.Query().Get("tutorialId")
	if tutorialID == "" {
		respondBadRequest(w, "tutorial ID required")
		return
	}
	h.GetExercisesByTutorialID(w, r, tutorialID)
}

// respondJSON sends a JSON response with 200 OK status
func respondJSON(w http.ResponseWriter, logger *slog.Logger, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		logger.Error("failed to encode JSON response", "error", err)
	}
}
