<script setup>
import { ref, computed, onMounted, watch, toRaw } from 'vue';
import { useRoute } from "vue-router";
const route = useRoute();

const kelasId = route.query.kelasId;
import InputText from 'primevue/inputtext';
import Select from 'primevue/select';
import AutoComplete from 'primevue/autocomplete';
import Toast from 'primevue/toast';
import { useToast } from 'primevue/usetoast';
const toast = useToast();


import router from '@/router';
import { useStore } from "vuex";
const store = useStore();
import AnggotaKelas from '@/components/AnggotaKelas.vue';

const schemaName = ref('')
const selectedSemester = computed(() => {
    return store.getters["sekolahService/getSelectedSemester"]
})
let isEdit = false

// Opsi Dropdown
const sekolah = computed(() => store.getters["sekolahService/getSekolah"])
const selectedTingkat = ref(null)
const tingkatPendidikanOptions = ref([])
// Model Kelas
const rombel = ref({
    rombonganBelajarId: '',
    sekolahId: '',
    semesterId: '',
    jurusanId: '',
    ptkId: '',
    nmKelas: '',
    tingkatPendidikanId: '',
    jenisRombel: '',
    namaJurusanSp: '',
    jurusan_sp_id: '',
    kurikulumId: '',
    tingkatPendidikan: null,
    kurikulum: null,
    jurusan: null,
    anggotaKelas: {
        anggotaRombelId: '',
        pesertaDidikId: '',
        rombonganBelajarId: '',
        semesterId: ''
    }
});


const fetchTingkat = async () => {
    const payload = {
        jenjang_pendidikan_id: sekolah.value?.jenjangPendidikanId
    }
    const response = await store.dispatch("sekolahService/fetchTingkatPendidikan", payload)
    tingkatPendidikanOptions.value = response
}

// ===================================
// KURIKULUM
// ===================================
const selectedKurikulum = ref(null)
const kurikulumList = ref([])
const fetchKurikulum = async () => {
    try {
        const payload = {
            jurusan_id: selectedJurusan.value?.jurusanId
        }
        const response = await store.dispatch("sekolahService/fetchKurikulum", payload)
        kurikulumList.value = response
    } catch (error) {
        console.error(error)
    }
}
// ===================================

// ===================================
// JURUSAN
// ===================================
const selectedJurusan = ref(null)
const jurusanList = ref([])
const filteredJurusan = ref([])
const fetchJurusan = async () => {
    const payload = {}
    if (sekolah.value?.jenjangPendidikanId == 6) {
        switch (sekolah.value?.bentukPendidikanId) {
            case 15:
                payload.param = "untuk_smk"
                break;
            case 13:
                payload.param = "untuk_sma"
                break;

            default:
                break;
        }
    }
    payload.jenjang_pendidikan_id = sekolah.value?.jenjangPendidikanId
    const response = await store.dispatch("sekolahService/fetchJurusan", payload)
    jurusanList.value = response
}
const searchJurusan = (event) => {
    setTimeout(() => {
        let query = event.query.toLowerCase();

        filteredJurusan.value = jurusanList.value.filter((jurusan) =>
            jurusan.namaJurusan.toLowerCase().includes(query)
        );
    }, 250);
}
const handleKeydown = (event) => {
    if (event.key === " ") {
        selectedJurusan.value += " "; // Menambahkan spasi ke query
    }
};

watch(selectedJurusan, (newVal) => {
    if (typeof (newVal) === "object") {
        fetchKurikulum()
    }
})
// ===================================


const selectedGuru = ref(null)
const guruList = ref([])
// const filteredGuru = ref([])
const fetchGuru = async () => {
    try {
        const cek = await store.getters["sekolahService/getPTKTerdaftar"]
        // console.log(cek)
        if (cek == null) {
            let payload = {
                tahunAjaranId: selectedSemester.value?.tahunAjaranId,
                schemaname: schemaName.value
            }
            const response = await store.dispatch("sekolahService/fetchPTKTerdaftar", payload)
            // console.log(response)
            guruList.value = response
        } else {
            guruList.value = cek
        }

    } catch (error) {
        console.error(error)
    }

}




// ===============================
// const emit = defineEmits(["add"]);

const fetchKelas = async () => {
    try {
        const payload = {
            schema_name: schemaName.value,
            semester_id: selectedSemester.value?.semesterId,
            kelas_id: kelasId
        }
        const response = await store.dispatch("sekolahService/fetchRombel", payload);
        rombel.value = { ...response[0] };
        selectedGuru.value = guruList.value.find((item) => item.ptk.ptkId === rombel.value.ptkId)
        // selectedJurusan.value = jurusanList.value.find((item) => item.jurusanId === rombel.value.jurusanId)
        selectedJurusan.value = rombel.value.jurusan
        selectedTingkat.value = rombel.value.tingkatPendidikan
        selectedKurikulum.value = rombel.value.kurikulum
    } catch (error) {
        console.error("Gagal mengambil data kelas:", error);
    }
};
// const rombelAnggota = ref()
const generateUUID = () => crypto.randomUUID();
// Handle Submit Form
const submitForm = async () => {
    // Cek apakah kelas sedang dibuat
    // mengisi data input
    console.log("submitForm execute!")
    if (rombel.value.nmKelas != "") {
        addKelas()
        let anggota = addSiswaBaru()
        // simpanKelas(rombel.value, anggota)
    }

    // // Cek apakah siswa
    // const cek = await addSiswaBaru()
    // // console.log(cek)
    // if (cek) {
    //     // fungsi untuk menambahkan siswa ke anggota kelas dimana siswa sudah ada di database sebelumnya.    
    // } else {
    //     await addAnggotaKelas()
    //     toast.add({ severity: 'warn', summary: 'Gagal', detail: 'Terjadi error pada saat penambahan siswa baru', life: 3000 });
    // }
};
const addKelas = async () => {
    // console.log("addKelas execute!")
    rombel.value.ptkId = selectedGuru.value.ptk?.ptkId
    rombel.value.semesterId = await store.getters["sekolahService/getSelectedSemester"]?.semesterId
    rombel.value.tingkatPendidikanId = selectedTingkat.value?.tingkatPendidikanId
    rombel.value.kurikulumId = selectedKurikulum.value?.kurikulumId
    rombel.value.jurusanId = selectedJurusan.value?.jurusanId
    rombel.value.namaJurusanSp = selectedJurusan.value?.namaJurusan
    rombel.value.sekolahId = await store.getters["sekolahService/getSekolah"]?.sekolahId
    rombel.value.jenisRombel = Number(rombel.value.jenisRombel) || 1;
    rombel.value.sekolahId = sekolah.value.sekolahId
    rombel.value.ptk = {}
}
const addSiswaBaru = async () => {
    const savedData = JSON.parse(localStorage.getItem("unsavedPesertaDidikBaru"));
    // console.log(savedData)
    if (savedData) {
        try {
            const payload = {
                schemaname: schemaName.value,
                siswa: savedData
            }
            const results = await store.dispatch("sekolahService/createBanyakSiswa", payload)
            //console.log(results)
            localStorage.removeItem("unsavedPesertaDidikBaru");
            return results
        } catch (error) {
            console.log(error)
            return []
        }
    } else {
        return false
    }
}

const simpanKelas = async (kelas, anggotaKelas) => {
    try {
        const payload = {
            schema_name: schemaName.value,
            kelas: kelas,
            anggota_kelas: anggotaKelas
        }

        let results = null
        if (isEdit) {
            payload.kelas_id = kelasId
            results = await store.dispatch("sekolahService/editKelas", payload)
        } else {
            results = await store.dispatch("sekolahService/createKelas", payload)
            Object.keys(rombel.value).forEach(key => {
                rombel.value[key] = '';
            });
            selectedJurusan.value = {}
            selectedKurikulum.value = {}
            selectedGuru.value = {}
            selectedTingkat.value = ""
        }
        if (results != null) {
            toast.add({ severity: 'success', summary: 'Sukses', detail: 'Data berhasil disimpan', life: 3000 });
            localStorage.removeItem("unsavedPesertaDidik");

            // router.push({ name: 'readKelas' })
        } else {
            toast.add({ severity: 'error', summary: 'Gagal', detail: 'Data gagal disimpan', life: 3000 });
        }
    } catch (error) {
        console.error(error)
    }
}


const batal = () => {
    router.push({ name: 'readKelas' })
}

onMounted(() => {
    schemaName.value = store.getters["sekolahService/getTabeltenant"]?.schemaName
    fetchTingkat()
    fetchJurusan()
    fetchGuru()
    if (kelasId) {
        // console.log("onMounted")
        isEdit = true
        fetchKelas();
    }

})

</script>

<template>
    <div class="container bg-white p-8 rounded-lg shadow-md">
        <h1 class="text-2xl font-bold mb-6">{{ isEdit ? 'Form Edit Kelas' : 'Form Tambah Kelas' }}</h1>

        <section class="mb-6">
            <h2 class="text-xl font-semibold mb-4">Informasi Kelas</h2>
            <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
                <div class="flex space-x-1">
                    <div class="w-full">
                        <label class="block text-gray-700" for="nmKelas">Nama Kelas</label>
                        <InputText v-model="rombel.nmKelas" fluid name="nmKelas" id="nmKelas"
                            placeholder="contoh: x tbsm a" />
                    </div>
                    <div class="w-full md:w-40">
                        <label class="block text-gray-700">Tingkat</label>
                        <Select v-model="selectedTingkat" :options="tingkatPendidikanOptions"
                            placeholder="Pilih tingkat" optionLabel="nama" class="w-full" />
                    </div>
                </div>
                <div class="">
                    <label class="block text-gray-700">Wali kelas</label>
                    <div class="relative">
                        <Select v-model="selectedGuru" :options="guruList" placeholder="Pilih Wali kelas"
                            optionLabel="ptk.nama" class="w-full" />
                    </div>
                </div>
            </div>

            <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
                <div>
                    <label class="block text-gray-700">Jurusan</label>
                    <div class="relative">
                        <AutoComplete v-model="selectedJurusan" :suggestions="filteredJurusan" optionLabel="namaJurusan"
                            @complete="searchJurusan" @keydown.space.prevent="handleKeydown"
                            placeholder="Cari Jurusan..." class="w-full" fluid dropdown />
                    </div>
                </div>
                <div>
                    <label class="block text-gray-700">Kurikulum</label>
                    <div class="relative">
                        <Select v-model="selectedKurikulum" :options="kurikulumList" placeholder="Pilih Kurikulum"
                            optionLabel="namaKurikulum" class="w-full" />
                    </div>
                </div>
            </div>
        </section>


        <Toast />
        <!-- Daftar anggota rombel -->
        <div>
            <h2 class="text-xl font-semibold mb-4">Anggota Kelas</h2>
            <!-- Anggota Kelas -->
            <AnggotaKelas :rombongan-belajar-id="kelasId" :is-edit="isEdit" />
            <!-- End of Anggota Kelas -->
        </div>

        <div class="flex justify-end space-x-4">
            <button class="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600"
                @click="submitForm">Simpan</button>
            <button class="bg-gray-300 text-gray-700 px-4 py-2 rounded hover:bg-gray-400" @click="batal">Batal</button>
        </div>
    </div>
</template>

<style scoped>
label {
    font-weight: bold;
    display: block;
    margin-bottom: 0.5rem;
}

select option:disabled {
    color: gray;
    font-weight: bold;
}
</style>
