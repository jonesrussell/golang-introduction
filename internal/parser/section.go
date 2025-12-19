package parser

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/jonesrussell/go-fundamentals-best-practices/pkg/models"
	"gopkg.in/yaml.v3"
)

// TutorialConfig represents the tutorial.yaml configuration
type TutorialConfig struct {
	ID              string   `yaml:"id"`
	Title           string   `yaml:"title"`
	Duration        string   `yaml:"duration"`
	Difficulty      string   `yaml:"difficulty"`
	Level           string   `yaml:"level"`
	Prerequisites   []string `yaml:"prerequisites"`
	TableOfContents string   `yaml:"tableOfContents"`
}

// DirectoryParser handles parsing of directory-based tutorials
type DirectoryParser struct {
	tutorialsDir string
}

// NewDirectoryParser creates a new directory-based parser
func NewDirectoryParser(tutorialsDir string) *DirectoryParser {
	return &DirectoryParser{
		tutorialsDir: tutorialsDir,
	}
}

// IsTutorialDirectory checks if a path is a directory-based tutorial
func (p *DirectoryParser) IsTutorialDirectory(tutorialID string) bool {
	dirPath := filepath.Join(p.tutorialsDir, fmt.Sprintf("tutorial-%s", tutorialID))
	configPath := filepath.Join(dirPath, "tutorial.yaml")

	if _, err := os.Stat(configPath); err == nil {
		return true
	}
	return false
}

// ListTutorialDirectories returns all tutorial directories
func (p *DirectoryParser) ListTutorialDirectories() ([]string, error) {
	entries, err := os.ReadDir(p.tutorialsDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read tutorials directory: %w", err)
	}

	var tutorials []string
	for _, entry := range entries {
		if entry.IsDir() && strings.HasPrefix(entry.Name(), "tutorial-") {
			// Check if it has tutorial.yaml
			configPath := filepath.Join(p.tutorialsDir, entry.Name(), "tutorial.yaml")
			if _, statErr := os.Stat(configPath); statErr == nil {
				// Extract ID from directory name
				id := strings.TrimPrefix(entry.Name(), "tutorial-")
				tutorials = append(tutorials, id)
			}
		}
	}

	// Sort by ID
	sort.Slice(tutorials, func(i, j int) bool {
		return tutorials[i] < tutorials[j]
	})

	return tutorials, nil
}

// LoadTutorialConfig loads the tutorial.yaml configuration
func (p *DirectoryParser) LoadTutorialConfig(tutorialID string) (*TutorialConfig, error) {
	configPath := filepath.Join(p.tutorialsDir, fmt.Sprintf("tutorial-%s", tutorialID), "tutorial.yaml")

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read tutorial config: %w", err)
	}

	var config TutorialConfig
	if unmarshalErr := yaml.Unmarshal(data, &config); unmarshalErr != nil {
		return nil, fmt.Errorf("failed to parse tutorial config: %w", unmarshalErr)
	}

	// Ensure ID is set
	if config.ID == "" {
		config.ID = tutorialID
	}

	return &config, nil
}

// ListSectionFiles returns all section files in order
func (p *DirectoryParser) ListSectionFiles(tutorialID string) ([]string, error) {
	sectionsDir := filepath.Join(p.tutorialsDir, fmt.Sprintf("tutorial-%s", tutorialID), "sections")

	entries, err := os.ReadDir(sectionsDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read sections directory: %w", err)
	}

	var files []string
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".md") {
			files = append(files, entry.Name())
		}
	}

	// Sort alphabetically (01-, 02- prefix ensures order)
	sort.Strings(files)

	return files, nil
}

// ParseSectionFile parses a single section markdown file
func (p *DirectoryParser) ParseSectionFile(tutorialID, filename string, order int) (*models.Section, error) {
	filePath := filepath.Join(p.tutorialsDir, fmt.Sprintf("tutorial-%s", tutorialID), "sections", filename)

	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read section file: %w", err)
	}

	contentStr := string(content)

	section := &models.Section{
		ID:             fmt.Sprintf("section-%d", order),
		Order:          order,
		Topics:         []string{},
		CodeExamples:   []models.CodeExample{},
		TeachingPoints: []string{},
		Content:        contentStr,
	}

	// Extract title from first heading
	section.Title = extractSectionTitle(contentStr)

	// Parse topics
	section.Topics = extractTopics(contentStr)

	// Parse code examples with attributes
	section.CodeExamples = extractCodeExamples(contentStr)

	// Parse teaching points
	section.TeachingPoints = extractTeachingPoints(contentStr)

	return section, nil
}

// ParseTutorialFromDirectory parses a complete tutorial from directory structure
func (p *DirectoryParser) ParseTutorialFromDirectory(tutorialID string, includeInstructorNotes bool) (*models.Tutorial, error) {
	// Load config
	config, err := p.LoadTutorialConfig(tutorialID)
	if err != nil {
		return nil, err
	}

	tutorial := &models.Tutorial{
		ID:              config.ID,
		Title:           config.Title,
		Duration:        config.Duration,
		Difficulty:      config.Difficulty,
		Level:           config.Level,
		Prerequisites:   config.Prerequisites,
		TableOfContents: config.TableOfContents,
		Sections:        []models.Section{},
	}

	// If level is empty, determine from ID
	if tutorial.Level == "" {
		tutorial.Level = DetermineLevel(tutorialID)
	}

	// Load sections
	sectionFiles, err := p.ListSectionFiles(tutorialID)
	if err != nil {
		return nil, err
	}

	for i, filename := range sectionFiles {
		section, parseErr := p.ParseSectionFile(tutorialID, filename, i+1)
		if parseErr != nil {
			// Log error but continue
			fmt.Fprintf(os.Stderr, "Warning: failed to parse section %s: %v\n", filename, parseErr)
			continue
		}

		// Load instructor notes if requested
		if includeInstructorNotes {
			notes, _ := p.LoadInstructorNotes(tutorialID, filename)
			section.InstructorNotes = notes
		}

		tutorial.Sections = append(tutorial.Sections, *section)
	}

	return tutorial, nil
}

// LoadInstructorNotes loads instructor notes for a section
func (p *DirectoryParser) LoadInstructorNotes(tutorialID, sectionFilename string) (string, error) {
	notesPath := filepath.Join(p.tutorialsDir, fmt.Sprintf("tutorial-%s", tutorialID), "instructor", sectionFilename)

	content, err := os.ReadFile(notesPath)
	if err != nil {
		return "", err // File doesn't exist or can't be read
	}

	return string(content), nil
}

// GetTutorialMetadataFromDirectory gets metadata for a directory-based tutorial
func (p *DirectoryParser) GetTutorialMetadataFromDirectory(tutorialID string) (*models.TutorialMetadata, error) {
	config, err := p.LoadTutorialConfig(tutorialID)
	if err != nil {
		return nil, err
	}

	sectionFiles, err := p.ListSectionFiles(tutorialID)
	if err != nil {
		return nil, err
	}

	return &models.TutorialMetadata{
		ID:            config.ID,
		Title:         config.Title,
		Duration:      config.Duration,
		Difficulty:    config.Difficulty,
		Level:         config.Level,
		Prerequisites: config.Prerequisites,
		SectionCount:  len(sectionFiles),
	}, nil
}

// Helper functions

func extractSectionTitle(content string) string {
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "# ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "#"))
		}
		if strings.HasPrefix(line, "## ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "##"))
		}
	}
	return "Untitled Section"
}

func extractTopics(content string) []string {
	var topics []string
	lines := strings.Split(content, "\n")
	inTopics := false

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		// Check for topics section header
		if strings.Contains(strings.ToLower(trimmed), "topics to cover") ||
			strings.Contains(strings.ToLower(trimmed), "## topics") {
			inTopics = true
			continue
		}

		// Stop at next section
		if inTopics && (strings.HasPrefix(trimmed, "##") || strings.HasPrefix(trimmed, "```")) {
			break
		}

		// Extract topic items
		if inTopics && strings.HasPrefix(trimmed, "-") {
			topic := strings.TrimSpace(strings.TrimPrefix(trimmed, "-"))
			topic = strings.Trim(topic, "`")
			if topic != "" {
				topics = append(topics, topic)
			}
		}
	}

	return topics
}

// codeBlockMatchGroups is the expected number of capture groups in code block regex:
// full match, language, attribute (optional), code content
const codeBlockMatchGroups = 4

func extractCodeExamples(content string) []models.CodeExample {
	var examples []models.CodeExample

	// Regex to match code blocks with optional attributes
	// Matches: ```go, ```go runnable, ```go snippet
	codeBlockRegex := regexp.MustCompile("(?s)```(\\w+)(?:\\s+(runnable|snippet))?\\s*\\n(.*?)```")
	matches := codeBlockRegex.FindAllStringSubmatch(content, -1)

	for i, match := range matches {
		if len(match) < codeBlockMatchGroups {
			continue
		}

		language := match[1]
		attribute := match[2] // "runnable", "snippet", or ""
		code := strings.TrimSpace(match[3])

		// Determine if runnable
		runnable := false
		snippet := false

		if attribute == "runnable" {
			runnable = true
		} else if attribute == "snippet" {
			runnable = true
			snippet = true
		} else if language == "go" {
			// Auto-detect: if has package main, it's runnable
			if strings.Contains(code, "package main") {
				runnable = true
			}
		}

		example := models.CodeExample{
			ID:       fmt.Sprintf("code-%d", i),
			Code:     code,
			Language: language,
			Runnable: runnable,
			Snippet:  snippet,
		}
		examples = append(examples, example)
	}

	return examples
}

func extractTeachingPoints(content string) []string {
	var points []string
	lines := strings.Split(content, "\n")
	inPoints := false

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		// Check for teaching points section header
		if strings.Contains(strings.ToLower(trimmed), "key teaching points") ||
			strings.Contains(strings.ToLower(trimmed), "teaching points") ||
			strings.Contains(strings.ToLower(trimmed), "## key") {
			inPoints = true
			continue
		}

		// Stop at next section
		if inPoints && strings.HasPrefix(trimmed, "##") {
			break
		}

		// Extract teaching point items
		if inPoints && strings.HasPrefix(trimmed, "-") {
			point := strings.TrimSpace(strings.TrimPrefix(trimmed, "-"))
			point = strings.Trim(point, "`")
			if point != "" {
				points = append(points, point)
			}
		}
	}

	return points
}
