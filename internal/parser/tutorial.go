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
	tutorialsDir string
}

// NewTutorialParser creates a new tutorial parser
func NewTutorialParser(tutorialsDir string) *TutorialParser {
	return &TutorialParser{
		tutorialsDir: tutorialsDir,
	}
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
			
			// Parse duration
			if strings.Contains(line, "Duration") {
				parts := strings.Split(line, "Duration")
				if len(parts) > 1 {
					metadata.Duration = strings.TrimSpace(strings.Trim(strings.Split(parts[1], "-")[0], "Target:"))
				}
			}
			
			// Parse difficulty
			if strings.Contains(line, "Difficulty:") {
				parts := strings.Split(line, "Difficulty:")
				if len(parts) > 1 {
					metadata.Difficulty = strings.TrimSpace(parts[1])
				}
			}
			
			// Parse prerequisites
			if strings.Contains(line, "Prerequisites:") {
				parts := strings.Split(line, "Prerequisites:")
				if len(parts) > 1 {
					prereqs := strings.TrimSpace(parts[1])
					if prereqs != "" {
						metadata.Prerequisites = strings.Split(prereqs, ",")
						for i, p := range metadata.Prerequisites {
							metadata.Prerequisites[i] = strings.TrimSpace(p)
						}
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

