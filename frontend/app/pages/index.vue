<script setup lang="ts">
const count = ref(0)
const name = ref('')

const reset = () => {
  count.value = 0
  name.value = ''
}
</script>

<template>
  <UContainer class="py-10">
    
    <UCard>
      <template #header>
        <div class="flex justify-between items-center">
          <h1 class="text-xl font-bold text-primary-500">Hello Nuxt UI</h1>
          <UBadge variant="subtle">Vue 3 + Nuxt 4</UBadge>
        </div>
      </template>

      <div class="space-y-6">
        <p class="text-gray-600 dark:text-gray-300">
          Example page
        </p>

        <div class="flex flex-col gap-4 p-4 border rounded-lg border-gray-200 dark:border-gray-700">
          <div class="flex items-center gap-4">
            <UButton @click="count++">
              Add count
            </UButton>
            
            <span class="font-mono text-2xl">{{ count }}</span>
          </div>

          <UInput 
            v-model="name" 
            icon="i-heroicons-user"
            placeholder="Input your name..." 
          />
          
          <p v-if="name" class="text-sm text-blue-500">
            welcome: {{ name }}！
          </p>
        </div>
      </div>

      <template #footer>
        <div class="flex justify-end">
          <UButton
            variant="ghost" 
            @click="reset"
          >
            Reset status
          </UButton>
        </div>
      </template>
    </UCard>

  </UContainer>
</template>