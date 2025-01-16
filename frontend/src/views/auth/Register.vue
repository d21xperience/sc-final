<script setup>
import { ref, computed, watch, onMounted, reactive, toRaw } from 'vue';
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
// ============AUTOCOMPLETE STATIS=======================
import CountryService from "@/service/CountryService";
onMounted(() => {
    CountryService.getSekolah().then((data) => (countries.value = data));
});
const initialValues = ref({
    country: { name: '' }
});

const countries = ref();
const filteredCountries = ref();
// const toast = useToast();

const search = (event) => {
    setTimeout(() => {
        if (!event.query.trim().length) {
            filteredCountries.value = [...countries.value];
        } else {
            filteredCountries.value = countries.value.filter((country) =>
                country.nama_sekolah.toLowerCase().includes(event.query.toLowerCase())
            );
        }
    }, 250);
};
// ======================================================


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
        errorDialog.value = true
        errorInfo.value = "Sepertinya internet tidak terhubung"
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
const registerForm = reactive({
    username: '',
    email: '',
    password: '',
})

// Fungsi untuk menghapus spasi dan konversi ke string
const formatValues = obj => {
    return Object.fromEntries(
        Object.entries(obj).map(([key, value]) => [key, String(value).trim()])
    );
};
const errorDialog = ref(false)
const errorInfo = ref("")
const onFormSubmit = async ({ valid, values }) => {
    let dataReg = {
        user: {
            username: values.username,
            email: values.email,
            password: values.password,
            role: 'admin'
        },
        sekolah:
        {
            sekolah_id_enkrip: searchTerm.value.sekolah_id_enkrip,
            kecamatan: searchTerm.value.kecamatan,
            kabupaten: searchTerm.value.kabupaten,
            propinsi: searchTerm.value.propinsi,
            kode_kecamatan: searchTerm.value.kode_kecamatan,
            kode_kab: searchTerm.value.kode_kab,
            kode_prop: searchTerm.value.kode_prop,
            nama_sekolah: searchTerm.value.nama_sekolah,
            npsn: searchTerm.value.npsn,
            alamat_jalan: searchTerm.value.alamat_jalan,
            status: searchTerm.value.status
        }
    }
    dataReg.sekolah = formatValues(dataReg.sekolah);
    // console.log(dataReg.sekolah)
    // return
    if (valid) {
        try {
            const resp = await store.dispatch('authService/registerAdmin', dataReg);
            // Jika sukses, arahkan ke beranda
            if (resp.ok) {
                router.push({ name: 'admin' })
            }
            // success.value = 'Admin registered successfully!';
        } catch (error) {
            errorDialog.value = true
            errorInfo.value = error
            console.error(error)
            // error.value = err.error || 'Registration failed';
        }
        // return { name, email, password, schoolName, register, error, success };
    }

}
const resolver = ref(zodResolver(
    z.object({
        email: z.string().min(1, { message: 'Email harus diisi.' }).email({ message: 'Invalid email address.' }),
        password: z.string().min(1, { message: 'Password harus diisi.' }),
        username: z.string().min(1, { message: 'Username harus diisi.' }),
    })
));

// -----------------------------------

// ---Agrement
const agreement = ref(false);

// Cek sekolah
const npsn = ref(""); // Input untuk NPSN
const sekolah = ref(null); // Data sekolah dari API
const error = ref(""); // Menyimpan pesan error
const infoSekolah = ref(false)
const statusSekolahTerdaftar = ref(false)
const cekSekolah = async () => {
    npsn.value = searchTerm.value?.npsn
    // console.log(npsn./value)
    try {
        error.value = ""; // Reset error
        sekolah.value = null; // Reset data sekolah
        // Panggil fungsi ceknpsn dari Vuex storex
        const data = await store.dispatch("authService/ceknpsn", npsn.value);
        // console.log(data)
        if (data) {
            sekolah.value = data; // Tampilkan data sekolah
            infoSekolah.value = true
        } else if (data === null) {
            statusSekolahTerdaftar.value = true
        }
        else {
            statusSekolahTerdaftar.value = true
        }
    } catch (e) {
        error.value = "Terjadi kesalahan saat mengambil data sekolah.";
    }
}
const agrement = ref(false)
const terimaPersetujuan = () => {
    showPesyaratan.value = false
    agreement.value = true
}
const batalPersetujuan = () => {
    // showPesyaratan.value = true
    agreement.value = false
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
                                        <AutoComplete name="nama_sekolah" optionLabel="nama_sekolah"
                                            v-model="searchTerm" :suggestions="filteredCountries" @complete="search"
                                            fluid size="small" />
                                        <InputIcon class="pi pi-building-columns" />
                                    </IconField>
                                    <label for="sekolah">NPSN/Nama Sekolah</label>
                                </FloatLabel>
                            </div>
                            <div>
                                <button type="button" class="bg-blue-600 text-white text-sm p-2 rounded-lg"
                                    :disabled="searchTerm.length <= 0" @click="cekSekolah">Cek</button>
                            </div>
                        </div>
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
                <Form v-else v-slot="$form" :initialValues="registerForm" :resolver="resolver" @submit="onFormSubmit"
                    class="w-full justify-center">
                    <div class="mt-4 mb-6">
                        <label class="block w-full rounded-lg p-2 text-slate-800 bg-slate-500" for="nm-sekolah">{{
                            searchTerm?.nama_sekolah ?? '-' }}</label>
                        <!-- <input type="text" name="nama-sekolah" disabled :value="searchTerm?.nama_sekolah ?? '-'"
                            class="w-full rounded-lg p-2 text-slate-800 bg-slate-500"> -->
                    </div>
                    <div class="my-4">
                        <!-- <label class="text-slate-700 text-sm block">Email Dapodik</label> -->
                        <FloatLabel variant="on">
                            <IconField>
                                <InputText name="username" type="text" size="small" class="w-full"
                                    autocomplete="false" />
                                <InputIcon class="pi pi-user" />
                            </IconField>
                            <label for="email">Username</label>
                        </FloatLabel>
                        <Message v-if="$form.username?.invalid" severity="error" size="small" variant="simple">{{
                            $form.username.error?.message }}</Message>
                        <Message size="small" severity="secondary" variant="simple">Isi dengan username.</Message>
                    </div>
                    <div class="my-4">
                        <!-- <label class="text-slate-700 text-sm block">Email Dapodik</label> -->
                        <FloatLabel variant="on">
                            <IconField>
                                <InputText name="email" type="email" size="small" class="w-full" autocomplete="false" />
                                <InputIcon class="pi pi-envelope" />
                            </IconField>
                            <label for="email">Email</label>
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
                        <input id="agrement" name="agrement" type="checkbox" v-model="agrement"
                            @click="showPesyaratan = true"
                            class="h-4 w-4 shrink-0 text-blue-600 focus:ring-blue-500 border-gray-300 rounded-md" />
                        <label for="agrement" class="ml-3 block text-sm">
                            Saya menyetujui <button type="button" @click="showPesyaratan = true"
                                class="text-blue-600 font-semibold hover:underline ml-1">Persyaratan & Aturan</button>
                        </label>
                    </div>
                    <!-- <p>{{ $form.email.error?.message }}</p> -->
                    <div class="flex justify-between mt-1 mb-2 space-x-2">
                        <!-- <div class="w-1/2"> -->
                        <button type="button" @click="statusSekolahTerdaftar = false"
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
                <button class="bg-yellow-600 block p-2 rounded-lg w-64 hover:opacity-75"
                    @click="showPesyaratan = false">Batal</button>
                <button class="bg-blue-600 block p-2 rounded-lg w-64 hover:opacity-75"
                    @click="terimaPersetujuan">Terima</button>
            </div>
        </div>
    </Dialog>
    <!-- End of dialog term -->
    <!-- Dialog Info register -->
    <Dialog v-model:visible="infoSekolah" header="Warning">
        <div class="text-slate-500">
            <p>Maaf sekolah yang Anda pilih telah terdaftar di sistem kami.</p>
            <p>Anda tidak bisa melanjutkan proses pendaftaran. </p>
            <p>Jika ini suatu kesalahan, silahkan hubungi admin kami</p>
        </div>
        <div class="flex my-4">
            <button class="bg-yellow-600 p-2 rounded-lg text-white" @click="infoSekolah = false">Ok, Terima
                kasih.</button>
        </div>
    </Dialog>


    <Dialog v-model:visible="errorDialog" header="Warning">
        <div>
            {{ errorInfo }}
        </div>
    </Dialog>

    <!-- End ofDialog Info register -->
    <!-- Dialog end -->
</template>