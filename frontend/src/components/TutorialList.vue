<template>
  <div class="flex-1 overflow-y-auto p-4">
    <!-- Loading state -->
    <div v-if="loading" class="flex flex-col items-center justify-center py-12 px-4 text-neutral-600 dark:text-neutral-400">
      <div class="w-8 h-8 border-2 border-neutral-200 dark:border-neutral-800 border-t-[#00ADD8] rounded-full animate-spin mb-4"></div>
      <p>Loading tutorials...</p>
    </div>

    <!-- Error state -->
    <div v-else-if="error" class="flex flex-col items-center py-8 px-4 text-red-500 text-center">
      <svg class="w-10 h-10 mb-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"/>
      </svg>
      <p>{{ error }}</p>
    </div>

    <!-- Tutorial list -->
    <div v-else class="flex flex-col gap-6">
      <div v-for="(tutorials, level) in tutorialsByLevel" :key="level" class="flex flex-col gap-3">
        <h3 class="m-0 px-1">
          <span
            :class="[
              'inline-flex items-center px-3.5 py-1.5 text-xs font-semibold uppercase tracking-wide rounded-md',
              getLevelClass(level as string) === 'level-beginner'
                ? 'bg-green-100 dark:bg-green-950/30 text-green-700 dark:text-green-300'
                : getLevelClass(level as string) === 'level-intermediate'
                ? 'bg-amber-100 dark:bg-amber-950/30 text-amber-700 dark:text-amber-300'
                : getLevelClass(level as string) === 'level-advanced'
                ? 'bg-red-100 dark:bg-red-950/30 text-red-700 dark:text-red-300'
                : 'bg-neutral-100 dark:bg-neutral-800 text-neutral-600 dark:text-neutral-400'
            ]"
          >{{ level }}</span>
        </h3>
        <div class="flex flex-col gap-2">
          <button
            type="button"
            v-for="tutorial in tutorials"
            :key="tutorial.id"
            @click="selectTutorial(tutorial.id)"
            :class="[
              'flex flex-col gap-2 w-full p-4 bg-white dark:bg-neutral-900 border rounded-xl text-left transition-all duration-150',
              currentTutorialId === tutorial.id
                ? 'border-[#00ADD8] bg-[#e6f7fb] dark:bg-neutral-800 shadow-[0_0_0_2px_rgba(0,173,216,0.2)]'
                : 'border-neutral-200 dark:border-neutral-800 hover:border-[#00ADD8] hover:shadow-sm'
            ]"
          >
            <div class="flex items-start justify-between gap-2">
              <h4 class="text-base font-semibold text-neutral-900 dark:text-neutral-100 m-0 leading-snug">{{ tutorial.title }}</h4>
              <span
                v-if="getTutorialStatus(tutorial.id) === 'completed'"
                class="inline-flex items-center gap-1 px-2.5 py-1 text-xs font-semibold rounded-sm bg-green-100 dark:bg-green-950/30 text-green-700 dark:text-green-300 whitespace-nowrap flex-shrink-0"
              >
                <svg class="w-3 h-3" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"/>
                </svg>
                Done
              </span>
              <span
                v-else-if="getTutorialStatus(tutorial.id) === 'in-progress'"
                class="inline-flex items-center gap-1 px-2.5 py-1 text-xs font-semibold rounded-sm bg-[#e6f7fb] dark:bg-neutral-800 text-[#007D9C] dark:text-[#5DC9E2] whitespace-nowrap flex-shrink-0"
              >
                In Progress
              </span>
            </div>

            <div class="flex flex-wrap gap-3">
              <span class="inline-flex items-center gap-1.5 text-sm text-neutral-600 dark:text-neutral-400">
                <svg class="w-4 h-4 text-neutral-500 dark:text-neutral-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"/>
                </svg>
                {{ tutorial.duration }}
              </span>
              <span class="inline-flex items-center gap-1.5 text-sm text-neutral-600 dark:text-neutral-400">
                <svg class="w-4 h-4 text-neutral-500 dark:text-neutral-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"/>
                </svg>
                {{ tutorial.sectionCount }} sections
              </span>
            </div>

            <!-- Progress bar -->
            <div class="flex items-center gap-3 mt-1">
              <div class="flex-1 h-1.5 bg-neutral-200 dark:bg-neutral-800 rounded-full overflow-hidden">
                <div
                  class="h-full bg-[#00ADD8] rounded-full transition-all duration-300"
                  :style="{ width: `${getTutorialProgress(tutorial.id).progressPercent}%` }"
                ></div>
              </div>
              <span class="text-xs text-neutral-500 dark:text-neutral-500 whitespace-nowrap">
                {{ getTutorialProgress(tutorial.id).completedCount }} / {{ tutorial.sectionCount }}
              </span>
            </div>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { useTutorial } from '../composables/useTutorial';
import { useProgressStore } from '../stores/progress';

defineProps<{
  currentTutorialId?: string;
}>();

const router = useRouter();
const { tutorials, tutorialsByLevel, loading, error, loadTutorials } = useTutorial();
const progressStore = useProgressStore();

const selectTutorial = (tutorialId: string) => {
  router.push({ name: 'tutorial', params: { id: tutorialId } });
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

const getLevelClass = (level: string): string => {
  const levelLower = level.toLowerCase();
  if (levelLower.includes('beginner')) return 'level-beginner';
  if (levelLower.includes('intermediate')) return 'level-intermediate';
  if (levelLower.includes('advanced')) return 'level-advanced';
  return 'level-default';
};

onMounted(async () => {
  await loadTutorials();
  progressStore.loadFromLocalStorage();
});
</script>

