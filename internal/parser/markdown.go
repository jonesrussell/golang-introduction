package parser

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/jonesrussell/go-fundamentals-best-practices/pkg/models"
)

// splitLimit is used when splitting strings with a maximum number of parts.
const splitLimit = 2

// sectionHeaderPrefix is the markdown prefix for section headers.
const sectionHeaderPrefix = "###"

// isSectionHeader checks if a line is a valid section header.
func isSectionHeader(trimmed string) bool {
	return strings.HasPrefix(trimmed, sectionHeaderPrefix) && !strings.Contains(trimmed, "Video")
}

// extractTitleFromSectionHeader extracts and cleans the title from a section header line.
func extractTitleFromSectionHeader(trimmed string) string {
	title := strings.TrimPrefix(trimmed, sectionHeaderPrefix)
	title = strings.TrimSpace(title)
	// Remove markdown bold formatting
	title = strings.TrimPrefix(title, "**")
	title = strings.TrimSuffix(title, "**")
	// Remove number prefix if present
	title = regexp.MustCompile(`^\d+\.\s*`).ReplaceAllString(title, "")
	return title
}

// createNewSection creates a new section with the given order and title.
func createNewSection(order int, title string) models.Section {
	return models.Section{
		ID:             fmt.Sprintf("section-%d", order),
		Title:          title,
		Order:          order,
		Topics:         []string{},
		CodeExamples:   []models.CodeExample{},
		TeachingPoints: []string{},
		Content:        "",
	}
}

// parseListItems extracts list items from lines starting at the given index.
// It stops when encountering an empty line or a line starting with specific prefixes.
func parseListItems(lines []string, startIndex int, stopPrefixes []string) []string {
	var items []string
	for j := startIndex; j < len(lines); j++ {
		itemLine := strings.TrimSpace(lines[j])
		if itemLine == "" || shouldStopParsing(itemLine, stopPrefixes) {
			break
		}
		if strings.HasPrefix(itemLine, "-") {
			item := strings.TrimPrefix(itemLine, "-")
			item = strings.TrimSpace(item)
			item = strings.Trim(item, "`")
			if item != "" {
				items = append(items, item)
			}
		}
	}
	return items
}

// shouldStopParsing checks if parsing should stop based on line prefixes.
func shouldStopParsing(line string, stopPrefixes []string) bool {
	for _, prefix := range stopPrefixes {
		if strings.HasPrefix(line, prefix) {
			return true
		}
	}
	return false
}

// isTopicsLine checks if the line indicates a topics section.
func isTopicsLine(trimmed string) bool {
	return strings.Contains(trimmed, "Topics to cover:") || strings.Contains(trimmed, "Topics:")
}

// isTeachingPointsLine checks if the line indicates a teaching points section.
func isTeachingPointsLine(trimmed string) bool {
	return strings.Contains(trimmed, "Key teaching points:") || strings.Contains(trimmed, "Teaching points:")
}

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

// parseSections extracts all sections from markdown content
func (p *TutorialParser) parseSections(content string) []models.Section {
	var sections []models.Section

	lines := strings.Split(content, "\n")
	currentSection := models.Section{Order: 0}
	sectionOrder := 0
	inSection := false

	topicStopPrefixes := []string{"**", sectionHeaderPrefix}
	teachingStopPrefixes := []string{"**", sectionHeaderPrefix, "---"}

	for i, line := range lines {
		trimmed := strings.TrimSpace(line)

		if isSectionHeader(trimmed) {
			sections, currentSection, sectionOrder, inSection = p.handleSectionHeader(
				sections, currentSection, sectionOrder, inSection, trimmed,
			)
			continue
		}

		if !inSection {
			continue
		}

		currentSection.Content += line + "\n"
		p.parseSectionContent(&currentSection, lines, i, trimmed, topicStopPrefixes, teachingStopPrefixes)
	}

	// Add last section
	if inSection && currentSection.Title != "" {
		sections = append(sections, currentSection)
	}

	return sections
}

// handleSectionHeader processes a section header line and returns updated state.
func (p *TutorialParser) handleSectionHeader(
	sections []models.Section,
	currentSection models.Section,
	sectionOrder int,
	inSection bool,
	trimmed string,
) (updatedSections []models.Section, newSection models.Section, newOrder int, nowInSection bool) {
	// Save previous section if it exists
	if inSection && currentSection.Title != "" {
		sections = append(sections, currentSection)
	}

	// Start new section
	sectionOrder++
	title := extractTitleFromSectionHeader(trimmed)
	newSection = createNewSection(sectionOrder, title)

	return sections, newSection, sectionOrder, true
}

// parseSectionContent parses content within a section (topics, code, teaching points).
func (p *TutorialParser) parseSectionContent(
	section *models.Section,
	lines []string,
	lineIndex int,
	trimmed string,
	topicStopPrefixes, teachingStopPrefixes []string,
) {
	if isTopicsLine(trimmed) {
		topics := parseListItems(lines, lineIndex+1, topicStopPrefixes)
		section.Topics = append(section.Topics, topics...)
	}

	if strings.HasPrefix(trimmed, "```") {
		if codeExample := p.parseCodeBlock(lines, lineIndex); codeExample != nil {
			section.CodeExamples = append(section.CodeExamples, *codeExample)
		}
	}

	if isTeachingPointsLine(trimmed) {
		points := parseListItems(lines, lineIndex+1, teachingStopPrefixes)
		section.TeachingPoints = append(section.TeachingPoints, points...)
	}
}

// parseCodeBlock extracts a code block from markdown
func (p *TutorialParser) parseCodeBlock(lines []string, startIndex int) *models.CodeExample {
	if startIndex >= len(lines) {
		return nil
	}

	firstLine := strings.TrimSpace(lines[startIndex])
	if !strings.HasPrefix(firstLine, "```") {
		return nil
	}

	// Extract language
	language := strings.TrimPrefix(firstLine, "```")
	language = strings.TrimSpace(language)
	if language == "" {
		language = "go" // Default to Go
	}

	var codeLines []string

	// Find the closing ```
	for i := startIndex + 1; i < len(lines); i++ {
		if strings.HasPrefix(strings.TrimSpace(lines[i]), "```") {
			break
		}
		codeLines = append(codeLines, lines[i])
	}

	if len(codeLines) == 0 {
		return nil
	}

	code := strings.Join(codeLines, "\n")
	code = strings.TrimSpace(code)

	// Determine if runnable (Go code blocks are typically runnable)
	runnable := language == "go" && strings.Contains(code, "package main")

	return &models.CodeExample{
		ID:       fmt.Sprintf("code-%d", startIndex),
		Code:     code,
		Language: language,
		Runnable: runnable,
	}
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

// isExerciseSectionStart checks if the line starts an exercise section.
func isExerciseSectionStart(trimmed string) bool {
	return strings.Contains(trimmed, "Practice Exercises:") ||
		strings.Contains(trimmed, "Homework/Practice suggestions:") ||
		strings.Contains(trimmed, "Practice suggestions:")
}

// extractDifficultyAndCleanText extracts difficulty level and cleans the exercise text.
func extractDifficultyAndCleanText(text string) (difficulty, cleanedText string) {
	difficulty = "Easy"

	if strings.Contains(text, "Challenge:") || strings.Contains(text, "Challenge") {
		difficulty = "Challenge"
		text = strings.ReplaceAll(text, "Challenge:", "")
		text = strings.ReplaceAll(text, "Challenge", "")
	} else if strings.Contains(text, "Medium:") || strings.Contains(text, "Medium") {
		difficulty = "Medium"
		text = strings.ReplaceAll(text, "Medium:", "")
		text = strings.ReplaceAll(text, "Medium", "")
	} else if strings.Contains(text, "Easy:") || strings.Contains(text, "Easy") {
		text = strings.ReplaceAll(text, "Easy:", "")
		text = strings.ReplaceAll(text, "Easy", "")
	}

	return difficulty, strings.TrimSpace(text)
}

// createExercise creates a new exercise with the given parameters.
func createExercise(id int, tutorialID, text, difficulty string) models.Exercise {
	return models.Exercise{
		ID:          fmt.Sprintf("exercise-%d", id),
		TutorialID:  tutorialID,
		Title:       text,
		Description: text,
		Difficulty:  difficulty,
	}
}

// isBulletPoint checks if the line is a bullet point.
func isBulletPoint(trimmed string) bool {
	return strings.HasPrefix(trimmed, "-") || strings.HasPrefix(trimmed, "*")
}

// parseBulletExercise parses an exercise from a bullet point line.
func parseBulletExercise(trimmed string) (difficulty, text string) {
	text = strings.TrimPrefix(trimmed, "-")
	text = strings.TrimPrefix(text, "*")
	text = strings.TrimSpace(text)
	return extractDifficultyAndCleanText(text)
}

// isNumberedLine checks if the line starts with a number followed by a period.
func isNumberedLine(trimmed string) bool {
	matched, _ := regexp.MatchString(`^\d+\.`, trimmed)
	return matched
}

// difficultyCleanupRegex is a compiled regex for cleaning difficulty markers.
var difficultyCleanupRegex = regexp.MustCompile(`(?i)(Easy|Medium|Challenge):?\s*`)

// parseNumberedExercise parses an exercise from a numbered line.
func parseNumberedExercise(trimmed string) (difficulty, text string) {
	parts := strings.SplitN(trimmed, ".", splitLimit)
	if len(parts) != splitLimit {
		return "", ""
	}

	text = strings.TrimSpace(parts[1])
	difficulty = "Easy"

	if strings.Contains(text, "Challenge") {
		difficulty = "Challenge"
	} else if strings.Contains(text, "Medium") {
		difficulty = "Medium"
	}

	text = difficultyCleanupRegex.ReplaceAllString(text, "")
	return difficulty, strings.TrimSpace(text)
}

// ParseExercises extracts exercises from tutorial content
func (p *TutorialParser) ParseExercises(tutorialID, content string) []models.Exercise {
	var exercises []models.Exercise
	lines := strings.Split(content, "\n")
	inExerciseSection := false
	exerciseID := 0

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		if isExerciseSectionStart(trimmed) {
			inExerciseSection = true
			continue
		}

		if inExerciseSection && strings.HasPrefix(trimmed, "##") {
			break
		}

		if !inExerciseSection {
			continue
		}

		if exercise, ok := p.tryParseExercise(trimmed, tutorialID, &exerciseID); ok {
			exercises = append(exercises, exercise)
		}
	}

	return exercises
}

// tryParseExercise attempts to parse an exercise from a line.
func (p *TutorialParser) tryParseExercise(trimmed, tutorialID string, exerciseID *int) (models.Exercise, bool) {
	var difficulty, text string

	if isBulletPoint(trimmed) {
		difficulty, text = parseBulletExercise(trimmed)
	} else if isNumberedLine(trimmed) {
		difficulty, text = parseNumberedExercise(trimmed)
	}

	if text == "" {
		return models.Exercise{}, false
	}

	*exerciseID++
	return createExercise(*exerciseID, tutorialID, text, difficulty), true
}

// ParseExercisesFromFile reads a tutorial file and extracts exercises
func (p *TutorialParser) ParseExercisesFromFile(tutorialID, filename string) []models.Exercise {
	filePath := filepath.Join(p.tutorialsDir, filename)
	content, err := os.ReadFile(filePath)
	if err != nil {
		return []models.Exercise{}
	}

	return p.ParseExercises(tutorialID, string(content))
}
