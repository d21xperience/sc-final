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
                                placeholder="Tahun Pelajaran" class="w-full md:w-52 mr-2" />

                        </div>
                    </div>
                </div>
                <!-- Breadcrumb -->
                <div>
                    <Breadcrumb :home="home" :model="breadcrumbItems">
                        <template #item="{ item, props }">
                            <router-link v-if="item.route" v-slot="{ href, navigate }" :to="item.route" custom>
                                <a :href="href" v-bind="props.action" @click="navigate">
                                    <span :class="[item.icon, 'text-color']" />
                                    <span class="text-primary font-semibold">{{ item.label }}</span>
                                </a>
                            </router-link>
                            <a v-else :href="item.url" :target="item.target" v-bind="props.action">
                                <span class="text-surface-700 dark:text-surface-0">{{ item.label }}</span>
                            </a>
                        </template>

                    </Breadcrumb>
                </div>
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



// Breadcrumb
// ==============================
const home = ref({
    icon: 'pi pi-home',
    route: `${route.fullPath}`,
});
const breadcrumbItems = computed(() => {
    const pathArray = route.path.split('/').filter((p) => p);
    let path = '';

    return pathArray.map((segment, index) => {
        path += `/${segment}`;
        return {
            label: segment.charAt(0).toUpperCase() + segment.slice(1),
            to: index < pathArray.length - 1 ? path : undefined // Hanya halaman terakhir yang tidak memiliki link
        };
    });
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
