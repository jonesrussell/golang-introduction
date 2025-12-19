<template>
  <TutorialViewer 
    :tutorial-id="id" 
    :section-index="sectionIndex ? parseInt(sectionIndex) - 1 : undefined"
  />
</template>

<script setup lang="ts">
import TutorialViewer from '../components/TutorialViewer.vue'

defineProps<{
  id: string
  sectionIndex?: string
}>()
</script>
