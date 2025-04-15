<script setup>
import { ref, computed, onMounted, watch, nextTick, toRaw } from 'vue';
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
import { result } from 'lodash';


const isEdit = ref(false)

const sekolah = computed(() => store.getters["sekolahService/getSekolah"])
const selectedTingkat = ref(null)
const tingkatPendidikanOptions = ref([])
const submitted = ref(false)
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
    // anggotaKelas: {
    //     anggotaRombelId: '',
    //     pesertaDidikId: '',
    //     rombonganBelajarId: '',
    //     semesterId: ''
    // }
});

// const anggotaKelas = ref({
//     anggotaRombelId: '',
//     pesertaDidikId: '',
//     rombonganBelajarId: '',
//     semesterId: '',
//     pesertaDidik: [],
//     rombonganBelajar: {}
// })
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
// ================================
// composable
// ================================
import { useSekolahService } from '@/composables/useSekolahService'
const selectedSemester = computed(() => store.getters["sekolahService/getSelectedSemester"])
const schemaname = computed(() => store.getters["sekolahService/getTabeltenant"]?.schemaname)

const { fetchKelas, kelasList, fetchGuruTerdaftar, guruTerdaftarList } = useSekolahService(schemaname, selectedSemester)
// ================================

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

watch(selectedJurusan, async (newVal) => {
    // console.log("tes Jursauan")
    if (typeof (newVal) === "object") {
        if (newVal) {
            await fetchKurikulum()
            const kur = (kurikulumList._rawValue)
            selectedKurikulum.value = kur.find(item => item.jurusanId === newVal.jurusanId)
        }
    }
})
// ===================================


const selectedGuruTerdaftar = ref(null)

const generateUUID = () => crypto.randomUUID();
// Handle Submit Form
// const submitForm = async () => {
//     submitted.value = true
//     if (!isEdit.value) {
//         if (rombel.value.nmKelas) {
//             tambahKelasBaru()
//         }

//     } else {
//         console.log("in here!!")
//         addAnggotaKelas()
//     }

// };
// const addAnggotaKelas = async () => {
//     const savedData = JSON.parse(localStorage.getItem("unsavedPesertaDidikBaru"));
//     console.log(savedData)
//     anggotaKelas.value.anggotaRombelId = generateUUID()
//     anggotaKelas.value.pesertaDidik = savedData

//     // return
//     if (savedData) {
//         try {
//             const payload = {
//                 schemaname: schemaname.value,
//                 siswa: savedData
//             }
//             const results = await store.dispatch("sekolahService/createBanyakSiswa", payload)
//             //console.log(results)
//             // localStorage.removeItem("unsavedPesertaDidikBaru");
//             return results
//         } catch (error) {
//             console.log(error)
//             return []
//         }
//     } else {
//         return false
//     }

// }
const angotaKelas = computed(() => JSON.parse(localStorage.getItem("unsavedPesertaDidik")) || null)
const simpanKelas = async () => {
    try {
        rombel.value.tingkatPendidikanId = selectedTingkat.value?.tingkatPendidikanId
        rombel.value.semesterId = await store.getters["sekolahService/getSelectedSemester"]?.semesterId
        rombel.value.jurusanId = selectedJurusan.value?.jurusanId
        rombel.value.ptkId = selectedGuruTerdaftar.value.ptk?.ptkId
        rombel.value.jenisRombel = Number(rombel.value.jenisRombel) || 1;
        rombel.value.namaJurusanSp = selectedJurusan.value?.namaJurusan
        rombel.value.kurikulumId = selectedKurikulum.value?.kurikulumId
        rombel.value.sekolahId = sekolah.value?.sekolah_id
        let result = null
        const payload = {
            schemaname: schemaname.value,
            kelas: rombel.value,
            // anggota_kelas: [anggotaKelas.value]
        }
        if (isEdit.value) {
            // Cek apakah anggota kelas diisi
            if (angotaKelas.value) {
                payload.anggota_kelas = angotaKelas.value
            }
            // anggotaKelas.value.anggotaRombelId = generateUUID()
            // anggotaKelas.value.pesertaDidik = savedData
            // console.log(payload)
            result = await store.dispatch("sekolahService/editKelas", payload)
            if (result) {
                toast.add({ severity: 'success', summary: 'Sukses', detail: 'Data berhasil diperbaharui', life: 3000 });
            }

        } else {
            result = await store.dispatch("sekolahService/createKelas", payload)
            if (result) {
                toast.add({ severity: 'success', summary: 'Sukses', detail: 'Data berhasil disimpan', life: 3000 });
                submitted.value = false
                Object.keys(rombel.value).forEach(key => {
                    rombel.value[key] = '';
                });
                selectedJurusan.value = null
                selectedKurikulum.value = null
                selectedGuruTerdaftar.value = {}
                selectedTingkat.value = ""
            }
        }
    } catch (error) {
        toast.add({ severity: 'error', summary: 'Gagal', detail: 'Data gagal disimpan', life: 3000 });

        console.error(error)
    }
}


const batal = async () => {
    await nextTick();
    router.push({ name: 'readKelas' })
}

onMounted(async () => {
    await fetchTingkat()
    await fetchJurusan()
    await fetchGuruTerdaftar()
    if (kelasId) {
        isEdit.value = true
        await fetchKelas(kelasId);
        if (kelasList.value && kelasList.value.length > 0) {
            const kelas = kelasList.value[0]
            rombel.value = {
                rombonganBelajarId: kelas.rombonganBelajarId || '',
                sekolahId: kelas.sekolahId || '',
                semesterId: kelas.semesterId || '',
                jurusanId: kelas.jurusanId || '',
                ptkId: kelas.ptkId || '',
                nmKelas: kelas.nmKelas || '',
                tingkatPendidikanId: kelas.tingkatPendidikanId || '',
                jenisRombel: kelas.jenisRombel || '',
                namaJurusanSp: kelas.namaJurusan || '',
                jurusan_sp_id: kelas.namaJurusanSpId || '',
                kurikulumId: kelas.kurikulumId || '',
                // tingkatPendidikan: null,
                // kurikulum: null,
                // jurusan: null,
            }
            selectedTingkat.value = tingkatPendidikanOptions.value.find(item => item.tingkatPendidikanId === kelas.tingkatPendidikanId)
            selectedJurusan.value = jurusanList.value.find(item => item.jurusanId === kelas.jurusanId)
            selectedGuruTerdaftar.value = guruTerdaftarList.value.find(item => item.ptk.ptkId === kelas.ptk.ptkId)
        }
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
                            placeholder="contoh: x tbsm a" :invalid="submitted && !rombel.nmKelas" />
                        <small v-if="submitted && !rombel.nmKelas" class="text-red-500">Nama harus
                            diisi.</small>
                    </div>
                    <div class="w-full md:w-40">
                        <label class="block text-gray-700">Tingkat</label>
                        <Select v-model="selectedTingkat" :options="tingkatPendidikanOptions"
                            placeholder="Pilih tingkat" optionLabel="nama" class="w-full"
                            :invalid="submitted && !selectedTingkat" />
                        <small v-if="submitted && !selectedTingkat" class="text-red-500">Tingkat harus
                            diisi.</small>
                    </div>
                </div>
                <div class="">
                    <label class="block text-gray-700">Wali kelas</label>
                    <div class="relative">
                        <Select v-model="selectedGuruTerdaftar" :options="guruTerdaftarList"
                            placeholder="Pilih Wali kelas" optionLabel="ptk.nama" class="w-full"
                            :invalid="submitted && !selectedGuruTerdaftar" />
                        <small v-if="submitted && !selectedGuruTerdaftar" class="text-red-500">Wali kelas harus
                            diisi.</small>
                    </div>
                </div>
            </div>

            <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
                <div>
                    <label class="block text-gray-700">Jurusan</label>
                    <div class="relative">
                        <AutoComplete v-model="selectedJurusan" :suggestions="filteredJurusan" optionLabel="namaJurusan"
                            @complete="searchJurusan" @keydown.space.prevent="handleKeydown"
                            placeholder="Cari Jurusan..." class="w-full" fluid dropdown
                            :invalid="submitted && !selectedJurusan" />
                        <small v-if="submitted && !selectedJurusan" class="text-red-500">Jurusan harus
                            diisi.</small>
                    </div>
                </div>
                <div>
                    <label class="block text-gray-700">Kurikulum</label>
                    <div class="relative">
                        <Select v-model="selectedKurikulum" :options="kurikulumList" placeholder="Pilih Kurikulum"
                            optionLabel="namaKurikulum" class="w-full" :invalid="submitted && !selectedKurikulum" />
                        <small v-if="submitted && !selectedKurikulum" class="text-red-500">Kurikulum harus
                            diisi.</small>
                    </div>
                </div>
            </div>
        </section>


        <Toast />
        <!-- Daftar anggota rombel -->
        <div v-show=isEdit>
            <h2 class="text-xl font-semibold mb-4">Anggota Kelas</h2>
            <!-- Anggota Kelas -->
            <AnggotaKelas :rombongan-belajar-id="kelasId" :is-edit="isEdit" />
            <!-- End of Anggota Kelas -->
        </div>

        <div class="flex justify-end space-x-4">
            <button class="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600"
                @click="simpanKelas">Simpan</button>
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
