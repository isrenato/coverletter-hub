<script setup lang="ts">
import type { Vacancy } from '../../types'

defineProps<{ vacancy: Vacancy }>()
defineEmits<{ delete: [id: string] }>()
</script>

<template>
  <div class="bg-white rounded-lg shadow p-6">
    <div class="flex justify-between items-start">
      <div>
        <router-link :to="`/vacancies/${vacancy.id}`" class="text-lg font-semibold text-blue-600 hover:underline">
          {{ vacancy.title }}
        </router-link>
        <p class="text-gray-600">{{ vacancy.company }}</p>
        <p class="text-sm text-gray-500">{{ vacancy.location }}</p>
      </div>
      <button @click="$emit('delete', vacancy.id)" class="text-red-500 hover:text-red-700 text-sm">Delete</button>
    </div>
  </div>
</template>
