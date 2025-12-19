<template>
  <div class="section-viewer">
    <!-- Section header -->
    <div class="section-header">
      <div class="header-content">
        <h2 class="section-title">{{ section.title }}</h2>
        <div class="section-counter">
          Section {{ sectionIndex + 1 }} of {{ totalSections }}
        </div>
      </div>

      <!-- Progress bar -->
      <div class="progress-bar-container">
        <div class="progress-bar">
          <div
            class="progress-fill"
            :style="{ width: `${((sectionIndex + 1) / totalSections) * 100}%` }"
          ></div>
        </div>
      </div>
    </div>

    <!-- Section content -->
    <div class="section-content">
      <!-- Topics -->
      <div v-if="section.topics.length > 0" class="content-block">
        <h3 class="block-title">
          <svg class="block-icon" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-3 7h3m-3 4h3m-6-4h.01M9 16h.01"/>
          </svg>
          Topics to Cover
        </h3>
        <ul class="topic-list">
          <li v-for="topic in section.topics" :key="topic" class="topic-item">
            <svg class="topic-bullet" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/>
            </svg>
            {{ topic }}
          </li>
        </ul>
      </div>

      <!-- Code examples -->
      <div v-if="section.codeExamples.length > 0" class="content-block code-examples">
        <div v-for="example in section.codeExamples" :key="example.id" class="code-example">
          <CodeRunner
            :code="example.code"
            :language="example.language"
            :editable="example.runnable"
          />
        </div>
      </div>

      <!-- Teaching points -->
      <div v-if="section.teachingPoints.length > 0" class="content-block">
        <h3 class="block-title">
          <svg class="block-icon" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.663 17h4.673M12 3v1m6.364 1.636l-.707.707M21 12h-1M4 12H3m3.343-5.657l-.707-.707m2.828 9.9a5 5 0 117.072 0l-.548.547A3.374 3.374 0 0014 18.469V19a2 2 0 11-4 0v-.531c0-.895-.356-1.754-.988-2.386l-.548-.547z"/>
          </svg>
          Key Takeaways
        </h3>
        <ul class="teaching-list">
          <li v-for="point in section.teachingPoints" :key="point" class="teaching-item">
            <svg class="teaching-bullet" fill="none" viewBox="0 0 24 24" stroke="currentColor">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"/>
            </svg>
            {{ point }}
          </li>
        </ul>
      </div>
    </div>

    <!-- Navigation footer -->
    <div class="section-footer">
      <button
        @click="$emit('previous')"
        :disabled="sectionIndex === 0"
        class="nav-button nav-button-secondary"
      >
        <svg class="nav-icon" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"/>
        </svg>
        <span class="nav-text">Previous</span>
      </button>

      <div class="footer-center">
        <button
          @click="handleComplete"
          :class="['complete-button', { 'complete-button-done': isComplete }]"
        >
          <svg v-if="isComplete" class="complete-icon" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13l4 4L19 7"/>
          </svg>
          <svg v-else class="complete-icon" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"/>
          </svg>
          {{ isComplete ? 'Completed' : 'Mark Complete' }}
        </button>
      </div>

      <button
        @click="$emit('next')"
        :disabled="sectionIndex >= totalSections - 1"
        class="nav-button nav-button-primary"
      >
        <span class="nav-text">Next</span>
        <svg class="nav-icon" fill="none" viewBox="0 0 24 24" stroke="currentColor">
          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"/>
        </svg>
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { useProgressStore } from '../stores/progress';
import CodeRunner from './CodeRunner.vue';
import type { Section } from '../types/tutorial';

const props = defineProps<{
  section: Section;
  sectionIndex: number;
  totalSections: number;
  tutorialId?: string;
}>();

const emit = defineEmits<{
  (e: 'next'): void;
  (e: 'previous'): void;
  (e: 'complete'): void;
}>();

const progressStore = useProgressStore();

const isComplete = computed(() => {
  if (!props.tutorialId) return false;
  return progressStore.isSectionComplete(props.tutorialId, props.section.id);
});

const handleComplete = () => {
  emit('complete');
};
</script>

<style scoped>
.section-viewer {
  background-color: var(--color-surface);
  border: 1px solid var(--color-border);
  border-radius: var(--radius-xl);
  box-shadow: var(--shadow-sm);
  overflow: hidden;
}

/* Section header */
.section-header {
  padding: 1.5rem;
  background: linear-gradient(135deg, var(--color-primary-light), var(--color-surface));
  border-bottom: 1px solid var(--color-border);
}

.header-content {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  gap: 1rem;
  margin-bottom: 1rem;
}

.section-title {
  font-size: 1.375rem;
  font-weight: 700;
  color: var(--color-text);
  margin: 0;
  line-height: 1.3;
}

.section-counter {
  font-size: 0.8125rem;
  color: var(--color-text-muted);
  white-space: nowrap;
  padding: 0.25rem 0.75rem;
  background-color: var(--color-surface);
  border-radius: var(--radius-md);
  border: 1px solid var(--color-border);
}

.progress-bar-container {
  padding-top: 0.5rem;
}

.progress-bar {
  height: 0.5rem;
  background-color: var(--color-neutral-200);
  border-radius: 9999px;
  overflow: hidden;
}

.progress-fill {
  height: 100%;
  background: linear-gradient(90deg, var(--color-primary), var(--color-go-blue-light));
  border-radius: 9999px;
  transition: width var(--transition-slow);
}

/* Section content */
.section-content {
  padding: 1.5rem;
  display: flex;
  flex-direction: column;
  gap: 2rem;
}

.content-block {
  animation: slideUp 0.3s ease-out;
}

@keyframes slideUp {
  from {
    opacity: 0;
    transform: translateY(10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.block-title {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  font-size: 1rem;
  font-weight: 600;
  color: var(--color-text);
  margin: 0 0 1rem;
}

.block-icon {
  width: 1.25rem;
  height: 1.25rem;
  color: var(--color-primary);
}

/* Topics */
.topic-list {
  list-style: none;
  padding: 0;
  margin: 0;
  display: flex;
  flex-direction: column;
  gap: 0.625rem;
}

.topic-item {
  display: flex;
  align-items: flex-start;
  gap: 0.625rem;
  font-size: 0.9375rem;
  color: var(--color-text);
  line-height: 1.5;
}

.topic-bullet {
  width: 1rem;
  height: 1rem;
  color: var(--color-primary);
  flex-shrink: 0;
  margin-top: 0.25rem;
}

/* Code examples */
.code-examples {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}

.code-example {
  border-radius: var(--radius-lg);
  overflow: hidden;
}

/* Teaching points */
.teaching-list {
  list-style: none;
  padding: 0;
  margin: 0;
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.teaching-item {
  display: flex;
  align-items: flex-start;
  gap: 0.625rem;
  font-size: 0.9375rem;
  color: var(--color-text);
  line-height: 1.5;
  padding: 0.75rem 1rem;
  background-color: var(--color-success-light);
  border-radius: var(--radius-md);
}

.teaching-bullet {
  width: 1.125rem;
  height: 1.125rem;
  color: var(--color-success);
  flex-shrink: 0;
  margin-top: 0.125rem;
}

/* Navigation footer */
.section-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 1rem;
  padding: 1.25rem 1.5rem;
  background-color: var(--color-background);
  border-top: 1px solid var(--color-border);
}

.nav-button {
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.625rem 1rem;
  font-size: 0.875rem;
  font-weight: 500;
  border-radius: var(--radius-lg);
  border: none;
  transition: all var(--transition-fast);
}

.nav-button:disabled {
  opacity: 0.4;
  cursor: not-allowed;
}

.nav-button-secondary {
  background-color: var(--color-surface);
  color: var(--color-text);
  border: 1px solid var(--color-border);
}

.nav-button-secondary:hover:not(:disabled) {
  background-color: var(--color-neutral-100);
  border-color: var(--color-neutral-300);
}

.nav-button-primary {
  background-color: var(--color-primary);
  color: white;
}

.nav-button-primary:hover:not(:disabled) {
  background-color: var(--color-primary-hover);
}

.nav-icon {
  width: 1.125rem;
  height: 1.125rem;
}

.footer-center {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.complete-button {
  display: inline-flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.625rem 1.25rem;
  font-size: 0.875rem;
  font-weight: 600;
  border-radius: var(--radius-lg);
  border: none;
  background-color: var(--color-success);
  color: white;
  transition: all var(--transition-fast);
}

.complete-button:hover {
  background-color: #059669;
  transform: translateY(-1px);
}

.complete-button-done {
  background-color: #059669;
}

.complete-icon {
  width: 1.125rem;
  height: 1.125rem;
}

/* Responsive */
@media (max-width: 640px) {
  .section-header {
    padding: 1.25rem;
  }

  .header-content {
    flex-direction: column;
    gap: 0.5rem;
  }

  .section-title {
    font-size: 1.25rem;
  }

  .section-content {
    padding: 1.25rem;
    gap: 1.5rem;
  }

  .section-footer {
    flex-wrap: wrap;
    padding: 1rem;
  }

  .nav-button {
    padding: 0.5rem 0.75rem;
    font-size: 0.8125rem;
  }

  .nav-text {
    display: none;
  }

  .complete-button {
    order: -1;
    width: 100%;
    justify-content: center;
    margin-bottom: 0.5rem;
  }
}

/* Dark mode */
@media (prefers-color-scheme: dark) {
  .teaching-item {
    background-color: color-mix(in srgb, var(--color-success) 15%, var(--color-surface));
  }
}
</style>
