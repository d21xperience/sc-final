<script setup>
import { ref, computed } from "vue";
import PanelMenu from 'primevue/panelmenu';
// import router from "@/router";
// import store from "@/store";
import 'primeicons/primeicons.css'
import Dialog from "primevue/dialog";
import Button from "primevue/button";
import { useRouter } from "vue-router";
import { useStore } from "vuex";
const router = useRouter();
const store = useStore();

const userRole = ref(store.state.authService.userRole); // Ambil role user dari Vuex

import Menubar from 'primevue/menubar';

// computed({
//     ...mapGetters("authService", ["userRole"]),
// })
// Penel Menu

const menuItems = ref([
    {
        label: 'Home',
        icon: 'pi pi-home',
        command: () => {
            router.push({ path: '/' })
        }

    },
    {
        label: 'Dashboard',
        icon: 'pi pi-objects-column',
        command: () => {
            router.push({ name: 'admin' })
        }

    },
    {
        label: 'Profile',
        icon: 'pi pi-user',
        command: () => {
            router.push({ name: 'profile' })
        }
    },
    {
        label: 'Data Akademik',
        icon: 'pi pi-file',
        // command: () => {
        //     router.push({ name: 'admin' })
        // },
        items: [
            {
                label: 'Ketuntasan Rapor',
                icon: 'pi pi-file',
                command: () => {
                    router.push({ name: 'ketuntasanRapor' })
                }
            },
            {
                label: 'Ijazah',
                icon: 'pi pi-image',
                command: () => {
                    router.push({ name: 'dataIjazah' })
                }
                // items: [
                //     {
                //         label: 'Logos',
                //         icon: 'pi pi-image'
                //     }
                // ]
            }
        ]
    },
    {
        label: 'Blockchain',
        icon: 'pi pi-ethereum',
        items: [
            {
                label: 'Seting',
                icon: 'pi pi-cog',
                command: () => {
                    router.push({ name: 'setingBlockchain' })
                }
            },
            {
                label: 'Daftar Jaringan',
                icon: 'pi pi-cloud',
                command: () => {
                    router.push({ name: 'listBCNetwork' })
                }
            },
            {
                label: 'SC-Ijazah',
                icon: 'pi pi-file',
                command: () => {
                    router.push({ name: 'scIjazah' })
                }
            },
            {
                label: 'Transaksi',
                icon: 'pi pi-chart-line',
                command: () => {
                    router.push({ name: 'scIjazah' })
                }
            },

        ]
    },
    {
        label: 'IPFS',
        icon: 'pi pi-desktop',
        items: [
            {
                label: 'Seting',
                icon: 'pi pi-cog',
                command: () => {
                    router.push({ name: 'setingBlockchain' })
                }
            },
            {
                label: 'URI Ijazah',
                icon: 'pi pi-file',
                command: () => {
                    router.push({ name: 'scIjazah' })
                }
            },
            {
                label: 'Transaksi',
                icon: 'pi pi-chart-line',
                command: () => {
                    router.push({ name: 'scIjazah' })
                }
            },

        ]
    },
    {
        label: 'Data Dapodik',
        icon: 'pi pi-tag',
        items: [

            {
                label: 'Data Sekolah',
                icon: 'pi pi-building-columns',
                command: () => {
                    router.push({ name: 'dapodikSekolah' })
                }
            },
            {
                label: 'Data Guru',
                icon: 'pi pi-graduation-cap',
                command: () => {
                    router.push({ name: 'dapodikGuru' })
                }
            },
            {
                label: 'Data Siswa',
                icon: 'pi pi-angle-double-right',

                command: () => {
                    router.push({ name: 'dapodikSiswa' })
                }
            },
            {
                label: 'Data Kelas',
                icon: 'pi pi-angle-double-right',
                command: () => {
                    router.push({ name: 'dapodikKelas' })
                }
            },
            {
                label: 'Data Mata Pelajaran',
                icon: 'pi pi-angle-double-right'
            },
            {
                label: 'Sync Dapodik',
                icon: 'pi pi-refresh',
                command: () => {
                    router.push({ name: 'syncDapodik' })
                }
            },

        ]
    },
    // {
    //     label: 'Sign Out',
    //     icon: 'pi pi-sign-out',
    //     command: () => {
    //         dialogSignOut.value = !dialogSignOut.value
    //     }
    // }
]);

// Filter menu berdasarkan role
const items = computed(() => {
    let excludedLabels = []
    if (userRole.value === "siswa") {
        excludedLabels = ["Home", "Blockchain", "IPFS", "Data Dapodik"]; // Kategori yang akan dihapus
    } else if (userRole.value === "admin") {
        excludedLabels = ["Data Akademik"]; // Kategori yang akan dihapus
    }
    return menuItems.value.filter(item => !excludedLabels.includes(item.label))
});



// SignOut
const dialogSignOut = ref(false)
const signOut = async () => {
    const resp = await store.dispatch('authService/logout');
    router.push({ name: 'home' })
}

// Logout
const onLogout = async () => {
    await store.dispatch('authService/logout');
    dialogSignOut.value = false
    router.push({ name: 'login' }); // Arahkan ke halaman login
};


</script>


<template>
    <nav
        class=" h-screen fixed top-0 left-0 min-w-[250px] py-2 px-4 hidden lg:block border-r-2 overflow-y-auto overflow-x-hidden">
        <div class="relative">
            <div class="">
                <div class="flex flex-row cent justify-center">
                    <div class="cursor-pointer" @click="router.push({ name: 'profile' })">
                        <img alt="User profile picture" class="rounded-full w-16 h-16" height="100"
                            src="https://storage.googleapis.com/a1aa/image/bS2nVV2XXD7IPNInU3xBCiOkojwX5fyy6Ne1rlaANQ9Rqs0TA.jpg"
                            width="100" />
                        <!-- <div class="absolute bottom-0 right-0 bg-blue-500 rounded-full p-1">
                            <i class="fas fa-camera text-white">
                            </i>
                        </div> -->

                    </div>

                </div>
            </div>
            <div class="border-b-1 shadow-sm mb-2">
                <h2 class="flex flex-wrap justify-center">
                    {{ store.state.authService.user?.username }}
                </h2>
                <h3 class="flex justify-center">{{ store.state.authService.user?.role ?? 'Anonymous' }}</h3>
            </div>
            <div
                class="absolute -right-6 top-0 h-6 w-6 p-[6px] cursor-pointer bg-[#007bff] flex items-center justify-center rounded-full">
                <svg xmlns="http://www.w3.org/2000/svg" fill="#fff" class="w-4 h-4" viewBox="0 0 55.752 55.752">
                    <path
                        d="M43.006 23.916a5.36 5.36 0 0 0-.912-.727L20.485 1.581a5.4 5.4 0 0 0-7.637 7.638l18.611 18.609-18.705 18.707a5.398 5.398 0 1 0 7.634 7.635l21.706-21.703a5.35 5.35 0 0 0 .912-.727 5.373 5.373 0 0 0 1.574-3.912 5.363 5.363 0 0 0-1.574-3.912z"
                        data-original="#000000" />
                </svg>

            </div>
        </div>

        <div class="py-6">
            <ul class="space-y-1">
                <!-- <li>
                    <router-link :to="{ name: 'admin' }" exact-active-class="bg-blue-100"
                        class="text-black hover:text-blue-600 text-[15px] flex items-center hover:bg-white rounded px-4 py-3 transition-all">
                        <svg xmlns="http://www.w3.org/2000/svg" fill="currentColor" class="w-[18px] h-[18px] mr-4"
                            viewBox="0 0 512 512">
                            <path
                                d="M197.332 170.668h-160C16.746 170.668 0 153.922 0 133.332v-96C0 16.746 16.746 0 37.332 0h160c20.59 0 37.336 16.746 37.336 37.332v96c0 20.59-16.746 37.336-37.336 37.336zM37.332 32A5.336 5.336 0 0 0 32 37.332v96a5.337 5.337 0 0 0 5.332 5.336h160a5.338 5.338 0 0 0 5.336-5.336v-96A5.337 5.337 0 0 0 197.332 32zm160 480h-160C16.746 512 0 495.254 0 474.668v-224c0-20.59 16.746-37.336 37.332-37.336h160c20.59 0 37.336 16.746 37.336 37.336v224c0 20.586-16.746 37.332-37.336 37.332zm-160-266.668A5.337 5.337 0 0 0 32 250.668v224A5.336 5.336 0 0 0 37.332 480h160a5.337 5.337 0 0 0 5.336-5.332v-224a5.338 5.338 0 0 0-5.336-5.336zM474.668 512h-160c-20.59 0-37.336-16.746-37.336-37.332v-96c0-20.59 16.746-37.336 37.336-37.336h160c20.586 0 37.332 16.746 37.332 37.336v96C512 495.254 495.254 512 474.668 512zm-160-138.668a5.338 5.338 0 0 0-5.336 5.336v96a5.337 5.337 0 0 0 5.336 5.332h160a5.336 5.336 0 0 0 5.332-5.332v-96a5.337 5.337 0 0 0-5.332-5.336zm160-74.664h-160c-20.59 0-37.336-16.746-37.336-37.336v-224C277.332 16.746 294.078 0 314.668 0h160C495.254 0 512 16.746 512 37.332v224c0 20.59-16.746 37.336-37.332 37.336zM314.668 32a5.337 5.337 0 0 0-5.336 5.332v224a5.338 5.338 0 0 0 5.336 5.336h160a5.337 5.337 0 0 0 5.332-5.336v-224A5.336 5.336 0 0 0 474.668 32zm0 0"
                                data-original="#000000" />
                        </svg>
                        <span>Dashboard</span>
                    </router-link>
                </li> -->
                <li>
                    <PanelMenu :model="items" multiple class="w-full" />
                </li>

            </ul>
        </div>
    </nav>
    <div class="ml-0 lg:ml-[250px] pl-4 pt-2 pr-4 ">
        <!-- <div class="card">
            <Menubar :model="menuItems">
                <template #start>
                    <svg width="35" height="40" viewBox="0 0 35 40" fill="none" xmlns="http://www.w3.org/2000/svg"
                        class="h-8">
                        <path
                            d="M25.87 18.05L23.16 17.45L25.27 20.46V29.78L32.49 23.76V13.53L29.18 14.73L25.87 18.04V18.05ZM25.27 35.49L29.18 31.58V27.67L25.27 30.98V35.49ZM20.16 17.14H20.03H20.17H20.16ZM30.1 5.19L34.89 4.81L33.08 12.33L24.1 15.67L30.08 5.2L30.1 5.19ZM5.72 14.74L2.41 13.54V23.77L9.63 29.79V20.47L11.74 17.46L9.03 18.06L5.72 14.75V14.74ZM9.63 30.98L5.72 27.67V31.58L9.63 35.49V30.98ZM4.8 5.2L10.78 15.67L1.81 12.33L0 4.81L4.79 5.19L4.8 5.2ZM24.37 21.05V34.59L22.56 37.29L20.46 39.4H14.44L12.34 37.29L10.53 34.59V21.05L12.42 18.23L17.45 26.8L22.48 18.23L24.37 21.05ZM22.85 0L22.57 0.69L17.45 13.08L12.33 0.69L12.05 0H22.85Z"
                            fill="var(--p-primary-color)" />
                        <path
                            d="M30.69 4.21L24.37 4.81L22.57 0.69L22.86 0H26.48L30.69 4.21ZM23.75 5.67L22.66 3.08L18.05 14.24V17.14H19.7H20.03H20.16H20.2L24.1 15.7L30.11 5.19L23.75 5.67ZM4.21002 4.21L10.53 4.81L12.33 0.69L12.05 0H8.43002L4.22002 4.21H4.21002ZM21.9 17.4L20.6 18.2H14.3L13 17.4L12.4 18.2L12.42 18.23L17.45 26.8L22.48 18.23L22.5 18.2L21.9 17.4ZM4.79002 5.19L10.8 15.7L14.7 17.14H14.74H15.2H16.85V14.24L12.24 3.09L11.15 5.68L4.79002 5.2V5.19Z"
                            fill="var(--p-text-color)" />
                    </svg>
                </template>
<template #item="{ item, props, hasSubmenu, root }">
                    <a v-ripple class="flex items-center" v-bind="props.action">
                        <span>{{ item.label }}</span>
                        <Badge v-if="item.badge" :class="{ 'ml-auto': !root, 'ml-2': root }" :value="item.badge" />
                        <span v-if="item.shortcut"
                            class="ml-auto border border-surface rounded bg-emphasis text-muted-color text-xs p-1">{{
                                item.shortcut }}</span>
                        <i v-if="hasSubmenu"
                            :class="['pi pi-angle-down ml-auto', { 'pi-angle-down': root, 'pi-angle-right': !root }]"></i>
                    </a>
                </template>
<template #end>
                    <div class="flex items-center gap-2">
                        <InputText placeholder="Search" type="text" class="w-32 sm:w-auto" />
                        <Avatar image="https://primefaces.org/cdn/primevue/images/avatar/amyelsner.png"
                            shape="circle" />
                    </div>
                </template>
</Menubar>
</div> -->
        <RouterView></RouterView>
    </div>

    <footer class="bg-gray-100 text-gray-600 p-4 fixed bottom-0 w-full">
        <button type="button" @click="dialogSignOut = true" class="w-[150px] hover:text-blue-400"><i class="pi pi-sign-out mr-2"></i>Logout</button>
    </footer>


    <!-- Dialog start -->
    <Dialog v-model:visible="dialogSignOut" :style="{ width: '450px' }" header="Sign Out" :modal="true" position="top">
        <div class="flex items-center gap-4">
            <i class="pi pi-exclamation-triangle !text-3xl" />
            <span>Yakin akan ke luar?</span>
        </div>
        <template #footer>
            <Button label="Tidak" icon="pi pi-times" text @click="dialogSignOut = false" severity="warn" />
            <Button label="Ya" icon="pi pi-check" text @click="signOut" />
        </template>
    </Dialog>

    <!-- Dialog end -->





</template>

<style scoped>
.p-accordionheader {
    font-weight: lighter;
    color: black;
    padding-bottom: 12px;
    padding-top: 12px;
    /* background-color: white */
}
</style>
