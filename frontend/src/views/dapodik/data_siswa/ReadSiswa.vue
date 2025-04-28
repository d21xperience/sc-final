<template>

    <div class="">
        <div class="card">
            <div class="w-full my-2 container">
                <div class=" ">
                    <div class="mb-2">
                        <h2 class="text-xl mb-2">Data Siswa</h2>
                        <Toolbar>
                            <template #start>
                                <!-- <Button icon="pi pi-plus" severity="success" class="mr-2 text-lg" @click="openNew"
                                        v-tooltip.bottom="'Tambah siswa'" /> -->
                                <Button icon="pi pi-pencil" severity="warn" @click="confirmDeleteSelected"
                                    :disabled="!dataLulusan || !dataLulusan.length || dataLulusan.length > 2"
                                    class="mr-2" v-tooltip.bottom="'Edit siswa'" />
                                <Button icon="pi pi-trash" severity="danger" class="mr-2 text-lg"
                                    @click="confirmDeleteSelected" :disabled="!dataLulusan || !dataLulusan.length"
                                    v-tooltip.bottom="'Hapus siswa'" />
                                <!-- <div v-show="selectedSemester.semester === 2">

                                    <Button label="Lulus" severity="help" class="mr-2 text-sm"
                                        @click="dialogStatus = true" :disabled="!isLulus"
                                        v-tooltip.bottom="'Luluskan siswa'" />
                                    <Button label="Naik" severity="success" class="mr-2 text-sm" @click="openNew"
                                        :disabled="!isNaik" v-tooltip.bottom="'Naikan siswa'" />
                                </div> -->
                            </template>
                            <template #end>
                                <Button label="Import" icon="pi pi-download" severity="warn"
                                    @click="dialogImport = true" class="mr-2 text-sm"
                                    v-tooltip.bottom="'Import siswa'" />
                                <Button label="Export" icon="pi pi-upload" severity="help" @click="exportCSV($event)"
                                    class="mr-2 text-sm" />
                                <!-- <Button label="Proses" icon="pi pi-send" severity="info"
                                            @click="exportCSV($event)" /> -->
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
                                    <!-- <Select v-model="selectedJurusan" :options="jurusanOptions" optionLabel="label"
                                        placeholder="Rombel" class="mr-2" /> -->
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


                <DataTable ref="dt" v-model:selection="selectedSiswa" stripedRows :value="siswa" scrollable
                    scrollHeight="450px" dataKey="pesertaDidikId" :paginator="true" :rows="10" :filters="filters"
                    paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
                    :rowsPerPageOptions="[10, 20, 50]"
                    currentPageReportTemplate="Showing {first} to {last} of {totalRecords} siswa">
                    <template #empty>
                        <p class="text-xl flex justify-center font-bold text-red-500">Nilai tidak ditemukan.</p>
                    </template>

                    <Column selectionMode="multiple" style="width: 3rem;" :exportable="false"></Column>
                    <Column field="nmSiswa" header="Nama" sortable></Column>
                    <Column field="jenisKelamin" header="JK"></Column>
                    <Column field="nisn" header="NISN"></Column>
                    <Column field="nis" header="NIS" sortable></Column>
                    <Column field="agama" header="Agama"></Column>
                    <Column field="tempatLahir" header="Tpt Lahir"></Column>
                    <Column field="tanggalLahir" header="Tgl Lahir">
                        <template #body="slotProps">
                            {{ formatterDateID(slotProps.data.tanggalLahir) }}
                        </template>
                    </Column>
                    <Column field="tingkatPendidikanId" header="Tingkat" sortable></Column>
                    <Column field="nmKelas" header="Rombel" sortable></Column>

                </DataTable>

            </div>

        </div>


        <!-- DIALOGBOX FOR EDIT DATA -->
        <Dialog v-model:visible="productDialog" :style="{ height: '650px', width: '450px' }" header="Edit Data"
            :modal="true">
            <div class="flex flex-wrap gap-6">
                <div>
                    <label for="name" class="block font-bold">NISN</label>
                    <InputText id="name" v-model.trim="product.code" required="true" autofocus
                        :invalid="submitted && !product.code" fluid />
                    <small v-if="submitted && !product.code" class="text-red-500">NISN is required.</small>
                </div>
                <div>
                    <label for="name" class="block font-bold ">Nama</label>
                    <InputText id="name" v-model.trim="product.name" required="true" autofocus
                        :invalid="submitted && !product.name" fluid />
                    <small v-if="submitted && !product.name" class="text-red-500">Nama is required.</small>
                </div>
                <div>
                    <label for="name" class="block font-bold ">Rerata Nilai</label>
                    <InputText id="name" v-model.trim="product.price" required="true" autofocus
                        :invalid="submitted && !product.price" fluid />
                    <small v-if="submitted && !product.price" class="text-red-500">Nilai is required.</small>
                </div>
                <div>
                    <label for="name" class="block font-bold ">Thn Lulus</label>
                    <InputText id="name" v-model.trim="product.category" required="true" autofocus
                        :invalid="submitted && !product.category" fluid />
                    <small v-if="submitted && !product.category" class="text-red-500">Thn lulus is required.</small>
                </div>
            </div>

            <template #footer>
                <Button label="Cancel" icon="pi pi-times" text @click="hideDialog" />
                <Button label="Save" icon="pi pi-check" @click="saveProduct" />
            </template>
        </Dialog>

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

        <Dialog v-model:visible="deleteProductsDialog" :style="{ width: '450px' }" header="Confirm" :modal="true">
            <div class="flex items-center gap-4">
                <i class="pi pi-exclamation-triangle !text-3xl" />
                <span v-if="product">Apakah data lulusan akan dihapus?</span>
            </div>
            <template #footer>
                <Button label="Tidak" icon="pi pi-times" text @click="deleteProductsDialog = false" />
                <Button label="Ya" icon="pi pi-check" text @click="deletedataLulusan" />
            </template>
        </Dialog>

        <!-- Dialog Status kenaikan/ lulus -->
        <Dialog v-model:visible="dialogStatus" :style="{ width: '450px' }" header="Confirm" :modal="true">
            <div class="flex items-center gap-4">
                <i class="pi pi-exclamation-triangle !text-3xl" />
                <span v-if="product">Apakah siswa akan diluluskan?</span>
            </div>
            <template #footer>
                <Button label="Tidak" icon="pi pi-times" text @click="dialogStatus = false" />
                <Button label="Ya" icon="pi pi-check" text @click="deletedataLulusan" />
            </template>
        </Dialog>


        <!-- import data -->
        <!-- DIALOG IMPORT -->
        <DialogImport v-model:visible="dialogImport" @save="saveImport" @cancel="cancelImport" template-type="siswa"
            :schema-name="schemaname" />

        <!-- end of import data -->
        <DialogLoading v-model="isLoading"> Memuat data, harap tunggu... </DialogLoading>

    </div>
</template>

<script setup>
import { ref, onMounted, watch, computed } from 'vue';
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


// ==============================
onMounted(async () => {
    //semester.value = store.getters["sekolahService/getSemester"]
    // schemaname.value = store.getters["sekolahService/getTabeltenant"].schemaname
    // siswa.value = store.getters["sekolahService/getSiswaAktif"]
    // console.log(siswa.value)
    // if (!siswa.value) {
    siswa.value = await fetchSiswaAktif()
    // }
    tingkatPendidikanOptions.value = await fetchTingkat()
    // console.log(jurusanOptions.value)
});
import DialogLoading from "@/components/DialogLoading.vue";

const isLoading = ref(false);

const dataConnected = ref(true)
const toast = useToast();
const dt = ref();
const products = ref();
const productDialog = ref(false);
const deleteProductDialog = ref(false);
const deleteProductsDialog = ref(false);
const product = ref({});
const dataLulusan = ref();
const filters = ref({
    'global': { value: null, matchMode: FilterMatchMode.CONTAINS },
    'tingkatPendidikanId': { value: null, matchMode: FilterMatchMode.EQUALS }
});
const submitted = ref(false);

const openNew = () => {
    router.push({ name: "inputSiswa" })
};
const hideDialog = () => {
    productDialog.value = false;
    submitted.value = false;
};
const saveProduct = () => {
    submitted.value = true;

    if (product?.value.name?.trim()) {
        if (product.value.id) {
            product.value.inventoryStatus = product.value.inventoryStatus.value ? product.value.inventoryStatus.value : product.value.inventoryStatus;
            products.value[findIndexById(product.value.id)] = product.value;
            toast.add({ severity: 'success', summary: 'Successful', detail: 'Product Updated', life: 3000 });
        }
        else {
            product.value.id = createId();
            product.value.code = createId();
            product.value.image = 'product-placeholder.svg';
            product.value.inventoryStatus = product.value.inventoryStatus ? product.value.inventoryStatus.value : 'INSTOCK';
            products.value.push(product.value);
            toast.add({ severity: 'success', summary: 'Successful', detail: 'Product Created', life: 3000 });
        }

        productDialog.value = false;
        product.value = {};
    }
};
const deleteProduct = () => {
    products.value = products.value.filter(val => val.id !== product.value.id);
    deleteProductDialog.value = false;
    product.value = {};
    toast.add({ severity: 'success', summary: 'Successful', detail: 'Product Deleted', life: 3000 });
};
const findIndexById = (id) => {
    let index = -1;
    for (let i = 0; i < products.value.length; i++) {
        if (products.value[i].id === id) {
            index = i;
            break;
        }
    }

    return index;
};
const createId = () => {
    let id = '';
    var chars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
    for (var i = 0; i < 5; i++) {
        id += chars.charAt(Math.floor(Math.random() * chars.length));
    }
    return id;
}
const exportCSV = () => {
    isLoading.value = true
    // alert("hello")
    // dt.value.exportCSV();
};
const confirmDeleteSelected = () => {
    deleteProductsDialog.value = true;
};
const deletedataLulusan = () => {
    products.value = products.value.filter(val => !dataLulusan.value.includes(val));
    deleteProductsDialog.value = false;
    dataLulusan.value = null;
    toast.add({ severity: 'success', summary: 'Successful', detail: 'Products Deleted', life: 3000 });
};


import Select from 'primevue/select';

// status siswa naik atau lulus
const dialogStatus = ref(false)


// ==================================
// =======SEMESTER=============
const selectedSiswa = ref();
const siswa = ref([]);
// ================================
// composable
// ================================
import { useSekolahService } from '@/composables/useSekolahService'
const selectedSemester = computed(() => store.getters["sekolahService/getSelectedSemester"])
const schemaname = computed(() => store.getters["sekolahService/getTabeltenant"]?.schemaname)
const { fetchSiswaAktif, fetchTingkat } = useSekolahService(schemaname, selectedSemester)
// ================================
watch(selectedSemester, async (e, b) => {
    siswa.value = await fetchSiswaAktif()
})
// ========IMPORT DATA========
import DialogImport from '@/components/DialogImport.vue'
import router from '@/router';
import { result } from 'lodash';
const dialogImport = ref(false)
const saveImport = (e) => {
    // console.log("Data disimpan:", e);
    dialogImport.value = false;
};

const cancelImport = () => {
    console.log("Import dibatalkan");
    dialogImport.value = false;
};
// ===========================================
// import { useRombelOptions } from '@/composables/useOptions'
// const { jurusanOptions, tingkatPendidikanOptions } = useRombelOptions(siswa)
// const tes = ref()
// const selectedJurusan = ref();
// const jurusan = ref([
//     { name: 'Teknik Kendaraan Ringan', code: 'TKR' },
//     { name: 'Teknik Mesin Sepeda Motor', code: 'TSM' },
//     { name: 'Teknik Komputer dan Jaringan', code: 'TKJ' },
//     { name: 'Otomatisasi Perkantoran', code: 'OTKP' },
//     { name: 'AKuntansi Lembaga', code: 'AKL' }
// ]);

import { useUtils } from '@/composables/useUtils'
const { formatterDateID } = useUtils()

const tingkatPendidikanOptions = ref([]) // return Array.from(unique.values())
const selectedTingkat = ref();
const selectedJurusan = ref();
// const jurusanOptions = computed(() => {
//     const unique = new Set();
//     return siswa.value
//         .filter(item => {
//             if (!unique.has(item.rombonganBelajar.namaJurusanSp)) {
//                 unique.add(item.rombonganBelajar.namaJurusanSp);
//                 return true;
//             }
//             return false;
//         })
//         .map(item => ({
//             label: item.rombonganBelajar.namaJurusanSp,
//             value: item.rombonganBelajar.namaJurusanSp
//         }));
// });

// lulus
const isLulus = ref(false)
const isNaik = ref(false)
</script>
