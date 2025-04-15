<script setup>
import { ref, onMounted, watch, computed, nextTick } from 'vue';
import { useSekolahService } from '@/composables/useSekolahService'
import { useStore } from "vuex";
const store = useStore();

// import FileUpload from 'primevue/fileupload';

import DataTable from 'primevue/datatable';
import Column from 'primevue/column';

import Button from 'primevue/button';

import Dialog from 'primevue/dialog';

import Toolbar from 'primevue/toolbar';
import { FilterMatchMode } from '@primevue/core/api';
import { useToast } from 'primevue/usetoast';
import InputText from 'primevue/inputtext';
import IconField from 'primevue/iconfield';
import InputIcon from 'primevue/inputicon';
import DialogLoading from "@/components/DialogLoading.vue";

const isLoading = ref(false);

const dataConnected = ref(true)
const toast = useToast();
const dt = ref();
// const products = ref();
// const productDialog = ref(false);
const deleteProductDialog = ref(false);
const deleteProductsDialog = ref(false);
const product = ref({});
const filters = ref({
    'global': { value: null, matchMode: FilterMatchMode.CONTAINS },
});
const submitted = ref(false);

// const formatCurrency = (value) => {
//     if (value)
//         return value.toLocaleString('en-US', { style: 'currency', currency: 'USD' });
//     return;
// };
const openNew = async () => {
    await nextTick()
    router.push({ name: "inputGuru" })
};

const deleteProduct = () => {
    products.value = products.value.filter(val => val.id !== product.value.id);
    deleteProductDialog.value = false;
    product.value = {};
    toast.add({ severity: 'success', summary: 'Successful', detail: 'Product Deleted', life: 3000 });
};

const exportCSV = () => {
    isLoading.value = true
    // alert("hello")
    // dt.value.exportCSV();
};
const confirmDeleteSelected = () => {
    deleteProductsDialog.value = true;
};

import Select from 'primevue/select';
import EmptyData from '@/components/EmptyData.vue';
onMounted(async () => {
    await fetchGuruTerdaftar()
});
// Fungsi yang menangkap event emit dari child
const handleProfileFetched = (data) => {
    dataConnected.value = data;
    console.log("Data sekolah diterima di parent:", data);
};

const handleFetchError = (error) => {
    dataConnected.value = data;
    console.error("Error diterima di parent:", error);
};

// ==================================
// =======composable=============
const selectedSemester = computed(() => store.getters["sekolahService/getSelectedSemester"])
const schemaname = computed(() => store.getters["sekolahService/getTabeltenant"].schemaname)
const selectedGuru = ref();
const { guruTerdaftarList, fetchGuruTerdaftar } = useSekolahService(schemaname, selectedSemester)
// ==================================

watch(selectedSemester, async (e, b) => {
    await fetchGuruTerdaftar()
})

// ========IMPORT DATA========
import DialogImport from '@/components/DialogImport1.vue'
import router from '@/router';
const dialogImport = ref(false)
const saveImport = (e) => {
    // console.log("Data disimpan:", e);
    dialogImport.value = false;
};

const cancelImport = () => {
    console.log("Import dibatalkan");
    dialogImport.value = false;
};





const editGuru = async () => {
    await nextTick();

    router.push({
        name: "editGuru",
        query: { ptkId: selectedGuru.value[0]?.ptk.ptkId.toString() }
    })
};
</script>

<template>

    <div class="">
        <div class="card">
            <div v-if="dataConnected">
                <div class="w-full my-2 container">
                    <div class=" ">
                        <div class="mb-2">
                            <h2 class="text-xl mb-2">Data Guru</h2>
                            <Toolbar>
                                <template #start>
                                    <Button icon="pi pi-plus" severity="success" class="mr-2 text-lg" @click="openNew"
                                        v-tooltip.bottom="'Tambah Guru Baru'" />
                                    <Button icon="pi pi-pencil" severity="warn" @click="editGuru"
                                        :disabled="!selectedGuru || !selectedGuru.length || selectedGuru.length > 1"
                                        class="mr-2" v-tooltip.bottom="'Edit Guru'" />
                                    <Button icon="pi pi-trash" severity="danger" class="mr-2 text-lg"
                                        @click="confirmDeleteSelected" :disabled="!selectedGuru || !selectedGuru.length"
                                        v-tooltip.bottom="'Hapus Guru'" />
                                </template>
                                <template #end>
                                    <Button label="Import" icon="pi pi-download" severity="warn"
                                        @click="dialogImport = true" class="mr-2 text-sm"
                                        v-tooltip.bottom="'Import Guru'" />
                                    <Button label="Export" icon="pi pi-upload" severity="help"
                                        @click="exportCSV($event)" class="mr-2 text-sm" />
                                    <!-- <Button label="Proses" icon="pi pi-send" severity="info"
                                            @click="exportCSV($event)" /> -->
                                </template>

                            </Toolbar>
                        </div>

                        <Toolbar>
                            <!-- <template #start>
                                <div class="flex flex-wrap gap-2 items-center justify-between">
                                    <div class="flex">
                                        <Select v-model="selectedJurusan" :options="jurusan" optionLabel="name"
                                            placeholder="Rombel" class="w-full md:w-56 mr-2" />
                                        <Select v-model="selectedJurusan" :options="jurusan" optionLabel="name"
                                            placeholder="Tingkat" class="mr-2" />
                                    </div>
                                </div>
                            </template> -->
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


                    <DataTable ref="dt" v-model:selection="selectedGuru" stripedRows size="small" :value="guruTerdaftarList"
                        dataKey="ptkTerdaftarId" :paginator="true" :rows="10" :filters="filters"
                        paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
                        :rowsPerPageOptions="[10, 20, 50]"
                        currentPageReportTemplate="Showing {first} to {last} of {totalRecords} Guru">
                        <Column selectionMode="multiple" style="width: 3rem;" :exportable="false"></Column>
                        <!-- <Column field="name" header="Foto">
                            <template #body="slotProps">
                                <Image
                                    :src="`https://primefaces.org/cdn/primevue/images/product/${slotProps.data.image}`"
                                    :alt="slotProps.data.image" preview image-class="w-16 h-16 rounded-full" />
                            </template>
                        </Column> -->
                        <Column field="ptk.nama" header="Nama" sortable>
                        </Column>
                        <Column field="ptk.gelarBelakang" header="Gelar belakang">
                        </Column>
                        <Column field="ptk.jenisKelamin" header="JK">
                        </Column>
                        <Column field="ptk.nip" header="NIP">
                        </Column>
                        <Column field="ptk.nuptk" header="NUPTK">
                        </Column>
                        <Column field="ptk.tempatLahir" header="Tpt.Lahir">
                        </Column>

                        <Column field="ptk.tanggalLahir" header="Tgl.Lahir">
                        </Column>
                    </DataTable>

                </div>
            </div>
            <div v-else>
                <EmptyData @profileFetched="handleProfileFetched" @fetchError="handleFetchError" />
            </div>
        </div>

        <Dialog v-model:visible="deleteProductDialog" :style="{ width: '450px' }" header="Confirm" :modal="true">
            <div class="flex items-center gap-4">
                <i class="pi pi-exclamation-triangle !text-3xl" />
                <span v-if="product">Are you sure you want to delete <b>{{ product.name }}</b>?</span>
            </div>
            <template #footer>
                <Button label="No" icon="pi pi-times" text @click="deleteProductDialog = false" />
                <Button label="Yes" icon="pi pi-check" @click="deleteProduct" />
            </template>
        </Dialog>

        <!-- import data -->
        <!-- DIALOG IMPORT -->
        <DialogImport v-model:visible="dialogImport" @save="saveImport" @cancel="cancelImport" template-type="guru"
            :schema-name="schemaname" />

        <!-- end of import data -->
        <DialogLoading v-model="isLoading"> Memuat data, harap tunggu... </DialogLoading>

    </div>
</template>
