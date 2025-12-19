import { defineStore } from 'pinia';
import { ref, computed } from 'vue';
import type { Progress, TutorialProgress } from '../types/progress';
import { progressApi } from '../services/api';

export const useProgressStore = defineStore('progress', () => {
  const progress = ref<Progress | null>(null);
  const loading = ref(false);
  const error = ref<string | null>(null);

  const loadProgress = async (userId: string = 'default') => {
    loading.value = true;
    error.value = null;
    try {
      progress.value = await progressApi.getProgress(userId);
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to load progress';
    } finally {
      loading.value = false;
    }
  };

  const updateProgress = async (newProgress: Progress, userId: string = 'default') => {
    // Optimistic update
    const previousProgress = progress.value;
    progress.value = newProgress;
    localStorage.setItem('tutorial-progress', JSON.stringify(newProgress));
    
    loading.value = true;
    error.value = null;
    try {
      await progressApi.updateProgress(newProgress, userId);
    } catch (err) {
      // Revert on error
      progress.value = previousProgress;
      if (previousProgress) {
        localStorage.setItem('tutorial-progress', JSON.stringify(previousProgress));
      }
      error.value = err instanceof Error ? err.message : 'Failed to update progress';
    } finally {
      loading.value = false;
    }
  };

  const markSectionComplete = async (tutorialId: string, sectionId: string, userId: string = 'default') => {
    // Optimistic update - update UI immediately
    if (!progress.value) {
      progress.value = {
        userId: userId,
        completedSections: {},
        completedExercises: {},
        lastAccessed: new Date().toISOString(),
      };
    }
    
    // TypeScript now knows progress.value is not null after the check above
    const currentProgress = progress.value;
    
    if (!currentProgress.completedSections[tutorialId]) {
      currentProgress.completedSections[tutorialId] = [];
    }
    
    // Only add if not already completed (idempotent)
    if (!currentProgress.completedSections[tutorialId].includes(sectionId)) {
      currentProgress.completedSections[tutorialId].push(sectionId);
    }
    
    currentProgress.currentTutorial = tutorialId;
    currentProgress.currentSection = sectionId;
    currentProgress.lastAccessed = new Date().toISOString();
    
    // Save to localStorage immediately
    localStorage.setItem('tutorial-progress', JSON.stringify(currentProgress));
    
    // Then sync with server (non-blocking)
    try {
      await progressApi.markSectionComplete(tutorialId, sectionId, userId);
    } catch (err) {
      // Log error but don't revert optimistic update
      // In production, you might want to queue this for retry
      console.error('Failed to sync progress to server:', err);
      error.value = err instanceof Error ? err.message : 'Failed to sync progress';
    }
  };

  const setCurrentSection = (tutorialId: string, sectionId: string, userId: string = 'default') => {
    // Initialize progress if it doesn't exist
    if (!progress.value) {
      progress.value = {
        userId: userId,
        completedSections: {},
        completedExercises: {},
        lastAccessed: new Date().toISOString(),
      };
    }
    
    // Update current tutorial and section
    progress.value.currentTutorial = tutorialId;
    progress.value.currentSection = sectionId;
    progress.value.lastAccessed = new Date().toISOString();
    
    // Save to localStorage immediately
    localStorage.setItem('tutorial-progress', JSON.stringify(progress.value));
  };

  const isSectionComplete = (tutorialId: string, sectionId: string): boolean => {
    if (!progress.value) return false;
    const sections = progress.value.completedSections[tutorialId] || [];
    return sections.includes(sectionId);
  };

  const getTutorialProgress = (tutorialId: string, totalSections: number): TutorialProgress => {
    if (!progress.value) {
      return {
        tutorialId,
        totalSections,
        completedCount: 0,
        sectionProgress: [],
        progressPercent: 0,
      };
    }

    const completedSections = progress.value.completedSections[tutorialId] || [];
    const progressPercent = totalSections > 0 ? (completedSections.length / totalSections) * 100 : 0;

    return {
      tutorialId,
      totalSections,
      completedCount: completedSections.length,
      sectionProgress: [],
      progressPercent,
    };
  };

  // Load from localStorage on init
  const loadFromLocalStorage = () => {
    const stored = localStorage.getItem('tutorial-progress');
    if (stored) {
      try {
        progress.value = JSON.parse(stored);
      } catch (err) {
        console.error('Failed to parse stored progress', err);
      }
    }
  };

  return {
    progress: computed(() => progress.value),
    loading: computed(() => loading.value),
    error: computed(() => error.value),
    loadProgress,
    updateProgress,
    markSectionComplete,
    setCurrentSection,
    isSectionComplete,
    getTutorialProgress,
    loadFromLocalStorage,
  };
});

