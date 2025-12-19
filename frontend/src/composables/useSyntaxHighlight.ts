import { ref, onMounted } from 'vue';
import { createHighlighter, type Highlighter } from 'shiki';

let highlighterInstance: Highlighter | null = null;

export function useSyntaxHighlight() {
  const loading = ref(true);
  const error = ref<string | null>(null);

  const initHighlighter = async () => {
    if (highlighterInstance) {
      loading.value = false;
      return highlighterInstance;
    }

    try {
      highlighterInstance = await createHighlighter({
        themes: ['github-dark'],
        langs: ['go', 'javascript', 'typescript', 'json', 'markdown'],
      });
      loading.value = false;
      return highlighterInstance;
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
      const lang = highlighter.getLoadedLanguages().includes(language as any)
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
