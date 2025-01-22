<template>

    <div>

        <div class="card">
            <div class="flex flex-wrap justify-between my-2">
                <h4 class="font-bold text-2xl lg:text-lg my-2">Data Peserta Didik Penerima Ijazah </h4>
                <div>
                    <Select v-model="selectedsemester" :options="semester" optionLabel="namaSemester" placeholder="Tahun Pelajaran"
                        class="w-full md:w-56 mr-2" />
                </div>

            </div>
            <Toolbar class="">
                <template #start>

                </template>
                <template #end>
                    <Button label="Export" icon="pi pi-upload" severity="help" @click="exportCSV($event)" />
                </template>
            </Toolbar>

            <DataTable ref="dt" v-model:selection="dataLulusan" :value="students" dataKey="id" :paginator="true"
                :rows="10" :filters="filters"
                paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
                :rowsPerPageOptions="[5, 10, 25]"
                currentPageReportTemplate="Showing {first} to {last} of {totalRecords} products">
                <template #header>
                    <div class="flex flex-wrap gap-2 items-center justify-between">
                        <!-- <h4 class="m-0 font-bold text-2xl lg:text-lg">Data Peserta Didik Penerima Ijazah </h4> -->
                        <div>
                            <Select v-model="selectedJurusan" :options="jurusan" optionLabel="name"
                                placeholder="Jaringan" class="w-full md:w-56 mr-2" />
                        </div>
                        <div class="flex">

                            <Select v-model="selectedJurusan" :options="jurusan" optionLabel="name" placeholder="Rombel"
                                class="w-full md:w-56 mr-2" />
                            <IconField>
                                <InputIcon>
                                    <i class="pi pi-search" />
                                </InputIcon>
                                <InputText v-model="filters['global'].value" placeholder="Search..." />
                            </IconField>
                        </div>

                    </div>
                </template>

                <!-- <Column selectionMode="multiple" style="width: 3rem" pasundan123 :exportable="false"></Column> -->
                <Column field="contract_address" header="ID Transaksi">
                    <template #body="slotProps">
                        {{ slotProps.data.certificates[0].contract_address }}
                    </template>
                </Column>
                <Column field="transaction_hash" header="Number & Hash">
                    <template #body="slotProps">
                        {{ slotProps.data.certificates[0].transaction_hash }}
                    </template>
                </Column>

                <Column field="issue_date" header="Tgl.TRX" sortable>
                    <template #body="slotProps">
                        {{ formatDate(slotProps.data.certificates[0].issue_date) }}
                    </template>
                </Column>
                <Column field="blockchain_type" header="Jaringan" sortable>
                    <template #body="slotProps">
                        {{ slotProps.data.certificates[0].blockchain_type }}
                    </template>
                </Column>
                <Column field="nisn" header="NISN"></Column>
                <Column field="nis" header="NIS" sortable></Column>
                <Column field="nama" header="Nama" sortable></Column>
                <Column field="name" header="No Ijazah"></Column>
                <Column field="nilai_rata" header="Rerata Nilai">
                    <template #body="slotProps">
                        {{ slotProps.data.certificates[0].nilai_rata }}
                    </template>
                </Column>
                <Column field="category" header="Thn Lulus" sortable></Column>
            </DataTable>
        </div>

        <Dialog v-model:visible="productDialog" :style="{ width: '450px' }" header="Tambah data lulusan" :modal="true">
            <div class="flex flex-col gap-6">
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
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useStore } from 'vuex';
const store = useStore();
import axios from 'axios';
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

// import DataLulusanService from '@/service/DataLulusanService.js';
// Fungsi untuk mendapatkan semua siswa
// State untuk menyimpan daftar siswa dan error
const DataLulusanService = ref([]);
const error = ref("");

// Fetch data ketika komponen dimount
onMounted(() => {
    fetchSemester()
    // fetchStudents() 
});

const fetchStudents = async () => {
    try {
        // Reset error
        error.value = "";
        // Request ke endpoint backend
        const response = await axios.get('http://localhost:8081/api/v1/students').then(response => {
            console.log(response.data)
            students.value = response.data
            // students.value = transformData(response.data)
        })
        // Simpan data siswa
        // console.log(response.data)
        // DataLulusanService.value = response.data;
    } catch (err) {
        // Tangani error
        error.value =
            err.response?.data?.error || "An error occurred while fetching data.";
    }
}
// Fungsi untuk memproses data dari backend
const transformData = (data) => {
    return data.map((student) => {
        // Tambahkan properti dari certificates ke student
        const certificates = student.certificates.map((certificate) => ({
            ...student, // Salin semua properti student
            ...certificate, // Gabungkan dengan properti certificate
        }));
        return certificates; // Mengembalikan array objek hasil transformasi
    }).flat(); // Flatten hasil transformasi ke array satu dimensi
};


// Fungsi untuk memformat tanggal
const formatDate = (isoDate) => {
    const date = new Date(isoDate); // Buat objek Date
    const day = String(date.getDate()).padStart(2, "0"); // Ambil hari dengan 2 digit
    const month = String(date.getMonth() + 1).padStart(2, "0"); // Ambil bulan (0-indexed, tambah 1) dengan 2 digit
    const year = date.getFullYear(); // Ambil tahun
    return `${day}-${month}-${year}`; // Gabungkan menjadi format dd-mm-yyyy
}
const toast = useToast()
const dt = ref()
const students = ref()
const productDialog = ref(false)
const deleteProductDialog = ref(false)
const deleteProductsDialog = ref(false)
const product = ref({})
const dataLulusan = ref()
const filters = ref({
    'global': { value: null, matchMode: FilterMatchMode.CONTAINS },
})
const submitted = ref(false)
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


// select tahun ijazah
const selectedsemester = ref();
const semester = ref()
const fetchSemester = async () => {
    const param = {
        limit: 5,
        offset: 56,
    }
    try {
        const results = await store.dispatch("sekolahService/getSemester",param)
        // console.log(results)
        semester.value = results.semester
    } catch (error) {
        console.log("Gagal mengambil data semester", error)
    } finally {
        console.log("Fungsi fetchSemester selesai")
    }
}

const selectedJurusan = ref();
const jurusan = ref([
    { name: 'Teknik Kendaraan Ringan', code: 'TKR' },
    { name: 'Teknik Mesin Sepeda Motor', code: 'TSM' },
    { name: 'Teknik Komputer dan Jaringan', code: 'TKJ' },
    { name: 'Otomatisasi Perkantoran', code: 'OTKP' },
    { name: 'AKuntansi Lembaga', code: 'AKL' }
]);

</script>
