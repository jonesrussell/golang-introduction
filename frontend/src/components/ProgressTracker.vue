<template>
  <div class="p-4 bg-white dark:bg-neutral-900 rounded-lg shadow-sm border border-neutral-200 dark:border-neutral-800">
    <div v-if="tutorialProgress" class="space-y-4">
      <div>
        <div class="flex justify-between items-center mb-2">
          <h3 class="text-lg font-semibold">Progress</h3>
          <span class="text-sm text-gray-600">
            {{ tutorialProgress.completedCount }} / {{ tutorialProgress.totalSections }} sections
          </span>
        </div>
        <div class="w-full bg-gray-200 rounded-full h-3">
          <div
            class="bg-blue-600 h-3 rounded-full transition-all duration-300"
            :style="{ width: `${tutorialProgress.progressPercent}%` }"
          ></div>
        </div>
        <div class="text-xs text-gray-500 mt-1">
          {{ Math.round(tutorialProgress.progressPercent) }}% complete
        </div>
      </div>

      <div v-if="tutorialProgress.sectionProgress.length > 0" class="space-y-2">
        <h4 class="text-sm font-semibold">Sections:</h4>
        <div class="space-y-1">
          <div
            v-for="section in tutorialProgress.sectionProgress"
            :key="section.sectionId"
            class="flex items-center gap-2 text-sm"
          >
            <span
              :class="[
                'w-4 h-4 rounded-full flex items-center justify-center',
                section.completed
                  ? 'bg-green-500 text-white'
                  : 'bg-gray-300 text-gray-600'
              ]"
            >
              <span v-if="section.completed" class="text-xs">✓</span>
            </span>
            <span :class="section.completed ? 'text-gray-700' : 'text-gray-500'">
              Section {{ section.sectionId }}
            </span>
          </div>
        </div>
      </div>
    </div>
    
    <div v-else class="text-sm text-gray-500">
      No progress data available
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { useProgressStore } from '../stores/progress';
import type { TutorialProgress } from '../types/progress';

const props = defineProps<{
  tutorialId: string;
  totalSections: number;
}>();

const progressStore = useProgressStore();

const tutorialProgress = computed((): TutorialProgress | null => {
  if (!progressStore.progress) {
    return null;
  }
  return progressStore.getTutorialProgress(props.tutorialId, props.totalSections);
});
</script>


