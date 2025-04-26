<script setup>
import { ref, computed, onMounted } from "vue";
import { useStore } from "vuex";
const store = useStore();
import Dialog from "primevue/dialog";
import Select from "primevue/select";
import FileUpload from "primevue/fileupload";
import Button from "primevue/button";

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
</script>

<template>
    <Toast />
    <Dialog v-model:visible="isVisible"  header="Tambah Data" :modal="true">
        <div>
            <DataTable ref="dt" v-model:selection="selectedKelas" stripedRows size="small" :value="kelasList" scrollable
                scrollHeight="400px" dataKey="rombonganBelajarId" :paginator="true" :rows="10" :filters="filters"
                tableStyle="min-width: 50rem"
                paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
                :rowsPerPageOptions="[10, 20, 30]"
                currentPageReportTemplate="Showing {first} to {last} of {totalRecords} kelas" class="mt-2">
                <Column selectionMode="multiple" style="width: 3rem;" :exportable="false"></Column>
                <Column field="nmKelas" header="Nama Siswa"></Column>
                <Column field="tingkatPendidikanId" header="Jk" sortable></Column>
                <Column field="kurikulum.namaKurikulum" header="nis"></Column>
                <!-- Jika SMK/MAK Program Keahlian & Kompetensi Keahlian akan muncul-->
                <div v-if="['smk', 'mak'].includes(bentukPendidikan)">
                    <Column field="namaJurusanSp" header="Jurusan" sortable></Column>
                </div>
                <Column field="ptk.nama" header="Wali kelas"></Column>
                <Column field="jumlahAnggota" header="Jml.Anggota"></Column>
                <Column header="Anggota Kelas">
                    <template #body="slotProps">
                        <Button icon="pi pi-bullseye" outlined rounded class="mr-2"
                            @click="dialogAnggotaRombel(slotProps.data)" />
                    </template>
                </Column>

            </DataTable>
        </div>

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
    </Dialog>
</template>
