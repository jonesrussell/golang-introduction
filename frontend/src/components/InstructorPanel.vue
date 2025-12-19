<template>
  <div v-if="notes" class="instructor-panel mb-6">
    <button
      type="button"
      @click="expanded = !expanded"
      class="w-full flex items-center justify-between gap-3 p-4 bg-amber-50 dark:bg-amber-950/30 border border-amber-200 dark:border-amber-800 rounded-xl text-left transition-all duration-150 hover:bg-amber-100 dark:hover:bg-amber-950/50"
    >
      <div class="flex items-center gap-3">
        <div class="p-2 bg-amber-100 dark:bg-amber-900/50 rounded-lg">
          <svg class="w-5 h-5 text-amber-600 dark:text-amber-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z"/>
          </svg>
        </div>
        <div>
          <h3 class="text-base font-semibold text-amber-900 dark:text-amber-100 m-0">Instructor Notes</h3>
          <p class="text-sm text-amber-700 dark:text-amber-300 m-0">Teaching tips and talking points</p>
        </div>
      </div>
      <svg 
        class="w-5 h-5 text-amber-600 dark:text-amber-400 transition-transform duration-200"
        :class="{ 'rotate-180': expanded }"
        fill="none" 
        viewBox="0 0 24 24" 
        stroke="currentColor"
      >
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9l-7 7-7-7"/>
      </svg>
    </button>

    <div 
      v-show="expanded"
      class="mt-2 p-6 bg-amber-50/50 dark:bg-amber-950/20 border border-amber-200 dark:border-amber-800 rounded-xl"
    >
      <div class="prose prose-amber dark:prose-invert max-w-none instructor-content" v-html="renderedNotes"></div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue';

const props = defineProps<{
  notes: string;
}>();

const expanded = ref(true);

// Simple markdown rendering (basic conversion)
const renderedNotes = computed(() => {
  if (!props.notes) return '';
  
  let html = props.notes;
  
  // Convert headers
  html = html.replace(/^### (.+)$/gm, '<h3 class="text-lg font-semibold text-amber-900 dark:text-amber-100 mt-6 mb-3">$1</h3>');
  html = html.replace(/^## (.+)$/gm, '<h2 class="text-xl font-bold text-amber-900 dark:text-amber-100 mt-6 mb-4">$1</h2>');
  html = html.replace(/^# (.+)$/gm, '<h1 class="text-2xl font-bold text-amber-900 dark:text-amber-100 mb-4">$1</h1>');
  
  // Convert bold and italic
  html = html.replace(/\*\*(.+?)\*\*/g, '<strong class="font-semibold text-amber-900 dark:text-amber-100">$1</strong>');
  html = html.replace(/\*(.+?)\*/g, '<em>$1</em>');
  
  // Convert inline code
  html = html.replace(/`([^`]+)`/g, '<code class="px-1.5 py-0.5 bg-amber-100 dark:bg-amber-900/50 text-amber-800 dark:text-amber-200 rounded text-sm font-mono">$1</code>');
  
  // Convert code blocks
  html = html.replace(/```(\w*)\n([\s\S]*?)```/g, (_, lang, code) => {
    return `<pre class="p-4 bg-neutral-900 rounded-lg overflow-x-auto my-4"><code class="text-sm font-mono text-neutral-100">${escapeHtml(code.trim())}</code></pre>`;
  });
  
  // Convert bullet lists
  html = html.replace(/^- (.+)$/gm, '<li class="text-amber-800 dark:text-amber-200 ml-4">$1</li>');
  html = html.replace(/(<li.*?<\/li>\n?)+/g, '<ul class="list-disc space-y-1 my-3">$&</ul>');
  
  // Convert paragraphs (lines not already converted)
  html = html.split('\n\n').map(para => {
    if (para.trim() && !para.startsWith('<')) {
      return `<p class="text-amber-800 dark:text-amber-200 leading-relaxed my-3">${para}</p>`;
    }
    return para;
  }).join('\n');
  
  return html;
});

function escapeHtml(text: string): string {
  const div = document.createElement('div');
  div.textContent = text;
  return div.innerHTML;
}
</script>

<style scoped>
.instructor-content :deep(h1),
.instructor-content :deep(h2),
.instructor-content :deep(h3) {
  &:first-child {
    margin-top: 0;
  }
}

.instructor-content :deep(ul) {
  list-style-type: disc;
  padding-left: 1.5rem;
}

.instructor-content :deep(li) {
  margin-left: 0;
}

.instructor-content :deep(pre) {
  margin: 1rem 0;
}

.instructor-content :deep(p:first-child) {
  margin-top: 0;
}

.instructor-content :deep(p:last-child) {
  margin-bottom: 0;
}
</style>
