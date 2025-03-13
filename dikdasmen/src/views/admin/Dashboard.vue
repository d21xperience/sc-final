<template>
    <h1 class="text-2xl mb-2">Welcome back, <span class="font-bold ">{{ store.state.authService.user?.username
            }}</span> 🖐</h1>
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-6">
        <card-info title="Total blocks" date="22 - 29 May 2016" :value="totalBlocks" :percentage="33.87"
            @updateData="handleUpdate"></card-info>
        <!-- <div class="bg-white p-6 rounded-lg shadow-md">
            <div class="flex justify-between items-center mb-4">
                <h2 class="text-gray-600 text-lg font-semibold">
                    Total blocks
                </h2>
                <span class="text-gray-400 text-sm">
                    22 - 29 May 2016
                </span>
            </div>
            <div class="text-2xl font-bold text-gray-800 mb-2">
                10,215,845
            </div>
            <div class="flex items-center text-green-500">
                <i class="fas fa-arrow-up">
                </i>
                <span class="ml-1">
                    33.87%
                </span>
                <span class="text-gray-400 ml-2">
                    from last week
                </span>
            </div>
        </div> -->
        <div class="bg-white p-6 rounded-lg shadow-md">
            <div class="flex justify-between items-center mb-4">
                <h2 class="text-gray-600 text-lg font-semibold">
                    Total transaksi
                </h2>
                <span class="text-gray-400 text-sm">
                    22 - 29 May 2016
                </span>
            </div>
            <div class="text-2xl font-bold text-gray-800 mb-2">
                128,980
            </div>
            <div class="flex items-center text-red-500">
                <i class="fas fa-arrow-down">
                </i>
                <span class="ml-1">
                    33.87%
                </span>
                <span class="text-gray-400 ml-2">
                    from last week
                </span>
            </div>
        </div>
        <div class="bg-white p-6 rounded-lg shadow-md">
            <div class="flex justify-between items-center mb-4">
                <h2 class="text-gray-600 text-lg font-semibold">
                    Transaksi per hari
                </h2>
                <span class="text-gray-400 text-sm">
                    22 - 29 May 2016
                </span>
            </div>
            <div class="text-2xl font-bold text-gray-800 mb-2">
                2,905,897
            </div>
            <div class="flex items-center text-green-500">
                <i class="fas fa-arrow-up">
                </i>
                <span class="ml-1">
                    33.87%
                </span>
                <span class="text-gray-400 ml-2">
                    from last week
                </span>
            </div>
        </div>
    </div>
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-6">
        <card-info title="Jumlah Siswa" date="22 - 29 May 2016" :value="totalSiswa" :percentage="33.87"
            @updateData="handleUpdate" bgColor="bg-blue-100" routePath="/admin/data-siswa"></card-info>
        <card-info title="Jumlah Guru Aktif" date="22 - 29 May 2016" :value="totalGuru" :percentage="33.87"
            @updateData="handleUpdate" bgColor="bg-red-200" routePath="/dashboard"></card-info>
        <card-info title="Jumlah Kelas" date="22 - 29 May 2016" :value="totalKelas" :percentage="33.87"
            @updateData="handleUpdate" bgColor="bg-blue-100" routePath="/dashboard"></card-info>
    </div>
    <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div class="bg-white p-6 rounded-lg shadow-md">
            <div class="flex justify-between items-center mb-4">
                <h2 class="text-gray-600 text-lg font-semibold">
                    Report Statistics
                </h2>
                <!-- <div class="flex space-x-4">
                    <button class="text-gray-600 font-semibold">
                        Items Sold
                    </button>
                    <button class="text-gray-600 font-semibold">
                        Revenue
                    </button>
                </div> -->
            </div>
            <div class="text-gray-400 text-sm mb-4">
                Period: 22 - 29 May 2016
            </div>
            <div class="h-64">
                <Chart type="bar" :data="chartData" :options="chartOptions" />
            </div>
        </div>
        <div class="bg-white p-6 rounded-lg shadow-md">
            <div class="flex justify-between items-center mb-4">
                <h2 class="text-gray-600 text-lg font-semibold">
                    Latest Transactions
                </h2>
                <button class="text-gray-600 font-semibold" @click="router.push({ name: 'scIjazah' })">
                    see all transactions
                </button>
            </div>
            <div class="space-y-4">
                <!-- <div class="flex items-center justify-between">
                    <div class="flex items-center">
                        <img alt="Profile picture of Samantha William" class="w-10 h-10 rounded-full mr-4" height="40"
                            src="https://storage.googleapis.com/a1aa/image/vJm5ZVU5XYIfZafwxDE8BYteqO1c6t4Ib6tNbbDlQ1lggmpnA.jpg"
                            width="40" />
                        <div>
                            <div class="text-gray-800 font-semibold">
                                Samantha William
                            </div>
                            <div class="text-gray-400 text-sm">
                                on May 22, 2016
                            </div>
                        </div>
                    </div>
                    <div class="text-gray-800 font-semibold">
                        IDR 250,875
                    </div>
                    <div class="flex items-center">
                        <span class="w-3 h-3 bg-green-500 rounded-full mr-2">
                        </span>
                        <span class="text-gray-400 text-sm">
                            confirmed
                        </span>
                    </div>
                </div> -->
                <div class="flex items-center justify-between">
                    <div class="flex items-center">
                        <img alt="Profile picture of Kevin Anderson" class="w-10 h-10 rounded-full mr-4" height="40"
                            src="https://storage.googleapis.com/a1aa/image/gfBIQO4Djy3nSaTrIM2IyC3fJTAA2B9tKHstIDEYGGpSQz0TA.jpg"
                            width="40" />
                        <div>
                            <div class="text-gray-800 font-semibold">
                                Jhon Doe
                            </div>
                            <div class="text-gray-400 text-sm">
                                on May 22, 2016
                            </div>
                        </div>
                    </div>
                    <!-- <div class="text-gray-800 font-semibold">
                        IDR 350,620
                    </div> -->
                    <div class="flex items-center">
                        <span class="w-3 h-3 bg-yellow-500 rounded-full mr-2">
                        </span>
                        <span class="text-gray-400 text-sm">
                            waiting confirmation
                        </span>
                    </div>
                </div>
                <!--  <div class="flex items-center justify-between">
                    <div class="flex items-center">
                        <img alt="Profile picture of Angela Samantha" class="w-10 h-10 rounded-full mr-4" height="40"
                            src="https://storage.googleapis.com/a1aa/image/8DcswEf53PyWeEKre4TNdzm2StGMyHvHBgOPJeOcwipTBNTPB.jpg"
                            width="40" />
                        <div>
                            <div class="text-gray-800 font-semibold">
                                Angela Samantha
                            </div>
                            <div class="text-gray-400 text-sm">
                                on May 22, 2016
                            </div>
                        </div>
                    </div>
                    <div class="text-gray-800 font-semibold">
                        IDR 870,563
                    </div>
                    <div class="flex items-center">
                        <span class="w-3 h-3 bg-green-500 rounded-full mr-2">
                        </span>
                        <span class="text-gray-400 text-sm">
                            confirmed
                        </span>
                    </div>
                </div>
                <div class="flex items-center justify-between">
                    <div class="flex items-center">
                        <img alt="Profile picture of Michael Smith" class="w-10 h-10 rounded-full mr-4" height="40"
                            src="https://storage.googleapis.com/a1aa/image/CVwvONjeAITBCS4lLxRkxe5VbSD0EcyrAJN3GOGhOfrbgmpnA.jpg"
                            width="40" />
                        <div>
                            <div class="text-gray-800 font-semibold">
                                Michael Smith
                            </div>
                            <div class="text-gray-400 text-sm">
                                on May 22, 2016
                            </div>
                        </div>
                    </div>
                    <div class="text-gray-800 font-semibold">
                        IDR 653,975
                    </div>
                    <div class="flex items-center">
                        <span class="w-3 h-3 bg-red-500 rounded-full mr-2">
                        </span>
                        <span class="text-gray-400 text-sm">
                            payment expired
                        </span>
                    </div>
                </div>
                <div class="flex items-center justify-between">
                    <div class="flex items-center">
                        <img alt="Profile picture of Jonathan Sebastian" class="w-10 h-10 rounded-full mr-4" height="40"
                            src="https://storage.googleapis.com/a1aa/image/uuflnZm8SnTOFSYnrs4B3PS2LOMNiTAC2SceGJbueiHeANTPB.jpg"
                            width="40" />
                        <div>
                            <div class="text-gray-800 font-semibold">
                                Jonathan Sebastian
                            </div>
                            <div class="text-gray-400 text-sm">
                                on May 22, 2016
                            </div>
                        </div>
                    </div>
                    <div class="text-gray-800 font-semibold">
                        IDR 950,000
                    </div>
                    <div class="flex items-center">
                        <span class="w-3 h-3 bg-green-500 rounded-full mr-2">
                        </span>
                        <span class="text-gray-400 text-sm">
                            confirmed
                        </span>
                    </div>
                </div> -->
            </div>
        </div>
    </div>

</template>

<script setup>

import router from '@/router';
import Chart from 'primevue/chart';

import { ref, onMounted, computed } from "vue";
import CardInfo from '@/components/CardInfo.vue';

import { useStore } from "vuex";
const store = useStore();
onMounted(() => {
    fetchTabelTenant()
    chartData.value = setChartData();
    chartOptions.value = setChartOptions();
    fetchCountSiswa()
});

const chartData = ref();
const chartOptions = ref();

const setChartData = () => {
    return {
        labels: ['2007', '2008', '2009', '2010'],
        datasets: [
            {
                label: 'Data Lulusan',
                data: [240, 325, 302, 220],
                backgroundColor: ['rgba(249, 115, 22, 0.2)', 'rgba(6, 182, 212, 0.2)', 'rgb(107, 114, 128, 0.2)', 'rgba(139, 92, 246 0.2)'],
                borderColor: ['rgb(249, 115, 22)', 'rgb(6, 182, 212)', 'rgb(107, 114, 128)', 'rgb(139, 92, 246)'],
                borderWidth: 1
            }
        ]
    };
};
const setChartOptions = () => {
    const documentStyle = getComputedStyle(document.documentElement);
    const textColor = documentStyle.getPropertyValue('--p-text-color');
    const textColorSecondary = documentStyle.getPropertyValue('--p-text-muted-color');
    const surfaceBorder = documentStyle.getPropertyValue('--p-content-border-color');

    return {
        plugins: {
            legend: {
                labels: {
                    color: textColor
                }
            }
        },
        scales: {
            x: {
                ticks: {
                    color: textColorSecondary
                },
                grid: {
                    color: surfaceBorder
                }
            },
            y: {
                beginAtZero: true,
                ticks: {
                    color: textColorSecondary
                },
                grid: {
                    color: surfaceBorder
                }
            }
        }
    };
}

// ambil data untuk tampilan di dashboard beserta variabel yang dibutuhkan
const totalBlocks = ref(10215845);
const totalSiswa = ref(800);
const totalGuru = ref(50);
const totalKelas = ref(28);
const fetchCountSiswa = async () => {
    // console.log(schemaname.value)
    let payload = {
        schemaname: schemaname.value,
        semester_id: "20241"
    }
    const resp = await store.dispatch("sekolahService/fetchSiswaCount", payload)
    // console.log(resp)
}
// Tenant Table
const schemaname = computed(() => {
    let res = store.getters["sekolahService/getTabeltenant"]
    return res?.schemaName
})
// Perikasa dan perbaiki kode sesuai dengan praktek terbaik.
const fetchTabelTenant = async () => {
    let sekolahId = store.getters["authService/getUserProfile"]
    const te = await store.dispatch("sekolahService/fetchTabeltenant", sekolahId?.sekolahId)
}
const handleUpdate = (newValue) => {
    totalBlocks.value = newValue;
};
</script>
