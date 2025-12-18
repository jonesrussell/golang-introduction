<template>
  <div id="app" class="min-h-screen bg-gray-50">
    <div class="flex h-screen">
      <!-- Sidebar -->
      <aside class="w-80 bg-white border-r border-gray-200 overflow-y-auto">
        <div class="p-4 border-b border-gray-200">
          <h1 class="text-2xl font-bold text-gray-800">Go Tutorials</h1>
          <p class="text-sm text-gray-600 mt-1">Interactive Learning Platform</p>
        </div>
        <TutorialList
          :current-tutorial-id="currentTutorialId"
          @select="handleTutorialSelect"
        />
      </aside>

      <!-- Main Content -->
      <main class="flex-1 overflow-y-auto">
        <TutorialViewer
          v-if="currentTutorialId"
          :tutorial-id="currentTutorialId"
          @home="handleHome"
        />
        <div v-else class="flex items-center justify-center h-full text-gray-500">
          <div class="text-center">
            <h2 class="text-2xl font-semibold mb-2">Welcome to Go Tutorials</h2>
            <p>Select a tutorial from the sidebar to get started</p>
          </div>
        </div>
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import TutorialList from './components/TutorialList.vue';
import TutorialViewer from './components/TutorialViewer.vue';

const currentTutorialId = ref<string>('');

const handleTutorialSelect = (tutorialId: string) => {
  currentTutorialId.value = tutorialId;
};

const handleHome = () => {
  currentTutorialId.value = '';
};
</script>

<style>
#app {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif;
}
</style>
