<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRoute } from 'vue-router'
import { useCoverLetterStore } from '../stores/coverletter'
import StatusBadge from '../components/coverletter/StatusBadge.vue'

const route = useRoute()
const store = useCoverLetterStore()
const editText = ref('')

onMounted(async () => {
  await store.fetchOne(route.params.id as string)
  if (store.current) {
    editText.value = store.current.edited_text || store.current.generated_text
  }
})

async function save() {
  if (store.current) {
    await store.updateText(store.current.id, editText.value)
  }
}

async function approve() {
  if (store.current) {
    await store.updateStatus(store.current.id, 'approved')
  }
}

async function reject() {
  if (store.current) {
    await store.updateStatus(store.current.id, 'rejected')
  }
}

function resetToOriginal() {
  if (store.current) {
    editText.value = store.current.generated_text
  }
}
</script>

<template>
  <div v-if="store.current">
    <div class="flex justify-between items-center mb-6">
      <div class="flex items-center gap-4">
        <router-link to="/cover-letters" class="text-blue-600 hover:underline text-sm">&larr; Back</router-link>
        <h1 class="text-2xl font-bold">Edit Cover Letter</h1>
        <StatusBadge :status="store.current.status" />
      </div>
      <div class="flex gap-2">
        <button @click="resetToOriginal" class="px-3 py-2 text-sm border rounded hover:bg-gray-50">Reset to Original</button>
        <button @click="save" class="px-3 py-2 text-sm bg-blue-600 text-white rounded hover:bg-blue-700">Save</button>
        <button @click="approve" class="px-3 py-2 text-sm bg-green-600 text-white rounded hover:bg-green-700">Approve</button>
        <button @click="reject" class="px-3 py-2 text-sm bg-red-600 text-white rounded hover:bg-red-700">Reject</button>
      </div>
    </div>

    <textarea
      v-model="editText"
      rows="20"
      class="w-full p-4 border rounded-lg font-serif text-lg leading-relaxed"
    />

    <div v-if="store.error" class="mt-4 p-4 bg-red-50 text-red-700 rounded">{{ store.error }}</div>
  </div>
  <div v-else-if="store.loading" class="text-center py-8 text-gray-500">Loading...</div>
</template>
