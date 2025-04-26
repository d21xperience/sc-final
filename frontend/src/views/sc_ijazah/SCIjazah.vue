<template>
    <div class="mt-2">
        <!-- <div class="w-full"> -->
        <!-- <div class=""> -->
        <div class=" ">
            <div class="flex flex-wrap justify-between items-center mb-2">
                <h4 class="font-bold text-xl md:text-2xl">Data Ijazah</h4>
                <div class="md:flex md:items-center md:space-x-2">
                    <h3 class="text-slate-500 md:text-base text-sm">Tahun Lulus</h3>
                    <div>
                        <Select v-model="selectedTahunAjaran" :options="tahunAjaranOptions" optionLabel="label"
                            placeholder="Tahun Pelajaran" class="w-full md:w-52 mr-2" />

                    </div>
                </div>
            </div>
            <div class="mb-2">
                <Toolbar>
                    <template #start>
                        <Button icon="pi pi-plus" severity="success" class="mr-2" @click="visible = true"
                            v-tooltip.bottom="'Tambah data'" />
                        <Button icon="pi pi-pencil" severity="warn" @click="confirmDeleteSelected"
                            :disabled="!dataLulusan || !dataLulusan.length || dataLulusan.length > 2" class="mr-2"
                            v-tooltip.bottom="'Edit data'" />
                        <Button icon="pi pi-trash" severity="danger" class="mr-2" @click="confirmDeleteSelected"
                            :disabled="!dataLulusan || !dataLulusan.length" />

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
        <!-- </div> -->
        <!-- </div> -->
        <DataTable ref="dt" v-model:selection="selectedSiswa" stripedRows size="small" :value="siswa"
            dataKey="anggotaRombelId" :paginator="true" :rows="10" :filters="filters"
            paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
            :rowsPerPageOptions="[10, 20, 50]"
            currentPageReportTemplate="Showing {first} to {last} of {totalRecords} products">
            <Column selectionMode="multiple" style="width: 3rem;" :exportable="false"></Column>
            <!-- <Column field="name" header="File Ijazah">
                <template #body="slotProps">
                    <Image :src="`https://primefaces.org/cdn/primevue/images/product/${slotProps.data.image}`"
                        :alt="slotProps.data.image" preview image-class="w-16 h-16 rounded-full" />
                </template>
            </Column> -->
            <Column field="anggotaKelas.nmSiswa" header="Nama" sortable></Column>
            <Column field="anggotaKelas.pesertaDidik.jenisKelamin" header="JK"></Column>
            <Column field="anggotaKelas.pesertaDidik.nis" header="NIS"></Column>
            <Column field="anggotaKelas.pesertaDidik.nisn" header="NISN"></Column>
            <Column field="anggotaKelas.nmKelas" header="Rombel"></Column>
            <Column field="anggotaKelas.pesertaDidik.tempatLahir" header="Tpt. Lahir"></Column>
            <!-- <Column field="" header="Tgl. Lahir">
                <template #body="slotProps">
                    {{ formatterDateID(slotProps.data.pesertaDidik.tanggalLahir) }}
                </template>

            </Column>
            <Column field="jk" header="Nama Ortu/Wali"></Column>
            <Column field="jk" header="Sekolah Asal">
                <template #body="slotProps">
                    {{ "SMK Pasundan Jatinangor" }}
                </template>
            </Column>
            <Column field="jk" header="Tgl. Terbit">
            </Column>
            <Column field="jk" header="No.Ijazah"> 

            </Column>-->
            <!-- Jika SMK/MAK Program Keahlian & Kompetensi Keahlian akan muncul-->
            <!-- <div v-if="['smk', 'mak'].includes(bentukPendidikan)">
                <Column field="jk" header="Prog.Keahlian">
                    <template #body="slotProps">
                        {{ slotProps.data.pesertaDidik.jenisKelamin }}
                    </template>
                </Column>
                <Column field="jk" header="Komp.Keahlian">
                    <template #body="slotProps">
                        {{ slotProps.data.pesertaDidik.jenisKelamin }}
                    </template>
                </Column>
            </div> -->

        </DataTable>
        <Dialog v-model:visible="visible" modal header="Tambah data ijazah">
            <IssueDegreeForm />

        </Dialog>

    </div>
</template>

<script setup>
import { ref, onMounted, watch, computed } from 'vue';
import { useStore } from "vuex";
const store = useStore();
import AddIjazah from './AddIjazah.vue';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Image from 'primevue/image';

import Button from 'primevue/button';
import Select from 'primevue/select';





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
const semester = ref()
const selectedTahunAjaran = ref()
const schemaname = computed(() => store.getters["sekolahService/getTabeltenant"]?.schemaname)

// ================================
const tahunAjaranOptions = ref()
const fetchSemester = async () => {
    try {
        semester.value = await store.getters["sekolahService/getSemester"]
        if (!semester.value) {
            semester.value = await store.dispatch("sekolahService/fetchSemester")
        }
        tahunAjaranOptions.value = getTahunAjaran(semester.value)
        // Ambil tahun ajaran terbaru berdasarkan tahun terbesar
        selectedTahunAjaran.value = tahunAjaranOptions.value.reduce((latest, current) =>
            current.tahunAjaranId > latest.tahunAjaranId ? current : latest
        );
    } catch (error) {
        console.log(error)
    }
}
watch(selectedTahunAjaran, async () => {
    // Panggil data untuk mengumpulkan siswa
    try {
        let payload = {
            schemaname: schemaname.value,
            semester_id: selectedTahunAjaran.value.value,
            tipe_kenaikan: 12
        }
        const results = await store.dispatch("sekolahService/fetchProsesIjazah", payload)
        if (results) {
            // console.log(results.anggotaKelas)
            siswa.value = results.kenaikan
        }
    } catch (error) {
        console.log(error)
    }
})

// ==================================
const confirmDeleteSelected = () => {
    deleteProductsDialog.value = true;
};
// ==================================

const dataLulusan = ref();
// const selectedJurusan = ref();
import { ethers } from 'ethers';
// Dummy data (bisa kamu ambil dari API atau input form)
const degreeData = ref({
    nama: "Andi Wijaya",
    nisn: "1234567890",
    nik: "3211142109820004",
    tahun_lulus: 2024,
    major: "Rekayasa Perangkat Lunak"
});
const sekolah = ref("SMK PASUNDAN JATINANGOR");

const ipfsUrl = ref("https://ipfs.io/ipfs/Qm...examplehash");
const transcript = ref({
    mapel: ["Matematika", "Pemrograman", "Basis Data"],
    nilai: [85, 90, 88]
});
const contract = null;

watch(selectedSiswa, (newVal) => {
    if (newVal.length === 1) {
        degreeData.value = { ...newVal[0].pesertaDidik }; // Salin object pertama
    }
});




onMounted(async () => {
    // await initContract();
    await fetchSemester()
});
// ==================================
const getTahunAjaran = (semesterArray) => {
    const unique = new Set();
    // console.log(semesterArray)
    return semesterArray
        .filter(item => {
            if (!unique.has(item.tahunAjaranId)) {
                unique.add(item.tahunAjaranId);
                return true;
            }
            return false;
        })
        .map(item => ({
            label: item.tahunAjaranId,
            value: item.tahunAjaranId + "2"
        }));
};
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
</script>