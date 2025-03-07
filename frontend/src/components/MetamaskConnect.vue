<!-- <script setup>
import { ref } from "vue";
import { loadWeb3, getBalance } from "@/utils/metamask";
import Dialog from 'primevue/dialog';
const visible = ref(false);
// Variabel reaktif untuk menyimpan data akun, saldo, dan error
const accounts = ref([]);
const balance = ref("0");
const errorMessage = ref(null);

// Fungsi untuk menghubungkan ke MetaMask
const connectMetaMask = async () => {
    errorMessage.value = null; // Reset error sebelum mencoba koneksi
    const { accounts: accs, error } = await loadWeb3();

    if (error) {
        errorMessage.value = error;
        visible.value = true
        return;
    }

    if (accs.length > 0) {
        accounts.value = accs;

        const { balance: bal, error: balanceError } = await getBalance(accs[0]);
        if (balanceError) {
            errorMessage.value = balanceError;
        } else {
            balance.value = bal;
        }
    }
}

</script>

<template>
    <div>
        <h2>MetaMask Connection</h2>
        <button @click="connectMetaMask">Connect to MetaMask</button>
        <p v-if="accounts.length">Connected Account: {{ accounts[0] }}</p>
        <p v-if="accounts.length">Balance: {{ balance }} ETH</p>
    </div>

    <Dialog v-model:visible="visible" modal header="Warning" :style="{ width: '25rem' }">
        <div>{{ errorMessage }}</div>
    </Dialog>
</template> -->
<script setup>
import { ref, onMounted } from "vue";
import { loadWeb3, getBalance, getNetwork, disconnectMetaMask, listenForAccountChange } from "@/utils/web3";

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

        // Ambil saldo dan informasi jaringan
        updateAccountData(selectedAccount.value);
    }
};

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
    }
};

// Event listener untuk perubahan akun atau jaringan
onMounted(() => {
    listenForAccountChange((accs) => {
        if (accs.length === 0) {
            disconnect();
        } else {
            accounts.value = accs;
            selectedAccount.value = accs[0]; // Pilih akun pertama saat berubah
            updateAccountData(selectedAccount.value);
        }
    });
});
</script>

<template>
    <div>
        <h2>MetaMask Connection</h2>

        <button @click="connectMetaMask" v-if="accounts.length === 0">Connect to MetaMask</button>
        <button @click="disconnect" v-if="accounts.length > 0">Disconnect</button>

        <div v-if="accounts.length > 0">
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
    </div>
</template>

<style>
.error {
    color: red;
    font-weight: bold;
}
</style>
