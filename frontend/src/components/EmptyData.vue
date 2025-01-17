<template>
    <div class="min-h-[350px]">
        <div class="text-center">
            <h2>Selamat datang <span class="font-bold">admin</span> {{ sekolah?.nama }}</h2>
            <p>Sebelum memulai mengirimkan data Sekolah ke jaringan Blockhain, silahkan siapkan data sekolah, klik
                tombol Siapkan Data</p>
            <div>
                <button type="button" class="bg-blue-200 p-2 rounded-lg hover:bg-blue-400" @click="siapkanData">Siapkan
                    Data</button>
            </div>
            <div v-show="visible" class="h-full  w-full static top-0">
                <ProgressSpinner />
            </div>
        </div>
        <!-- <div class="text-center flex justify-center items-center space-x-2 ">
            <h3 class="text-xl">Data belum ditarik,
            </h3>
            <p>Silakan lakukan tarik data dengan mendownload aplikasi terlebih
                dahulu</p>
        </div>
        <div class="text-center mt-6">
            <p>Aplikasi hanya untuk operator aplikasi</p>
            <button type="button" class="bg-blue-200 p-2 rounded-lg hover:bg-blue-400"
                @click="router.push({ name: 'syncDapodik' })">Download disini</button>
        </div> -->
    </div>

    <!-- <Dialog v-model:visible="visible" pt:root:class="!border-0 !bg-transparent" pt:mask:class="backdrop-blur-sm"> -->

    <!-- </Dialog> -->
</template>

<script setup>
import { ref, onMounted, computed } from "vue";
import ProgressSpinner from 'primevue/progressspinner';
// 🔥 Deklarasi emit untuk mengirim event ke parent
import { useStore } from "vuex";
const emit = defineEmits(["profileFetched", "fetchError"]);
const store = useStore();
const sekolah = ref(null)
const fetchSekolah = async (sekolahID) => {
    try {
        if (!sekolahID) throw new Error("User ID not found");
        // Dispatch untuk mendapatkan nama sekolah dari tabel tenant
        await store.dispatch("authService/getSekolahByID", sekolahID);
        // Ambil data terbaru dari store
        sekolah.value = store.getters["authService/getSekolah"];
    } catch (error) {
        console.error("Failed to fetch user profile:", error.message);
    }
};

onMounted(async() => {
    try {
        const sekolahID = store.state.authService.user?.sekolahId;
        const tes = await store.dispatch("sekolahService/getTabeltenant", sekolahID)
        // const tes = store.getters["sekolahService/getTabeltenant"]
        console.log(tes)
        if (tes) {
            // 🔥 Emit event ke parent bahwa data berhasil diambil
            emit("profileFetched", true);
        } else {
            fetchSekolah(sekolahID)
        }
    } catch (error) {
        // 🔥 Emit event ke parent bahwa ada error
        emit("fetchError", false);
    }

});

const fetchTabelTenant = async () => {
    try {
        const res = await store.dispatch("sekolahService/createTabeltenant", sekolah.value)
        if (res) {
            console.log(res)
        }
    } catch (error) {
        console.log(error)
    }
}
const siapkanData = async () => {
    fetchTabelTenant()
}

const visible = ref(true)
</script>

<style lang="scss" scoped></style>