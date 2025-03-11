<template>
    <div class="container">
        <div class="w-full bg-white shadow-md p-2">
            <div class="flex justify-between p-2">
                <div class="flex space-x-2 items-center">
                    <p class="text-gray-500">Platform</p>
                    <button @click="dialogSelectplatforms = true"
                        class="rounded-full bg-slate-300 py-2 px-4 hover:opacity-80">
                        <span v-if="platformsActivate">{{ platformsActivate?.name ?? "Pilih Platform" }}</span>
                        <i class="ml-2 pi pi-angle-up "></i></button>
                </div>
                <div v-show="platformsActivate?.name" class="flex items-center">
                    <button class="bg-red-400 py-2 px-3 rounded-full hover:opacity-70 flex items-center gap-2"
                        @click="platformDiactive"><i class="pi pi-times"></i> Disconect</button>
                </div>
            </div>
        </div>
        <div>
            <RouterView></RouterView>
        </div>
    </div>
    <Dialog v-model:visible="dialogSelectplatforms" modal header="Pilih Platform" position="top">
        <div>
            <div v-for="platform in platforms" :key="platform.id" class="my-2 flex space-x-1">
                <input type="radio" :name="platform.id" :id="platform.id" :value="platform.name"
                    v-model="selectedPlatform">
                <label :for="platform.id">{{ platform?.name }}</label>
            </div>

            <div class="flex justify-center">
                <button type="button" class="bg-blue-400 p-2 rounded-md text-white" @click="saveSelection">
                    Simpan
                </button>
            </div>
        </div>
    </Dialog>
</template>

<script setup>
import { useStore } from "vuex";
const store = useStore();
import { computed, ref, watch, onMounted } from 'vue';

// Dialog
import Dialog from 'primevue/dialog';

const platformsActivate = ref({})
const platforms = ref(null)

const dialogSelectplatforms = ref(false)
const selectedPlatform = ref()
const fetchPlatforms = async () => {
    payload.schemaname = store.getters["sekolahService/getTabeltenant"]?.schemaName
    const res = await store.dispatch("scService/fetchBCPlatform", payload)
    platforms.value = res.bcPlatform
    platforms.value.filter(platform => {
        if (platform.active) {
            selectedPlatform.value = platform.name
            platformsActivate.value = platform
        }
    })
}
let payload = {}
// Fungsi untuk mengubah platform yang aktif
const setPlatform = async () => {
    if (!platformsActivate.value) return;

    payload.bc_platform = { ...platformsActivate.value }; // Pastikan objek tidak reaktif
    console.log("📡 Payload yang dikirim:", JSON.stringify(payload));

    try {
        const res = await store.dispatch("scService/setBCPlatform", payload);
        console.log("✅ Response backend:", res);
    } catch (error) {
        console.error("❌ Error mengirim ke backend:", error);
    }
};
// Fungsi untuk memilih platform baru
const saveSelection = async () => {
    dialogSelectplatforms.value = false;

    // Set semua platform ke `false`
    platforms.value.forEach((platform) => (platform.active = false));

    // Temukan platform yang dipilih dan aktifkan
    platformsActivate.value =
        platforms.value.find((p) => p.name === selectedPlatform.value) || null;

    if (platformsActivate.value) {
        platformsActivate.value.active = true;
    }

    // Update ke backend
    await setPlatform();
};
// Set diaktive current Blockchain network
const platformDiactive = async () => {
    if (!platformsActivate.value) return;

    // Set semua platform ke `false`
    platforms.value.forEach((platform) => (platform.active = false));

    // Pastikan platform yang aktif diubah menjadi false
    platformsActivate.value.active = false;

    console.log("🚀 Sebelum mengirim ke backend:", JSON.stringify(platformsActivate.value));


    // Kirim perubahan ke backend
    await setPlatform();

    // Pastikan platform ter-reset setelah update sukses
    platformsActivate.value = {};
    selectedPlatform.value = '';
}
onMounted(() => {
    fetchPlatforms()
});

</script>

<style lang="scss" scoped></style>