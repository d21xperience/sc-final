<template>
    <div class="container">
        <div class="card">
            <div v-if="dataConnected">
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
                                    </template>
                                    <template #end>
                                        <!-- <Button label="Import" icon="pi pi-download" severity="warn"
                                            @click="dialogImport = true" class="mr-2" /> -->
                                        <Button label="Export" icon="pi pi-upload" severity="help"
                                            @click="exportCSV($event)" class="mr-2" />
                                        <!-- <Button label="Proses" icon="pi pi-send" severity="info"
                                            @click="exportCSV($event)" v-tooltip.bottom="'Menyimpan ke database'"
                                            badge="2" /> -->
                                    </template>

                                </Toolbar>
                            </div>

                            <Toolbar>
                                <template #start>
                                    <div class="flex flex-wrap gap-2 items-center justify-between">
                                        <div class="flex">
                                            <!-- <MultiSelect v-model="selectedJurusan" :options="jurusan" optionLabel="name"
                                                filter placeholder="Jurusan" :maxSelectedLabels="1"
                                                class="w-full md:w-80 mr-2" showClear /> -->
                                            <!-- <Select v-model="selectedTingkat" showClear :options="tingkat"
                                                optionLabel="name" placeholder="Tingkat" class="mr-2" /> -->
                                        </div>
                                    </div>
                                </template>
                                <!-- <template #end>
                                    <IconField>
                                        <InputIcon>
                                            <i class="pi pi-search" />
                                        </InputIcon>
                                        <InputText v-model="filters['global'].value" placeholder="Search..." />
                                    </IconField>
                                </template> -->
                            </Toolbar>
                        </div>
                    </div>
                </div>


                <DataTable ref="dt" v-model:selection="selectedKelas" stripedRows size="small" :value="kelasList"
                    scrollable scrollHeight="400px" dataKey="rombonganBelajarId" :paginator="true" :rows="10"
                    :filters="filters" tableStyle="min-width: 50rem"
                    paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
                    :rowsPerPageOptions="[10, 20, 30]"
                    currentPageReportTemplate="Showing {first} to {last} of {totalRecords} kelas" class="mt-2">
                    <Column selectionMode="multiple" style="width: 3rem;" :exportable="false"></Column>
                    <!-- <Column field="name" header="Foto">
                        <template #body="slotProps">
                            <Image :src="`https://primefaces.org/cdn/primevue/images/product/${slotProps.data.image}`"
                                :alt="slotProps.data.image" preview image-class="w-16 h-16 rounded-full" />
                        </template>
                    </Column> -->
                    <Column field="nmKelas" header="Nama Kelas"></Column>
                    <Column field="tingkatPendidikanId" header="Tingkat" sortable></Column>
                    <Column field="kurikulum.namaKurikulum" header="Kurikulum"></Column>
                    <!-- Jika SMK/MAK Program Keahlian & Kompetensi Keahlian akan muncul-->
                    <div v-if="['smk', 'mak'].includes(bentukPendidikan)">
                        <Column field="namaJurusanSp" header="Jurusan" sortable></Column>
                    </div>
                    <Column field="ptk.nama" header="Wali kelas"></Column>
                    <Column field="jumlahAnggota" header="Jml.Anggota"></Column>
                    <Column field="code" header="Anggota Kelas">
                        <template #body="slotProps">
                            <!-- <Button icon="pi pi-bullseye" outlined rounded class="mr-2" @click="editProduct(slotProps.data)" /> -->
                            <Button icon="pi pi-bullseye" outlined rounded class="mr-2"
                                @click="dialogAnggotaRombel(slotProps.data)" />
                        </template>
                    </Column>
                    <!--<Column field="name" header="JK"></Column> -->
                    <!-- <Column field="name" header="Tpt.Lahir"></Column>
                    <Column field="name" header="Tgl.Lahir"></Column>
                    <Column field="name" header="Agama"></Column>
                    <Column field="category" header="Ayah"></Column>
                    <Column field="category" header="Ibu"></Column> -->
                    <!-- <Column field="category" header="Pekerjaan Ayah"></Column>
                    <Column field="category" header="Pekerjaan Ibu"></Column> -->
                    <!-- <Column field="category" header="Alamat"></Column> -->

                    <!-- <Column field="inventoryStatus" header="Status" sortable>
                        <template #body="slotProps">
                            <Tag :value="slotProps.data.inventoryStatus"
                                :severity="getStatusLabel(slotProps.data.inventoryStatus)" />
                        </template>
                    </Column> -->
                </DataTable>

            </div>
            <div v-else>
                <EmptyData @profileFetched="handleProfileFetched" @fetchError="handleFetchError" />
            </div>
        </div>
    </div>


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

    <DialogLoading :model-value=isLoading />
    <!-- import data -->
    <!-- DIALOG IMPORT -->
    <!-- <DialogImport v-model:visible="dialogImport" @save="saveImport" @cancel="cancelImport"
        template-type="kelas" :schema-name="schemaname" /> -->
</template>

<script setup>
import { ref, onMounted, watch, computed, nextTick } from 'vue';
import { useStore } from "vuex";
const store = useStore();
import DialogImport from '@/components/DialogImport.vue'
import FileUpload from 'primevue/fileupload';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Button from 'primevue/button';
import Dialog from 'primevue/dialog';
import Toolbar from 'primevue/toolbar';
import { FilterMatchMode } from '@primevue/core/api';
import { useToast } from 'primevue/usetoast';
import Select from 'primevue/select';
import MultiSelect from 'primevue/multiselect';
import DialogLoading from '@/components/DialogLoading.vue';
import EmptyData from '@/components/EmptyData.vue';
import router from '@/router';
// ================================
// composable
// ================================
import { useSekolahService } from '@/composables/useSekolahService'
const selectedSemester = computed(() => store.getters["sekolahService/getSelectedSemester"])
const schemaname = computed(() => store.getters["sekolahService/getTabeltenant"]?.schemaname)

const { fetchKelas, kelasList } = useSekolahService(schemaname, selectedSemester)
// ================================
defineProps({
    kelasData: String
});

const isLoading = ref(false)


watch(selectedSemester, (e, b) => {
    fetchK()
})
onMounted(() => {
    // console.log("onMounted")
    
    fetchK()
    // getParamDialogImport()
});

const fetchK = async () => {
    // Loading
    isLoading.value = true
    await fetchKelas()
    isLoading.value = false

}
// const getParamDialogImport = () => {
//     schemaname.value = store.getters["sekolahService/getTabeltenant"].schemaname
//     // semester.value = store.getters["sekolahService/getSemester"]

// }

const dataConnected = ref(true)
const toast = useToast();
const dt = ref();
const rombel = ref();
const kelas = ref({})
const kelasDialog = ref(false);
const deleteKelasDialog = ref(false);
const product = ref({});
const selectedKelas = ref();
const filters = ref({
    'global': { value: null, matchMode: FilterMatchMode.CONTAINS },
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


// ==================================

// Fungsi yang menangkap event emit dari child
const handleProfileFetched = (data) => {
    dataConnected.value = data;
    console.log("Data sekolah diterima di parent:", data);
};

const handleFetchError = (error) => {
    dataConnected.value = data;
    console.error("Error diterima di parent:", error);
};

// status siswa naik atau lulus
// const dialogStatus = ref(false)
const dialogImport = ref(false)

const dialogAnggotaRombel = (d) => {
    console.log(d)
}
const bentukPendidikan = ref("smk")
</script>
