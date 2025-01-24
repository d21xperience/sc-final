<script setup>
// Dialog


// import Menu from 'primevue/menu';

import Dialog from 'primevue/dialog';
import { useDialog } from 'primevue/usedialog';
const dialog = useDialog();

import Button from 'primevue/button';
import { computed, ref, watch, onMounted } from 'vue';

// Ambil dari database
const BCPlatform = ref()
const dialogSelectNetwork = ref(false)

// Akun
const akunBC = ref(null)
const dialogListAccounts = ref(false)
const dialogCreateAccount = ref(false)
const dialogCreateNewAccount = ref(false)
const dialogImportAccount = ref(false)

import Tes from "./Tes.vue"
import router from '@/router';
import QrCodeAccount from './QrCodeAccount.vue';
import axios from 'axios';


function openDialog() {
    dialog.open(Tes, {
        props: {
            header: 'Product List',
            style: {
                width: '50vw',
            },
            breakpoints: {
                '960px': '75vw',
                '640px': '90vw'
            },
            modal: true
        }
    })
}

function func_openMyQrCodeAccount() {
    dialog.open(QrCodeAccount, {
        props: {
            // header: 'Product List',
            style: {
                width: '30vw',
            },
            breakpoints: {
                '960px': '75vw',
                '640px': '90vw'
            },
            modal: true,
            // maximizable: true
        }
    })
}

// -----------
// const tes = () => {
//     alert('hello world!')
// }

// const addNewAccount = () => {
//     alert('new account')
// }

// const coba = (e) => {
//     alert(console.log(e))
// }


const sendKrypto = () => {
    router.push({ name: "sendKrypto" })
}
const menu = ref();
const items = ref([
    {
        label: 'Options',
        items: [
            {
                label: 'Refresh',
                icon: 'pi pi-refresh'
            },
            {
                label: 'Export',
                icon: 'pi pi-upload'
            },
            {
                label: 'Pengaturan',
                icon: 'pi pi-upload'
            }
        ]
    }
]);
const toggle = (event) => {
    menu.value.toggle(event);
};





const addBCNetwork = () => {
    router.push({ name: 'listBCNetwork' })
    dialogSelectNetwork.value = false
}

const networkActive = ref(false)
const networkName = ref(null)
const networkIndex = ref(0)
const selectNetwork = (e) => {
    // console.log(e.target.value)
    networkIndex.value = e.target.value
    // console.log(networkIndex.value)
    const text = e.target.textContent.trim();
    // console.log(text);
    networkName.value = text
    networkActive.value = true
    dialogSelectNetwork.value = false
    setActiveCurrentBC(BCPlatform.value[networkIndex.value].network_id)

    // ambil data akun dari backend
    // set aktive blockhain network
    // getBcPlatform()
}
onMounted(() => {
    getAccount()
    fetchBlockchainNetworks();
});

const api = axios.create({
    baseURL: 'http://localhost:8081', // URL utama API
    timeout: 5000, // Waktu maksimal request (ms)
    headers: {
        'Content-Type': 'application/json', // Header default
    },
});
// Mengambil bcplatform dari backend
const fetchBlockchainNetworks = async () => {
    // menggunakan axios
    try {
        const result = await api.get('/api/v1/accounts')
            .then(response => {
                // console.log(response.data)
                BCPlatform.value = response.data
                BCPlatform.value.forEach((bc, i) => {
                    // console.log(bc.applicable)
                    if (bc.activate) {
                        if (bc.applicable) {
                            console.log(i)
                            networkActive.value = true
                            networkIndex.value = i
                        }
                    }
                })
            })
    } catch (error) {
        console.error('Error connecting to blockchain network:', error);

    }
}

// Set aktive current Blockchain network
const setActiveCurrentBC = async (id) => {
    // console.log(id)
    try {
        const resp = await api.put(`/api/v1/blockchain-networks/${id}`, {
            Applicable: true
        })
        console.log(resp)
    } catch (error) {
        console.log('error', error)
    }
}
// Set diaktive current Blockchain network
const networkDiactive = async () => {
    networkActive.value = false
    let id = BCPlatform.value[networkIndex.value].network_id
    // console.log(id)
    try {
        const resp = await api.put(`/api/v1/blockchain-networks/${id}`, {
            Applicable: false
        })
        console.log(resp)
    } catch (error) {
        console.log('error', error)
    }
}

const currentAccount = ref(0)
// Buat akun
const getAccount = async (id) => {
    try {
        const resp = await api.get(`/api/v1/accounts`)
        akunBC.value = resp
    } catch (error) {
        console.log('error', error)
    }
}

const createAccount = async (id) => {
    try {
        const resp = await api.post(`/api/v1/account`, {
            netword_id: id,
            name: 'test',
            organization: '',
            isActive: false,
        })

        console.log(resp)
    } catch (error) {
        console.log("error", error)
    }
}




// -------------------------------------------------------------------------------


// Mempersingkat address
const shortenText = (text) => {
    if (text.length <= 10) return text; // Tidak dipersingkat jika terlalu pendek
    return `0x${text.slice(0, 5)}...${text.slice(-5)}`;
};

// Fungsi untuk mengonversi ETH ke Fiat (USD)
const fiatAmount = ref(null); // Menyimpan hasil konversi
async function convertToFiatCurrency(amount) {
    try {
        const response = await fetch(
            'https://api.coingecko.com/api/v3/simple/price?ids=ethereum&vs_currencies=usd'
        );
        const data = await response.json();
        const ethPrice = data.ethereum.usd;
        return (amount * ethPrice).toFixed(2);
    } catch (error) {
        console.error('Error fetching Ethereum price:', error);
        return null;
    }
}

// Watcher untuk memperbarui fiatAmount saat networkIndex berubah
// watch(networkIndex, async () => {
//     const amount = BCPlatform[networkIndex.value].account[0].amount;
//     fiatAmount.value = await convertToFiatCurrency(amount);
// });

// Inisialisasi nilai awal
// (async () => {
//   const amount = BCPlatform[0].account[0].amount;
//   fiatAm

// copy to clipboard
const textToCopy = ref(null); // Reference for the text element
const isCopied = ref(false); // State for success message visibility

const copyText = async () => {
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


</script>


<template>
    <div class="w-full bg-white shadow-md p-2">
        <div class="flex justify-between p-2">
            <div class="">
                <button @click="dialogSelectNetwork = true"
                    class="rounded-full bg-slate-300 py-2 px-4 hover:opacity-80">
                    <span v-if="!networkActive">Pilih Jaringan</span>
                    <span v-else>{{ BCPlatform[networkIndex].network_name }}</span>
                    <i class="ml-2 pi pi-angle-up "></i></button>
            </div>
            <div v-show="networkActive" class="flex items-center">
                <button class="bg-red-400 py-2 px-3 rounded-full hover:opacity-70 flex items-center gap-2"
                    @click="networkDiactive"><i class="pi pi-times"></i> Disconect</button>
            </div>
        </div>
    </div>
    <div class="p-4">
        <template v-if="networkActive">
            <h3 class="text-center font-bold text-3xl">{{ BCPlatform[networkIndex].network_name }} </h3>
            <div>
                <template v-if="BCPlatform[networkIndex].accounts.length > 1">
                    <div class="text-center">
                        <button type="button" @click="openDialog" class="hover:opacity-70 border-b-2 shadow-sm">
                            <!--   <span>{{ BCPlatform[networkIndex].accounts[0].name }}</span> <i class="pi pi-angle-up "></i>-->
                        </button>
                        <div class=" text-center">
                            <select name="bc-account" id="bc-account" v-model="currentAccount" class="outline-none">
                                <option v-for="(bc, index) in BCPlatform[networkIndex].accounts" :key="index"
                                    :value="index" class="text-3xl text-slate-600">
                                    {{ bc.name }}
                                </option>
                            </select>
                            <div class="text-sm text-gray-500  flex items-center justify-center my-2">
                                <p>{{ shortenText(BCPlatform[networkIndex].accounts[currentAccount].address) }}</p>
                                <button @click="copyText"><i class="pi pi-copy "></i></button>
                                <div class="relative">
                                    <p class="hidden" ref="textToCopy"> {{
                                        BCPlatform[networkIndex].accounts[currentAccount].address }}</p>
                                    <p v-if="isCopied"
                                        class="font-bold text-green-500 w-32 mt-2 absolute top-0 transition ease-in-out duration-500">
                                        Text copied!</p>
                                </div>
                            </div>
                        </div>
                        <h3 class="text-3xl font-bold mt-2">
                            <p> {{ BCPlatform[networkIndex].accounts[0].amount }} ETH</p>
                        </h3>
                        <p class="text-gray-500">

                            <!-- konversikan ke mata uang -->
                            <!-- $ {{ convertToFiatCurrency(BCPlatform[networkIndex].account[0].amount) }} USD -->
                            <!-- $ {{ fiatAmount ?? 'Loading...' }} USD -->
                        </p>
                    </div>
                    <div class="flex justify-center space-x-4 mt-4 pb-2 border-b-2">
                        <div class="flex flex-col items-center">
                            <Button icon="pi pi-plus" severity="info" aria-label="User" rounded />
                            <p class="text-sm mt-1">
                                Jual &amp;...
                            </p>
                        </div>
                        <div class="flex flex-col items-center">
                            <Button icon="pi pi-arrow-right" severity="info" aria-label="Send" rounded
                                @click="sendKrypto" />
                            <p class="text-sm mt-1">
                                Kirim
                            </p>
                        </div>
                        <div class="flex flex-col items-center">
                            <Button icon="pi pi-exchange" severity="info" aria-label="pertukaran" rounded />
                            <p class="text-xs mt-1">
                                Smartcontract
                            </p>
                        </div>
                        <!-- <div class="flex flex-col items-center">
                            <Button icon="pi pi-random" severity="info" aria-label="User" rounded />
                            <p class="text-sm mt-1">
                                Bridge
                            </p>
                        </div> -->
                        <div class="flex flex-col items-center">
                            <Button icon="pi pi-briefcase" severity="info" aria-label="User" rounded />
                            <p class="text-sm mt-1">
                                Portofolio
                            </p>
                        </div>
                        <!-- <div v-show="networkIndex <= 1" class="flex flex-col items-center">
                            <Button icon="pi pi-qrcode" severity="info" aria-label="User" rounded
                                @click="func_openMyQrCodeAccount" />
                            <p class="text-sm mt-1">
                                Terima
                            </p>
                        </div> -->

                    </div>
                    <!-- <div class="flex justify-center mt-6">
                        <div class="flex space-x-8">
                            <a class="text-blue-500 border-b-2 border-blue-500 pb-1" href="#">
                                Token
                            </a>
                            <a class="text-gray-500" href="#">
                                NFT
                            </a>
                            <a class="text-gray-500" href="#">
                                Aktivitas
                            </a>
                        </div>
                    </div> -->
                    <div class="mt-4">
                        <RouterView></RouterView>
                        <!-- <div class="flex items-center justify-between py-2 border-b">
                    <div class="flex items-center space-x-2">
                        <img alt="BNB logo" class="w-6 h-6" height="24"
                            src="https://storage.googleapis.com/a1aa/image/LUcCNgXFGubHChpik2eW4Mc5yj1CDFfeEHNfe7AQt4Wx8SneE.jpg"
                            width="24" />
                        <div>
                            <p class="font-medium">
                                BNB
                            </p>
                            <p class="text-sm text-gray-500">
                                BNB
                            </p>
                        </div>
                    </div>
                    <div class="text-right">
                        <p class="font-medium">
                            0 BNB
                        </p>
                        <p class="text-sm text-gray-500">
                            $0.00 USD
                        </p>
                    </div>
                </div>
                <div class="flex items-center justify-between py-2 border-b">
                    <div class="flex items-center space-x-2">
                        <font-awesome-icon :icon="faEthereum" />
                        <div>
                            <p class="font-medium">
                                USDT
                            </p>
                            <p class="text-sm text-gray-500">
                                Tether USD
                            </p>
                        </div>
                    </div>
                    <div class="text-right">
                        <p class="font-medium">
                            $4.76 USD
                        </p>
                        <p class="text-sm text-gray-500">
                            4.78 USDT
                        </p>
                    </div>
                </div>
                <div class="flex items-center justify-between py-2">
                    <div class="flex items-center space-x-2">
                        <img alt="JMPT logo" class="w-6 h-6" height="24"
                            src="https://storage.googleapis.com/a1aa/image/wl0kcOq8C9bCFxGqmEAaMHBDutEqt5nebLUmpeeTsfqkeSneE.jpg"
                            width="24" />
                        <div>
                            <p class="font-medium">
                                JMPT
                            </p>
                            <p class="text-sm text-gray-500">
                                Tether USD
                            </p>
                        </div>
                    </div>
                    <div class="text-right">
                        <p class="font-medium">
                            $4.74 USD
                        </p>
                        <p class="text-sm text-gray-500">
                            4.74 USD
                        </p>
                    </div>
                </div> -->
                    </div>
                </template>
                <template v-else>
                    <h2 class="text-center text-2xl">Anda belum mempunyai akun, silakan buat akun terlebih dahulu
                    </h2>
                    <div class="flex justify-center mt-6">
                        <Button @click="dialogCreateAccount = true" icon="pi pi-user" aria-label="addAccount"
                            label="Buat Akun" rounded style="background-color:orange;border: none; color:black" />
                    </div>
                </template>
            </div>
        </template>
        <template v-else>
            <div class="text-center">
                <h2 class="text-3xl uppercase">Anda Belum terhubung ke jaringan Blockahain</h2>
                <p class="my-3">Silahkankan <span class="font-semibold">pilih jaringan</span> yang ingin digunakan</p>
                <p class="mt-5"><i class="pi pi-sort-alt-slash" style="font-size: 3rem;"></i></p>
                <p class="my-6">atau lihat <span class="font-bold">daftar jaringan untuk </span>menambahkan</p>
                <Button @click="router.push({ name: 'listBCNetwork' })" icon="pi pi-cloud" label="Daftar Jaringan"
                    severity="warn" />
            </div>
        </template>
        <!-- </div> -->
    </div>

    <!-- <Dialog v-model:visible="dialogListAccounts" modal header="Pilih Akun" :style="{ width: '25rem' }" position="top">
       
    </Dialog> -->
    <Dialog v-model:visible="dialogCreateAccount" modal header="Tambahkan Akun" :style="{ width: '25rem' }"
        position="top">
        <div class="flex flex-col">
            <ul>
                <li class="mb-2 hover:text-blue-500 py-2 rounded-sm ">
                    <button @click="dialogCreateNewAccount = true">
                        <i class="pi pi-plus mr-2">
                        </i>
                        Tambah akun
                    </button>
                </li>
                <li @click="dialogImportAccount = true" class="mb-2 hover:text-blue-500 py-2 rounded-sm"><button>
                        <i class="pi pi-cloud-download mr-2">
                        </i>
                        Impor akun
                    </button></li>
            </ul>
        </div>
    </Dialog>



    <!-- Dialog Create New Acount -->
    <Dialog v-model:visible="dialogCreateNewAccount" modal header="Tambahkan Akun" :style="{ width: '25rem' }"
        position="top">
        <form action="" method="post">
            <div class="flex flex-col">
                <div class="mb-2  p-2 rounded-lg cursor-pointer">
                    <label for="">Nama akun</label>
                    <!-- Placeholder diisi dengan nama akun yang sedang aktif jika tidak aktif tambahkan nama akun secara default -->
                    <input v-model="akunBC" type="text" placeholder="Account - 1" class="w-full p-2">
                </div>

                <div class="flex justify-between gap-2 mt-2">
                    <Button label="Batal" icon="pi pi-times" class="w-full hover:opacity-75"
                        style="background-color: white; color: black;" @click="dialogCreateNewAccount = false" />
                    <Button label="Buat" icon="pi pi-check" class="w-full" @click="createAccount" />
                </div>
            </div>
        </form>
    </Dialog>
    <!-- ########################################### -->
    <!-- Dialog Import Akun -->
    <Dialog v-model:visible="dialogImportAccount" modal header="Impor Akun" :style="{ width: '25rem' }" position="top">
        <div class="flex flex-col">
            <p>Akun yang diimpor tidak akan ditautkan dengan Frasa Pemulihan Rahasia Akun ini.</p>
            <p>Pelajari selengkapnya</p>
            <div class="flex justify-between mt-2">
                <label for="jenis-kunci">Pilih jenis kunci</label>
                <select name="jenis-kunci" id="jenis-kunci" class="p-2">
                    <option value="0">Kebijakan pribadi</option>
                    <option value="0">File JSON</option>
                </select>
            </div>
            <div class="flex flex-col mb-4">
                <label for="string-kunci">Tempel string kunci</label>
                <div>
                    <input name="string-kunci" id="string-kunci" class="p-2 w-full border border-solid">
                    </input>
                </div>
            </div>

            <div class="flex justify-between gap-2 mt-2">
                <Button label="Batal" icon="pi pi-times" class="w-full hover:opacity-75"
                    style="background-color: white; color: black;" @click="dialogImportAccount = false" />
                <Button label="Impor" icon="pi pi-check" class="w-full" />
            </div>
        </div>
    </Dialog>

    <!-- ########################################### -->
    <!-- Dialog Select Network -->
    <Dialog v-model:visible="dialogSelectNetwork" modal header="Pilih Jaringan" position="top">
        <ul v-show="BCPlatform !== undefined">
            <li class="text-sm text-slate-400">Public Network</li>
            <template v-for="(bc, index) in BCPlatform" :key="index">
                <li v-if="bc.blockchain_type === 'Public' && bc.activate === true"
                    class="my-2 hover:bg-slate-500 p-2 cursor-pointer hover:text-white  items-center" :value="index"
                    @click="selectNetwork($event)">
                    {{ bc.network_name }}
                </li>
            </template>
            <li class="text-sm text-slate-400">Private Network</li>
            <template v-for="(bc, index) in BCPlatform" :key="index">
                <li v-if="bc.blockchain_type === 'Private' && bc.activate === true"
                    class="my-2 hover:bg-slate-500 p-2 cursor-pointer hover:text-white  items-center" :value="index"
                    @click="selectNetwork($event)">
                    {{ bc.network_name }}
                </li>
            </template>
            <!-- <li v-for="(bc, index) in BCPlatform" :key="index"
                class="my-2 hover:bg-slate-500 p-2 cursor-pointer hover:text-white  items-center" :value="index"
                @click="selectNetwork($event)">
                <template v-if="bc.blockchain_type === 'private'">
                    {{ bc.network_name }}

                </template>
            </li> -->
        </ul>

        <ul>
            <!-- <li v-for="(bc, index) in BCPlatform" :key="index"
                class="my-2 hover:bg-slate-500 p-2 cursor-pointer hover:text-white  items-center"
                :value="index" @click="selectNetwork($event)">
                <template v-if="bc.blockchain_type == 'public'">
                        <p>Public Network</p>                    
                        {{ bc.network_name }}

                </template>
                <template></template>
            </li> -->
            <!-- <li v-for="(bc, index) in BCPlatform" :key="index"
                class="my-2 hover:bg-slate-500 p-2 cursor-pointer hover:text-white flex justify-between items-center"
                :value="index" @click="selectNetwork($event)">
                    <span>Private</span>
                {{
                    bc.network_name }} </li> -->
            <li class=""><button @click="addBCNetwork"
                    class="px-4 py-2 hover:bg-blue-600 hover:text-white border  rounded-lg text-blue-600 font-bold"><i
                        class="pi pi-plus"></i>Tambahkan Jaringan</button></li>
        </ul>

    </Dialog>
</template>