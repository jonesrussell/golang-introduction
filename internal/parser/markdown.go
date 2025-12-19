package parser

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
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

// parseSections extracts all sections from markdown content
func (p *TutorialParser) parseSections(content string) []models.Section {
	var sections []models.Section

	// Split by section headers (###)
	sectionRegex := regexp.MustCompile(`(?m)^###\s+\d+\.\s+(.+?)$`)
	matches := sectionRegex.FindAllStringSubmatch(content, -1)

	if len(matches) == 0 {
		// Try alternative format without numbers
		sectionRegex = regexp.MustCompile(`(?m)^###\s+(.+?)$`)
		matches = sectionRegex.FindAllStringSubmatch(content, -1)
	}

	lines := strings.Split(content, "\n")
	currentSection := models.Section{
		Order: 0,
	}
	sectionOrder := 0
	inSection := false

	for i, line := range lines {
		trimmed := strings.TrimSpace(line)

		// Check if this is a section header
		if strings.HasPrefix(trimmed, "###") && !strings.Contains(trimmed, "Video") {
			// Save previous section if it exists
			if inSection && currentSection.Title != "" {
				sections = append(sections, currentSection)
			}

			// Start new section
			sectionOrder++
			title := strings.TrimPrefix(trimmed, "###")
			title = strings.TrimSpace(title)
			// Remove markdown formatting
			title = strings.Trim(title, "**")
			// Remove number prefix if present
			title = regexp.MustCompile(`^\d+\.\s*`).ReplaceAllString(title, "")

			currentSection = models.Section{
				ID:             fmt.Sprintf("section-%d", sectionOrder),
				Title:          title,
				Order:          sectionOrder,
				Topics:         []string{},
				CodeExamples:   []models.CodeExample{},
				TeachingPoints: []string{},
				Content:        "",
			}
			inSection = true
			continue
		}

		if !inSection {
			continue
		}

		// Accumulate content
		currentSection.Content += line + "\n"

		// Parse topics
		if strings.Contains(trimmed, "Topics to cover:") || strings.Contains(trimmed, "Topics:") {
			// Collect topics from following lines
			for j := i + 1; j < len(lines); j++ {
				topicLine := strings.TrimSpace(lines[j])
				if topicLine == "" || strings.HasPrefix(topicLine, "**") || strings.HasPrefix(topicLine, "###") {
					break
				}
				if strings.HasPrefix(topicLine, "-") {
					topic := strings.TrimPrefix(topicLine, "-")
					topic = strings.TrimSpace(topic)
					topic = strings.Trim(topic, "`")
					if topic != "" {
						currentSection.Topics = append(currentSection.Topics, topic)
					}
				}
			}
		}

		// Parse code examples
		if strings.HasPrefix(trimmed, "```") {
			codeExample := p.parseCodeBlock(lines, i)
			if codeExample != nil {
				currentSection.CodeExamples = append(currentSection.CodeExamples, *codeExample)
			}
		}

		// Parse teaching points
		if strings.Contains(trimmed, "Key teaching points:") || strings.Contains(trimmed, "Teaching points:") {
			for j := i + 1; j < len(lines); j++ {
				pointLine := strings.TrimSpace(lines[j])
				if pointLine == "" || strings.HasPrefix(pointLine, "**") || strings.HasPrefix(pointLine, "###") || strings.HasPrefix(pointLine, "---") {
					break
				}
				if strings.HasPrefix(pointLine, "-") {
					point := strings.TrimPrefix(pointLine, "-")
					point = strings.TrimSpace(point)
					point = strings.Trim(point, "`")
					if point != "" {
						currentSection.TeachingPoints = append(currentSection.TeachingPoints, point)
					}
				}
			}
		}
	}

	// Add last section
	if inSection && currentSection.Title != "" {
		sections = append(sections, currentSection)
	}

	return sections
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

// ListTutorials returns all tutorial files in the tutorials directory
func (p *TutorialParser) ListTutorials() ([]string, error) {
	files, err := os.ReadDir(p.tutorialsDir)
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

// LoadAllTutorials loads and parses all tutorial files
func (p *TutorialParser) LoadAllTutorials() ([]*models.Tutorial, error) {
	tutorialFiles, err := p.ListTutorials()
	if err != nil {
		return nil, err
	}

	var tutorials []*models.Tutorial
	for _, filename := range tutorialFiles {
		tutorial, err := p.ParseTutorial(filename)
		if err != nil {
			// Log error but continue
			fmt.Fprintf(os.Stderr, "Error parsing %s: %v\n", filename, err)
			continue
		}
		tutorials = append(tutorials, tutorial)
	}

	return tutorials, nil
}

// ParseExercises extracts exercises from tutorial content
func (p *TutorialParser) ParseExercises(tutorialID string, content string) []models.Exercise {
	var exercises []models.Exercise

	lines := strings.Split(content, "\n")
	inExerciseSection := false
	exerciseID := 0

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		// Look for exercise/practice sections
		if strings.Contains(trimmed, "Practice Exercises:") ||
			strings.Contains(trimmed, "Homework/Practice suggestions:") ||
			strings.Contains(trimmed, "Practice suggestions:") {
			inExerciseSection = true
			continue
		}

		// Stop at next major section
		if inExerciseSection && strings.HasPrefix(trimmed, "##") {
			break
		}

		if !inExerciseSection {
			continue
		}

		// Parse exercise items (usually bullet points)
		if strings.HasPrefix(trimmed, "-") || strings.HasPrefix(trimmed, "*") {
			exerciseText := strings.TrimPrefix(trimmed, "-")
			exerciseText = strings.TrimPrefix(exerciseText, "*")
			exerciseText = strings.TrimSpace(exerciseText)

			// Check for difficulty markers
			difficulty := "Easy"
			if strings.Contains(exerciseText, "Challenge:") || strings.Contains(exerciseText, "Challenge") {
				difficulty = "Challenge"
				exerciseText = strings.ReplaceAll(exerciseText, "Challenge:", "")
				exerciseText = strings.ReplaceAll(exerciseText, "Challenge", "")
			} else if strings.Contains(exerciseText, "Medium:") || strings.Contains(exerciseText, "Medium") {
				difficulty = "Medium"
				exerciseText = strings.ReplaceAll(exerciseText, "Medium:", "")
				exerciseText = strings.ReplaceAll(exerciseText, "Medium", "")
			} else if strings.Contains(exerciseText, "Easy:") || strings.Contains(exerciseText, "Easy") {
				exerciseText = strings.ReplaceAll(exerciseText, "Easy:", "")
				exerciseText = strings.ReplaceAll(exerciseText, "Easy", "")
			}

			exerciseText = strings.TrimSpace(exerciseText)

			if exerciseText != "" {
				exerciseID++
				exercise := models.Exercise{
					ID:          fmt.Sprintf("exercise-%d", exerciseID),
					TutorialID:  tutorialID,
					Title:       exerciseText,
					Description: exerciseText,
					Difficulty:  difficulty,
				}
				exercises = append(exercises, exercise)
			}
		}

		// Look for numbered exercises (e.g., "1. Easy: ...")
		if matched, _ := regexp.MatchString(`^\d+\.`, trimmed); matched {
			parts := strings.SplitN(trimmed, ".", 2)
			if len(parts) == 2 {
				exerciseText := strings.TrimSpace(parts[1])

				difficulty := "Easy"
				if strings.Contains(exerciseText, "Challenge") {
					difficulty = "Challenge"
				} else if strings.Contains(exerciseText, "Medium") {
					difficulty = "Medium"
				}

				// Clean up difficulty markers
				exerciseText = regexp.MustCompile(`(?i)(Easy|Medium|Challenge):?\s*`).ReplaceAllString(exerciseText, "")
				exerciseText = strings.TrimSpace(exerciseText)

				if exerciseText != "" {
					exerciseID++
					exercise := models.Exercise{
						ID:          fmt.Sprintf("exercise-%d", exerciseID),
						TutorialID:  tutorialID,
						Title:       exerciseText,
						Description: exerciseText,
						Difficulty:  difficulty,
					}
					exercises = append(exercises, exercise)
				}
			}
		}
	}

	return exercises
}

// ParseExercisesFromFile reads a tutorial file and extracts exercises
func (p *TutorialParser) ParseExercisesFromFile(tutorialID string, filename string) []models.Exercise {
	filePath := filepath.Join(p.tutorialsDir, filename)
	content, err := os.ReadFile(filePath)
	if err != nil {
		return []models.Exercise{}
	}

	return p.ParseExercises(tutorialID, string(content))
}
