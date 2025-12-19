<template>
  <div class="tutorial-list">
    <!-- Loading state -->
    <div v-if="loading" class="loading-state">
      <div class="loading-spinner"></div>
      <p>Loading tutorials...</p>
    </div>

    <!-- Error state -->
    <div v-else-if="error" class="error-state">
      <svg class="error-icon" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z"/>
      </svg>
      <p>{{ error }}</p>
    </div>

    <!-- Tutorial list -->
    <div v-else class="tutorial-groups">
      <div v-for="(tutorials, level) in tutorialsByLevel" :key="level" class="tutorial-group">
        <h3 class="group-title">
          <span class="level-badge" :class="getLevelClass(level as string)">{{ level }}</span>
        </h3>
        <div class="tutorial-cards">
          <button
            v-for="tutorial in tutorials"
            :key="tutorial.id"
            @click="selectTutorial(tutorial.id)"
            :class="['tutorial-card', { 'tutorial-card-active': currentTutorialId === tutorial.id }]"
          >
            <div class="card-header">
              <h4 class="card-title">{{ tutorial.title }}</h4>
              <span
                v-if="getTutorialStatus(tutorial.id) === 'completed'"
                class="status-badge status-completed"
              >
                <svg class="status-icon" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"/>
                </svg>
                Done
              </span>
              <span
                v-else-if="getTutorialStatus(tutorial.id) === 'in-progress'"
                class="status-badge status-progress"
              >
                In Progress
              </span>
            </div>

            <div class="card-meta">
              <span class="meta-item">
                <svg class="meta-icon" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z"/>
                </svg>
                {{ tutorial.duration }}
              </span>
              <span class="meta-item">
                <svg class="meta-icon" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"/>
                </svg>
                {{ tutorial.sectionCount }} sections
              </span>
            </div>

            <!-- Progress bar -->
            <div class="progress-container">
              <div class="progress-bar">
                <div
                  class="progress-fill"
                  :style="{ width: `${getTutorialProgress(tutorial.id).progressPercent}%` }"
                ></div>
              </div>
              <span class="progress-text">
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
import { useTutorial } from '../composables/useTutorial';
import { useProgressStore } from '../stores/progress';

defineProps<{
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

<style scoped>
.tutorial-list {
  flex: 1;
  overflow-y: auto;
  padding: 1rem;
}

/* Loading state */
.loading-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 3rem 1rem;
  color: var(--color-text-muted);
}

.loading-spinner {
  width: 2rem;
  height: 2rem;
  border: 2px solid var(--color-border);
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
  padding: 2rem 1rem;
  color: var(--color-error);
  text-align: center;
}

.error-icon {
  width: 2.5rem;
  height: 2.5rem;
  margin-bottom: 0.75rem;
}

/* Tutorial groups */
.tutorial-groups {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}

.tutorial-group {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.group-title {
  margin: 0;
  padding: 0 0.25rem;
}

.level-badge {
  display: inline-flex;
  align-items: center;
  padding: 0.375rem 0.875rem;
  font-size: var(--text-xs);
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.025em;
  border-radius: var(--radius-md);
}

.level-beginner {
  background-color: var(--color-success-light);
  color: #059669;
}

.level-intermediate {
  background-color: var(--color-warning-light);
  color: #d97706;
}

.level-advanced {
  background-color: var(--color-error-light);
  color: #dc2626;
}

.level-default {
  background-color: var(--color-neutral-100);
  color: var(--color-text-muted);
}

/* Tutorial cards */
.tutorial-cards {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.tutorial-card {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
  width: 100%;
  padding: 1rem;
  background-color: var(--color-surface);
  border: 1px solid var(--color-border);
  border-radius: var(--radius-lg);
  text-align: left;
  transition: all var(--transition-fast);
}

.tutorial-card:hover {
  border-color: var(--color-primary);
  box-shadow: var(--shadow-sm);
}

.tutorial-card-active {
  border-color: var(--color-primary);
  background-color: var(--color-primary-light);
  box-shadow: 0 0 0 2px color-mix(in srgb, var(--color-primary) 20%, transparent);
}

.card-header {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 0.5rem;
}

.card-title {
  font-size: var(--text-base);
  font-weight: 600;
  color: var(--color-text);
  margin: 0;
  line-height: 1.4;
}

/* Status badges */
.status-badge {
  display: inline-flex;
  align-items: center;
  gap: 0.25rem;
  padding: 0.25rem 0.625rem;
  font-size: var(--text-xs);
  font-weight: 600;
  border-radius: var(--radius-sm);
  white-space: nowrap;
  flex-shrink: 0;
}

.status-completed {
  background-color: var(--color-success-light);
  color: #059669;
}

.status-progress {
  background-color: var(--color-primary-light);
  color: var(--color-primary-hover);
}

.status-icon {
  width: 0.75rem;
  height: 0.75rem;
}

/* Meta info */
.card-meta {
  display: flex;
  flex-wrap: wrap;
  gap: 0.75rem;
}

.meta-item {
  display: inline-flex;
  align-items: center;
  gap: 0.375rem;
  font-size: var(--text-sm);
  color: var(--color-text-muted);
}

.meta-icon {
  width: 1rem;
  height: 1rem;
  color: var(--color-text-subtle);
}

/* Progress bar */
.progress-container {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  margin-top: 0.25rem;
}

.progress-bar {
  flex: 1;
  height: 0.375rem;
  background-color: var(--color-neutral-200);
  border-radius: 9999px;
  overflow: hidden;
}

.progress-fill {
  height: 100%;
  background-color: var(--color-primary);
  border-radius: 9999px;
  transition: width var(--transition-slow);
}

.progress-text {
  font-size: var(--text-xs);
  color: var(--color-text-subtle);
  white-space: nowrap;
}

/* Dark mode adjustments */
@media (prefers-color-scheme: dark) {
  .level-beginner {
    color: #34d399;
  }

  .level-intermediate {
    color: #fbbf24;
  }

  .level-advanced {
    color: #f87171;
  }

  .status-completed {
    color: #34d399;
  }
}
</style>
