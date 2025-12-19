package executor

import (
	"regexp"
	"strings"
)

// commonImports maps package prefixes to their import paths
var commonImports = map[string]string{
	"fmt.":      "fmt",
	"strings.":  "strings",
	"strconv.":  "strconv",
	"math.":     "math",
	"sort.":     "sort",
	"time.":     "time",
	"errors.":   "errors",
	"bytes.":    "bytes",
	"bufio.":    "bufio",
	"io.":       "io",
	"log.":      "log",
	"rand.":     "math/rand",
	"json.":     "encoding/json",
	"regexp.":   "regexp",
	"unicode.":  "unicode",
	"reflect.":  "reflect",
	"sync.":     "sync",
	"atomic.":   "sync/atomic",
	"context.":  "context",
	"filepath.": "path/filepath",
	"path.":     "path",
	"base64.":   "encoding/base64",
	"hex.":      "encoding/hex",
	"binary.":   "encoding/binary",
	"csv.":      "encoding/csv",
	"xml.":      "encoding/xml",
	"gob.":      "encoding/gob",
	"hash.":     "hash",
	"md5.":      "crypto/md5",
	"sha256.":   "crypto/sha256",
}

// IsSnippet checks if the code is a snippet (lacks package declaration)
func IsSnippet(code string) bool {
	code = strings.TrimSpace(code)
	// Check if code starts with package declaration
	packageRegex := regexp.MustCompile(`(?m)^package\s+\w+`)
	return !packageRegex.MatchString(code)
}

// HasMainFunc checks if the code has a main function
func HasMainFunc(code string) bool {
	mainRegex := regexp.MustCompile(`(?m)^func\s+main\s*\(\s*\)`)
	return mainRegex.MatchString(code)
}

// DetectImports scans the code for package usage and returns required imports
func DetectImports(code string) []string {
	imports := make(map[string]bool)

	for prefix, importPath := range commonImports {
		if strings.Contains(code, prefix) {
			imports[importPath] = true
		}
	}

	// Convert map to slice
	result := make([]string, 0, len(imports))
	for imp := range imports {
		result = append(result, imp)
	}

	return result
}

// WrapSnippet wraps a code snippet in a complete Go program
// It auto-detects required imports and adds package main + func main()
func WrapSnippet(code string) string {
	code = strings.TrimSpace(code)

	// If it already has a package declaration, return as-is
	if !IsSnippet(code) {
		return code
	}

	// Detect required imports
	imports := DetectImports(code)

	// Build the wrapped code
	var builder strings.Builder

	builder.WriteString("package main\n\n")

	// Add imports if any
	if len(imports) > 0 {
		if len(imports) == 1 {
			builder.WriteString("import \"")
			builder.WriteString(imports[0])
			builder.WriteString("\"\n\n")
		} else {
			builder.WriteString("import (\n")
			for _, imp := range imports {
				builder.WriteString("\t\"")
				builder.WriteString(imp)
				builder.WriteString("\"\n")
			}
			builder.WriteString(")\n\n")
		}
	}

	// Check if the code already has a main function (but no package)
	if HasMainFunc(code) {
		builder.WriteString(code)
	} else {
		// Wrap in main function
		builder.WriteString("func main() {\n")

		// Indent the code
		lines := strings.Split(code, "\n")
		for _, line := range lines {
			if strings.TrimSpace(line) != "" {
				builder.WriteString("\t")
				builder.WriteString(line)
				builder.WriteString("\n")
			} else {
				builder.WriteString("\n")
			}
		}

		builder.WriteString("}\n")
	}

	return builder.String()
}

// NeedsWrapping returns true if the code needs to be wrapped before execution
func NeedsWrapping(code string) bool {
	return IsSnippet(code)
}

// PrepareForExecution prepares code for execution, wrapping if necessary
func PrepareForExecution(code string, isSnippet bool) string {
	// If explicitly marked as snippet, always wrap
	if isSnippet {
		return WrapSnippet(code)
	}

	// Otherwise, check if it needs wrapping
	if NeedsWrapping(code) {
		return WrapSnippet(code)
	}

	return code
}
