<template>
    <div class="container">
        <Toolbar>
            <template #start>
                <h2 class="text-xl mb-2 font-bold">Data Users</h2>
            </template>
            <template #end>
                <Select v-model="selectedSemester" :options="semester" optionLabel="namaSemester"
                    placeholder="Tahun Pelajaran" class="w-full md:w-52 mr-2" />
            </template>
        </Toolbar>
        <div class="card">
            <Tabs value="/dashboard">
                <TabList>
                    <Tab v-for="(tab, index) in items" :key="index" :value="tab.value">
                        <router-link v-if="tab.route" v-slot="{ href, navigate }" :to="tab.route" custom>
                            <a v-ripple :href="href" @click="navigate" class="flex items-center gap-2 text-inherit">
                                <i :class="tab.icon" />
                                <span>{{ tab.label }}</span>
                            </a>
                        </router-link>
                    </Tab>
                </TabList>
            </Tabs>
        </div>
        <router-view></router-view>
    </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue';
// STORE
import { useStore } from "vuex";
const store = useStore();
// ===========================
// ========================
// PRIMEVUE
import Toolbar from 'primevue/toolbar';
import Button from 'primevue/button';
import Select from 'primevue/select';
import Dialog from 'primevue/dialog';
import Tabs from 'primevue/tabs';
import TabList from 'primevue/tablist';
import Tab from 'primevue/tab';
// ========================
const items = ref([
    // { route: { name: 'usersAdmin' }, label: 'Administrator', icon: 'pi pi-users' },
    { route: { name: "usersSiswa" }, label: 'Siswa', icon: 'pi pi-users', value: 0 },
    // { route: '/products', label: 'Products', icon: 'pi pi-list' },
    // { route: '/messages', label: 'Messages', icon: 'pi pi-inbox' }
]);


// ==================================
// =======SEMESTER=============
const selectedSemester = ref(null);
const semester = ref(null);
const fetchSemester = async () => {
    try {
        semester.value = await store.getters["sekolahService/getSemester"]
        if (!semester.value || semester.value.length === 0) {
            semester.value = store.dispatch("sekolahService/fetchSemester")
            // Cek apakah di vuex ada nilai
        }
        // tetapkan semester yang dipilih
        selectedSemester.value = await store.getters["sekolahService/getSelectedSemester"]
        if (!selectedSemester.value) {
            selectedSemester.value = semester.value.reduce((latest, current) =>
                current.semesterId > latest.semesterId ? current : latest
            );
        }
    } catch (error) {
        console.lod(error)
    }
}
watch(selectedSemester, () => {
    store.commit("sekolahService/SET_SELECTEDSEMESTER", selectedSemester.value)
})
onMounted(() => {
    fetchSemester()
});
// ==================================





</script>

<style lang="scss" scoped></style>