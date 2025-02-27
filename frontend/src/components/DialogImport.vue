<script setup>
import { ref, computed } from "vue";
import Dialog from "primevue/dialog";
import Select from "primevue/select";
import FileUpload from "primevue/fileupload";
import Button from "primevue/button";
// ============toast============
import Toast from 'primevue/toast';
import { useToast } from "primevue/usetoast";
const toast = useToast();
// ========================
// Props dari parent
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

// Menggunakan computed agar bisa mengupdate prop.visible
const isVisible = computed({
    get: () => props.visible,
    set: (value) => emit("update:visible", value)
});

// Function untuk menutup dialog
const closeDialog = () => {
    isVisible.value = false;
};

// Refs untuk FileUpload dan file yang diunggah
// const fileupload = ref();
const uploadedFiles = ref();
const uploadUrl = "http://localhost:8183/api/v1/ss/upload/rest"

// Function untuk menyimpan data
const saveData = async () => {
    if (uploadedFiles.value.files.length == 0) {
        toast.add({ severity: 'warn', summary: 'Gagal', detail: 'Silakan unggah file terlebih dahulu!', life: 3000 });
        return;
    }
    const file = uploadedFiles.value.file
    // console.log(uploadedFiles.value.files.length)
    const formData = new FormData();
    formData.append("file", file);

    try {
        const response = await fetch(uploadUrl, {
            method: "POST",
            body: formData,
        });

        if (!response.ok) {
            throw new Error("Gagal mengunggah file");
        }

        const result = await response.json();
        uploadedFiles.value.push(file);

        toast.add({ severity: 'success', summary: 'Sukses', detail: 'File berhasil diunggah!', life: 3000 });

        // Reset input file setelah upload selesai
        if (uploadedFiles.value) {
            uploadedFiles.value.clear();
        }
    } catch (error) {
        toast.add({ severity: 'error', summary: 'Error', detail: 'Gagal mengunggah file', life: 3000 });
        console.error("Upload error:", error);
    }
    // Reset input file setelah upload selesai
    if (uploadedFiles.value) {
        uploadedFiles.value.clear();
    }
    emit("save", uploadedFiles.value);
    closeDialog();
};
const upload = () => {
    uploadedFiles.value.upload()
}
// Handle sebelum upload (validasi file)
const onBeforeUpload = (event) => {
    const file = event.files[0]; // Ambil file pertama
    const allowedExtensions = ["xlsx"];
    const maxFileSize = 2 * 1024 * 1024; // 2MB

    const fileExtension = file.name.split(".").pop().toLowerCase();
    if (!allowedExtensions.includes(fileExtension)) {
        toast.add({ severity: 'error', summary: 'Error', detail: 'Format file harus .xlsx!', life: 3000 });
        return false;
    }

    if (file.size > maxFileSize) {
        toast.add({ severity: 'error', summary: 'Error', detail: 'Ukuran file tidak boleh lebih dari 2MB!', life: 3000 });
        return false;
    }

    return true; // Izinkan unggahan
};

// Handle upload file
const onUpload = async () => {
    console.log("sedang upload files")

    // uploadedFiles.value = event.files;
    toast.add({ severity: 'info', summary: 'Success', detail: 'File berhasil diunggah!', life: 3000 });

};
// Function untuk mengunduh template
const downloadTemplate = async () => {
    try {
        const response = await fetch(props.downloadUrl);

        if (!response.ok) {
            throw new Error("Gagal mengunduh file");
        }

        const blob = await response.blob();
        const url = window.URL.createObjectURL(blob);

        const a = document.createElement("a");
        a.href = url;
        a.download = props.fileName;
        document.body.appendChild(a);
        a.click();
        document.body.removeChild(a);

        window.URL.revokeObjectURL(url);
    } catch (error) {
        toast.add({ severity: 'error', summary: 'Error', detail: 'Terjadi kesalahan saat mengunduh file', life: 3000 });
    }
};
</script>

<template>
    <Toast />
    <Dialog v-model:visible="isVisible" :style="{ width: '450px' }" header="Tambah Data" :modal="true">
        <div>
            <div class="mb-4">
                <label class="block text-sm font-medium text-gray-700">
                    Tahun Pelajaran <span class="text-red-500">*</span>
                </label>
                <Select v-model="props.selectedSemester" :options="props.semester" optionLabel="namaSemester"
                    placeholder="Pilih Tahun Pelajaran" class="w-full mr-2" />
            </div>

            <div class="mb-4">
                <label class="block text-sm font-medium text-gray-700">
                    Unggah File Excel (Pastikan sesuai dengan Template yang disediakan)
                </label>
                <div class="mt-2 flex flex-col gap-6 items-center justify-center">
                    <FileUpload ref="uploadedFiles" mode="basic" name="file" accept=".xlsx" :maxFileSize="2000000"
                        :customUpload="true" @before-upload="onBeforeUpload" @upload="onUpload" severity="secondary" />
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
