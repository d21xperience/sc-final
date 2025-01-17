<script setup>
import { ref } from 'vue';
import { useStore } from 'vuex';
import { Form } from '@primevue/forms';
import { zodResolver } from '@primevue/forms/resolvers/zod';
import { z } from 'zod';
import IconField from 'primevue/iconfield';
import InputIcon from 'primevue/inputicon';
import InputText from 'primevue/inputtext';
import Password from 'primevue/password';

import { FloatLabel, Message, Select } from 'primevue';
import router from '@/router';


const store = useStore()
// membuat login form
const credential = ref({
    username: '',
    password: '',
    rememberMe: false
})

const resolver = ref(zodResolver(
    z.object({
        username: z.string().min(1, { message: 'Username harus diisi.' }),
        password: z.string().min(1, { message: 'Password harus diisi.' }),
    })
));

const onFormSubmit = async ({ valid, values }) => {
    if (valid) {
        const success = await store.dispatch('authService/login', { username: values.username, password: values.password });
        if (!success) {
            // alert(success.message)
            // error.value = 'Invalid login credentials';
            alert("Invalid login credentials")
        } else {
            router.push({ name: 'admin' })
        }

    }
}

</script>



<template>
    <div class=" bg-slate-900 flex items-center h-screen p-4 text-slate-200">
        <div class="w-full max-w-md max-md:max-w-md mx-auto">
            <div class="bg-slate-700 grid gap-8 w-full sm:p-8 p-6 shadow-md rounded-md overflow-hidden">
                <Form v-slot="$form" :initialValues="credential" :resolver="resolver" @submit="onFormSubmit"
                    class="w-full justify-center">
                    <div class="mb-4">
                        <h3 class="text-2xl font-bold text-center">Form. Login</h3>
                    </div>
                    <div class="my-6">
                        <FloatLabel variant="on">
                            <IconField>
                                <InputText name="username" type="text" size="small" class="w-full" autocomplete="on" />
                                <InputIcon class="pi pi-user" />
                            </IconField>
                            <label for="email">Username</label>
                        </FloatLabel>
                        <Message v-if="$form.username?.invalid" severity="error" size="small" variant="simple">{{
                            $form.username.error?.message }}</Message>
                    </div>
                    <div class="mt-4">
                        <FloatLabel variant="on">
                            <Password name="password" v-model="credential.password" toggleMask size="small" fluid
                                :feedback="false">
                            </Password>
                            <label for="password">Password</label>

                        </FloatLabel>
                        <Message v-if="$form.password?.invalid" severity="error" size="small" variant="simple">{{
                            $form.password.error?.message }}</Message>
                    </div>
                    <div class="flex flex-wrap items-center justify-between gap-4 mt-1 mb-6">
                        <div class="flex items-center">
                            <input id="remember-me" name="remember-me" type="checkbox"
                                class="h-4 w-4 shrink-0 border-gray-300 rounded" />
                            <label for="remember-me" class="ml-3 block text-sm ">
                                Remember me
                            </label>
                        </div>
                        <div>
                            <a href="jajvascript:void(0);" class="text-blue-600 text-sm font-semibold hover:underline">
                                Lupa Password?
                            </a>
                        </div>
                    </div>
                    <div class="!mt-2">
                        <button type="submit"
                            class="w-full py-2.5 px-4 text-sm tracking-wider font-semibold rounded-md bg-blue-600 hover:bg-blue-700 text-white focus:outline-none">
                            Log In
                        </button>
                    </div>
                </Form>
                <!-- login SSO -->
                <!-- <p class="text-center text-sm">---------------- Atau ----------------</p> -->
                <hr class="!my-1 border-gray-400" />
                <div class="space-x-8 flex justify-center">
                    <button type="button" class="border-none outline-none">
                        <svg xmlns="http://www.w3.org/2000/svg" width="30px" class="inline" viewBox="0 0 512 512">
                            <path fill="#fbbd00"
                                d="M120 256c0-25.367 6.989-49.13 19.131-69.477v-86.308H52.823C18.568 144.703 0 198.922 0 256s18.568 111.297 52.823 155.785h86.308v-86.308C126.989 305.13 120 281.367 120 256z"
                                data-original="#fbbd00" />
                            <path fill="#0f9d58"
                                d="m256 392-60 60 60 60c57.079 0 111.297-18.568 155.785-52.823v-86.216h-86.216C305.044 385.147 281.181 392 256 392z"
                                data-original="#0f9d58" />
                            <path fill="#31aa52"
                                d="m139.131 325.477-86.308 86.308a260.085 260.085 0 0 0 22.158 25.235C123.333 485.371 187.62 512 256 512V392c-49.624 0-93.117-26.72-116.869-66.523z"
                                data-original="#31aa52" />
                            <path fill="#3c79e6"
                                d="M512 256a258.24 258.24 0 0 0-4.192-46.377l-2.251-12.299H256v120h121.452a135.385 135.385 0 0 1-51.884 55.638l86.216 86.216a260.085 260.085 0 0 0 25.235-22.158C485.371 388.667 512 324.38 512 256z"
                                data-original="#3c79e6" />
                            <path fill="#cf2d48"
                                d="m352.167 159.833 10.606 10.606 84.853-84.852-10.606-10.606C388.668 26.629 324.381 0 256 0l-60 60 60 60c36.326 0 70.479 14.146 96.167 39.833z"
                                data-original="#cf2d48" />
                            <path fill="#eb4132"
                                d="M256 120V0C187.62 0 123.333 26.629 74.98 74.98a259.849 259.849 0 0 0-22.158 25.235l86.308 86.308C162.883 146.72 206.376 120 256 120z"
                                data-original="#eb4132" />
                        </svg>
                    </button>

                    <button type="button" class="border-none outline-none">
                        <svg xmlns="http://www.w3.org/2000/svg" width="30px" fill="#007bff"
                            viewBox="0 0 167.657 167.657">
                            <path
                                d="M83.829.349C37.532.349 0 37.881 0 84.178c0 41.523 30.222 75.911 69.848 82.57v-65.081H49.626v-23.42h20.222V60.978c0-20.037 12.238-30.956 30.115-30.956 8.562 0 15.92.638 18.056.919v20.944l-12.399.006c-9.72 0-11.594 4.618-11.594 11.397v14.947h23.193l-3.025 23.42H94.026v65.653c41.476-5.048 73.631-40.312 73.631-83.154 0-46.273-37.532-83.805-83.828-83.805z"
                                data-original="#010002"></path>
                        </svg>
                    </button>
                </div>






                <div class="flex flex-wrap items-center justify-between gap-4">
                    <!-- <div class="flex items-center">
                        <input id="remember-me" name="remember-me" type="checkbox"
                            class="h-4 w-4 shrink-0 border-gray-300 rounded" />
                        <label for="remember-me" class="ml-3 block text-sm text-gray-800">
                            Remember me
                        </label>
                    </div> -->
                    <div>
                        <RouterLink to="/" class=" text-sm font-semibold hover:underline">
                            Kembali ke halaman utama
                        </RouterLink>
                    </div>
                </div>

            </div>
        </div>
    </div>
</template>