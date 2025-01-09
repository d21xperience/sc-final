<template>
    <div class="container mx-auto p-8">
        <div class="flex justify-between items-center mb-8">
            <h1 class="text-2xl font-bold">Data Sekolah</h1>
            <!-- <div class="bg-blue-800 text-white px-4 py-2 rounded">
                Tanggal : 26-11-2024
            </div> -->
        </div>
        <div>
            <div v-if="dataConnected">
                <div class=" p-6 rounded-lg shadow-lg">
                    <h2 class="text-xl font-bold mb-4"><i class="fas fa-school"></i> Detail Data Sekolah</h2>
                    <div class="mb-4">
                        <h3 class="text-lg font-semibold mb-2">IDENTITAS SEKOLAH</h3>
                        <div class="grid grid-cols-2 gap-4">
                            <div>Nama</div>
                            <div><input type="text" v-model="sekolah.nama" class="w-full"></div>
                            <div>Jenjang</div>
                            <div>SMK</div>
                            <div>NSS</div>
                            <div>402021015029</div>
                            <div>NPSN</div>
                            <div>20254180</div>
                        </div>
                    </div>
                    <div class="mb-4">
                        <h3 class="text-lg font-semibold mb-2">ALAMAT</h3>
                        <div class="grid grid-cols-2 gap-4">
                            <div>Jalan</div>
                            <div>Jalan Kolonel Ahmad Syam Jatinangor</div>
                            <div>Desa/Kelurahan</div>
                            <div>Cikeruh</div>
                            <div>Kecamatan</div>
                            <div>Kec. Jatinangor</div>
                        </div>
                    </div>
                    <div class="mb-4">
                        <h3 class="text-lg font-semibold mb-2">KONTAK</h3>
                        <div class="grid grid-cols-2 gap-4">
                            <div>Telp./Fax.</div>
                            <div>0227782242 / 0227782242</div>
                            <div>email</div>
                            <div>smkspasundanjatinangor@gmail.com</div>
                            <div>website</div>
                            <div>http://smkpasundanjatinangor.sch.id</div>
                        </div>
                    </div>
                    <div class="mb-4">
                        <h3 class="text-lg font-semibold mb-2">Kepala Sekolah</h3>
                        <div class="grid grid-cols-2 gap-4">
                            <div>Nama Kepsek</div>
                            <div>Nunung Kurnia, S.T., M.Pd.</div>
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
            <div v-else>
                <EmptyData />
            </div>
        </div>
    </div>

</template>

<script setup>
import EmptyData from '@/components/EmptyData.vue';
import { ref } from 'vue';

const sekolah = ref({
    sekolah_id: '8a5bd957-66bc-4096-9ff1-fee096b87ba0',
    nama: 'SMK PASUNDAN JATINANGOR',
    nss: '402021015029',
    npsn: '20254180',
    bentuk_pendidikan_id: '',
    alamat_jalan: '',
    rt: '',
    rw: '',
    nama_dusun: '',
    desa_kelurahan: '',
    kode_wilayah: '',
    kode_pos: ''

})

const dataConnected = ref(false)
</script>

<style lang="scss" scoped></style>