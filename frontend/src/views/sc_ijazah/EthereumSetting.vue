<template>

    <!-- ETHEREUM -->



    <MetamaskConnect/>

    <!-- <div>
        <input type="text">
        <label for="address"></label>
    </div> -->
</template>

<script setup>

defineProps({
    platformSelected: {
        type: Object,

    }
})
// Dialog

import Dialog from 'primevue/dialog';
import { useDialog } from 'primevue/usedialog';
const dialog = useDialog();

import Button from 'primevue/button';
import { computed, ref, toRaw, onMounted } from 'vue';

// Ambil dari database
const BCPlatformAvailable = ref([])
const BCPlatformActivate = ref({})

const dialogSelectNetwork = ref(false)
import { useStore } from "vuex";
const store = useStore();
// Akun
const doesntHaveAccount = ref(true)
const akunBC = ref(null)
const dialogListAccounts = ref(false)
const dialogCreateAccount = ref(false)
const dialogCreateNewAccount = ref(false)
const dialogImportAccount = ref(false)

import Tes from "./Tes.vue"
import router from '@/router';
import QrCodeAccount from './QrCodeAccount.vue';
import MetamaskConnect from '@/components/MetamaskConnect.vue';


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

const bcPlatformSelected = computed(() => {
    return store.getters["scService/getBCPlatformSelected"]
    // console.log(tes)
    // if (Object.keys(tes).length > 0) {
    //     return true
    // }
    // else {
    //     return false
    // }

})
const networkActive = ref(false)
// const networkName = ref(null)
const networkIndex = ref(0)
// const networkAvailable = ref(null)
const selectNetwork = (e) => {
    let index = e.target.value
    // console.log(e.target.value)
    networkIndex.value = e.target.value
    BCPlatformActivate.value = BCPlatformAvailable.value[index]
    // simpan ke state
    store.dispatch("scService/updateBCPlatformActivate", BCPlatformActivate.value)
    networkActive.value = true
    dialogSelectNetwork.value = false
    // Ambil akun sesuai denga nama blockahain yang dipilih
    getAccount(BCPlatformActivate.value)
}
onMounted(() => {
    // fetchBlockchainNetworks();
    // getAccount()
    user.value = toRaw(store.getters["authService/getUserProfile"])
    sekolah.value = toRaw(store.getters["authService/getSekolah"])
});

const user = ref(null)
const sekolah = ref({})
// Mengambil bcplatform dari backend
// const fetchBlockchainNetworks = async () => {
//     try {
//         const response = await store.dispatch("scService/fetchBlockchainNetworks");
//         if (!response) {
//             return;
//         }
//         const { network } = response;
//         BCPlatform.value = network
//         network.forEach((bc, i) => {
//             if (bc.Available) {
//                 BCPlatformAvailable.value.push(bc)
//                 if (bc.Activate) {
//                     BCPlatformActivate.value = bc
//                     networkActive.value = true
//                 }
//             }
//         });
//     } catch (error) {
//         console.error("Error connecting to blockchain network:", error);
//     }
// }


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

const pVKey = ref('')
const importAccount = async () => {
    try {
        let payload = {
            admin: {
                sekolah_id: sekolah.sekolahData.sekolah_id,
                user_id: user.userId,
                // password: ,
                nama_sekolah: sekolah.nama,
                sekolah_id_enkrip: sekolah.sekolahData.sekolah_id_enkrip

            },
            network: toRaw(store.getters["scService/getBCPlatformActivate"]),
            private_key: pVKey.value,
            schemaname: sekolah.sekolahData.sekolah_id_enkrip
        }
        // console.log(payload)
        // let schemaname =
        //     let network = 
        // console.log(network)
        // console.log(pVKey)
        const response = await store.dispatch("scService/importBCAccount", JSON.parse(JSON.stringify(payload)));
        console.log(response)
        dialogImportAccount.value = false
    } catch (error) {
        console.log(error)
    }
}
const akun = computed(() => store.getters["scService/getBCAccount"])
// Dapatkan akun
const getAccount = async (network) => {
    try {
        doesntHaveAccount.value = true
        let payload = {
            user_id: user.value.userId,
            schemaname: sekolah.value.sekolahData.sekolah_id_enkrip,
            network_id: toRaw(store.getters["scService/getBCPlatformActivate"]).Id
        }
        const response = await store.dispatch("scService/fetchBCAccount", payload);
        console.log(response)
        if (response.code == 2) {
            // Jika schema belum dibuat, tampilkan pilihan buat akun
            doesntHaveAccount.value = true
            return;
        } else if (response.status = true) {
            doesntHaveAccount.value = false
        }

    } catch (error) {
        console.log('error', error)
    }
}

// -------------------------------------------------------------------------------




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



const dialogBuildSmartContract = ref(false)
const buildSmartContract = async () => {
    try {
        const payload = {
            // ambil akun
            account_type: akun.value.type,
            gas_limit: 300000,
            // private_key: 
        }

        const response = await store.dispatch("scService/deployIjazahContract", payload)


    } catch (error) {

    }
}
// blockchain platform
</script>