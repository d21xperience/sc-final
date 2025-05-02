<template>
    <div class="p2 my-2">
        <div>
            <DataTable stripedRows :value="siswaList">
                <Column field="ijazah.nama" header="Nama"></Column>
                <Column field="ipfsUrl" header="IPFS URL"></Column>
                <Column field="degreeHash" header="Degree Hash"></Column>
                <Column field="txHash" header="Trx Hash"></Column>
                <Column field="bcType" header="BC Type"></Column>
                <Column header="Created"></Column>
            </DataTable>
        </div>
        <Toast />
    </div>
</template>
<script setup>
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Toast from 'primevue/toast';

import { computed, onMounted, ref, watch } from 'vue';
import { useStore } from 'vuex'
import { useToast } from "primevue/usetoast";
const toast = useToast();

const store = useStore();
// Ambil semester
const tahunAjaranId = computed(() => store.getters["sekolahService/getSelectedTahunAjaran"].slice(0, 4))

onMounted(async () => {
    await fetchTransaksi()
})
const siswaList = ref()
const fetchTransaksi = async () => {
    try {
        let payload = {
            sekolah_id: await store.getters["sekolahService/getSekolah"]?.sekolah_id,
            tahun_ajaran_id: tahunAjaranId.value
        }
        const results = await store.dispatch("scService/fetchIjazahBC", payload)
        siswaList.value = results.degreeData
        console.log(results)
    } catch (error) {
        toast.add({ severity: 'error', summary: 'Error', detail: error, life: 3000 });
    }
}

watch(tahunAjaranId, async (newVal) => {
    await fetchTransaksi()
})


</script>