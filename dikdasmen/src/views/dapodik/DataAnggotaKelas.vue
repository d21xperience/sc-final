<template>

    <div class="">
        <div class="card">
            <div v-if="dataConnected">
                <div class="fixed top-0 w-full left-0 z-20 bg-white">
                    <div class="lg:ml-[250px] my-2 ">
                        <div class="container ">
                            <div class="flex flex-wrap justify-between items-center mb-2">
                                <h4 class="font-bold text-xl md:text-2xl">Data Siswa </h4>
                                <div class="md:flex md:items-center md:space-x-2">
                                    <h3 class="text-slate-500 md:text-base text-sm">Tahun Pelajaran</h3>
                                    <div>
                                        <Select v-model="selectedSemester" :options="semester"
                                            optionLabel="namaSemester" placeholder="Tahun Pelajaran"
                                            class="w-full md:w-52 mr-2" />

                                    </div>
                                </div>
                            </div>
                            <div class="mb-2">
                                <Toolbar>
                                    <template #start>
                                        <Button icon="pi pi-pencil" severity="warn" @click="confirmDeleteSelected"
                                            :disabled="!dataLulusan || !dataLulusan.length || dataLulusan.length > 2"
                                            class="mr-2" />
                                        <Button icon="pi pi-trash" severity="danger" class="mr-2"
                                            @click="confirmDeleteSelected"
                                            :disabled="!dataLulusan || !dataLulusan.length" />
                                        <Button label="Lulus" severity="warn" class="mr-2" @click="dialogStatus = true"
                                            :disabled="!dataLulusan || !dataLulusan.length" />
                                        <Button label="Naik" severity="warn" class="mr-2" @click="openNew"
                                            :disabled="!dataLulusan || !dataLulusan.length" />
                                    </template>
                                    <template #end>
                                        <FileUpload mode="basic" accept="xlsx/*" :maxFileSize="1000000" label="Import"
                                            chooseLabel="Import" class="mr-2" auto />
                                        <Button label="Export" icon="pi pi-upload" severity="help"
                                            @click="exportCSV($event)" class="mr-2" />
                                        <!-- <Button label="Proses" icon="pi pi-send" severity="info"
                                            @click="exportCSV($event)" /> -->
                                    </template>

                                </Toolbar>
                            </div>

                            <Toolbar>
                                <template #start>
                                    <div class="flex flex-wrap gap-2 items-center justify-between">
                                        <div class="flex">
                                            <Select v-model="selectedJurusan" :options="jurusan" optionLabel="name"
                                                placeholder="Rombel" class="w-full md:w-56 mr-2" />
                                            <Select v-model="selectedJurusan" :options="jurusan" optionLabel="name"
                                                placeholder="Tingkat" class="mr-2" />
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
                    </div>
                </div>


                <DataTable ref="dt" v-model:selection="selectedSiswa" stripedRows size="small" :value="siswa"
                    dataKey="id" :paginator="true" :rows="10" :filters="filters"
                    paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
                    :rowsPerPageOptions="[10, 20, 50]"
                    currentPageReportTemplate="Showing {first} to {last} of {totalRecords} products" class="mt-56">
                    <Column selectionMode="multiple" style="width: 3rem;" :exportable="false"></Column>
                    <Column field="name" header="Foto">
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
                    <Column field="jk" header="JK">
                        <template #body="slotProps">
                            {{ slotProps.data.pesertaDidik.jenisKelamin }}
                        </template>
                    </Column>
                    <Column field="nisn" header="NISN">
                        <template #body="slotProps">
                            {{ slotProps.data.pesertaDidik.nisn }}
                        </template>
                    </Column>


                    <Column field="nis" header="NIS" sortable>
                        <template #body="slotProps">
                            {{ slotProps.data.pesertaDidik.nis }}
                        </template>
                    </Column>
                    <Column field="tingkat" header="Tingkat" sortable>
                        <!-- rombonganBelajar -->
                        <template #body="slotProps">
                            {{ slotProps.data.rombonganBelajar.tingkatPendidikanId }}
                        </template>
                    </Column>
                    <Column field="rombel" header="Rombel" sortable>
                        <template #body="slotProps">
                            {{ slotProps.data.rombonganBelajar.nmKelas }}
                        </template>
                    </Column>
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
    </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue';
import { useStore } from "vuex";
const store = useStore();
import FileUpload from 'primevue/fileupload';

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
import DataLulusanService from '@/service/ProductService.js';
// =============UJI FITUR FOTO========================
import Image from 'primevue/image';
// =====================================


onMounted(() => {
    fetchSemester()
    fetchSiswa()
    // DataLulusanService.getProducts().then((data) => (products.value = data));

});
// watch(selectedSemester, (newVal, oldVal) => {
//     console.log(newVal)
//     // fetchRombel()
// })


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
});
const submitted = ref(false);
const statuses = ref([
    { label: 'INSTOCK', value: 'instock' },
    { label: 'LOWSTOCK', value: 'lowstock' },
    { label: 'OUTOFSTOCK', value: 'outofstock' }
]);

const formatCurrency = (value) => {
    if (value)
        return value.toLocaleString('en-US', { style: 'currency', currency: 'USD' });
    return;
};
const openNew = () => {
    product.value = {};
    submitted.value = false;
    productDialog.value = true;
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
const editProduct = (prod) => {
    product.value = { ...prod };
    productDialog.value = true;
};
const confirmDeleteProduct = (prod) => {
    product.value = prod;
    deleteProductDialog.value = true;
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
    dt.value.exportCSV();
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

const getStatusLabel = (status) => {
    switch (status) {
        case 'INSTOCK':
            return 'success';

        case 'LOWSTOCK':
            return 'warn';

        case 'OUTOFSTOCK':
            return 'danger';

        default:
            return null;
    }
};



import Select from 'primevue/select';
import EmptyData from '@/components/EmptyData.vue';

// select tahun ijazah
const selectedCity = ref();
const cities = ref([
    { name: '2023/2024 Ganjil', code: '20231' },
    { name: '2023/2024 Genap', code: '20232' },
    { name: '2022/2023', code: '20222' },
    { name: '2021/2022', code: '20212' },
    { name: '2022/2021', code: '20202' },
    { name: '2019/2020', code: '20192' }
]);
const selectedJurusan = ref();
const jurusan = ref([
    { name: 'Teknik Kendaraan Ringan', code: 'TKR' },
    { name: 'Teknik Mesin Sepeda Motor', code: 'TSM' },
    { name: 'Teknik Komputer dan Jaringan', code: 'TKJ' },
    { name: 'Otomatisasi Perkantoran', code: 'OTKP' },
    { name: 'AKuntansi Lembaga', code: 'AKL' }
]);

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
const dialogStatus = ref(false)


// ==================================
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
// =======Siswa=============
const selectedSiswa = ref();
const siswa = ref(null);
const fetchSiswa = async () => {
    try {
        console.log("fethcSiswa")
        let payload = {
            semesterId: 20232,
            schemaName: "tabel_D4DA6B98FCFD71C58F5A"
        }
        console.log(payload)
        const results = await store.dispatch("sekolahService/fetchSiswa", payload)
        // console.log(results)
        if (results) {
            siswa.value = store.getters["sekolahService/getSiswa"]
            // // Ambil siswa terbaru berdasarkan ID terbesar
            // selectedSiswa.value = siswa.value.reduce((latest, current) =>
            //     current.siswaId > latest.siswaId ? current : latest
            // );
        }
    } catch (error) {

    }
}

// ==================================




</script>
