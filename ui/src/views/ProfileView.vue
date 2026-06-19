<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useProfileStore } from '../stores/profile'

const store = useProfileStore()
const activeTab = ref<'manual' | 'upload' | 'linkedin'>('manual')
const fileInput = ref<HTMLInputElement>()

onMounted(() => {
  store.fetchProfile()
})

async function handleUpload() {
  const file = fileInput.value?.files?.[0]
  if (file) {
    await store.uploadCV(file)
  }
}

async function handleSave() {
  if (store.profile) {
    await store.updateProfile(store.profile)
  }
}
</script>

<template>
  <div>
    <h1 class="text-2xl font-bold mb-6">CV Profile</h1>

    <div class="border-b border-gray-200 mb-6">
      <nav class="-mb-px flex space-x-8">
        <button
          v-for="tab in (['manual', 'upload', 'linkedin'] as const)"
          :key="tab"
          @click="activeTab = tab"
          :class="[
            activeTab === tab ? 'border-blue-500 text-blue-600' : 'border-transparent text-gray-500 hover:text-gray-700',
            'whitespace-nowrap py-4 px-1 border-b-2 font-medium text-sm capitalize'
          ]"
        >
          {{ tab === 'linkedin' ? 'Import from LinkedIn' : tab === 'upload' ? 'Upload CV' : 'Manual Entry' }}
        </button>
      </nav>
    </div>

    <div v-if="store.loading" class="text-center py-8 text-gray-500">Loading...</div>

    <div v-else-if="activeTab === 'upload'" class="space-y-4">
      <div class="border-2 border-dashed border-gray-300 rounded-lg p-8 text-center">
        <input ref="fileInput" type="file" accept=".pdf,.docx" class="hidden" @change="handleUpload" />
        <button @click="fileInput?.click()" class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700">
          Choose File (PDF or DOCX)
        </button>
      </div>
    </div>

    <div v-else-if="activeTab === 'manual' && store.profile" class="space-y-4">
      <div>
        <label class="block text-sm font-medium text-gray-700">Full Name</label>
        <input v-model="store.profile.full_name" class="mt-1 block w-full rounded border-gray-300 shadow-sm" />
      </div>
      <div>
        <label class="block text-sm font-medium text-gray-700">Headline</label>
        <input v-model="store.profile.headline" class="mt-1 block w-full rounded border-gray-300 shadow-sm" />
      </div>
      <div>
        <label class="block text-sm font-medium text-gray-700">Summary</label>
        <textarea v-model="store.profile.summary" rows="4" class="mt-1 block w-full rounded border-gray-300 shadow-sm" />
      </div>
      <button @click="handleSave" class="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700">
        Save Profile
      </button>
    </div>

    <div v-else-if="activeTab === 'linkedin'" class="text-center py-8">
      <p class="text-gray-500">LinkedIn import coming soon</p>
    </div>

    <div v-if="store.error" class="mt-4 p-4 bg-red-50 text-red-700 rounded">{{ store.error }}</div>
  </div>
</template>
