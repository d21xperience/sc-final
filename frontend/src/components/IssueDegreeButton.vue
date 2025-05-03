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


const handleSubmit = async () => {
    isLoading.value = true;
    try {
        // console.log("props.transcript:", props.transcript);
        // console.log("subjects:", props.transcript?.subjects);
        // console.log("grades:", props.transcript?.grades);

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
        saveToBackend(tx.hash, degreeHash)
    } catch (err) {
        console.error(err);
        alert(`Gagal memproses: ${err.message}`);
    } finally {
        isLoading.value = false;
    }
};
const handleBatchSubmit = async () => {
    isLoading.value = true;
    try {
        if (!contract.value) await loadContract();

        for (const student of studentsData) {
            // Buat degreeHash dari data (hash dari ipfs + sekolah + waktu sekarang)
            const degreeHash = generateDegreeHash(student); // pastikan fungsi ini konsisten

            const issueDate = Math.floor(Date.now() / 1000);

            // Validasi input
            if (
                !Array.isArray(student.subjects) ||
                !Array.isArray(student.grades) ||
                student.subjects.length !== student.grades.length
            ) {
                console.warn("Data tidak valid:", student);
                continue; // Skip data tidak valid
            }

            console.log("Mengirim untuk:", student.sekolah, degreeHash);

            const tx = await contract.value.issueDegree(
                degreeHash,
                student.sekolah,
                issueDate,
                student.ipfsUrl,
                student.subjects,
                student.grades
            );

            await tx.wait(); // Tunggu satu per satu
            console.log("Sukses simpan:", degreeHash);
        }

        alert("Semua data siswa berhasil dikirim ke smart contract.");
    } catch (err) {
        console.error("Gagal:", err);
        alert(`Terjadi error: ${err.message}`);
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
    try {
        console.log(props.degreeData)
        let degreData = {
            degree_hash: degreeHash,
            tx_hash: txHash,
            ipfs_url: props.ipfsUrl,
            bc_type: "Ethereum",
            link_bc_explorer: "http://localhost:26000",
            tahun_ajaran_id: props.degreeData?.tahun_lulus,
            ijazah: {
                nama: props.degreeData?.nama,
                nisn: props.degreeData?.nisn,
                program_keahlian: props.degreeData?.major
            },
            schemaname: await store.getters["sekolahService/getTabeltenant"]?.schemaname
        }

        const payload = {
            degree_data: [degreData]
        }
        // console.log(payload)
        // return
        const response = await store.dispatch("scService/createIjazahBC", payload)
        console.log(response)
    } catch (error) {
        console.log(error)
    }

};
</script>