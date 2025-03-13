<script setup>
// import { defineEmits } from 'vue';
import { useRouter } from 'vue-router';

const props = defineProps({
    title: String,
    date: String,
    value: Number,
    percentage: Number,
    bgColor: { type: String, default: 'bg-white' }, // Warna default
    routePath: { type: String, default: '/' }, // Default ke homepage
});

const emit = defineEmits(['updateData']);
const router = useRouter();

const handleClick = () => {
    emit('updateData', props.value + 1000); // Contoh event emit
    router.push(props.routePath); // Redirect ke halaman yang ditentukan
};
</script>

<template>
    <div :class="[bgColor, 'p-6 rounded-lg shadow-md cursor-pointer transition-all hover:shadow-lg']"
        @click="handleClick">
        <div class="flex justify-between items-center mb-4">
            <h2 class="text-gray-600 text-lg font-semibold">
                {{ title }}
            </h2>
            <span class="text-gray-400 text-sm">
                {{ date }}
            </span>
        </div>
        <div class="text-2xl font-bold text-gray-800 mb-2">
            {{ value.toLocaleString() }}
        </div>
        <div class="flex items-center" :class="percentage > 0 ? 'text-green-500' : 'text-red-500'">
            <i :class="percentage > 0 ? 'fas fa-arrow-up' : 'fas fa-arrow-down'"></i>
            <span class="ml-1">
                {{ percentage }}%
            </span>
            <span class="text-gray-400 ml-2">
                from last week
            </span>
        </div>
    </div>
</template>
