<template>
  <div class="bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 rounded-2xl shadow-sm overflow-hidden">
    <!-- Section header -->
    <div class="p-6 bg-gradient-to-br from-[#e6f7fb] to-white dark:from-neutral-800 dark:to-neutral-900 border-b border-neutral-200 dark:border-neutral-800 sm:p-5">
      <div class="flex justify-between items-start gap-4 mb-4 sm:flex-col sm:gap-2">
        <h2 class="text-2xl font-bold text-neutral-900 dark:text-neutral-100 m-0 leading-tight sm:text-xl">{{ section.title }}</h2>
        <div class="text-sm text-neutral-600 dark:text-neutral-400 whitespace-nowrap py-1.5 px-3.5 bg-white dark:bg-neutral-900 rounded-md border border-neutral-200 dark:border-neutral-800">
          Section {{ sectionIndex + 1 }} of {{ totalSections }}
        </div>
      </div>

      <!-- Progress bar -->
      <div class="pt-2">
        <div class="h-2 bg-neutral-200 dark:bg-neutral-800 rounded-full overflow-hidden">
          <div
            class="h-full bg-gradient-to-r from-[#00ADD8] to-[#5DC9E2] rounded-full transition-all duration-300"
            :style="{ width: `${((sectionIndex + 1) / totalSections) * 100}%` }"
          ></div>
        </div>
      </div>
    </div>

    <!-- Section content -->
    <div class="p-6 flex flex-col gap-8 sm:p-5 sm:gap-6">
      <!-- Table of Contents (from content) -->
      <div v-if="tableOfContents" class="animate-slide-up p-6 bg-gradient-to-br from-blue-50 to-indigo-50 dark:from-blue-950/20 dark:to-indigo-950/20 rounded-xl border border-blue-200 dark:border-blue-900/50">
        <!-- eslint-disable-next-line vue/no-v-html -->
        <div v-html="tableOfContents" class="toc-content prose prose-blue dark:prose-invert max-w-none"></div>
      </div>

      <!-- Topics -->
      <div v-if="section.topics && section.topics.length > 0" class="animate-slide-up">
        <h3 class="flex items-center gap-2 text-lg font-semibold text-neutral-900 dark:text-neutral-100 m-0 mb-4">
          <svg class="w-5.5 h-5.5 text-[#00ADD8]" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-3 7h3m-3 4h3m-6-4h.01M9 16h.01"/>
          </svg>
          Topics to Cover
        </h3>
        <ul class="list-none p-0 m-0 flex flex-col gap-2.5">
          <li v-for="topic in section.topics" :key="topic" class="flex items-start gap-3 text-base text-neutral-900 dark:text-neutral-100 leading-relaxed">
            <svg class="w-4.5 h-4.5 text-[#00ADD8] flex-shrink-0 mt-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/>
            </svg>
            <!-- eslint-disable-next-line vue/no-v-html -->
            <span v-html="renderMarkdown(topic)"></span>
          </li>
        </ul>
      </div>

      <!-- Code examples -->
      <div v-if="section.codeExamples && section.codeExamples.length > 0" class="flex flex-col gap-6 animate-slide-up">
        <div v-for="example in section.codeExamples" :key="example.id" class="rounded-xl overflow-hidden">
          <CodeRunner
            :code="example.code"
            :language="example.language"
            :editable="example.runnable"
            :snippet="example.snippet"
          />
        </div>
      </div>

      <!-- Teaching points -->
      <div v-if="section.teachingPoints && section.teachingPoints.length > 0" class="animate-slide-up">
        <h3 class="flex items-center gap-2 text-lg font-semibold text-neutral-900 dark:text-neutral-100 m-0 mb-4">
          <svg class="w-5.5 h-5.5 text-[#00ADD8]" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z"/>
          </svg>
          Key Takeaways
        </h3>
        <ul class="list-none p-0 m-0 flex flex-col gap-3">
          <li v-for="point in section.teachingPoints" :key="point" class="flex items-start gap-3 text-base text-neutral-900 dark:text-neutral-100 leading-relaxed p-4 bg-green-50 dark:bg-green-950/20 rounded-md">
            <svg class="w-4.5 h-4.5 text-green-500 flex-shrink-0 mt-0.5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"/>
            </svg>
            <!-- eslint-disable-next-line vue/no-v-html -->
            <span v-html="renderMarkdown(point)"></span>
          </li>
        </ul>
      </div>
    </div>

    <!-- Navigation footer -->
    <div class="flex justify-between items-center gap-4 px-6 py-5 bg-neutral-50 dark:bg-neutral-950 border-t border-neutral-200 dark:border-neutral-800 sm:flex-wrap sm:p-4">
      <button
        type="button"
        :disabled="sectionIndex === 0"
        class="inline-flex items-center gap-2 px-5 py-3 text-base font-medium rounded-xl border-none transition-all duration-150 sm:px-3 sm:py-2 sm:text-sm disabled:opacity-40 disabled:cursor-not-allowed bg-white dark:bg-neutral-900 text-neutral-900 dark:text-neutral-100 border border-neutral-200 dark:border-neutral-800 hover:bg-neutral-100 dark:hover:bg-neutral-800 hover:border-neutral-300 dark:hover:border-neutral-700"
        @click="$emit('previous')"
      >
        <svg class="w-4.5 h-4.5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"/>
        </svg>
        <span class="hidden sm:inline">Previous</span>
      </button>

      <div class="flex items-center gap-4">
        <button
          type="button"
          class="inline-flex items-center gap-2 px-6 py-3 text-base font-semibold rounded-xl border-none transition-all duration-150 sm:w-full sm:justify-center sm:mb-2 sm:order-first" :class="[
            isComplete
              ? 'bg-green-600 text-white'
              : 'bg-green-500 text-white hover:bg-green-600 hover:-translate-y-px'
          ]"
          @click="handleComplete"
        >
          <svg v-if="isComplete" class="w-4.5 h-4.5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"/>
          </svg>
          <svg v-else class="w-4.5 h-4.5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"/>
          </svg>
          {{ isComplete ? 'Completed' : 'Mark Complete' }}
        </button>
      </div>

      <button
        type="button"
        :disabled="sectionIndex >= totalSections - 1"
        class="inline-flex items-center gap-2 px-5 py-3 text-base font-medium rounded-xl border-none transition-all duration-150 sm:px-3 sm:py-2 sm:text-sm disabled:opacity-40 disabled:cursor-not-allowed bg-[#00ADD8] text-white hover:bg-[#007D9C]"
        @click="$emit('next')"
      >
        <span class="hidden sm:inline">Next</span>
        <svg class="w-4.5 h-4.5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/>
        </svg>
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { useProgressStore } from '../stores/progress';
import CodeRunner from './CodeRunner.vue';
import type { Section } from '../types/tutorial';

const props = defineProps<{
  section: Section;
  sectionIndex: number;
  totalSections: number;
  tutorialId?: string;
  instructorMode?: boolean;
}>();

/* eslint-disable no-unused-vars */
const emit = defineEmits<{
  (e: 'next'): void;
  (e: 'previous'): void;
  (e: 'complete'): void;
}>();
/* eslint-enable no-unused-vars */

const progressStore = useProgressStore();

const isComplete = computed(() => {
  if (!props.tutorialId) return false;
  return progressStore.isSectionComplete(props.tutorialId, props.section.id);
});

const handleComplete = () => {
  emit('complete');
};

// Extract table of contents from section content
const tableOfContents = computed(() => {
  if (!props.section.content) return null;
  
  const content = props.section.content;
  const tocStart = content.indexOf('## Table of Contents');
  if (tocStart === -1) return null;
  
  // Find the end of the TOC section (next ## header that's not part of TOC)
  const afterToc = content.slice(tocStart);
  // Look for next section header after the TOC
  const nextSectionMatch = afterToc.match(/\n## (?!Table of Contents)/);
  const tocEnd = nextSectionMatch ? tocStart + nextSectionMatch.index : content.length;
  
  const tocContent = content.slice(tocStart, tocEnd).trim();
  if (!tocContent) return null;
  
  return renderMarkdownContent(tocContent);
});

// Convert markdown links and inline code to HTML
function renderMarkdown(text: string): string {
  let html = text;
  
  // Convert markdown links [text](url) to HTML anchors
  html = html.replace(
    /\[([^\]]+)\]\(([^)]+)\)/g,
    '<a href="$2" target="_blank" rel="noopener noreferrer" class="text-[#00ADD8] hover:text-[#007D9C] underline underline-offset-2">$1</a>'
  );
  
  // Convert inline code `code` to styled spans
  html = html.replace(
    /`([^`]+)`/g,
    '<code class="px-1.5 py-0.5 bg-neutral-100 dark:bg-neutral-800 text-[#00ADD8] rounded text-sm font-mono">$1</code>'
  );
  
  return html;
}

// Render markdown content with full support for headers, lists, etc.
function renderMarkdownContent(content: string): string {
  const lines = content.split('\n');
  const processedLines: string[] = [];
  let inList = false;
  let listItems: string[] = [];
  
  for (let i = 0; i < lines.length; i++) {
    const line = lines[i];
    const trimmed = line.trim();
    
    // Skip empty lines
    if (!trimmed) {
      if (listItems.length > 0) {
        processedLines.push(`<ol class="list-decimal list-inside space-y-2.5 my-4 pl-6 marker:text-[#00ADD8] marker:font-semibold">${listItems.join('')}</ol>`);
        listItems = [];
        inList = false;
      }
      continue;
    }
    
    // Check for headers
    if (trimmed.startsWith('## ')) {
      if (listItems.length > 0) {
        processedLines.push(`<ol class="list-decimal list-inside space-y-2.5 my-4 pl-6 marker:text-[#00ADD8] marker:font-semibold">${listItems.join('')}</ol>`);
        listItems = [];
        inList = false;
      }
      const headerText = trimmed.replace(/^##\s+/, '');
      processedLines.push(`<h2 class="text-xl font-bold text-neutral-900 dark:text-neutral-100 mt-6 mb-4">${headerText}</h2>`);
      continue;
    }
    
    if (trimmed.startsWith('### ')) {
      if (listItems.length > 0) {
        processedLines.push(`<ol class="list-decimal list-inside space-y-2.5 my-4 pl-6 marker:text-[#00ADD8] marker:font-semibold">${listItems.join('')}</ol>`);
        listItems = [];
        inList = false;
      }
      const headerText = trimmed.replace(/^###\s+/, '');
      processedLines.push(`<h3 class="text-lg font-semibold text-neutral-900 dark:text-neutral-100 mt-4 mb-3">${headerText}</h3>`);
      continue;
    }
    
    // Check for numbered list items (1. **Title** - description)
    const listMatch = trimmed.match(/^(\d+)\.\s+(.+)$/);
    if (listMatch) {
      inList = true;
      let itemContent = listMatch[2];
      
      // Convert bold text
      itemContent = itemContent.replace(/\*\*(.+?)\*\*/g, '<strong class="font-semibold">$1</strong>');
      
      // Convert markdown links
      itemContent = itemContent.replace(
        /\[([^\]]+)\]\(([^)]+)\)/g,
        '<a href="$2" target="_blank" rel="noopener noreferrer" class="text-[#00ADD8] hover:text-[#007D9C] underline underline-offset-2">$1</a>'
      );
      
      listItems.push(`<li class="mb-2.5 text-base text-neutral-900 dark:text-neutral-100 leading-relaxed">${itemContent}</li>`);
      continue;
    }
    
    // If we were in a list but hit a non-list line, close the list
    if (inList && listItems.length > 0) {
      processedLines.push(`<ol class="list-decimal list-inside space-y-2.5 my-4 pl-6 marker:text-[#00ADD8] marker:font-semibold">${listItems.join('')}</ol>`);
      listItems = [];
      inList = false;
    }
    
    // Regular paragraph text
    let paraText = trimmed;
    
    // Convert bold text (but not if already converted)
    paraText = paraText.replace(/\*\*(.+?)\*\*/g, '<strong class="font-semibold">$1</strong>');
    
    // Convert markdown links
    paraText = paraText.replace(
      /\[([^\]]+)\]\(([^)]+)\)/g,
      '<a href="$2" target="_blank" rel="noopener noreferrer" class="text-[#00ADD8] hover:text-[#007D9C] underline underline-offset-2">$1</a>'
    );
    
    // Convert inline code
    paraText = paraText.replace(
      /`([^`]+)`/g,
      '<code class="px-1.5 py-0.5 bg-neutral-100 dark:bg-neutral-800 text-[#00ADD8] rounded text-sm font-mono">$1</code>'
    );
    
    processedLines.push(`<p class="text-base text-neutral-900 dark:text-neutral-100 leading-relaxed my-3">${paraText}</p>`);
  }
  
  // Close any remaining list
  if (listItems.length > 0) {
    processedLines.push(`<ol class="list-decimal list-inside space-y-2.5 my-4 pl-6 marker:text-[#00ADD8] marker:font-semibold">${listItems.join('')}</ol>`);
  }
  
  return processedLines.join('\n');
}
</script>

