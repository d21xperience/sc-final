<template>
    <div class="container">
        <div class="mb-2">
            <Toolbar>
                <!-- <template #start>
                    <Button icon="pi pi-pencil" severity="warn" @click="confirmDeleteSelected"
                        :disabled="!dataLulusan || !dataLulusan.length || dataLulusan.length > 2" class="mr-2" />

                </template> -->
                <template #end>
                    <FileUpload mode="basic" accept="xlsx/*" :maxFileSize="1000000" label="Import" chooseLabel="Import"
                        class="mr-2" auto />
                    <Button label="Tambah user" icon="pi pi-plus" severity="help" @click="exportCSV($event)"
                        class="mr-2" />
                    <Button label="Generate Otomatis" icon="pi pi-users" severity="info" @click="exportCSV($event)" />
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
import Toolbar from 'primevue/toolbar';
import Button from 'primevue/button';

import Dialog from 'primevue/dialog';
import Tabs from 'primevue/tabs';
import TabList from 'primevue/tablist';
import Tab from 'primevue/tab';
import { ref } from 'vue';

const items = ref([
    { route: { name: 'usersAdmin' }, label: 'Administrator', icon: 'pi pi-users' },
    { route: { name: "usersSiswa" }, label: 'Siswa', icon: 'pi pi-users' },
    // { route: '/products', label: 'Products', icon: 'pi pi-list' },
    // { route: '/messages', label: 'Messages', icon: 'pi pi-inbox' }
]);
</script>

<style lang="scss" scoped></style>