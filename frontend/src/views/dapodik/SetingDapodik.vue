<script setup>
import axios from 'axios';
import { onMounted, ref } from 'vue';


import Select from 'primevue/select';

import ProgressBar from 'primevue/progressbar';
import Button from 'primevue/button';
import Dialog from 'primevue/dialog';


const dataDapodikDefault = ref({
    baseUrl: "localhost",
    port: "5774",
    username: "d21xperience@gmail.com",
    password: "Pasja123*",
    semester_id: "20221",
    sekolah_id: "",
    semester_id: ""
})

// Data didapat dari Database
// const semester = ref([
//     {
//         id: '20231',
//         nama: '2022/2023 Ganjil',
//     },
//     {
//         id: '20232',
//         nama: '2022/2023 Ganjil',
//     },
//     {
//         id: '20221',
//         nama: '2021/2022 Ganjil',
//     },
//     {
//         id: '20222',
//         nama: '2021/2022 Genap',
//     },
//     {
//         id: '20231',
//         nama: '2022/2023 Ganjil',
//     },

// ])
const semester = ref(null)
const isConnect = ref(false)

const connecting = () => {
    // isConnect.value = !isConnect.value
    if (!isConnect) {
        alert("Aplikasi dapodik sudah terhubung")
    } else {
        alert("Aplikasi dapodik belum terhubung")
    }
}


onMounted(() => {
    cekKoneksiDapodik()
})

const api = axios.create({
    baseURL: 'http://localhost:8082',
    timeout: 5000,
    headers: {
        'Content-Type': 'application/json',
    },

});
const data = ref(null)
const cekDapodikNetwork = ref(false)
const userHasLoginToDapodik = ref(true)
const loading = ref(false)
const cekKoneksiDapodik = async () => {
    try {
        const response = await api.get("/api/v1/dapodik", {
            params: {
                ip_dapodik: "localhost",
                port: 5774
            }

        })

        // console.log(response.data)
        loading.value = true
        if (response.status == 200) {
            data.value = response.data
            semester.value = data.value.data.data
            cekDapodikNetwork.value = true
        }
        setTimeout(() => {
            loading.value = false;
        }, 2000);
    } catch (error) {
        console.log('err', error)
        dialogConnectionDapodik.value = true
        switch (error.code) {
            case "ERR_BAD_RESPONSE":
                errorMessageConnectionDapodik.value = "Aplikasi Dapodik tidak terdeteksi"
                break;
            case "ERR_NETWORK":
                "Server perantara tidak aktif, silahkan aktifkan terlebih dahulu"
                break;
            default:
                break;
        }
        // if (error.code === "ERR_NETWORK") {
        //     errorMessageConnectionDapodik.value = "Server perantara tidak aktif, silahkan aktifkan terlebih dahulu"

        // }
    }
}

// Dialog
const dialogConnectionDapodik = ref(false)
const errorMessageConnectionDapodik = ref("")

const loginDapodik = async () => {
    try {
        const resp = await api.post("/api/v1/dapodik/login", {
            params: {
                username: dataDapodikDefault.username,
                password: dataDapodikDefault.password,
                semester_id: dataDapodikDefault.semester_id
            }
        })
        // userHasLoginToDapodik.value = true
        // console.log(resp.statusText)
        switch (resp.statusText) {
            case "OK":
                userHasLoginToDapodik.value = true
                break;

            default:
                break;
        }

        console.log(resp)
    } catch (error) {
        console.log("Terjadi Error: ", error)
    }
}

const getPesertaDidik = async () => {
    try {
        const resp = await api.get("/api/v1/dapodik/GetPesertaDidik")
        console.log(resp)


    } catch (error) {
        console.log(error)
    }
}
</script>



<template>
    <div class="p-2">
        <div v-if="cekDapodikNetwork">
            <div v-if="userHasLoginToDapodik">
                <h3>Welcome back, Jhon Doe. 🖐</h3>
                <p>Data Dapodik sudah tersedia, silakan pilih semester yang ingin diakses.</p>

                <div>
                    <table class="table table-condensed">
                        <thead class="bg-primary text-light">
                            <tr>
                                <th>No</th>
                                <th>Jenis Data</th>
                                <th>Data Lokal Rapor</th>
                                <th>Update</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr>
                                <td>1</td>
                                <td>Data Semester</td>
                                <td>2023/2024 Genap</td>
                                <td>

                                    <button type="button" class="btn btn-success btn-sm" onclick="updatesemester()"><i
                                            class="fa fa-refresh"></i> Update Data Semester</button>
                                </td>
                            </tr>

                            <tr>
                                <td>2</td>
                                <td>Data Sekolah</td>
                                <td>
                                    SMKS PASUNDAN JATINANGOR </td>
                                <td>

                                    <button type="button" class="btn btn-success btn-sm" onclick="update('sekolah')"><i
                                            class="fa fa-refresh"></i> Update Data Sekolah</button>
                                </td>
                            </tr>
                            <tr>
                                <td>3</td>
                                <td>Data Referensi Mapel Dapo</td>
                                <td>
                                    5508 Data </td>
                                <td>

                                    <button type="button" class="btn btn-success btn-sm" onclick="update('mapel')"><i
                                            class="fa fa-refresh"></i> Update Data Referensi</button>
                                </td>
                            </tr>
                            <tr>
                                <td rowspan="3">4</td>
                                <td>Data Guru</td>
                                <td>65 Data</td>
                                <td rowspan="3" style="vertical-align: middle;">
                                    <button type="button" class="btn btn-success btn-sm" onclick="update('gtk')"><i
                                            class="fa fa-refresh"></i> Update Data Guru</button>


                                </td>
                            </tr>
                            <tr>

                                <td>Data Pelengkap Guru</td>
                                <td>65 Data</td>

                            </tr>
                            <tr>

                                <td>Data Guru Terdaftar Tahun Ajaran 2023/2024 Genap</td>
                                <td>64 Data</td>

                            </tr>
                            <tr>
                                <td rowspan="2">5</td>
                                <td>Data Siswa</td>
                                <td>1195 Data</td>
                                <td rowspan="2" style="vertical-align: middle;">
                                    <button type="button" class="bg-blue-600 text-sm text-white p-2 rounded-xl"
                                        @click="getPesertaDidik"><i class="fa fa-refresh"></i> Update Data
                                        Siswa</button>



                                </td>
                            </tr>
                            <tr>

                                <td>Data Pelengkap Siswa</td>
                                <td>1197 Data</td>

                            </tr>
                            <tr>
                                <td rowspan="4">6</td>
                                <td>Data Rombel Tahun Ajaran 2023/2024 Genap</td>
                                <td>42 Data</td>
                                <td rowspan="4" style="vertical-align: middle;">
                                    <button type="button" class="btn btn-success btn-sm" onclick="update('rombel')"><i
                                            class="fa fa-refresh"></i> Update Data Transaksional</button>

                                </td>
                            </tr>
                            <tr>

                                <td>Data Anggota Rombel Tahun Ajaran 2023/2024 Genap</td>
                                <td>1178 Data</td>

                            </tr>
                            <tr>

                                <td>Data Pembelajaran Tahun Ajaran 2023/2024 Genap</td>
                                <td>289 Data</td>

                            </tr>
                            <tr>

                                <td>Data Kelas Ekskul</td>
                                <td>12 Data</td>

                            </tr>


                        </tbody>
                    </table>
                </div>
            </div>
            <div v-else>
                <div class="flex justify-center items-center">
                    <div class="p-2 rounded-md  w-96">
                        <form action="" class=" border p-6">
                            <h2 class="text-center text-2xl mb-2">Login</h2>
                            <div class="flex items-center justify-between mb-2">
                                <label for="username">Username</label>
                                <input type="text" name="username" id="username" v-model="dataDapodikDefault.username"
                                    class="border p-2">
                            </div>
                            <div class="flex items-center justify-between mb-2">
                                <label for="password">Password</label>
                                <input type="password" name="password" id="password"
                                    v-model="dataDapodikDefault.username" class="border p-2">
                            </div>
                            <div class="flex items-center justify-between mb-2">
                                <label for="semester" class="w-2/4">Semester</label>
                                <select name="semester" id="semester" class="w-2/4 p-2 outline-none border"
                                    v-model="dataDapodikDefault.semester_id">
                                    <option v-for="(smt, index) in semester" :key="index" :value="smt">{{ smt }}
                                    </option>
                                </select>
                            </div>
                            <div class="mt-2">
                                <button @click="loginDapodik"
                                    class="bg-blue-300 p-2 rounded-lg w-full hover:bg-blue-600 hover:text-white"
                                    type="button">Login</button>
                            </div>
                        </form>
                    </div>
                </div>
            </div>

            <!-- <div>
                <div class="card-body">
                    <div class="text-end">
                    </div>
                    <div class="my-3 border-top"></div>

                    <table class="table-auto border-collapse border border-slate-400" id="tb-koneksi">
                        <thead class="bg-primary text-light">
                            <tr>
                                <th>No</th>
                                <th>Nama Koneksi</th>
                                <th>IP Address Server Rapor SP</th>
                                <th>IP Address Server Dapodik</th>
                                <th>Key</th>
                                <th>NPSN Sekolah</th>
                                <th>Opsi</th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr>
                                <td>1</td>
                                <td>RAPORSP</td>
                                <td>localhost</td>
                                <td>localhost</td>
                                <td>mkegY2Ps2W7b4HM</td>
                                <td>20254180</td>

                                <td>
                                    <button type="button" class="btn btn-outline-warning btn-sm"
                                        onclick="edit_koneksi('ded48229-f449-43fb-9e91-60d930ee6b74')"><i
                                            class="fa fa-edit"></i> Edit </button>
                                    <button type="button" class="btn btn-outline-danger btn-sm"
                                        onclick="hapus_koneksi('ded48229-f449-43fb-9e91-60d930ee6b74')"><i
                                            class="fa fa-trash"></i> Hapus </button>

                                </td>
                            </tr>

                            <tr>
                                <td>&nbsp;</td>
                                <td>&nbsp;</td>
                                <td>&nbsp;</td>
                                <td>&nbsp;</td>
                                <td>&nbsp;</td>
                                <td>&nbsp;</td>
                                <td>&nbsp;</td>
                            </tr>
                        </tbody>
                        <tfoot>

                            <tr>
                                <td colspan="7">
                                    <div class="text-end"> <button type="button" class="btn btn-success btn-sm"
                                            onclick="tes_koneksi()"><i class="fa fa-refresh"></i> Tes Koneksi </button>
                                    </div>
                                </td>
                            </tr>
                        </tfoot>
                    </table>
                </div>
            </div> -->
            <!-- Daftar syncron -->
        </div>
        <div v-else class="grid place-items-center h-screen">
            <div class="border lg:w-[500px] p-4 min-w-72">
                <div class="flex text-center justify-center items-center space-x-2 mb-6">
                    <i class="pi pi-exclamation-triangle !text-3xl"></i>
                    <h3 class="lg:text-md font-medium"> Sistem tidak mendeteksi
                        aplikasi Dapodik</h3>
                </div>
                <p class="text-sm text-red-700">*Silahkan isi IP (localhost untuk Dapodik yang diinstall dalam satu
                    komputer) dan port dapodik untuk mengecek koneksi </p>
                <div class="flex flex-col my-2">
                    <label for="ip-dapodik" class=text-sm>IP Dapodik</label>
                    <div class="flex justify-between space-x-3 mt-1">
                        <input class="w-3/4 border p-2" type="text" placeholder="Masukan IP dapodik"
                            v-model="dataDapodikDefault.baseUrl" id="ip-dapodik">
                        <input class="w-1/4 border p-2" type="text" placeholder="Port dapodik"
                            v-model="dataDapodikDefault.port">
                    </div>
                </div>
                <div>
                    <!-- <Button type="button" label="Search" icon="pi pi-search" :loading="loading" @click="cekKoneksiDapodik" /> -->
                    <Button type="button" @click="cekKoneksiDapodik" icon="pi pi-search" :loading="loading"
                        label="Cek Koneksi"
                        class="bg-blue-600 text-white py-2 px-4 rounded-full hover:bg-blue-900 cursor-pointer" />

                </div>

            </div>
        </div>
    </div>

    <!-- Dialog start -->
    <Dialog v-model:visible="dialogConnectionDapodik" modal header="Error">
        <div class="text-center">
            <div class="flex justify-center items-center space-x-2 mb-1">
                <i class="pi pi-exclamation-triangle !text-3xl"></i>
                <h3 class="lg:text-md font-medium">{{ errorMessageConnectionDapodik }}</h3>
            </div>
            <!-- <p>{{ errorMessageConnectionDapodik }}</p> -->
            <!-- <p class="text-sm">Silahkan cek koneksi ke Dapodik</p> -->
            <div class="my-2">
                <button @click="dialogConnectionDapodik = false"
                    class="bg-blue-600 text-white py-2 px-4 rounded-full hover:bg-blue-900 cursor-pointer w-32">Siap!</button>
            </div>
        </div>

    </Dialog>





    <!-- Dialog End -->

</template>
