package parser

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/jonesrussell/go-fundamentals-best-practices/pkg/models"
)

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
