<template>
    <div class="card">
        <div class="flex flex-wrap justify-between my-2">
            <h4 class="font-bold text-2xl lg:text-lg my-2">Daftar Jaringan Blockhain </h4>
        </div>
        <Toolbar class="mb-6">
            <template #end>
                <Button label="New" icon="pi pi-plus" severity="success" class="mr-2"
                    @click="router.push({ name: 'addBCNetworks' })" />
                <Button label="Edit" icon="pi pi-pencil" severity="warn" @click="editBCNetwork"
                    :disabled="!selectedBCNetworks || !selectedBCNetworks.length || selectedBCNetworks.length >= 2"
                    class="mr-2" />
                <Button label="Delete" icon="pi pi-trash" severity="danger" class="mr-2" @click="confirmDeleteSelected"
                    :disabled="!selectedBCNetworks || !selectedBCNetworks.length" />
                <Button label="Export" icon="pi pi-upload" severity="help" @click="exportCSV($event)" class="mr-2" />
                <!-- <Button label="Aktifkan" icon="pi pi-check" severity="info" @click="activeBCNetworkDialog = true"
                    class="mr-2"
                    :disabled="!selectedBCNetworks || !selectedBCNetworks.length || selectedBCNetworks.length >= 2" /> -->
                <Button v-slot="slotProps" asChild>
                    <button @click="activeBCNetworkDialog = true"
                        class="mr-2 border rounded-lg p-2 hover:bg-slate-400 cursor-pointer"
                        :disabled="!selectedBCNetworks || !selectedBCNetworks.length || selectedBCNetworks.length >= 2">
                        <span class="font-medium">Aktifkan</span>
                    </button>
                </Button>

            </template>
        </Toolbar>

        <DataTable ref="dt" v-model:selection="selectedBCNetworks" :value="BCNetworks" dataKey="network_id"
            :paginator="true" :rows="10" :filters="filters"
            paginatorTemplate="FirstPageLink PrevPageLink PageLinks NextPageLink LastPageLink CurrentPageReport RowsPerPageDropdown"
            :rowsPerPageOptions="[5, 10, 25]"
            currentPageReportTemplate="Showing {first} to {last} of {totalRecords} BCNetworks">
            <template #header>
                <div class="flex flex-wrap gap-2 items-center justify-between">
                    <div class="flex">
                        <!-- <Select v-model="selectedJurusan" :options="jurusan" optionLabel="name" placeholder="Rombel"
                                class="w-full md:w-56 mr-2" />
                            <Select v-model="selectedJurusan" :options="jurusan" optionLabel="name"
                                placeholder="Kompetensi Keahlian" class="w-full md:w-56 mr-2" /> -->
                        <IconField>
                            <InputIcon>
                                <i class="pi pi-search" />
                            </InputIcon>
                            <InputText v-model="filters['global'].value" placeholder="Search..." />
                        </IconField>
                    </div>

                </div>
            </template>

            <Column selectionMode="multiple" style="width: 3rem" :exportable="false"></Column>
            <Column field="network_name" header="Nama Jaringan">
                <template #body="slotProps">
                    {{ slotProps.data.network_name }}
                </template>
            </Column>
            <Column field="activate" header="Active" sortable>
                <template #body="slotProps">
                    <!-- <button @click="activeBCNetworkDialog = true" class="rounded-full border border-slate-500 w-4 h-4"
                        :class="{ 'bg-green-800': slotProps.data.activate }"></button> -->
                    <div class="rounded-full border border-slate-500 w-4 h-4 "
                        :class="[slotProps.data.activate ? 'bg-green-400' : 'bg-slate-400']"></div>

                </template>
            </Column>
            <Column field="blockchain_type" header="Tipe Jaringan" sortable>
                <template #body="slotProps">
                    {{ slotProps.data.blockchain_type }}
                </template>
            </Column>
            <Column field="chain_id" header="Chain ID">
                <template #body="slotProps">
                    {{ slotProps.data.chain_id }}
                </template>
            </Column>
            <Column field="rpc_url" header="RPC URL">
                <template #body="slotProps">
                    {{ slotProps.data.rpc_url }}
                </template>
            </Column>
            <Column field="block_explorer" header="Explorer URL">
                <template #body="slotProps">
                    {{ slotProps.data.block_explorer }}
                </template>
            </Column>
            <!-- <Column header="Image">
                    <template #body="slotProps">
                        <img :src="`https://primefaces.org/cdn/primevue/images/product/${slotProps.data.image}`"
                            :alt="slotProps.data.image" class="rounded" style="width: 64px" />
                    </template>
                </Column> -->
            <!--<Column field="rating" header="Reviews" sortable style="min-width: 12rem">
                    <template #body="slotProps">
                        <Rating :modelValue="slotProps.data.rating" :readonly="true" />
                    </template>
                </Column>-->
            <!-- <Column field="inventoryStatus" header="Status" sortable>
                <template #body="slotProps">
                    <Tag :value="slotProps.data.inventoryStatus"
                        :severity="getStatusLabel(slotProps.data.inventoryStatus)" />
                </template>
            </Column> -->
            <!-- <Column header="Aksi" :exportable="false" style="min-width: 12rem">
                    <template #body="slotProps">
                        <Button icon="pi pi-pencil" outlined rounded class="mr-2"
                            @click="editProduct(slotProps.data)" />
                        <Button icon="pi pi-trash" outlined rounded severity="danger"
                            @click="confirmDeleteProduct(slotProps.data)" />
                    </template>
                </Column> -->
        </DataTable>
    </div>

    <!--  -------------------------------------------- -->
    <!--  -------------------------------------------- -->
    <!-- EDIT BCNETWORK DIALOG -->
    <Dialog v-model:visible="BCNetworkDialog" :style="{ height: '650px', width: '450px' }" header="Edit Data"
        :modal="true">
        <div class="flex flex-wrap gap-6">
            <!-- <img v-if="BCNetwork.image" :src="`https://primefaces.org/cdn/primevue/images/product/${BCNetwork.image}`"
                    :alt="BCNetwork.image" class="block m-auto pb-4" /> -->
            <div>
                <label for="name" class="block font-bold">Nama Jaringan</label>
                <InputText id="name" v-model.trim="selectedBCNetworks[0].network_name" required="true" autofocus
                    :invalid="submitted && !selectedBCNetworks[0].network_name" fluid />
                <small v-if="submitted && !selectedBCNetworks[0].network_name" class="text-red-500">NISN is
                    required.</small>
            </div>
            <div>
                <label for="name" class="block font-bold ">Chain ID</label>
                <InputText id="name" v-model.trim="BCNetworks.network_name" required="true" autofocus
                    :invalid="submitted && !BCNetworks.network_name" fluid />
                <small v-if="submitted && !BCNetworks.network_name" class="text-red-500">Nama is required.</small>
            </div>
            <div>
                <label for="name" class="block font-bold ">Tipe Jaringan</label>
                <InputText id="name" v-model.trim="BCNetwork.price" required="true" autofocus
                    :invalid="submitted && !BCNetwork.price" fluid />
                <small v-if="submitted && !BCNetwork.price" class="text-red-500">Nilai is required.</small>
            </div>
            <div>
                <label for="name" class="block font-bold ">RPC URL</label>
                <InputText id="name" v-model.trim="selectedBCNetworks[0].rpc_url" required="true" autofocus
                    :invalid="submitted && !selectedBCNetworks[0].rpc_url" fluid />
                <small v-if="submitted && !selectedBCNetworks[0].rpc_url" class="text-red-500">Thn lulus is
                    required.</small>
            </div>
            <!-- <div>
                    <label for="description" class="block font-bold mb-3">Description</label>
                    <Textarea id="description" v-model="BCNetwork.description" required="true" rows="3" cols="20" fluid />
                </div> -->
            <!-- <div>
                    <label for="inventoryStatus" class="block font-bold mb-3">Inventory Status</label>
                    <Select id="inventoryStatus" v-model="BCNetwork.inventoryStatus" :options="statuses"
                        optionLabel="label" placeholder="Select a Status" fluid></Select>
                </div> -->

            <!-- <div>
                    <span class="block font-bold mb-4">Category</span>
                    <div class="grid grid-cols-12 gap-4">
                        <div class="flex items-center gap-2 col-span-6">
                            <RadioButton id="category1" v-model="BCNetwork.category" name="category"
                                value="Accessories" />
                            <label for="category1">Accessories</label>
                        </div>
                        <div class="flex items-center gap-2 col-span-6">
                            <RadioButton id="category2" v-model="BCNetwork.category" name="category" value="Clothing" />
                            <label for="category2">Clothing</label>
                        </div>
                        <div class="flex items-center gap-2 col-span-6">
                            <RadioButton id="category3" v-model="BCNetwork.category" name="category"
                                value="Electronics" />
                            <label for="category3">Electronics</label>
                        </div>
                        <div class="flex items-center gap-2 col-span-6">
                            <RadioButton id="category4" v-model="BCNetwork.category" name="category" value="Fitness" />
                            <label for="category4">Fitness</label>
                        </div>
                    </div>
                </div> -->

            <!-- <div class="grid grid-cols-12 gap-4">
                    <div class="col-span-6">
                        <label for="price" class="block font-bold mb-3">Price</label>
                        <InputNumber id="price" v-model="BCNetwork.price" mode="currency" currency="USD" locale="en-US"
                            fluid />
                    </div>
                    <div class="col-span-6">
                        <label for="quantity" class="block font-bold mb-3">Quantity</label>
                        <InputNumber id="quantity" v-model="BCNetwork.quantity" integeronly fluid />
                    </div>
                </div> -->
        </div>

        <template #footer>
            <Button label="Cancel" icon="pi pi-times" text @click="hideDialog" />
            <Button label="Save" icon="pi pi-check" @click="saveProduct" />
        </template>
    </Dialog>

    <Dialog v-model:visible="activeBCNetworkDialog" :style="{ width: '450px' }" header="Confirm" :modal="true">
        <div class="flex items-center gap-4">
            <i class="pi pi-exclamation-triangle !text-3xl" />
            <!-- <span v-if="BCNetwork">Are you sure you want to delete <b>{{ BCNetwork.network_name }}</b>?</span> -->
            <span>Are you sure you want to Activing this network?</span>
        </div>
        <template #footer>
            <Button label="No" icon="pi pi-times" text @click="activeBCNetworkDialog = false" />
            <Button label="Yes" icon="pi pi-check" @click="activingBcNetwork" />
        </template>
    </Dialog>
    <Dialog v-model:visible="activeBCNetworkDialogInfo" :style="{ width: '450px' }" header="Confirm" :modal="true">
        <div class="flex items-center gap-4">
            <i class="pi pi-times !text-3xl" />
            <!-- <span v-if="BCNetwork">Are you sure you want to delete <b>{{ BCNetwork.network_name }}</b>?</span> -->
            <span>{{ activeBCNetworkInfo }}</span>
            <!-- <span>Sudah ada</span> -->
        </div>
        <template #footer>
            <Button label="Ok" icon="pi pi-check" text @click="activeBCNetworkDialogInfo = false" />
            <!-- <Button label="Yes" icon="pi pi-check" @click="activingBcNetwork" /> -->
        </template>
    </Dialog>
    <Toast />
    <Dialog v-model:visible="deleteBCNetworksDialog" :style="{ width: '450px' }" header="Confirm" :modal="true">
        <div class="flex items-center gap-4">
            <i class="pi pi-exclamation-triangle !text-3xl" />
            <span v-if="selectedBCNetworks">Apakah jaringan <span class="font-medium">{{
                BCNetworkName
                    }} </span> akan dihapus?</span>
        </div>
        <template #footer>
            <Button label="Tidak" icon="pi pi-times" text @click="deleteBCNetworksDialog = false" />
            <Button label="Ya" icon="pi pi-check" text @click="deleteBCNetwork" />
        </template>
    </Dialog>
    <!-- END OF DIALOGBOX FOR EDIT DATA -->
    <!--  -------------------------------------------- -->

</template>

<script setup>
import { ref, onMounted, watch, watchEffect } from 'vue';
import Toast from 'primevue/toast';
import router from '@/router';
import DataTable from 'primevue/datatable';
import Column from 'primevue/column';
import Button from 'primevue/button';
import Dialog from 'primevue/dialog';
import Toolbar from 'primevue/toolbar';
import { FilterMatchMode } from '@primevue/core/api';
import { useToast } from 'primevue/usetoast';
import InputText from 'primevue/inputtext';
import IconField from 'primevue/iconfield';
import InputIcon from 'primevue/inputicon';

const toast = useToast();
// import BCNetworkService from '@/service/BCNetworkService.js';
// const BCNetworks = ref([
//     {
//         network_id: 1,
//         network_name: 'Hyperledger fabric',
//         blockchain_type: 'Private',
//         chain_id: '1',
//         rpc_url: 'localhot:6786',
//         block_explorer: 'https://blockscout',
//         description: ''
//     },
//     {
//         network_id: 2,
//         network_name: 'Quorum',
//         blockchain_type: 'Private',
//         chain_id: '12',
//         rpc_url: '',
//         block_explorer: '',
//         description: ''
//     },
//     {
//         network_id: 3,
//         network_name: "Ethereum Mainnet",
//         blockchain_type: "Public",
//         chain_id: 1,
//         rpc_url: "https://mainnet.infura.io/v3/",
//         block_explorer: "https://etherscan.io"
//     },
//     {
//         network_id: 4,
//         network_name: "Binance Smart Chain (BSC)",
//         blockchain_type: "Public",
//         chain_id: 56,
//         rpc_url: "https://bsc-dataseed.binance.org",
//         block_explorer: "https://bscscan.com"
//     },
//     {
//         network_id: 5,
//         network_name: "Polygon (Matic) Mainnet",
//         blockchain_type: "Public",
//         chain_id: 137,
//         rpc_url: "https://polygon-rpc.com",
//         block_explorer: "https://polygonscan.com"
//     },
//     {
//         network_id: 6,
//         network_name: "Avalanche C-Chain",
//         blockchain_type: "Public",
//         chain_id: 43114,
//         rpc_url: "https://api.avax.network/ext/bc/C/rpc",
//         block_explorer: "https://snowtrace.io"
//     },
//     {
//         network_id: 7,
//         network_name: "Fantom Opera",
//         blockchain_type: "Public",
//         chain_id: 250,
//         rpc_url: "https://rpc.ftm.tools",
//         block_explorer: "https://ftmscan.com"
//     },
//     {
//         network_id: 8,
//         network_name: "Arbitrum One",
//         blockchain_type: "Public",
//         chain_id: 42161,
//         rpc_url: "https://arb1.arbitrum.io/rpc",
//         block_explorer: "https://arbiscan.io"
//     },
//     {
//         network_id: 9,
//         network_name: "Optimism",
//         blockchain_type: "Public",
//         chain_id: 10,
//         rpc_url: "https://mainnet.optimism.io",
//         block_explorer: "https://optimistic.etherscan.io"
//     },
//     {
//         network_id: 10,
//         network_name: "Ethereum Goerli Testnet",
//         blockchain_type: "Public",
//         chain_id: 5,
//         rpc_url: "https://goerli.infura.io/v3/",
//         block_explorer: "https://goerli.etherscan.io"
//     },
//     {
//         network_id: 11,
//         network_name: "Binance Smart Chain Testnet",
//         blockchain_type: "Public",
//         chain_id: 97,
//         rpc_url: "https://data-seed-prebsc-1-s1.binance.org:8545",
//         block_explorer: "https://testnet.bscscan.com"
//     },


// ])


const BCNetworks = ref(null)
import axios from 'axios';

// Membuat instance Axios
// Pindahkan ke vuex
const api = axios.create({
    baseURL: 'http://localhost:8081', // URL utama API
    timeout: 5000, // Waktu maksimal request (ms)
    headers: {
        'Content-Type': 'application/json', // Header default
    },
});

const fetchData = async () => {
    try {
        const { data } = await api.get('/api/v1/blockchain-networks'); // Destructuring response
        // console.log(data);
        //BCNetworks.value = data.filter(item => !item.activate)
        BCNetworks.value = data

        // let obj ={...data}
        // console.log(obj)
        // BCNetwork.value = obj
    } catch (error) {
        if (error.response) {
            console.error(`Error ${error.response.status}: ${error.response.data}`);
        } else {
            console.error('Error:', error.message);
        }
    }
};


onMounted(() => {
    fetchData()
    // BCNetworkService.getBCNetworks().then((data) => (BCNetworks.value = data));
});


const dt = ref();
const BCNetwork = ref();
const BCNetworkName = ref('');
const BCNetworkDialog = ref(false);
const addBCNetworkDialog = ref(false)
const deleteBCNetworkDialog = ref(false);
const deleteBCNetworksDialog = ref(false);
const product = ref({});
const selectedBCNetworks = ref(null);
const filters = ref({
    'global': { value: null, matchMode: FilterMatchMode.CONTAINS },
});
const submitted = ref(false);
const statuses = ref([
    { label: 'INSTOCK', value: 'instock' },
    { label: 'LOWSTOCK', value: 'lowstock' },
    { label: 'OUTOFSTOCK', value: 'outofstock' }
]);

const formatCurrency = (value) => {
    if (value)
        return value.toLocaleString('en-US', { style: 'currency', currency: 'USD' });
    return;
};


// watch(selectedBCNetworks, (newVal) => {
//     // console.log(newVal.length ===1)
//     if (newVal.length === 1) {
//         BCNetworkName.value = newVal[0].network_name
//     } else {
//         BCNetworkName.value = `${newVal[0].network_name} dan jaringan lainnya `
//     }
// })
// const openNew = () => {
//     selectedBCNetworks.value = undefined
//     BCNetwork.value = {};
//     submitted.value = false;
//     addBCNetworkDialog.value = true;
// };
const hideDialog = () => {
    BCNetworkDialog.value = false;
    submitted.value = false;
};
const saveProduct = () => {
    submitted.value = true;

    if (product?.value.name?.trim()) {
        if (BCNetwork.value.id) {
            BCNetwork.value.inventoryStatus = BCNetwork.value.inventoryStatus.value ? BCNetwork.value.inventoryStatus.value : BCNetwork.value.inventoryStatus;
            BCNetworks.value[findIndexById(BCNetwork.value.id)] = BCNetwork.value;
            toast.add({ severity: 'success', summary: 'Successful', detail: 'Product Updated', life: 3000 });
        }
        else {
            BCNetwork.value.id = createId();
            BCNetwork.value.code = createId();
            BCNetwork.value.image = 'product-placeholder.svg';
            BCNetwork.value.inventoryStatus = BCNetwork.value.inventoryStatus ? BCNetwork.value.inventoryStatus.value : 'INSTOCK';
            BCNetworks.value.push(BCNetwork.value);
            toast.add({ severity: 'success', summary: 'Successful', detail: 'Product Created', life: 3000 });
        }

        BCNetworkDialog.value = false;
        BCNetwork.value = {};
    }
};
const editBCNetwork = (prod) => {
    BCNetwork.value = { ...prod };
    BCNetworkDialog.value = true;
};
// const confirmDeleteProduct = (prod) => {
//     BCNetworks.value = prod;
//     deleteBCNetworkDialog.value = true;
// };
// const deleteProduct = () => {
//     BCNetworks.value = BCNetworks.value.filter(val => val.id !== BCNetwork.value.network_id);
//     deleteBCNetworkDialog.value = false;
//     BCNetwork.value = {};
//     toast.add({ severity: 'success', summary: 'Successful', detail: 'Product Deleted', life: 3000 });
// };
const findIndexById = (id) => {
    let index = -1;
    for (let i = 0; i < BCNetworks.value.length; i++) {
        if (BCNetworks.value[i].id === id) {
            index = i;
            break;
        }
    }

    return index;
};
const createId = () => {
    let id = '';
    var chars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
    for (var i = 0; i < 5; i++) {
        id += chars.charAt(Math.floor(Math.random() * chars.length));
    }
    return id;
}
const exportCSV = () => {
    dt.value.exportCSV();
};
const confirmDeleteSelected = (b) => {
    deleteBCNetworksDialog.value = true;
};
const deleteBCNetwork = () => {
    BCNetworks.value = BCNetworks.value.filter(val => !selectedBCNetworks.value.includes(val))
    deleteBCNetworksDialog.value = false;
    selectedBCNetworks.value = null;
    toast.add({ severity: 'success', summary: 'Successful', detail: 'BCNetworks Deleted', life: 3000 });
};

const getStatusLabel = (status) => {
    switch (status) {
        case 'INSTOCK':
            return 'success';

        case 'LOWSTOCK':
            return 'warn';

        case 'OUTOFSTOCK':
            return 'danger';

        default:
            return null;
    }
};

// Active BC Network
const activeBCNetwork = ref(null);
const activeBCNetworkDialog = ref(false);
const activeBCNetworkInfo = ref('');
const activeBCNetworkDialogInfo = ref(false);
const activingBcNetwork = () => {
    activeBCNetworkDialog.value = !activeBCNetworkDialog.value
    let cekNetwork = false
    let id = ''
    selectedBCNetworks.value.forEach(bc => {
        // console.log(bc.chain_id)
        if (bc.chain_id != 1 && bc.chain_id != 12) {
            cekNetwork = false
            activeBCNetworkInfo.value = "Maaf, saat ini jaringan belum tersedia"
        } else {
            cekNetwork = true
            id = bc.network_id
            activeBCNetworkInfo.value = "Jaringan sudah tersedia"
        }
    })
    selectedBCNetworks.value = null
    // console.log(cekNetwork)
    activeBCNetworkDialogInfo.value = true
    if (cekNetwork) {
        const tes = sendActiveBCNetwork(id)
        console.log(tes)
    }
}


const sendActiveBCNetwork = async (id) => {
    try {
        // console.log(id)
        const response = await api.put(`/api/v1/blockchain-networks/${id}`, {
            Activate: true
        })
        // console.log(response.data);
    } catch (error) {
        console.error('Error:', error);
    }
}





</script>
