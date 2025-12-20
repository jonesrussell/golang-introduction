package parser

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/jonesrussell/go-fundamentals-best-practices/pkg/models"
)

// ParseTutorial parses a complete tutorial markdown file
func (p *TutorialParser) ParseTutorial(filename string) (*models.Tutorial, error) {
	filePath := filepath.Join(p.tutorialsDir, filename)
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read tutorial file: %w", err)
	}

	contentStr := string(content)
	tutorial := &models.Tutorial{
		ID:            ExtractTutorialID(filename),
		Level:         DetermineLevel(ExtractTutorialID(filename)),
		Prerequisites: []string{},
		Sections:      []models.Section{},
	}

	// Parse metadata
	metadata, err := p.ParseTutorialMetadata(filename, contentStr)
	if err == nil {
		tutorial.Title = metadata.Title
		tutorial.Duration = metadata.Duration
		tutorial.Difficulty = metadata.Difficulty
		tutorial.Prerequisites = metadata.Prerequisites
	}

	// Parse sections
	sections := p.parseSections(contentStr)
	tutorial.Sections = sections

	return tutorial, nil
}

// listLegacyTutorials returns only legacy .md tutorial files (not directories)
func listLegacyTutorials(tutorialsDir string) ([]string, error) {
	files, err := os.ReadDir(tutorialsDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read tutorials directory: %w", err)
	}

	var tutorials []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".md") && strings.HasPrefix(file.Name(), "Tutorial-") {
			tutorials = append(tutorials, file.Name())
		}
	}

	return tutorials, nil
}

// ListTutorials returns all tutorial IDs (both directory-based and file-based)
func (p *TutorialParser) ListTutorials() ([]string, error) {
	tutorialIDs := make(map[string]bool)

	// Get directory-based tutorials
	dirTutorials, err := p.directoryParser.ListTutorialDirectories()
	if err == nil {
		for _, id := range dirTutorials {
			tutorialIDs[id] = true
		}
	}

	// Get file-based tutorials
	files, err := listLegacyTutorials(p.tutorialsDir)
	if err != nil {
		return nil, err
	}

	for _, filename := range files {
		id := ExtractTutorialID(filename)
		// Don't include if directory version exists
		if !tutorialIDs[id] {
			tutorialIDs[id] = true
		}
	}

	// Convert to slice
	var result []string
	for id := range tutorialIDs {
		result = append(result, id)
	}

	// Sort by ID
	sortTutorialIDs(result)

	return result, nil
}

// sortTutorialIDs sorts tutorial IDs numerically
func sortTutorialIDs(ids []string) {
	sort.Slice(ids, func(i, j int) bool {
		var numI, numJ int
		_, _ = fmt.Sscanf(ids[i], "%d", &numI)
		_, _ = fmt.Sscanf(ids[j], "%d", &numJ)
		return numI < numJ
	})
}

// LoadAllTutorials loads and parses all tutorials (both formats)
func (p *TutorialParser) LoadAllTutorials() ([]*models.Tutorial, error) {
	tutorialIDs, err := p.ListTutorials()
	if err != nil {
		return nil, err
	}

	var tutorials []*models.Tutorial
	for _, id := range tutorialIDs {
		tutorial, loadErr := p.GetTutorial(id, false)
		if loadErr != nil {
			// Log error but continue
			fmt.Fprintf(os.Stderr, "Error loading tutorial %s: %v\n", id, loadErr)
			continue
		}
		tutorials = append(tutorials, tutorial)
	}

	return tutorials, nil
}

// GetTutorialMetadataFromFile gets metadata from a legacy file
func (p *TutorialParser) GetTutorialMetadataFromFile(filename string) (*models.TutorialMetadata, error) {
	filePath := filepath.Join(p.tutorialsDir, filename)
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read tutorial file: %w", err)
	}

	metadata, err := p.ParseTutorialMetadata(filename, string(content))
	if err != nil {
		return nil, err
	}

	// Count sections
	sections := p.parseSections(string(content))
	metadata.SectionCount = len(sections)

	return metadata, nil
}
