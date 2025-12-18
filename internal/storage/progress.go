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
		if err := storage.load(); err != nil {
			return nil, fmt.Errorf("failed to load progress: %w", err)
		}
	}
	
	return storage, nil
}

// GetProgress retrieves progress for a user
func (s *ProgressStorage) GetProgress(userID string) *models.Progress {
	s.mu.RLock()
	defer s.mu.RUnlock()
	
	if p, exists := s.progress[userID]; exists {
		return p
	}
	
	// Return empty progress
	return &models.Progress{
		UserID:             userID,
		CompletedSections:  make(map[string][]string),
		CompletedExercises: make(map[string][]string),
	}
}

// UpdateProgress updates progress for a user
func (s *ProgressStorage) UpdateProgress(userID string, progress *models.Progress) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	progress.UserID = userID
	s.progress[userID] = progress
	
	return s.save()
}

// MarkSectionComplete marks a section as completed
func (s *ProgressStorage) MarkSectionComplete(userID, tutorialID, sectionID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	if s.progress[userID] == nil {
		s.progress[userID] = &models.Progress{
			UserID:             userID,
			CompletedSections:  make(map[string][]string),
			CompletedExercises: make(map[string][]string),
		}
	}
	
	progress := s.progress[userID]
	if progress.CompletedSections == nil {
		progress.CompletedSections = make(map[string][]string)
	}
	
	// Check if already completed
	sections := progress.CompletedSections[tutorialID]
	for _, id := range sections {
		if id == sectionID {
			return nil // Already completed
		}
	}
	
	progress.CompletedSections[tutorialID] = append(sections, sectionID)
	progress.CurrentTutorial = tutorialID
	progress.CurrentSection = sectionID
	
	return s.save()
}

// MarkExerciseComplete marks an exercise as completed
func (s *ProgressStorage) MarkExerciseComplete(userID, tutorialID, exerciseID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	if s.progress[userID] == nil {
		s.progress[userID] = &models.Progress{
			UserID:             userID,
			CompletedSections:  make(map[string][]string),
			CompletedExercises: make(map[string][]string),
		}
	}
	
	progress := s.progress[userID]
	if progress.CompletedExercises == nil {
		progress.CompletedExercises = make(map[string][]string)
	}
	
	// Check if already completed
	exercises := progress.CompletedExercises[tutorialID]
	for _, id := range exercises {
		if id == exerciseID {
			return nil // Already completed
		}
	}
	
	progress.CompletedExercises[tutorialID] = append(exercises, exerciseID)
	
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
	if err := os.MkdirAll(filepath.Dir(s.filePath), 0755); err != nil {
		return err
	}
	
	return os.WriteFile(s.filePath, data, 0644)
}

