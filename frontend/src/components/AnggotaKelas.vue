<template>
    <div>
        <div class="card">
            <Toolbar class="mb-1">
                <template #start>
                    <Button label="New" icon="pi pi-plus" class="mr-2" @click="openNew" />
                    <Button label="Delete" icon="pi pi-trash" severity="danger" outlined @click="confirmDeleteSelected"
                        :disabled="!selectedSiswa || !selectedSiswa.length" class="mr-2" />
                    <Button label="Cari semua siswa" icon="pi pi-search" severity="info" @click="semuaSiswaOpen" />

                </template>

                <template #end>
                    <a href="#" @click.prevent="downloadTemplate"
                        class="text-indigo-600 hover:text-indigo-500 mr-2">Download template</a>
                    <!-- <FileUpload mode="basic" accept="image/*" :maxFileSize="1000000" label="Import" customUpload
                        chooseLabel="Import" class="mr-2" auto :chooseButtonProps="{ severity: 'secondary' }" /> -->
                    <input type="file" @change="handleFileUpload" />
                    <Button label="Upload" @click="uploadSiswa" />
                    <Button label="Export" icon="pi pi-upload" severity="secondary" @click="exportCSV($event)"
                        v-show="isExport" />
                </template>
            </Toolbar>

            <DataTable ref="dt" v-model:selection="selectedSiswa" :value="anggotaKelasList" dataKey="anggotaRombelId"
                :paginator="true" :rows="10" :filters="filters"
                paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
                :rowsPerPageOptions="[5, 10, 25]"
                currentPageReportTemplate="Showing {first} to {last} of {totalRecords} siswa">
                <template #header>
                    <div class="flex flex-wrap gap-2 items-center justify-end">
                        <IconField>
                            <InputIcon>
                                <i class="pi pi-search" />
                            </InputIcon>
                            <InputText v-model="filters['global'].value" placeholder="Search..." />
                        </IconField>
                    </div>
                </template>

                <Column selectionMode="multiple" style="width: 3rem" :exportable="false"></Column>
                <Column field="pesertaDidik.nmSiswa" header="Nama" sortable style="min-width: 16rem"></Column>
                <Column field="pesertaDidik.jenisKelamin" header="JK" sortable style="min-width: 6rem"></Column>
                <Column field="pesertaDidik.tempatLahir" header="Tpt Lahir" style="min-width: 12rem"></Column>
                <Column field="pesertaDidik.tanggalLahir" header="Tgl Lahir" style="min-width: 12rem"></Column>
                <Column field="pesertaDidik.agama" header="Agama" style="min-width: 12rem"></Column>
                <Column field="pesertaDidik.nis" header="NISN" style="min-width: 12rem"></Column>
                <Column field="pesertaDidik.nisn" header="NISN" style="min-width: 12rem"></Column>
                <Column field="pesertaDidik.sekolahAsal" header="Asal sekolah" style="min-width: 16rem"></Column>
            </DataTable>
        </div>

        <Dialog v-model:visible="siswaDialog" :style="{ width: '750px' }" header="Tambah siswa" :modal="true">
            <section class="mb-2">
                <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
                    <div>
                        <label class="block text-gray-700" for="nmSiswa">Nama</label>
                        <InputText v-model="pesertaDidik.nmSiswa" fluid name="nmSiswa" id="nmSiswa"
                            placeholder="Masukan nama" :invalid="submitted && !pesertaDidik.nmSiswa" />
                        <small v-if="submitted && !pesertaDidik.nmSiswa" class="text-red-500">Name is required.</small>
                    </div>
                    <div class="w-full">
                        <label class="block text-gray-700">Jenis Kelamin</label>
                        <Select v-model="selectedjenisKelaminOptions" :options="jenisKelaminOptions"
                            placeholder="Pilih jenis kelamin" optionLabel="label" class="w-full"
                            :invalid="submitted && !pesertaDidik.jenisKelamin" />
                        <small v-if="submitted && !pesertaDidik.jenisKelamin" class="text-red-500">Gender is
                            required.</small>

                    </div>
                    <div>
                        <div class="md:flex md:space-x-1">
                            <div class="w-full">
                                <label class="block text-gray-700" for="tempatLahir">Tpt Lahir</label>
                                <InputText v-model="pesertaDidik.tempatLahir" fluid name="tempatLahir" id="tempatLahir"
                                    placeholder="Masukan tempat lahir" class=" w-full md:w-64"
                                    :invalid="submitted && !pesertaDidik.tempatLahir" />
                                <small v-if="submitted && !pesertaDidik.tempatLahir" class="text-red-500">Place of birth
                                    is
                                    required.</small>

                            </div>
                            <div>
                                <label class="block text-gray-700">Tgl Lahir</label>
                                <input type="date" placeholder="DD-MM-YYYY"
                                    class=" w-full p-2 border border-gray-300 rounded"
                                    v-model="pesertaDidik.tanggalLahir">
                                <small v-if="submitted && !pesertaDidik.tanggalLahir" class="text-red-500">Date of birth
                                    is
                                    required.</small>
                            </div>
                        </div>
                    </div>

                    <div>
                        <label class="block text-gray-700">Agama</label>
                        <Select v-model="selectedAgamaOptions" :options="agamaOptions" placeholder="Pilih Agama"
                            optionLabel="label" fluid class="w-full" :invalid="submitted && !selectedAgamaOptions" />
                        <small v-if="submitted && !selectedAgamaOptions" class="text-red-500">Religion is
                            required.</small>

                    </div>
                    <div class="flex space-x-2">
                        <div>
                            <label class="block text-gray-700" for="nis">NIS</label>
                            <InputText v-model="pesertaDidik.nis" fluid name="nis" id="nis" placeholder="Masukan NIS"
                                :invalid="submitted && !pesertaDidik.nis" />
                            <small v-if="submitted && !pesertaDidik.nis" class="text-red-500">NIS is
                                required.</small>

                        </div>
                        <div>
                            <label class="block text-gray-700" for="nisn">NISN</label>
                            <InputText v-model="pesertaDidik.nisn" fluid name="nisn" id="nisn"
                                placeholder="Masukan NISN" :invalid="submitted && !pesertaDidik.nisn" />
                            <small v-if="submitted && !pesertaDidik.nisn" class="text-red-500">NISN is
                                required.</small>

                        </div>
                    </div>
                    <div>
                        <label class="block text-gray-700" for="asal-sekolah">Asal sekolah</label>
                        <InputText v-model="pesertaDidikPelengkap.sekolahAsal" fluid name="asal-sekolah"
                            id="asal-sekolah" placeholder="Masukan asal sekolah"
                            :invalid="submitted && !pesertaDidikPelengkap.sekolahAsal" />
                        <small v-if="submitted && !pesertaDidikPelengkap.sekolahAsal" class="text-red-500">Asal sekolah
                            is
                            required.</small>

                    </div>
                    <div>
                        <label class="block text-gray-700" for="nm-ibu">Nama Ibu Kandung</label>
                        <InputText v-model="pesertaDidik.nm_ibu" fluid name="nm-ibu" id="nm-ibu"
                            placeholder="Masukan nama" :invalid="submitted && !pesertaDidik.nm_ibu" />
                        <small v-if="submitted && !selectedAgamaOptions" class="text-red-500">Mother name is
                            required.</small>
                    </div>
                    <div>
                        <label class="block text-gray-700" for="nm-ayah">Nama Ayah Kandung</label>
                        <InputText v-model="pesertaDidik.nm_ayah" fluid name="nm-ayah" id="nm-ayah"
                            placeholder="Masukan nama" :invalid="submitted && !pesertaDidik.nm_ayah" />
                        <small v-if="submitted && !selectedAgamaOptions" class="text-red-500">Father name is
                            required.</small>
                    </div>
                </div>

                <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
                    <div>
                        <label class="block text-gray-700" for="alamatJln">Alamat Jalan</label>
                        <InputText v-model="alamatLengkap.alamatJalan" fluid name="alamatJln" id="alamatJln"
                            placeholder="contoh: Dusun Antah berantah; Perum.Antah berantah" />
                    </div>
                    <div class="flex space-x-1">
                        <div class="w-1/2">
                            <label class="block text-gray-700" for="rt">RT</label>
                            <InputText v-model="alamatLengkap.rt" fluid name="rt" id="rt" placeholder="Masukan RT" />
                        </div>
                        <div class="w-1/2">
                            <label class="block text-gray-700" for="rw">RW</label>
                            <InputText v-model="alamatLengkap.rw" fluid name="rw" id="rw" placeholder="Masukan RW" />
                        </div>
                    </div>
                    <div>
                        <label class="block text-gray-700" for="prov">Prov.</label>
                        <InputText v-model="alamatLengkap.prov" fluid name="prov" id="prov"
                            placeholder="Masukan nama" />
                    </div>
                    <div>
                        <label class="block text-gray-700" for="kab">Kab</label>
                        <InputText v-model="alamatLengkap.kab" fluid name="kab" id="kab" placeholder="Masukan nama" />
                    </div>
                    <div>
                        <label class="block text-gray-700" for="kec">Kecamatan</label>
                        <InputText v-model="alamatLengkap.kec" fluid name="kec" id="kec"
                            placeholder="Masukan nama kecamatan" />
                    </div>
                    <div>
                        <label class="block text-gray-700" for="desa">Desa</label>
                        <InputText v-model="alamatLengkap.desa" fluid name="desa" id="desa"
                            placeholder="Masukan nama desa" />
                    </div>
                </div>
            </section>
            <template #footer>
                <Button label="Cancel" icon="pi pi-times" text @click="hideDialog" />
                <Button label="Tambah" icon="pi pi-check" @click="addAnggotaSiswaFromInput" />
            </template>
        </Dialog>



        <Dialog v-model:visible="deleteSiswasDialog" :style="{ width: '450px' }" header="Confirm" :modal="true">
            <div class="flex items-center gap-4">
                <i class="pi pi-exclamation-triangle !text-3xl" />
                <span v-if="selectedSiswa">Yakin akan dihapus?</span>
            </div>
            <template #footer>
                <Button label="No" icon="pi pi-times" text @click="deleteSiswasDialog = false" />
                <Button label="Yes" icon="pi pi-check" text @click="deleteselectedSiswa" />
            </template>
        </Dialog>


        <!-- Cari siswa dialog -->
        <Dialog v-model:visible="semuaSiswaDialog" :style="{ width: '70%' }" header="Cari siswa" :modal="true"
            position="top">
            <div>
                <IconField>
                    <InputIcon>
                        <i class="pi pi-search" />
                    </InputIcon>
                    <InputText v-model="pesertaDidikQuery" placeholder="Cari berdasarkan nama..." class="w-full"
                        fluid />
                </IconField>
                <DataTable v-model:selection="pesertaDidikSelected" :value="resultsQuerySiswa" dataKey="pesertaDidikId">
                    <Column selectionMode="multiple" style="width: 3rem" :exportable="false"></Column>
                    <Column field="nmSiswa" header="Nama" sortable style="min-width: 16rem"></Column>
                    <Column field="jenisKelamin" header="JK" sortable style="min-width: 3rem"></Column>
                    <Column field="tempatLahir" header="Tpt Lahir" style="min-width: 8rem"></Column>
                    <Column field="tanggalLahir" header="Tgl Lahir" style="min-width: 3rem"></Column>
                    <Column field="nis" header="NISN" style="min-width: 6rem"></Column>
                    <Column field="nisn" header="NISN" style="min-width: 6rem"></Column>
                    <Column field="sekolahAsal" header="Asal sekolah" style="min-width: 16rem"></Column>
                </DataTable>

            </div>
            <template #footer>
                <Button label="Batal" icon="pi pi-times" text @click="batalSemuaSiswa" />
                <Button label="Tambahkan" icon="pi pi-check" @click="addSemuaSiswa" />
            </template>
        </Dialog>
        <!-- End of Cari siswa dialog -->



    </div>
</template>

<script setup>
import { ref, onMounted, watch, computed, onUnmounted } from 'vue';
import { useStore } from "vuex";
const store = useStore();
import debounce from 'lodash/debounce';

// import ProductService from '@/service/ProductService';
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
import Select from 'primevue/select';
import { onBeforeRouteLeave } from 'vue-router';
import { isEmpty } from 'lodash';

const toast = useToast();
const dt = ref();
// const products = ref();
const siswaDialog = ref(false);
// const deletesiswaDialog = ref(false);
const deleteSiswasDialog = ref(false);
// const product = ref({});
const pesertaDidik = ref({});
const pesertaDidikQuery = ref("")
const pesertaDidikSelected = ref([])
const pesertaDidikBaru = ref([])
const pesertaDidikAnggotaBaru = ref([])

const pesertaDidikPelengkap = ref({});
const selectedSiswa = ref();
const filters = ref({
    'global': { value: null, matchMode: FilterMatchMode.CONTAINS },
});
const submitted = ref(false);
let schemaname = ''
const selectedSemester = computed(() => store.getters["sekolahService/getSelectedSemester"])

const anggotaKelasList = ref([])

const openNew = () => {
    pesertaDidik.value = {};
    submitted.value = false;
    siswaDialog.value = true;
};

let fromManual = true
const semuaSiswaDialog = ref(false)
const resultsQuerySiswa = ref([])
const semuaSiswaOpen = () => {
    semuaSiswaDialog.value = true;
}
const cariSemuaSiswa = async (val) => {
    try {
        const payload = {
            schemaname: schemaname,
            nm_siswa: val
        }
        const results = await store.dispatch("sekolahService/searchSiswa", payload)
        resultsQuerySiswa.value = results.siswa
    } catch (error) {

    }
};

const addSemuaSiswa = async () => {
    let payloadBase = {
        schemaname: schemaname,
        semester_id: selectedSemester.value?.semesterId,
    };
    for (const e of pesertaDidikSelected.value) {
        const isExist = anggotaKelasList.value.some(
            (i) => i.pesertaDidikId === e.pesertaDidikId
        );
        if (!isExist) {
            fromManual = false;
            const payload = {
                ...payloadBase,
                peserta_didik_id: e.pesertaDidikId
            };

            try {
                const response = await store.dispatch(
                    "sekolahService/searchAnggotaKelas",
                    payload
                );
                // console.log(response)
                if (response.length === 0) {
                    const anggotaTes = {
                        anggotaRombelId: generateUUID(),
                        semesterId: selectedSemester.value?.semesterId,
                        rombonganBelajarId: props.rombonganBelajarId,
                        pesertaDidikId: e.pesertaDidikId,
                        pesertaDidik: e
                    }
                    // console.log(anggotaTes)
                    anggotaKelasList.value.push(anggotaTes);
                    pesertaDidikAnggotaBaru.value.push(anggotaTes)
                    // Lakukan sesuatu dengan respons jika perlu
                    toast.add({
                        severity: 'success',
                        summary: 'Berhasil',
                        detail: `Siswa ${e.nmSiswa} ditambahkan`,
                        life: 3000
                    });
                } else {
                    toast.add({
                        severity: "error",
                        summary: "Gagal",
                        detail: `${e.nmSiswa} sudah ada di kelas ${response[0]?.rombonganBelajar.nmKelas} `,
                        life: 3000,
                    });
                    localStorage.removeItem("unsavedPesertaDidik");
                    clearSearch()

                    return
                }
            } catch (error) {
                console.error("Gagal mengambil data anggota kelas:", error);
                toast.add({
                    severity: "error",
                    summary: "Gagal",
                    detail: `Gagal menambahkan ${e.nmSiswa}`,
                    life: 3000,
                });
            }
        } else {
            toast.add({
                severity: "error",
                summary: "Gagal",
                detail: `Siswa A.N. ${e?.nmSiswa.toLowerCase().replace(/\b\w/g, (char) => char.toUpperCase())} sudah ada di kelas ini`,
                life: 3000,
            });

        }
    }


    let savedData = JSON.parse(localStorage.getItem("unsavedPesertaDidik"));
    console.log(savedData === null)
    if (savedData === null) {
        localStorage.setItem("unsavedPesertaDidik", JSON.stringify(pesertaDidikAnggotaBaru.value));
        pesertaDidikAnggotaBaru.value = []
    } else {
        savedData.forEach(item => pesertaDidikAnggotaBaru.value.push(item))
        localStorage.setItem("unsavedPesertaDidik", JSON.stringify(pesertaDidikAnggotaBaru.value));
        pesertaDidikAnggotaBaru.value = []

    }

    clearSearch()
}

const addAnggotaSiswaFromInput = () => {
    submitted.value = true;
    pesertaDidik.value.pesertaDidikId = generateUUID()
    pesertaDidikSelected.value.push(pesertaDidik.value)
    addSemuaSiswa()
    addPesertaDidikBaru()
    siswaDialog.value = false;
    pesertaDidikPelengkap.value = {}
    selectedAgamaOptions.value = {}
    selectedjenisKelaminOptions.value = {}
    pesertaDidik.value = {};
};

import { useExcelUpload } from "@/composables/useExcelUpload";
const { handleFileUpload, uploadToStore, error, excelData } = useExcelUpload();

const addAnggotaSiswaFromSearch = () => {

}
const addAnggotaSiswaFromImpor = () => {

}

const addPesertaDidikBaru = () => {
    pesertaDidikBaru.value.push(pesertaDidik.value)
    localStorage.setItem("unsavedPesertaDidikBaru", JSON.stringify(pesertaDidikBaru.value));
}
const batalSemuaSiswa = () => {
    semuaSiswaDialog.value = false
    clearSearch()
}

const clearSearch = () => {
    pesertaDidikQuery.value = ""
    pesertaDidikSelected.value = []
    resultsQuerySiswa.value = []
}


const hideDialog = () => {
    siswaDialog.value = false;
    submitted.value = false;
};
// const emit = defineEmits(["rombelAnggota"]);

// const deleteSiswa = () => {
//     pesertaDidik.value = anggotaKelasList.value.filter(val => val.pesertaDidikId !== product.value.pesertaDidikId);
//     deletesiswaDialog.value = false;
//     pesertaDidik.value = {};
//     // emit("rombelAnggota", anggotaKelasList)
//     localStorage.setItem("unsavedPesertaDidik", JSON.stringify(anggotaKelasList.value));
//     toast.add({ severity: 'success', summary: 'Successful', detail: 'Student Deleted', life: 3000 });
// };
// const findIndexById = (id) => {
//     let index = -1;
//     for (let i = 0; i < products.value.length; i++) {
//         if (products.value[i].id === id) {
//             index = i;
//             break;
//         }
//     }

//     return index;
// };
const generateUUID = () => crypto.randomUUID();
// const createId = () => {
//     return generateUUID()
// }
const exportCSV = () => {
    dt.value.exportCSV();
};
const confirmDeleteSelected = () => {
    deleteSiswasDialog.value = true;
};
const deleteselectedSiswa = () => {
    anggotaKelasList.value = anggotaKelasList.value.filter(val => !selectedSiswa.value.includes(val));
    deleteSiswasDialog.value = false;
    // localStorage.setItem("unsavedPesertaDidik", JSON.stringify(anggotaKelasList.value));
    const payload = {
        schemaname: schemaname,
    }
    selectedSiswa.value.forEach(async (item) => {
        payload.anggota_rombel_id = item?.anggotaRombelId
        const response = await store.dispatch("sekolahService/deleteAnggotaKelas", payload)
        if (response) {
            toast.add({ severity: 'success', summary: 'Successful', detail: `${item.pesertaDidik?.nmSiswa} telah dihapus`, life: 3000 });
            selectedSiswa.value = null;

        }

    })
};

// ================================
// Props dari parent
const props = defineProps({
    visible: Boolean,
    semester: Array,
    templateType: String,
    rombonganBelajarId: String,
    isEdit: Boolean,
});

// ambil data anggota kelas
const fetchAnggotaKelas = async () => {
    const payload = {
        schemaname: await store.getters["sekolahService/getTabeltenant"]?.schemaname,
        semester_id: selectedSemester.value?.semesterId,
        rombongan_belajar_id: props.rombonganBelajarId,
    }
    const results = await store.dispatch("sekolahService/fetchAnggotaKelas", payload)
    results.forEach(item => {
        anggotaKelasList.value.push(item)
    })
}




onMounted(() => {
    //ProductService.getProducts().then((data) => (products.value = data));
    if (props.rombonganBelajarId) {
        fetchAnggotaKelas()
    }
    const savedData = JSON.parse(localStorage.getItem("unsavedPesertaDidik"));
    if (savedData) {
        // console.log(savedData)
        savedData.forEach(item => anggotaKelasList.value.push(item));
    }
    schemaname = store.getters["sekolahService/getTabeltenant"]?.schemaname
});
const selectedAgamaOptions = ref()
const agamaOptions = ref([
    { label: 'Islam', value: 'Islam' },
    { label: 'Kristen', value: 'Kristen' },
    { label: 'Katolik', value: 'Katolik' },
    { label: 'Hindu', value: 'Hindu' },
    { label: 'Buddha', value: 'Buddha' },
    { label: 'Konghucu', value: 'Konghucu' }
]);

const selectedjenisKelaminOptions = ref()
const jenisKelaminOptions = ref([
    { label: 'Laki-Laki', value: 'L' },
    { label: 'Perempuan', value: 'P' }
]);


const alamatLengkap = ref({
    alamatJalan: '',
    rt: '',
    rw: '',
    desa: '',
    kec: '',
    kab: '',
    prov: ''
})

watch(selectedjenisKelaminOptions, (newVal) => {
    pesertaDidik.value.jenisKelamin = newVal?.value
})
watch(selectedAgamaOptions, (newVal) => {
    pesertaDidik.value.agama = newVal?.value
})

watch(pesertaDidikQuery, debounce((newVal) => {
    if (newVal.length < 2) {
        return
    }
    // console.log(newVal)
    cariSemuaSiswa(newVal)
}, 500))
// import { onBeforeRouteLeave } from "vue-router";
onBeforeRouteLeave((to, from, next) => {

    if (localStorage.getItem("unsavedPesertaDidik")) {
        const confirmLeave = confirm("Data belum disimpan. Apakah Anda ingin meninggalkan halaman?");
        if (confirmLeave) {
            localStorage.removeItem("unsavedPesertaDidik"); // Hapus data jika user mengabaikan
            localStorage.removeItem("unsavedPesertaDidikBaru"); // Hapus data jika user mengabaikan
            next();
        } else {
            next(false);
        }
    } else {
        next();
    }
})

const isExport = computed(() => {
    if (typeof anggotaKelasList.value === 'undefined') {
        return false
    } else if (typeof anggotaKelasList.value === 'object') {
        if (isEmpty(anggotaKelasList.value)) {
            return false
        } else {
            return true
        }
    } else {
        return true

    }
})

const baseUrl = "http://localhost:8183/api/v1/ss"; // Disimpan di child
const templateType = 'siswa'
const templateUrl = computed(() => {
    return `${baseUrl}/download/template?template_type=${templateType}&schemaname=${schemaname}&semesterId=${selectedSemester.value?.semesterId}`;
});
const downloadTemplate = async () => {
    // if (isEmpty(selectedSemester.value)) {
    //     // alert("Pilih tahun pelajaran")
    //     isErr.value = true
    //     return
    // }
    try {
        const response = await fetch(templateUrl.value, {
            method: "GET",
            headers: {
                "Accept": "application/octet-stream",
            },
        });

        if (!response.ok) {
            throw new Error("Gagal mengunduh file");
        }

        // Coba ambil nama file dari header Content-Disposition
        const contentDisposition = response.headers.get("Content-Disposition");
        let fileName = "downloaded_file.xlsx"; // Default jika tidak ditemukan
        if (contentDisposition) {
            const match = contentDisposition.match(/filename="([^"]+)"/);
            if (match && match[1]) {
                fileName = match[1];
            }
        }

        const blob = await response.blob();
        const url = window.URL.createObjectURL(blob);

        const a = document.createElement("a");
        a.href = url;
        a.download = fileName;
        document.body.appendChild(a);
        a.click();
        document.body.removeChild(a);

        window.URL.revokeObjectURL(url);
    } catch (error) {
        toast.add({ severity: 'error', summary: 'Error', detail: 'Terjadi kesalahan saat mengunduh file', life: 3000 });
    }
};

const uploadSiswa = () => {
    uploadToStore("sekolahService/CreateBanyakSiswa", "siswa");
};

// Mapping function
const mapExcelToPesertaDidikArray = (data) => {
    // Lewati baris pertama jika itu header
    const rows = data.slice(1);

    pesertaDidik.value = rows.map((row) => ({
        pesertaDidikId: row[0] || '',
        nis: row[1] || '',
        nisn: row[2] || '',
        nmSiswa: row[3] || '',
        tempatLahir: row[4] || '',
        tanggalLahir: row[5] || '',
        jenisKelamin: row[6] || '',
        agama: row[7] || '',
        alamatSiswa: row[8] || '',
        teleponSiswa: row[9] || '',
        diterimaTanggal: row[10] || '',
        nmAyah: row[11] || '',
        nmIbu: row[12] || '',
        pekerjaanAyah: row[13] || '',
        pekerjaanIbu: row[14] || '',
        nmWali: row[15] || '',
        pekerjaanWali: row[16] || '',
    }));
};

// Pantau perubahan excelData
watch(excelData, (newData) => {
    console.log(newData.length > 0)
    if (newData.length > 0) {
        mapExcelToPesertaDidikArray(newData);
    }
});

</script>