<template>
  <TutorialViewer :tutorial-id="id" />
</template>

<script setup lang="ts">
import TutorialViewer from '../components/TutorialViewer.vue'

defineProps<{
  id: string
}>()
</script>
