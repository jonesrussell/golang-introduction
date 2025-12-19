import { ref, computed } from 'vue';
import type { Tutorial, TutorialMetadata } from '../types/tutorial';
import { tutorialApi } from '../services/api';
import { useTutorialCache } from './useTutorialCache';

export function useTutorial() {
  const tutorials = ref<TutorialMetadata[]>([]);
  const currentTutorial = ref<Tutorial | null>(null);
  const loading = ref(false);
  const error = ref<string | null>(null);
  const cache = useTutorialCache();

  const loadTutorials = async (forceRefresh = false) => {
    // Check cache first
    if (!forceRefresh) {
      const cached = cache.getCachedMetadata();
      if (cached.length > 0) {
        tutorials.value = cached;
        // Load fresh data in background
        loadTutorials(true).catch(() => {
          // Silently fail background refresh
        });
        return;
      }
    }

    loading.value = true;
    error.value = null;
    try {
      const data = await tutorialApi.listTutorials();
      tutorials.value = data;
      cache.setCachedMetadata(data);
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to load tutorials';
      // Try to use cached data if available
      const cached = cache.getCachedMetadata();
      if (cached.length > 0) {
        tutorials.value = cached;
      }
    } finally {
      loading.value = false;
    }
  };

  const loadTutorial = async (id: string, forceRefresh = false) => {
    // Check cache first
    if (!forceRefresh) {
      const cached = cache.getCachedTutorial(id);
      if (cached) {
        currentTutorial.value = cached;
        // Load fresh data in background
        loadTutorial(id, true).catch(() => {
          // Silently fail background refresh
        });
        return;
      }
    }

    loading.value = true;
    error.value = null;
    try {
      const tutorial = await tutorialApi.getTutorial(id);
      currentTutorial.value = tutorial;
      cache.setCachedTutorial(id, tutorial);
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to load tutorial';
      // Try to use cached data if available
      const cached = cache.getCachedTutorial(id);
      if (cached) {
        currentTutorial.value = cached;
      }
    } finally {
      loading.value = false;
    }
  };

  const getTutorialById = (id: string): TutorialMetadata | undefined => {
    return tutorials.value.find(t => t.id === id);
  };

  const tutorialsByLevel = computed(() => {
    const grouped: Record<string, TutorialMetadata[]> = {
      Beginner: [],
      Intermediate: [],
      Advanced: [],
    };

    tutorials.value.forEach((tutorial: TutorialMetadata) => {
      const level = tutorial.level || 'Beginner';
      if (grouped[level]) {
        grouped[level].push(tutorial);
      }
    });

    return grouped;
  });

  return {
    tutorials: computed(() => tutorials.value),
    currentTutorial: computed(() => currentTutorial.value),
    loading: computed(() => loading.value),
    error: computed(() => error.value),
    tutorialsByLevel,
    loadTutorials,
    loadTutorial,
    getTutorialById,
  };
}

