<template>
  <div class="flex items-center justify-center min-h-screen p-8 lg:p-16">
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
</template>
