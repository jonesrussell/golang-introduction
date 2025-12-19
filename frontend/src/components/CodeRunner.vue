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
          v-if="editable"
          @click="toggleEdit"
          class="px-3 py-1 text-xs bg-gray-200 hover:bg-gray-300 rounded"
        >
          {{ editing ? 'View' : 'Edit' }}
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
    
    <!-- Editable Code Editor -->
    <CodeEditor
      v-if="editable && editing"
      v-model="editableCode"
      placeholder="Edit the code here..."
    />
    
    <!-- Syntax Highlighted Code Display -->
    <div v-else-if="highlightedCode" class="mb-2">
      <div
        class="bg-gray-900 rounded-lg overflow-x-auto"
        v-html="highlightedCode"
      ></div>
    </div>
    
    <!-- Fallback plain code -->
    <div v-else class="mb-2">
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
import { ref, computed, watch, onMounted } from 'vue';
import { useCodeExecution } from '../composables/useCodeExecution';
import { useSyntaxHighlight } from '../composables/useSyntaxHighlight';
import CodeEditor from './CodeEditor.vue';

const props = defineProps<{
  code: string;
  language?: string;
  editable?: boolean;
}>();

const { executing, result, error: executionError, executeCode: execCode, clearResult } = useCodeExecution();
const { highlightCode, loading: highlightLoading } = useSyntaxHighlight();

const editing = ref(false);
const editableCode = ref(props.code);
const highlightedCode = ref<string>('');

const currentCode = computed(() => {
  return editing.value && props.editable ? editableCode.value : props.code;
});

const loadHighlightedCode = async () => {
  if (!editing.value) {
    try {
      highlightedCode.value = await highlightCode(props.code, props.language || 'go');
    } catch (err) {
      console.error('Failed to highlight code', err);
      highlightedCode.value = '';
    }
  }
};

const toggleEdit = () => {
  editing.value = !editing.value;
  if (!editing.value) {
    // Reload highlighted code when switching back to view mode
    loadHighlightedCode();
  }
};

const executeCode = async () => {
  clearResult();
  const codeToExecute = editing.value && props.editable ? editableCode.value : props.code;
  await execCode(codeToExecute);
};

const copyCode = async () => {
  try {
    await navigator.clipboard.writeText(currentCode.value);
    // Could show a toast notification here
  } catch (err) {
    console.error('Failed to copy code', err);
  }
};

watch(() => props.code, () => {
  editableCode.value = props.code;
  if (!editing.value) {
    loadHighlightedCode();
  }
});

onMounted(() => {
  if (!props.editable) {
    loadHighlightedCode();
  }
});
</script>

<style scoped>
.code-runner {
  margin: 1rem 0;
}
</style>

