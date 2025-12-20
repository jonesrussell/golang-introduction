/**
 * Composable for copying text to clipboard with feedback.
 * Consolidates copy logic that was duplicated in CodeRunner and CodeEditor.
 */
import { ref } from 'vue';

/** Duration in ms to show "Copied!" feedback */
const COPIED_FEEDBACK_DURATION = 2000;

/**
 * Composable that provides clipboard copy functionality with feedback state.
 */
export function useCopyToClipboard() {
  const copied = ref(false);
  let timeoutId: ReturnType<typeof setTimeout> | null = null;

  /**
   * Copies text to clipboard and sets copied state.
   * @param text The text to copy
   * @returns Promise that resolves to true on success, false on failure
   */
  const copyToClipboard = async (text: string): Promise<boolean> => {
    try {
      await navigator.clipboard.writeText(text);
      copied.value = true;

      // Clear any existing timeout
      if (timeoutId) {
        clearTimeout(timeoutId);
      }

      // Reset copied state after delay
      timeoutId = setTimeout(() => {
        copied.value = false;
        timeoutId = null;
      }, COPIED_FEEDBACK_DURATION);

      return true;
    } catch (err) {
      console.error('Failed to copy to clipboard:', err);
      return false;
    }
  };

  /**
   * Resets the copied state manually.
   */
  const resetCopied = () => {
    if (timeoutId) {
      clearTimeout(timeoutId);
      timeoutId = null;
    }
    copied.value = false;
  };

  return {
    copied,
    copyToClipboard,
    resetCopied,
  };
}
