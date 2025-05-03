<template>
    <div>
        <h2 class="text-xl font-semibold mb-4">Informasi Ijazah</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
            <div>
                <label class="block text-gray-700" for="prov">Prov.</label>
                <InputText v-model="ijazahParam.prov" fluid name="prov" id="prov" placeholder="Masukan nama" />
            </div>
            <div>
                <label class="block text-gray-700" for="kab">Kab</label>
                <InputText v-model="ijazahParam.kab" fluid name="kab" id="kab" placeholder="Masukan nama" />
            </div>
            <div>
                <label class="block text-gray-700" for="tpt-ijazah">Tempat Ijazah</label>
                <InputText v-model="tempatIjazah" fluid name="tpt-ijazah" id="tpt-ijazah" placeholder="Masukan nama" />
            </div>
            <div>
                <label class="block text-gray-700" for="tgl-dikelurkan-ijazah">Tgl dikeluarkan
                    Ijazah</label>
                <input type="date" placeholder="YYYY-MM-DD" class=" w-full p-2 border border-gray-300 rounded"
                    v-model="tglIjazah">
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue';

import InputText from 'primevue/inputtext';

const ijazahParam = ref({
    tanggalIjazah: "",
    tempatIjazah: "Sumedang",
    kab: "Sumedang",
    prov: "Jawa Barat"

})

const tglIjazah = ref()
const tempatIjazah = ref()
</script>