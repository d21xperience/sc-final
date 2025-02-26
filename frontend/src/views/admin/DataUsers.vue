<template>
    <div class="container">
        <div class="mb-2">
            <Toolbar>
                <template #start>
                    <Button label="Tambah user" icon="pi pi-plus" severity="help" @click="exportCSV($event)"
                        class="mr-2" />
                    <Button label="Generate Otomatis" icon="pi pi-users" severity="info" @click="exportCSV($event)" />
                </template>
                <template #end>
                    <Select v-model="selectedSemester" :options="semester" optionLabel="namaSemester"
                        placeholder="Tahun Pelajaran" class="w-full md:w-52 mr-2" />
                </template>
            </Toolbar>
        </div>

        <div class="card">
            <Tabs value="/dashboard">
                <TabList>
                    <Tab v-for="tab in items" :key="tab.label" :value="tab.route">
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
import { ref, onMounted } from 'vue';

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
    { route: { name: "usersSiswa" }, label: 'Siswa', icon: 'pi pi-users' },
    // { route: '/products', label: 'Products', icon: 'pi pi-list' },
    // { route: '/messages', label: 'Messages', icon: 'pi pi-inbox' }
]);

onMounted(() => {
    fetchSemester()
});

// ==================================
// =======SEMESTER=============
const selectedSemester = ref();
const semester = ref(null);
const fetchSemester = async () => {
    try {
        const results = await store.dispatch("sekolahService/fetchSemester")
        console.log(results)
        if (results) {
            semester.value = store.getters["sekolahService/getSemester"]
            // Ambil semester terbaru berdasarkan ID terbesar
            selectedSemester.value = semester.value.reduce((latest, current) =>
                current.semesterId > latest.semesterId ? current : latest
            );
        }
    } catch (error) {

    }
}
// ==================================





</script>

<style lang="scss" scoped></style>