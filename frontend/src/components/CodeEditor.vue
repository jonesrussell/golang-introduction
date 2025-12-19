<template>
  <div class="border border-neutral-200 dark:border-neutral-800 rounded-xl overflow-hidden bg-neutral-900">
    <!-- Editor header -->
    <div class="flex justify-between items-center gap-4 px-4 py-2.5 bg-neutral-800 border-b border-neutral-700">
      <div class="flex items-center gap-2 text-sm font-semibold text-neutral-300">
        <svg class="w-4.5 h-4.5 text-go-yellow" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z"/>
        </svg>
        <span>Code Editor</span>
      </div>
      <div class="flex items-center gap-2">
        <button
          type="button"
          class="inline-flex items-center gap-1.5 px-3 py-2 text-sm font-medium text-neutral-300 bg-neutral-700 border-none rounded-md transition-all duration-150 hover:bg-neutral-600 hover:text-white"
          title="Copy code"
          @click="copyCode"
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
          type="button"
          class="inline-flex items-center gap-1.5 px-3 py-2 text-sm font-medium text-neutral-300 bg-neutral-700 border-none rounded-md transition-all duration-150 hover:bg-neutral-600 hover:text-white"
          title="Reset to original"
          @click="resetCode"
        >
          <svg class="w-3.5 h-3.5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"/>
          </svg>
          <span class="hidden sm:inline">Reset</span>
        </button>
      </div>
    </div>

    <!-- Editor area -->
    <div class="flex min-h-64 max-h-96 overflow-auto">
      <div class="flex flex-col py-5 px-3.5 bg-neutral-950 border-r border-neutral-800 select-none text-right min-w-14">
        <span
          v-for="line in lineCount"
          :key="line"
          class="font-mono text-sm leading-[1.7] text-neutral-600"
        >{{ line }}</span>
      </div>
      <textarea
        ref="textareaRef"
        v-model="localCode"
        class="flex-1 py-5 px-5 font-mono text-sm leading-[1.7] text-neutral-100 bg-transparent border-none outline-none resize-none whitespace-pre overflow-wrap-normal overflow-x-auto placeholder:text-neutral-600 code-editor-textarea"
        :placeholder="placeholder"
        spellcheck="false"
        autocomplete="off"
        autocorrect="off"
        autocapitalize="off"
        @input="handleInput"
        @keydown="handleKeydown"
      ></textarea>
    </div>

    <!-- Editor footer -->
    <div class="flex justify-between items-center gap-4 px-4 py-2 bg-neutral-800 border-t border-neutral-700">
      <span class="flex items-center gap-1.5 text-xs text-neutral-500">
        <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"/>
        </svg>
        Edit the code and click "Run" to execute
      </span>
      <span class="text-xs font-mono text-neutral-500">
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

/* eslint-disable no-unused-vars */
const emit = defineEmits<{
  (e: 'update:modelValue', value: string): void;
}>();
/* eslint-enable no-unused-vars */

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
.code-editor-textarea {
  tab-size: 2;
}
</style>

