<!-- src/components/FormPeserta.vue -->
<template>
    <div class="max-w-4xl mx-auto p-6 bg-white rounded shadow-md">
        <h2 class="text-2xl font-semibold mb-6 text-gray-800">Form Data Peserta</h2>

        <form @submit.prevent="submitForm" class="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div v-for="field in fields" :key="field.name" class="flex flex-col gap-1">
                <label :for="field.name" class="font-medium text-gray-700">
                    {{ field.label }}
                </label>

                <InputText v-if="field.type === 'text'" v-model="form[field.name]" :id="field.name"
                    class="p-inputtext-sm w-full" />

                <DatePicker v-else-if="field.type === 'date'" v-model="form[field.name]" :id="field.name"
                    dateFormat="yy-mm-dd" showIcon class="p-inputtext-sm w-full" />
            </div>

            <div class="col-span-1 md:col-span-2 mt-4">
                <Button label="Submit" class="w-full p-button-primary" />
            </div>
        </form>
    </div>
</template>

<script setup>
import { ref } from 'vue'
import InputText from 'primevue/inputtext'

import DatePicker from 'primevue/datepicker';

import Button from 'primevue/button'

const fields = [
    // { name: 'id', label: 'ID', type: 'text' },
    // { name: 'peserta_didik_id', label: 'Peserta Didik ID', type: 'text' },
    { name: 'program_keahlian', label: 'Program Keahlian', type: 'text' },
    { name: 'paket_keahlian', label: 'Paket Keahlian', type: 'text' },
    // { name: 'sekolah_id', label: 'Sekolah ID', type: 'text' },
    // { name: 'npsn', label: 'NPSN', type: 'text' },
    // { name: 'kabupaten_kota', label: 'Kabupaten/Kota', type: 'text' },
    // { name: 'provinsi', label: 'Provinsi', type: 'text' },
    { name: 'nama', label: 'Nama', type: 'text' },
    { name: 'tempat_lahir', label: 'Tempat Lahir', type: 'text' },
    { name: 'tanggal_lahir', label: 'Tanggal Lahir', type: 'date' },
    { name: 'nis', label: 'NIS', type: 'text' },
    { name: 'nisn', label: 'NISN', type: 'text' },
    // { name: 'nama_ortu_wali', label: 'Nama Ortu/Wali', type: 'text' },
    // { name: 'sekolah_penyelenggara_ujian_us', label: 'Sekolah Ujian US', type: 'text' },
    // { name: 'sekolah_penyelenggara_ujian_un', label: 'Sekolah Ujian UN', type: 'text' },
    // { name: 'asal_sekolah', label: 'Asal Sekolah', type: 'text' },
    // { name: 'nomor_ijazah', label: 'Nomor Ijazah', type: 'text' },
    // { name: 'tempat_ijazah', label: 'Tempat Ijazah', type: 'text' },
    { name: 'tanggal_ijazah', label: 'Tanggal Ijazah', type: 'date' },
    { name: 'no_ijazah', label: 'No Ijazah', type: 'text' },
]

const form = ref({})
fields.forEach(field => (form.value[field.name] = null))

const submitForm = () => {
    console.log('Form submitted:', form.value)
    // Kirim ke API atau emit ke parent
}
</script>

<style scoped>
.field {
    margin-bottom: 1rem;
}
</style>