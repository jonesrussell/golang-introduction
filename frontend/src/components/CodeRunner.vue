<template>
  <div class="code-runner">
    <!-- Header -->
    <div class="runner-header">
      <div class="header-title">
        <svg class="title-icon" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 20l4-16m4 4l4 4-4 4M6 16l-4-4 4-4"/>
        </svg>
        <span>{{ editable ? 'Interactive Code' : 'Code Example' }}</span>
      </div>
      <div class="header-actions">
        <button @click="copyCode" class="action-button" title="Copy code">
          <svg v-if="!copied" class="action-icon" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16H6a2 2 0 01-2-2V6a2 2 0 012-2h8a2 2 0 012 2v2m-6 12h8a2 2 0 002-2v-8a2 2 0 00-2-2h-8a2 2 0 00-2 2v8a2 2 0 002 2z"/>
          </svg>
          <svg v-else class="action-icon action-icon-success" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"/>
          </svg>
          <span class="action-text">{{ copied ? 'Copied!' : 'Copy' }}</span>
        </button>
        <button
          v-if="editable"
          @click="toggleEdit"
          :class="['action-button', { 'action-button-active': editing }]"
          title="Toggle edit mode"
        >
          <svg class="action-icon" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"/>
          </svg>
          <span class="action-text">{{ editing ? 'View' : 'Edit' }}</span>
        </button>
        <button
          @click="executeCode"
          :disabled="executing"
          :class="['run-button', { 'run-button-executing': executing }]"
        >
          <svg v-if="!executing" class="run-icon" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.752 11.168l-3.197-2.132A1 1 0 0010 9.87v4.263a1 1 0 001.555.832l3.197-2.132a1 1 0 000-1.664z"/>
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
          </svg>
          <div v-else class="run-spinner"></div>
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
    <div v-else-if="highlightedCode" class="code-display">
      <div class="code-content" v-html="highlightedCode"></div>
    </div>

    <!-- Fallback plain code -->
    <div v-else class="code-display">
      <pre class="code-content code-plain"><code>{{ code }}</code></pre>
    </div>

    <!-- Results -->
    <div v-if="result || executionError" class="results-container">
      <!-- Success output -->
      <div v-if="result && result.output" class="result-block result-success">
        <div class="result-header">
          <svg class="result-icon" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"/>
          </svg>
          <span>Output</span>
        </div>
        <pre class="result-content">{{ result.output }}</pre>
      </div>

      <!-- Error output from execution -->
      <div v-if="result && result.error" class="result-block result-error">
        <div class="result-header">
          <svg class="result-icon" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4m0 4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
          </svg>
          <span>Error</span>
        </div>
        <pre class="result-content">{{ result.error }}</pre>
      </div>

      <!-- Execution error -->
      <div v-if="executionError" class="result-block result-error">
        <div class="result-header">
          <svg class="result-icon" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"/>
          </svg>
          <span>Execution Error</span>
        </div>
        <pre class="result-content">{{ executionError }}</pre>
      </div>

      <!-- Execution time -->
      <div v-if="result" class="execution-time">
        <svg class="time-icon" fill="none" viewBox="0 0 24 24" stroke="currentColor">
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
  await execCode(codeToExecute);
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
.code-runner {
  border: 1px solid var(--color-border);
  border-radius: var(--radius-lg);
  overflow: hidden;
  background-color: var(--color-surface);
}

/* Header */
.runner-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 1rem;
  padding: 0.75rem 1rem;
  background-color: var(--color-neutral-800);
  border-bottom: 1px solid var(--color-neutral-700);
}

.header-title {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  font-size: var(--text-sm);
  font-weight: 600;
  color: var(--color-neutral-300);
}

.title-icon {
  width: 1.125rem;
  height: 1.125rem;
  color: var(--color-go-blue-light);
}

.header-actions {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

/* Action buttons */
.action-button {
  display: inline-flex;
  align-items: center;
  gap: 0.375rem;
  padding: 0.5rem 0.75rem;
  font-size: var(--text-sm);
  font-weight: 500;
  color: var(--color-neutral-300);
  background-color: var(--color-neutral-700);
  border: none;
  border-radius: var(--radius-md);
  transition: all var(--transition-fast);
}

.action-button:hover {
  background-color: var(--color-neutral-600);
  color: white;
}

.action-button-active {
  background-color: var(--color-primary);
  color: white;
}

.action-icon {
  width: 0.875rem;
  height: 0.875rem;
}

.action-icon-success {
  color: var(--color-success);
}

.action-text {
  display: none;
}

/* Run button */
.run-button {
  display: inline-flex;
  align-items: center;
  gap: 0.375rem;
  padding: 0.5rem 1rem;
  font-size: var(--text-sm);
  font-weight: 600;
  color: white;
  background-color: var(--color-success);
  border: none;
  border-radius: var(--radius-md);
  transition: all var(--transition-fast);
}

.run-button:hover:not(:disabled) {
  background-color: #059669;
  transform: translateY(-1px);
}

.run-button:disabled {
  opacity: 0.7;
  cursor: not-allowed;
}

.run-button-executing {
  background-color: var(--color-warning);
}

.run-icon {
  width: 1rem;
  height: 1rem;
}

.run-spinner {
  width: 0.875rem;
  height: 0.875rem;
  border: 2px solid rgba(255, 255, 255, 0.3);
  border-top-color: white;
  border-radius: 50%;
  animation: spin 0.8s linear infinite;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

/* Code display */
.code-display {
  background-color: var(--color-neutral-900);
  overflow-x: auto;
}

.code-content {
  padding: 1.25rem;
  font-family: var(--font-mono);
  font-size: var(--text-sm);
  line-height: 1.7;
}

.code-plain {
  margin: 0;
  color: var(--color-neutral-100);
}

/* Deep styling for syntax highlighting */
:deep(.shiki) {
  background-color: transparent !important;
  padding: 1.25rem;
  margin: 0;
}

:deep(.shiki code) {
  font-family: var(--font-mono);
  font-size: var(--text-sm);
  line-height: 1.7;
}

/* Results */
.results-container {
  border-top: 1px solid var(--color-border);
  padding: 1rem;
  background-color: var(--color-background);
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.result-block {
  border-radius: var(--radius-md);
  overflow: hidden;
}

.result-header {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.625rem 1rem;
  font-size: var(--text-sm);
  font-weight: 600;
}

.result-icon {
  width: 1.125rem;
  height: 1.125rem;
}

.result-content {
  margin: 0;
  padding: 1rem;
  font-family: var(--font-mono);
  font-size: var(--text-sm);
  line-height: 1.6;
  white-space: pre-wrap;
  word-break: break-word;
}

.result-success {
  background-color: var(--color-success-light);
}

.result-success .result-header {
  color: #065f46;
  background-color: color-mix(in srgb, var(--color-success) 20%, var(--color-success-light));
}

.result-success .result-content {
  color: #047857;
}

.result-error {
  background-color: var(--color-error-light);
}

.result-error .result-header {
  color: #991b1b;
  background-color: color-mix(in srgb, var(--color-error) 20%, var(--color-error-light));
}

.result-error .result-content {
  color: #b91c1c;
}

.execution-time {
  display: flex;
  align-items: center;
  gap: 0.375rem;
  font-size: var(--text-sm);
  color: var(--color-text-subtle);
}

.time-icon {
  width: 1rem;
  height: 1rem;
}

/* Responsive */
@media (min-width: 640px) {
  .action-text {
    display: inline;
  }
}

/* Dark mode adjustments */
@media (prefers-color-scheme: dark) {
  .result-success {
    background-color: color-mix(in srgb, var(--color-success) 15%, var(--color-surface));
  }

  .result-success .result-header {
    color: #6ee7b7;
    background-color: color-mix(in srgb, var(--color-success) 25%, var(--color-surface));
  }

  .result-success .result-content {
    color: #a7f3d0;
  }

  .result-error {
    background-color: color-mix(in srgb, var(--color-error) 15%, var(--color-surface));
  }

  .result-error .result-header {
    color: #fca5a5;
    background-color: color-mix(in srgb, var(--color-error) 25%, var(--color-surface));
  }

  .result-error .result-content {
    color: #fecaca;
  }
}
</style>
