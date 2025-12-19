import { ref } from 'vue';
import type { ExecutionResult } from '../types/progress';
import { executionApi } from '../services/api';

export function useCodeExecution() {
  const executing = ref(false);
  const result = ref<ExecutionResult | null>(null);
  const error = ref<string | null>(null);

  const executeCode = async (code: string, snippet: boolean = false) => {
    executing.value = true;
    error.value = null;
    result.value = null;

    try {
      result.value = await executionApi.executeCode(code, snippet);
      if (result.value.error) {
        error.value = result.value.error;
      }
    } catch (err) {
      error.value = err instanceof Error ? err.message : 'Failed to execute code';
    } finally {
      executing.value = false;
    }
  };

  const clearResult = () => {
    result.value = null;
    error.value = null;
  };

  return {
    executing,
    result,
    error,
    executeCode,
    clearResult,
  };
}

