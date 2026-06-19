<script setup lang="ts">
import { ref, computed } from 'vue'

interface FieldPair {
  source: string
  target: string
}

interface Props {
  sourceFields?: string[]
  targetOptions?: string[]
}

const props = withDefaults(defineProps<Props>(), {
  sourceFields: () => [],
  targetOptions: () => []
})

const emit = defineEmits<{
  close: []
  save: [pairs: FieldPair[]]
}>()

const search = ref<string>('')

const pairs = ref<FieldPair[]>(
  props.sourceFields.map((src, i) => ({
    source: src,
    target: props.targetOptions[i] ?? ''
  }))
)

const filteredPairs = computed<FieldPair[]>(() => {
  if (!search.value.trim()) return pairs.value
  const q = search.value.toLowerCase()
  return pairs.value.filter(p => p.source.toLowerCase().includes(q))
})

function addPair(): void {
  pairs.value.push({ source: '', target: '' })
}

function removePair(index: number): void {
  const pair = filteredPairs.value[index]
  const realIndex = pairs.value.indexOf(pair)
  if (realIndex !== -1) pairs.value.splice(realIndex, 1)
}

function saveMapping(): void {
  emit('save', pairs.value.map(p => ({ source: p.source, target: p.target })))
}
</script>

<template>
  <div class="flex flex-col gap-4 p-4">
    <div class="flex items-center justify-between border-b border-gray-200 dark:border-gray-800 pb-4">
      <p class="text-base font-semibold text-gray-900 dark:text-white">Map Fields</p>
      <UButton
        color="neutral"
        variant="ghost"
        icon="i-lucide-x"
        size="sm"
        @click="$emit('close')"
      />
    </div>

    <UInput
      v-model="search"
      icon="i-lucide-search"
      placeholder="Search source fields..."
      size="sm"
    />

    <div class="grid grid-cols-2 px-1">
      <span class="text-xs font-semibold uppercase tracking-wider text-gray-400">Source Field</span>
      <span class="text-xs font-semibold uppercase tracking-wider text-gray-400">Target Key</span>
    </div>

    <div class="flex flex-col divide-y divide-gray-100 dark:divide-gray-800">
      <div
        v-for="(pair, index) in filteredPairs"
        :key="index"
        class="grid grid-cols-2 items-center gap-3 py-2"
      >
        <span class="truncate text-sm text-gray-800 dark:text-gray-200">{{ pair.source }}</span>

        <div class="flex items-center gap-2">
          <USelect
            v-if="targetOptions && targetOptions.length"
            v-model="pair.target"
            :options="[{ label: 'Select target...', value: '' }, ...targetOptions.map(o => ({ label: o, value: o }))]"
            size="sm"
            class="flex-1"
          />
          <UInput
            v-else
            v-model="pair.target"
            placeholder="Enter target key..."
            size="sm"
            class="flex-1"
          />
          <UButton
            color="neutral"
            variant="ghost"
            icon="i-lucide-x"
            size="xs"
            @click="removePair(index)"
          />
        </div>
      </div>
    </div>

    <UButton
      variant="ghost"
      color="primary"
      icon="i-lucide-plus"
      size="sm"
      class="self-start"
      @click="addPair"
    >
      Add Field Pair
    </UButton>

    <div class="flex justify-end gap-2 border-t border-gray-200 dark:border-gray-800 pt-4">
      <UButton variant="ghost" color="neutral" @click="$emit('close')">Cancel</UButton>
      <UButton color="primary" @click="saveMapping">Save Mapping</UButton>
    </div>
  </div>
</template>
