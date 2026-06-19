<script setup lang="ts">
import { reactive } from 'vue'

const emit = defineEmits<{ submit: [data: { title: string; company: string; description: string; location: string; linkedin_url: string }] }>()

const form = reactive({
  title: '',
  company: '',
  description: '',
  location: '',
  linkedin_url: '',
})

function handleSubmit() {
  emit('submit', { ...form })
  Object.assign(form, { title: '', company: '', description: '', location: '', linkedin_url: '' })
}
</script>

<template>
  <form @submit.prevent="handleSubmit" class="space-y-4 bg-white p-6 rounded-lg shadow">
    <div>
      <label class="block text-sm font-medium text-gray-700">Job Title</label>
      <input v-model="form.title" required class="mt-1 block w-full rounded border-gray-300 shadow-sm" />
    </div>
    <div>
      <label class="block text-sm font-medium text-gray-700">Company</label>
      <input v-model="form.company" required class="mt-1 block w-full rounded border-gray-300 shadow-sm" />
    </div>
    <div>
      <label class="block text-sm font-medium text-gray-700">Description</label>
      <textarea v-model="form.description" required rows="4" class="mt-1 block w-full rounded border-gray-300 shadow-sm" />
    </div>
    <div class="grid grid-cols-2 gap-4">
      <div>
        <label class="block text-sm font-medium text-gray-700">Location</label>
        <input v-model="form.location" class="mt-1 block w-full rounded border-gray-300 shadow-sm" />
      </div>
      <div>
        <label class="block text-sm font-medium text-gray-700">LinkedIn URL</label>
        <input v-model="form.linkedin_url" type="url" class="mt-1 block w-full rounded border-gray-300 shadow-sm" />
      </div>
    </div>
    <button type="submit" class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700">Add Vacancy</button>
  </form>
</template>
