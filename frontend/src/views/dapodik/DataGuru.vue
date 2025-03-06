<template>
    <div class="">
        <div class="card">
            <div v-if="dataConnected">
                <div class="fixed top-0 w-full left-0 z-20 bg-white">
                    <div class="lg:ml-[250px] my-2 ">
                        <div class="container ">
                            <div class="flex flex-wrap justify-between items-center mb-2">
                                <h4 class="font-bold text-xl md:text-2xl">Data Guru </h4>
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
                                        <Button label="Import" icon="pi pi-download" severity="warn"
                                            @click="dialogImport = true" class="mr-2" />
                                        <Button label="Export" icon="pi pi-upload" severity="help"
                                            @click="exportCSV($event)" class="mr-2" />
                                        <!-- <Button label="Proses" icon="pi pi-send" severity="info"
                                            @click="exportCSV($event)" v-tooltip.right="'Menyimpan ke database'"
                                            badge="2" /> -->
                                    </template>

                                </Toolbar>
                            </div>

                            <Toolbar>
                                <!-- <template #start>
                                    <div class="flex flex-wrap gap-2 items-center justify-between">
                                        <div class="flex">
                                            <MultiSelect v-model="selectedJurusan" :options="jurusan" optionLabel="name"
                                                filter placeholder="Jurusan" :maxSelectedLabels="1"
                                                class="w-full md:w-80 mr-2" showClear />
                                            <Select v-model="selectedTingkat" showClear :options="tingkat"
                                                optionLabel="name" placeholder="Tingkat" class="mr-2" />
                                        </div>
                                    </div>
                                </template> -->
                                <template #end>
                                    <IconField>
                                        <InputIcon>
                                            <i class="pi pi-search" />
                                        </InputIcon>
                                        <InputText v-model="filters['global'].value" placeholder="Cari nama..." />
                                    </IconField>
                                </template>
                            </Toolbar>
                        </div>
                    </div>
                </div>


                <DataTable ref="dt" v-model:selection="selectedKelas" stripedRows size="small" :value="rombel"
                    scrollable scrollHeight="400px" dataKey="ptkTerdaftarId" :paginator="true" :rows="10"
                    :filters="filters" tableStyle="min-width: 50rem"
                    paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
                    :rowsPerPageOptions="[10, 20, 30]"
                    currentPageReportTemplate="Showing {first} to {last} of {totalRecords} kelas" class="mt-56">
                    <Column selectionMode="multiple" style="width: 3rem;" :exportable="false"></Column>
                    <Column field="ptk.nama" header="Nama" sortable></Column>
                    <Column field="ptk.jenisKelamin" header="JK" sortable></Column>
                    <Column field="ptk.tempatLahir" header="Tpt Lahir"></Column>
                    <Column field="tglLahir" header="Tgl Lahir">
                        <template #body="slotProps">
                            {{ new Date(slotProps.data.ptk.tanggalLahir).toLocaleDateString("id-ID", {
                                day: "2-digit",
                                month: "2-digit", year: "numeric"
                            }) }}
                        </template>
                    </Column>
                    <Column field="ptk.nuptk" header="NUPTK"></Column>
                </DataTable>

            </div>
            <div v-else>
                <EmptyData @profileFetched="handleProfileFetched" @fetchError="handleFetchError" />
            </div>
        </div>
    </div>

    <!-- DIALOGBOX UNTUK TAMBAH DATA -->
    <Dialog v-model:visible="kelasDialog" :style="{ width: '450px' }" :header="judulHeader" :modal="true">
        <div class="">
            <div class="flex space-x-2">
                <div class="w-full mb-2">
                    <!-- <p>{{ product.name }}</p> -->
                    <label for="name" class="block font-bold">Nama</label>
                    <InputText id="name" v-model.trim="kelas.nama" required="true" :invalid="submitted && !kelas.nama"
                        fluid />
                    <small v-if="submitted && !kelas.nama" class="text-red-500">Nama harus diisi.</small>


                </div>
                <div>
                    <label for="name" class="block font-bold ">JK</label>
                    <Select v-model.trim="kelas.tingkat" showClear :options="tingkat" optionLabel="name"
                        placeholder="JK" class="mr-2" />
                    <!-- <InputText id="name" v-model.trim="product.name" required="true" 
                    :invalid="submitted && !product.name" fluid /> -->
                    <small v-if="submitted && !kelas.tingkat" class="text-red-500">Kelas is required.</small>
                </div>
            </div>
            <div class="flex">
                <div class="w-1/2">
                    <label for="name" class="block font-bold ">Wali kelas</label>
                    <Select v-model.trim="kelas.tingkat" showClear :options="tingkat" optionLabel="name"
                        placeholder="Tingkat" class="mr-2" fluid />
                    <!-- <InputText id="name" v-model.trim="product.name" required="true" 
                    :invalid="submitted && !product.name" fluid /> -->
                    <small v-if="submitted && !kelas.tingkat" class="text-red-500">Kelas is required.</small>
                </div>
                <div class="w-1/2">
                    <label for="name" class="block font-bold ">Jurusan</label>
                    <Select v-model="kelas.jurusan" showClear :options="jurusan" optionLabel="name"
                        placeholder="Jurusan" class="w-full mr-2" />
                    <!-- <InputText id="name" v-model.trim="product.price" required="true" 
                    :invalid="submitted && !product.price" fluid /> -->
                    <small v-if="submitted && !kelas.jurusan" class="text-red-500">Nilai is required.</small>
                </div>
            </div>
            <!--<div>
                <label for="name" class="block font-bold ">Thn Lulus</label>
                <InputText id="name" v-model.trim="product.category" required="true" 
                    :invalid="submitted && !product.category" fluid />
                <small v-if="submitted && !product.category" class="text-red-500">Thn lulus is required.</small>
            </div> -->
        </div>

        <template #footer>
            <Button label="Cancel" icon="pi pi-times" text @click="hideDialog" />
            <Button label="Save" icon="pi pi-check" @click="saveProduct" />
        </template>
    </Dialog>

    <Dialog v-model:visible="deleteKelasDialog" :style="{ width: '450px' }" header="Confirm" :modal="true">
        <div class="flex items-center gap-4">
            <i class="pi pi-exclamation-triangle !text-3xl" />
            <span v-if="product">Apakah data ini akan dihapus?</span>
        </div>
        <template #footer>
            <Button label="Tidak" icon="pi pi-times" text @click="deleteKelasDialog = false" />
            <Button label="Ya" icon="pi pi-check" text @click="deletedKelas" />
        </template>
    </Dialog>

    <!-- DIALOG IMPORT -->
    <DialogImport v-model:visible="dialogImport" :semester="semester" :selectedSemester="selectedSemester"
        @save="saveImport" @cancel="cancelImport" :downloadUrl="templateUrl" fileName="template_siswa.xlsx" />
</template>

<script setup>
import { ref, onMounted, watch } from 'vue';
import { useStore } from "vuex";
import DialogImport from '../../components/DialogImport.vue'
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
import RuangKelasService from '@/service/ProductService.js';


onMounted(() => {
    fetchSemester()
    fetchData()
    // RuangKelasService.getProducts().then((data) => (products.value = data));
});

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
// const statuses = ref([
//     { label: 'INSTOCK', value: 'instock' },
//     { label: 'LOWSTOCK', value: 'lowstock' },
//     { label: 'OUTOFSTOCK', value: 'outofstock' }
// ]);

// const formatCurrency = (value) => {
//     if (value)
//         return value.toLocaleString('en-US', { style: 'currency', currency: 'USD' });
//     return;
// };

const judulHeader = ref("")
const openNew = () => {
    judulHeader.value = "Tambah Data"
    product.value = {};
    selectedKelas.value = []
    submitted.value = false;
    kelasDialog.value = true;
};
const hideDialog = () => {
    kelasDialog.value = false;
    submitted.value = false;
};
const saveProduct = () => {
    submitted.value = true;
    if (kelas?.value.nama?.trim()) {
        // Untuk menyimpan edit
        if (kelas.value.id) {
            // kelas.value.tingkat = kelas.value.inventoryStatus.value ? kelas.value.inventoryStatus.value : kelas.value.inventoryStatus;
            rombel.value[findIndexById(product.value.id)] = kelas.value;
            toast.add({ severity: 'success', summary: 'Successful', detail: 'Product Updated', life: 3000 });
        }
        else {
            kelas.value.id = createId();
            kelas.value.code = createId();
            kelas.value.image = 'product-placeholder.svg';
            // product.value.inventoryStatus = product.value.inventoryStatus ? product.value.inventoryStatus.value : 'INSTOCK';
            rombel.value.push(kelas.value);
            toast.add({ severity: 'success', summary: 'Successful', detail: 'Product Created', life: 3000 });
        }

        kelasDialog.value = false;
        kelas.value = {};
    }
};
const editKelas = (prod) => {
    judulHeader.value = "Edit Data"
    // product.value = { ...prod };
    product.value = prod.reduce((acc, item) => {
        acc[item.name] = item.value;
        return acc;
    })
    kelasDialog.value = true;
};
// const confirmDeleteProduct = (prod) => {
//     product.value = prod;
//     deletekelasDialog.value = true;
// };
// const deleteProduct = () => {
//     products.value = products.value.filter(val => val.id !== product.value.id);
//     deletekelasDialog.value = false;
//     product.value = {};
//     toast.add({ severity: 'success', summary: 'Successful', detail: 'Product Deleted', life: 3000 });
// };
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
    deleteKelasDialog.value = true;
};
const deletedKelas = () => {
    rombel.value = rombel.value.filter(val => !selectedKelas.value.includes(val));
    deleteKelasDialog.value = false;
    selectedKelas.value = null;
    toast.add({ severity: 'success', summary: 'Successful', detail: 'Products Deleted', life: 3000 });
};

// const getStatusLabel = (status) => {
//     switch (status) {
//         case 'INSTOCK':
//             return 'success';

//         case 'LOWSTOCK':
//             return 'warn';

//         case 'OUTOFSTOCK':
//             return 'danger';

//         default:
//             return null;
//     }
// };



import Select from 'primevue/select';
import MultiSelect from 'primevue/multiselect';

import EmptyData from '@/components/EmptyData.vue';

const selectedSemester = ref();
const semester = ref(null);
const fetchSemester = async () => {
    try {
        const results = await store.dispatch("sekolahService/fetchSemester")
        // console.log(results)
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
watch(selectedSemester, (newVal, oldVal) => {
    console.log(newVal)
    fetchData()
})
const fetchData = async () => {
    try {
        let payload = {
            schema_name: "tabel_D4DA6B98FCFD71C58F5A",
            tahun_ajaran_id: selectedSemester.value?.tahunAjaranId
        }
        const results = await store.dispatch("sekolahService/fetchPTK", payload)
        console.log(results)
        rombel.value = results
        // if (results) {
        //     rombel.value = store.getters["sekolahService/fetchData"]
        //     // Ambil semester terbaru berdasarkan ID terbesar
        //     // selectedSemester.value = semester.value.reduce((latest, current) =>
        //     //     current.semesterId > latest.semesterId ? current : latest
        //     // );
        //     console.log()
        // }
    } catch (error) {

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
const selectedTingkat = ref();
const tingkat = ref([
    { name: '10' },
    { name: '11' },
    { name: '12' },
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
const dialogImport = ref(false)
const fileupload = ref();

const onUpload = () => {
    fileupload.value.upload();
};

const downloadTemplate = async () => {
    const response = await store.dispatch("sekolahService/getTemplate")
    console.log(response)
}

const dialogAnggotaRombel = (d) => {
    console.log(d)
}
const saveImport = () => {
    console.log("Data disimpan:", selectedSemester.value);
    dialogImport.value = false;
};

const cancelImport = () => {
    console.log("Import dibatalkan");
    dialogImport.value = false;
};
const templateUrl = ref("http://localhost:8183/api/v1/ss/download/template?template-type=siswa");

</script>
