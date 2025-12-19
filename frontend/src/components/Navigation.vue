<template>
  <div class="p-4 bg-white dark:bg-neutral-900 rounded-lg shadow-sm border border-neutral-200 dark:border-neutral-800">
    <div class="flex justify-between items-center">
      <button
        type="button"
        :disabled="!hasPrevious"
        class="px-4 py-2 rounded-lg flex items-center gap-2 transition-colors"
        :class="[
          hasPrevious
            ? 'bg-[#00ADD8] text-white hover:bg-[#007D9C]'
            : 'bg-neutral-200 dark:bg-neutral-800 text-neutral-400 dark:text-neutral-600 cursor-not-allowed'
        ]" @click="$emit('previous')"
      >
        <span>←</span>
        <span>Previous</span>
      </button>

      <div class="flex items-center gap-2">
        <span class="text-sm text-neutral-600 dark:text-neutral-400">
          Section {{ currentIndex + 1 }} of {{ totalSections }}
        </span>
        <div class="flex gap-1">
          <span
            v-for="i in totalSections"
            :key="i"
            class="w-2 h-2 rounded-full" :class="[
              i - 1 < currentIndex
                ? 'bg-green-500'
                : i - 1 === currentIndex
                  ? 'bg-[#00ADD8]'
                  : 'bg-neutral-300 dark:bg-neutral-700'
            ]"
          ></span>
        </div>
      </div>

      <button
        type="button"
        :disabled="!hasNext"
        class="px-4 py-2 rounded-lg flex items-center gap-2 transition-colors"
        :class="[
          hasNext
            ? 'bg-[#00ADD8] text-white hover:bg-[#007D9C]'
            : 'bg-neutral-200 dark:bg-neutral-800 text-neutral-400 dark:text-neutral-600 cursor-not-allowed'
        ]" @click="$emit('next')"
      >
        <span>Next</span>
        <span>→</span>
      </button>
    </div>

    <div v-if="showBreadcrumbs" class="mt-4 text-sm text-neutral-600 dark:text-neutral-400">
      <nav class="flex items-center gap-2">
        <span class="cursor-pointer hover:text-[#00ADD8] transition-colors" @click="$emit('home')">Home</span>
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

/* eslint-disable no-unused-vars */
defineEmits<{
  (e: 'next'): void;
  (e: 'previous'): void;
  (e: 'home'): void;
}>();
/* eslint-enable no-unused-vars */

const hasPrevious = computed(() => props.currentIndex > 0);
const hasNext = computed(() => props.currentIndex < props.totalSections - 1);
</script>

