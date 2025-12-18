<template>
  <div class="exercise-view">
    <div v-if="loading" class="text-center py-8">Loading exercises...</div>
    <div v-else-if="error" class="text-red-500 py-8">{{ error }}</div>
    <div v-else-if="exercises.length === 0" class="text-center py-8 text-gray-500">
      No exercises available for this tutorial yet.
    </div>
    <div v-else>
      <div v-for="exercise in exercises" :key="exercise.id" class="mb-8 p-6 bg-white rounded-lg shadow">
        <div class="mb-4">
          <h3 class="text-xl font-bold mb-2">{{ exercise.title }}</h3>
          <span :class="[
            'px-2 py-1 rounded text-xs font-semibold',
            exercise.difficulty === 'Easy' ? 'bg-green-100 text-green-800' :
            exercise.difficulty === 'Medium' ? 'bg-yellow-100 text-yellow-800' :
            'bg-red-100 text-red-800'
          ]">
            {{ exercise.difficulty }}
          </span>
        </div>
        
        <div class="mb-4">
          <p class="text-gray-700 whitespace-pre-wrap">{{ exercise.description }}</p>
        </div>

        <div v-if="exercise.starterCode" class="mb-4">
          <h4 class="font-semibold mb-2">Starter Code:</h4>
          <pre class="bg-gray-100 p-4 rounded-lg overflow-x-auto"><code>{{ exercise.starterCode }}</code></pre>
        </div>

        <div class="mb-4">
          <h4 class="font-semibold mb-2">Your Solution:</h4>
          <textarea
            v-model="solutions[exercise.id]"
            class="w-full h-48 p-3 border border-gray-300 rounded-lg font-mono text-sm"
            :placeholder="exercise.starterCode || 'Write your solution here...'"
          ></textarea>
        </div>

        <div class="flex gap-2">
          <button
            @click="executeSolution(exercise.id)"
            :disabled="executing"
            :class="[
              'px-4 py-2 rounded',
              executing
                ? 'bg-gray-400 text-gray-600 cursor-not-allowed'
                : 'bg-blue-500 text-white hover:bg-blue-600'
            ]"
          >
            {{ executing ? 'Running...' : 'Run Solution' }}
          </button>
          
          <button
            @click="checkSolution(exercise.id)"
            class="px-4 py-2 rounded bg-green-500 text-white hover:bg-green-600"
          >
            Check Solution
          </button>
        </div>

        <div v-if="executionResults[exercise.id]" class="mt-4">
          <div v-if="executionResults[exercise.id]?.output" class="bg-green-50 border border-green-200 rounded p-3 mb-2">
            <div class="text-sm font-semibold text-green-800 mb-1">Output:</div>
            <pre class="text-sm text-green-700 whitespace-pre-wrap">{{ executionResults[exercise.id]?.output }}</pre>
          </div>
          
          <div v-if="executionResults[exercise.id]?.error" class="bg-red-50 border border-red-200 rounded p-3">
            <div class="text-sm font-semibold text-red-800 mb-1">Error:</div>
            <pre class="text-sm text-red-700 whitespace-pre-wrap">{{ executionResults[exercise.id]?.error }}</pre>
          </div>
        </div>

        <div v-if="exercise.hints && exercise.hints.length > 0" class="mt-4">
          <details>
            <summary class="cursor-pointer text-sm text-blue-600 hover:text-blue-800">Show Hints</summary>
            <ul class="mt-2 list-disc list-inside space-y-1 text-sm text-gray-700">
              <li v-for="(hint, index) in exercise.hints" :key="index">{{ hint }}</li>
            </ul>
          </details>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { useCodeExecution } from '../composables/useCodeExecution';
import { exerciseApi } from '../services/api';
import type { ExecutionResult } from '../types/progress';
import type { Exercise } from '../types/tutorial';

const props = defineProps<{
  tutorialId: string;
}>();

const exercises = ref<Exercise[]>([]);
const loading = ref(false);
const error = ref<string | null>(null);
const solutions = ref<Record<string, string>>({});
const executionResults = ref<Record<string, ExecutionResult>>({});
const { executing, result, executeCode } = useCodeExecution();

const loadExercises = async () => {
  loading.value = true;
  error.value = null;
  try {
    exercises.value = await exerciseApi.getExercises(props.tutorialId);
    // Initialize solutions with starter code
    exercises.value.forEach((ex: Exercise) => {
      if (ex.starterCode) {
        solutions.value[ex.id] = ex.starterCode;
      }
    });
  } catch (err) {
    error.value = err instanceof Error ? err.message : 'Failed to load exercises';
  } finally {
    loading.value = false;
  }
};

const executeSolution = async (exerciseId: string) => {
  const code = solutions.value[exerciseId] || '';
  if (!code.trim()) {
    error.value = 'Please write a solution first';
    return;
  }
  
  await executeCode(code);
  // Store result for this specific exercise
  if (result.value) {
    executionResults.value[exerciseId] = result.value;
  }
};

const checkSolution = (exerciseId: string) => {
  // In a real implementation, this would validate the solution
  // For now, just show a message
  alert('Solution checking will be implemented with validation logic');
};

onMounted(() => {
  loadExercises();
});
</script>

<style scoped>
.exercise-view {
  padding: 2rem;
  max-width: 1200px;
  margin: 0 auto;
}
</style>

