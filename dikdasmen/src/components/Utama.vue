<script setup>
import { loadWeb3, getAccounts } from '@/utils/web3';
import { ref } from 'vue';

const account = ref("")
const connectToWallet = async () => {
    await loadWeb3()
    const accounts = await getAccounts()
    if (accounts.length > 0) {
        account.value = accounts[0]
        console.log(account.value)
    } else {
        console.error('No accounts found')
    }
}

</script>

<template>
    <h1>Hello WEB3!</h1>
    <div>
        <button @click="connectToWallet">Connect to Metamask</button>
    </div>
</template>