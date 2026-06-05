<template>
  <div style="font-family: 'DM Sans', system-ui, sans-serif; max-width: 560px; background: #fafaf9; border: 1px solid #e5e3dc; border-radius: 16px; padding: 2rem;">

    <!-- Header -->
    <div style="margin-bottom: 1.5rem;">
      <p style="font-size: 1.0625rem; font-weight: 600; color: #1a1917; letter-spacing: -0.02em; margin: 0 0 0.25rem;">{{ title }}</p>
      <p style="font-size: 0.8125rem; color: #888680; margin: 0; line-height: 1.5;">{{ subtitle }}</p>
    </div>

    <!-- Entry list -->
    <TransitionGroup tag="div" name="entry" style="display: flex; flex-direction: column; gap: 6px; margin-bottom: 8px;">
      <div
        v-for="(entry, index) in entries"
        :key="entry.id"
        style="display: flex; align-items: center; gap: 10px; padding: 4px 8px 4px 12px; background: #fff; border-radius: 10px; border: 1px solid #e5e3dc; transition: border-color 0.15s, box-shadow 0.15s;"
        :style="focusedIndex === index ? { borderColor: '#9f85e8', boxShadow: '0 0 0 3px rgba(139,115,224,0.1)' } : {}"
      >
        <!-- Row number badge -->
        <span style="font-size: 0.6875rem; font-weight: 600; color: #c2bfb8; min-width: 16px; text-align: center; user-select: none; font-variant-numeric: tabular-nums;">
          {{ index + 1 }}
        </span>

        <!-- PrimeVue InputText -->
        <InputText
          :ref="el => setInputRef(el, index)"
          v-model="entry.value"
          :placeholder="placeholder"
          :aria-label="`Entry ${index + 1}`"
          variant="filled"
          style="flex: 1; border: none; outline: none; background: transparent; box-shadow: none; padding: 6px 0; font-size: 0.9375rem; font-family: inherit; color: #1a1917;"
          @focus="focusedIndex = index"
          @blur="focusedIndex = null"
          @keydown.enter="handleEnter(index)"
          @keydown.backspace="handleBackspace(index, $event)"
          @keydown.up.prevent="focusEntry(index - 1)"
          @keydown.down.prevent="focusEntry(index + 1)"
        />

        <!-- PrimeVue remove Button -->
        <Button
          icon="pi pi-times"
          text
          rounded
          severity="secondary"
          size="small"
          :aria-label="`Remove entry ${index + 1}`"
          :disabled="entries.length === 1 && !entry.value"
          style="width: 28px; height: 28px; padding: 0; color: #c2bfb8; flex-shrink: 0;"
          @click="removeEntry(index)"
        />
      </div>
    </TransitionGroup>

    <!-- Add entry button -->
    <Button
      label="Add entry"
      icon="pi pi-plus"
      text
      style="width: 100%; justify-content: center; border: 1px dashed #d4d0c8; border-radius: 10px; color: #888680; font-size: 0.8125rem; font-weight: 500; padding: 8px; margin-top: 2px;"
      @click="addEntry()"
    />

    <!-- Divider -->
    <Divider style="margin: 1.25rem 0;" />

    <!-- Footer -->
    <div style="display: flex; align-items: center; justify-content: space-between;">
      <span style="font-size: 0.8125rem; color: #b0ada5; font-variant-numeric: tabular-nums;">
        {{ filledCount }} / {{ entries.length }} filled
      </span>

      <div style="display: flex; gap: 8px; align-items: center;">
        <!-- PrimeVue clear Button -->
        <Button
          label="Clear all"
          text
          severity="secondary"
          :disabled="entries.length <= 1 && !entries[0]?.value"
          style="font-size: 0.8125rem; font-weight: 500; padding: 7px 13px; color: #888680;"
          @click="clearAll"
        />

        <!-- PrimeVue submit Button -->
        <Button
          :label="filledCount > 0 ? `Submit (${filledCount})` : 'Submit'"
          :disabled="filledCount === 0"
          style="font-size: 0.8125rem; font-weight: 600; padding: 7px 16px; background: #5b3ec8; border-color: #5b3ec8; letter-spacing: -0.01em;"
          @click="submit"
        />
      </div>
    </div>

    <!-- PrimeVue Toast for submit feedback -->
    <Toast position="bottom-center" />
  </div>
</template>

<script setup lang="ts">
import { ref, computed, nextTick } from 'vue'
import InputText from 'primevue/inputtext'
import Button from 'primevue/button'
import Divider from 'primevue/divider'
import Toast from 'primevue/toast'
import { useToast } from 'primevue/usetoast'

interface Entry {
  id: number
  value: string
}

interface Props {
  title?: string
  subtitle?: string
  placeholder?: string
}

const props = withDefaults(defineProps<Props>(), {
  title: 'Your entries',
  subtitle: 'Add as many items as you need. Press Enter to add a new one.',
  placeholder: 'Type something…',
})

const emit = defineEmits<{
  submit: [values: string[]]
}>()

const toast = useToast()

let nextId = 0
const entries = ref<Entry[]>([{ id: nextId++, value: '' }])
const focusedIndex = ref<number | null>(null)
const inputRefs = ref<(InstanceType<typeof InputText> | null)[]>([])

const filledCount = computed(() => entries.value.filter(e => e.value.trim()).length)

function setInputRef(el: InstanceType<typeof InputText> | null, index: number) {
  inputRefs.value[index] = el
}

function focusEntry(index: number) {
  if (index >= 0 && index < entries.value.length) {
    nextTick(() => {
      const ref = inputRefs.value[index]
      // PrimeVue InputText exposes $el which is the native input
      const native = (ref as any)?.$el as HTMLInputElement | undefined
      native?.focus()
    })
  }
}

function addEntry() {
  entries.value.push({ id: nextId++, value: '' })
  nextTick(() => focusEntry(entries.value.length - 1))
}

function removeEntry(index: number) {
  if (entries.value.length === 1) {
    entries.value[0].value = ''
    return
  }
  entries.value.splice(index, 1)
  inputRefs.value.splice(index, 1)
  nextTick(() => focusEntry(Math.min(index, entries.value.length - 1)))
}

function handleEnter(index: number) {
  if (index === entries.value.length - 1) {
    addEntry()
  } else {
    focusEntry(index + 1)
  }
}

function handleBackspace(index: number, event: KeyboardEvent) {
  if (entries.value[index].value === '' && entries.value.length > 1) {
    event.preventDefault()
    removeEntry(index)
  }
}

function clearAll() {
  entries.value = [{ id: nextId++, value: '' }]
  inputRefs.value = []
  nextTick(() => focusEntry(0))
}

function submit() {
  const values = entries.value.map(e => e.value.trim()).filter(Boolean)
  emit('submit', values)
  toast.add({
    severity: 'success',
    summary: 'Submitted',
    detail: `${values.length} entr${values.length === 1 ? 'y' : 'ies'} submitted successfully.`,
    life: 3000,
  })
}
</script>

<!--
  No <style> block — all styling is done via style="" inline attributes.

  Required setup in main.ts:
    import PrimeVue from 'primevue/config'
    import Aura from '@primevue/themes/aura'
    import ToastService from 'primevue/toastservice'
    import 'primeicons/primeicons.css'

    app.use(PrimeVue, { theme: { preset: Aura } })
    app.use(ToastService)
-->
