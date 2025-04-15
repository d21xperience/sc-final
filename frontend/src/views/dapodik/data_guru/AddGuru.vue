<script setup>
import { ref, computed, watch, onMounted, toRaw } from 'vue';

import InputText from 'primevue/inputtext';
import Select from 'primevue/select';

import DatePicker from 'primevue/datepicker';

import Button from 'primevue/button';


import Toast from 'primevue/toast';
import { useToast } from 'primevue/usetoast';
const toast = useToast();
// ================================
// composable
// ================================
import { useSekolahService } from '@/composables/useSekolahService'

const selectedSemester = computed(() => store.getters["sekolahService/getSelectedSemester"])
const schemaName = computed(() => store.getters["sekolahService/getTabeltenant"]?.schemaname)

const { guruList, fetchGuru } = useSekolahService(schemaName, selectedSemester)
// ================================

import Textarea from 'primevue/textarea';

import FileUpload from 'primevue/fileupload';

import InputMask from 'primevue/inputmask';

import Card from 'primevue/card';
import router from '@/router';
import { useFormOptions } from '@/composables/useFormOptions'
const {
    selectedJenisKelamin,
    jenisKelaminOptions,
} = useFormOptions()

import DialogLoading from '@/components/DialogLoading.vue';
import store from '@/store';


const alamatLengkap = ref({
    alamatJalan: '',
    rt: '',
    rw: '',
    desa: '',
    kec: '',
    kab: '',
    prov: ''
})

// Model Peserta Didik Pelengkap
const tabel_ptk_terdaftar = ref({
    ptk_terdaftar_id: '',
    ptk_id: '',
    tahun_ajaran_id: '',
    jenis_keluar_id: '',
    soft_delete: '',
    ptk: {
        ptk_id: "",
        nama: "",
        nip: "",
        jenis_ptk_id: 8,
        jenis_kelamin: "",
        tempat_lahir: "",
        tanggal_lahir: "",
        nuptk: "",
        alamat_jalan: "",
        status_keaktifan_id: 1,
    }
});
import { useRoute } from "vue-router";
const route = useRoute();
const ptkId = route.query.ptkId;
const isEdit = ref(false)
onMounted(async () => {
    if (ptkId) {
        isEdit.value = true
        await fetchGuru(ptkId);
        if (guruList.value && guruList.value.length > 0) {
            const guru = guruList.value[0];
            if (guru.jenisKelamin === "L") {
                selectedJenisKelamin.value = jenisKelaminOptions.value[0]
            } else if (guru.jenisKelamin === "P") {
                selectedJenisKelamin.value = jenisKelaminOptions.value[1]
            }
            tabel_ptk_terdaftar.value.ptk = {
                ptk_id: guru.ptkId || "",
                nama: guru.nama || "",
                nip: guru.nip || "",
                jenis_ptk_id: guru.jenisPtkId || "",
                tempat_lahir: guru.tempatLahir || "",
                tanggal_lahir: guru.tanggalLahir || "",
                nuptk: guru.nuptk || "",
                alamat_jalan: guru.alamatJalan || "",
                status_keaktifan_id: guru.statusKeaktifanId || ""
            };

            tabel_ptk_terdaftar.ptk_id = guru.ptkId; // jangan lupa juga isi yang di luar
        }

    }
})
// Opsi Dropdown
watch(selectedJenisKelamin, () => tabel_ptk_terdaftar.value.ptk.jenis_kelamin = selectedJenisKelamin.value.value)
// Handle Submit Form
// onMounted(async () => await store.dispatch("sekolahService/fetchTabeltenant"))
const isLoading = ref(false)
const submitForm = () => {
    submitted.value = true
    isLoading.value = true
    if (tabel_ptk_terdaftar.value.ptk.nama?.trim()) {
        if (saveGuru()) {
            toast.add({ severity: 'success', summary: 'Sukses', detail: 'Data berhasil disimpan', life: 3000 });
        } else {
            toast.add({ severity: 'error', summary: 'Error', detail: 'Data gagal disimpan', life: 3000 });
        }
    }
    isLoading.value = false

};
const generateUUID = () => crypto.randomUUID();
const saveGuru = async () => {
    try {
        tabel_ptk_terdaftar.value.ptk_terdaftar_id = generateUUID()
        tabel_ptk_terdaftar.value.ptk_id = generateUUID()
        tabel_ptk_terdaftar.value.ptk.ptk_id = tabel_ptk_terdaftar.value.ptk_id
        tabel_ptk_terdaftar.value.tahun_ajaran_id = `${selectedSemester.value?.tahunAjaranId}`
        const payload = {
            schema_name: schemaName.value,
            ptk_terdaftar: [tabel_ptk_terdaftar.value]
        }
        // console.log(payload)
        // simpan ptk ke database
        const response = await store.dispatch("sekolahService/savePTKTerdaftar", payload)
        if (response) {

        }
        // simpan ptk terdaftar ke database

    } catch (error) {

    }

}

const batal = () => {
    router.push({ name: 'readGuru' })
}

const submitted = ref(false)

const tes = (e) => {
    console.log(e)
}


</script>

<template>
    <div class="container bg-white p-8 rounded-lg shadow-md">
        <DialogLoading :modelValue=isLoading @update:modelValue="tes" />
        <h1 class="text-2xl font-bold mb-6">{{ isEdit ? 'Form Edit Guru' : 'Form Tambah Guru' }}</h1>
        <section class="mb-8">
            <h2 class="text-xl font-semibold mb-4">Informasi Guru</h2>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
                <div>
                    <label class="block text-gray-700" for="nmGuru">Nama Lengkap</label>
                    <InputText v-model="tabel_ptk_terdaftar.ptk.nama" fluid name="nmGuru" id="nmGuru"
                        placeholder="Masukan nama" :invalid="submitted && !tabel_ptk_terdaftar.ptk.nama" />
                    <small v-if="submitted && !tabel_ptk_terdaftar.ptk.nama" class="text-red-500">Nama harus
                        diisi.</small>
                </div>
                <div class="w-full">
                    <label class="block text-gray-700">Jenis Kelamin</label>
                    <Select v-model="selectedJenisKelamin" :options="jenisKelaminOptions"
                        placeholder="Pilih jenis kelamin" optionLabel="label" class="w-full"
                        :invalid="submitted && !selectedJenisKelamin" />
                    <small v-if="submitted && !selectedJenisKelamin" class="text-red-500">Jenis kelamin harus
                        diisi.</small>
                </div>
                <div>
                    <div class="md:flex md:space-x-1">

                        <div class="w-3/4">
                            <label class="block text-gray-700" for="tempatLahir">Tpt Lahir</label>
                            <InputText v-model="tabel_ptk_terdaftar.ptk.tempat_lahir" fluid name="tempatLahir"
                                id="tempatLahir" placeholder="Masukan tempat lahir"
                                :invalid="submitted && !tabel_ptk_terdaftar.ptk.tempat_lahir" />
                            <small v-if="submitted && !tabel_ptk_terdaftar.ptk.tempat_lahir"
                                class="text-red-500">Tpt.Lahir harus
                                diisi.</small>

                        </div>
                        <div class="w-1/4">
                            <label class="block text-gray-700">Tgl Lahir</label>
                            <input type="date" placeholder="YYYY-MM-DD"
                                class=" w-full p-2 border border-gray-300 rounded"
                                v-model="tabel_ptk_terdaftar.ptk.tanggal_lahir"
                                :class="{ 'border-red-400': submitted && !tabel_ptk_terdaftar.ptk.tanggal_lahir, 'text-red-400': submitted && !tabel_ptk_terdaftar.ptk.tanggal_lahir }">
                            <small v-if="submitted && !tabel_ptk_terdaftar.ptk.tanggal_lahir"
                                class="text-red-500">Tgl.Lahir harus
                                diisi.</small>

                        </div>
                    </div>
                </div>
                <div>
                    <label class="block text-gray-700" for="alamat-jalan">Alamat Jalan</label>
                    <InputText v-model="tabel_ptk_terdaftar.ptk.alamat_jalan" fluid name="alamat-jalan"
                        id="alamat-jalan" placeholder="Masukan alamat"
                        :invalid="submitted && !tabel_ptk_terdaftar.ptk.alamat_jalan" />
                    <small v-if="submitted && !tabel_ptk_terdaftar.ptk.alamat_jalan" class="text-red-500">Alamat harus
                        diisi.</small>
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
