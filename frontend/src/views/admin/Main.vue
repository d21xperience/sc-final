<script setup>
import { ref, computed } from "vue";
defineProps(['username']); // Parameter username dari route diterima sebagai props
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
            // router.push(`${username}`)
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
                label: 'Data Mapel',
                icon: 'pi pi-angle-double-right'
            },
            // {
            //     label: 'Sync Dapodik',
            //     icon: 'pi pi-refresh',
            //     command: () => {
            //         router.push({ name: 'syncDapodik' })
            //     }
            // },

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
                    router.push({ name: 'ipfsIjazah' })
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
                    router.push({ name: 'daftarTrx' })
                }
            },

        ]
    },
    {
        label: 'Sign Out',
        icon: 'pi pi-sign-out',
        command: () => {
            dialogSignOut.value = !dialogSignOut.value
        }
    }
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
        class="fixed top-0 w-[250px] h-screen py-1 px-4 hidden lg:block border-r-2 overflow-y-auto overflow-x-hidden z-30">
        <div class="mt-2">
            <PanelMenu :model="items" multiple class="" />
        </div>
    </nav>
    <div class="lg:ml-[250px]">
        <div class="container mt-2">
            <RouterView></RouterView>
        </div>
    </div>


    <!-- Dialog start -->
    <Dialog v-model:visible="dialogSignOut" :style="{ width: '450px' }" header="Keluar" :modal="true" position="top">
        <div class="flex items-center gap-4">
            <i class="pi pi-exclamation-triangle !text-3xl" />
            <span>Yakin akan keluar?</span>
        </div>
        <template #footer>
            <Button label="Tidak" icon="pi pi-times" text @click="dialogSignOut = false" severity="warn" />
            <Button label="Ya" icon="pi pi-check" text @click="signOut" />
        </template>
    </Dialog>
    <!-- Dialog end -->
</template>

<style scoped></style>
