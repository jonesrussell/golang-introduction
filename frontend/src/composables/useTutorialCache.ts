import { ref } from 'vue';
import type { Tutorial, TutorialMetadata } from '../types/tutorial';

// Simple in-memory cache for tutorials
const tutorialCache = new Map<string, Tutorial>();
const metadataCache = ref<TutorialMetadata[]>([]);
const cacheTimestamp = new Map<string, number>();
const CACHE_TTL = 5 * 60 * 1000; // 5 minutes

export function useTutorialCache() {
  const getCachedTutorial = (id: string): Tutorial | null => {
    const cached = tutorialCache.get(id);
    const timestamp = cacheTimestamp.get(id);
    
    if (cached && timestamp && Date.now() - timestamp < CACHE_TTL) {
      return cached;
    }
    
    // Remove expired cache
    if (cached) {
      tutorialCache.delete(id);
      cacheTimestamp.delete(id);
    }
    
    return null;
  };

  const setCachedTutorial = (id: string, tutorial: Tutorial) => {
    tutorialCache.set(id, tutorial);
    cacheTimestamp.set(id, Date.now());
  };

  const getCachedMetadata = (): TutorialMetadata[] => {
    return metadataCache.value;
  };

  const setCachedMetadata = (metadata: TutorialMetadata[]) => {
    metadataCache.value = metadata;
  };

  const clearCache = () => {
    tutorialCache.clear();
    cacheTimestamp.clear();
    metadataCache.value = [];
  };

  const clearTutorial = (id: string) => {
    tutorialCache.delete(id);
    cacheTimestamp.delete(id);
  };

  return {
    getCachedTutorial,
    setCachedTutorial,
    getCachedMetadata,
    setCachedMetadata,
    clearCache,
    clearTutorial,
  };
}
