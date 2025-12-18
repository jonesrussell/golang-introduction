package models

// Progress represents user progress through tutorials
type Progress struct {
	UserID              string                `json:"userId"`
	CompletedSections  map[string][]string   `json:"completedSections"`  // tutorialID -> []sectionID
	CompletedExercises  map[string][]string   `json:"completedExercises"` // tutorialID -> []exerciseID
	CurrentTutorial    string                `json:"currentTutorial,omitempty"`
	CurrentSection     string                `json:"currentSection,omitempty"`
	LastAccessed       string                `json:"lastAccessed"`
}

// SectionProgress represents progress for a specific section
type SectionProgress struct {
	SectionID    string `json:"sectionId"`
	Completed    bool   `json:"completed"`
	CompletedAt  string `json:"completedAt,omitempty"`
}

// TutorialProgress represents overall progress for a tutorial
type TutorialProgress struct {
	TutorialID      string            `json:"tutorialId"`
	TotalSections   int               `json:"totalSections"`
	CompletedCount  int               `json:"completedCount"`
	SectionProgress []SectionProgress `json:"sectionProgress"`
	ProgressPercent float64           `json:"progressPercent"`
}

