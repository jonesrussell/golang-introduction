<template>
  <div class="p-6 max-w-4xl mx-auto animate-fade-in sm:p-4">
    <!-- Loading state -->
    <div v-if="loading" class="flex flex-col items-center justify-center py-16 px-8 text-neutral-600 dark:text-neutral-400">
      <div class="w-10 h-10 border-[3px] border-neutral-200 dark:border-neutral-800 border-t-[#00ADD8] rounded-full animate-spin mb-4"></div>
      <p>Loading tutorial...</p>
    </div>

    <!-- Error state -->
    <div v-else-if="error" class="flex flex-col items-center py-16 px-8 text-red-500 text-center">
      <svg class="w-12 h-12 mb-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"/>
      </svg>
      <p>{{ error }}</p>
    </div>

    <!-- Tutorial content -->
    <div v-else-if="tutorial">
      <!-- Breadcrumb Navigation -->
      <nav class="flex items-center gap-2 mb-6 text-sm flex-wrap">
        <button
          @click="emit('home')"
          class="inline-flex items-center gap-1.5 text-neutral-600 dark:text-neutral-400 bg-transparent border-none py-1.5 px-2 -my-1.5 -mx-2 rounded-md transition-all duration-150 hover:text-[#00ADD8] hover:bg-[#e6f7fb] dark:hover:bg-neutral-800"
        >
          <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"/>
          </svg>
          Home
        </button>
        <svg class="w-4 h-4 text-neutral-500 dark:text-neutral-500 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/>
        </svg>
        <span class="text-neutral-900 dark:text-neutral-100 font-medium">{{ tutorial.title }}</span>
        <template v-if="currentSection">
          <svg class="w-4 h-4 text-neutral-500 dark:text-neutral-500 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/>
          </svg>
          <span class="text-neutral-900 dark:text-neutral-100 font-medium">{{ currentSection.title }}</span>
        </template>
      </nav>

      <!-- Tutorial header -->
      <header class="mb-8 pb-6 border-b border-neutral-200 dark:border-neutral-800">
        <h1 class="text-3xl font-bold text-neutral-900 dark:text-neutral-100 m-0 mb-4 leading-tight sm:text-2xl">{{ tutorial.title }}</h1>
        <div class="flex items-center flex-wrap gap-3">
          <span class="inline-flex items-center gap-1.5 text-base text-neutral-600 dark:text-neutral-400">
            <svg class="w-4.5 h-4.5 text-neutral-500 dark:text-neutral-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"/>
            </svg>
            {{ tutorial.duration }}
          </span>
          <span class="w-1 h-1 bg-neutral-300 dark:bg-neutral-700 rounded-full"></span>
          <span class="inline-flex items-center gap-1.5 text-base text-neutral-600 dark:text-neutral-400">
            <svg class="w-4.5 h-4.5 text-neutral-500 dark:text-neutral-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"/>
            </svg>
            {{ tutorial.difficulty }}
          </span>
          <span class="w-1 h-1 bg-neutral-300 dark:bg-neutral-700 rounded-full"></span>
          <span class="inline-flex items-center gap-1.5 text-base text-neutral-600 dark:text-neutral-400">
            <svg class="w-4.5 h-4.5 text-neutral-500 dark:text-neutral-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"/>
            </svg>
            {{ tutorial.level }}
          </span>
        </div>
      </header>

      <!-- Section content -->
      <SectionViewer
        v-if="currentSection"
        :section="currentSection"
        :section-index="currentSectionIndex"
        :total-sections="tutorial.sections.length"
        :tutorial-id="tutorial.id"
        @next="nextSection"
        @previous="previousSection"
        @complete="markComplete"
      />

      <!-- Empty state -->
      <div v-else class="flex flex-col items-center py-16 px-8 text-neutral-600 dark:text-neutral-400 text-center">
        <svg class="w-12 h-12 mb-4 text-neutral-500 dark:text-neutral-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"/>
        </svg>
        <p>No sections available for this tutorial</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch, onMounted } from 'vue';
import { useTutorial } from '../composables/useTutorial';
import { useProgressStore } from '../stores/progress';
import SectionViewer from './SectionViewer.vue';
import type { Section } from '../types/tutorial';

const props = defineProps<{
  tutorialId: string;
}>();

const emit = defineEmits<{
  (e: 'home'): void;
}>();

const { currentTutorial, loading, error, loadTutorial } = useTutorial();
const progressStore = useProgressStore();

const currentSectionIndex = ref(0);

const tutorial = computed(() => currentTutorial.value);

const currentSection = computed((): Section | null => {
  if (!tutorial.value || tutorial.value.sections.length === 0) {
    return null;
  }
  return tutorial.value.sections[currentSectionIndex.value] || null;
});

const nextSection = () => {
  if (tutorial.value && currentSectionIndex.value < tutorial.value.sections.length - 1) {
    currentSectionIndex.value++;
  }
};

const previousSection = () => {
  if (currentSectionIndex.value > 0) {
    currentSectionIndex.value--;
  }
};

const markComplete = () => {
  if (tutorial.value && currentSection.value) {
    progressStore.markSectionComplete(tutorial.value.id, currentSection.value.id);
  }
};

watch(() => props.tutorialId, async (newId) => {
  if (newId) {
    currentSectionIndex.value = 0;
    await loadTutorial(newId);
    await progressStore.loadProgress();
    if (tutorial.value && progressStore.progress?.currentTutorial === newId) {
      const lastSectionId = progressStore.progress.currentSection;
      if (lastSectionId) {
        const sectionIndex = tutorial.value.sections.findIndex((s: Section) => s.id === lastSectionId);
        if (sectionIndex >= 0) {
          currentSectionIndex.value = sectionIndex;
        }
      }
    }
  }
}, { immediate: true });

onMounted(async () => {
  await progressStore.loadFromLocalStorage();
});
</script>

