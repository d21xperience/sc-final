<script setup>
import { ref, computed, watch } from 'vue';
import { useStore } from 'vuex';

import AutoComplete from 'primevue/autocomplete';
import IconField from 'primevue/iconfield';
import InputIcon from 'primevue/inputicon';
import InputText from 'primevue/inputtext';
import Password from 'primevue/password';
import { Form } from '@primevue/forms';
import { zodResolver } from '@primevue/forms/resolvers/zod';
import { z } from 'zod';


import Dialog from 'primevue/dialog';
import { FloatLabel, Message } from 'primevue';
import router from '@/router';



// -----------------------------------------

const showPesyaratan = ref(false)

// Autocomplete ---------------------------------------
const store = useStore();
const items = ref([])
const searchTerm = ref('');
const loading = ref(false);
const onSearch = async () => {
    // console.log("onSearch")
    if (searchTerm.value.trim().length < 4) {
        console.log("Please enter at least 4 characters");
        return;
    }

    loading.value = true;
    try {
        items.value = await store.dispatch('search/fetchResults', searchTerm.value);
        // console.log("Results:", items.value);
    } catch (error) {
        console.error("Error fetching results:", error);
    } finally {
        loading.value = false;
    }
    // console.log(searchTerm.value.length < 3)
    // loading.value = true;
    // await store.dispatch('search/fetchResults', searchTerm.value);
    // loading.value = false;
    // console.log('Results from Vuex:', store.state.search.results);
    // items.value = store.state.search.results.data
};

// ---------------------------------------



// Mmembuat Register form ---------------------
const registerForm = ref({
    email: "",
    password: "",
})


// watch(searchTerm, (oldVal, newVal) => {
//     if (newVal.length > 0) {
//         console.log("hello")
//     }
// })
const onFormSubmit = async (e) => {
    // console.log(e.values)
    // Mengirim data ke Server
    let dataReg = {
        email: e.values['email'],
        password: e.values['password'],
        sekolah: items.value.length > 0 ? toRaw(items.value[0]) : null,
    }

    console.log(dataReg)
    try {
        await store.dispatch('authService/registerAdmin', dataReg);
        // success.value = 'Admin registered successfully!';
    } catch (error) {
        console.error(error)
        // error.value = err.error || 'Registration failed';
    }
    // return { name, email, password, schoolName, register, error, success };
}
const resolver = ref(zodResolver(
    z.object({
        email: z.string().min(1, { message: 'Email harus diisi.' }).email({ message: 'Invalid email address.' }),
        password: z.string().min(1, { message: 'Password harus diisi.' }),
    })
));

// -----------------------------------

// ---Agrement
const agreement = ref(false);

// Cek sekolah
const statusSekolahTerdaftar = ref(false)
const infoError = ref(false)
const cekSekolah = () => {
    // infoError.value = !infoError.value
    statusSekolahTerdaftar.value = !statusSekolahTerdaftar.value
}

</script>



<template>
    <div class=" bg-slate-900 flex items-center h-screen p-4 text-slate-200">
        <div class="w-full max-w-md max-md:max-w-md mx-auto">
            <div class="bg-slate-700 grid gap-16 w-full sm:p-8 p-6 shadow-md rounded-md overflow-hidden">
                <div v-if="!statusSekolahTerdaftar">
                    <div class="mb-4">
                        <h3 class="text-2xl font-bold text-center">Form. Register</h3>
                        <div class="ml-4">
                            <ul class="text-sm list-disc">
                                <li>Formulir Register diperuntukan untuk <strong>Admin Sekolah.</strong></li>
                                <li>Admin Sekolah adalah <strong>Guru</strong> atau <strong>
                                        Tendik</strong> yang telah terdaftar di Dapodik.</li>
                                <li>Isi nama Sekolah atau NPSN pada kolom di bawah kemudian click tombol
                                    <strong>cek</strong>.
                                </li>
                                <li>Jika sekolah <strong>belum terdaftar</strong> maka akan dilanjutkan untuk mengsisi
                                    formulir register.
                                </li>
                            </ul>

                        </div>
                    </div>

                    <div class="my-6">
                        <div class="flex justify-between items-center space-x-1">
                            <div class="w-full">
                                <FloatLabel variant="on">
                                    <IconField>
                                        <AutoComplete name="sekolah" optionLabel="nama_sekolah" v-model="searchTerm"
                                            :suggestions="items" @complete="onSearch" fluid size="small" />
                                        <InputIcon class="pi pi-building-columns" />
                                    </IconField>
                                    <label for="sekolah">NPSN/Nama Sekolah</label>
                                </FloatLabel>
                            </div>
                            <div>
                                <button type="button" class="bg-blue-600 text-white text-sm p-2 rounded-lg"
                                    @click="cekSekolah">Cek</button>
                            </div>
                        </div>
                        <!-- <Message v-if="$form.sekolah?.invalid" severity="error" size="small" variant="simple">{{
                            $form.sekolah.error?.message }}</Message> -->
                        <!-- <Message size="small" severity="secondary" variant="simple">Diisi NPSN dan Nama Sekolah.
                        </Message> -->
                    </div>

                    <div class="flex justify-between mt-6">
                        <div class="flex justify-center flex-col">
                            <RouterLink to="/" class="xs:text-[10px] text-sm text-blue-600 hover:underline">Ke halaman
                                utama
                            </RouterLink>
                        </div>
                        <div class="flex justify-center flex-col">
                            <p class=" text-sm">Sudah punya akun ? <RouterLink :to="{ name: 'login' }"
                                    class="text-blue-600 font-semibold hover:underline ml-1">Login
                                    disini</RouterLink>
                            </p>
                        </div>
                    </div>
                </div>


                <Form v-slot="$form" :initialValues="registerForm" :resolver="resolver" @submit="onFormSubmit"
                    class="w-full justify-center" v-else>


                    <div class="my-4">
                        <input type="text" name="" id="" disabled value="SMKS Pasundan Jatinangor"
                            class="w-full rounded-lg p-2 text-slate-800 bg-slate-500">
                    </div>
                    <div class="my-4">
                        <!-- <label class="text-slate-700 text-sm block">Email Dapodik</label> -->
                        <FloatLabel variant="on">
                            <IconField>
                                <InputText name="email" type="email" size="small" class="w-full" autocomplete="false" />
                                <InputIcon class="pi pi-envelope" />
                            </IconField>
                            <label for="email">Email Dapodik</label>
                        </FloatLabel>
                        <Message v-if="$form.email?.invalid" severity="error" size="small" variant="simple">{{
                            $form.email.error?.message }}</Message>
                        <Message size="small" severity="secondary" variant="simple">Diisi email yg terdaftar di
                            Dapodik.</Message>
                    </div>
                    <div class="my-4">
                        <!-- <label class="text-slate-700 text-sm block">Password</label> -->
                        <FloatLabel variant="on">

                            <Password name="password" toggleMask size="small" fluid>
                                <template #header>
                                    <div class="font-semibold text-xm mb-4">Isi password</div>
                                </template>
                                <template #footer>
                                    <!-- <Divider /> -->
                                    <ul class="pl-2 ml-2 my-0 leading-normal text-sm w-full">
                                        <li>Terdiri dari huruf kecil</li>
                                        <li>Satu huruf besar</li>
                                        <li>Sekurangnya satu karakter angka</li>
                                        <li>Minimum 8 karakter</li>
                                    </ul>
                                </template>
                            </Password>
                            <label for="password">Password</label>

                        </FloatLabel>
                        <Message v-if="$form.password?.invalid" severity="error" size="small" variant="simple">{{
                            $form.password.error?.message }}</Message>
                        <Message size="small" severity="secondary" variant="simple">Buat Password baru.</Message>
                    </div>
                    <div class=" flex items-center">
                        <input id="aggrement" name="aggrement" type="checkbox" @click="agreement = !agreement"
                            class="h-4 w-4 shrink-0 text-blue-600 focus:ring-blue-500 border-gray-300 rounded-md" />
                        <label for="aggrement" class="ml-3 block text-sm">
                            Saya menyetujui <button type="button" @click="showPesyaratan = true"
                                class="text-blue-600 font-semibold hover:underline ml-1">Persyaratan & Aturan</button>
                        </label>
                    </div>
                    <!-- <p>{{ $form.email.error?.message }}</p> -->
                    <div class="flex justify-between mt-1 mb-2 space-x-2">
                        <!-- <div class="w-1/2"> -->
                        <button type="button" @click="cekSekolah"
                            class="block w-1/2 py-2.5 px-4 text-sm tracking-wider font-semibold rounded-md bg-green-600  text-white focus:outline-none">
                            Kembali
                        </button>

                        <!-- </div> -->
                        <button type="submit" :disabled="!agreement"
                            :class="agreement ? ['cursor-pointer', 'hover:bg-blue-700'] : ['cursor-not-allowed', 'bg-slate-500']"
                            class="block w-1/2  py-2.5 px-4 text-sm tracking-wider font-semibold rounded-md bg-blue-600  text-white focus:outline-none">
                            Buat Akun
                        </button>

                    </div>

                </Form>
            </div>
        </div>
    </div>





    <!-- Dialog start -->
    <!-- Dialog term -->
    <Dialog v-model:visible="showPesyaratan" header="Pernyataan dan aturan" :style="{ width: '50rem' }"
        :breakpoints="{ '1199px': '75vw', '575px': '90vw' }">
        <div>
            <h3>Pesyaratan dan Aturan</h3>
            <ol>
                <li>
                    <p>Lorem ipsum dolor sit, amet consectetur adipisicing elit. Aperiam, est necessitatibus! Nulla
                        rerum
                        animi in laudantium at cum!</p>
                </li>
                <li>
                    <p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Doloribus, libero asperiores quidem
                        corporis, reprehenderit, hic officia praesentium velit fugiat pariatur dolore!</p>
                </li>
                <li>
                    <p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Quo eos corrupti odio cum, similique
                        facere
                        at sed distinctio minus doloremque!</p>
                </li>
                <li>
                    <p>Lorem ipsum dolor sit, amet consectetur adipisicing elit. Illo molestias et ab explicabo ad
                        excepturi
                        corrupti eaque modi adipisci fugiat. Fugiat dolores officiis cum id quae eius impedit, dolore,
                        error
                        quis harum debitis fuga.</p>
                </li>
                <li>
                    <p>Lorem ipsum dolor, sit amet consectetur adipisicing elit. Nesciunt excepturi quos, tempore
                        incidunt
                        quisquam adipisci sint reiciendis, officiis illum possimus provident suscipit labore ullam nobis
                        quia earum veritatis expedita dolores necessitatibus totam doloribus? Dolorum nulla autem minima
                        illo nihil. Quae consequuntur tenetur alias, quis cum, obcaecati in omnis ab repellendus eum
                        corrupti error eos quam harum atque laborum quos optio fuga id praesentium sunt ea quo nemo
                        quisquam. Dolore porro quos fugit voluptate. Corporis, voluptatibus labore!</p>
                </li>
            </ol>
            <div class="flex items-center space-x-2 text-white my-3">
                <button class="bg-yellow-600 block p-2 rounded-lg w-64 hover:opacity-75">Batal</button>
                <button class="bg-blue-600 block p-2 rounded-lg w-64 hover:opacity-75">Terima</button>
            </div>
        </div>
    </Dialog>
    <!-- End of dialog term -->
    <!-- Dialog Info register -->
    <Dialog v-model:visible="infoError" header="Warning">
        <div class="text-slate-500">
            <p>Maaf sekolah yang Anda pilih telah terdaftar di sistem kami.</p>
            <p>Anda tidak bisa melanjutkan proses pendaftaran. </p>
            <p>Jika ini suatu kesalahan, silahkan hubungi admin kami</p>
        </div>
        <div class="flex my-4">
            <button class="bg-yellow-600 p-2 rounded-lg text-white" @click="router.push('/')">Ok, Terima kasih.</button>
        </div>
    </Dialog>


    <!-- End ofDialog Info register -->
    <!-- Dialog end -->
</template>