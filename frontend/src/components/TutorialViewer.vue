<template>
  <div class="tutorial-viewer">
    <div v-if="loading" class="text-center py-8">Loading tutorial...</div>
    <div v-else-if="error" class="text-red-500 py-8">{{ error }}</div>
    <div v-else-if="tutorial">
      <div class="mb-6">
        <h1 class="text-3xl font-bold mb-2">{{ tutorial.title }}</h1>
        <div class="flex gap-4 text-sm text-gray-600">
          <span>{{ tutorial.duration }}</span>
          <span>•</span>
          <span>{{ tutorial.difficulty }}</span>
          <span>•</span>
          <span>{{ tutorial.level }}</span>
        </div>
      </div>

      <SectionViewer
        v-if="currentSection"
        :section="currentSection"
        :section-index="currentSectionIndex"
        :total-sections="tutorial.sections.length"
        @next="nextSection"
        @previous="previousSection"
        @complete="markComplete"
      />

      <div v-else class="text-center py-8 text-gray-500">
        No sections available
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
    await loadTutorial(newId);
    // Load progress for this tutorial
    await progressStore.loadProgress();
    // Set current section based on progress
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
  padding: 2rem;
  max-width: 1200px;
  margin: 0 auto;
}
</style>

