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
      <div v-if="parsedTableOfContents" class="animate-slide-up p-6 bg-gradient-to-br from-blue-50 to-indigo-50 dark:from-blue-950/20 dark:to-indigo-950/20 rounded-xl border border-blue-200 dark:border-blue-900/50">
        <h2 class="text-xl font-bold text-neutral-900 dark:text-neutral-100 mt-0 mb-4">Table of Contents</h2>
        <p v-if="parsedTableOfContents.intro" class="text-base text-neutral-900 dark:text-neutral-100 leading-relaxed my-3" v-html="renderMarkdown(parsedTableOfContents.intro)"></p>
        <ol class="list-decimal list-inside space-y-2.5 my-4 pl-6 marker:text-[#00ADD8] marker:font-semibold">
          <li
            v-for="(item, index) in parsedTableOfContents.items"
            :key="index"
            class="mb-2.5 text-base leading-relaxed cursor-pointer transition-all duration-150 rounded-md -ml-2 pl-2 py-1"
            :class="{
              'text-[#00ADD8] font-semibold bg-blue-100 dark:bg-blue-900/30 border-l-4 border-[#00ADD8]': index === sectionIndex,
              'text-neutral-900 dark:text-neutral-100 hover:text-[#00ADD8] hover:bg-blue-50 dark:hover:bg-blue-950/20': index !== sectionIndex
            }"
            @click="handleTocClick(index)"
          >
            <strong class="font-semibold">{{ item.title }}</strong> - <span v-html="renderMarkdown(item.description)"></span>
          </li>
        </ol>
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
import { useMarkdownRenderer } from '../composables/useMarkdownRenderer';
import CodeRunner from './CodeRunner.vue';
import type { Section } from '../types/tutorial';

const { renderMarkdown } = useMarkdownRenderer();

const props = defineProps<{
  section: Section;
  sectionIndex: number;
  totalSections: number;
  tutorialId?: string;
  instructorMode?: boolean;
  tableOfContents?: string;
}>();

/* eslint-disable no-unused-vars */
const emit = defineEmits<{
  (e: 'next'): void;
  (e: 'previous'): void;
  (e: 'complete'): void;
  (e: 'navigate-to-section', index: number): void;
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

// Parse table of contents from tutorial
interface TocItem {
  title: string;
  description: string;
}

interface ParsedToc {
  intro?: string;
  items: TocItem[];
}

const parsedTableOfContents = computed((): ParsedToc | null => {
  if (!props.tableOfContents) return null;
  
  const content = props.tableOfContents.trim();
  const lines = content.split('\n');
  
  const result: ParsedToc = {
    items: []
  };
  
  let introLines: string[] = [];
  let foundFirstItem = false;
  
  for (const line of lines) {
    const trimmed = line.trim();
    
    // Skip empty lines at the start
    if (!trimmed && !foundFirstItem) continue;
    
    // Check for numbered list item (1. **Title** - description)
    const itemMatch = trimmed.match(/^(\d+)\.\s+\*\*(.+?)\*\*\s+-\s+(.+)$/);
    if (itemMatch) {
      foundFirstItem = true;
      // Store intro before first item
      if (introLines.length > 0 && result.intro === undefined) {
        result.intro = introLines.join(' ').trim();
        introLines = [];
      }
      
      result.items.push({
        title: itemMatch[2],
        description: itemMatch[3]
      });
    } else if (!foundFirstItem && trimmed) {
      // Collect intro text before first item
      introLines.push(trimmed);
    }
  }
  
  // If we have intro lines but no items yet, set intro
  if (introLines.length > 0 && result.intro === undefined) {
    result.intro = introLines.join(' ').trim();
  }
  
  return result.items.length > 0 ? result : null;
});

const handleTocClick = (index: number) => {
  emit('navigate-to-section', index);
};
</script>

