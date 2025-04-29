<script setup>
import { ref, computed, onMounted } from "vue";
import { useStore } from "vuex";
const store = useStore();
import Dialog from "primevue/dialog";
import Select from "primevue/select";
import FileUpload from "primevue/fileupload";
import Button from "primevue/button";
import InputText from 'primevue/inputtext';

// ============toast============
import Toast from 'primevue/toast';
import { useToast } from "primevue/usetoast";
import DialogLoading from "./DialogLoading.vue";
import { isEmpty } from "lodash";
import { useExcelUpload } from "@/composables/useExcelUpload";
import { Column, DataTable } from "primevue";

const { excelData, loading, error, handleFileUpload, uploadToBackend } = useExcelUpload();
// const { excelData, loading, error, handleFileUpload } = useExcelUpload();
// import { isError } from "lodash";
const toast = useToast();

const isLoading = ref(false);

;
const selectedSemester = ref()


// ========================
// Props dari parent
const props = defineProps({
    visible: Boolean,
    // templateType: String,
    // schemaName: String,
});

// Emit event ke parent
const emit = defineEmits(["update:visible", "save", "cancel"]);

// Menggunakan computed agar bisa mengupdate prop.visible
const isVisible = computed({
    get: () => props.visible,
    set: (value) => emit("update:visible", value)
});

// Function untuk menutup dialog
const closeDialog = () => {
    isVisible.value = false;
};

// Function untuk mengunduh template
const isErr = ref(false)

const semester = computed(() => store.getters["sekolahService/getSemester"])

const pesertaDidik = ref({
    pesertaDidikId: '',
    nis: '',
    nisn: '',
    nmSiswa: '',
    tempatLahir: '',
    tanggalLahir: '',
    jenisKelamin: '',
    agama: '',
    alamatSiswa: '',
    teleponSiswa: '',
    diterimaTanggal: '',
    nmAyah: '',
    nmIbu: '',
    pekerjaanAyah: '',
    pekerjaanIbu: '',
    nmWali: '',
    pekerjaanWali: ''
});

const selectedjenisKelaminOptions = ref()
const jenisKelaminOptions = ref([
    { label: 'Laki-Laki', value: 'L' },
    { label: 'Perempuan', value: 'P' }
]);
const selectedAgamaOptions = ref()
const agamaOptions = ref([
    { label: 'Islam', value: 'Islam' },
    { label: 'Kristen', value: 'Kristen' },
    { label: 'Katolik', value: 'Katolik' },
    { label: 'Hindu', value: 'Hindu' },
    { label: 'Buddha', value: 'Buddha' },
    { label: 'Konghucu', value: 'Konghucu' }
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


import Tabs from 'primevue/tabs';
import TabList from 'primevue/tablist';
import Tab from 'primevue/tab';
import TabPanels from 'primevue/tabpanels';
import TabPanel from 'primevue/tabpanel';


</script>

<template>
    <Toast />
    <!-- <Dialog v-model:visible="isVisible"  header="Tambah Data" :modal="true">
        

<template #footer>
            <Button label="Batal" icon="pi pi-times" text @click="closeDialog" />
            <Button label="Simpan" icon="pi pi-check" text @click="uploadToBackend"
                :disabled="loading || !excelData.length" />
        </template>
</Dialog>

<DialogLoading v-model="isLoading"> Memuat data, harap tunggu... </DialogLoading>
<Dialog v-model:visible="isErr" header="Warning!">
    <div>
        Pilih <b>Tahun Pelajaran</b> terlebih dahulu!
    </div>
</Dialog> -->
    <Tabs value="0">
        <TabList>
            <Tab value="0">Ijazah</Tab>
            <Tab value="1">Transkrip</Tab>
        </TabList>
        <TabPanels>
            <TabPanel value="0">
                <div class="container bg-white p-8 rounded-lg shadow-md">
                    <section class="mb-8">
                        <h2 class="text-xl font-semibold mb-4">Informasi diri</h2>
                        <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
                            <div>
                                <label class="block text-gray-700" for="nmSiswa">Nama</label>
                                <InputText v-model="pesertaDidik.nmSiswa" fluid name="nmSiswa" id="nmSiswa"
                                    placeholder="Masukan nama" />
                            </div>
                            <div class="w-full">
                                <label class="block text-gray-700">Jenis Kelamin</label>
                                <Select v-model="selectedjenisKelaminOptions" :options="jenisKelaminOptions"
                                    placeholder="Pilih jenis kelamin" optionLabel="label" class="w-full" />
                            </div>
                            <div>
                                <div class="md:flex md:space-x-1">

                                    <div class="w-full">
                                        <label class="block text-gray-700" for="tempatLahir">Tpt Lahir</label>
                                        <InputText v-model="pesertaDidik.tempatLahir" fluid name="tempatLahir"
                                            id="tempatLahir" placeholder="Masukan tempat lahir"
                                            class=" w-full md:w-64" />
                                    </div>
                                    <div>
                                        <label class="block text-gray-700">Tgl Lahir</label>
                                        <input type="date" placeholder="YYYY-MM-DD"
                                            class=" w-full p-2 border border-gray-300 rounded"
                                            v-model="pesertaDidik.tanggalLahir">
                                    </div>
                                </div>
                            </div>

                            <div>
                                <label class="block text-gray-700">Agama</label>
                                <Select v-model="selectedAgamaOptions" :options="agamaOptions" placeholder="Pilih Agama"
                                    optionLabel="label" fluid class="w-full" />
                            </div>
                            <div>
                                <label class="block text-gray-700" for="nis">NIS</label>
                                <InputText v-model="pesertaDidik.nis" fluid name="nis" id="nis"
                                    placeholder="Masukan NIS" />
                            </div>
                            <div>
                                <label class="block text-gray-700" for="nisn">NISN</label>
                                <InputText v-model="pesertaDidik.nisn" fluid name="nisn" id="nisn"
                                    placeholder="Masukan NISN" />
                            </div>
                        </div>
                        <h2 class="text-xl font-semibold mb-4">Informasi Alamat</h2>
                        <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
                            <div>
                                <label class="block text-gray-700" for="nmSiswa">Alamat Jalan</label>
                                <InputText v-model="alamatLengkap.alamatJalan" fluid name="nmSiswa" id="nmSiswa"
                                    placeholder="Masukan nama" />
                            </div>
                            <div class="flex space-x-1">
                                <div class="w-1/2">
                                    <label class="block text-gray-700" for="rt">RT</label>
                                    <InputText v-model="alamatLengkap.rt" fluid name="rt" id="rt"
                                        placeholder="Masukan RT" />
                                </div>
                                <div class="w-1/2">
                                    <label class="block text-gray-700" for="rw">RW</label>
                                    <InputText v-model="alamatLengkap.rw" fluid name="rw" id="rw"
                                        placeholder="Masukan RW" />
                                </div>
                            </div>
                            <div>
                                <label class="block text-gray-700" for="prov">Prov.</label>
                                <InputText v-model="alamatLengkap.prov" fluid name="prov" id="prov"
                                    placeholder="Masukan nama" />
                            </div>
                            <div>
                                <label class="block text-gray-700" for="kab">Kab</label>
                                <InputText v-model="alamatLengkap.kab" fluid name="kab" id="kab"
                                    placeholder="Masukan nama" />
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
                        <div class="flex">
                            <div class="mb-4">
                                <label class="block text-gray-700">Phone Number</label>
                                <div class="relative">
                                    <input type="text" placeholder="Enter student's phone number"
                                        class="w-full p-2 border border-gray-300 rounded">
                                    <i class="fas fa-phone-alt absolute right-3 top-3 text-gray-400"></i>
                                </div>
                            </div>
                            <div class="mb-4">
                                <label class="block text-gray-700">Email</label>
                                <div class="relative">
                                    <input type="text" placeholder="Enter student's phone number"
                                        class="w-full p-2 border border-gray-300 rounded">
                                    <i class="fas fa-phone-alt absolute right-3 top-3 text-gray-400"></i>
                                </div>
                            </div>
                        </div>

                    </section>

                    <section class="mb-8">
                        <h2 class="text-xl font-semibold mb-4">Informasi Orang tua</h2>
                        <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
                            <div>
                                <label class="block text-gray-700" for="nmAyah">Nama Ayah Kandung</label>
                                <InputText v-model="pesertaDidik.nmAyah" fluid name="nmAyah" id="nmAyah"
                                    placeholder="Masukan nama" />
                            </div>
                            <div>
                                <label class="block text-gray-700">Pekerjaan Ayah</label>
                                <input type="text" placeholder="Enter father's occupation"
                                    class="w-full p-2 border border-gray-300 rounded">
                            </div>
                            <div>
                                <label class="block text-gray-700">Nama Ibu Kandung</label>
                                <input type="text" placeholder="Enter mother's name"
                                    class="w-full p-2 border border-gray-300 rounded">
                            </div>
                            <div>
                                <label class="block text-gray-700">Pekerjaan Ibu</label>
                                <input type="text" placeholder="Enter mother's occupation"
                                    class="w-full p-2 border border-gray-300 rounded">
                            </div>
                        </div>
                        <div class="mb-4">
                            <label class="block text-gray-700">Alamat Orang tua</label>
                            <textarea placeholder="Enter parents' address (if different from student)"
                                class="w-full p-2 border border-gray-300 rounded"></textarea>
                        </div>
                    </section>

                    <section class="mb-8">

                        <div>
                            <label class="block text-gray-700">Upload Ijazah</label>
                            <div class="relative">
                                <input type="file" class="w-full p-2 border border-gray-300 rounded">
                                <i class="fas fa-upload absolute right-3 top-3 text-gray-400"></i>
                            </div>
                        </div>
                    </section>

                    <div class="flex justify-end space-x-4">
                        <button class="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600"
                            @click="submitForm">Simpan</button>
                        <button class="bg-gray-300 text-gray-700 px-4 py-2 rounded hover:bg-gray-400"
                            @click="batal">Batal</button>
                    </div>

                </div>
            </TabPanel>
            <TabPanel value="1">
                <div>
                    <DataTable ref="dt" v-model:selection="selectedKelas" stripedRows size="small" :value="kelasList"
                        scrollable scrollHeight="400px" dataKey="rombonganBelajarId" :paginator="true" :rows="10"
                        :filters="filters" tableStyle="min-width: 50rem"
                        paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
                        :rowsPerPageOptions="[10, 20, 30]"
                        currentPageReportTemplate="Showing {first} to {last} of {totalRecords} kelas" class="mt-2">
                        <!-- <Column selectionMode="multiple" style="width: 3rem;" :exportable="false"></Column> -->
                        <Column field="nmKelas" header="Nama Mapel"></Column>
                        <Column field="tingkatPendidikanId" header="SMT 1"></Column>
                        <Column field="kurikulum.namaKurikulum" header="SMT 2"></Column>

                        <div v-if="['smk', 'mak'].includes(bentukPendidikan)">
                            <Column field="namaJurusanSp" header="Jurusan"></Column>
                        </div>
                        <Column field="ptk.nama" header="SMT 3"></Column>
                        <Column field="jumlahAnggota" header="SMT 4"></Column>
                        <Column field="jumlahAnggota" header="SMT 5"></Column>
                        <Column field="jumlahAnggota" header="SMT 6"></Column>
                        <Column header="Rata-rata">
                            <template #body="slotProps">
                                <Button icon="pi pi-bullseye" outlined rounded class="mr-2"
                                    @click="dialogAnggotaRombel(slotProps.data)" />
                            </template>
                        </Column>

                    </DataTable>
                </div>
            </TabPanel>


        </TabPanels>
    </Tabs>
</template>
