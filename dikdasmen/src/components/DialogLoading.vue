<template>
    <Dialog :visible="modelValue" @update:visible="(newValue) => $emit('update:modelValue', newValue)" :pt="{
        root: { class: '!border-0 !bg-transparent' },
        mask: { class: 'backdrop-blur-sm' }
    }">
        <template #container>
            <div class="flex flex-col px-8 py-8 gap-6 rounded-2xl">
                <ProgressSpinner style="width: 100px; height: 100px" strokeWidth="8" fill="transparent"
                    animationDuration=".5s" aria-label="Custom ProgressSpinner" />
                <slot> Sedang mengambil data </slot>
            </div>
        </template>
    </Dialog>
</template>

<script setup>
import Dialog from 'primevue/dialog';

import ProgressSpinner from 'primevue/progressspinner';

defineProps({
    modelValue: Boolean, // Model untuk v-model
});
defineEmits(["update:modelValue"]); // Mengizinkan emit untuk update
</script>
