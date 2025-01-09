<script setup>
import { ref, onMounted } from 'vue';

import Tabs from 'primevue/tabs';
import TabList from 'primevue/tablist';
import Tab from 'primevue/tab';
import TabPanels from 'primevue/tabpanels';
import TabPanel from 'primevue/tabpanel';

import InputText from 'primevue/inputtext';
const value = ref(null);


import DatePicker from 'primevue/datepicker';

// import { ProductService } from '@/service/ProductService';

onMounted(() => {
    // ProductService.getProductsMini().then((data) => (products.value = data));
});

import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import ColumnGroup from 'primevue/columngroup';   // optional
import Row from 'primevue/row';                   // optional

const products = ref([
    { code: "001", name: "Terser", category: "tes", quantity: "1" },
    { code: "002", name: "Terser", category: "tes", quantity: "1" },
    { code: "003", name: "Terser", category: "tes", quantity: "1" },
    { code: "004", name: "Terser", category: "tes", quantity: "1" },
    { code: "005", name: "Terser", category: "tes", quantity: "1" },
    { code: "006", name: "Terser", category: "tes", quantity: "1" },
]);
const trankrip = ref([
    { code: "001", name: "Terser", nilaiRata: "87" },
    { code: "002", name: "Terser", nilaiRata: "88" },
    { code: "003", name: "Terser", nilaiRata: "89" },
    { code: "004", name: "Terser", nilaiRata: "89" },
    { code: "005", name: "Terser", nilaiRata: "86" },
    { code: "006", name: "Terser", nilaiRata: "87" },
]);

// Button

import Button from 'primevue/button';

import 'primeicons/primeicons.css'

// Dialog

import Dialog from 'primevue/dialog';
const visible = ref(false);
const inputManual = ref(false);
// Upload

import FileUpload from 'primevue/fileupload';





</script>


<template>
    <div class="card">
        <Tabs value="0">
            <TabList>
                <Tab value="0">Data Ijazah</Tab>
                <Tab value="1">Data Transkrip</Tab>
                <!-- <Tab value="2">Header III</Tab> -->

            </TabList>
            <TabPanels>
                <TabPanel value="0">
                    <div class="card">
                        <div class="flex space-x-2">
                            <Button label="Import" icon="pi pi-upload" iconPos="right" @click="visible = true" />
                            <Button label="Tambah Data" icon="pi pi-plus" iconPos="left" class="bg-blue-500"
                                @click="inputManual = true" />
                            <!-- <Button label="Save" icon="pi pi-check" iconPos="right" />
                            <Button label="Save" icon="pi pi-check" iconPos="right" /> -->
                        </div>
                        <DataTable :value="products" tableStyle="min-width: 50rem">
                            <Column field="name" header="NISN" sortable style="width: 25%"></Column>
                            <Column field="name" header="NIS" sortable style="width: 25%"></Column>
                            <Column field="name" header="Nama" sortable style="width: 25%"></Column>
                            <Column field="code" header="Sekolah" sortable style="width: 25%"></Column>
                            <Column field="category" header="Nama Ortu/Wali" sortable style="width: 25%"></Column>
                            <Column field="quantity" header="Tpt.Lahir" sortable style="width: 25%"></Column>
                            <Column field="quantity" header="Tgl.Lahir" sortable style="width: 25%"></Column>
                        </DataTable>
                    </div>




                    <!--  -->
                </TabPanel>
                <TabPanel value="1">
                    <div class="card">
                        <div class="flex space-x-2">
                            <Button label="Import" icon="pi pi-upload" iconPos="right" />
                            <Button label="Tambah Data" icon="pi pi-plus" iconPos="left" class="bg-blue-500" />
                            <Button label="Tarik Data" icon="pi pi-download" iconPos="left" class="bg-blue-500" />
                        </div>
                        <DataTable :value="trankrip" tableStyle="min-width: 50rem">
                            <Column field="code" header="NISN" sortable style="width: 25%"></Column>
                            <Column field="code" header="NIS" sortable style="width: 25%"></Column>
                            <Column field="name" header="Nama" sortable style="width: 25%"></Column>
                            <Column field="nilaiRata" header="Rata-rata Nilai" sortable style="width: 25%"></Column>
                            <!-- <Column field="quantity" header="Quantity" sortable style="width: 25%"></Column> -->
                        </DataTable>
                    </div>
                </TabPanel>
                <!-- <TabPanel value="2">
                    <p class="m-0">
                        At vero eos et accusamus et iusto odio dignissimos ducimus qui blanditiis praesentium voluptatum
                        deleniti atque corrupti quos dolores et quas molestias excepturi sint occaecati cupiditate non
                        provident, similique sunt in culpa
                        qui officia deserunt mollitia animi, id est laborum et dolorum fuga. Et harum quidem rerum
                        facilis est et expedita distinctio. Nam libero tempore, cum soluta nobis est eligendi optio
                        cumque nihil impedit quo minus.
                    </p>
                </TabPanel> -->
            </TabPanels>
        </Tabs>
    </div>


    <div class="">
        <Dialog v-model:visible="visible" modal header="Import Data" :style="{ width: '35rem' }">
            <div class="">
                <div>
                    <label for="username" class="text-sm font-semibold">Unggah File Excel (Pastikan sesuai dengan
                        Template
                        yang
                        disediakan)</label>
                </div>
                <div>
                    <Toast />
                    <FileUpload ref="fileupload" mode="basic" name="demo[]" url="/api/upload" accept="image/*"
                        :maxFileSize="1000000" @upload="onUpload" />
                </div>
                <p class="text-sm">Unduh Template Import data Penerima Ijazah <a href="#"
                        class="text-blue-400">Disini</a></p>
            </div>
            <!-- <div class="flex items-center gap-4 mb-8">
                <label for="email" class="font-semibold w-24">Email</label>
                <InputText id="email" class="flex-auto" autocomplete="off" />
            </div> -->
            <div class="flex justify-end gap-2">
                <Button type="button" label="Cancel" severity="secondary" @click="visible = false"></Button>
                <Button type="button" label="Save" @click="visible = false"></Button>
            </div>
        </Dialog>
    </div>

    <!-- input data ijazah manual -->
    <div class="">
        <Dialog v-model:visible="inputManual" modal header="Input data manual" :style="{ width: '35rem' }">
            <div class="">
                <form>
                    <div>
                        <div>
                            <label for="">Nama</label>

                        </div>
                        <InputText type="text" v-model="value" />
                    </div>
                    <div>
                        <div>
                            <label for="">Tempat Lahir</label>
                        </div>
                        <InputText type="text" v-model="value" />
                    </div>
                    <div>
                        <div>
                            <label for="">Tanggal Lahir</label>
                        </div>
                        <DatePicker v-model="date" />
                    </div>
                    <div>
                        <div>
                            <label for="">Nama Orangtua/Wali</label>
                        </div>
                        <InputText type="text" v-model="value" />
                    </div>
                    <div>
                        <div>
                            <label for="">No. Ijazah</label>
                        </div>
                        <InputText type="text" v-model="value" />
                    </div>
                </form>
            </div>
            <!-- <div class="flex items-center gap-4 mb-8">
                <label for="email" class="font-semibold w-24">Email</label>
                <InputText id="email" class="flex-auto" autocomplete="off" />
            </div> -->
            <div class="flex justify-end gap-2">
                <Button type="button" label="Cancel" severity="secondary" @click="visible = false"></Button>
                <Button type="button" label="Save" @click="visible = false"></Button>
            </div>
        </Dialog>
    </div>




</template>
