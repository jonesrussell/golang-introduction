<template>
  <div class="code-editor">
    <div class="flex justify-between items-center mb-2">
      <h4 class="text-sm font-semibold text-gray-700">Code Editor</h4>
      <div class="flex gap-2">
        <button
          @click="copyCode"
          class="px-3 py-1 text-xs bg-gray-200 hover:bg-gray-300 rounded"
        >
          Copy
        </button>
        <button
          @click="resetCode"
          class="px-3 py-1 text-xs bg-gray-200 hover:bg-gray-300 rounded"
        >
          Reset
        </button>
      </div>
    </div>
    
    <textarea
      v-model="localCode"
      @input="handleInput"
      class="w-full h-64 p-4 font-mono text-sm border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
      :placeholder="placeholder"
      spellcheck="false"
    ></textarea>
    
    <div class="mt-2 text-xs text-gray-500">
      Edit the code above and click "Run Code" to execute
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue';

const props = defineProps<{
  modelValue: string;
  placeholder?: string;
}>();

const emit = defineEmits<{
  (e: 'update:modelValue', value: string): void;
}>();

const localCode = ref(props.modelValue);
const originalCode = ref(props.modelValue);

const handleInput = () => {
  emit('update:modelValue', localCode.value);
};

const copyCode = async () => {
  try {
    await navigator.clipboard.writeText(localCode.value);
    // Could show a toast notification here
  } catch (err) {
    console.error('Failed to copy code', err);
  }
};

const resetCode = () => {
  localCode.value = originalCode.value;
  emit('update:modelValue', originalCode.value);
};

watch(() => props.modelValue, (newValue) => {
  localCode.value = newValue;
  originalCode.value = newValue;
});
</script>

<style scoped>
.code-editor {
  margin: 1rem 0;
}
</style>

