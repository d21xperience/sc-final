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




const alamatLengkap = ref({
    alamatJalan: '',
    rt: '',
    rw: '',
    desa: '',
    kec: '',
    kab: '',
    prov: ''
})
// Model Peserta Didik
const mapel = ref({
    pembelajaran_id: "",
    rombongan_belajar_id: "",
    mata_pelajaran_id: "",
    semester_id: "",
    ptk_terdaftar_id: "",
    status_di_kurikulum: "",
    nama_mata_pelajaran: "",
    induk_pembelajaran: "",

});

// Model Peserta Didik Pelengkap
const mapelPelengkap = ref({
    pelengkapSiswaId: '',
    mapelId: '',
    statusDalamKel: '',
    anakKe: '',
    sekolahAsal: '',
    diterimaKelas: '',
    alamatOrtu: '',
    teleponOrtu: '',
    alamatWali: '',
    teleponWali: '',
    fotoSiswa: null
});

// Opsi Dropdown
const selectedjenisKelaminOptions = ref()
const jenisKelaminOptions = ref([
    { label: 'Laki-Laki', value: 'L' },
    { label: 'Perempuan', value: 'P' }
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
    console.log(mapel.value.alamatSiswa);
    console.log('Peserta Didik:', mapel.value);
    console.log('Peserta Didik Pelengkap:', mapelPelengkap.value);

    toast.add({ severity: 'success', summary: 'Sukses', detail: 'Data berhasil disimpan', life: 3000 });
};

// // Handle Upload Foto
// const onUpload = (event) => {
//     const file = event.files[0];
//     mapelPelengkap.value.fotoSiswa = URL.createObjectURL(file);
//     toast.add({ severity: 'info', summary: 'Foto Diunggah', detail: file.name, life: 3000 });
// };

const batal = () => {
    router.push({ name: 'readSiswa' })
}






</script>

<template>
    <div class="container bg-white p-8 rounded-lg shadow-md">
        <!-- <nav class="text-sm text-gray-500 mb-6">
            <a href="#" class="hover:underline">Dashboard</a> &gt; 
            <a href="#" class="hover:underline">Students</a> &gt; 
            <span>Registration Form</span>
        </nav> -->
        <h1 class="text-2xl font-bold mb-6">Form Registrasi Mata pelajaran</h1>

        <section class="mb-8">
            <h2 class="text-xl font-semibold mb-4">Informasi Mata pelajaran</h2>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
                <div>
                    <label class="block text-gray-700" for="nama_mata_pelajaran">Nama Mata pelajaran</label>
                    <InputText v-model="mapel.nama_mata_pelajaran" fluid name="nama_mata_pelajaran"
                        id="nama_mata_pelajaran" placeholder="Masukan nama" />
                </div>
                <div class="w-full">
                    <label class="block text-gray-700">Guru</label>
                    <Select v-model="selectedjenisKelaminOptions" :options="jenisKelaminOptions"
                        placeholder="Pilih Guru" optionLabel="label" class="w-full" />
                </div>
                <div class="w-full">
                    <label class="block text-gray-700">Kelas</label>
                    <Select v-model="selectedjenisKelaminOptions" :options="jenisKelaminOptions"
                        placeholder="Pilih Guru" optionLabel="label" class="w-full" />
                </div>
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
