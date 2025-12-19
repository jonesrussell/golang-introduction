<template>
  <div class="app-container">
    <!-- Mobile menu toggle -->
    <button
      @click="sidebarOpen = !sidebarOpen"
      class="mobile-menu-toggle"
      aria-label="Toggle navigation"
    >
      <svg v-if="!sidebarOpen" class="icon" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16"/>
      </svg>
      <svg v-else class="icon" fill="none" viewBox="0 0 24 24" stroke="currentColor">
        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>
      </svg>
    </button>

    <!-- Sidebar -->
    <aside :class="['sidebar', { 'sidebar-open': sidebarOpen }]">
      <div class="sidebar-header">
        <div class="logo">
          <svg class="logo-icon" viewBox="0 0 24 24" fill="currentColor">
            <path d="M1.811 10.715c-.404.258-.378.746.086.911l5.753 2.145c.302.112.553-.038.672-.318l1.183-3.235c.124-.295.032-.583-.239-.69L3.513 7.383c-.453-.178-.882.031-1.14.333l-.562.999zM6.837 15.091c-.302-.112-.553.038-.672.318l-1.183 3.235c-.124.295-.032.583.239.69l5.753 2.145c.464.173.882-.112 1.086-.411l.562-.999c.404-.258.378-.746-.086-.911l-5.699-4.067zM22.189 10.715c.404.258.378.746-.086.911l-5.753 2.145c-.302.112-.553-.038-.672-.318l-1.183-3.235c-.124-.295-.032-.583.239-.69l5.753-2.145c.453-.178.882.031 1.14.333l.562.999zM17.163 15.091c.302-.112.553.038.672.318l1.183 3.235c.124.295.032.583-.239.69l-5.753 2.145c-.464.173-.882-.112-1.086-.411l-.562-.999c-.404-.258-.378-.746.086-.911l5.699-4.067z"/>
          </svg>
          <div class="logo-text">
            <h1>Go Tutorials</h1>
            <span class="logo-tagline">Interactive Learning</span>
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
      class="sidebar-overlay"
      @click="sidebarOpen = false"
    ></div>

    <!-- Main Content -->
    <main class="main-content">
      <TutorialViewer
        v-if="currentTutorialId"
        :tutorial-id="currentTutorialId"
        @home="handleHome"
      />
      <div v-else class="welcome-screen">
        <div class="welcome-content">
          <div class="welcome-icon">
            <svg viewBox="0 0 80 80" fill="none">
              <circle cx="40" cy="40" r="38" stroke="currentColor" stroke-width="2" opacity="0.2"/>
              <circle cx="40" cy="40" r="28" stroke="currentColor" stroke-width="2" opacity="0.4"/>
              <path d="M25 35c0-8.284 6.716-15 15-15s15 6.716 15 15" stroke="currentColor" stroke-width="3" stroke-linecap="round"/>
              <circle cx="32" cy="38" r="3" fill="currentColor"/>
              <circle cx="48" cy="38" r="3" fill="currentColor"/>
              <path d="M32 50c0 0 4 6 8 6s8-6 8-6" stroke="currentColor" stroke-width="2.5" stroke-linecap="round"/>
            </svg>
          </div>
          <h2 class="welcome-title">Welcome to Go Tutorials</h2>
          <p class="welcome-subtitle">
            Master the Go programming language with interactive, hands-on tutorials.
            Run code directly in your browser and track your progress.
          </p>
          <div class="welcome-features">
            <div class="feature">
              <svg class="feature-icon" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 10V3L4 14h7v7l9-11h-7z"/>
              </svg>
              <span>Interactive Code Examples</span>
            </div>
            <div class="feature">
              <svg class="feature-icon" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"/>
              </svg>
              <span>Progress Tracking</span>
            </div>
            <div class="feature">
              <svg class="feature-icon" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253"/>
              </svg>
              <span>Beginner to Advanced</span>
            </div>
          </div>
          <p class="welcome-cta">Select a tutorial from the sidebar to get started</p>
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

<style scoped>
.app-container {
  display: flex;
  min-height: 100vh;
  background-color: var(--color-background);
}

/* Mobile menu toggle */
.mobile-menu-toggle {
  display: none;
  position: fixed;
  top: 1rem;
  left: 1rem;
  z-index: 60;
  padding: 0.75rem;
  background-color: var(--color-surface);
  border: 1px solid var(--color-border);
  border-radius: var(--radius-lg);
  box-shadow: var(--shadow-md);
  color: var(--color-text);
  transition: all var(--transition-fast);
}

.mobile-menu-toggle:hover {
  background-color: var(--color-primary-light);
  border-color: var(--color-primary);
}

.mobile-menu-toggle .icon {
  width: 1.5rem;
  height: 1.5rem;
}

/* Sidebar */
.sidebar {
  position: fixed;
  top: 0;
  left: 0;
  width: 320px;
  height: 100vh;
  background-color: var(--color-surface);
  border-right: 1px solid var(--color-border);
  display: flex;
  flex-direction: column;
  z-index: 50;
  transition: transform var(--transition-normal);
}

.sidebar-header {
  padding: 1.5rem;
  border-bottom: 1px solid var(--color-border);
  background: linear-gradient(135deg, var(--color-primary-light), var(--color-surface));
}

.logo {
  display: flex;
  align-items: center;
  gap: 0.75rem;
}

.logo-icon {
  width: 2.5rem;
  height: 2.5rem;
  color: var(--color-primary);
}

.logo-text h1 {
  font-size: 1.25rem;
  font-weight: 700;
  color: var(--color-text);
  margin: 0;
  line-height: 1.2;
}

.logo-tagline {
  font-size: 0.75rem;
  color: var(--color-text-muted);
  font-weight: 500;
}

.sidebar-overlay {
  display: none;
  position: fixed;
  inset: 0;
  background-color: rgba(0, 0, 0, 0.5);
  z-index: 40;
  backdrop-filter: blur(2px);
}

/* Main content */
.main-content {
  flex: 1;
  margin-left: 320px;
  min-height: 100vh;
  background-color: var(--color-background);
}

/* Welcome screen */
.welcome-screen {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
  padding: 2rem;
}

.welcome-content {
  text-align: center;
  max-width: 32rem;
  animation: fadeIn 0.5s ease-out;
}

.welcome-icon {
  width: 6rem;
  height: 6rem;
  margin: 0 auto 2rem;
  color: var(--color-primary);
}

.welcome-icon svg {
  width: 100%;
  height: 100%;
}

.welcome-title {
  font-size: 2rem;
  font-weight: 700;
  color: var(--color-text);
  margin: 0 0 1rem;
  text-wrap: balance;
}

.welcome-subtitle {
  font-size: 1.125rem;
  color: var(--color-text-muted);
  line-height: 1.7;
  margin: 0 0 2rem;
}

.welcome-features {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  margin-bottom: 2rem;
}

.feature {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.75rem;
  padding: 0.75rem 1.5rem;
  background-color: var(--color-surface);
  border: 1px solid var(--color-border);
  border-radius: var(--radius-lg);
  font-weight: 500;
  color: var(--color-text);
  transition: all var(--transition-fast);
}

.feature:hover {
  border-color: var(--color-primary);
  background-color: var(--color-primary-light);
}

.feature-icon {
  width: 1.25rem;
  height: 1.25rem;
  color: var(--color-primary);
}

.welcome-cta {
  font-size: 0.875rem;
  color: var(--color-text-subtle);
}

/* Responsive design */
@media (max-width: 1024px) {
  .mobile-menu-toggle {
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .sidebar {
    transform: translateX(-100%);
  }

  .sidebar-open {
    transform: translateX(0);
  }

  .sidebar-overlay {
    display: block;
  }

  .main-content {
    margin-left: 0;
  }

  .welcome-screen {
    padding: 4rem 1.5rem 2rem;
  }
}

@media (max-width: 640px) {
  .sidebar {
    width: 100%;
  }

  .welcome-title {
    font-size: 1.5rem;
  }

  .welcome-subtitle {
    font-size: 1rem;
  }

  .welcome-icon {
    width: 4rem;
    height: 4rem;
    margin-bottom: 1.5rem;
  }
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(10px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>
