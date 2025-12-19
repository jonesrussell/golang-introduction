package models

// Tutorial represents a complete tutorial with all its sections
type Tutorial struct {
	ID            string    `json:"id"`
	Title         string    `json:"title"`
	Duration      string    `json:"duration"`
	Difficulty    string    `json:"difficulty"`
	Prerequisites []string  `json:"prerequisites"`
	Sections      []Section `json:"sections"`
	Level         string    `json:"level"` // Beginner, Intermediate, Advanced
}

// TutorialMetadata represents basic tutorial information without full content
type TutorialMetadata struct {
	ID            string   `json:"id"`
	Title         string   `json:"title"`
	Duration      string   `json:"duration"`
	Difficulty    string   `json:"difficulty"`
	Prerequisites []string `json:"prerequisites"`
	Level         string   `json:"level"`
	SectionCount  int      `json:"sectionCount"`
}

// Section represents a single section within a tutorial
type Section struct {
	ID              string        `json:"id"`
	Title           string        `json:"title"`
	Topics          []string      `json:"topics"`
	CodeExamples    []CodeExample `json:"codeExamples"`
	TeachingPoints  []string      `json:"teachingPoints"`
	Order           int           `json:"order"`
	Content         string        `json:"content"`                   // Markdown content for the section
	InstructorNotes string        `json:"instructorNotes,omitempty"` // Instructor-only notes (when instructor mode enabled)
}

// CodeExample represents a code example within a section
type CodeExample struct {
	ID             string `json:"id"`
	Code           string `json:"code"`
	Language       string `json:"language"`
	Runnable       bool   `json:"runnable"`
	Snippet        bool   `json:"snippet,omitempty"` // If true, code needs wrapping before execution
	ExpectedOutput string `json:"expectedOutput,omitempty"`
	Description    string `json:"description,omitempty"`
}

// Exercise represents a practice exercise
type Exercise struct {
	ID          string   `json:"id"`
	TutorialID  string   `json:"tutorialId"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Difficulty  string   `json:"difficulty"` // Easy, Medium, Challenge
	Hints       []string `json:"hints,omitempty"`
	Solution    string   `json:"solution,omitempty"`
	StarterCode string   `json:"starterCode,omitempty"`
}
