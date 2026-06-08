<template>
  <!-- PrimeVue Drawer — slides in from the right -->
  <Drawer
    :visible="visible"
    position="right"
    style="width: 420px;"
    :header="drawerTitle"
    @update:visible="emit('update:visible', $event)"
  >
    <!-- ── Step 1: Pick a component type ── -->
    <div v-if="step === 'pick-type'" style="display: flex; flex-direction: column; gap: 12px; padding: 0.25rem 0;">
      <p style="font-size: 0.8125rem; color: #888680; margin: 0 0 0.5rem;">
        Choose the type of component for this entry.
      </p>

      <div
        v-for="(typeKey) in availableTypes"
        :key="typeKey"
        style="display: flex; align-items: center; gap: 14px; padding: 14px 16px; background: #fff; border: 1px solid #e5e3dc; border-radius: 12px; cursor: pointer; transition: border-color 0.15s, box-shadow 0.15s;"
        :style="hoveredType === typeKey ? { borderColor: '#9f85e8', boxShadow: '0 0 0 3px rgba(139,115,224,0.1)' } : {}"
        @click="onTypePicked(typeKey)"
        @mouseenter="hoveredType = typeKey"
        @mouseleave="hoveredType = null"
      >
        <span
          style="display: flex; align-items: center; justify-content: center; width: 36px; height: 36px; border-radius: 8px; background: #f3f0fd; flex-shrink: 0;"
        >
          <i :class="nodeTypeIcons[typeKey]" style="color: #6b52d9; font-size: 0.9375rem;" />
        </span>
        <div>
          <p style="font-size: 0.9375rem; font-weight: 600; color: #1a1917; margin: 0 0 2px;">
            {{ nodeTypeLabels[typeKey] }}
          </p>
          <p style="font-size: 0.75rem; color: #b0ada5; margin: 0;">
            {{ nodeTypeDescriptions[typeKey] }}
          </p>
        </div>
        <i class="pi pi-chevron-right" style="margin-left: auto; color: #c2bfb8; font-size: 0.75rem;" />
      </div>
    </div>

    <!-- ── Step 2: Configure the node via NodePropertyEditor ── -->
    <div v-else-if="step === 'configure' && draftNode" style="display: flex; flex-direction: column; height: 100%;">
      <!-- Back button -->
      <button
        style="display: inline-flex; align-items: center; gap: 6px; background: none; border: none; cursor: pointer; color: #888680; font-size: 0.8125rem; font-weight: 500; padding: 0 0 1.25rem; font-family: inherit;"
        @click="step = 'pick-type'"
      >
        <i class="pi pi-arrow-left" style="font-size: 0.75rem;" />
        Back
      </button>

      <!-- NodePropertyEditor handles all config fields for the chosen type -->
      <NodePropertyEditor v-model:node="draftNode" style="flex: 1;" />

      <Divider style="margin: 1.25rem 0 1rem;" />

      <!-- Footer actions -->
      <div style="display: flex; gap: 8px; justify-content: flex-end;">
        <Button
          label="Cancel"
          text
          severity="secondary"
          style="font-size: 0.8125rem;"
          @click="onCancel"
        />
        <Button
          label="Save entry"
          icon="pi pi-check"
          style="font-size: 0.8125rem; background: #5b3ec8; border-color: #5b3ec8;"
          @click="onSave"
        />
      </div>
    </div>
  </Drawer>
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue'
import Drawer from 'primevue/drawer'
import Button from 'primevue/button'
import Divider from 'primevue/divider'
import NodePropertyEditor from './NodePropertyEditor.vue'
import {
  type AnyNode,
  nodeDefaults,
  nodeTypeLabels,
  nodeTypeIcons,
} from './nodeDefaults'

const nodeTypeDescriptions: Record<string, string> = {
  inputtext:   'A single line of free text',
  inputnumber: 'A numeric value with optional min/max',
  dropdown:    'A selection from a predefined list',
  repeater:    'A nested list of any component',
}

const availableTypes = Object.keys(nodeDefaults)

// ── Props & emits ────────────────────────────────────────────────────────────

interface Props {
  visible: boolean
  // Pass an existing node to edit, or leave undefined to create a new one
  editingNode?: AnyNode | null
}

const props = withDefaults(defineProps<Props>(), {
  editingNode: null,
})

const emit = defineEmits<{
  'update:visible': [value: boolean]
  save: [node: AnyNode]
}>()

// ── Internal state ───────────────────────────────────────────────────────────

type DrawerStep = 'pick-type' | 'configure'

const step = ref<DrawerStep>('pick-type')
const draftNode = ref<AnyNode | null>(null)
const hoveredType = ref<string | null>(null)

const drawerTitle = computed(() =>
  step.value === 'pick-type' ? 'Choose component' : `Configure ${draftNode.value ? nodeTypeLabels[draftNode.value.type] : ''}`
)

// If an existing node is passed in (edit mode), jump straight to configure
watch(() => props.visible, (open) => {
  if (open && props.editingNode) {
    // Deep clone so edits don't mutate the original until Save
    draftNode.value = JSON.parse(JSON.stringify(props.editingNode))
    step.value = 'configure'
  } else if (open) {
    step.value = 'pick-type'
    draftNode.value = null
  }
})

// ── Handlers ─────────────────────────────────────────────────────────────────

function onTypePicked(typeKey: string) {
  draftNode.value = nodeDefaults[typeKey]()
  step.value = 'configure'
}

function onSave() {
  if (!draftNode.value) return
  emit('save', JSON.parse(JSON.stringify(draftNode.value)))
  emit('update:visible', false)
}

function onCancel() {
  emit('update:visible', false)
}
</script>





<template>
  <div style="font-family: 'DM Sans', system-ui, sans-serif; max-width: 600px;">

    <!-- Repeater label -->
    <p style="font-size: 0.8125rem; font-weight: 600; color: #1a1917; margin: 0 0 0.75rem; letter-spacing: -0.01em;">
      {{ node.label || 'Repeater' }}
    </p>

    <!-- Entry list -->
    <TransitionGroup tag="div" name="entry" style="display: flex; flex-direction: column; gap: 8px; margin-bottom: 8px;">
      <div
        v-for="(entry, index) in node.entries"
        :key="entryKeys[index]"
        style="display: flex; align-items: flex-start; gap: 10px; padding: 14px 14px 14px 16px; background: #fff; border: 1px solid #e5e3dc; border-radius: 12px;"
      >
        <!-- Type badge -->
        <span
          style="display: flex; align-items: center; justify-content: center; width: 30px; height: 30px; border-radius: 7px; background: #f3f0fd; flex-shrink: 0; margin-top: 1px;"
        >
          <i :class="nodeTypeIcons[entry.type]" style="color: #6b52d9; font-size: 0.8125rem;" />
        </span>

        <!-- Rendered node -->
        <div style="flex: 1; min-width: 0;">
          <p style="font-size: 0.6875rem; font-weight: 600; color: #b0ada5; margin: 0 0 6px; text-transform: uppercase; letter-spacing: 0.05em;">
            {{ nodeTypeLabels[entry.type] }}
          </p>
          <NodeRenderer :node="entry" />
        </div>

        <!-- Edit & remove -->
        <div style="display: flex; gap: 4px; flex-shrink: 0;">
          <Button
            icon="pi pi-pencil"
            text
            rounded
            severity="secondary"
            size="small"
            :aria-label="`Edit entry ${index + 1}`"
            style="width: 28px; height: 28px; padding: 0; color: #b0ada5;"
            @click="openEditDrawer(index)"
          />
          <Button
            icon="pi pi-times"
            text
            rounded
            severity="secondary"
            size="small"
            :aria-label="`Remove entry ${index + 1}`"
            style="width: 28px; height: 28px; padding: 0; color: #b0ada5;"
            @click="removeEntry(index)"
          />
        </div>
      </div>
    </TransitionGroup>

    <!-- Add entry -->
    <Button
      label="Add entry"
      icon="pi pi-plus"
      text
      style="width: 100%; justify-content: center; border: 1px dashed #d4d0c8; border-radius: 10px; color: #888680; font-size: 0.8125rem; font-weight: 500; padding: 9px; margin-top: 2px;"
      @click="openAddDrawer"
    />

    <!-- Drawer — handles both add and edit -->
    <RepeaterEntryDrawer
      v-model:visible="drawerVisible"
      :editing-node="editingNode"
      @save="onSave"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import Button from 'primevue/button'
import NodeRenderer from './NodeRenderer.vue'
import RepeaterEntryDrawer from './RepeaterEntryDrawer.vue'
import { type AnyNode, type RepeaterNode, nodeTypeLabels, nodeTypeIcons } from './nodeDefaults'

// ── Props & model ─────────────────────────────────────────────────────────────

const node = defineModel<RepeaterNode>({ required: true })

// ── Stable keys for TransitionGroup ──────────────────────────────────────────
// We generate a key per entry once on creation so TransitionGroup can track
// them correctly even when entries are spliced out and re-added.

let keyCounter = 0
const entryKeys = ref<number[]>(node.value.entries.map(() => keyCounter++))

// ── Drawer state ──────────────────────────────────────────────────────────────

const drawerVisible = ref(false)
const editingNode = ref<AnyNode | null>(null)
const editingIndex = ref<number | null>(null)

function openAddDrawer() {
  editingNode.value = null   // null = new entry, drawer starts at pick-type step
  editingIndex.value = null
  drawerVisible.value = true
}

function openEditDrawer(index: number) {
  editingNode.value = node.value.entries[index]  // existing node → drawer jumps to configure
  editingIndex.value = index
  drawerVisible.value = true
}

function removeEntry(index: number) {
  node.value.entries.splice(index, 1)
  entryKeys.value.splice(index, 1)
}

// Called by the drawer when the user hits Save
function onSave(saved: AnyNode) {
  if (editingIndex.value !== null) {
    // Edit — replace in place
    node.value.entries.splice(editingIndex.value, 1, saved)
  } else {
    // Add — append with a fresh stable key
    node.value.entries.push(saved)
    entryKeys.value.push(keyCounter++)
  }
  editingIndex.value = null
}
</script>
