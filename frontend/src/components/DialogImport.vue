<script setup>
import { ref } from "vue";
import Dialog from "primevue/dialog";
import Select from "primevue/select";
import FileUpload from "primevue/fileupload";
import Button from "primevue/button";

// Props untuk menerima data dari parent
const props = defineProps({
    visible: Boolean,
    semester: Array,
    selectedSemester: Object,
    downloadUrl: {
        type: String,
        required: true
    },
    fileName: {
        type: String,
        default: "template.xlsx"
    }
});

// Emit event ke parent
const emit = defineEmits(["update:visible", "save", "cancel"]);

// Function untuk menutup dialog
const closeDialog = () => {
    emit("update:visible", false);
};

// Function untuk menyimpan data
const saveData = () => {
    emit("save");
};

const onUpload = () => {
    console.log("File berhasil diunggah");
};
const downloadTemplate = async () => {
    try {
        const response = await fetch("http://localhost:8082/api/v1/ss/download/template?template-type=siswa");

        if (!response.ok) {
            throw new Error("Gagal mengunduh file");
        }

        const blob = await response.blob(); // Mengubah response menjadi file Blob
        const url = window.URL.createObjectURL(blob); // Membuat URL objek

        // Membuat elemen `<a>` untuk mengunduh file secara otomatis
        const a = document.createElement("a");
        a.href = url;
        a.download = "template_siswa.xlsx"; // Nama file saat diunduh
        document.body.appendChild(a);
        a.click();
        document.body.removeChild(a);

        // Membersihkan URL objek setelah download selesai
        window.URL.revokeObjectURL(url);
    } catch (error) {
        console.error("Terjadi kesalahan saat mengunduh file:", error);
    }
};
</script>

<template>
    <Dialog v-model:visible="props.visible" :style="{ width: '450px' }" header="Tambah Data" :modal="true">
        <div>
            <div class="mb-4">
                <label class="block text-sm font-medium text-gray-700">
                    Tahun Pelajaran <span class="text-red-500">*</span>
                </label>
                <Select v-model="props.selectedSemester" :options="props.semester" optionLabel="namaSemester"
                    placeholder="Tahun Pelajaran" class="w-full mr-2" />
            </div>

            <div class="mb-4">
                <label class="block text-sm font-medium text-gray-700">
                    Unggah File Excel (Pastikan sesuai dengan Template yang disediakan)
                </label>
                <div class="mt-2 flex flex-col gap-6 items-center justify-center">
                    <FileUpload ref="fileupload" mode="basic" name="demo[]" url="/api/upload" accept="xlsx/*"
                        :maxFileSize="1000000" @upload="onUpload" severity="secondary" />
                </div>
                <p class="mt-2 text-sm text-gray-500">
                    Unduh Template Import data Penerima Ijazah
                    <a href="#" @click.prevent="downloadTemplate"
                        class="text-indigo-600 hover:text-indigo-500">Disini</a>
                </p>
            </div>
        </div>

        <template #footer>
            <Button label="Batal" icon="pi pi-times" text @click="closeDialog" />
            <Button label="Simpan" icon="pi pi-check" text @click="saveData" />
        </template>
    </Dialog>
</template>
