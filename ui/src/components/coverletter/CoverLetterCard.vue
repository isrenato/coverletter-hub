<script setup lang="ts">
import type { CoverLetter } from '../../types'
import StatusBadge from './StatusBadge.vue'

defineProps<{ coverLetter: CoverLetter }>()
</script>

<template>
  <router-link :to="`/cover-letters/${coverLetter.id}`" class="block bg-white rounded-lg shadow p-6 hover:shadow-md transition">
    <div class="flex justify-between items-start">
      <div>
        <p class="text-sm text-gray-500">Generated {{ new Date(coverLetter.generated_at).toLocaleDateString() }}</p>
        <p class="mt-1 text-gray-700 line-clamp-2">{{ coverLetter.generated_text.substring(0, 150) }}...</p>
      </div>
      <StatusBadge :status="coverLetter.status" />
    </div>
  </router-link>
</template>
