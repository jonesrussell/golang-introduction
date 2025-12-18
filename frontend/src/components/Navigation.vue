<template>
  <div class="navigation">
    <div class="flex justify-between items-center">
      <button
        @click="$emit('previous')"
        :disabled="!hasPrevious"
        :class="[
          'px-4 py-2 rounded flex items-center gap-2',
          hasPrevious
            ? 'bg-blue-500 text-white hover:bg-blue-600'
            : 'bg-gray-200 text-gray-400 cursor-not-allowed'
        ]"
      >
        <span>←</span>
        <span>Previous</span>
      </button>

      <div class="flex items-center gap-2">
        <span class="text-sm text-gray-600">
          Section {{ currentIndex + 1 }} of {{ totalSections }}
        </span>
        <div class="flex gap-1">
          <span
            v-for="i in totalSections"
            :key="i"
            :class="[
              'w-2 h-2 rounded-full',
              i - 1 < currentIndex
                ? 'bg-green-500'
                : i - 1 === currentIndex
                ? 'bg-blue-500'
                : 'bg-gray-300'
            ]"
          ></span>
        </div>
      </div>

      <button
        @click="$emit('next')"
        :disabled="!hasNext"
        :class="[
          'px-4 py-2 rounded flex items-center gap-2',
          hasNext
            ? 'bg-blue-500 text-white hover:bg-blue-600'
            : 'bg-gray-200 text-gray-400 cursor-not-allowed'
        ]"
      >
        <span>Next</span>
        <span>→</span>
      </button>
    </div>

    <div v-if="showBreadcrumbs" class="mt-4 text-sm text-gray-600">
      <nav class="flex items-center gap-2">
        <span @click="$emit('home')" class="cursor-pointer hover:text-blue-600">Home</span>
        <span>/</span>
        <span>{{ tutorialTitle }}</span>
        <span v-if="sectionTitle">/</span>
        <span v-if="sectionTitle">{{ sectionTitle }}</span>
      </nav>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';

const props = defineProps<{
  currentIndex: number;
  totalSections: number;
  tutorialTitle?: string;
  sectionTitle?: string;
  showBreadcrumbs?: boolean;
}>();

const emit = defineEmits<{
  (e: 'next'): void;
  (e: 'previous'): void;
  (e: 'home'): void;
}>();

const hasPrevious = computed(() => props.currentIndex > 0);
const hasNext = computed(() => props.currentIndex < props.totalSections - 1);
</script>

<style scoped>
.navigation {
  padding: 1rem;
  background: white;
  border-radius: 8px;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
}
</style>
