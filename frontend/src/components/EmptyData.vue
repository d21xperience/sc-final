<template>
    <div class="min-h-[350px]">
        <div class="text-center flex justify-center items-center space-x-2 ">
            <h3 class="text-xl">Data belum ditarik,
            </h3>
            <p>Silakan lakukan tarik data dengan mendownload aplikasi terlebih
                dahulu</p>
        </div>
        <div class="text-center mt-6">
            <p>Aplikasi hanya untuk operator aplikasi</p>
            <button type="button" class="bg-blue-200 p-2 rounded-lg hover:bg-blue-400"
                @click="router.push({ name: 'syncDapodik' })">Download disini</button>
        </div>
    </div>
</template>

<script setup>
import router from '@/router';

</script>

<style lang="scss" scoped></style>