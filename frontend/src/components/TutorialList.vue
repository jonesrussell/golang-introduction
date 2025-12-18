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
              <div class="flex-1">
                <div class="flex items-center gap-2 mb-1">
                  <h4 class="font-semibold text-lg">{{ tutorial.title }}</h4>
                  <span
                    v-if="getTutorialStatus(tutorial.id) === 'completed'"
                    class="text-xs bg-green-100 text-green-800 px-2 py-0.5 rounded-full font-semibold"
                  >
                    ✓ Completed
                  </span>
                  <span
                    v-else-if="getTutorialStatus(tutorial.id) === 'in-progress'"
                    class="text-xs bg-blue-100 text-blue-800 px-2 py-0.5 rounded-full font-semibold"
                  >
                    In Progress
                  </span>
                </div>
                <p class="text-sm text-gray-600 mt-1">
                  {{ tutorial.duration }} • {{ tutorial.difficulty }}
                </p>
                <p class="text-xs text-gray-500 mt-1">
                  {{ tutorial.sectionCount }} sections
                </p>
                <!-- Progress bar -->
                <div class="mt-2">
                  <div class="w-full bg-gray-200 rounded-full h-1.5">
                    <div
                      class="bg-blue-600 h-1.5 rounded-full transition-all duration-300"
                      :style="{ width: `${getTutorialProgress(tutorial.id).progressPercent}%` }"
                    ></div>
                  </div>
                  <p class="text-xs text-gray-500 mt-1">
                    {{ getTutorialProgress(tutorial.id).completedCount }} / {{ tutorial.sectionCount }} sections completed
                  </p>
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

const getTutorialProgress = (tutorialId: string) => {
  const tutorial = tutorials.value.find(t => t.id === tutorialId);
  if (!tutorial) {
    return { completedCount: 0, progressPercent: 0 };
  }
  return progressStore.getTutorialProgress(tutorialId, tutorial.sectionCount);
};

const getTutorialStatus = (tutorialId: string): 'not-started' | 'in-progress' | 'completed' => {
  const progress = getTutorialProgress(tutorialId);
  if (progress.progressPercent === 0) {
    return 'not-started';
  } else if (progress.progressPercent === 100) {
    return 'completed';
  } else {
    return 'in-progress';
  }
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

