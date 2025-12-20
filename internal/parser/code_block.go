package parser

import (
	"fmt"
	"strings"

	"github.com/jonesrussell/go-fundamentals-best-practices/pkg/models"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

// parseCodeBlock extracts a code block from markdown using goldmark AST for robust parsing.
func (p *TutorialParser) parseCodeBlock(lines []string, startIndex int) *models.CodeExample {
	// Find the full code block range by looking for closing ```
	if startIndex >= len(lines) {
		return nil
	}

	firstLine := strings.TrimSpace(lines[startIndex])
	if !strings.HasPrefix(firstLine, "```") {
		return nil
	}

	// Extract the code block content as a string
	var codeBlockLines []string
	codeBlockLines = append(codeBlockLines, lines[startIndex])

	endIndex := startIndex + 1
	for endIndex < len(lines) {
		line := strings.TrimSpace(lines[endIndex])
		codeBlockLines = append(codeBlockLines, lines[endIndex])
		if strings.HasPrefix(line, "```") {
			break
		}
		endIndex++
	}

	codeBlockContent := strings.Join(codeBlockLines, "\n")

	// Use goldmark to parse just this code block for robust attribute extraction
	return p.extractCodeBlockWithGoldmark(codeBlockContent, startIndex)
}

// extractCodeBlockWithGoldmark uses goldmark AST to extract code block with proper attribute parsing.
func (p *TutorialParser) extractCodeBlockWithGoldmark(content string, idOffset int) *models.CodeExample {
	source := []byte(content)
	md := goldmark.New()
	doc := md.Parser().Parse(text.NewReader(source))

	var codeBlock *ast.FencedCodeBlock

	// Find the code block node in AST
	if err := ast.Walk(doc, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}
		if cb, ok := n.(*ast.FencedCodeBlock); ok {
			codeBlock = cb
			return ast.WalkStop, nil
		}
		return ast.WalkContinue, nil
	}); err != nil {
		return nil
	}

	if codeBlock == nil {
		return nil
	}

	// Extract language and attributes from info string
	info := string(codeBlock.Language(source))

	// Extract code using Lines() instead of deprecated Text()
	var codeBuilder strings.Builder
	lines := codeBlock.Lines()
	lineCount := lines.Len()
	for i := range lineCount {
		line := lines.At(i)
		codeBuilder.Write(line.Value(source))
	}
	code := codeBuilder.String()

	if code == "" {
		return nil
	}

	// Parse info string (e.g., "go", "go snippet", "go runnable")
	parts := strings.Fields(strings.TrimSpace(info))
	if len(parts) == 0 {
		return nil
	}

	language := parts[0]
	attribute := ""
	if len(parts) > 1 {
		attribute = parts[1] // "runnable" or "snippet"
	}

	runnable := false
	snippet := false

	switch attribute {
	case "runnable":
		runnable = true
	case "snippet":
		runnable = true
		snippet = true
	default:
		// Auto-detect: if Go code has "package main", it's runnable
		if language == "go" && strings.Contains(code, "package main") {
			runnable = true
		}
	}

	return &models.CodeExample{
		ID:       fmt.Sprintf("code-%d", idOffset),
		Code:     strings.TrimSpace(code),
		Language: language,
		Runnable: runnable,
		Snippet:  snippet,
	}
}
