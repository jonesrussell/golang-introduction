package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/jonesrussell/go-fundamentals-best-practices/pkg/models"
)

// ProgressStorage handles storage of user progress
type ProgressStorage struct {
	mu       sync.RWMutex
	progress map[string]*models.Progress // userID -> progress
	filePath string
}

// NewProgressStorage creates a new progress storage
func NewProgressStorage(dataDir string) (*ProgressStorage, error) {
	filePath := filepath.Join(dataDir, "progress.json")

	storage := &ProgressStorage{
		progress: make(map[string]*models.Progress),
		filePath: filePath,
	}

	// Load existing progress if file exists
	if _, err := os.Stat(filePath); err == nil {
		if loadErr := storage.load(); loadErr != nil {
			return nil, fmt.Errorf("failed to load progress: %w", loadErr)
		}
	}

	return storage, nil
}

// newEmptyProgress creates a new empty progress for a user.
func newEmptyProgress(userID string) *models.Progress {
	return &models.Progress{
		UserID:             userID,
		CompletedSections:  make(map[string][]string),
		CompletedExercises: make(map[string][]string),
	}
}

// ensureUserProgress ensures a progress entry exists for the user.
// Must be called with the lock held.
func (s *ProgressStorage) ensureUserProgress(userID string) *models.Progress {
	if s.progress[userID] == nil {
		s.progress[userID] = newEmptyProgress(userID)
	}
	return s.progress[userID]
}

// GetProgress retrieves progress for a user
func (s *ProgressStorage) GetProgress(userID string) *models.Progress {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if p, exists := s.progress[userID]; exists {
		return p
	}

	// Return empty progress
	return newEmptyProgress(userID)
}

// UpdateProgress updates progress for a user
func (s *ProgressStorage) UpdateProgress(userID string, progress *models.Progress) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	progress.UserID = userID
	s.progress[userID] = progress

	return s.save()
}

// containsID checks if the given ID is already in the slice.
func containsID(ids []string, id string) bool {
	for _, existingID := range ids {
		if existingID == id {
			return true
		}
	}
	return false
}

// MarkSectionComplete marks a section as completed
func (s *ProgressStorage) MarkSectionComplete(userID, tutorialID, sectionID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	progress := s.ensureUserProgress(userID)
	if progress.CompletedSections == nil {
		progress.CompletedSections = make(map[string][]string)
	}

	// Check if already completed
	if containsID(progress.CompletedSections[tutorialID], sectionID) {
		return nil
	}

	progress.CompletedSections[tutorialID] = append(progress.CompletedSections[tutorialID], sectionID)
	progress.CurrentTutorial = tutorialID
	progress.CurrentSection = sectionID

	return s.save()
}

// MarkExerciseComplete marks an exercise as completed
func (s *ProgressStorage) MarkExerciseComplete(userID, tutorialID, exerciseID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	progress := s.ensureUserProgress(userID)
	if progress.CompletedExercises == nil {
		progress.CompletedExercises = make(map[string][]string)
	}

	// Check if already completed
	if containsID(progress.CompletedExercises[tutorialID], exerciseID) {
		return nil
	}

	progress.CompletedExercises[tutorialID] = append(progress.CompletedExercises[tutorialID], exerciseID)

	return s.save()
}

// load loads progress from file
func (s *ProgressStorage) load() error {
	data, err := os.ReadFile(s.filePath)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, &s.progress)
}

// save saves progress to file
func (s *ProgressStorage) save() error {
	data, err := json.MarshalIndent(s.progress, "", "  ")
	if err != nil {
		return err
	}

	// Ensure directory exists
	if mkdirErr := os.MkdirAll(filepath.Dir(s.filePath), 0755); mkdirErr != nil {
		return mkdirErr
	}

	return os.WriteFile(s.filePath, data, 0600)
}
