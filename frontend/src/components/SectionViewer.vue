<template>
  <div class="section-viewer">
    <div class="mb-6">
      <div class="flex justify-between items-center mb-4">
        <h2 class="text-2xl font-bold">{{ section.title }}</h2>
        <div class="text-sm text-gray-600">
          Section {{ sectionIndex + 1 }} of {{ totalSections }}
        </div>
      </div>
      
      <div class="w-full bg-gray-200 rounded-full h-2 mb-4">
        <div
          class="bg-blue-600 h-2 rounded-full transition-all"
          :style="{ width: `${((sectionIndex + 1) / totalSections) * 100}%` }"
        ></div>
      </div>
    </div>

    <div class="prose max-w-none mb-6">
      <div v-if="section.topics.length > 0" class="mb-6">
        <h3 class="text-lg font-semibold mb-2">Topics to cover:</h3>
        <ul class="list-disc list-inside space-y-1">
          <li v-for="topic in section.topics" :key="topic">{{ topic }}</li>
        </ul>
      </div>

      <div v-if="section.codeExamples.length > 0" class="mb-6">
        <div v-for="example in section.codeExamples" :key="example.id" class="mb-4">
          <CodeRunner
            v-if="example.runnable"
            :code="example.code"
            :language="example.language"
          />
          <div v-else>
            <pre class="bg-gray-100 p-4 rounded-lg overflow-x-auto"><code>{{ example.code }}</code></pre>
          </div>
        </div>
      </div>

      <div v-if="section.teachingPoints.length > 0" class="mb-6">
        <h3 class="text-lg font-semibold mb-2">Key teaching points:</h3>
        <ul class="list-disc list-inside space-y-1">
          <li v-for="point in section.teachingPoints" :key="point">{{ point }}</li>
        </ul>
      </div>
    </div>

    <div class="flex justify-between mt-8 pt-6 border-t">
      <button
        @click="$emit('previous')"
        :disabled="sectionIndex === 0"
        :class="[
          'px-4 py-2 rounded',
          sectionIndex === 0
            ? 'bg-gray-200 text-gray-400 cursor-not-allowed'
            : 'bg-blue-500 text-white hover:bg-blue-600'
        ]"
      >
        Previous Section
      </button>
      
      <button
        @click="handleComplete"
        class="px-4 py-2 rounded bg-green-500 text-white hover:bg-green-600"
      >
        Mark Complete
      </button>
      
      <button
        @click="$emit('next')"
        :disabled="sectionIndex >= totalSections - 1"
        :class="[
          'px-4 py-2 rounded',
          sectionIndex >= totalSections - 1
            ? 'bg-gray-200 text-gray-400 cursor-not-allowed'
            : 'bg-blue-500 text-white hover:bg-blue-600'
        ]"
      >
        Next Section
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useProgressStore } from '../stores/progress';
import CodeRunner from './CodeRunner.vue';
import type { Section } from '../types/tutorial';

const props = defineProps<{
  section: Section;
  sectionIndex: number;
  totalSections: number;
}>();

const emit = defineEmits<{
  (e: 'next'): void;
  (e: 'previous'): void;
  (e: 'complete'): void;
}>();

const progressStore = useProgressStore();

const handleComplete = () => {
  emit('complete');
};
</script>

<style scoped>
.section-viewer {
  background: white;
  padding: 2rem;
  border-radius: 8px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
}
</style>

