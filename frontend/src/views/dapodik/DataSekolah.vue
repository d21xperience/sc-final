<template>
    <div class="container mx-auto p-8">
        <div class="flex justify-between items-center mb-8">
            <h1 class="text-2xl font-bold">Data Sekolah</h1>
            <!-- <div class="bg-blue-800 text-white px-4 py-2 rounded">
                Tanggal : 26-11-2024
            </div> -->
            <div>
                <Button icon="pi pi-pencil" @click="editSekolah" class="mr-2" :style="isEdit?'background-color:blue;border:none':'background-color:gray;border:none'"
                    v-tooltip.bottom="'Edit data'" />
            </div>
        </div>
        <div v-if="!isEdit">
            <div class=" p-6 rounded-lg shadow-lg">
                <h2 class="text-xl font-bold mb-4"><i class="fas fa-school"></i> Detail Data Sekolah</h2>
                <div class="mb-4">
                    <h3 class="text-lg font-semibold mb-2">IDENTITAS SEKOLAH</h3>
                    <div class="grid grid-cols-2 gap-4">
                        <div>Nama</div>
                        <div>{{ sekolah.value?.nama }}</div>
                        <div>Jenjang</div>
                        <div>{{ jenjangPendidikanSelected?.nama }}</div>
                        <div>Bentuk Pendidikan</div>
                        <div>{{ bentukPendidikanSelected?.nama }}</div>
                        <div>NSS</div>
                        <div>{{ sekolah.value?.nss }}</div>
                        <div>NPSN</div>
                        <div>{{ sekolah.value?.npsn }}</div>
                    </div>
                </div>
                <div class="mb-4">
                    <h3 class="text-lg font-semibold mb-2">ALAMAT</h3>
                    <div class="grid grid-cols-2 gap-4">
                        <div>Jalan</div>
                        <div>{{ sekolah.value?.alamat }}</div>
                        <div>Desa/Kelurahan</div>
                        <div>{{ sekolah.value?.kelurahan }}</div>
                        <div>Kecamatan</div>
                        <div>{{ sekolah.value?.kecamatan }}</div>
                        <div>Provinsi</div>
                        <div>{{ sekolah.value?.propinsi }}</div>
                    </div>
                </div>
                <div class="mb-4">
                    <h3 class="text-lg font-semibold mb-2">KONTAK</h3>
                    <div class="grid grid-cols-2 gap-4">
                        <div>Telp./Fax.</div>
                        <div>{{ sekolah.value?.telepon }}</div>
                        <div>email</div>
                        <div>{{ sekolah.value?.email }}</div>
                        <div>website</div>
                        <div> <a v-if="sekolah.value?.website" :href="getWebsiteUrl(sekolah.value?.website)"
                                target="_blank" rel="noopener noreferrer" class="text-blue-700 underline">
                                {{ sekolah.value?.website }}
                            </a></div>
                    </div>
                </div>
                <div class="mb-4">
                    <h3 class="text-lg font-semibold mb-2">Kepala Sekolah</h3>
                    <div class="grid grid-cols-2 gap-4">
                        <div>Nama Kepsek</div>
                        <div>{{ sekolah.value?.nmKepsek }}</div>
                        <div>NIP Kepsek</div>
                        <div>{{ sekolah.value?.nipKepsek }}</div>
                    </div>
                </div>
            </div>

        </div>
        <div v-else>
            <div class=" p-6 rounded-lg shadow-lg">
                <h2 class="text-xl font-bold mb-4"><i class="fas fa-school"></i> Detail Data Sekolah</h2>
                <div class="mb-4">
                    <h3 class="text-lg font-semibold mb-2">IDENTITAS SEKOLAH</h3>
                    <div class="grid grid-cols-2 gap-4">
                        <div>Nama</div>
                        <div>{{ sekolah.value?.nama }}</div>
                        <div>Jenjang</div>
                        <div> <Select v-model="jenjangPendidikanSelected" filter :options="jenjangPendidikanFiltered"
                                optionLabel="nama" placeholder="Pilih jenjang pendidikan" fluid showClear
                                class="w-full md:w-56" /></div>
                        <div>Bentuk Pendidikan</div>
                        <div><Select v-model="bentukPendidikanSelected" :options="bentukPendidikan" filter
                                optionLabel="nama" placeholder="Pilih bentuk pendidikan" fluid showClear
                                class="w-full md:w-56" /></div>
                        <div>NSS</div>
                        <div>{{ sekolah.value?.nss }}</div>
                        <div>NPSN</div>
                        <div>{{ sekolah.value?.npsn }}</div>
                    </div>
                </div>
                <div class="mb-4">
                    <h3 class="text-lg font-semibold mb-2">ALAMAT</h3>
                    <div class="grid grid-cols-2 gap-4">
                        <div>Jalan</div>
                        <div>{{ sekolah.value?.alamat }}</div>
                        <div>Desa/Kelurahan</div>
                        <div>{{ sekolah.value?.kelurahan }}</div>
                        <div>Kecamatan</div>
                        <div>{{ sekolah.value?.kecamatan }}</div>
                    </div>
                </div>
                <div class="mb-4">
                    <h3 class="text-lg font-semibold mb-2">KONTAK</h3>
                    <div class="grid grid-cols-2 gap-4">
                        <div>Telp./Fax.</div>
                        <div>{{ sekolah.nama }}</div>
                        <div>email</div>
                        <div>{{ sekolah.value?.email }}</div>
                        <div>website</div>
                        <div>{{ sekolah.value?.website }}</div>
                    </div>
                </div>
                <div class="mb-4">
                    <h3 class="text-lg font-semibold mb-2">Kepala Sekolah</h3>
                    <div class="grid grid-cols-2 gap-4">
                        <div>Nama Kepsek</div>
                        <div><input v-model="sekolah.value.nmKepsek" class="w-full" name="nmKepsek"></div>
                        <div>NIP Kepsek</div>
                        <div><input v-model="sekolah.value.nipKepsek" class="w-full" name="nipKepsek"></div>
                    </div>
                </div>
                <div class="flex justify-end">
                    <button class="bg-blue-800 text-white px-4 py-2 rounded flex items-center" @click="updateData">
                        <i class="fas fa-edit mr-2"></i> Update Data
                    </button>
                </div>
            </div>
        </div>
        <Toast />
    </div>
</template>

<script setup>
import { ref, onMounted, reactive, watch, computed } from "vue";
import { useStore } from "vuex";
const store = useStore();
// ambil data dari backend
import Button from 'primevue/button';
import Select from 'primevue/select';
import Toast from 'primevue/toast';
import { useToast } from 'primevue/usetoast';
const toast = useToast();
const sekolah = reactive({})
const dataSekolah = ref()
const fetchSekolah = async () => {
    // get tabel tenant
    // ambil sekolah id dari user yang sedang login
    const sekolahID = await store.state.authService.user?.sekolahId;
    const tTenant = await store.dispatch("sekolahService/fetchTabeltenant", sekolahID)
    // console.log(tTenant)
    dataSekolah.value = await store.dispatch("sekolahService/fetchSekolah", { schemaname: tTenant.schemaname, namaSekolah: tTenant.namaSekolah })
    // console.log(dataSekolah)
    sekolah.value = dataSekolah
}

const fetchRefTable = async () => {
    bentukPendidikan.value = await store.dispatch("sekolahService/fetchBentukPendidikan")
    jenjangPendidikan.value = await store.dispatch("sekolahService/fetchJenjangPendidikan")
}

const jenjangPendidikanFiltered = computed(() => {
    return jenjangPendidikan.value.filter(
        (item) => item.jenjangLembaga === 1
    );
});

onMounted(async () => {
    await fetchSekolah()
    await fetchRefTable()
    jenjangPendidikanSelected.value = jenjangPendidikan.value.find(item => item.jenjangPendidikanId == sekolah.value.jenjangPendidikanId)
    bentukPendidikanSelected.value = bentukPendidikan.value.find(item => item.bentukPendidikanId == sekolah.value.bentukPendidikanId)
})



// Edit
const isEdit = ref(false)
const editSekolah = () => {
    isEdit.value = !isEdit.value
}

const bentukPendidikan = ref([])
const jenjangPendidikan = ref([])
const bentukPendidikanSelected = ref()
const jenjangPendidikanSelected = ref()

const changedData = ref({})
watch(sekolah, (newVal) => {
    changedData.value = {}; // Reset perubahan
    for (const key in newVal) {
        if (newVal[key] !== dataSekolah.value[key]) {
            changedData.value[key] = newVal[key]; // Simpan hanya data yang berubah
        }
    }
    // console.log("Perubahan terdeteksi:", changedData.value);
},
    { deep: true } // Perlu `deep: true` untuk melacak perubahan dalam objek
)

watch(jenjangPendidikanSelected, (newVal) => sekolah.value.jenjangPendidikanId = newVal.jenjangPendidikanId)
watch(bentukPendidikanSelected, (newVal) => sekolah.value.bentukPendidikanId = newVal.bentukPendidikanId)
const getWebsiteUrl = (url) => {
    if (!url.startsWith("http://") && !url.startsWith("https://")) {
        return `https://${url}`; // Tambahkan https jika belum ada
    }
    return url;
};

const updateData = async () => {
    if (Object.keys(changedData.value).length > 0) {
        console.log("Mengirim data ke backend:", changedData.value);
        // Kirim `changedData.value` ke backend via API
        const payload = {
            schemaname: store.getters["sekolahService/getTabeltenant"]?.schemaName,
            sekolah: sekolah.value
        }
        const resp = await store.dispatch("sekolahService/updateSekolah", payload)
        toast.add({ severity: 'info', summary: 'Info', detail: resp?.message, life: 3000 });
        // console.log(resp)
    } else {
        console.log("Tidak ada perubahan, tidak perlu update.");
    }
    isEdit.value = false
};
</script>

<style scoped>
    edit-class{
        background-color: red;

    }
</style>