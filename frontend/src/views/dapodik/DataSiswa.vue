<template>

    <div class="container mx-auto p-8">

        <div class="card">
            <div class="flex flex-wrap justify-between my-2">
                <h4 class="font-bold text-2xl lg:text-lg my-2">Data Siswa </h4>
            </div>
            <div v-if="dataConnected">
                <Toolbar class="mb-6">
                    <template #end>
                        <!-- <Select v-model="selectedCity" :options="cities" optionLabel="name" placeholder="Tahun Pelajaran"
                            class="w-full md:w-56 mr-2" /> -->
                        <Select v-model="selectedCity" :options="cities" optionLabel="name"
                            placeholder="Tahun Pelajaran" class="w-full md:w-56 mr-2" />

                    </template>
                    <template #start>
                        <!-- <Button label="New" icon="pi pi-plus" severity="success" class="mr-2" @click="openNew" /> -->
                        <Button label="Edit" icon="pi pi-pencil" severity="warn" @click="confirmDeleteSelected"
                            :disabled="!dataLulusan || !dataLulusan.length || dataLulusan.length > 2" class="mr-2" />
                        <Button label="Delete" icon="pi pi-trash" severity="danger" class="mr-2"
                            @click="confirmDeleteSelected" :disabled="!dataLulusan || !dataLulusan.length" />

                        <!-- <FileUpload mode="basic" accept="image/*" :maxFileSize="1000000" label="Import" chooseLabel="Import"
                            class="mr-2" auto /> -->
                        <Button label="Export" icon="pi pi-upload" severity="help" @click="exportCSV($event)"
                            class="mr-2" />
                        <Button label="Kirim Blockchain" icon="pi pi-send" severity="help" @click="exportCSV($event)" />
                    </template>
                </Toolbar>

                <DataTable ref="dt" v-model:selection="dataLulusan" :value="products" dataKey="id" :paginator="true"
                    :rows="10" :filters="filters"
                    paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
                    :rowsPerPageOptions="[5, 10, 25]"
                    currentPageReportTemplate="Showing {first} to {last} of {totalRecords} products">
                    <template #header>
                        <div class="flex flex-wrap gap-2 items-center justify-between">
                            <div class="flex">
                                <Select v-model="selectedJurusan" :options="jurusan" optionLabel="name"
                                    placeholder="Rombel" class="w-full md:w-56 mr-2" />
                                <Select v-model="selectedJurusan" :options="jurusan" optionLabel="name"
                                    placeholder="Kompetensi Keahlian" class="w-full md:w-56 mr-2" />
                                <IconField>
                                    <InputIcon>
                                        <i class="pi pi-search" />
                                    </InputIcon>
                                    <InputText v-model="filters['global'].value" placeholder="Search..." />
                                </IconField>
                            </div>

                        </div>
                    </template>

                    <Column selectionMode="multiple" style="width: 3rem" :exportable="false"></Column>
                    <Column field="code" header="NISN"></Column>
                    <Column field="code" header="NIS" sortable></Column>
                    <Column field="code" header="Tingkat" sortable></Column>
                    <Column field="code" header="Rombel" sortable></Column>
                    <Column field="name" header="Nama" sortable></Column>
                    <Column field="name" header="JK"></Column>
                    <Column field="name" header="Tpt.Lahir"></Column>
                    <Column field="name" header="Tgl.Lahir"></Column>
                    <Column field="name" header="Agama"></Column>
                    <!-- <Column field="price" header="Nilai Rerata">
                        <template #body="slotProps">
                            {{ formatCurrency(slotProps.data.price) }}
                        </template>
                    </Column> -->
                    <Column field="category" header="Ayah"></Column>
                    <Column field="category" header="Ibu"></Column>
                    <Column field="category" header="Pekerjaan Ayah"></Column>
                    <Column field="category" header="Pekerjaan Ibu"></Column>
                    <Column field="category" header="Alamat"></Column>
                    <!-- <Column header="Image">
                        <template #body="slotProps">
                            <img :src="`https://primefaces.org/cdn/primevue/images/product/${slotProps.data.image}`"
                                :alt="slotProps.data.image" class="rounded" style="width: 64px" />
                        </template>
                    </Column> -->
                    <!--<Column field="rating" header="Reviews" sortable style="min-width: 12rem">
                        <template #body="slotProps">
                            <Rating :modelValue="slotProps.data.rating" :readonly="true" />
                        </template>
                    </Column>-->
                    <Column field="inventoryStatus" header="Status" sortable>
                        <template #body="slotProps">
                            <Tag :value="slotProps.data.inventoryStatus"
                                :severity="getStatusLabel(slotProps.data.inventoryStatus)" />
                        </template>
                    </Column>
                    <!-- <Column header="Aksi" :exportable="false" style="min-width: 12rem">
                        <template #body="slotProps">
                            <Button icon="pi pi-pencil" outlined rounded class="mr-2"
                                @click="editProduct(slotProps.data)" />
                            <Button icon="pi pi-trash" outlined rounded severity="danger"
                                @click="confirmDeleteProduct(slotProps.data)" />
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
                <!-- <img v-if="product.image" :src="`https://primefaces.org/cdn/primevue/images/product/${product.image}`"
                    :alt="product.image" class="block m-auto pb-4" /> -->
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
                <!-- <div>
                    <label for="description" class="block font-bold mb-3">Description</label>
                    <Textarea id="description" v-model="product.description" required="true" rows="3" cols="20" fluid />
                </div> -->
                <!-- <div>
                    <label for="inventoryStatus" class="block font-bold mb-3">Inventory Status</label>
                    <Select id="inventoryStatus" v-model="product.inventoryStatus" :options="statuses"
                        optionLabel="label" placeholder="Select a Status" fluid></Select>
                </div> -->

                <!-- <div>
                    <span class="block font-bold mb-4">Category</span>
                    <div class="grid grid-cols-12 gap-4">
                        <div class="flex items-center gap-2 col-span-6">
                            <RadioButton id="category1" v-model="product.category" name="category"
                                value="Accessories" />
                            <label for="category1">Accessories</label>
                        </div>
                        <div class="flex items-center gap-2 col-span-6">
                            <RadioButton id="category2" v-model="product.category" name="category" value="Clothing" />
                            <label for="category2">Clothing</label>
                        </div>
                        <div class="flex items-center gap-2 col-span-6">
                            <RadioButton id="category3" v-model="product.category" name="category"
                                value="Electronics" />
                            <label for="category3">Electronics</label>
                        </div>
                        <div class="flex items-center gap-2 col-span-6">
                            <RadioButton id="category4" v-model="product.category" name="category" value="Fitness" />
                            <label for="category4">Fitness</label>
                        </div>
                    </div>
                </div> -->

                <!-- <div class="grid grid-cols-12 gap-4">
                    <div class="col-span-6">
                        <label for="price" class="block font-bold mb-3">Price</label>
                        <InputNumber id="price" v-model="product.price" mode="currency" currency="USD" locale="en-US"
                            fluid />
                    </div>
                    <div class="col-span-6">
                        <label for="quantity" class="block font-bold mb-3">Quantity</label>
                        <InputNumber id="quantity" v-model="product.quantity" integeronly fluid />
                    </div>
                </div> -->
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

const dataConnected = ref(false)
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
    { name: '2023/2024', code: '20232' },
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
</script>
