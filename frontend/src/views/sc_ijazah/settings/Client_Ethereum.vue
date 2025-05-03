<template>
    <div>
        <h2 class="text-xl font-semibold mb-4">Informasi Ethereum</h2>
        <div class="grid grid-cols-1 md:grid-cols-2 gap-4 mb-4">
            <div>
                <label class="block text-gray-700" for="contract-address">Contract Address</label>
                <InputText v-model="ethereumParam.contractAddress" fluid name="contract-address" id="contract-address"
                    placeholder="Masukan alamat smart contract" />
            </div>
            <div>
                <label class="block text-gray-700" for="block-explorer-url">Block Explorer URL</label>
                <InputText v-model="ethereumParam.blockExplorerURL" fluid name="block-explorer-url"
                    id="block-explorer-url" placeholder="Masukan nama" />
            </div>
            <div>
                <label class="block text-gray-700" for="ethereumNetwork">Network</label>
                <InputText v-model="ethereumParam.network" fluid name="ethereumNetwork" id="ethereumNetwork"
                    placeholder="Masukan nama" />
            </div>
            <!-- <div>
                <label class="block text-gray-700" for="tpt-ijazah"></label>
                <InputText v-model="tempatIjazah" fluid name="tpt-ijazah" id="tpt-ijazah" placeholder="Masukan nama" />
            </div> -->
            <!-- <div>
                <label class="block text-gray-700" for="tgl-dikelurkan-ijazah">Tgl dikeluarkan
                    Ijazah</label>
                <input type="date" placeholder="YYYY-MM-DD" class=" w-full p-2 border border-gray-300 rounded"
                    v-model="tglIjazah">
            </div> -->
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue';

import InputText from 'primevue/inputtext';

const ethereumParam = ref({
    blockExplorerURL: "http://localhost:26000",
    contractAddress: "0xdc64a140aa3e981100a9beca4e685f962f0cf6c9",
    network: "localhost:8545"

})

</script>