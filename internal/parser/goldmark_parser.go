package parser

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/jonesrussell/go-fundamentals-best-practices/pkg/models"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

// sectionHeadingLevel is the markdown heading level for sections (###).
const sectionHeadingLevel = 3

// GoldmarkParser uses goldmark for more robust markdown parsing
type GoldmarkParser struct {
	md goldmark.Markdown
}

// NewGoldmarkParser creates a new goldmark-based parser
func NewGoldmarkParser() *GoldmarkParser {
	md := goldmark.New(
		goldmark.WithExtensions(),
	)
	return &GoldmarkParser{md: md}
}

// ParseWithGoldmark parses markdown content using goldmark
func (gp *GoldmarkParser) ParseWithGoldmark(content string) (*models.Tutorial, error) {
	reader := text.NewReader([]byte(content))
	doc := gp.md.Parser().Parse(reader)

	tutorial := &models.Tutorial{
		Sections: []models.Section{},
	}

	// Walk the AST to extract content
	err := ast.Walk(doc, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		switch v := n.(type) {
		case *ast.Heading:
			// Extract heading text
			headingText := extractText(v)
			if v.Level == sectionHeadingLevel {
				// This is a section (###)
				section := models.Section{
					Title:          headingText,
					Topics:         []string{},
					CodeExamples:   []models.CodeExample{},
					TeachingPoints: []string{},
					Content:        "",
				}
				tutorial.Sections = append(tutorial.Sections, section)
			}
		case *ast.FencedCodeBlock:
			// Extract code blocks
			language := string(v.Language([]byte(content)))
			code := extractCodeBlock(v, []byte(content))

			runnable := language == "go" && strings.Contains(code, "package main")

			if len(tutorial.Sections) > 0 {
				lastSection := &tutorial.Sections[len(tutorial.Sections)-1]
				codeExample := models.CodeExample{
					ID:       fmt.Sprintf("code-%d", len(lastSection.CodeExamples)),
					Code:     code,
					Language: language,
					Runnable: runnable,
				}
				lastSection.CodeExamples = append(lastSection.CodeExamples, codeExample)
			}
		case *ast.List:
			// Extract lists (topics, teaching points, exercises)
			items := extractListItems(v)
			if len(tutorial.Sections) > 0 {
				lastSection := &tutorial.Sections[len(tutorial.Sections)-1]
				// Determine if this is topics, teaching points, or exercises based on context
				// This is a simplified version - in production, you'd track context better
				if strings.Contains(strings.ToLower(content), "topics") {
					lastSection.Topics = append(lastSection.Topics, items...)
				} else if strings.Contains(strings.ToLower(content), "teaching") {
					lastSection.TeachingPoints = append(lastSection.TeachingPoints, items...)
				}
			}
		}

		return ast.WalkContinue, nil
	})

	return tutorial, err
}

// extractText extracts text content from a node
func extractText(n ast.Node) string {
	var buf bytes.Buffer
	for child := n.FirstChild(); child != nil; child = child.NextSibling() {
		if textNode, ok := child.(*ast.Text); ok {
			buf.Write(textNode.Segment.Value(nil))
		}
	}
	return strings.TrimSpace(buf.String())
}

// extractCodeBlock extracts code from a fenced code block
func extractCodeBlock(n *ast.FencedCodeBlock, source []byte) string {
	var buf bytes.Buffer
	lines := n.Lines()
	for i := range lines.Len() {
		line := lines.At(i)
		buf.Write(line.Value(source))
	}
	return strings.TrimSpace(buf.String())
}

// extractListItems extracts items from a list
func extractListItems(n *ast.List) []string {
	var items []string
	for child := n.FirstChild(); child != nil; child = child.NextSibling() {
		if item, ok := child.(*ast.ListItem); ok {
			itemText := extractText(item)
			if itemText != "" {
				items = append(items, itemText)
			}
		}
	}
	return items
}
