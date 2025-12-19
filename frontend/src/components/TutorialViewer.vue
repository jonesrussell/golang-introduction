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
        <router-link
          to="/"
          class="inline-flex items-center gap-1.5 text-neutral-600 dark:text-neutral-400 bg-transparent border-none py-1.5 px-2 -my-1.5 -mx-2 rounded-md transition-all duration-150 hover:text-[#00ADD8] hover:bg-[#e6f7fb] dark:hover:bg-neutral-800"
        >
          <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"/>
          </svg>
          Home
        </router-link>
        <svg class="w-4 h-4 text-neutral-500 dark:text-neutral-500 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/>
        </svg>
        <router-link
          :to="{ name: 'tutorial', params: { id: tutorial.id } }"
          class="text-neutral-900 dark:text-neutral-100 font-medium hover:text-[#00ADD8] transition-colors"
        >
          {{ tutorial.title }}
        </router-link>
        <template v-if="currentSection">
          <svg class="w-4 h-4 text-neutral-500 dark:text-neutral-500 flex-shrink-0" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/>
          </svg>
          <span class="text-neutral-900 dark:text-neutral-100 font-medium">{{ currentSection.title }}</span>
        </template>
      </nav>

      <!-- Tutorial header -->
      <header class="mb-8 pb-6 border-b border-neutral-200 dark:border-neutral-800">
        <div class="flex items-start justify-between gap-4 mb-4">
          <h1 class="text-3xl font-bold text-neutral-900 dark:text-neutral-100 m-0 leading-tight sm:text-2xl">{{ tutorial.title }}</h1>
          
          <!-- Instructor Mode Toggle -->
          <button
            type="button"
            @click="toggleInstructorMode"
            :class="[
              'inline-flex items-center gap-2 px-3 py-2 text-sm font-medium rounded-lg border transition-all duration-150',
              instructorMode
                ? 'bg-amber-100 dark:bg-amber-900/30 border-amber-300 dark:border-amber-700 text-amber-800 dark:text-amber-200'
                : 'bg-neutral-100 dark:bg-neutral-800 border-neutral-200 dark:border-neutral-700 text-neutral-600 dark:text-neutral-400 hover:border-amber-300 dark:hover:border-amber-700'
            ]"
            title="Toggle instructor mode"
          >
            <svg class="w-4 h-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 10l4.553-2.276A1 1 0 0121 8.618v6.764a1 1 0 01-1.447.894L15 14M5 18h8a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v8a2 2 0 002 2z"/>
            </svg>
            <span class="hidden sm:inline">{{ instructorMode ? 'Instructor Mode' : 'Instructor' }}</span>
          </button>
        </div>
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
        :instructor-mode="instructorMode"
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
import { useRouter } from 'vue-router';
import { useTutorial } from '../composables/useTutorial';
import { useProgressStore } from '../stores/progress';
import SectionViewer from './SectionViewer.vue';
import type { Section } from '../types/tutorial';

const props = defineProps<{
  tutorialId: string;
  sectionIndex?: number;
}>();

const router = useRouter();
const { currentTutorial, loading, error, loadTutorial } = useTutorial();
const progressStore = useProgressStore();

const currentSectionIndex = ref(0);

// Instructor mode state - persisted in localStorage
const instructorMode = ref(localStorage.getItem('instructor-mode') === 'true');

const toggleInstructorMode = async () => {
  instructorMode.value = !instructorMode.value;
  localStorage.setItem('instructor-mode', instructorMode.value.toString());
  
  // Reload tutorial with/without instructor notes
  if (props.tutorialId) {
    await loadTutorial(props.tutorialId, true, instructorMode.value);
  }
};

const tutorial = computed(() => currentTutorial.value);

const currentSection = computed((): Section | null => {
  if (!tutorial.value || tutorial.value.sections.length === 0) {
    return null;
  }
  return tutorial.value.sections[currentSectionIndex.value] || null;
});

const navigateToSection = (index: number) => {
  if (!tutorial.value) return;
  
  // Clamp index to valid range
  const clampedIndex = Math.max(0, Math.min(index, tutorial.value.sections.length - 1));
  
  // Save progress
  const section = tutorial.value.sections[clampedIndex];
  if (section) {
    progressStore.setCurrentSection(tutorial.value.id, section.id);
  }
  
  // Navigate to the section route (1-based in URL)
  router.push({
    name: 'tutorial-section',
    params: {
      id: tutorial.value.id,
      sectionIndex: (clampedIndex + 1).toString()
    }
  });
};

const nextSection = () => {
  if (tutorial.value && currentSectionIndex.value < tutorial.value.sections.length - 1) {
    navigateToSection(currentSectionIndex.value + 1);
  }
};

const previousSection = () => {
  if (currentSectionIndex.value > 0) {
    navigateToSection(currentSectionIndex.value - 1);
  }
};

const markComplete = () => {
  if (tutorial.value && currentSection.value) {
    progressStore.markSectionComplete(tutorial.value.id, currentSection.value.id);
  }
};

// Watch for tutorial ID changes
watch(() => props.tutorialId, async (newId) => {
  if (newId) {
    await loadTutorial(newId, false, instructorMode.value);
    await progressStore.loadProgress();
    
    // Set section index from prop if provided, otherwise use saved progress or default to 0
    if (props.sectionIndex !== undefined && props.sectionIndex >= 0) {
      currentSectionIndex.value = props.sectionIndex;
    } else if (tutorial.value && progressStore.progress?.currentTutorial === newId) {
      const lastSectionId = progressStore.progress.currentSection;
      if (lastSectionId) {
        const sectionIndex = tutorial.value.sections.findIndex((s: Section) => s.id === lastSectionId);
        if (sectionIndex >= 0) {
          currentSectionIndex.value = sectionIndex;
        }
      }
    } else {
      currentSectionIndex.value = 0;
    }
  }
}, { immediate: true });

// Watch for section index prop changes (when navigating via route)
watch(() => props.sectionIndex, (newIndex) => {
  if (newIndex !== undefined && newIndex >= 0 && tutorial.value) {
    // Clamp to valid range
    const clampedIndex = Math.max(0, Math.min(newIndex, tutorial.value.sections.length - 1));
    currentSectionIndex.value = clampedIndex;
    
    // Update progress to reflect current section
    const section = tutorial.value.sections[clampedIndex];
    if (section) {
      progressStore.setCurrentSection(tutorial.value.id, section.id);
    }
  }
});

onMounted(async () => {
  await progressStore.loadFromLocalStorage();
});
</script>

