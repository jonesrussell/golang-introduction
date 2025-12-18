<template>
  <div class="tutorial-list">
    <h2 class="text-2xl font-bold mb-4">Tutorials</h2>
    
    <div v-if="loading" class="text-center py-8">Loading tutorials...</div>
    <div v-else-if="error" class="text-red-500 py-8">{{ error }}</div>
    <div v-else>
      <div v-for="(tutorials, level) in tutorialsByLevel" :key="level" class="mb-8">
        <h3 class="text-xl font-semibold mb-3 text-gray-700">{{ level }}</h3>
        <div class="space-y-2">
          <div
            v-for="tutorial in tutorials"
            :key="tutorial.id"
            @click="selectTutorial(tutorial.id)"
            :class="[
              'p-4 border rounded-lg cursor-pointer transition-colors',
              currentTutorialId === tutorial.id
                ? 'bg-blue-50 border-blue-500'
                : 'bg-white border-gray-200 hover:border-gray-300'
            ]"
          >
            <div class="flex justify-between items-start">
              <div>
                <h4 class="font-semibold text-lg">{{ tutorial.title }}</h4>
                <p class="text-sm text-gray-600 mt-1">
                  {{ tutorial.duration }} • {{ tutorial.difficulty }}
                </p>
                <p class="text-xs text-gray-500 mt-1">
                  {{ tutorial.sectionCount }} sections
                </p>
              </div>
              <div v-if="progressStore.getTutorialProgress(tutorial.id, tutorial.sectionCount).progressPercent > 0">
                <div class="w-16 h-16 rounded-full border-4 border-blue-500 flex items-center justify-center"
                     :style="{ background: `conic-gradient(from 0deg, #3b82f6 0% ${progressStore.getTutorialProgress(tutorial.id, tutorial.sectionCount).progressPercent}%, #e5e7eb ${progressStore.getTutorialProgress(tutorial.id, tutorial.sectionCount).progressPercent}% 100%)` }">
                  <span class="text-xs font-semibold">
                    {{ Math.round(progressStore.getTutorialProgress(tutorial.id, tutorial.sectionCount).progressPercent) }}%
                  </span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, computed } from 'vue';
import { useTutorial } from '../composables/useTutorial';
import { useProgressStore } from '../stores/progress';

const props = defineProps<{
  currentTutorialId?: string;
}>();

const emit = defineEmits<{
  (e: 'select', tutorialId: string): void;
}>();

const { tutorials, tutorialsByLevel, loading, error, loadTutorials } = useTutorial();
const progressStore = useProgressStore();

const selectTutorial = (tutorialId: string) => {
  emit('select', tutorialId);
};

onMounted(async () => {
  await loadTutorials();
  progressStore.loadFromLocalStorage();
});
</script>

<style scoped>
.tutorial-list {
  max-height: calc(100vh - 2rem);
  overflow-y: auto;
  padding: 1rem;
}
</style>

