<script setup>
import { ref, onMounted } from "vue";
import { loadWeb3, getBalance, getNetwork, disconnectMetaMask, listenForAccountChange } from "@/utils/web3";

// PRIMVEVUE
import Button from 'primevue/button';

// Variabel reaktif
const accounts = ref([]);
const selectedAccount = ref(null);
const balance = ref("0");
const networkId = ref(null);
const chainId = ref(null);
const errorMessage = ref(null);

// Fungsi untuk menghubungkan ke MetaMask
const connectMetaMask = async () => {
    errorMessage.value = null;
    const { accounts: accs, error } = await loadWeb3();

    if (error) {
        errorMessage.value = error;
        return;
    }

    if (accs.length > 0) {
        accounts.value = accs;
        selectedAccount.value = accs[0]; // Pilih akun pertama secara default

        // Simpan status koneksi di localStorage
        localStorage.setItem('isMetaMaskConnected', 'true');

        // Ambil saldo dan informasi jaringan
        updateAccountData(selectedAccount.value);
    }
};

// Fungsi untuk disconnect dari MetaMask
const disconnect = async () => {
    const { success, error } = await disconnectMetaMask();
    if (error) {
        errorMessage.value = error;
        return;
    }
    if (success) {
        accounts.value = [];
        selectedAccount.value = null;
        balance.value = "0";
        networkId.value = null;
        chainId.value = null;
        // Hapus status koneksi dari localStorage
        localStorage.removeItem('isMetaMaskConnected');
    }
};

// Periksa koneksi saat komponen dimuat
const checkConnection = async () => {
    if (localStorage.getItem('isMetaMaskConnected') === 'true') {
        await connectMetaMask();
    }
};

// Fungsi untuk menghubungkan ke MetaMask
// const connectMetaMask = async () => {
//     errorMessage.value = null;
//     const { accounts: accs, error } = await loadWeb3();

//     if (error) {
//         errorMessage.value = error;
//         return;
//     }

//     if (accs.length > 0) {
//         accounts.value = accs;
//         selectedAccount.value = accs[0]; // Pilih akun pertama secara default

//         // Ambil saldo dan informasi jaringan
//         updateAccountData(selectedAccount.value);
//     }
// };

// // Fungsi untuk disconnect dari MetaMask
// const disconnect = async () => {
//     const { success, error } = await disconnectMetaMask();
//     if (error) {
//         errorMessage.value = error;
//         return;
//     }
//     if (success) {
//         accounts.value = [];
//         selectedAccount.value = null;
//         balance.value = "0";
//         networkId.value = null;
//         chainId.value = null;
//     }
// };

// Fungsi untuk memperbarui saldo & jaringan saat akun dipilih
const updateAccountData = async (account) => {
    if (!account) return;

    const { balance: bal, error: balanceError } = await getBalance(account);
    if (balanceError) errorMessage.value = balanceError;
    else balance.value = bal;

    const { networkId: netId, chainId: cId, error: networkError } = await getNetwork();
    if (networkError) errorMessage.value = networkError;
    else {
        networkId.value = netId;
        chainId.value = cId;
    }
};

// Fungsi untuk menangani perubahan akun dari dropdown
const handleAccountChange = (event) => {
    selectedAccount.value = event.target.value;
    updateAccountData(selectedAccount.value);
};



// Event listener untuk perubahan akun atau jaringan
// onMounted(() => {
//     listenForAccountChange((accs) => {
//         if (accs.length === 0) {
//             disconnect();
//         } else {
//             accounts.value = accs;
//             selectedAccount.value = accs[0]; // Pilih akun pertama saat berubah
//             updateAccountData(selectedAccount.value);
//         }
//     });
// });

onMounted(async () => {
    // Periksa koneksi yang tersimpan
    await checkConnection();

    // Listener untuk perubahan akun
    listenForAccountChange((accs) => {
        if (accs.length === 0) {
            disconnect();
        } else {
            accounts.value = accs;
            selectedAccount.value = accs[0];
            updateAccountData(selectedAccount.value);
        }
    });

    // Listener untuk perubahan rantai
    if (window.ethereum) {
        window.ethereum.on('chainChanged', () => {
            window.location.reload(); // Reload halaman saat rantai berubah
        });
    }
});





// Mempersingkat address
const shortenText = (text) => {
    if (text.length <= 10) return text; // Tidak dipersingkat jika terlalu pendek
    return `${text.slice(0, 5)}...${text.slice(-5)}`;
};
// copy to clipboard
const textToCopy = ref(null); // Reference for the text element
const isCopied = ref(false); // State for success message visibility

const copyText = async () => {
    console.log("copied")
    try {
        if (textToCopy.value) {
            // Use clipboard API to copy text
            await navigator.clipboard.writeText(textToCopy.value.textContent);
            isCopied.value = true;

            // Hide the success message after 2 seconds
            setTimeout(() => {
                isCopied.value = false;
            }, 2000);
        }
    } catch (err) {
        console.error("Failed to copy text:", err);
    }
};


import { ethers } from 'ethers';

// async function issueDegree(contractAddress, degreeHash, sekolah, issueDate) {
//     // 1. Dapatkan provider dari MetaMask
//     if (!window.ethereum) {
//         throw new Error("MetaMask tidak terdeteksi!");
//     }

//     const provider = new ethers.providers.Web3Provider(window.ethereum);
//     const signer = provider.getSigner();

//     // // 2. Load ABI contract (pastikan ABI sudah didefinisikan)
//     // const abi = [...]; // Isi dengan ABI contract Anda

//     // // 3. Buat instance contract
//     // const contract = new ethers.Contract(contractAddress, abi, signer);

//     try {
//         // 4. Kirim transaksi
//         console.log("Mengirim transaksi...");
//         const tx = await contract.issueDegree(
//             ethers.utils.hexlify(degreeHash), // Convert [32]byte ke hex
//             sekolah,
//             issueDate
//         );

//         // 5. Tunggu konfirmasi transaksi
//         await tx.wait();
//         console.log("Transaksi berhasil:", tx.hash);
//         return tx.hash;
//     } catch (error) {
//         console.error("Error:", error);
//         throw new Error("Gagal mengirim transaksi: " + error.message);
//     }
// }

const txHash = ref(null);
const error = ref(null);

// async function handleIssueDegree() {
//     try {
//         // Data contoh (sesuaikan dengan kebutuhan)
//         const degreeHash = "0x123..."; // Hash ijazah dalam format hex
//         const sekolah = "Universitas Contoh";
//         const issueDate = Math.floor(Date.now() / 1000); // Timestamp Unix

//         txHash.value = await issueDegree(
//             "0xContractAddress", // Alamat contract
//             degreeHash,
//             sekolah,
//             issueDate
//         );
//     } catch (err) {
//         error.value = err.message;
//     }
// }

// // Setelah transaksi berhasil
// contract.on("DegreeIssued", (degreeHash, sekolah, issueDate, event) => {
//     console.log("Event DegreeIssued:", degreeHash, sekolah, issueDate);
//     alert(`Ijazah berhasil diterbitkan! Hash: ${degreeHash}`);
// });

// // Setelah tx berhasil
// async function saveToBackend(txHash, degreeData) {
//     const response = await fetch('/api/save-degree', {
//         method: 'POST',
//         headers: { 'Content-Type': 'application/json' },
//         body: JSON.stringify({ txHash, ...degreeData }),
//     });
//     return response.json();
// }

// // Contoh pemanggilan:
// await saveToBackend(tx.hash, {
//     degreeHash: "0x123...",
//     sekolah: "Universitas Contoh",
//     issueDate: 1234567890
// });
</script>

<template>
    <!-- <div>
        <h2>MetaMask Connection</h2> -->

    <Button @click="connectMetaMask" v-if="accounts.length === 0" label="Connect to MetaMask" icon="pi pi-user" />
    <Button @click="disconnect" v-if="accounts.length > 0" label="Disconnect" />


    <div class="mt-4" v-if="accounts.length > 0">
        <button @click="handleIssueDegree">Issue Ijazah</button>
        <p v-if="txHash">Transaksi berhasil! Hash: {{ txHash }}</p>
        <p v-if="error">{{ error }}</p>
    </div>

    <!--    <div v-if="accounts.length > 0">
            <label for="accountSelect">Select Account:</label>
            <select id="accountSelect" @change="handleAccountChange" v-model="selectedAccount">
                <option v-for="account in accounts" :key="account" :value="account">
                    {{ account }}
                </option>
            </select>
        </div>

        <p v-if="selectedAccount">Connected Account: {{ selectedAccount }}</p>
        <p v-if="selectedAccount">Balance: {{ balance }} ETH</p>
        <p v-if="selectedAccount">Network ID: {{ networkId }}</p>
        <p v-if="selectedAccount">Chain ID: {{ chainId }}</p>

        <p v-if="errorMessage" class="error">{{ errorMessage }}</p>
    </div> -->

    <!-- <div class="max-w-4xl mx-auto p-4"> -->
    <!-- Header -->
    <div class="flex items-center justify-between bg-white p-4 rounded-lg shadow">
        <div class="flex items-center space-x-2">
            <div class="bg-gray-200 p-2 rounded-full">
                <i class="fas fa-l text-gray-600">
                </i>
            </div>
            <span class="text-gray-700">
                Localhost 8545
            </span>
            <i class="fas fa-chevron-down text-gray-500">
            </i>
        </div>
        <div class="flex flex-col items-center">
            <div class="flex space-x-2 items-center">
                <div class="bg-yellow-400 p-2 rounded-full w-2 h-2">
                </div>
                <span class="text-gray-700">
                    Account 2
                </span>
                <i class="pi pi-chevron-down" style="color: green"></i>
            </div>

            <div>
                <span class="text-gray-500">
                    <!-- 0x14dC7...d9955 -->
                    <p v-if="selectedAccount">{{ shortenText(selectedAccount) }} <i
                            class="pi pi-copy text-gray-500 cursor-pointer" @click="copyText"></i></p>

                </span>

            </div>
        </div>
        <div class="flex item-center justify-center">
            <i class="pi pi-ellipsis-v" style="color: green"></i>
        </div>
    </div>
    <!-- Balance -->
    <div class="bg-white mt-4 p-6 rounded-lg shadow text-center">
        <div class="text-3xl font-bold text-gray-800">
            <!-- $191,154.00 USD -->
            <p v-if="selectedAccount">{{ balance }} ETH</p>
        </div>
        <div class="text-green-500">
            +$0 (0.00%)
            <!-- <a class="text-blue-500" href="#">
                Portfolio
            </a> -->
        </div>
        <div class="flex justify-center space-x-4 mt-4">
            <div class="flex flex-col items-center">
                <div class="bg-gray-200 p-3 rounded-full">
                    <i class="fas fa-exchange-alt text-blue-500">
                    </i>
                </div>
                <span class="text-gray-700 mt-2">
                    Buy &amp; Sell
                </span>
            </div>
            <div class="flex flex-col items-center">
                <div class="bg-gray-200 p-3 rounded-full w-10 h-10 flex items-center">
                    <i class="pi pi-paper-plane text-blue-500">
                    </i>
                </div>
                <span class="text-gray-700 mt-2">
                    Send
                </span>
            </div>
            <div class="flex flex-col items-center">
                <div class="bg-gray-200 p-3 rounded-full w-10 h-10 flex items-center">
                    <i class="pi pi-download text-blue-500">
                    </i>
                </div>
                <span class="text-gray-700 mt-2">
                    Receive
                </span>
            </div>
        </div>
    </div>
    <!-- Tabs -->
    <div class="bg-white mt-4 p-4 rounded-lg shadow">
        <!-- <div class="flex justify-around border-b">
            <div class="text-blue-500 pb-2 border-b-2 border-blue-500">
                Tokens
            </div>
            <div class="text-gray-500 pb-2">
                NFTs
            </div>
            <div class="text-gray-500 pb-2">
                Activity
            </div>
        </div> -->
        <!-- <div class="mt-4">
            <div class="flex items-center justify-between">
                <div class="flex items-center space-x-2">
                    <i class="fas fa-chevron-down text-gray-500">
                    </i>
                    <span class="text-gray-700">
                        Localhost 8545
                    </span>
                </div>
                <i class="fas fa-ellipsis-v text-gray-500">
                </i>
            </div>
            <div class="flex items-center justify-between mt-4">
                <div class="flex items-center space-x-2">
                    <img alt="Ethereum logo" class="w-8 h-8" height="32"
                        src="https://storage.googleapis.com/a1aa/image/J_56IYPQE9Z_LFn5vJVlTL2dDSEHutJ639u_uhEc0XY.jpg"
                        width="32" />
                    <span class="text-gray-700">
                        Ethereum
                    </span>
                </div>
                <span class="text-gray-700">
                    100 ETH
                </span>
            </div>
        </div> -->
    </div>
</template>

<style>
.error {
    color: red;
    font-weight: bold;
}
</style>
