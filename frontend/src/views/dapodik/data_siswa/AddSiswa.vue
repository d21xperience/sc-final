<script setup>
import { ref, computed } from 'vue';

import InputText from 'primevue/inputtext';
import Select from 'primevue/select';

import DatePicker from 'primevue/datepicker';

import Button from 'primevue/button';


import Toast from 'primevue/toast';
import { useToast } from 'primevue/usetoast';
const toast = useToast();

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
const pesertaDidik = ref({
    pesertaDidikId: '',
    nis: '',
    nisn: '',
    nmSiswa: '',
    tempatLahir: '',
    tanggalLahir: '',
    jenisKelamin: '',
    agama: '',
    alamatSiswa: computed(() => `${alamatLengkap.value.alamatJalan} RT.${alamatLengkap.value.rt} RW.${alamatLengkap.value.rw} Desa ${alamatLengkap.value.desa} Kec. ${alamatLengkap.value.kec} Kab. ${alamatLengkap.value.kab} Prov. ${alamatLengkap.value.prov}`),
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
    console.log(pesertaDidik.value.alamatSiswa);
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
        <h1 class="text-2xl font-bold mb-6">Form Registrasi Siswa</h1>

        <section class="mb-8">
            <h2 class="text-xl font-semibold mb-4">Informasi Siswa</h2>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
                <div>
                    <label class="block text-gray-700" for="nmSiswa">Nama</label>
                    <InputText v-model="pesertaDidik.nmSiswa" fluid name="nmSiswa" id="nmSiswa"
                        placeholder="Masukan nama" />
                </div>
                <div class="w-full">
                    <label class="block text-gray-700">Jenis Kelamin</label>
                    <Select v-model="selectedjenisKelaminOptions" :options="jenisKelaminOptions"
                        placeholder="Pilih jenis kelamin" optionLabel="label" class="w-full" />
                </div>
                <div>
                    <div class="md:flex md:space-x-1">

                        <div class="w-full">
                            <label class="block text-gray-700" for="tempatLahir">Tpt Lahir</label>
                            <InputText v-model="pesertaDidik.tempatLahir" fluid name="tempatLahir" id="tempatLahir"
                                placeholder="Masukan tempat lahir" class=" w-full md:w-64" />
                        </div>
                        <div>
                            <label class="block text-gray-700">Tgl Lahir</label>
                            <input type="date" placeholder="YYYY-MM-DD"
                                class=" w-full p-2 border border-gray-300 rounded" v-model="pesertaDidik.tanggalLahir">
                        </div>
                    </div>
                </div>

                <div>
                    <label class="block text-gray-700">Agama</label>
                    <Select v-model="selectedAgamaOptions" :options="agamaOptions" placeholder="Pilih Agama"
                        optionLabel="label" fluid class="w-full" />
                </div>
                <div>
                    <label class="block text-gray-700" for="nis">NIS</label>
                    <InputText v-model="pesertaDidik.nis" fluid name="nis" id="nis" placeholder="Masukan NIS" />
                </div>
                <div>
                    <label class="block text-gray-700" for="nisn">NISN</label>
                    <InputText v-model="pesertaDidik.nisn" fluid name="nisn" id="nisn" placeholder="Masukan NISN" />
                </div>
            </div>

            <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
                <div>
                    <label class="block text-gray-700" for="nmSiswa">Alamat Jalan</label>
                    <InputText v-model="alamatLengkap.alamatJalan" fluid name="nmSiswa" id="nmSiswa"
                        placeholder="Masukan nama" />
                </div>
                <div class="flex space-x-1">
                    <div class="w-1/2">
                        <label class="block text-gray-700" for="rt">RT</label>
                        <InputText v-model="alamatLengkap.rt" fluid name="rt" id="rt" placeholder="Masukan RT" />
                    </div>
                    <div class="w-1/2">
                        <label class="block text-gray-700" for="rw">RW</label>
                        <InputText v-model="alamatLengkap.rw" fluid name="rw" id="rw" placeholder="Masukan RW" />
                    </div>
                </div>
                <div>
                    <label class="block text-gray-700" for="prov">Prov.</label>
                    <InputText v-model="alamatLengkap.prov" fluid name="prov" id="prov" placeholder="Masukan nama" />
                </div>
                <div>
                    <label class="block text-gray-700" for="kab">Kab</label>
                    <InputText v-model="alamatLengkap.kab" fluid name="kab" id="kab" placeholder="Masukan nama" />
                </div>
                <div>
                    <label class="block text-gray-700" for="kec">Kecamatan</label>
                    <InputText v-model="alamatLengkap.kec" fluid name="kec" id="kec"
                        placeholder="Masukan nama kecamatan" />
                </div>
                <div>
                    <label class="block text-gray-700" for="desa">Desa</label>
                    <InputText v-model="alamatLengkap.desa" fluid name="desa" id="desa"
                        placeholder="Masukan nama desa" />
                </div>



                <!-- <div class="mb-4">
                    <label class="block text-gray-700">Address</label>
                    <textarea placeholder="Enter student's address"
                        class="w-full p-2 border border-gray-300 rounded"></textarea>
                </div> -->

            </div>
            <div class="mb-4">
                <label class="block text-gray-700">Phone Number</label>
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

        <section class="mb-8">
            <h2 class="text-xl font-semibold mb-4">Informasi Orang tua</h2>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
                <div>
                    <label class="block text-gray-700" for="nmAyah">Nama Ayah Kandung</label>
                    <InputText v-model="pesertaDidik.nmAyah" fluid name="nmAyah" id="nmAyah"
                        placeholder="Masukan nama" />
                </div>
                <div>
                    <label class="block text-gray-700">Pekerjaan Ayah</label>
                    <input type="text" placeholder="Enter father's occupation"
                        class="w-full p-2 border border-gray-300 rounded">
                </div>
                <div>
                    <label class="block text-gray-700">Nama Ibu Kandung</label>
                    <input type="text" placeholder="Enter mother's name"
                        class="w-full p-2 border border-gray-300 rounded">
                </div>
                <div>
                    <label class="block text-gray-700">Pekerjaan Ibu</label>
                    <input type="text" placeholder="Enter mother's occupation"
                        class="w-full p-2 border border-gray-300 rounded">
                </div>
            </div>
            <div class="mb-4">
                <label class="block text-gray-700">Alamat Orang tua</label>
                <textarea placeholder="Enter parents' address (if different from student)"
                    class="w-full p-2 border border-gray-300 rounded"></textarea>
            </div>
            <div>
                <label class="block text-gray-700">No.Tlp Ortu</label>
                <div class="relative">
                    <input type="text" placeholder="Enter parents' phone number"
                        class="w-full p-2 border border-gray-300 rounded">
                    <i class="fas fa-phone-alt absolute right-3 top-3 text-gray-400"></i>
                </div>
            </div>
        </section>

        <section class="mb-8">
            <h2 class="text-xl font-semibold mb-4">Informasi Wali (Optional)</h2>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
                <div>
                    <label class="block text-gray-700">Nama Wali</label>
                    <input type="text" placeholder="Enter guardian's name (if applicable)"
                        class="w-full p-2 border border-gray-300 rounded">
                </div>
                <div>
                    <label class="block text-gray-700">Pekerjaan wali</label>
                    <input type="text" placeholder="Enter guardian's occupation"
                        class="w-full p-2 border border-gray-300 rounded">
                </div>
            </div>
            <div class="mb-4">
                <label class="block text-gray-700">Alamat Wali</label>
                <textarea placeholder="Enter guardian's address"
                    class="w-full p-2 border border-gray-300 rounded"></textarea>
            </div>
            <div>
                <label class="block text-gray-700">No.Tlp. Wali</label>
                <div class="relative">
                    <input type="text" placeholder="Enter guardian's phone number"
                        class="w-full p-2 border border-gray-300 rounded">
                    <i class="fas fa-phone-alt absolute right-3 top-3 text-gray-400"></i>
                </div>
            </div>
        </section>

        <section class="mb-8">
            <h2 class="text-xl font-semibold mb-4">Informasi Tambahan</h2>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
                <div>
                    <label class="block text-gray-700">Status dalam Keluarga</label>
                    <select class="w-full p-2 border border-gray-300 rounded">
                        <option>Select status</option>
                    </select>
                </div>
                <div>
                    <label class="block text-gray-700">Anak Ke-</label>
                    <input type="text" placeholder="Enter child number (e.g., 1, 2)"
                        class="w-full p-2 border border-gray-300 rounded">
                </div>
                <div>
                    <label class="block text-gray-700">Asal Sekolah</label>
                    <input type="text" placeholder="Enter previous school name"
                        class="w-full p-2 border border-gray-300 rounded">
                </div>
                <div>
                    <label class="block text-gray-700">Diterima di kelas</label>
                    <input type="text" placeholder="Enter class level"
                        class="w-full p-2 border border-gray-300 rounded">
                </div>
            </div>
            <div>
                <label class="block text-gray-700">Student Photo</label>
                <div class="relative">
                    <input type="file" class="w-full p-2 border border-gray-300 rounded">
                    <i class="fas fa-upload absolute right-3 top-3 text-gray-400"></i>
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
