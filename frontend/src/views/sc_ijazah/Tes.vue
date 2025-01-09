<template>
    <div class="py-2">
        <div v-for="tes in BCPlatform[0].accounts"
            class="flex items-center justify-between bg-blue-50 p-3 rounded-lg mb-4 cursor-pointer">
            <div class="flex items-center ">
                <img alt="Account icon" class="w-10 h-10 rounded-full mr-3" height="40"
                    src="https://storage.googleapis.com/a1aa/image/Xiq2NrsXBGrZHx49nxfacliC5jqcSqfWysyee9U6IhxvsfreE.jpg"
                    width="40" />
                <div>
                    <div class="font-semibold">
                        {{ tes.accountName }}
                    </div>
                    <div class="text-sm text-gray-500">
                        0x06664...DEc14
                    </div>
                </div>
            </div>
            <div class="text-right">
                <div class="font-semibold">
                    {{ tes.amount }} USD
                </div>
                <div class="flex items-center justify-end">
                    <i class="fas fa-coins text-yellow-500 mr-1">
                    </i>
                    <i class="fas fa-university text-blue-500">
                    </i>
                </div>
            </div>
            <button type="button" @click="tes">
                <i class="pi pi-ellipsis-v text-gray-500 ml-2 hover:opacity-50">
                </i>

            </button>
        </div>
        <!-- <button @click="dialogCreateAccount = true"
            class="w-full flex items-center justify-center py-2 border border-blue-500 text-blue-500 rounded-lg hover:bg-blue-500 hover:text-white">
            <i class="pi pi-plus mr-2">
            </i>
            Tambahkan akun
        </button> -->
    </div>
</template>

<script setup>
import { ref } from 'vue';

const BCPlatform = ref([
    {
        BCPlatformID: 1,
        name: "Ethereum",
        unit: "ETH",
        accounts: [
            {
                accountID: "",
                accountName: "Account 1",
                accountAddress: "39TqcPQbBh9bXvSkuxJC1Gi8GnGtAf41xCHevt85PgTp",
                amount: 300
            },
            {
                accountID: "",
                accountName: "Account 2",
                accountAddress: "9TqcPQbBh9bXvSkuxJC1Gi8GnGtAf41xCHevt85PgTpU",
                amount: 100
            }
        ]
    },
    {
        BCPlatformID: 2,
        name: "Quorum",
        seleted: false,
        unit: "ETH",
        amount: "10",
    },
    {
        BCPlatformID: 3,
        name: "Hyperledger Fabric",
        seleted: false,
        unit: "HF",
        amount: "6",
    },
])
</script>

<style lang="scss" scoped></style>