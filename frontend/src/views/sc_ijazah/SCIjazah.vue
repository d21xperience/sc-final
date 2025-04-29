<template>
    <div class="mt-2">
        <!-- <div class="w-full"> -->
        <!-- <div class=""> -->
        <div class=" ">
            <div class="flex flex-wrap justify-between items-center mb-2">
                <h4 class="font-bold text-xl md:text-2xl">Data Ijazah</h4>
                <!-- <div class="md:flex md:items-center md:space-x-2">
                    <h3 class="text-slate-500 md:text-base text-sm">Tahun Lulus</h3>
                    <div>
                        <Select v-model="selectedTahunAjaran" :options="tahunAjaranOptions" optionLabel="label"
                            optionValue="value" placeholder="Tahun Pelajaran" class="w-full md:w-52 mr-2" />

                    </div>
                </div> -->
            </div>
            <div class="mb-2">
                <Toolbar>
                    <template #start>
                        <!-- <Button icon="pi pi-plus" severity="success" class="mr-2" @click="visible = true"
                            v-tooltip.bottom="'Tambah data'" /> -->
                        <Button icon="pi pi-pencil" severity="warn" @click="visible = true"
                            :disabled="!selectedSiswa || !selectedSiswa.length || selectedSiswa.length > 1" class="mr-2"
                            v-tooltip.bottom="'Edit data'" />
                        <Button icon="pi pi-trash" severity="danger" class="mr-2" @click="confirmDeleteSelected"
                            :disabled="!selectedSiswa || !selectedSiswa.length" />

                        <!-- <Button label="Lulus" severity="warn" class="mr-2" @click="dialogStatus = true"
                                            :disabled="!dataLulusan || !dataLulusan.length" />
                                        <Button label="Naik" severity="warn" class="mr-2" @click="openNew"
                                            :disabled="!dataLulusan || !dataLulusan.length" /> -->
                    </template>
                    <template #end>
                        <Button label="Import" icon="pi pi-download" severity="warn" @click="dialogImport = true"
                            class="mr-2" />
                        <Button label="Export" icon="pi pi-upload" severity="help" @click="exportCSV($event)"
                            class="mr-2" />
                        <!-- <Button label="Proses" icon="pi pi-send" severity="info" @click="exportCSV($event)" /> -->
                        <IssueDegreeButton :degreeData="degreeData" :sekolah="sekolah" :ipfsUrl="ipfsUrl"
                            :transcript="transcript" :contract="contract" class="bg-blue-600 p-3 rounded-lg text-white"
                            :disabled="!selectedSiswa"
                            :class="{ 'bg-slate-500': !selectedSiswa || selectedSiswa.length === 0 || selectedSiswa.length > 2 }" />
                    </template>

                </Toolbar>
            </div>

            <Toolbar>
                <template #start>
                    <div class="flex flex-wrap gap-2 items-center justify-between">
                        <div class="flex">
                            <!-- <Select v-model="selectedJurusan" :options="jurusan" optionLabel="name" placeholder="Rombel"
                                class="w-full md:w-56 mr-2" /> -->
                            <!-- <Select v-model="selectedJurusan" :options="jurusan" optionLabel="name"
                                                placeholder="Tingkat" class="mr-2" /> -->
                        </div>
                    </div>
                </template>
                <template #end>
                    <IconField>
                        <InputIcon>
                            <i class="pi pi-search" />
                        </InputIcon>
                        <InputText v-model="filters['global'].value" placeholder="Search..." />
                    </IconField>
                </template>
            </Toolbar>
        </div>
        <DataTable ref="dt" v-model:selection="selectedSiswa" stripedRows size="small" :value="siswa" scrollable
            scrollHeight="29rem" dataKey="anggotaRombelId" :paginator="true" :rows="10" :filters="filters"
            paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
            :rowsPerPageOptions="[10, 20, 50]"
            currentPageReportTemplate="Showing {first} to {last} of {totalRecords} products">
            <Column selectionMode="multiple" style="width: 3rem;" :exportable="false"></Column>
            <!-- <Column field="anggotaKelas.nmKelas" header="Kelas"></Column> -->
            <Column field="pesertaDidik.nis" header="NIS"></Column>
            <Column field="pesertaDidik.nisn" header="NISN"></Column>
            <Column field="pesertaDidik.nmSiswa" header="Nama" sortable></Column>
            <Column field="pesertaDidik.jenisKelamin" header="JK"></Column>
            <Column field="pesertaDidik.tempatLahir" header="Tpt. Lahir"></Column>
            <Column field="pesertaDidik.nmAyah" header="Nama Wali"></Column>
            <Column field="cidUrl" header="CID Ijazah"></Column>
            <Column field="nomorIjazah" header="No. Ijazah"></Column>
            <Column field="" header="Status">
                <template #body>
                    belum terkirim
                </template>
            </Column>
            <Column field="" header="Kelengkapan">
                <template #body="{ data: siswaRow }">
                    <div class="card flex flex-col items-center gap-6">
                        <p class="text-sm text-red-500">Belum lengkap</p>
                        <!-- <FileUpload mode="basic" customUpload auto accept="image/*" chooseLabel="Pilih File"
                            @select="(e) => onFileSelect(e, siswaRow)" @uploader="(e) => onFileUpload(e, siswaRow)"
                            class="p-button-outlined" />
                        <img v-if="siswaRow.preview" :src="siswaRow.preview" alt="Image"
                            class="shadow-md rounded-xl w-full sm:w-64" style="filter: grayscale(100%)" /> -->
                    </div>
                </template>

            </Column>
        </DataTable>
        <Dialog v-model:visible="visible" modal header="Data ijazah" :style="{ width: '100rem', height: '100rem' }">
            <DialogIjazah />
        </Dialog>
    </div>
</template>

<script setup>
import { ref, onMounted, watch, computed } from 'vue';
import { useStore } from "vuex";
const store = useStore();

import FileUpload from 'primevue/fileupload';

import AddIjazah from './AddIjazah.vue';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Image from 'primevue/image';

import Button from 'primevue/button';
import Dialog from 'primevue/dialog';
import Toolbar from 'primevue/toolbar';
import { FilterMatchMode } from '@primevue/core/api';
import { useToast } from 'primevue/usetoast';
import InputText from 'primevue/inputtext';
import IconField from 'primevue/iconfield';
import InputIcon from 'primevue/inputicon';
import IssueDegreeForm from './IssueDegreeForm.vue';
import IssueDegreeButton from '@/components/IssueDegreeButton.vue';
import { useUtils } from '@/composables/useUtils'
const { formatterDateID } = useUtils()

const visible = ref(false)

const selectedSiswa = ref();
const siswa = ref();
const bentukPendidikan = ref("smk")
const filters = ref({
    'global': { value: null, matchMode: FilterMatchMode.CONTAINS },
});

// ================================
const selectedTahunAjaran = computed(() => {
    fetchSiswaLulus()
    return store.getters["sekolahService/getSelectedTahunAjaran"]
})
const schemaname = computed(() => store.getters["sekolahService/getTabeltenant"]?.schemaname)

const fetchSiswaLulus = async () => {
    try {
        let payload = {
            schemaname: await store.getters["sekolahService/getTabeltenant"]?.schemaname,
            tahun_ajaran_id: 2023//`${selectedTahunAjaran.value}`, 
        }
        // console.log("fetchSiswaLulus", payload)
        const results = await store.dispatch("sekolahService/fetchProsesIjazah", payload)
        // console.log(results)
        siswa.value = results

        // if (results) {
        //     // console.log(results.anggotaKelas)
        //     siswa.value = results.kenaikan
        // }
        // store.commit("sekolahService/SET_SELECTEDSEMESTER", selectedTahunAjaran.value)

    } catch (error) {
        console.log(error)
    }
}
// ==================================
const confirmDeleteSelected = () => {
    deleteProductsDialog.value = true;
};
// ==================================
const scData = ref({
    degreeData: null,
    sekolah: null,
    ipfsUrl: null,
    transcript: null
})
const dataLulusan = ref();
// const selectedJurusan = ref();
import { ethers } from 'ethers';
import DialogIjazah from '@/components/DialogIjazah.vue';
// Dummy data (bisa kamu ambil dari API atau input form)
const degreeData = ref({
    nama: "",
    nisn: "",
    nik: "3211142109820004",
    tahun_lulus: 2024,
    major: "Rekayasa Perangkat Lunak"
});
const sekolah = ref("SMK PASUNDAN JATINANGOR");

const ipfsUrl = ref("https://ipfs.io/ipfs/Qm...examplehash");
const transcript = ref({
    subjects: ["Matematika", "Pemrograman", "Basis Data"],
    grades: [85, 90, 88]
});
const contract = null;

watch(selectedSiswa, (newVal) => {
    if (newVal.length === 1) {
        console.log(newVal[0].pesertaDidik.nmSiswa)
        degreeData.value.nama = newVal[0].pesertaDidik.nmSiswa
        degreeData.value.nisn = newVal[0].pesertaDidik.nisn
        degreeData.value.tahun_lulus = 2023
        // scData.value.degreeData = { ...newVal[0].pesertaDidik }; // Salin object pertama
        // scData.value.degreeData = {
        //     nama: newVal.value.pesertaDidik.nmSiswa,
        //     nisn: newVal.value.pesertaDidik.nisn,
        //     tahun_lulus: 2023,
        //     nik: "",
        //     major: newVal.value.programKeahlian
        // }
        // scData.value.transcript = transcript.value
        // scData.value.sekolah = sekolah.value
        // scData.value.ipfsUrl = ipfsUrl
    }
});




onMounted(async () => {
    await fetchSiswaLulus()
    // await initContract();
    // await fetchSemester()
});
// ==================================
// watch(selectedTahunAjaran, async () => {
//     // Panggil data untuk mengumpulkan siswa
//     try {
//         let payload = {
//             schemaname: schemaname.value,
//             semester_id: selectedTahunAjaran.value,
//             tipe_kenaikan: 14
//         }
//         const results = await store.dispatch("sekolahService/fetchProsesIjazah", payload)
//         if (results) {
//             // console.log(results.anggotaKelas)
//             siswa.value = results.kenaikan
//         }
//         // store.commit("sekolahService/SET_SELECTEDSEMESTER", selectedTahunAjaran.value)

//     } catch (error) {
//         console.log(error)
//     }
// })
// const initContract = async () => {
//     try {
//         if (window.ethereum) {
//             await window.ethereum.request({ method: 'eth_requestAccounts' });
//             const provider = new ethers.BrowserProvider(window.ethereum);
//             const signer = await provider.getSigner();
//             const contractAddress = '0xYourContractAddressHere'; // Ganti dengan alamat kontrakmu
//             contract.value = new ethers.Contract(contractAddress, DegreeContractABI, signer);
//         } else {
//             alert('Metamask tidak ditemukan. Harap instal terlebih dahulu.');
//         }
//     } catch (error) {
//         console.error('Gagal memuat kontrak:', error);
//     }
// };

const src = ref(null);

const onFileSelect = (event, siswaRow) => {
    const file = event.files[0];
    if (file) {
        const reader = new FileReader();
        reader.onload = (e) => {
            siswaRow.preview = e.target.result; // Set ke row-nya!
        };
        reader.readAsDataURL(file);
    }
};
const editSelectedSiswa = () => {

}
</script>