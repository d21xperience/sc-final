<template>
    <div class="container mx-auto p-8">
        <div class="flex justify-between items-center mb-8">
            <h1 class="text-2xl font-bold">Data Sekolah</h1>
            <!-- <div class="bg-blue-800 text-white px-4 py-2 rounded">
                Tanggal : 26-11-2024
            </div> -->
        </div>
        <div>
            <div class=" p-6 rounded-lg shadow-lg">
                <h2 class="text-xl font-bold mb-4"><i class="fas fa-school"></i> Detail Data Sekolah</h2>
                <div class="mb-4">
                    <h3 class="text-lg font-semibold mb-2">IDENTITAS SEKOLAH</h3>
                    <div class="grid grid-cols-2 gap-4">
                        <div>Nama</div>
                        <div><input type="text" v-model="sekolah.nama" class="w-full"></div>
                        <div>Jenjang</div>
                        <div></div>
                        <div>NSS</div>
                        <div>sekolah.nss</div>
                        <div>NPSN</div>
                        <div>{{sekolah?.npsn}}</div>
                    </div>
                </div>
                <div class="mb-4">
                    <h3 class="text-lg font-semibold mb-2">ALAMAT</h3>
                    <div class="grid grid-cols-2 gap-4">
                        <div>Jalan</div>
                        <div>{{ sekolah?.alamat }}</div>
                        <div>Desa/Kelurahan</div>
                        <div></div>
                        <div>Kecamatan</div>
                        <div>{{ sekolah?.kecamatan }}</div>
                    </div>
                </div>
                <div class="mb-4">
                    <h3 class="text-lg font-semibold mb-2">KONTAK</h3>
                    <div class="grid grid-cols-2 gap-4">
                        <div>Telp./Fax.</div>
                        <div></div>
                        <div>email</div>
                        <div></div>
                        <div>website</div>
                        <div></div>
                    </div>
                </div>
                <div class="mb-4">
                    <h3 class="text-lg font-semibold mb-2">Kepala Sekolah</h3>
                    <div class="grid grid-cols-2 gap-4">
                        <div>Nama Kepsek</div>
                        <div></div>
                        <div>NIP Kepsek</div>
                        <div></div>
                    </div>
                </div>
                <div class="flex justify-end">
                    <button class="bg-blue-800 text-white px-4 py-2 rounded flex items-center">
                        <i class="fas fa-edit mr-2"></i> Update Data Kepsek
                    </button>
                </div>
            </div>

        </div>
    </div>



</template>

<script setup>
import { ref, onMounted, computed } from "vue";
import { useStore } from "vuex";
const store = useStore();
// ambil data dari backend

const fetchSekolah = async () => {
    // get tabel tenant
    // ambil sekolah id dari user yang sedang login
    const sekolahID = await store.state.authService.user?.sekolahId;
    const tTenant = await store.dispatch("sekolahService/fetchTabeltenant", sekolahID)
    // console.log(tTenant)
    const dataSekolah = await store.dispatch("sekolahService/fetchSekolah", { schemaName: tTenant.schemaName, namaSekolah: tTenant.namaSekolah })
    // console.log(dataSekolah)
    sekolah.value = dataSekolah
}

onMounted(() => {
    fetchSekolah()
})
const sekolah = ref({})

</script>

<style lang="scss" scoped></style>