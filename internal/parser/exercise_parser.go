package parser

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/jonesrussell/go-fundamentals-best-practices/pkg/models"
)

// splitLimit is used when splitting strings with a maximum number of parts.
const splitLimit = 2

// isExerciseSectionStart checks if the line starts an exercise section.
func isExerciseSectionStart(trimmed string) bool {
	return strings.Contains(trimmed, "Practice Exercises:") ||
		strings.Contains(trimmed, "Homework/Practice suggestions:") ||
		strings.Contains(trimmed, "Practice suggestions:")
}

// extractDifficultyAndCleanText extracts difficulty level and cleans the exercise text.
func extractDifficultyAndCleanText(exerciseText string) (difficulty, cleanedText string) {
	difficulty = "Easy"

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

	return difficulty, strings.TrimSpace(exerciseText)
}

// createExercise creates a new exercise with the given parameters.
func createExercise(id int, tutorialID, exerciseText, difficulty string) models.Exercise {
	return models.Exercise{
		ID:          fmt.Sprintf("exercise-%d", id),
		TutorialID:  tutorialID,
		Title:       exerciseText,
		Description: exerciseText,
		Difficulty:  difficulty,
	}
}

// isBulletPoint checks if the line is a bullet point.
func isBulletPoint(trimmed string) bool {
	return strings.HasPrefix(trimmed, "-") || strings.HasPrefix(trimmed, "*")
}

// parseBulletExercise parses an exercise from a bullet point line.
func parseBulletExercise(trimmed string) (difficulty, exerciseText string) {
	exerciseText = strings.TrimPrefix(trimmed, "-")
	exerciseText = strings.TrimPrefix(exerciseText, "*")
	exerciseText = strings.TrimSpace(exerciseText)
	return extractDifficultyAndCleanText(exerciseText)
}

// isNumberedLine checks if the line starts with a number followed by a period.
func isNumberedLine(trimmed string) bool {
	matched, _ := regexp.MatchString(`^\d+\.`, trimmed)
	return matched
}

// difficultyCleanupRegex is a compiled regex for cleaning difficulty markers.
var difficultyCleanupRegex = regexp.MustCompile(`(?i)(Easy|Medium|Challenge):?\s*`)

// parseNumberedExercise parses an exercise from a numbered line.
func parseNumberedExercise(trimmed string) (difficulty, exerciseText string) {
	parts := strings.SplitN(trimmed, ".", splitLimit)
	if len(parts) != splitLimit {
		return "", ""
	}

	exerciseText = strings.TrimSpace(parts[1])
	difficulty = "Easy"

	if strings.Contains(exerciseText, "Challenge") {
		difficulty = "Challenge"
	} else if strings.Contains(exerciseText, "Medium") {
		difficulty = "Medium"
	}

	exerciseText = difficultyCleanupRegex.ReplaceAllString(exerciseText, "")
	return difficulty, strings.TrimSpace(exerciseText)
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
	var difficulty, exerciseText string

	if isBulletPoint(trimmed) {
		difficulty, exerciseText = parseBulletExercise(trimmed)
	} else if isNumberedLine(trimmed) {
		difficulty, exerciseText = parseNumberedExercise(trimmed)
	}

	if exerciseText == "" {
		return models.Exercise{}, false
	}

	*exerciseID++
	return createExercise(*exerciseID, tutorialID, exerciseText, difficulty), true
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
