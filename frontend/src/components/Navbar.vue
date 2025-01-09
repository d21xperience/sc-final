<script setup>
import store from '@/store';
import { ref, onMounted, onUnmounted } from 'vue';
const cek = ref(false)

// Emit event ke parent
// const emit = defineEmits(['scrollStateChange']);

// // State untuk posisi navbar
// const isRelative = ref(false);

// // Fungsi untuk memeriksa scroll
// const handleScroll = () => {
//     isRelative.value = window.scrollY > 100; // Contoh: jika scroll lebih dari 100px
//     emit('scrollStateChange', isRelative.value); // Emit state ke parent
// };

// // Tambahkan dan hapus event listener
// onMounted(() => {
//     window.addEventListener('scroll', handleScroll);
// });
// onUnmounted(() => {
//     window.removeEventListener('scroll', handleScroll);
// });

const isAuthenticated = store.getters["authService/isAuthenticated"]
</script>


<template>
    <header class='bg-transparent navbar-fixed top-0 left-0 w-full flex items-center z-10'>
        <div class='container'>
            <div class="flex items-center justify-between relative">
                <div class="px-4">
                    <a href="#home" class="font-bold text-lg block py-6">AKA</a>
                </div>
                <div class='flex items-center px-4'>
                    <button id="hamburger" name="hamburger" type="button" class="absolute right-4 lg:hidden"
                        @click="cek = !cek" :class="{ 'hamburger-active': cek }">
                        <!-- tes -->
                        <span class="hamburger-line origin-top-left transition duration-300"></span>
                        <span class="hamburger-line transition duration-300"></span>
                        <span class="hamburger-line origin-bottom-left transition duration-300"></span>
                    </button>
                    <nav class="absolute p-5 bg-white shadow-lg rounded-lg max-w-[250px] w-full right-4 top-full lg:block lg:static lg:bg-transparent lg:max-w-full"
                        :class="{ 'hidden': !cek }">
                        <ul class="block lg:flex">
                            <li class="group">
                                <a href="#home" class='text-base py-2 mx-8 flex group-hover:text-slate-500'>
                                    Beranda
                                </a>
                            </li>
                            <li class="group">
                                <a href="#about" class='text-base py-2 mx-8 flex group-hover:text-slate-500'>
                                    Tentang Kami
                                </a>
                            </li>
                            <li class="group">
                                <a href="#our-patner" class='text-base py-2 mx-8 flex group-hover:text-slate-500'>
                                    Our Partner
                                </a>
                            </li>
                            <li class="group">
                                <a href="#contact" class='text-base py-2 mx-8 flex group-hover:text-slate-500'>
                                    Kontak
                                </a>
                            </li>
                            <template v-if="isAuthenticated">
                                <li class="group">
                                    <RouterLink :to="{ name: 'admin' }"
                                        class='text-base py-2 mx-8 flex group-hover:text-slate-500 items-center'>
                                        <i class="pi pi-microsoft"></i>&nbsp; Dashboard
                                    </RouterLink>
                                </li>
                            </template>
                            <template v-else>
                                <li class="group">
                                    <RouterLink :to="{ name: 'login' }"
                                        class='text-base py-2 mx-8 flex group-hover:text-slate-500'>
                                        Login
                                    </RouterLink>
                                </li>
                                <li class="group">
                                    <RouterLink :to="{ name: 'register' }"
                                        class='text-base py-2 mx-8 flex group-hover:text-slate-500'>
                                        Sign
                                        up</RouterLink>
                                </li>
                            </template>
                            <li></li>
                        </ul>
                    </nav>
                </div>
            </div>
        </div>
    </header>

    <!-- <p>Scroll Posisi: {{ scrollY }}</p>
    <div style="height: 2000px;">Scroll ke bawah untuk melihat efeknya.</div> -->
</template>

<style></style>