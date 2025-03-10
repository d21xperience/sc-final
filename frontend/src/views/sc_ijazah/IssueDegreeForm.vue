<template>
    <div class="mt-4"></div>
    <div class="max-w-2xl mx-auto p-6 bg-white shadow-md rounded-lg">
        <h2 class="text-2xl font-bold mb-6 text-center">Issue Degree</h2>
        <form @submit.prevent="submitForm" class="space-y-4">
            <!-- Degree Hash -->
            <div>
                <label for="degreeHash" class="block text-sm font-medium text-gray-700">Degree Hash:</label>
                <input type="text" id="degreeHash" v-model="formData.degreeHash" required
                    class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500" />
            </div>

            <!-- Sekolah -->
            <div>
                <label for="sekolah" class="block text-sm font-medium text-gray-700">Sekolah:</label>
                <input type="text" id="sekolah" v-model="formData.sekolah" required
                    class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500" />
            </div>

            <!-- Issue Date -->
            <div>
                <label for="issueDate" class="block text-sm font-medium text-gray-700">Issue Date (Unix
                    Timestamp):</label>
                <input type="number" id="issueDate" v-model="formData.issueDate" required
                    class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500" />
            </div>

            <!-- IPFS URL -->
            <div>
                <label for="ipfsUrl" class="block text-sm font-medium text-gray-700">IPFS URL:</label>
                <input type="text" id="ipfsUrl" v-model="formData.ipfsUrl" required
                    class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500" />
            </div>

            <!-- Mata Pelajaran -->
            <div>
                <label for="mataPelajaran" class="block text-sm font-medium text-gray-700">Mata Pelajaran
                    (comma-separated):</label>
                <input type="text" id="mataPelajaran" v-model="formData.mataPelajaran" required
                    class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500" />
            </div>

            <!-- Nilai -->
            <div>
                <label for="nilai" class="block text-sm font-medium text-gray-700">Nilai (comma-separated):</label>
                <input type="text" id="nilai" v-model="formData.nilai" required
                    class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500" />
            </div>

            <!-- Submit Button -->
            <div>
                <button type="submit"
                    class="w-full px-4 py-2 bg-indigo-600 text-white font-semibold rounded-md hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500">
                    Submit
                </button>
            </div>
        </form>

        <!-- Transaction Hash -->
        <div v-if="transactionHash" class="mt-6 p-4 bg-green-50 border border-green-200 rounded-md">
            <p class="text-sm text-green-700">Transaction Hash: <span class="font-mono">{{ transactionHash }}</span></p>
        </div>

        <!-- Error Message -->
        <div v-if="error" class="mt-6 p-4 bg-red-50 border border-red-200 rounded-md">
            <p class="text-sm text-red-700">{{ error }}</p>
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue';
import { ethers } from 'ethers';

// Data form
const formData = ref({
    degreeHash: '',
    sekolah: '',
    issueDate: '',
    ipfsUrl: '',
    mataPelajaran: '',
    nilai: '',
});

const transactionHash = ref('');
const error = ref('');

// ABI dan alamat kontrak
const contractABI = [
    [
        {
            "inputs": [],
            "stateMutability": "nonpayable",
            "type": "constructor"
        },
        {
            "anonymous": false,
            "inputs": [
                {
                    "indexed": true,
                    "internalType": "bytes32",
                    "name": "degreeHash",
                    "type": "bytes32"
                }
            ],
            "name": "DegreeDeleted",
            "type": "event"
        },
        {
            "anonymous": false,
            "inputs": [
                {
                    "indexed": true,
                    "internalType": "bytes32",
                    "name": "degreeHash",
                    "type": "bytes32"
                },
                {
                    "indexed": false,
                    "internalType": "bytes32",
                    "name": "sekolah",
                    "type": "bytes32"
                },
                {
                    "indexed": false,
                    "internalType": "uint256",
                    "name": "issueDate",
                    "type": "uint256"
                },
                {
                    "indexed": false,
                    "internalType": "bytes32",
                    "name": "ipfsUrl",
                    "type": "bytes32"
                }
            ],
            "name": "DegreeIssued",
            "type": "event"
        },
        {
            "inputs": [
                {
                    "internalType": "bytes32",
                    "name": "",
                    "type": "bytes32"
                }
            ],
            "name": "degrees",
            "outputs": [
                {
                    "internalType": "bytes32",
                    "name": "degreeHash",
                    "type": "bytes32"
                },
                {
                    "internalType": "bytes32",
                    "name": "sekolah",
                    "type": "bytes32"
                },
                {
                    "internalType": "uint256",
                    "name": "issueDate",
                    "type": "uint256"
                },
                {
                    "internalType": "bytes32",
                    "name": "ipfsUrl",
                    "type": "bytes32"
                },
                {
                    "internalType": "bool",
                    "name": "exists",
                    "type": "bool"
                }
            ],
            "stateMutability": "view",
            "type": "function"
        },
        {
            "inputs": [
                {
                    "internalType": "bytes32",
                    "name": "_degreeHash",
                    "type": "bytes32"
                }
            ],
            "name": "deleteDegree",
            "outputs": [],
            "stateMutability": "nonpayable",
            "type": "function"
        },
        {
            "inputs": [
                {
                    "internalType": "bytes32",
                    "name": "_degreeHash",
                    "type": "bytes32"
                }
            ],
            "name": "getDegree",
            "outputs": [
                {
                    "internalType": "bytes32",
                    "name": "",
                    "type": "bytes32"
                },
                {
                    "internalType": "uint256",
                    "name": "",
                    "type": "uint256"
                },
                {
                    "internalType": "bytes32",
                    "name": "",
                    "type": "bytes32"
                },
                {
                    "internalType": "bytes32[]",
                    "name": "",
                    "type": "bytes32[]"
                },
                {
                    "internalType": "uint8[]",
                    "name": "",
                    "type": "uint8[]"
                }
            ],
            "stateMutability": "view",
            "type": "function"
        },
        {
            "inputs": [
                {
                    "internalType": "bytes32",
                    "name": "_degreeHash",
                    "type": "bytes32"
                },
                {
                    "internalType": "bytes32",
                    "name": "_sekolah",
                    "type": "bytes32"
                },
                {
                    "internalType": "uint256",
                    "name": "_issueDate",
                    "type": "uint256"
                },
                {
                    "internalType": "bytes32",
                    "name": "_ipfsUrl",
                    "type": "bytes32"
                },
                {
                    "internalType": "bytes32[]",
                    "name": "_mataPelajaran",
                    "type": "bytes32[]"
                },
                {
                    "internalType": "uint8[]",
                    "name": "_nilai",
                    "type": "uint8[]"
                }
            ],
            "name": "issueDegree",
            "outputs": [],
            "stateMutability": "nonpayable",
            "type": "function"
        },
        {
            "inputs": [],
            "name": "owner",
            "outputs": [
                {
                    "internalType": "address",
                    "name": "",
                    "type": "address"
                }
            ],
            "stateMutability": "view",
            "type": "function"
        },
        {
            "inputs": [
                {
                    "internalType": "bytes32",
                    "name": "",
                    "type": "bytes32"
                },
                {
                    "internalType": "uint256",
                    "name": "",
                    "type": "uint256"
                }
            ],
            "name": "transcripts",
            "outputs": [
                {
                    "internalType": "bytes32",
                    "name": "name",
                    "type": "bytes32"
                },
                {
                    "internalType": "uint8",
                    "name": "grade",
                    "type": "uint8"
                }
            ],
            "stateMutability": "view",
            "type": "function"
        }
    ]
];
const contractAddress = '0xYourContractAddress'; // Ganti dengan alamat kontrak Anda

// Fungsi untuk mengirim transaksi
const submitForm = async () => {
    try {
        // Validasi input
        if (!formData.value.degreeHash || !formData.value.sekolah || !formData.value.issueDate || !formData.value.ipfsUrl || !formData.value.mataPelajaran || !formData.value.nilai) {
            throw new Error('Semua field harus diisi.');
        }

        // Konversi input ke format yang sesuai
        const mataPelajaranArray = formData.value.mataPelajaran.split(',').map(item => ethers.utils.formatBytes32String(item.trim()));
        const nilaiArray = formData.value.nilai.split(',').map(item => parseInt(item.trim(), 10));

        // Hubungkan ke provider Ethereum (MetaMask)
        const provider = new ethers.providers.Web3Provider(window.ethereum);
        await provider.send('eth_requestAccounts', []); // Meminta akses ke akun MetaMask
        const signer = provider.getSigner();

        // Buat instance kontrak
        const contract = new ethers.Contract(contractAddress, contractABI, signer);

        // Kirim transaksi
        const tx = await contract.issueDegree(
            ethers.utils.formatBytes32String(formData.value.degreeHash),
            ethers.utils.formatBytes32String(formData.value.sekolah),
            parseInt(formData.value.issueDate, 10),
            ethers.utils.formatBytes32String(formData.value.ipfsUrl),
            mataPelajaranArray,
            nilaiArray
        );

        // Tunggu transaksi selesai
        await tx.wait();
        transactionHash.value = tx.hash;
        error.value = '';
    } catch (err) {
        error.value = `Error: ${err.message}`;
        console.error(err);
    }
};
</script>