<template>
  <div class="code-editor">
    <!-- Editor header -->
    <div class="editor-header">
      <div class="header-title">
        <svg class="title-icon" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"/>
        </svg>
        <span>Code Editor</span>
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
        <button @click="resetCode" class="action-button" title="Reset to original">
          <svg class="action-icon" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"/>
          </svg>
          <span class="action-text">Reset</span>
        </button>
      </div>
    </div>

    <!-- Editor area -->
    <div class="editor-container">
      <div class="line-numbers">
        <span v-for="line in lineCount" :key="line" class="line-number">{{ line }}</span>
      </div>
      <textarea
        ref="textareaRef"
        v-model="localCode"
        @input="handleInput"
        @keydown="handleKeydown"
        class="editor-textarea"
        :placeholder="placeholder"
        spellcheck="false"
        autocomplete="off"
        autocorrect="off"
        autocapitalize="off"
      ></textarea>
    </div>

    <!-- Editor footer -->
    <div class="editor-footer">
      <span class="footer-hint">
        <svg class="hint-icon" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
        </svg>
        Edit the code and click "Run" to execute
      </span>
      <span class="footer-stats">
        {{ lineCount }} lines | {{ localCode.length }} chars
      </span>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue';

const props = defineProps<{
  modelValue: string;
  placeholder?: string;
}>();

const emit = defineEmits<{
  (e: 'update:modelValue', value: string): void;
}>();

const textareaRef = ref<HTMLTextAreaElement | null>(null);
const localCode = ref(props.modelValue);
const originalCode = ref(props.modelValue);
const copied = ref(false);

const lineCount = computed(() => {
  return localCode.value.split('\n').length;
});

const handleInput = () => {
  emit('update:modelValue', localCode.value);
};

const handleKeydown = (e: KeyboardEvent) => {
  if (e.key === 'Tab') {
    e.preventDefault();
    const textarea = textareaRef.value;
    if (!textarea) return;

    const start = textarea.selectionStart;
    const end = textarea.selectionEnd;
    const value = localCode.value;

    localCode.value = value.substring(0, start) + '\t' + value.substring(end);
    emit('update:modelValue', localCode.value);

    // Move cursor after tab
    setTimeout(() => {
      textarea.selectionStart = textarea.selectionEnd = start + 1;
    }, 0);
  }
};

const copyCode = async () => {
  try {
    await navigator.clipboard.writeText(localCode.value);
    copied.value = true;
    setTimeout(() => {
      copied.value = false;
    }, 2000);
  } catch (err) {
    console.error('Failed to copy code', err);
  }
};

const resetCode = () => {
  localCode.value = originalCode.value;
  emit('update:modelValue', originalCode.value);
};

watch(() => props.modelValue, (newValue) => {
  if (newValue !== localCode.value) {
    localCode.value = newValue;
    originalCode.value = newValue;
  }
});
</script>

<style scoped>
.code-editor {
  border: 1px solid var(--color-border);
  border-radius: var(--radius-lg);
  overflow: hidden;
  background-color: var(--color-neutral-900);
}

/* Header */
.editor-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 1rem;
  padding: 0.625rem 1rem;
  background-color: var(--color-neutral-800);
  border-bottom: 1px solid var(--color-neutral-700);
}

.header-title {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  font-size: 0.8125rem;
  font-weight: 600;
  color: var(--color-neutral-300);
}

.title-icon {
  width: 1rem;
  height: 1rem;
  color: var(--color-go-yellow);
}

.header-actions {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.action-button {
  display: inline-flex;
  align-items: center;
  gap: 0.375rem;
  padding: 0.375rem 0.625rem;
  font-size: 0.75rem;
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

/* Editor container */
.editor-container {
  display: flex;
  min-height: 16rem;
  max-height: 24rem;
  overflow: auto;
}

.line-numbers {
  display: flex;
  flex-direction: column;
  padding: 1rem 0.75rem;
  background-color: var(--color-neutral-950);
  border-right: 1px solid var(--color-neutral-800);
  user-select: none;
  text-align: right;
  min-width: 3rem;
}

.line-number {
  font-family: var(--font-mono);
  font-size: 0.875rem;
  line-height: 1.6;
  color: var(--color-neutral-600);
}

.editor-textarea {
  flex: 1;
  padding: 1rem;
  font-family: var(--font-mono);
  font-size: 0.875rem;
  line-height: 1.6;
  color: var(--color-neutral-100);
  background-color: transparent;
  border: none;
  outline: none;
  resize: none;
  tab-size: 2;
  white-space: pre;
  overflow-wrap: normal;
  overflow-x: auto;
}

.editor-textarea::placeholder {
  color: var(--color-neutral-600);
}

.editor-textarea:focus {
  outline: none;
}

/* Footer */
.editor-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 1rem;
  padding: 0.5rem 1rem;
  background-color: var(--color-neutral-800);
  border-top: 1px solid var(--color-neutral-700);
}

.footer-hint {
  display: flex;
  align-items: center;
  gap: 0.375rem;
  font-size: 0.6875rem;
  color: var(--color-neutral-500);
}

.hint-icon {
  width: 0.875rem;
  height: 0.875rem;
}

.footer-stats {
  font-size: 0.6875rem;
  font-family: var(--font-mono);
  color: var(--color-neutral-500);
}

/* Responsive */
@media (min-width: 640px) {
  .action-text {
    display: inline;
  }
}
</style>
