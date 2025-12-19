import { ref, readonly, type Ref } from 'vue';

export interface RetryOptions {
  maxRetries?: number;
  retryDelay?: number;
  onRetry?: (attempt: number) => void;
}

export function useRetry<T>(
  fn: () => Promise<T>,
  options: RetryOptions = {}
): {
  execute: () => Promise<T>;
  retrying: Readonly<Ref<boolean>>;
  attempt: Readonly<Ref<number>>;
} {
  const {
    maxRetries = 3,
    retryDelay = 1000,
    onRetry,
  } = options;

  const retrying = ref(false);
  const attempt = ref(0);

  const sleep = (ms: number) => new Promise(resolve => setTimeout(resolve, ms));

  const execute = async (): Promise<T> => {
    let lastError: Error | null = null;
    attempt.value = 0;

    for (let i = 0; i <= maxRetries; i++) {
      attempt.value = i + 1;
      
      if (i > 0) {
        retrying.value = true;
        if (onRetry) {
          onRetry(i);
        }
        await sleep(retryDelay * i); // Exponential backoff
      }

      try {
        const result = await fn();
        retrying.value = false;
        return result;
      } catch (err) {
        lastError = err instanceof Error ? err : new Error(String(err));
        if (i === maxRetries) {
          retrying.value = false;
          throw lastError;
        }
      }
    }

    retrying.value = false;
    throw lastError || new Error('Failed after retries');
  };

  return {
    execute,
    retrying: readonly(retrying),
    attempt: readonly(attempt),
  };
}
