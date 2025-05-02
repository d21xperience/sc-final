<template>
    <div>
        <h1>IPFS Docker Connection Check</h1>
        <button @click="checkIpfsConnection" :disabled="isChecking">
            {{ isChecking ? 'Checking...' : 'Check IPFS Connection' }}
        </button>

        <div v-if="connectionStatus" class="status-container">
            <p :class="['status', connectionStatus.toLowerCase()]">
                Status: {{ connectionStatus }}
            </p>
            <div v-if="nodeInfo.id" class="node-info">
                <p><strong>Node ID:</strong> {{ nodeInfo.id }}</p>
                <p><strong>Version:</strong> {{ nodeInfo.version }}</p>
                <p><strong>Public Key:</strong> {{ nodeInfo.publicKey }}</p>
            </div>
            <p v-if="error" class="error">Error: {{ error }}</p>
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue';
import { create } from 'ipfs-http-client';

const isChecking = ref(false);
const connectionStatus = ref('');
const nodeInfo = ref({
    id: '',
    version: '',
    publicKey: ''
});
const error = ref('');

// Konfigurasi koneksi ke IPFS Docker
// Sesuaikan dengan konfigurasi Docker Anda
const ipfs = create({
    host: 'localhost', // atau alamat IP Docker host Anda
    port: 5001,
    protocol: 'http'
});

async function checkIpfsConnection() {
    isChecking.value = true;
    connectionStatus.value = 'Connecting...';
    error.value = '';
    nodeInfo.value = { id: '', version: '', publicKey: '' };

    try {
        // Cek versi IPFS terlebih dahulu (endpoint yang ringan)
        const version = await ipfs.version();

        // Dapatkan info node lengkap
        const id = await ipfs.id();

        nodeInfo.value = {
            id: id.id,
            version: version.version,
            publicKey: id.publicKey
        };

        connectionStatus.value = 'Connected to IPFS Docker';
    } catch (err) {
        connectionStatus.value = 'Connection failed';
        error.value = err.message;
        console.error('IPFS Docker connection error:', err);

        // Saran troubleshooting
        if (err.message.includes('ECONNREFUSED')) {
            error.value += '. Pastikan IPFS Docker container sedang berjalan dan port 5001 terekspos.';
        }
    } finally {
        isChecking.value = false;
    }
}
</script>

<style scoped>
.status {
    font-weight: bold;
    padding: 0.5rem;
    border-radius: 4px;
}

.status.connected {
    background-color: #d4edda;
    color: #155724;
}

.status.connecting {
    background-color: #fff3cd;
    color: #856404;
}

.status.connection.failed {
    background-color: #f8d7da;
    color: #721c24;
}

.error {
    color: #dc3545;
    font-weight: bold;
}

.node-info {
    margin-top: 1rem;
    padding: 1rem;
    background-color: #f8f9fa;
    border-radius: 4px;
}

button:disabled {
    opacity: 0.6;
    cursor: not-allowed;
}
</style>