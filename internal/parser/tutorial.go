package parser

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/jonesrussell/go-fundamentals-best-practices/pkg/models"
)

// TutorialParser handles parsing of tutorial markdown files
type TutorialParser struct {
	tutorialsDir    string
	directoryParser *DirectoryParser
}

// NewTutorialParser creates a new tutorial parser
func NewTutorialParser(tutorialsDir string) *TutorialParser {
	return &TutorialParser{
		tutorialsDir:    tutorialsDir,
		directoryParser: NewDirectoryParser(tutorialsDir),
	}
}

// GetTutorial returns a tutorial by ID, checking directory format first, then file format
func (p *TutorialParser) GetTutorial(tutorialID string, includeInstructorNotes bool) (*models.Tutorial, error) {
	// Check if tutorial exists as directory
	if p.directoryParser.IsTutorialDirectory(tutorialID) {
		return p.directoryParser.ParseTutorialFromDirectory(tutorialID, includeInstructorNotes)
	}

	// Fall back to file-based tutorial
	filename := p.findTutorialFile(tutorialID)
	if filename == "" {
		return nil, fmt.Errorf("tutorial %s not found", tutorialID)
	}

	return p.ParseTutorial(filename)
}

// GetTutorialMetadata returns metadata for a tutorial
func (p *TutorialParser) GetTutorialMetadata(tutorialID string) (*models.TutorialMetadata, error) {
	// Check if tutorial exists as directory
	if p.directoryParser.IsTutorialDirectory(tutorialID) {
		return p.directoryParser.GetTutorialMetadataFromDirectory(tutorialID)
	}

	// Fall back to file-based tutorial
	filename := p.findTutorialFile(tutorialID)
	if filename == "" {
		return nil, fmt.Errorf("tutorial %s not found", tutorialID)
	}

	return p.GetTutorialMetadataFromFile(filename)
}

// findTutorialFile finds the file for a tutorial ID
func (p *TutorialParser) findTutorialFile(tutorialID string) string {
	tutorials, err := p.ListTutorialFiles()
	if err != nil {
		return ""
	}

	for _, filename := range tutorials {
		if ExtractTutorialID(filename) == tutorialID {
			return filename
		}
	}

	return ""
}

// ListTutorialFiles returns legacy .md tutorial files only
func (p *TutorialParser) ListTutorialFiles() ([]string, error) {
	return listLegacyTutorials(p.tutorialsDir)
}

// ExtractTutorialID extracts a tutorial ID from a filename
func ExtractTutorialID(filename string) string {
	// Extract number from filename like "Tutorial-1-Go-Basics-for-Beginners.md" -> "1"
	re := regexp.MustCompile(`Tutorial-(\d+)`)
	matches := re.FindStringSubmatch(filename)
	if len(matches) > 1 {
		return matches[1]
	}
	// Fallback: use filename without extension
	base := filepath.Base(filename)
	return strings.TrimSuffix(base, filepath.Ext(base))
}

// DetermineLevel determines the tutorial level based on tutorial number
func DetermineLevel(tutorialNum string) string {
	num := 0
	fmt.Sscanf(tutorialNum, "%d", &num)

	if num >= 1 && num <= 3 {
		return "Beginner"
	} else if num >= 4 && num <= 8 {
		return "Intermediate"
	} else if num >= 9 {
		return "Advanced"
	}
	return "Beginner"
}

// ParseTutorialMetadata extracts basic metadata from a tutorial file
func (p *TutorialParser) ParseTutorialMetadata(filename string, content string) (*models.TutorialMetadata, error) {
	id := ExtractTutorialID(filename)
	level := DetermineLevel(id)

	metadata := &models.TutorialMetadata{
		ID:            id,
		Level:         level,
		Prerequisites: []string{},
	}

	// Extract title from first line or header
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "##") && strings.Contains(line, "Video Tutorial Plan") {
			// Extract title from next non-empty line or from the line itself
			continue
		}
		if strings.HasPrefix(line, "### **Video Metadata**") {
			// Next section, start parsing metadata
			break
		}
		if line != "" && metadata.Title == "" {
			// Try to extract title
			if strings.HasPrefix(line, "##") {
				metadata.Title = strings.TrimSpace(strings.TrimPrefix(line, "##"))
				metadata.Title = strings.Trim(metadata.Title, "**")
			}
		}
	}

	// Parse metadata section
	inMetadata := false
	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.Contains(line, "Video Metadata") {
			inMetadata = true
			continue
		}

		if inMetadata {
			if strings.HasPrefix(line, "###") || strings.HasPrefix(line, "##") {
				break
			}

			// Parse title
			if strings.Contains(line, "Title:") {
				parts := strings.Split(line, "Title:")
				if len(parts) > 1 {
					metadata.Title = strings.TrimSpace(strings.Trim(parts[1], "*"))
				}
			}

			// Parse duration - extract just the time range like "25-35 minutes"
			// Line format: "- **Duration Target:** 25-35 minutes"
			if strings.Contains(line, "Duration") {
				re := regexp.MustCompile(`(\d+-\d+\s*(?:minutes?|min))`)
				matches := re.FindStringSubmatch(line)
				if len(matches) > 1 {
					metadata.Duration = matches[1]
				}
			}

			// Parse difficulty - extract just the level
			// Line format: "- **Difficulty:** Beginner (no prior Go experience needed)"
			if strings.Contains(line, "Difficulty:**") {
				parts := strings.Split(line, "Difficulty:**")
				if len(parts) > 1 {
					diff := strings.TrimSpace(parts[1])
					// Extract just the first word before any parenthesis
					if idx := strings.Index(diff, "("); idx > 0 {
						diff = strings.TrimSpace(diff[:idx])
					}
					// Take first word if multiple words
					words := strings.Fields(diff)
					if len(words) > 0 {
						metadata.Difficulty = words[0]
					}
				}
			}

			// Parse prerequisites
			// Line format: "- **Prerequisites:** Basic programming concepts helpful but not required"
			if strings.Contains(line, "Prerequisites:**") {
				parts := strings.Split(line, "Prerequisites:**")
				if len(parts) > 1 {
					prereqs := strings.TrimSpace(parts[1])
					if prereqs != "" {
						metadata.Prerequisites = []string{prereqs}
					}
				}
			}

			// Count sections for section count
			if strings.HasPrefix(line, "###") {
				metadata.SectionCount++
			}
		}
	}

	// If title still empty, use filename
	if metadata.Title == "" {
		base := filepath.Base(filename)
		metadata.Title = strings.TrimSuffix(base, filepath.Ext(base))
	}

	return metadata, nil
}
