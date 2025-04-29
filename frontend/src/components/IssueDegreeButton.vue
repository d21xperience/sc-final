<template>
    <button @click="handleSubmit" :disabled="isLoading">
        {{ isLoading ? 'Memproses...' : 'Verifikasi Ijazah' }}
    </button>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { ethers } from 'ethers';
import DegreeContractABI from '@/VerifikasiIjazahABI.json';
import { keccak256, toUtf8Bytes } from 'ethers';
import store from '@/store';
const contractAddress = '0xdc64a140aa3e981100a9beca4e685f962f0cf6c9';
const props = defineProps({
    degreeData: Object,  // { name, nisn, graduationYear, major }
    sekolah: String,
    ipfsUrl: String,
    transcript: Object,  // { subjects: ["B. Indonesia", "Matematika"], grades: [85, 90] }
});

// State
const isLoading = ref(false);
const contract = ref(null);

// Inisialisasi kontrak
const loadContract = async () => {
    if (window.ethereum) {
        const provider = new ethers.BrowserProvider(window.ethereum);
        const signer = await provider.getSigner();
        // const contractAddress = '0xdc64a140aa3e981100a9beca4e685f962f0cf6c9'; // ganti dengan address kontrak kamu
        contract.value = new ethers.Contract(contractAddress, DegreeContractABI, signer);
    } else {
        alert('Metamask tidak ditemukan.');
    }
};

// Fungsi untuk membuat hash ijazah (gunakan string JSON untuk konsistensi)
const generateDegreeHash = (data) => {
    const stringified = JSON.stringify(data);
    return keccak256(toUtf8Bytes(stringified)); // ethers v6
};

// Fungsi submit
// const handleSubmit = async () => {
//     isLoading.value = true;
//     try {
//         if (!contract.value) await loadContract();

//         // 1. Buat hash ijazah
//         const degreeHash = generateDegreeHash(props.degreeData);
//         const issueDate = Math.floor(Date.now() / 1000); // timestamp
//         // Validasi dan konversi array transcript
//         let subjects = props.transcript.subjects;
//         let grades = props.transcript.grades;

//         // Jika bentuknya string, parsing dulu
//         if (typeof subjects === 'string') {
//             subjects = JSON.parse(subjects);
//         }
//         if (typeof grades === 'string') {
//             grades = JSON.parse(grades);
//         }

//         // Pastikan grades berbentuk array number
//         grades = grades.map(n => parseInt(n));

//         // Validasi panjang array sama
//         if (subjects.length !== grades.length) {
//             throw new Error("Jumlah mata pelajaran dan nilai tidak sama.");
//         }

//         // 2. Estimasi gas (opsional, bisa dilewatkan jika tidak diperlukan)
//         // const gasEstimate = await contract.value.issueDegree.estimateGas(
//         //     degreeHash,
//         //     props.sekolah,
//         //     issueDate,
//         //     props.ipfsUrl,
//         //     props.transcript.subjects,
//         //     props.transcript.grades
//         // );
//         const gasEstimate = await contract.value.issueDegree.estimateGas(
//             degreeHash,
//             props.sekolah,
//             issueDate,
//             props.ipfsUrl,
//             subjects,
//             grades
//         );





//         const proceed = confirm(`Biaya gas kira-kira: ${ethers.formatUnits(gasEstimate, 'gwei')} Gwei. Lanjutkan?`);
//         if (!proceed) return;

//         // 3. Kirim transaksi
//         // const tx = await contract.value.issueDegree(
//         //     degreeHash,
//         //     props.sekolah,
//         //     issueDate,
//         //     props.ipfsUrl,
//         //     props.transcript.subjects,
//         //     props.transcript.grades
//         // );
//         const tx = await contract.value.issueDegree(
//             degreeHash,
//             props.sekolah,
//             issueDate,
//             props.ipfsUrl,
//             subjects,
//             grades
//         );

//         await tx.wait(); // tunggu konfirmasi

//         alert("Ijazah berhasil diverifikasi di blockchain!");
//         // await saveToBackend(tx.hash, degreeHash);

//     } catch (err) {
//         console.error(err);
//         alert(`Gagal memproses: ${err.message}`);
//     } finally {
//         isLoading.value = false;
//     }
// };

const handleSubmit = async () => {
    isLoading.value = true;
    try {
        console.log("props.transcript:", props.transcript);
        console.log("subjects:", props.transcript?.subjects);
        console.log("grades:", props.transcript?.grades);

        if (!contract.value) await loadContract();

        const degreeHash = generateDegreeHash(props.degreeData);
        const issueDate = Math.floor(Date.now() / 1000);

        // Validasi & parsing subjects dan grades
        let subjects = props.transcript?.subjects;
        let grades = props.transcript?.grades;

        if (!Array.isArray(subjects) || !Array.isArray(grades)) {
            throw new Error("Data transcript tidak valid. Pastikan subjects dan grades berupa array.");
        }

        grades = grades.map((n) => parseInt(n));

        if (subjects.length !== grades.length) {
            throw new Error("Jumlah mata pelajaran dan nilai tidak cocok.");
        }

        const gasEstimate = await contract.value.issueDegree.estimateGas(
            degreeHash,
            props.sekolah,
            issueDate,
            props.ipfsUrl,
            subjects,
            grades
        );

        const proceed = confirm(`Biaya gas kira-kira: ${ethers.formatUnits(gasEstimate, 'gwei')} Gwei. Lanjutkan?`);
        if (!proceed) return;

        const tx = await contract.value.issueDegree(
            degreeHash,
            props.sekolah,
            issueDate,
            props.ipfsUrl,
            subjects,
            grades
        );

        await tx.wait();
        alert("Ijazah berhasil diverifikasi di blockchain!");
    } catch (err) {
        console.error(err);
        alert(`Gagal memproses: ${err.message}`);
    } finally {
        isLoading.value = false;
    }
};

// const contract = ref(null);

onMounted(async () => {
    try {
        if (window.ethereum) {
            await window.ethereum.request({ method: 'eth_requestAccounts' });
            const provider = new ethers.BrowserProvider(window.ethereum);
            const signer = await provider.getSigner();

            // const contractAddress = '0x700b6A60ce7EaaEA56F065753d8dcB9653dbAD35'; // HARUS STRING BUKAN undefined/objek
            contract.value = new ethers.Contract(contractAddress, DegreeContractABI, signer);
        } else {
            alert('Metamask tidak ditemukan. Harap instal terlebih dahulu.');
        }
    } catch (err) {
        console.error('Gagal menginisialisasi kontrak:', err);
    }
});
// 🗄️ Simpan ke backend
const saveToBackend = async (txHash, degreeHash) => {
    // await fetch('/api/degrees', {
    //     method: 'POST',
    //     headers: { 'Content-Type': 'application/json' },
    //     body: JSON.stringify({
    //         txHash,
    //         degreeHash,
    //         degreeData: props.degreeData,
    //         sekolah: props.sekolah,
    //         ipfsUrl: props.ipfsUrl
    //     })
    // });
    try {
        const payload = {
            schemaname: await store.getters["sekolahService/getSchemaname"]?.schemaname

        }
    } catch (error) {
        console.log(error)
    }

};
</script>