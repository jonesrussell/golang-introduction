<template>
  <div class="border border-neutral-200 dark:border-neutral-800 rounded-xl overflow-hidden bg-white dark:bg-neutral-900">
    <!-- Header -->
    <div class="flex justify-between items-center gap-4 px-4 py-3 bg-neutral-800 border-b border-neutral-700">
      <div class="flex items-center gap-2 text-sm font-semibold text-neutral-300">
        <svg class="w-4.5 h-4.5 text-[#5DC9E2]" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4"/>
        </svg>
        <span>{{ editable ? 'Interactive Code' : 'Code Example' }}</span>
      </div>
      <div class="flex items-center gap-2">
        <button
          @click="copyCode"
          class="inline-flex items-center gap-1.5 px-3 py-2 text-sm font-medium text-neutral-300 bg-neutral-700 border-none rounded-md transition-all duration-150 hover:bg-neutral-600 hover:text-white"
          title="Copy code"
        >
          <svg v-if="!copied" class="w-3.5 h-3.5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"/>
          </svg>
          <svg v-else class="w-3.5 h-3.5 text-green-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"/>
          </svg>
          <span class="hidden sm:inline">{{ copied ? 'Copied!' : 'Copy' }}</span>
        </button>
        <button
          v-if="editable"
          @click="toggleEdit"
          :class="[
            'inline-flex items-center gap-1.5 px-3 py-2 text-sm font-medium border-none rounded-md transition-all duration-150',
            editing
              ? 'bg-[#00ADD8] text-white'
              : 'text-neutral-300 bg-neutral-700 hover:bg-neutral-600 hover:text-white'
          ]"
          title="Toggle edit mode"
        >
          <svg class="w-3.5 h-3.5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"/>
          </svg>
          <span class="hidden sm:inline">{{ editing ? 'View' : 'Edit' }}</span>
        </button>
        <button
          @click="executeCode"
          :disabled="executing"
          :class="[
            'inline-flex items-center gap-1.5 px-4 py-2 text-sm font-semibold text-white border-none rounded-md transition-all duration-150',
            executing
              ? 'bg-amber-500 cursor-not-allowed opacity-70'
              : 'bg-green-500 hover:bg-green-600 hover:-translate-y-px'
          ]"
        >
          <svg v-if="!executing" class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.752 11.168l-3.197-2.132A1 1 0 0010 9.87v4.263a1 1 0 001.555.832l3.197-2.132a1 1 0 000-1.664z"/>
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
          </svg>
          <div
            v-else
            class="w-3.5 h-3.5 border-2 border-white/30 border-t-white rounded-full animate-spin"
          ></div>
          <span>{{ executing ? 'Running...' : 'Run' }}</span>
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
    <div v-else-if="highlightedCode" class="bg-neutral-900 overflow-x-auto">
      <div class="p-5 [&:deep(.shiki)]:bg-transparent [&:deep(.shiki)]:p-5 [&:deep(.shiki)]:m-0 [&:deep(.shiki_code)]:font-mono [&:deep(.shiki_code)]:text-sm [&:deep(.shiki_code)]:leading-[1.7]" v-html="highlightedCode"></div>
    </div>

    <!-- Fallback plain code -->
    <div v-else class="bg-neutral-900 overflow-x-auto">
      <pre class="p-5 m-0 font-mono text-sm leading-relaxed text-neutral-100 whitespace-pre-wrap break-words"><code>{{ code }}</code></pre>
    </div>

    <!-- Results -->
    <div v-if="result || executionError" class="border-t border-neutral-200 dark:border-neutral-800 p-4 bg-neutral-50 dark:bg-neutral-950 flex flex-col gap-3">
      <!-- Success output -->
      <div v-if="result && result.output" class="rounded-md overflow-hidden bg-green-50 dark:bg-green-950/20">
        <div class="flex items-center gap-2 px-4 py-2.5 text-sm font-semibold text-green-800 dark:text-green-300 bg-green-100 dark:bg-green-900/30">
          <svg class="w-4.5 h-4.5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"/>
          </svg>
          <span>Output</span>
        </div>
        <pre class="m-0 p-4 font-mono text-sm leading-relaxed whitespace-pre-wrap break-words text-green-700 dark:text-green-200">{{ result.output }}</pre>
      </div>

      <!-- Error output from execution -->
      <div v-if="result && result.error" class="rounded-md overflow-hidden bg-red-50 dark:bg-red-950/20">
        <div class="flex items-center gap-2 px-4 py-2.5 text-sm font-semibold text-red-800 dark:text-red-300 bg-red-100 dark:bg-red-900/30">
          <svg class="w-4.5 h-4.5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
          </svg>
          <span>Error</span>
        </div>
        <pre class="m-0 p-4 font-mono text-sm leading-relaxed whitespace-pre-wrap break-words text-red-700 dark:text-red-200">{{ result.error }}</pre>
      </div>

      <!-- Execution error -->
      <div v-if="executionError" class="rounded-md overflow-hidden bg-red-50 dark:bg-red-950/20">
        <div class="flex items-center gap-2 px-4 py-2.5 text-sm font-semibold text-red-800 dark:text-red-300 bg-red-100 dark:bg-red-900/30">
          <svg class="w-4.5 h-4.5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"/>
          </svg>
          <span>Execution Error</span>
        </div>
        <pre class="m-0 p-4 font-mono text-sm leading-relaxed whitespace-pre-wrap break-words text-red-700 dark:text-red-200">{{ executionError }}</pre>
      </div>

      <!-- Execution time -->
      <div v-if="result" class="flex items-center gap-1.5 text-sm text-neutral-500 dark:text-neutral-500">
        <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"/>
        </svg>
        Executed in {{ result.duration }}
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
  snippet?: boolean;
}>();

const { executing, result, error: executionError, executeCode: execCode, clearResult } = useCodeExecution();
const { highlightCode } = useSyntaxHighlight();

const editing = ref(false);
const editableCode = ref(props.code);
const highlightedCode = ref<string>('');
const copied = ref(false);

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
    loadHighlightedCode();
  }
};

const executeCode = async () => {
  clearResult();
  const codeToExecute = editing.value && props.editable ? editableCode.value : props.code;
  await execCode(codeToExecute, props.snippet || false);
};

const copyCode = async () => {
  try {
    await navigator.clipboard.writeText(currentCode.value);
    copied.value = true;
    setTimeout(() => {
      copied.value = false;
    }, 2000);
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
  loadHighlightedCode();
});
</script>

<style scoped>
/* Deep styling for syntax highlighting */
:deep(.shiki) {
  background-color: transparent !important;
  padding: 1.25rem;
  margin: 0;
}

:deep(.shiki code) {
  font-family: var(--font-family-mono);
  font-size: 0.9375rem;
  line-height: 1.7;
}
</style>
