<template>
    <div>
        <div class="max-w-4xl mx-auto p-6 bg-white shadow-md rounded-lg mt-10">
            <div class="flex justify-end">
                <button @click="editProfile" class="hover:text-red-500" :class="{ editProfileClass: isProfileEdit }"><i
                        class="pi pi-user-edit" style="font-size: 1.5rem;"></i></button>
            </div>
            <div class="flex flex-wrap">
                <div class="w-full md:w-1/2">
                    <div class="flex items-center space-x-6">
                        <div class="relative">
                            <img alt="Profile picture of a person" class="w-32 h-32 rounded-full block" height="150"
                                src="https://storage.googleapis.com/a1aa/image/bS2nVV2XXD7IPNInU3xBCiOkojwX5fyy6Ne1rlaANQ9Rqs0TA.jpg"
                                width="150" />
                            <div
                                class="absolute  w-full h-full top-0 left-0 flex items-center justify-center opacity-0 rounded-full hover:opacity-100 bg-black bg-opacity-50 text-gray-100 transition ease-in-out duration-500">
                                <i @click="triggerFileInput" class="pi pi-camera cursor-pointer "
                                    style="font-size: 1.5rem;"></i>
                                <!-- Input file (disembunyikan) -->
                                <input id="fileInput" type="file" accept="image/*" style="display: none;"
                                    @change="handleFileSelect" />
                            </div>
                        </div>
                        <div class="text-center">
                            <h1 class="text-2xl font-bold">
                                {{ akun.username }}
                            </h1>
                            <p class="text-gray-600">
                                {{ akun.role.toUpperCase() }}
                            </p>
                        </div>
                    </div>
                    <div class="mt-6">
                        <h2 class="text-xl font-semibold">
                            Personal Information
                        </h2>
                        <div class="mt-4">
                            <p>
                                <span class="font-bold">
                                    Email:
                                </span>
                                {{ akun.email }}
                            </p>
                            <p>
                                <span class="font-bold">
                                    Phone:
                                </span>
                                +123 456 7890
                            </p>
                            <div>
                                <span class="font-bold">
                                    Nama:
                                </span>
                                <div class="inline-block">
                                    <div v-if="isProfileEdit">
                                        {{ akun.nama }}
                                    </div>
                                    <div v-else>
                                        {{ akun.nama }}
                                    </div>

                                </div>
                            </div>
                            <div>
                                <span class="font-bold">
                                    Tanggal Lahir:
                                </span>
                                <div class="inline-block">
                                    <div v-if="isProfileEdit">
                                        <DatePicker v-model="akun.tglLahir" dateFormat="dd/mm/yy"
                                            :inputClass="{ dateClass: isProfileEdit }" />
                                    </div>
                                    <div v-else>
                                        {{ akun.tglLahir }}
                                    </div>

                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <div class="w-1/2">
                    <div class="mt-6">
                        <h2 class="text-xl font-semibold">
                            Alamat
                        </h2>
                        <div class="mt-4">
                            <p>
                                <span class="font-bold">
                                    Nama Jalan:
                                </span>
                                <input type="text" name="nama-jalan" id="nama-jalan" v-model="akun.alamatJalan"
                                    :class="{ 'p-inputtext': isProfileEdit }">
                                <!-- 123 Main St -->
                            </p>
                            <p>
                                <span class="font-bold">
                                    Kota/Kab:
                                </span>
                                <input type="text" name="kota-kab" id="kota-kab" v-model="akun.kotaKab"
                                    :class="{ 'p-inputtext': isProfileEdit }">
                            </p>
                            <p>
                                <span class="font-bold">
                                    Provinsi:
                                </span>
                                <input type="text" name="prov" id="prov" v-model="akun.prov"
                                    :class="{ 'p-inputtext': isProfileEdit }">
                            </p>
                            <p>
                                <span class="font-bold">
                                    Kode Pos:
                                </span>
                                <input type="text" name="kodepos" id="kodepos" v-model="akun.kodePos"
                                    :class="{ 'p-inputtext': isProfileEdit }">
                            </p>
                        </div>
                    </div>
                    <div class="mt-6">
                        <h2 class="text-xl font-semibold">
                            Orang tua
                        </h2>
                        <div class="mt-4">
                            <p>
                                <span class="font-bold">
                                    Nama Ayah:
                                </span>
                                <input type="text" name="tgl-lahir" id="nm-ayah" v-model="akun.namaAyah"
                                    :class="{ 'p-inputtext': isProfileEdit }">
                            </p>
                            <p>
                                <span class="font-bold">
                                    Nama Ibu:
                                </span>
                                <input type="text" name="tgl-lahir" id="nama-ibu" v-model="akun.namaIbu"
                                    :class="{ 'p-inputtext': isProfileEdit }">
                            </p>

                        </div>
                    </div>
                </div>
            </div>
            <div v-show="isProfileEdit" class="flex space-x-4 mt-16">
                <div>
                    <button @click="editProfile"
                        class="bg-yellow-400 hover:bg-yellow-500 hover:text-white p-2 rounded-lg w-40">Batal</button>
                </div>
                <div>
                    <button class="bg-blue-400 hover:bg-blue-500 hover:text-white p-2 rounded-lg w-40"
                        @click="showUpdateProfile">Simpan</button>
                </div>
            </div>
        </div>
        <div class="max-w-4xl mx-auto p-2 bg-white shadow-md rounded-lg mt-10">
            <h3 class="font-bold text-xl">Change email & password</h3>
            <form action="" method="post">
                <div>
                    <label for="email">Email</label>
                    <input class="block border w-full" type="email" name="email" id="email">
                </div>
                <div>
                    <label for="password">Password</label>
                    <input class="block border w-full" type="password" name="password" id="password">
                </div>
                <div>
                    <button>Batal</button>
                    <button>Simpan</button>
                </div>

            </form>
        </div>
    </div>
</template>

<script setup>

import DatePicker from 'primevue/datepicker';

// import Tabs from 'primevue/tabs';
// import TabList from 'primevue/tablist';
// import Tab from 'primevue/tab';
// import TabPanels from 'primevue/tabpanels';
// import TabPanel from 'primevue/tabpanel';


// ==========[RPFOLE]-----------
import { ref, onMounted, watch, computed } from "vue";
import { useStore } from "vuex";

const store = useStore();
const fetchUserProfile = async () => {
    try {
        const userId = store.state.authService.user?.userId;
        if (!userId) throw new Error("User ID not found");

        // Dispatch untuk mendapatkan profil pengguna
        await store.dispatch("authService/getUserProfile", userId);

        // Ambil data terbaru dari store
        akun.value = store.getters["authService/getUserProfile"];

    } catch (error) {
        console.error("Failed to fetch user profile:", error.message);
    }
};
const showUpdateProfile = async () => {
    try {
        await store.dispatch("authService/updateUserProfile", akun.value);
    } catch (error) {
        console.log(error)
    }
}


onMounted(fetchUserProfile);
// watch(userProfile, (newVal) => {
//     console.log("UserProfile updated:", newVal);
// });
// ==============================
// State untuk menyimpan file yang dipilih
const selectedFile = ref(null);

// Fungsi untuk menangani pemilihan file
const handleFileSelect = (event) => {
    const file = event.target.files[0];
    if (file) {
        selectedFile.value = file; // Simpan file yang dipilih
        console.log("Selected file:", file);
    }
};

// Fungsi untuk memicu input file secara programatis
const triggerFileInput = () => {
    const fileInput = document.getElementById("fileInput");
    if (fileInput) fileInput.click();
};
// Biodata
const akun = ref({
    userId: "6",
    username: "administrator",
    email: "example@gmail.com",
    role: "admin",
    sekolahId: "",
    nama: "",
    jk: "",
    phone: "",
    tptLahir: "",
    tglLahir: "",
    alamatJalan: "",
    kotaKab: "",
    prov: "",
    kodePos: "",
    namaAyah: "",
    namaIbu: "",
    photoUrl: ""

})
const isProfileEdit = ref(false)
const editProfile = () => {
    isProfileEdit.value = !isProfileEdit.value
}
// const 'p-inputtext' = ref('edit')
// const editProfileClass = ref('tes')



</script>

<style scoped>
.p-inputext {
    padding: 0;
}

.editProfileClass {
    color: red
}
</style>