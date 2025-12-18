<template>
  <div class="code-runner">
    <div class="flex justify-between items-center mb-2">
      <h4 class="text-sm font-semibold text-gray-700">Code Example</h4>
      <div class="flex gap-2">
        <button
          @click="copyCode"
          class="px-3 py-1 text-xs bg-gray-200 hover:bg-gray-300 rounded"
        >
          Copy
        </button>
        <button
          @click="executeCode"
          :disabled="executing"
          :class="[
            'px-3 py-1 text-xs rounded',
            executing
              ? 'bg-gray-400 text-gray-600 cursor-not-allowed'
              : 'bg-blue-500 text-white hover:bg-blue-600'
          ]"
        >
          {{ executing ? 'Running...' : 'Run Code' }}
        </button>
      </div>
    </div>
    
    <div class="mb-2">
      <pre class="bg-gray-900 text-gray-100 p-4 rounded-lg overflow-x-auto"><code>{{ code }}</code></pre>
    </div>

    <div v-if="result || executionError" class="mt-4">
      <div v-if="result && result.output" class="bg-green-50 border border-green-200 rounded p-3 mb-2">
        <div class="text-sm font-semibold text-green-800 mb-1">Output:</div>
        <pre class="text-sm text-green-700 whitespace-pre-wrap">{{ result.output }}</pre>
      </div>
      
      <div v-if="result && result.error" class="bg-red-50 border border-red-200 rounded p-3 mb-2">
        <div class="text-sm font-semibold text-red-800 mb-1">Error:</div>
        <pre class="text-sm text-red-700 whitespace-pre-wrap">{{ result.error }}</pre>
      </div>
      
      <div v-if="executionError" class="bg-red-50 border border-red-200 rounded p-3">
        <div class="text-sm font-semibold text-red-800 mb-1">Error:</div>
        <pre class="text-sm text-red-700">{{ executionError }}</pre>
      </div>
      
      <div v-if="result" class="text-xs text-gray-500 mt-2">
        Execution time: {{ result.duration }}
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useCodeExecution } from '../composables/useCodeExecution';

const props = defineProps<{
  code: string;
  language?: string;
}>();

const { executing, result, error: executionError, executeCode: execCode, clearResult } = useCodeExecution();

const executeCode = async () => {
  clearResult();
  await execCode(props.code);
};

const copyCode = async () => {
  try {
    await navigator.clipboard.writeText(props.code);
    // Could show a toast notification here
  } catch (err) {
    console.error('Failed to copy code', err);
  }
};
</script>

<style scoped>
.code-runner {
  margin: 1rem 0;
}
</style>

