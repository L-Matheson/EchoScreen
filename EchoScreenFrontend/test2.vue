<template>
  <div class="modal-backdrop" @click.self="$emit('close')">
    <div class="modal">
      <!-- Header -->
      <div class="modal-header">
        <h2 class="modal-title">Map Fields</h2>
        <button class="close-btn" @click="$emit('close')" aria-label="Close">
          <svg width="16" height="16" viewBox="0 0 16 16" fill="none">
            <path d="M12 4L4 12M4 4l8 8" stroke="currentColor" stroke-width="1.75" stroke-linecap="round"/>
          </svg>
        </button>
      </div>

      <!-- Search -->
      <div class="search-wrapper">
        <svg class="search-icon" width="15" height="15" viewBox="0 0 15 15" fill="none">
          <circle cx="6.5" cy="6.5" r="5" stroke="currentColor" stroke-width="1.5"/>
          <path d="M10.5 10.5L14 14" stroke="currentColor" stroke-width="1.5" stroke-linecap="round"/>
        </svg>
        <input
          v-model="search"
          type="text"
          class="search-input"
          placeholder="Search source fields..."
        />
      </div>

      <!-- Column Labels -->
      <div class="column-labels">
        <span>SOURCE FIELD</span>
        <span>TARGET KEY</span>
      </div>

      <!-- Field Pairs -->
      <div class="field-list">
        <div
          v-for="(pair, index) in filteredPairs"
          :key="index"
          class="field-row"
        >
          <!-- Source Field (read-only) -->
          <div class="field-source">{{ pair.source }}</div>

          <!-- Target Key (dropdown if targetOptions given, else text input) -->
          <select
            v-if="targetOptions && targetOptions.length"
            v-model="pair.target"
            class="field-target"
            :class="{ placeholder: !pair.target }"
          >
            <option value="" disabled>Select target...</option>
            <option v-for="opt in targetOptions" :key="opt" :value="opt">{{ opt }}</option>
          </select>
          <input
            v-else
            v-model="pair.target"
            type="text"
            class="field-target"
            placeholder="Enter target key..."
          />

          <!-- Remove row -->
          <button class="remove-btn" @click="removePair(index)" aria-label="Remove field pair">
            <svg width="13" height="13" viewBox="0 0 13 13" fill="none">
              <path d="M10 3L3 10M3 3l7 7" stroke="currentColor" stroke-width="1.5" stroke-linecap="round"/>
            </svg>
          </button>
        </div>
      </div>

      <!-- Add Field Pair -->
      <button class="add-btn" @click="addPair">
        <svg width="14" height="14" viewBox="0 0 14 14" fill="none">
          <path d="M7 1v12M1 7h12" stroke="currentColor" stroke-width="1.75" stroke-linecap="round"/>
        </svg>
        Add Field Pair
      </button>

      <!-- Footer -->
      <div class="modal-footer">
        <button class="btn-cancel" @click="$emit('close')">Cancel</button>
        <button class="btn-save" @click="saveMapping">Save Mapping</button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'

const props = defineProps({
  /**
   * Array of source field name strings.
   * e.g. ['customer_id', 'transaction_date', 'order_value_usd', 'loyalty_tier']
   */
  sourceFields: {
    type: Array,
    default: () => []
  },
  /**
   * Array of target key strings.
   * If provided, target column renders as a <select> dropdown.
   * If empty/omitted, target column renders as a free-text <input>.
   * e.g. ['uid_pk', 'purchase_time', 'amount_gross']
   */
  targetOptions: {
    type: Array,
    default: () => []
  }
})

const emit = defineEmits(['close', 'save'])

const search = ref('')

// Build initial pairs from sourceFields prop
const pairs = ref(
  props.sourceFields.map((src, i) => ({
    source: src,
    target: props.targetOptions[i] ?? ''
  }))
)

const filteredPairs = computed(() => {
  if (!search.value.trim()) return pairs.value
  const q = search.value.toLowerCase()
  return pairs.value.filter(p => p.source.toLowerCase().includes(q))
})

function addPair() {
  pairs.value.push({ source: '', target: '' })
}

function removePair(index) {
  // Find the actual index in the full pairs array (search may filter)
  const pair = filteredPairs.value[index]
  const realIndex = pairs.value.indexOf(pair)
  if (realIndex !== -1) pairs.value.splice(realIndex, 1)
}

function saveMapping() {
  emit('save', pairs.value.map(p => ({ source: p.source, target: p.target })))
}
</script>

<style scoped>
/* ── Backdrop ── */
.modal-backdrop {
  position: fixed;
  inset: 0;
  background: rgba(0, 0, 0, 0.35);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

/* ── Modal shell ── */
.modal {
  background: #fff;
  border-radius: 10px;
  width: 480px;
  max-width: 95vw;
  max-height: 90vh;
  display: flex;
  flex-direction: column;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.14);
  overflow: hidden;
}

/* ── Header ── */
.modal-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 18px 20px 14px;
  border-bottom: 1px solid #ebebeb;
}

.modal-title {
  font-size: 15px;
  font-weight: 600;
  color: #111;
  margin: 0;
}

.close-btn {
  background: none;
  border: none;
  cursor: pointer;
  color: #888;
  padding: 4px;
  border-radius: 4px;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: color 0.15s;
}
.close-btn:hover { color: #111; }

/* ── Search ── */
.search-wrapper {
  position: relative;
  padding: 14px 20px 10px;
}

.search-icon {
  position: absolute;
  left: 34px;
  top: 50%;
  transform: translateY(-50%);
  color: #aaa;
  pointer-events: none;
}

.search-input {
  width: 100%;
  box-sizing: border-box;
  padding: 8px 12px 8px 36px;
  border: 1px solid #ddd;
  border-radius: 6px;
  font-size: 13.5px;
  color: #333;
  outline: none;
  background: #fafafa;
  transition: border-color 0.15s;
}
.search-input:focus { border-color: #2563eb; background: #fff; }
.search-input::placeholder { color: #bbb; }

/* ── Column labels ── */
.column-labels {
  display: grid;
  grid-template-columns: 1fr 1fr;
  padding: 4px 20px 6px;
  font-size: 10.5px;
  font-weight: 600;
  letter-spacing: 0.06em;
  color: #999;
  text-transform: uppercase;
}

/* ── Field list ── */
.field-list {
  flex: 1;
  overflow-y: auto;
  padding: 0 20px;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.field-row {
  display: grid;
  grid-template-columns: 1fr 1fr auto;
  align-items: center;
  gap: 10px;
  padding: 6px 0;
  border-bottom: 1px solid #f2f2f2;
}
.field-row:last-child { border-bottom: none; }

/* Source label */
.field-source {
  font-size: 13.5px;
  color: #222;
  font-weight: 450;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

/* Target input / select */
.field-target {
  width: 100%;
  box-sizing: border-box;
  padding: 7px 10px;
  border: 1px solid #ddd;
  border-radius: 6px;
  font-size: 13.5px;
  color: #222;
  background: #fff;
  outline: none;
  appearance: none;
  -webkit-appearance: none;
  transition: border-color 0.15s;
  background-image: none;
}
select.field-target {
  background-image: url("data:image/svg+xml,%3Csvg width='10' height='6' viewBox='0 0 10 6' fill='none' xmlns='http://www.w3.org/2000/svg'%3E%3Cpath d='M1 1l4 4 4-4' stroke='%23999' stroke-width='1.5' stroke-linecap='round' stroke-linejoin='round'/%3E%3C/svg%3E");
  background-repeat: no-repeat;
  background-position: right 10px center;
  padding-right: 28px;
  cursor: pointer;
}
select.field-target.placeholder { color: #bbb; }
.field-target:focus { border-color: #2563eb; }
.field-target::placeholder { color: #bbb; }

/* Remove button */
.remove-btn {
  background: none;
  border: none;
  cursor: pointer;
  color: #ccc;
  padding: 4px;
  border-radius: 4px;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: color 0.15s;
  flex-shrink: 0;
}
.remove-btn:hover { color: #e55; }

/* ── Add field pair ── */
.add-btn {
  display: flex;
  align-items: center;
  gap: 7px;
  margin: 14px 20px 4px;
  background: none;
  border: none;
  cursor: pointer;
  font-size: 13.5px;
  font-weight: 500;
  color: #2563eb;
  padding: 0;
  transition: opacity 0.15s;
}
.add-btn:hover { opacity: 0.75; }

/* ── Footer ── */
.modal-footer {
  display: flex;
  justify-content: flex-end;
  gap: 10px;
  padding: 16px 20px;
  border-top: 1px solid #ebebeb;
  margin-top: 10px;
}

.btn-cancel {
  background: none;
  border: none;
  cursor: pointer;
  font-size: 13.5px;
  font-weight: 500;
  color: #555;
  padding: 8px 16px;
  border-radius: 6px;
  transition: background 0.15s;
}
.btn-cancel:hover { background: #f3f4f6; }

.btn-save {
  background: #2563eb;
  color: #fff;
  border: none;
  cursor: pointer;
  font-size: 13.5px;
  font-weight: 600;
  padding: 8px 20px;
  border-radius: 6px;
  transition: background 0.15s;
}
.btn-save:hover { background: #1d4ed8; }
</style>
