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
      <div v-for="(tutorialsInLevel, level) in tutorialsByLevel" :key="level" class="flex flex-col gap-3">
        <h3 class="m-0 px-1">
          <span
            class="inline-flex items-center px-3.5 py-1.5 text-xs font-semibold uppercase tracking-wide rounded-md" :class="[
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
          <div
            v-for="tutorial in tutorialsInLevel"
            :key="tutorial.id"
            class="flex flex-col gap-2"
          >
            <button
              type="button"
              class="flex flex-col gap-2 w-full p-4 bg-white dark:bg-neutral-900 border rounded-xl text-left transition-all duration-150" :class="[
                currentTutorialId === tutorial.id
                  ? 'border-[#00ADD8] bg-[#e6f7fb] dark:bg-neutral-800 shadow-[0_0_0_2px_rgba(0,173,216,0.2)]'
                  : 'border-neutral-200 dark:border-neutral-800 hover:border-[#00ADD8] hover:shadow-sm'
              ]"
              @click="selectTutorial(tutorial.id)"
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
            
            <!-- Table of Contents submenu (only shown when tutorial is active) -->
            <div v-if="currentTutorialId === tutorial.id && parsedTableOfContents(tutorial.id)" class="ml-4 mt-2 border-l-2 border-neutral-200 dark:border-neutral-800 pl-4">
              <ol class="list-none m-0 p-0 flex flex-col gap-1.5">
                <li
                  v-for="(item, index) in parsedTableOfContents(tutorial.id)!.items"
                  :key="index"
                  class="cursor-pointer transition-all duration-150 rounded-md px-2 py-1.5 -ml-2"
                  :class="{
                    'text-[#00ADD8] font-semibold bg-blue-100 dark:bg-blue-900/30': currentSectionIndex === index,
                    'text-neutral-700 dark:text-neutral-300 hover:text-[#00ADD8] hover:bg-blue-50 dark:hover:bg-blue-950/20': currentSectionIndex !== index
                  }"
                  @click="navigateToSection(tutorial.id, index)"
                >
                  <div class="flex items-start gap-2 text-sm">
                    <span class="flex-shrink-0 w-5 text-[#00ADD8] font-semibold">{{ index + 1 }}.</span>
                    <div class="flex-1">
                      <strong class="font-semibold block">{{ item.title }}</strong>
                      <span class="text-xs text-neutral-600 dark:text-neutral-400 mt-0.5 block">{{ item.description }}</span>
                    </div>
                  </div>
                </li>
              </ol>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, ref, computed, watch } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import { useTutorial } from '../composables/useTutorial';
import { useProgressStore } from '../stores/progress';
import { tutorialApi } from '../services/api';

interface TocItem {
  title: string;
  description: string;
}

interface ParsedToc {
  intro?: string;
  items: TocItem[];
}

const props = defineProps<{
  currentTutorialId?: string;
}>();

const router = useRouter();
const route = useRoute();
const { tutorials, tutorialsByLevel, loading, error, loadTutorials } = useTutorial();
const progressStore = useProgressStore();

// Store loaded tutorials with table of contents
const tutorialsWithToc = ref<Map<string, ParsedToc | null>>(new Map());

// Get current section index from route
const currentSectionIndex = computed(() => {
  if (route.name === 'tutorial-section' && typeof route.params.sectionIndex === 'string') {
    const index = parseInt(route.params.sectionIndex, 10);
    return isNaN(index) ? -1 : index - 1; // Convert from 1-based to 0-based
  }
  return -1;
});

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

// Parse table of contents from tutorial content
const parseTableOfContents = (content: string): ParsedToc | null => {
  if (!content) return null;
  
  const lines = content.trim().split('\n');
  const result: ParsedToc = {
    items: []
  };
  
  let introLines: string[] = [];
  let foundFirstItem = false;
  
  for (const line of lines) {
    const trimmed = line.trim();
    
    // Skip empty lines at the start
    if (!trimmed && !foundFirstItem) continue;
    
    // Check for numbered list item (1. **Title** - description)
    const itemMatch = trimmed.match(/^(\d+)\.\s+\*\*(.+?)\*\*\s+-\s+(.+)$/);
    if (itemMatch) {
      foundFirstItem = true;
      // Store intro before first item
      if (introLines.length > 0 && result.intro === undefined) {
        result.intro = introLines.join(' ').trim();
        introLines = [];
      }
      
      result.items.push({
        title: itemMatch[2] ?? '',
        description: itemMatch[3] ?? ''
      });
    } else if (!foundFirstItem && trimmed) {
      // Collect intro text before first item
      introLines.push(trimmed);
    }
  }
  
  // If we have intro lines but no items yet, set intro
  if (introLines.length > 0 && result.intro === undefined) {
    result.intro = introLines.join(' ').trim();
  }
  
  return result.items.length > 0 ? result : null;
};

// Get parsed table of contents for a tutorial
const parsedTableOfContents = (tutorialId: string): ParsedToc | null => {
  return tutorialsWithToc.value.get(tutorialId) ?? null;
};

// Load tutorial with table of contents when it becomes active
watch(() => props.currentTutorialId, async (newId) => {
  if (newId && !tutorialsWithToc.value.has(newId)) {
    try {
      const tutorial = await tutorialApi.getTutorial(newId, false);
      if (tutorial?.tableOfContents) {
        tutorialsWithToc.value.set(newId, parseTableOfContents(tutorial.tableOfContents));
      } else {
        tutorialsWithToc.value.set(newId, null);
      }
    } catch (err) {
      console.error('Failed to load tutorial for TOC:', err);
      tutorialsWithToc.value.set(newId, null);
    }
  }
}, { immediate: true });

const navigateToSection = (tutorialId: string, sectionIndex: number) => {
  router.push({
    name: 'tutorial-section',
    params: {
      id: tutorialId,
      sectionIndex: (sectionIndex + 1).toString() // Convert to 1-based for URL
    }
  });
};

onMounted(async () => {
  await loadTutorials();
  progressStore.loadFromLocalStorage();
  
  // Load TOC for current tutorial if active
  if (props.currentTutorialId) {
    try {
      const tutorial = await tutorialApi.getTutorial(props.currentTutorialId, false);
      if (tutorial?.tableOfContents) {
        tutorialsWithToc.value.set(props.currentTutorialId, parseTableOfContents(tutorial.tableOfContents));
      }
    } catch (err) {
      console.error('Failed to load tutorial for TOC:', err);
    }
  }
});
</script>

