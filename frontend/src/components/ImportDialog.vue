<template>
        <!-- DIALOG IMPORT -->
        <Dialog v-model:visible="dialogImport" :style="{ width: '450px' }" header="Tambah Data" :modal="true">
        <div>
            <div class="mb-4">
                <label for="tahun-pelajaran" class="block text-sm font-medium text-gray-700">Tahun Pelajaran <span
                        class="text-red-500">*</span></label>
                <Select v-model="selectedSemester" :options="semester" optionLabel="namaSemester"
                    placeholder="Tahun Pelajaran" class="w-full mr-2" />
            </div>
            <div class="mb-4">
                <label class="block text-sm font-medium text-gray-700">Unggah File Excel (Pastikan sesuai dengan
                    Template yang disediakan)</label>
                <div class="mt-2 flex flex-col gap-6 items-center justify-center">
                    <FileUpload ref="fileupload" mode="basic" name="demo[]" url="/api/upload" accept="xlsx/*"
                        :maxFileSize="1000000" @upload="onUpload" severity="secondary" />
                </div>
                <p class="mt-2 text-sm text-gray-500">Unduh Template Import data Penerima Ijazah <a
                        href="http://localhost:8082/api/v1/ss/download/template?template-type=siswa"
                        class="text-indigo-600 hover:text-indigo-500">Disini</a></p>
            </div>
        </div>
        <template #footer>
            <Button label="Batal" icon="pi pi-times" text @click="dialogImport = false" />
            <Button label="Simpan" icon="pi pi-check" text @click="deletedKelas" />
        </template>
    </Dialog>
</template>

<script setup>
import Dialog from 'primevue/dialog';
import Button from 'primevue/button';
import FileUpload from 'primevue/fileupload';
import Select from 'primevue/select';

</script>