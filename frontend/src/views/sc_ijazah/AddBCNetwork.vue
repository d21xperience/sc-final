<template>
    <div class="flex justify-between border-b-2 pb-3">
        <h2 class="text-2xl font-bold">Jaringan</h2>
        <button class="bg-red-400 py-2 px-3 rounded-full hover:opacity-70 flex items-center gap-2" @click="bc"><i
                class="pi pi-times"></i> Tutup</button>
        <!-- <button class="rounded-full border w-10 p-2 hover:bg-slate-400 hover:text-white"
            @click="router.push({ name: 'setingBlockchain' })">X</button> -->
    </div>
    <div class="mt-4">
        <form action="" class="w-1/2 mx-auto border p-4">
            <div>
                <label for="tipe-jaringan" class="font-bold">Tipe Jaringan</label>
                <select name="tipe-jaringan" id="tip-jaringan" class="block border p-2 w-full  rounded-lg">
                    <option value="0" selected disabled><span>Pilih tipe jaringan</span></option>
                    <option value="public">Public</option>
                    <option value="private">Private</option>
                </select>
            </div>
            <div class="mt-1">
                <label for="nama-jaringan" class="font-bold">Nama Jaringan</label>
                <input id="nama-jaringan" type="text" class="block border  w-full  rounded-lg p-2">
            </div>
            <div class="mt-1">
                <label for="url-rpc" class="font-bold">URL RPC</label>
                <input id="url-rpc" type="text" class="block border  w-full  rounded-lg p-2">
            </div>
            <div class="mt-1">
                <label for="id-chain" class="font-bold">ID Chain</label>
                <input id="id-chain" type="text" class="block border  w-full  rounded-lg p-2">
            </div>
            <div class="mt-1">
                <label for="id-chain" class="font-bold">Simbol mata uang</label>
                <input id="id-chain" type="text" class="block border  w-full  rounded-lg p-2">
            </div>
            <div class="mt-1">
                <label for="id-chain" class="font-bold">ID Chain</label>
                <input id="id-chain" type="text" class="block border  w-full  rounded-lg p-2">
            </div>
            <div class="mt-1">
                <label for="id-chain"><span class="font-bold">URL block explorer</span> (opsional)</label>
                <input id="id-chain" type="text" class="block border  w-full  rounded-lg p-2">
            </div>
            <div class="mt-3 flex justify-between">
                <button @click="bc"
                    class="mr-2 p-2 border rounded-lg w-40 bg-yellow-800 text-white hover:bg-slate-400 hover:text-black">Batal</button>
                <button class="p-2 border rounded-lg w-40 bg-blue-500 text-white">Simpan &amp; gunakan</button>
            </div>
        </form>
    </div>
</template>

<script setup>
import router from '@/router';

const bc = () => {
    router.push({ name: 'setingBlockchain' })
}
</script>

<style lang="scss" scoped></style>