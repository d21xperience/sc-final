<script setup>
import { ref, watch } from "vue";
import Toast from 'primevue/toast';
// import TheWelcome from '../components/Utama.vue'
import Navbar from '../components/Navbar.vue';
import { Form } from '@primevue/forms';
import { useToast } from 'primevue/usetoast';
import Carousel from "@/components/Carousel.vue";
const toast = useToast();


const show = () => {
  toast.add({ severity: 'info', summary: 'Info', detail: 'Message Content', life: 3000 });
};




// Button Search
const loading = ref(false);
const searchBy = ref()


const load = () => {
  loading.value = true;
  setTimeout(() => {
    loading.value = false;
  }, 2000);
};

// NPSN INPUT TEXT
import axios from "axios";
import { debounce } from "lodash";
const selectedSekolah = ref()
const filteredSekolah = ref([])
const sekolah = ref(true)
const baseUrl = '/api/getHasilPencarian'
// const instansi = ref([
//   "Kementerian Pendidikan", "Kementerian Agama"
// ])
// const jenjangPendidikan = ref([
//   ["SD", "SMP", "SMA", "SMK"],
//   ["MI", "MTs", "MA", "MAK"],
// ])
// const status = ref(["Negeri", "Swasta"])

// const selectedInstansi = ref(null)
// const selectedJenjangPendidikan = ref(null)
// const selectedStatus = ref(null)



const fetchData = async (e) => {
  console.log(e.query)
  if (e.query.length <= 3) return
  try {
    const res = await axios.get(baseUrl, {
      params: {
        keyword: e.query
      }
    })
    // console.log(res.data)
    // console.log(res.data.map((item) => item.nama_sekolah))
    filteredSekolah.value = res.data.map((item) => `${item.npsn} - ${item.nama_sekolah}`)
  } catch (error) {
    console.error(error)
  }

}

const cariSekolah = debounce(fetchData, 500)
watch(selectedSekolah, (newFilteredSekolah, oldFilteredsekolah) => {
  // console.log(newFilteredSekolah.length)
  if (newFilteredSekolah == null) {
    sekolah.value = true
  } else {
    sekolah.value = false
  }
})
// Carousel data
const products = ref()


// State untuk menyimpan query, hasil, dan error
const query = ref("");
const student = ref(null);
const error = ref("");
// Fungsi untuk mencari student
const searchStudent = async () => {
  try {
    // Reset hasil dan error
    student.value = null;
    error.value = "";

    // Request ke endpoint backend
    const response = await axios.get("http://localhost:8081/api/v1/student", {
      params: { query: query.value },
    });

    // Simpan data student
    student.value = response.data;
  } catch (err) {
    // Tangani error
    error.value = err.response?.data?.error || "An error occurred while fetching data.";
  }
}
const dlg = ref(false)

import Dialog from 'primevue/dialog';
import Captcha from "@/components/Captcha.vue";
const showRobot = () => {
  dlg.value = !dlg.value
}


const dataSiswa = ref({})


</script>

<template>
  <Navbar />
  <!-- <main>
    <TheWelcome />

  </main> -->
  <div class="bg-slate-100 text-black mt-20">
    <div class="">
      <!-- <img src="https://readymadeui.com/cardImg.webp" alt="Background Image" class="w-full" /> -->
      <!-- <img src="https://readymadeui.com/cardImg.webp" alt="Background Image"
        class="w-full h-full object-cover opacity-50" /> -->
    </div>

    <!-- <div class="relative max-w-screen-xl mx-auto px-8 py-2 z-10 ">
      <h1 class="text-4xl md:text-5xl font-bold leading-tight mb-3 text-center">Selamat Datang di Website Kami</h1>
      <p class="text-lg md:text-xl mb-3 text-center">Verifikasi Ijazah & Transkrip Nilai semakin mudah.</p>
    </div>
    <div>
      <form action="" class="lg:w-[500px] p-3 border  rounded-lg">
        <h2 class="font-bold text-xl text-center mb-3">FORMULIR VERIFIKASI</h2>
        <div class="mt-1">
          <label for="nama_sekolah" class="text-xs">Masukan Nama Sekolah atau NPSN</label>
          <div>
            <AutoComplete inputId="nama_sekolah" v-model="selectedSekolah" :suggestions="filteredSekolah"
              @complete="cariSekolah" inputClass="font-bold" forceSelection fluid />
          </div>
        </div>
        <div class="mt-1">
          <label for="no_ijazah" class="text-xs">No Ijazah</label>
          <div>
            <input type="text" name="no_ijazah" id="no_ijazah" class="rounded-md w-full p-2 text-black font-bold"
              :disabled="sekolah">
          </div>
        </div>
        <div class="mt-4 flex">
          <Button type="button" label="Verifikasi" icon="pi pi-search" :loading="loading" @click="load"
            class="w-full" />
        </div>
      </form>
    </div> -->

    <Toast />
    <!-- <button @click="show()">Show</button> -->

    <div v-show="dlg">

      <Captcha />
    </div>
    <!-- SECTION -1  -->
    <!-- <div class="bg-blue-100"> -->
    <section id="home">
      <div class="bg-slate-800 ">
        <div class="container mx-auto p-4">
          <div class=" p-8 flex flex-wrap">
            <div class="lg:w-1/2 w-full mb-12">
              <h1 class="font-sans text-3xl font-bold mb-4 text-slate-50">
                Verifikasi Ijazah Dasar & Menengah
              </h1>

              <div class="mb-4">
                <Form v-slot="$form" :initialValues="dataSiswa">
                  <!-- Input untuk query -->
                  <label for="cari-data" class="text-white">Masukan NISN atau NIK</label>
                  <input v-model="query" name="cari-data" placeholder="Enter NISN or NIK" type="text"
                    class="mt-1 p-2 w-full border rounded" />
                  <div class="flex items-center gap-2 my-3">
                    <!-- <button @click="showRobot" class="bg-white" type="button">&nbsp;&nbsp;</button> -->
                    <input type="checkbox" name="robot" id="robot" @change="showRobot">
                    <label for="robot" class=" text-white">Saya bukan robot</label>
                  </div>
                  <!-- <button @click="searchStudent" class="bg-white text-black border p-2 rounded-full">Search</button> -->
                  <button @click="searchStudent" :disabled="query.length === 0"
                    :class="{ 'opacity-25 cursor-not-allowed hover:bg-slate-500': query.length === 0 }"
                    class="w-full bg-slate-500 font-bold p-2 rounded shadow-md hover:bg-slate-400 hover:text-white"
                    type="button">
                    <i class="fas fa-search">
                    </i>
                    Cek Ijazah
                  </button>
                </Form>

                <!-- Tampilkan hasil jika ada -->
                <div v-if="student">
                  <div class=" mt-12">
                    <div class="flex flex-col gap-2 flex-wrap  bg-slate-100 rounded-lg hover:opacity-75">
                      <div class="flex cursor-pointer select-none align-middle text-left" tabindex="0" role="button">
                        <div class="min-w-[56px] inline-flex flex-shrink-0 "><svg
                            class="text-green-800 w-[3em] inline-block fill-green-800 shrink-0" focusable="false"
                            aria-hidden="true" viewBox="0 0 24 24" data-testid="CheckCircleIcon">
                            <path
                              d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z">
                            </path>
                          </svg></div>
                        <div class="flex-auto mt-1 mb-1"><span class="block text-green-900">Data telah
                            Terverifikasi</span>
                          <p class="text-sm">{{ student.nama }} | {{ student.asal_sekolah }} </p>
                        </div>
                        <div class="flex flex-col">
                          <div class="css-td0mdc">
                            <p class="MuiTypography-root MuiTypography-body1 css-1cf0n95"><span
                                class="font-token"></span>&nbsp;2 Okt 2022</p>
                          </div><button class="inline-flex items-center" tabindex="0" type="button">Ethereum<span
                              class="flex"><svg class="text-[20px]" focusable="false" aria-hidden="true"
                                viewBox="0 0 24 24" data-testid="OpenInNewIcon">
                                <path
                                  d="M19 19H5V5h7V3H5c-1.11 0-2 .9-2 2v14c0 1.1.89 2 2 2h14c1.1 0 2-.9 2-2v-7h-2v7zM14 3v2h3.59l-9.83 9.83 1.41 1.41L19 6.41V10h2V3h-7z">
                                </path>
                              </svg></span><span class="MuiTouchRipple-root css-w0pj6f"></span></button>
                        </div>
                        <span class="MuiTouchRipple-root css-w0pj6f"></span>
                      </div>
                    </div>
                  </div>


                </div>
                <!-- JIka ditemukan data ditampilkan dibawah -->

                <!-- Tampilkan error jika ada -->
                <div v-if="error" style="color: red;">
                  {{ error }}
                </div>
              </div>

            </div>
            <div class="lg:w-1/2 w-full pl-8 text-slate-50">
              <h2 class="text-xl font-bold mb-4">
                Pengantar
                <button>Read more...</button>
              </h2>
              <ol class="list-decimal">
                <li class="mb-2 hidden md:block">
                  Pengertian Nomor Induk Siswa Nasional adalah, kode pengenal identitas siswa yang bersifat unik,
                  standar dan berlaku sepanjang masa yang dapat membedakan satu siswa dengan siswa lainnya di seluruh
                  sekolah Indonesia dan Sekolah Indonesia di Luar Negeri; Nomor Induk Siswa Nasional (NISN) diberikan
                  kepada setiap peserta didik yang bersekolah di satuan pendidikan yang memiliki NPSN dan terdaftar di
                  Referensi Kemendikbud. Sistem pengelolaan NISN secara nasional oleh Pusat Data dan Statistik
                  Pendidikan dan Kebudayaan (PDSPK) Kemendikbud yang merupakan bagian dari program Dapodik (Data Pokok
                  Pendidikan) Kementerian Pendidikan dan Kebudayaan. Hasil dari proses pemberian kode identifikasi oleh
                  PDSPK ditampilkan secara terbuka dalam batasan tertentu melalui situs NISN
                  (http://nisn.data.kemdikbud.go.id./).
                </li>
                <li class="mb-2 hidden md:block">
                  Tujuan dan Manfaat
                  <ul class="list-disc ml-4">
                    <li>
                      Mengidentifikasi setiap individu siswa (peserta didik) di seluruh sekolah se-Indonesia secara
                      standar, konsisten dan berkesinambungan.
                    </li>
                    <li>
                      Sebagai pusat layanan sistem pengelolaan nomor induk siswa secara online bagi Unit unit Kerja di
                      Kemendikbud, Dinas Pendidikan Daerah hingga Sekolah yang bersifat standar, terpadu dan akuntabel
                      berbasis Teknologi Informasi dan Komunikasi terkini.
                    </li>
                    <li>
                      Sebagai sistem pendukung program Dapodik dalam pengembangan dan penerapan program perencanaan
                      pendidikan, statistik pendidikan dan program pendidikan lainnya baik di tingkat pusat, propinsi,
                      kota, kabupaten hingga sekolah, serta unit-unit kerja lainnya di Kemendikbud.
                    </li>
                  </ul>
                </li>
              </ol>
            </div>
          </div>
        </div>
      </div>
      <!-- </div> -->
    </section>



    <!-- END SECTION -1  -->

    <!-- <div class="bg-gray-800 p-4 rounded-lg shadow-lg w-80">
      <div class="flex justify-between items-center text-white mb-4">
        <span>
          Select in this order:
        </span>
        <div class="flex space-x-2">
          <i class="fas fa-star">
          </i>
          <i class="fas fa-calendar-alt">
          </i>
          <i class="fas fa-shopping-cart">
          </i>
        </div>
      </div>
      <div class="relative mb-4">
        <img alt="A serene night landscape with a tree reflected in water" class="rounded-lg w-full" height="180"
          src="https://storage.googleapis.com/a1aa/image/VUgCU0IKDhauLlnhaCBtVtjbxH8CZWdZGzA3Ky1xW7pOXe7JA.jpg"
          width="320">
        <div class="absolute top-2 left-2">
          <div class="bg-blue-500 text-white rounded-full w-8 h-8 flex items-center justify-center">
            <i class="fas fa-alarm-clock">
            </i>
          </div>
        </div>
        <div class="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2">
          <div class="bg-purple-500 text-white rounded-full w-8 h-8 flex items-center justify-center">
            <i class="fas fa-calendar-alt">
            </i>
          </div>
        </div>
        <div class="absolute bottom-2 right-2">
          <div class="bg-teal-500 text-white rounded-full w-8 h-8 flex items-center justify-center">
            <i class="fas fa-shopping-cart">
            </i>
          </div>
        </div>
      </div>
      <button class="bg-blue-600 text-white py-2 px-4 rounded-lg w-full">
        Apply
      </button>
      <div class="flex justify-between items-center text-gray-400 mt-4">
        <i class="fas fa-times">
        </i>
        <i class="fas fa-sync-alt">
        </i>
        <i class="fas fa-info-circle">
        </i>
      </div>
    </div> -->


    <!-- About start -->
    <!-- Section 2 -->
    <section id="about" class="lg:min-h-[350px] p-3 flex flex-wrap">
      <div class="container">
        <div class="flex flex-wrap">
          <div class="w-full px-4 mb-10 lg:w-1/2">
            <h4 class="font-bold uppercase mb-3">Tentang Kami</h4>
            <h2 class="font-bold text-3xl mb-5 max-w-xl">Secure and easy process</h2>
            <p class="font-medium text-base text-slate-500">Verifikasi Ijazah provides a reliable service for students,
              employers, and educational institutions in
              Bandung, ID to verify academic certificates quickly and securely. Our streamlined process ensures that you
              can
              trust
              the authenticity of educational credentials, helping you make informed decisions. Whether you're seeking
              employment or validating qualifications, we are here to assist you with utmost efficiency and
              professionalism.
            </p>
            <p class="font-medium text-base text-slate-500">Lorem ipsum dolor sit, amet consectetur adipisicing elit.
              Accusantium laborum consequuntur, omnis placeat dicta excepturi cumque repellat hic accusamus quis quia
              rem autem iste reiciendis aut voluptates impedit pariatur ratione illum doloribus fuga minus nostrum.
              Quisquam dolores reprehenderit nostrum a praesentium! Voluptas consequuntur esse nulla laborum velit id
              sequi repellat.</p>
          </div>
          <div class="w-full px-4 lg:w-1/2">
            <h3 class="font-semibold text-2xl lg:pt-10">Terhubung dengan Kami</h3>
            <p class="font-medium text-base text-slate-500 mb-6">Lorem ipsum dolor sit amet consectetur adipisicing
              elit. Adipisci aspernatur harum minima molestiae quas
              sunt. Maiores, voluptatem. Architecto, iste corrupti.</p>
          </div>
        </div>
      </div>
    </section>
    <!-- End of Section 2 -->

    <section class="mb-8 container">
      <h2 class="text-xl font-bold mb-4">
        Platform yang Digunakan
      </h2>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
        <div class="bg-white p-4 rounded-lg shadow-md">
          <img alt="Ethereum logo with blockchain elements" class="w-full h-auto mb-4 rounded-lg" height="200"
            src="https://storage.googleapis.com/a1aa/image/FJjJhxyfdhymaCz6sBLvGJKSVWtXxv6f7bfIy29W4DVBOcfPB.jpg"
            width="300">
          <h3 class="text-lg font-bold mb-2">
            Ethereum
          </h3>
          <p>
            Ethereum adalah platform blockchain terdesentralisasi yang memungkinkan smart contracts dan aplikasi
            terdesentralisasi (DApps) untuk berjalan tanpa downtime, penipuan, kontrol, atau gangguan dari pihak ketiga.
          </p>
        </div>
        <div class="bg-white p-4 rounded-lg shadow-md">
          <img alt="Hyperledger Fabric logo with network nodes" class="w-full h-auto mb-4 rounded-lg" height="200"
            src="https://storage.googleapis.com/a1aa/image/l6eS2IgpBPTtDiYT0Bhlo1A8Cew90IY7ja9i3Yickh1IHufnA.jpg"
            width="300">
          <h3 class="text-lg font-bold mb-2">
            Hyperledger Fabric
          </h3>
          <p>
            Hyperledger Fabric adalah kerangka kerja blockchain yang diizinkan, yang dirancang untuk aplikasi
            perusahaan. Ini menawarkan modularitas dan fleksibilitas yang unik untuk berbagai kasus penggunaan industri.
          </p>
        </div>
        <div class="bg-white p-4 rounded-lg shadow-md">
          <img alt="Quorum logo with secure transaction elements" class="w-full h-auto mb-4 rounded-lg" height="200"
            src="https://storage.googleapis.com/a1aa/image/1TIdlaY1OZodJBYXBSr9X2skHNLlpMzR1fsT83zCHWzhD3fTA.jpg"
            width="300">
          <h3 class="text-lg font-bold mb-2">
            Quorum
          </h3>
          <p>
            Quorum adalah platform blockchain perusahaan yang didasarkan pada Ethereum. Ini dirancang untuk digunakan
            dalam lingkungan yang memerlukan kecepatan tinggi dan throughput tinggi dari jaringan pribadi.
          </p>
        </div>
      </div>
    </section>







    <!-- About End -->
    <!-- Section 3 -->
    <section id="our-patner" class="md:min-h-[350px] p-3 flex flex-wrap bg-slate-700">
      <div class="container">
        <h2 class="p-3 text-white uppercase font-semibold">Our Patner</h2>
        <!-- <Carousel /> -->
      </div>

    </section>
    <!-- End ofSection 3 -->


    <!-- Section 4 -->
    <section class="container mb-8">
      <h2 class="text-xl font-bold mb-4">
        Berita Terbaru
      </h2>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
        <div class="bg-white p-4 rounded-lg shadow-md">
          <img alt="News image showing a blockchain conference" class="w-full h-auto mb-4 rounded-lg" height="200"
            src="https://storage.googleapis.com/a1aa/image/8c54PjjFN3buIBwaaoi4uEVUUgGNmnvrmT43v3cQOiowh7fJA.jpg"
            width="300">
          <h3 class="text-lg font-bold mb-2">
            Konferensi Blockchain 2023
          </h3>
          <p>
            Konferensi blockchain terbesar tahun ini akan diadakan di Jakarta, membahas berbagai topik termasuk
            verifikasi ijazah menggunakan blockchain.
          </p>
        </div>
        <div class="bg-white p-4 rounded-lg shadow-md">
          <img alt="News image showing a new blockchain partnership" class="w-full h-auto mb-4 rounded-lg" height="200"
            src="https://storage.googleapis.com/a1aa/image/JJgHI2WEvq4KO59545kadeeOFONLTwnTEes568r92rPKOcfPB.jpg"
            width="300">
          <h3 class="text-lg font-bold mb-2">
            Kemitraan Baru dengan Universitas
          </h3>
          <p>
            Kami dengan bangga mengumumkan kemitraan baru dengan beberapa universitas terkemuka untuk
            mengimplementasikan sistem verifikasi ijazah berbasis blockchain.
          </p>
        </div>
        <div class="bg-white p-4 rounded-lg shadow-md">
          <img alt="News image showing a blockchain workshop" class="w-full h-auto mb-4 rounded-lg" height="200"
            src="https://storage.googleapis.com/a1aa/image/0OwVEfhSMvTSXyHN5SSRbCbSuVUsQmCdWGNRee591o2OOcfPB.jpg"
            width="300">
          <h3 class="text-lg font-bold mb-2">
            Workshop Blockchain
          </h3>
          <p>
            Ikuti workshop kami untuk mempelajari lebih lanjut tentang bagaimana blockchain dapat digunakan untuk
            verifikasi ijazah dan aplikasi lainnya.
          </p>
        </div>
      </div>
    </section>




    <!-- <section class="lg:h-[350px] p-3 text-slate-100">
      <h2 class="text-4xl text-center font-semibold">Join Us!</h2>
      <div class="flex text-center lg:max-w-8xl mt-4">
        <div>
          <p class="text-2xl">Bersama Menjadi Bagian dari perkembangan teknologi</p>
          <p class="text-xl text-white"><span class="text-red-500">Indonesia</span> Bisa!</p>
          <p>Dengan blockchain, pastikan dokumen akademik Anda selalu terjaga keasliannya. Tidak ada lagi ijazah palsu,
            tidak ada lagi kekhawatiran akan kehilangan data.</p>
          <div class="mt-4">
            <button
              class="bg-blue-950 text-slate-100 p-3 rounded-full hover:opacity-75 hover:text-blue-400 w-2/12">Gabung</button>
          </div>
        </div>
        <div>
          <img src="../assets/eth.jpg" alt="" srcset="">
        </div>
      </div>
    </section> -->
    <!-- End of Section 3 -->

    <!-- Section 4 -->
    <!-- <section class="lg:h-[200px]">
      <div class="card">
        <Carousel :value="products" :numVisible="3" :numScroll="3" :responsiveOptions="responsiveOptions">
          <template #item="slotProps">
            <div class="border border-surface-200 dark:border-surface-700 rounded m-2  p-4">
              <div class="mb-4">
                <div class="relative mx-auto">
                  <img :src="'https://primefaces.org/cdn/primevue/images/product/' + slotProps.data.image"
                    :alt="slotProps.data.name" class="w-full rounded" />
                  <Tag :value="slotProps.data.inventoryStatus" :severity="getSeverity(slotProps.data.inventoryStatus)"
                    class="absolute" style="left:5px; top: 5px" />
                </div>
              </div>
              <div class="mb-4 font-medium">{{ slotProps.data.name }}</div>
              <div class="flex justify-between items-center">
                <div class="mt-0 font-semibold text-xl">${{ slotProps.data.price }}</div>
                <span>
                  <Button icon="pi pi-heart" severity="secondary" outlined />
                  <Button icon="pi pi-shopping-cart" class="ml-2" />
                </span>
              </div>
            </div>
          </template>
</Carousel>
</div>
</section> -->
    <!-- Endi of Section 4 -->
  </div>
  <!-- footer -->
  <footer id="contact" class="bg-slate-100 p-10  tracking-wide">
    <!-- <footer class="bg-gray-900 p-10 font-[sans-serif] tracking-wide"> -->
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-8">
      <div class="lg:flex lg:items-center">
        <!-- <a href="javascript:void(0)">
          <img src="https://readymadeui.com/readymadeui-light.svg" alt="logo" class="w-52" />
        </a> -->
      </div>

      <div class="lg:flex lg:items-center">
        <ul class="flex space-x-6">
          <li>
            <a href="javascript:void(0)">
              <svg xmlns="http://www.w3.org/2000/svg" class="fill-black hover:fill-white w-7 h-7" viewBox="0 0 24 24">
                <path fill-rule="evenodd"
                  d="M19 3H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h7v-7h-2v-3h2V8.5A3.5 3.5 0 0 1 15.5 5H18v3h-2a1 1 0 0 0-1 1v2h3v3h-3v7h4a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2z"
                  clip-rule="evenodd" />
              </svg>
            </a>
          </li>
          <li>
            <a href="javascript:void(0)">
              <svg xmlns="http://www.w3.org/2000/svg" class="fill-black hover:fill-white w-7 h-7" viewBox="0 0 24 24">
                <path fill-rule="evenodd"
                  d="M21 5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V5zm-2.5 8.2v5.3h-2.79v-4.93a1.4 1.4 0 0 0-1.4-1.4c-.77 0-1.39.63-1.39 1.4v4.93h-2.79v-8.37h2.79v1.11c.48-.78 1.47-1.3 2.32-1.3 1.8 0 3.26 1.46 3.26 3.26zM6.88 8.56a1.686 1.686 0 0 0 0-3.37 1.69 1.69 0 0 0-1.69 1.69c0 .93.76 1.68 1.69 1.68zm1.39 1.57v8.37H5.5v-8.37h2.77z"
                  clip-rule="evenodd" />
              </svg>
            </a>
          </li>
          <li>
            <a href="javascript:void(0)">
              <svg xmlns="http://www.w3.org/2000/svg" fill="none" class="fill-black hover:fill-white w-7 h-7"
                viewBox="0 0 24 24">
                <path
                  d="M22.92 6c-.77.35-1.6.58-2.46.69.88-.53 1.56-1.37 1.88-2.38-.83.5-1.75.85-2.72 1.05C18.83 4.5 17.72 4 16.46 4c-2.35 0-4.27 1.92-4.27 4.29 0 .34.04.67.11.98-3.56-.18-6.73-1.89-8.84-4.48-.37.63-.58 1.37-.58 2.15 0 1.49.75 2.81 1.91 3.56-.71 0-1.37-.2-1.95-.5v.03c0 2.08 1.48 3.82 3.44 4.21a4.22 4.22 0 0 1-1.93.07 4.28 4.28 0 0 0 4 2.98 8.521 8.521 0 0 1-5.33 1.84c-.34 0-.68-.02-1.02-.06C3.9 20.29 6.16 21 8.58 21c7.88 0 12.21-6.54 12.21-12.21 0-.19 0-.37-.01-.56.84-.6 1.56-1.36 2.14-2.23z" />
              </svg>
            </a>
          </li>
        </ul>
      </div>

      <div>
        <h4 class="text-lg font-semibold mb-6 text-black">Contact Us</h4>
        <ul class="space-y-2">
          <li>
            <a href="javascript:void(0)" class="text-black hover:text-black text-sm">Email</a>
          </li>
          <li>
            <a href="javascript:void(0)" class="text-black hover:text-black text-sm">Phone</a>
          </li>
          <li>
            <a href="javascript:void(0)" class="text-black hover:text-black text-sm">Address</a>
          </li>
        </ul>
      </div>

      <div>
        <h4 class="text-lg font-semibold mb-6 text-black">Information</h4>
        <ul class="space-y-2">
          <li>
            <a href="javascript:void(0)" class="text-black hover:text-black text-sm">Ethereum</a>
          </li>
          <!-- <li>
            <a href="javascript:void(0)" class="text-black hover:text-black text-sm">Terms &amp; Conditions</a>
          </li> -->
          <li>
            <a href="javascript:void(0)" class="text-black hover:text-black text-sm">Quorum</a>
          </li>
          <li>
            <a href="javascript:void(0)" class="text-black hover:text-black text-sm">Hyperledger</a>
          </li>
          <li>
            <a href="javascript:void(0)" class="text-black hover:text-black text-sm">Privacy Policy</a>
          </li>
        </ul>
      </div>
    </div>

    <!-- <p class='text-black text-sm mt-10'>© ReadymadeUI. All rights reserved.</p> -->
  </footer>

  <!-- End of footer -->

</template>
