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

        <Button @click="connectMetaMask" v-if="accounts.length === 0" label="Connect to MetaMask" icon="pi pi-user" />
        <!-- <Button @click="connectMetaMask" v-if="accounts.length === 0" label="Connect to MetaMask"/> -->
        <Button @click="disconnect" v-if="accounts.length > 0" label="Disconnect" />

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
