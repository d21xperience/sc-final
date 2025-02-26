<template>

    <!-- <div class=""> -->
    <div class="">
        <div class="my-2">
            <Toolbar>
                <template #start>
                    <Button icon="pi pi-pencil" severity="warn" @click="confirmDeleteSelected"
                        :disabled="!dataLulusan || !dataLulusan.length || dataLulusan.length > 2" class="mr-2" />
                    <Button icon="pi pi-trash" severity="danger" class="mr-2" @click="confirmDeleteSelected"
                        :disabled="!dataLulusan || !dataLulusan.length" />
                    <Button icon="pi pi-key" severity="success" class="mr-2" @click="confirmDeleteSelected"
                        :disabled="!dataLulusan || !dataLulusan.length" />
                    <Button icon="pi pi-download" severity="warn" class="mr-2" @click="confirmDeleteSelected"
                         />
                </template>
                <template #end>

                    <div class="flex flex-wrap gap-2 items-center justify-between">
                        <div class="flex">
                            <!-- <Select v-model="selectedJurusan" :options="jurusan" optionLabel="name"
                                placeholder="Tahun Ajaran" class="md:w-48  mr-2" /> -->
                            <Select v-model="selectedJurusan" :options="jurusan" optionLabel="name" placeholder="Rombel"
                                class="w-full md:w-48 mr-1" />
                            <!-- <Button icon="pi pi-times" severity="danger" class="mr-2" @click="confirmDeleteSelected"
                                :disabled="!dataLulusan || !dataLulusan.length" /> -->
                        </div>
                    </div>
                    <IconField>
                        <InputIcon>
                            <i class="pi pi-search" />
                        </InputIcon>
                        <InputText v-model="filters['global'].value" placeholder="Search..." />
                    </IconField>
                </template>
            </Toolbar>
        </div>
        <DataTable ref="dt" v-model:selection="dataLulusan" stripedRows size="small" :value="products" dataKey="id"
            :paginator="true" :rows="5" :filters="filters"
            paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
            :rowsPerPageOptions="[5, 10, 25]"
            currentPageReportTemplate="Showing {first} to {last} of {totalRecords} Users" class="">
            <Column selectionMode="multiple" style="width: 3rem;" :exportable="false"></Column>
            <Column field="name" header="Username" sortable></Column>
            <Column field="name" header="Nama user" sortable></Column>
            <Column field="code" header="Login terakhir"></Column>
            <Column field="code" header="Online" sortable></Column>
            <Column field="code" header="Status" sortable></Column>
            <!-- <Column field="inventoryStatus" header="Status" sortable>
                <template #body="slotProps">
                    <Tag :value="slotProps.data.inventoryStatus"
                        :severity="getStatusLabel(slotProps.data.inventoryStatus)" />
                </template>
            </Column> -->
        </DataTable>




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
import FileUpload from 'primevue/fileupload';

import DataTable from 'primevue/datatable';
import Column from 'primevue/column';

import Button from 'primevue/button';

import Dialog from 'primevue/dialog';

import Toolbar from 'primevue/toolbar';

import ColumnGroup from 'primevue/columngroup';   // optional
import Row from 'primevue/row';                   // optional

import { ref, onMounted } from 'vue';
import { FilterMatchMode } from '@primevue/core/api';
import { useToast } from 'primevue/usetoast';
import InputText from 'primevue/inputtext';
import IconField from 'primevue/iconfield';
import InputIcon from 'primevue/inputicon';
import RadioButton from 'primevue/radiobutton';
import DataLulusanService from '@/service/ProductService.js';



onMounted(() => {
    DataLulusanService.getProducts().then((data) => (products.value = data));
});

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
</script>
