import { defineStore } from 'pinia';
import { ref, computed } from 'vue';
import type { Progress, TutorialProgress } from '../types';
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
    loading.value = true;
    error.value = null;
    try {
      await progressApi.updateProgress(newProgress, userId);
      progress.value = newProgress;
      // Save to localStorage as backup
      localStorage.setItem('tutorial-progress', JSON.stringify(newProgress));
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to update progress';
    } finally {
      loading.value = false;
    }
  };

  const markSectionComplete = async (tutorialId: string, sectionId: string, userId: string = 'default') => {
    try {
      await progressApi.markSectionComplete(tutorialId, sectionId, userId);
      if (progress.value) {
        if (!progress.value.completedSections[tutorialId]) {
          progress.value.completedSections[tutorialId] = [];
        }
        if (!progress.value.completedSections[tutorialId].includes(sectionId)) {
          progress.value.completedSections[tutorialId].push(sectionId);
        }
        progress.value.currentTutorial = tutorialId;
        progress.value.currentSection = sectionId;
        localStorage.setItem('tutorial-progress', JSON.stringify(progress.value));
      }
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to mark section complete';
    }
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
    isSectionComplete,
    getTutorialProgress,
    loadFromLocalStorage,
  };
});

