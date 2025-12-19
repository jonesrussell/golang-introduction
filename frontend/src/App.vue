<template>
  <div class="flex min-h-screen bg-neutral-50 dark:bg-neutral-950">
    <!-- Mobile menu toggle -->
    <button
      @click="sidebarOpen = !sidebarOpen"
      class="lg:hidden fixed top-4 left-4 z-[60] p-3 bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 rounded-xl shadow-md text-neutral-900 dark:text-neutral-100 transition-all duration-150 hover:bg-primary-light hover:border-go-blue items-center justify-center"
      aria-label="Toggle navigation"
    >
      <svg v-if="!sidebarOpen" class="w-6 h-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16"/>
      </svg>
      <svg v-else class="w-6 h-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>
      </svg>
    </button>

    <!-- Sidebar -->
    <aside
      :class="[
        'fixed top-0 left-0 w-80 h-screen bg-white dark:bg-neutral-900 border-r border-neutral-200 dark:border-neutral-800 flex flex-col z-50 transition-transform duration-200',
        'lg:translate-x-0',
        sidebarOpen ? 'translate-x-0' : '-translate-x-full'
      ]"
    >
      <div class="p-6 border-b border-neutral-200 dark:border-neutral-800 bg-gradient-to-br from-primary-light to-white dark:from-neutral-800 dark:to-neutral-900">
        <div class="flex items-center gap-3">
          <svg class="w-10 h-10 text-go-blue" viewBox="0 0 24 24" fill="currentColor">
            <path d="M1.811 10.715c-.404.258-.378.746.086.911l5.753 2.145c.302.112.553-.038.672-.318l1.183-3.235c.124-.295.032-.583-.239-.69L3.513 7.383c-.453-.178-.882.031-1.14.333l-.562.999zM6.837 15.091c-.302-.112-.553.038-.672.318l-1.183 3.235c-.124.295-.032.583.239.69l5.753 2.145c.464.173.882-.112 1.086-.411l.562-.999c.404-.258.378-.746-.086-.911l-5.699-4.067zM22.189 10.715c.404.258.378.746-.086.911l-5.753 2.145c-.302.112-.553-.038-.672-.318l-1.183-3.235c-.124-.295-.032-.583.239-.69l5.753-2.145c.453-.178.882.031 1.14.333l.562.999zM17.163 15.091c.302-.112.553.038.672.318l1.183 3.235c.124.295.032.583-.239.69l-5.753 2.145c-.464.173-.882-.112-1.086-.411l-.562-.999c-.404-.258-.378-.746.086-.911l5.699-4.067z"/>
          </svg>
          <div>
            <h1 class="text-xl font-bold text-neutral-900 dark:text-neutral-100 m-0 leading-tight">Go Tutorials</h1>
            <span class="text-sm text-neutral-600 dark:text-neutral-400 font-medium">Interactive Learning</span>
          </div>
        </div>
      </div>
      <TutorialList
        :current-tutorial-id="currentTutorialId"
        @select="handleTutorialSelect"
      />
    </aside>

    <!-- Overlay for mobile -->
    <div
      v-if="sidebarOpen"
      class="lg:hidden fixed inset-0 bg-black/50 z-40 backdrop-blur-sm"
      @click="sidebarOpen = false"
    ></div>

    <!-- Main Content -->
    <main class="flex-1 lg:ml-80 min-h-screen bg-neutral-50 dark:bg-neutral-950">
      <TutorialViewer
        v-if="currentTutorialId"
        :tutorial-id="currentTutorialId"
        @home="handleHome"
      />
      <div v-else class="flex items-center justify-center min-h-screen p-8 lg:p-16">
        <div class="text-center max-w-2xl animate-fade-in">
          <div class="w-24 h-24 mx-auto mb-8 text-go-blue lg:w-32 lg:h-32">
            <svg viewBox="0 0 80 80" fill="none" class="w-full h-full">
              <circle cx="40" cy="40" r="38" stroke="currentColor" stroke-width="2" opacity="0.2"/>
              <circle cx="40" cy="40" r="28" stroke="currentColor" stroke-width="2" opacity="0.4"/>
              <path d="M25 35c0-8.284 6.716-15 15-15s15 6.716 15 15" stroke="currentColor" stroke-width="3" stroke-linecap="round"/>
              <circle cx="32" cy="38" r="3" fill="currentColor"/>
              <circle cx="48" cy="38" r="3" fill="currentColor"/>
              <path d="M32 50c0 0 4 6 8 6s8-6 8-6" stroke="currentColor" stroke-width="2.5" stroke-linecap="round"/>
            </svg>
          </div>
          <h2 class="text-4xl lg:text-5xl font-bold text-neutral-900 dark:text-neutral-100 mb-4 text-balance">Welcome to Go Tutorials</h2>
          <p class="text-lg text-neutral-600 dark:text-neutral-400 leading-relaxed mb-8">
            Master the Go programming language with interactive, hands-on tutorials.
            Run code directly in your browser and track your progress.
          </p>
          <div class="flex flex-col gap-4 mb-8">
            <div class="flex items-center justify-center gap-3 p-4 bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 rounded-xl text-base font-medium text-neutral-900 dark:text-neutral-100 transition-all duration-150 hover:border-go-blue hover:bg-primary-light">
              <svg class="w-6 h-6 text-go-blue" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"/>
              </svg>
              <span>Interactive Code Examples</span>
            </div>
            <div class="flex items-center justify-center gap-3 p-4 bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 rounded-xl text-base font-medium text-neutral-900 dark:text-neutral-100 transition-all duration-150 hover:border-go-blue hover:bg-primary-light">
              <svg class="w-6 h-6 text-go-blue" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"/>
              </svg>
              <span>Progress Tracking</span>
            </div>
            <div class="flex items-center justify-center gap-3 p-4 bg-white dark:bg-neutral-900 border border-neutral-200 dark:border-neutral-800 rounded-xl text-base font-medium text-neutral-900 dark:text-neutral-100 transition-all duration-150 hover:border-go-blue hover:bg-primary-light">
              <svg class="w-6 h-6 text-go-blue" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253"/>
              </svg>
              <span>Beginner to Advanced</span>
            </div>
          </div>
          <p class="text-base text-neutral-500 dark:text-neutral-500">Select a tutorial from the sidebar to get started</p>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import TutorialList from './components/TutorialList.vue';
import TutorialViewer from './components/TutorialViewer.vue';

const currentTutorialId = ref<string>('');
const sidebarOpen = ref(false);

const handleTutorialSelect = (tutorialId: string) => {
  currentTutorialId.value = tutorialId;
  sidebarOpen.value = false;
};

const handleHome = () => {
  currentTutorialId.value = '';
};
</script>

