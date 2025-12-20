import { ref, onMounted } from 'vue';
import { createHighlighter, type Highlighter } from 'shiki';

// Singleton highlighter instance shared across all components
let highlighterInstance: Highlighter | null = null;
// Promise lock to prevent concurrent initialization
let initPromise: Promise<Highlighter> | null = null;

export function useSyntaxHighlight() {
  const loading = ref(true);
  const error = ref<string | null>(null);

  const initHighlighter = async () => {
    // Return existing instance if available
    if (highlighterInstance) {
      loading.value = false;
      return highlighterInstance;
    }

    // If initialization is in progress, wait for it to complete
    if (initPromise) {
      try {
        const instance = await initPromise;
        loading.value = false;
        return instance;
      } catch (err) {
        loading.value = false;
        throw err;
      }
    }

    // Start initialization with promise lock to prevent concurrent creation
    initPromise = (async () => {
      try {
        const instance = await createHighlighter({
          themes: ['github-dark'],
          langs: ['go', 'javascript', 'typescript', 'json', 'markdown'],
        });
        highlighterInstance = instance;
        return instance;
      } catch (err) {
        // Reset promise on error so retries are possible
        initPromise = null;
        throw err;
      }
    })();

    try {
      const instance = await initPromise;
      loading.value = false;
      return instance;
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to initialize syntax highlighter';
      loading.value = false;
      throw err;
    }
  };

  const highlightCode = async (code: string, language: string = 'go'): Promise<string> => {
    const highlighter = await initHighlighter();
    if (!highlighter) {
      return code;
    }

    try {
      const loadedLanguages = highlighter.getLoadedLanguages();
      const lang = loadedLanguages.includes(language)
        ? language
        : 'go';
      return highlighter.codeToHtml(code, {
        lang,
        theme: 'github-dark',
      });
    } catch (err) {
      console.error('Failed to highlight code', err);
      return code;
    }
  };

  onMounted(() => {
    initHighlighter();
  });

  return {
    loading,
    error,
    highlightCode,
    initHighlighter,
  };
}
