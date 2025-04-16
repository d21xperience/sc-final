<!-- File: src/components/IssueDegreeForm.vue -->
<template>
    <button @click="handleSubmit" :disabled="isLoading">
        {{ isLoading ? 'Memproses...' : 'Verifikasi Ijazah' }}
    </button>
</template>

<script setup>
import { ref } from 'vue';
import { ethers } from 'ethers';
import { prepareDegreeIssue } from '@/utils/degreeHandler';
import { Button } from 'primevue';
// Props atau inject data jika diperlukan dari parent component atau pinia
// Misalnya, kita asumsikan data ini di-passing sebagai props:
const props = defineProps({
    degreeData: Object,
    sekolah: String,
    ipfsUrl: String,
    transcript: Object,
    contract: Object
});

const isLoading = ref(false);

const handleSubmit = async () => {
    isLoading.value = true;
    try {
        // 1. Siapkan data
        const { degreeHash, gasEstimate } = await prepareDegreeIssue(
            props.degreeData,
            props.sekolah,
            props.ipfsUrl,
            props.transcript
        );

        // 2. Tampilkan gas fee ke user
        const proceed = confirm(`Biaya gas: ${ethers.utils.formatUnits(gasEstimate, 'gwei')} Gwei. Lanjutkan?`);
        if (!proceed) return;

        // 3. Kirim transaksi
        const tx = await props.contract.issueDegree(
            degreeHash,
            props.sekolah,
            Math.floor(Date.now() / 1000),
            props.ipfsUrl,
            props.transcript.subjects,
            props.transcript.grades
        );

        // 4. Tunggu konfirmasi
        await tx.wait();

        // 5. Simpan bukti ke backend
        await saveToBackend(tx.hash);

        alert("Ijazah berhasil diverifikasi!");
    } catch (error) {
        console.error("Error:", error);
        alert(`Gagal: ${error.message}`);
    } finally {
        isLoading.value = false;
    }
};

const saveToBackend = async (txHash) => {
    await fetch('/api/degrees', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
            txHash,
            degreeData: props.degreeData,
            sekolah: props.sekolah,
            ipfsUrl: props.ipfsUrl
        })
    });
};
</script>