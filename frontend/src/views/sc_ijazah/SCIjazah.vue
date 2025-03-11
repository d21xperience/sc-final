<template>
    <div class="mt-2">
        <!-- <div class="w-full"> -->
        <!-- <div class=""> -->
        <div class=" ">
            <div class="flex flex-wrap justify-between items-center mb-2">
                <h4 class="font-bold text-xl md:text-2xl">Data Ijazah</h4>
                <div class="md:flex md:items-center md:space-x-2">
                    <h3 class="text-slate-500 md:text-base text-sm">Tahun Pelajaran</h3>
                    <div>
                        <Select v-model="selectedSemester" :options="semester" optionLabel="namaSemester"
                            placeholder="Tahun Pelajaran" class="w-full md:w-52 mr-2" />

                    </div>
                </div>
            </div>
            <div class="mb-2">
                <Toolbar>
                    <template #start>
                        <Button icon="pi pi-plus" severity="success" class="mr-2" @click="openNew"
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
                        <Button label="Proses" icon="pi pi-send" severity="info" @click="exportCSV($event)" />
                    </template>

                </Toolbar>
            </div>

            <Toolbar>
                <template #start>
                    <div class="flex flex-wrap gap-2 items-center justify-between">
                        <div class="flex">
                            <Select v-model="selectedJurusan" :options="jurusan" optionLabel="name" placeholder="Rombel"
                                class="w-full md:w-56 mr-2" />
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
        <DataTable ref="dt" v-model:selection="selectedSiswa" stripedRows size="small" :value="siswa" dataKey="id"
            :paginator="true" :rows="10" :filters="filters"
            paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
            :rowsPerPageOptions="[10, 20, 50]"
            currentPageReportTemplate="Showing {first} to {last} of {totalRecords} products">
            <Column selectionMode="multiple" style="width: 3rem;" :exportable="false"></Column>
            <Column field="name" header="File Ijazah">
                <template #body="slotProps">
                    <Image :src="`https://primefaces.org/cdn/primevue/images/product/${slotProps.data.image}`"
                        :alt="slotProps.data.image" preview image-class="w-16 h-16 rounded-full" />
                </template>
            </Column>
            <Column field="nama" header="Nama" sortable>
                <template #body="slotProps">
                    {{ slotProps.data.pesertaDidik.nmSiswa }}
                </template>
            </Column>
            <Column field="rombonganBelajar.nmKelas" header="Rombel" sortable>
                <template #body="slotProps">
                    {{ slotProps.data.rombonganBelajar.nmKelas }}
                </template>
            </Column>
            <Column field="jk" header="Tahun Pelajaran">
                <template #body="slotProps">
                    {{ slotProps.data.pesertaDidik.jenisKelamin }}
                </template>
            </Column>
            <Column field="jk" header="Tempat,Tgl. Lahir">
                <template #body="slotProps">
                    {{ slotProps.data.pesertaDidik.jenisKelamin }}
                </template>
            </Column>
            <Column field="jk" header="Nama Ortu/Wali">
                <template #body="slotProps">
                    {{ slotProps.data.pesertaDidik.jenisKelamin }}
                </template>
            </Column>
            <Column field="jk" header="NIS">
                <template #body="slotProps">
                    {{ slotProps.data.pesertaDidik.jenisKelamin }}
                </template>
            </Column>
            <Column field="jk" header="NISN">
                <template #body="slotProps">
                    {{ slotProps.data.pesertaDidik.jenisKelamin }}
                </template>
            </Column>
            <Column field="jk" header="Sekolah Asal">
                <template #body="slotProps">
                    {{ slotProps.data.pesertaDidik.jenisKelamin }}
                </template>
            </Column>
            <Column field="jk" header="Tgl. Terbit">
                <template #body="slotProps">
                    {{ slotProps.data.pesertaDidik.jenisKelamin }}
                </template>
            </Column>
            <Column field="jk" header="No.Ijazah">
                <template #body="slotProps">
                    {{ slotProps.data.pesertaDidik.jenisKelamin }}
                </template>
            </Column>
            <!-- Jika SMK/MAK Program Keahlian & Kompetensi Keahlian akan muncul-->
            <div v-if="['smk', 'mak'].includes(bentukPendidikan)">
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
            </div>
        </DataTable>
    </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue';
import { useStore } from "vuex";
const store = useStore();



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



const selectedSiswa = ref();
const siswa = ref(null);
const bentukPendidikan = ref("smk")
const filters = ref({
    'global': { value: null, matchMode: FilterMatchMode.CONTAINS },
});



// =======SEMESTER=============
const selectedSemester = ref();
const semester = ref(null);
const fetchSemester = async () => {
    try {
        const results = await store.dispatch("sekolahService/fetchSemester")
        console.log(results)
        if (results) {
            semester.value = store.getters["sekolahService/getSemester"]
            // Ambil semester terbaru berdasarkan ID terbesar
            selectedSemester.value = semester.value.reduce((latest, current) =>
                current.semesterId > latest.semesterId ? current : latest
            );
        }
    } catch (error) {

    }
}

// ==================================

// ==================================
const confirmDeleteSelected = () => {
    deleteProductsDialog.value = true;
};
// ==================================

const dataLulusan = ref();
const selectedJurusan = ref();
const jurusan = ref([
    { name: 'Teknik Kendaraan Ringan', code: 'TKR' },
    { name: 'Teknik Mesin Sepeda Motor', code: 'TSM' },
    { name: 'Teknik Komputer dan Jaringan', code: 'TKJ' },
    { name: 'Otomatisasi Perkantoran', code: 'OTKP' },
    { name: 'AKuntansi Lembaga', code: 'AKL' }
]);


</script>