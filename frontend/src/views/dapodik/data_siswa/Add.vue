<script setup>
import { ref } from 'vue';

import InputText from 'primevue/inputtext';

import Button from 'primevue/button';


import Toast from 'primevue/toast';
import { useToast } from 'primevue/usetoast';
const toast = useToast();



import Textarea from 'primevue/textarea';

import FileUpload from 'primevue/fileupload';

import InputMask from 'primevue/inputmask';

import Card from 'primevue/card';





// Model Peserta Didik
const pesertaDidik = ref({
    pesertaDidikId: '',
    nis: '',
    nisn: '',
    nmSiswa: '',
    tempatLahir: '',
    tanggalLahir: '',
    jenisKelamin: '',
    agama: '',
    alamatSiswa: '',
    teleponSiswa: '',
    diterimaTanggal: '',
    nmAyah: '',
    nmIbu: '',
    pekerjaanAyah: '',
    pekerjaanIbu: '',
    nmWali: '',
    pekerjaanWali: ''
});

// Model Peserta Didik Pelengkap
const pesertaDidikPelengkap = ref({
    pelengkapSiswaId: '',
    pesertaDidikId: '',
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
const jenisKelaminOptions = ref([
    { label: 'Laki-Laki', value: 'L' },
    { label: 'Perempuan', value: 'P' }
]);

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
    console.log('Peserta Didik:', pesertaDidik.value);
    console.log('Peserta Didik Pelengkap:', pesertaDidikPelengkap.value);

    toast.add({ severity: 'success', summary: 'Sukses', detail: 'Data berhasil disimpan', life: 3000 });
};

// Handle Upload Foto
const onUpload = (event) => {
    const file = event.files[0];
    pesertaDidikPelengkap.value.fotoSiswa = URL.createObjectURL(file);
    toast.add({ severity: 'info', summary: 'Foto Diunggah', detail: file.name, life: 3000 });
};
</script>

<template>
    <div class="p-4">
        <Toast />
        <Card>
            <template #title>Form Input Peserta Didik</template>
            <template #content>
                <form @submit.prevent="submitForm">
                    <div class="grid p-fluid">
                        <!-- Nama Siswa -->
                        <div class="col-12 md:col-6">
                            <label for="nmSiswa">Nama Siswa</label>
                            <InputText id="nmSiswa" v-model="pesertaDidik.nmSiswa" required />
                        </div>

                        <!-- Tempat Lahir -->
                        <div class="col-12 md:col-6">
                            <label for="tempatLahir">Tempat Lahir</label>
                            <InputText id="tempatLahir" v-model="pesertaDidik.tempatLahir" required />
                        </div>

                        <!-- Tanggal Lahir -->
                        <div class="col-12 md:col-6">
                            <label for="tanggalLahir">Tanggal Lahir</label>
                            <Calendar id="tanggalLahir" v-model="pesertaDidik.tanggalLahir" dateFormat="dd/mm/yy"
                                required />
                        </div>

                        <!-- Jenis Kelamin -->
                        <div class="col-12 md:col-6">
                            <label for="jenisKelamin">Jenis Kelamin</label>
                            <Dropdown id="jenisKelamin" v-model="pesertaDidik.jenisKelamin"
                                :options="jenisKelaminOptions" optionLabel="label" optionValue="value"
                                placeholder="Pilih Jenis Kelamin" required />
                        </div>

                        <!-- Agama -->
                        <div class="col-12 md:col-6">
                            <label for="agama">Agama</label>
                            <Dropdown id="agama" v-model="pesertaDidik.agama" :options="agamaOptions"
                                optionLabel="label" optionValue="value" placeholder="Pilih Agama" required />
                        </div>

                        <!-- Alamat -->
                        <div class="col-12">
                            <label for="alamatSiswa">Alamat</label>
                            <Textarea id="alamatSiswa" v-model="pesertaDidik.alamatSiswa" rows="2" />
                        </div>

                        <!-- Telepon -->
                        <div class="col-12 md:col-6">
                            <label for="teleponSiswa">Telepon</label>
                            <InputMask id="teleponSiswa" v-model="pesertaDidik.teleponSiswa" mask="9999-9999-9999"
                                required />
                        </div>

                        <!-- Sekolah Asal -->
                        <div class="col-12 md:col-6">
                            <label for="sekolahAsal">Sekolah Asal</label>
                            <InputText id="sekolahAsal" v-model="pesertaDidikPelengkap.sekolahAsal" required />
                        </div>

                        <!-- Foto Siswa -->
                        <div class="col-12">
                            <label>Foto Siswa</label>
                            <FileUpload mode="basic" name="file" accept="image/*" :auto="true" @upload="onUpload" />
                            <div v-if="pesertaDidikPelengkap.fotoSiswa" class="mt-2">
                                <img :src="pesertaDidikPelengkap.fotoSiswa" alt="Foto Siswa"
                                    class="w-32 h-32 border-round" />
                            </div>
                        </div>

                        <!-- Tombol Submit -->
                        <div class="col-12">
                            <Button type="submit" label="Simpan" icon="pi pi-check" class="p-button-success" />
                        </div>
                    </div>
                </form>
            </template>
        </Card>
    </div>
</template>

<style scoped>
label {
    font-weight: bold;
    display: block;
    margin-bottom: 0.5rem;
}
</style>
