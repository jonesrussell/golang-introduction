export interface Progress {
  userId: string;
  completedSections: Record<string, string[]>;
  completedExercises: Record<string, string[]>;
  currentTutorial?: string;
  currentSection?: string;
  lastAccessed: string;
}

export interface SectionProgress {
  sectionId: string;
  completed: boolean;
  completedAt?: string;
}

export interface TutorialProgress {
  tutorialId: string;
  totalSections: number;
  completedCount: number;
  sectionProgress: SectionProgress[];
  progressPercent: number;
}

export interface ExecutionResult {
  output: string;
  error?: string;
  exitCode: number;
  duration: string;
}

