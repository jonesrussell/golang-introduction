<template>
  <div class="tutorial-viewer">
    <!-- Loading state -->
    <div v-if="loading" class="loading-state">
      <div class="loading-spinner"></div>
      <p>Loading tutorial...</p>
    </div>

    <!-- Error state -->
    <div v-else-if="error" class="error-state">
      <svg class="error-icon" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"/>
      </svg>
      <p>{{ error }}</p>
    </div>

    <!-- Tutorial content -->
    <div v-else-if="tutorial" class="tutorial-content">
      <!-- Breadcrumb Navigation -->
      <nav class="breadcrumb">
        <button @click="emit('home')" class="breadcrumb-link">
          <svg class="breadcrumb-icon" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"/>
          </svg>
          Home
        </button>
        <svg class="breadcrumb-separator" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/>
        </svg>
        <span class="breadcrumb-current">{{ tutorial.title }}</span>
        <template v-if="currentSection">
          <svg class="breadcrumb-separator" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/>
          </svg>
          <span class="breadcrumb-current">{{ currentSection.title }}</span>
        </template>
      </nav>

      <!-- Tutorial header -->
      <header class="tutorial-header">
        <h1 class="tutorial-title">{{ tutorial.title }}</h1>
        <div class="tutorial-meta">
          <span class="meta-item">
            <svg class="meta-icon" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"/>
            </svg>
            {{ tutorial.duration }}
          </span>
          <span class="meta-divider"></span>
          <span class="meta-item">
            <svg class="meta-icon" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"/>
            </svg>
            {{ tutorial.difficulty }}
          </span>
          <span class="meta-divider"></span>
          <span class="meta-item">
            <svg class="meta-icon" fill="none" viewBox="0 0 24 24" stroke="currentColor">
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
      <div v-else class="empty-state">
        <svg class="empty-icon" fill="none" viewBox="0 0 24 24" stroke="currentColor">
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

<style scoped>
.tutorial-viewer {
  padding: 1.5rem;
  max-width: 900px;
  margin: 0 auto;
  animation: fadeIn 0.3s ease-out;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

/* Loading state */
.loading-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 4rem 2rem;
  color: var(--color-text-muted);
}

.loading-spinner {
  width: 2.5rem;
  height: 2.5rem;
  border: 3px solid var(--color-border);
  border-top-color: var(--color-primary);
  border-radius: 50%;
  animation: spin 0.8s linear infinite;
  margin-bottom: 1rem;
}

@keyframes spin {
  to { transform: rotate(360deg); }
}

/* Error state */
.error-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 4rem 2rem;
  color: var(--color-error);
  text-align: center;
}

.error-icon {
  width: 3rem;
  height: 3rem;
  margin-bottom: 1rem;
}

/* Breadcrumb */
.breadcrumb {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  margin-bottom: 1.5rem;
  font-size: 0.875rem;
  flex-wrap: wrap;
}

.breadcrumb-link {
  display: inline-flex;
  align-items: center;
  gap: 0.375rem;
  color: var(--color-text-muted);
  background: none;
  border: none;
  padding: 0.375rem 0.5rem;
  margin: -0.375rem -0.5rem;
  border-radius: var(--radius-md);
  transition: all var(--transition-fast);
}

.breadcrumb-link:hover {
  color: var(--color-primary);
  background-color: var(--color-primary-light);
}

.breadcrumb-icon {
  width: 1rem;
  height: 1rem;
}

.breadcrumb-separator {
  width: 1rem;
  height: 1rem;
  color: var(--color-text-subtle);
  flex-shrink: 0;
}

.breadcrumb-current {
  color: var(--color-text);
  font-weight: 500;
}

/* Tutorial header */
.tutorial-header {
  margin-bottom: 2rem;
  padding-bottom: 1.5rem;
  border-bottom: 1px solid var(--color-border);
}

.tutorial-title {
  font-size: 1.875rem;
  font-weight: 700;
  color: var(--color-text);
  margin: 0 0 1rem;
  line-height: 1.3;
}

.tutorial-meta {
  display: flex;
  align-items: center;
  flex-wrap: wrap;
  gap: 0.75rem;
}

.meta-item {
  display: inline-flex;
  align-items: center;
  gap: 0.375rem;
  font-size: 0.875rem;
  color: var(--color-text-muted);
}

.meta-icon {
  width: 1rem;
  height: 1rem;
  color: var(--color-text-subtle);
}

.meta-divider {
  width: 4px;
  height: 4px;
  background-color: var(--color-neutral-300);
  border-radius: 50%;
}

/* Empty state */
.empty-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 4rem 2rem;
  color: var(--color-text-muted);
  text-align: center;
}

.empty-icon {
  width: 3rem;
  height: 3rem;
  margin-bottom: 1rem;
  color: var(--color-text-subtle);
}

/* Responsive */
@media (max-width: 640px) {
  .tutorial-viewer {
    padding: 1rem;
  }

  .tutorial-title {
    font-size: 1.5rem;
  }

  .breadcrumb {
    font-size: 0.8125rem;
  }
}
</style>
