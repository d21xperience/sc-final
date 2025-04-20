<template>

    <div class="container">
        <!-- <div class="fixed top-0 w-full left-0 z-20 bg-white"> -->
        <div class="my-2 ">
            <div class=" ">
                <div class="flex flex-wrap justify-between items-center mb-2">
                    <h4 class="text-xl md:text-2xl">Data Dapodik {{ tabelTenant?.namaSekolah }}</h4>
                    <div class="md:flex md:items-center md:space-x-2">
                        <h3 class="text-slate-500 md:text-base text-sm">Tahun Pelajaran</h3>
                        <div>
                            <Select v-model="selectedSemester" :options="semester" optionLabel="namaSemester"
                                placeholder="Tahun Pelajaran" class="w-full md:w-52 mr-2" :disabled="isDisabled" />

                        </div>
                    </div>
                </div>
                <!-- Breadcrumb -->
                <!-- <div>
                    <Breadcrumb :home="home" :model="breadcrumbItems">
                        <template #item="{ item, props }">
                            <router-link v-if="item.route" v-slot="{ href, navigate }" :to="item.route" custom>
                                <a :href="href" v-bind="props.action" @click="navigate">
                                    <span class="text-primary font-semibold">{{ item.label }}</span>
                                </a>
                            </router-link>
                            <span v-else class="text-surface-700 dark:text-surface-0">{{ item.label }}</span>
                        </template>
</Breadcrumb>
</div> -->
            </div>
        </div>
        <!-- </div> -->
    </div>
    <div>
        <RouterView />
    </div>
</template>

<script setup>
import { ref, onMounted, computed, watch } from 'vue';
import { useStore } from "vuex";
import { useRoute } from 'vue-router';

const route = useRoute();
const store = useStore();

// console.log(route.matched) // Cek apakah child route aktif

// =====================================
import Breadcrumb from 'primevue/breadcrumb';
const isDisabled = computed(() => route.meta.disableSelect);


// Breadcrumb
// ==============================
// Mencari parent route
const parentRoute = computed(() => {
    return route.matched.length > 1 ? route.matched[route.matched.length - 2] : null;
});

// Home Breadcrumb akan menunjuk ke parent route (data-kelas)
const home = computed(() => {
    return parentRoute.value
        ? {
            icon: "pi pi-home",
            route: parentRoute.value.path,
            label: parentRoute.value.meta.title || "Home",
        }
        : {
            icon: "pi pi-home",
            route: "/",
            label: "Home",
        };
});

// Breadcrumb items
const breadcrumbItems = computed(() => {
    return route.matched
        .filter((r) => r !== parentRoute.value) // Hapus parent agar tidak duplikasi
        .map((r) => ({
            label: r.meta.breadcrumb || r.meta.title || r.name, // Gunakan meta.breadcrumb atau title
            route: r.path !== route.path ? r.path : null, // Halaman aktif tidak diklik
        }));
});
// ==============================
onMounted(async () => {
    fetchSemester()
    fetchTabelTenant()
});

// const dataConnected = ref(true)

import Select from 'primevue/select';

// ==================================
// =======SEMESTER=============
const selectedSemester = ref(null);
const semester = ref(null);
const fetchSemester = async () => {
    try {
        const results = await store.dispatch("sekolahService/fetchSemester")
        if (results) {
            semester.value = store.getters["sekolahService/getSemester"]
            // Cek apakah di vuex ada nilai
            selectedSemester.value = await store.getters["sekolahService/getSelectedSemester"]
            if (selectedSemester.value == null) {
                // jika tidak ada, ambil semester terbaru berdasarkan ID terbesar
                selectedSemester.value = semester.value.reduce((latest, current) =>
                    current.semesterId > latest.semesterId ? current : latest
                );
            }
            // tetapkan semester yang dipilih
            store.commit("sekolahService/SET_SELECTEDSEMESTER", selectedSemester.value)
        }
    } catch (error) {
        console.lod(error)
    }
}


// ==================================
// =======DATA SEKOLAH=============
const tabelTenant = ref(null)
const fetchTabelTenant = async () => {
    try {
        tabelTenant.value = store.getters["sekolahService/getTabeltenant"]
        // console.log(tabelTenant.value)
        if (tabelTenant.value == null) {
            await store.dispatch("sekolahService/fetchTabeltenant")
            tabelTenant.value = store.getters["sekolahService/getTabeltenant"]
        }
    }
    catch (error) {

    }
}


watch(selectedSemester, () => {
    store.commit("sekolahService/SET_SELECTEDSEMESTER", selectedSemester.value)

})
</script>
