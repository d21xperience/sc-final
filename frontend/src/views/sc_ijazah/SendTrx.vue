<template>
    <div>
        <h1>Hello from send TRX</h1>
        <div class="border mx-auto p-3 w-1/2 rounded-sm">
            <div>
                <label for="sender">Pengirim</label>
                <input v-model="sender" type="text" name="sender" id="sender" class="border p-2 w-full"
                    placeholder="Masukan alamat pengirim">

            </div>
            <div>
                <label for="receiver">Penerima</label>
                <input v-model="receiver" type="text" name="receiver" id="receiver" class="border p-2 w-full"
                    placeholder="Masukan alamat penerima">
            </div>

            <div class="flex justify-center my-2">
                <Button label="Kirim" severity="success" icon="pi pi-briefcase" @click="sendCrypto"></Button>
            </div>
        </div>

    </div>
</template>

<script setup>
import axios from 'axios';
import { Button } from 'primevue';



const sender = ref('')
const receiver = ref('')
// ---------------Transasksi Blockchain
const api = axios.create({
    baseURL: 'http://localhost:8081', // URL utama API
    timeout: 5000, // Waktu maksimal request (ms)
    headers: {
        'Content-Type': 'application/json', // Header default
    },
});
const sendCrypto = async (id) => {
    try {
        const resp = api.post(`/api/v1/account/${id}/send-crypto`, {
            sender: sender.value,
            receiver: receiver.value,
            amount: 1
        })
    } catch (error) {
        console.log("error", error)
    }
}
</script>

<style lang="scss" scoped></style>