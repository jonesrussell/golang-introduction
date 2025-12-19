package parser

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/jonesrussell/go-fundamentals-best-practices/pkg/models"
)

// Level threshold constants for determining tutorial difficulty.
const (
	beginnerMaxTutorial     = 3
	intermediateMinTutorial = 4
	intermediateMaxTutorial = 8
	advancedMinTutorial     = 9
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
	_, _ = fmt.Sscanf(tutorialNum, "%d", &num)

	if num >= 1 && num <= beginnerMaxTutorial {
		return "Beginner"
	} else if num >= intermediateMinTutorial && num <= intermediateMaxTutorial {
		return "Intermediate"
	} else if num >= advancedMinTutorial {
		return "Advanced"
	}
	return "Beginner"
}

// durationRegex matches time ranges like "25-35 minutes".
var durationRegex = regexp.MustCompile(`(\d+-\d+\s*(?:minutes?|min))`)

// extractTitleFromHeader extracts title from a header line.
func extractTitleFromHeader(line string) string {
	title := strings.TrimSpace(strings.TrimPrefix(line, "##"))
	title = strings.TrimPrefix(title, "**")
	title = strings.TrimSuffix(title, "**")
	return title
}

// extractInitialTitle extracts the title from the content before the metadata section.
func extractInitialTitle(lines []string) string {
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "##") && strings.Contains(line, "Video Tutorial Plan") {
			continue
		}
		if strings.HasPrefix(line, "### **Video Metadata**") {
			break
		}
		if line != "" && strings.HasPrefix(line, "##") {
			return extractTitleFromHeader(line)
		}
	}
	return ""
}

// parseMetadataTitle parses the title from a metadata line.
func parseMetadataTitle(line string) string {
	parts := strings.Split(line, "Title:")
	if len(parts) > 1 {
		return strings.TrimSpace(strings.Trim(parts[1], "*"))
	}
	return ""
}

// parseMetadataDuration parses the duration from a metadata line.
func parseMetadataDuration(line string) string {
	matches := durationRegex.FindStringSubmatch(line)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

// parseMetadataDifficulty parses the difficulty from a metadata line.
func parseMetadataDifficulty(line string) string {
	parts := strings.Split(line, "Difficulty:**")
	if len(parts) <= 1 {
		return ""
	}

	diff := strings.TrimSpace(parts[1])
	// Extract just the first word before any parenthesis
	if idx := strings.Index(diff, "("); idx > 0 {
		diff = strings.TrimSpace(diff[:idx])
	}
	// Take first word if multiple words
	words := strings.Fields(diff)
	if len(words) > 0 {
		return words[0]
	}
	return ""
}

// parseMetadataPrerequisites parses the prerequisites from a metadata line.
func parseMetadataPrerequisites(line string) []string {
	parts := strings.Split(line, "Prerequisites:**")
	if len(parts) > 1 {
		prereqs := strings.TrimSpace(parts[1])
		if prereqs != "" {
			return []string{prereqs}
		}
	}
	return nil
}

// parseMetadataSection parses the metadata section and updates the metadata struct.
func parseMetadataSection(lines []string, metadata *models.TutorialMetadata) {
	inMetadata := false

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.Contains(line, "Video Metadata") {
			inMetadata = true
			continue
		}

		if !inMetadata {
			continue
		}

		if strings.HasPrefix(line, "###") || strings.HasPrefix(line, "##") {
			break
		}

		updateMetadataFromLine(line, metadata)
	}
}

// updateMetadataFromLine updates metadata fields based on a single line.
func updateMetadataFromLine(line string, metadata *models.TutorialMetadata) {
	if strings.Contains(line, "Title:") {
		if title := parseMetadataTitle(line); title != "" {
			metadata.Title = title
		}
	}

	if strings.Contains(line, "Duration") {
		if duration := parseMetadataDuration(line); duration != "" {
			metadata.Duration = duration
		}
	}

	if strings.Contains(line, "Difficulty:**") {
		if difficulty := parseMetadataDifficulty(line); difficulty != "" {
			metadata.Difficulty = difficulty
		}
	}

	if strings.Contains(line, "Prerequisites:**") {
		if prereqs := parseMetadataPrerequisites(line); prereqs != nil {
			metadata.Prerequisites = prereqs
		}
	}
}

// ParseTutorialMetadata extracts basic metadata from a tutorial file
func (p *TutorialParser) ParseTutorialMetadata(filename, content string) (*models.TutorialMetadata, error) {
	id := ExtractTutorialID(filename)

	metadata := &models.TutorialMetadata{
		ID:            id,
		Level:         DetermineLevel(id),
		Prerequisites: []string{},
	}

	lines := strings.Split(content, "\n")

	// Extract initial title from content
	metadata.Title = extractInitialTitle(lines)

	// Parse metadata section
	parseMetadataSection(lines, metadata)

	// Fallback to filename if no title found
	if metadata.Title == "" {
		base := filepath.Base(filename)
		metadata.Title = strings.TrimSuffix(base, filepath.Ext(base))
	}

	return metadata, nil
}
