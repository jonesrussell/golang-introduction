<template>
  <div class="flex min-h-screen bg-neutral-50 dark:bg-neutral-950">
    <!-- Mobile menu toggle -->
    <button
      type="button"
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
      <TutorialList :current-tutorial-id="currentTutorialId" />
    </aside>

    <!-- Overlay for mobile -->
    <div
      v-if="sidebarOpen"
      class="lg:hidden fixed inset-0 bg-black/50 z-40 backdrop-blur-sm"
      @click="sidebarOpen = false"
    ></div>

    <!-- Main Content -->
    <main class="flex-1 lg:ml-80 min-h-screen bg-neutral-50 dark:bg-neutral-950">
      <RouterView />
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue';
import { RouterView, useRoute } from 'vue-router';
import TutorialList from './components/TutorialList.vue';

const route = useRoute();
const sidebarOpen = ref(false);

const currentTutorialId = computed(() => {
  return route.name === 'tutorial' && typeof route.params.id === 'string' 
    ? route.params.id 
    : '';
});

// Close sidebar on route change (for mobile)
watch(() => route.path, () => {
  sidebarOpen.value = false;
});
</script>

