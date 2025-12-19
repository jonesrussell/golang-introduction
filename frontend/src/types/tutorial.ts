export interface Tutorial {
  id: string;
  title: string;
  duration: string;
  difficulty: string;
  prerequisites: string[];
  sections: Section[];
  level: string;
}

export interface TutorialMetadata {
  id: string;
  title: string;
  duration: string;
  difficulty: string;
  prerequisites: string[];
  level: string;
  sectionCount: number;
}

export interface Section {
  id: string;
  title: string;
  topics: string[];
  codeExamples: CodeExample[];
  teachingPoints: string[];
  order: number;
  content: string;
  instructorNotes?: string;
}

export interface CodeExample {
  id: string;
  code: string;
  language: string;
  runnable: boolean;
  snippet?: boolean;
  expectedOutput?: string;
  description?: string;
}

export interface Exercise {
  id: string;
  tutorialId: string;
  title: string;
  description: string;
  difficulty: string;
  hints?: string[];
  solution?: string;
  starterCode?: string;
  expectedOutput?: string;
}

