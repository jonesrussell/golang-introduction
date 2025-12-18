import { ref, computed } from 'vue';
import type { Tutorial, TutorialMetadata } from '../types/tutorial';
import { tutorialApi } from '../services/api';

export function useTutorial() {
  const tutorials = ref<TutorialMetadata[]>([]);
  const currentTutorial = ref<Tutorial | null>(null);
  const loading = ref(false);
  const error = ref<string | null>(null);

  const loadTutorials = async () => {
    loading.value = true;
    error.value = null;
    try {
      tutorials.value = await tutorialApi.listTutorials();
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to load tutorials';
    } finally {
      loading.value = false;
    }
  };

  const loadTutorial = async (id: string) => {
    loading.value = true;
    error.value = null;
    try {
      currentTutorial.value = await tutorialApi.getTutorial(id);
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to load tutorial';
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

