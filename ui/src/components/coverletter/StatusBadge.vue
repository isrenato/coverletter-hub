<script setup lang="ts">
defineProps<{ status: string }>()

const colors: Record<string, string> = {
  draft: 'bg-yellow-100 text-yellow-800',
  approved: 'bg-green-100 text-green-800',
  rejected: 'bg-red-100 text-red-800',
}
</script>

<template>
  <span :class="[colors[status] || 'bg-gray-100 text-gray-800', 'px-2 py-1 rounded-full text-xs font-medium capitalize']">
    {{ status }}
  </span>
</template>
