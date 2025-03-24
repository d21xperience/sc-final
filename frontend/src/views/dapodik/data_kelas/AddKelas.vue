<script setup>
import { ref, computed } from 'vue';

import InputText from 'primevue/inputtext';
import Select from 'primevue/select';

import DatePicker from 'primevue/datepicker';

import Button from 'primevue/button';


import Toast from 'primevue/toast';
import { useToast } from 'primevue/usetoast';
const toast = useToast();



import Textarea from 'primevue/textarea';

import FileUpload from 'primevue/fileupload';

import InputMask from 'primevue/inputmask';

import Card from 'primevue/card';
import router from '@/router';

// Model Kelas
const rombel = ref({
    rombongan_belajar_id: '',
    sekolah_id: '',
    semester_id: '',
    jurusan_id: '',
    ptk_id: '',
    nm_kelas: '',
    tingkat_pendidikan_id: '',
    jenis_rombel: '',
    nama_jurusan_sp: '',
    jurusan_sp_id: '',
    kurikulum_id: ''
});



// Opsi Dropdown
const selectedTingkat = ref()
const jenisKelaminOptions = ref([
    { label: '1', value: '1' },
    { label: '2', value: '2' },
    { label: '3', value: '3' }
]);
const selectedAgamaOptions = ref()
const agamaOptions = ref([
    { label: 'Islam', value: 'Islam' },
    { label: 'Kristen', value: 'Kristen' },
    { label: 'Katolik', value: 'Katolik' },
    { label: 'Hindu', value: 'Hindu' },
    { label: 'Buddha', value: 'Buddha' },
    { label: 'Konghucu', value: 'Konghucu' }
]);

// Handle Submit Form
const submitForm = () => {
    console.log(rombel.value.alamatSiswa);
    console.log('Peserta Didik:', rombel.value);
    console.log('Peserta Didik Pelengkap:', rombelPelengkap.value);

    toast.add({ severity: 'success', summary: 'Sukses', detail: 'Data berhasil disimpan', life: 3000 });
};

// Handle Upload Foto
const onUpload = (event) => {
    const file = event.files[0];
    rombelPelengkap.value.fotoSiswa = URL.createObjectURL(file);
    toast.add({ severity: 'info', summary: 'Foto Diunggah', detail: file.name, life: 3000 });
};

const batal = () => {
    router.push({ name: 'readSiswa' })
}






</script>

<template>
    <div class="container bg-white p-8 rounded-lg shadow-md">
        <h1 class="text-2xl font-bold mb-6">Form Tambah Kelas</h1>

        <section class="mb-8">
            <h2 class="text-xl font-semibold mb-4">Informasi Kelas</h2>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
                <div>
                    <label class="block text-gray-700" for="nmSiswa">Nama</label>
                    <InputText v-model="rombel.nm_kelas" fluid name="nmSiswa" id="nmSiswa" placeholder="Masukan nama" />
                </div>
                <div class="w-full">
                    <label class="block text-gray-700">Tingkat</label>
                    <Select v-model="selectedTingkat" :options="jenisKelaminOptions" placeholder="Pilih jenis kelamin"
                        optionLabel="label" class="w-full" />
                </div>

            </div>

            <div class="mb-4">
                <label class="block text-gray-700">Wali Kelas</label>
                <div class="relative">
                    <input type="text" placeholder="Enter student's phone number"
                        class="w-full p-2 border border-gray-300 rounded">
                    <i class="fas fa-phone-alt absolute right-3 top-3 text-gray-400"></i>
                </div>
            </div>
            <div class="mb-4">
                <label class="block text-gray-700">Email</label>
                <div class="relative">
                    <input type="text" placeholder="Enter student's phone number"
                        class="w-full p-2 border border-gray-300 rounded">
                    <i class="fas fa-phone-alt absolute right-3 top-3 text-gray-400"></i>
                </div>
            </div>
            <div>
                <label class="block text-gray-700">Admission Date</label>
                <input type="date" placeholder="YYYY-MM-DD" class="w-full p-2 border border-gray-300 rounded">
            </div>
        </section>

        <div class="flex justify-end space-x-4">
            <button class="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600"
                @click="submitForm">Simpan</button>
            <button class="bg-gray-300 text-gray-700 px-4 py-2 rounded hover:bg-gray-400" @click="batal">Batal</button>
        </div>
    </div>

    <Toast />
</template>

<style scoped>
label {
    font-weight: bold;
    display: block;
    margin-bottom: 0.5rem;
}

select option:disabled {
    color: gray;
    font-weight: bold;
}
</style>
