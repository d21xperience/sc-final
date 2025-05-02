<template>
    <div class="container">
        <div class="card">
            <div class=" bg-white">
                <div class="my-2 ">
                    <div class=" ">
                        <h2 class="text-xl mb-2">Data Kelas</h2>
                        <div class="mb-2">
                            <Toolbar>
                                <template #start>
                                    <Button icon="pi pi-plus" severity="success" class="mr-2" @click="openNew"
                                        v-tooltip.bottom="'Tambah data'" />
                                    <Button icon="pi pi-pencil" severity="warn" @click="editKelas(selectedKelas)"
                                        :disabled="!selectedKelas || !selectedKelas.length || selectedKelas.length > 1"
                                        class="mr-2" v-tooltip.bottom="'Edit data'" />
                                    <Button icon="pi pi-trash" severity="danger" class="mr-2"
                                        @click="confirmDeleteSelected"
                                        :disabled="!selectedKelas || !selectedKelas.length"
                                        v-tooltip.bottom="'Hapus data'" />
                                    <div v-show="selectedSemester.semester === 2">
                                        <Button label="Lulus" severity="help" class="mr-2 text-sm"
                                            @click="isDialogKelulusan = true" :disabled="!isLulus"
                                            v-tooltip.bottom="'Luluskan siswa'" />
                                        <Button label="Naik" severity="success" class="mr-2 text-sm"
                                            @click="isDialogKenaikan = true" :disabled="!isNaik"
                                            v-tooltip.bottom="'Naikan siswa'" />
                                    </div>
                                </template>
                                <template #end>
                                    <Button label="Export" icon="pi pi-upload" severity="help"
                                        @click="exportCSV($event)" class="mr-2" />
                                </template>
                            </Toolbar>
                        </div>

                        <Toolbar>
                            <template #start>
                                <div class="flex flex-wrap gap-2 items-center justify-between">
                                    <div class="flex">
                                        <Select v-model="filters['tingkatPendidikanId'].value"
                                            :options="tingkatPendidikanOptions" optionLabel="nama" optionValue="kode"
                                            placeholder="Tingkat" class="w-full md:w-56 mr-2" checkmark show-clear />
                                    </div>
                                </div>
                            </template>
                        </Toolbar>
                    </div>
                </div>
            </div>

            <!-- <Skeleton class="mb-2" borderRadius="16px"></Skeleton> -->
            <DataTable ref="dt" v-model:selection="selectedKelas" stripedRows size="small" :value="kelasList" scrollable
                scrollHeight="400px" dataKey="rombonganBelajarId" :paginator="true" :rows="10" :filters="filters"
                tableStyle="min-width: 50rem"
                paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
                :rowsPerPageOptions="[10, 20, 30]"
                currentPageReportTemplate="Showing {first} to {last} of {totalRecords} kelas" class="mt-2">
                <Column selectionMode="multiple" style="width: 3rem;" :exportable="false"></Column>
                <Column field="nmKelas" header="Nama Kelas">
                    <template #loading>
                        <div class="flex items-center"
                            :style="{ height: '17px', 'flex-grow': '1', overflow: 'hidden' }">
                            <Skeleton width="40%" height="1rem" />
                        </div>
                    </template>
                </Column>
                <Column field="tingkatPendidikanId" header="Tingkat" sortable>
                    <template #loading>
                        <div class="flex items-center"
                            :style="{ height: '17px', 'flex-grow': '1', overflow: 'hidden' }">
                            <Skeleton width="40%" height="1rem" />
                        </div>
                    </template>
                </Column>
                <Column field="kurikulum.namaKurikulum" header="Kurikulum"></Column>
                <!-- Jika SMK/MAK Program Keahlian & Kompetensi Keahlian akan muncul-->
                <div v-if="['smk', 'mak'].includes(bentukPendidikan)">
                    <Column field="namaJurusanSp" header="Jurusan" sortable></Column>
                </div>
                <Column field="jumlahAnggota" header="Jml.Anggota"></Column>
                <Column header="Anggota Kelas">
                    <template #body="slotProps">
                        <Button icon="pi pi-bullseye" outlined rounded class="mr-2"
                            @click="dialogAnggotaRombel(slotProps.data)" />
                    </template>
                </Column>

            </DataTable>


        </div>
    </div>

    <!-- DIALOG  -->
    <Dialog v-model:visible="deleteKelasDialog" :style="{ width: '450px' }" header="Confirm" :modal="true">
        <div class="flex items-center gap-4">
            <i class="pi pi-exclamation-triangle !text-3xl" />
            <span v-if="product">Apakah kelas ini akan dihapus?</span>
        </div>
        <template #footer>
            <Button label="Tidak" icon="pi pi-times" text @click="deleteKelasDialog = false" />
            <Button label="Ya" icon="pi pi-check" text @click="deletedKelas" />
        </template>
    </Dialog>
    <!-- kelulusan -->
    <Dialog v-model:visible="isDialogKelulusan" :style="{ width: '450px' }" header="Confirm" :modal="true">
        <div class="flex items-center gap-4">
            <i class="pi pi-exclamation-triangle !text-3xl" />
            <span v-if="product">Apakah semua siswa di kelas ini akan diluluskan?</span>
        </div>
        <template #footer>
            <Button label="Tidak" icon="pi pi-times" text @click="isDialogKelulusan = false" />
            <Button label="Ya" icon="pi pi-check" text @click="luluskan" />
        </template>
    </Dialog>


    <DialogAnggotaKelas :visible=showAnggotaKelas />

    <DialogLoading :model-value=isLoading />
</template>

<script setup>
import { ref, onMounted, watch, computed, nextTick } from 'vue';
import { useStore } from "vuex";
const store = useStore();
// import DialogImport from '@/components/DialogImport.vue'
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Button from 'primevue/button';
import Dialog from 'primevue/dialog';
import Toolbar from 'primevue/toolbar';
import { FilterMatchMode } from '@primevue/core/api';
import { useToast } from 'primevue/usetoast';
import Select from 'primevue/select';
import DialogLoading from '@/components/DialogLoading.vue';
import router from '@/router';

import Skeleton from 'primevue/skeleton';

// ================================
// composable
// ================================
import { useSekolahService } from '@/composables/useSekolahService'
import DialogAnggotaKelas from '@/components/DialogAnggotaKelas.vue';
const selectedSemester = computed(() => store.getters["sekolahService/getSelectedSemester"])
const schemaname = computed(() => store.getters["sekolahService/getTabeltenant"]?.schemaname)

const { fetchKelas, fetchTingkat } = useSekolahService(schemaname, selectedSemester)
// ================================
const kelasList = ref()
const isLoading = ref(false)
const tingkatPendidikanOptions = ref()

watch(selectedSemester, async (e, b) => {
    // isLoading.value = true
    kelasList.value = await fetchKelas()
    // isLoading.value = false
})
onMounted(async () => {
    await fetchK()
    tingkatPendidikanOptions.value = await fetchTingkat()

});

const fetchK = async () => {
    try {
        const res = store.getters["sekolahService/getKelas"](selectedSemester.value?.semesterId)
        if (!res || res.length === 0) {
            // console.log("fetch to backend!")
            kelasList.value = await fetchKelas()
        } else {
            kelasList.value = res
        }
    } catch (error) {
        console.log(error)
    }
}

const toast = useToast();
const dt = ref();
const rombel = ref();
const deleteKelasDialog = ref(false);
const product = ref({});
const selectedKelas = ref();
const filters = ref({
    'global': { value: null, matchMode: FilterMatchMode.CONTAINS },
    'tingkatPendidikanId': { value: null, matchMode: FilterMatchMode.CONTAINS },
});
const submitted = ref(false);
const openNew = async () => {
    await nextTick();
    router.push({ name: "addKelas" })
};

const editKelas = async () => {
    await nextTick();

    router.push({
        name: "editKelas",
        query: { kelasId: selectedKelas.value[0]?.rombonganBelajarId.toString() }
    })
};

const exportCSV = () => {
    dt.value.exportCSV();
};
const confirmDeleteSelected = () => {
    deleteKelasDialog.value = true;
};
const deletedKelas = () => {
    rombel.value = rombel.value.filter(val => !selectedKelas.value.includes(val));
    deleteKelasDialog.value = false;
    selectedKelas.value = null;
    toast.add({ severity: 'success', summary: 'Successful', detail: 'Products Deleted', life: 3000 });
};

const showAnggotaKelas = ref(false)
const dialogAnggotaRombel = (d) => {

    showAnggotaKelas.value = true
    // console.log(d)
}
const bentukPendidikan = ref("smk")

const isLulus = ref(false)
const isNaik = ref(false)
const selectedKelasLulus = ref()
const selectedKelasNaik = ref()
watch(selectedKelas, (item) => {
    const adaKelas12 = selectedKelas.value.some(item => item.tingkatPendidikanId === 12);
    if (adaKelas12) {
        selectedKelasLulus.value = item
    } else {
        selectedKelasNaik.value = item
    }
});

watch(selectedKelasLulus, () => {
    if (!selectedKelasLulus.value || selectedKelasLulus.value.length === 0) {
        isLulus.value = false
    } else {
        isLulus.value = true
    }
})
watch(selectedKelasNaik, () => {
    if (!selectedKelasNaik.value || selectedKelasNaik.value.length === 0) {
        isNaik.value = false
    } else {
        isNaik.value = true
    }
})
const isDialogKelulusan = ref(false)
const luluskan = async () => {
    try {
        isDialogKelulusan.value = false
        const anggotaKelas = selectedKelas.value.flatMap(kelas => kelas?.anggotaKelas || []);
        const payload = {
            schemaname: await store.getters["sekolahService/getTabeltenant"]?.schemaname,
            tahun_ajaran_id: `${selectedSemester.value?.tahunAjaranId + 1}`,
            anggota_kelas: anggotaKelas,
            sekolah_id: await store.getters["sekolahService/getSekolah"]?.sekolah_id
        }
        const res = await store.dispatch("sekolahService/createProsesIjazah", payload)
        console.log(res)
        if (res) {
            toast.add({ severity: 'success', summary: 'Successful', detail: 'Data Ijazah berhasil ditambahkan', life: 3000 });
            selectedKelas.value = []
        }
    } catch (error) {

        toast.add({ severity: 'error', summary: 'Gagal', detail: 'Gagal menambahkan data', life: 3000 });
    }
    isDialogKelulusan.value = false
}
const dialogKenaikan = async () => {

}
</script>
