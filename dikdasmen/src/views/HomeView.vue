<script setup>
import { ref, watch, onMounted } from "vue";
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
const imgRef = ref(null);
const isVisible = ref(false);

onMounted(() => {
  const observer = new IntersectionObserver(
    ([entry]) => {
      if (entry.isIntersecting) {
        isVisible.value = true;
        observer.disconnect(); // Hentikan observasi setelah gambar dimuat
      }
    },
    { threshold: 0.1 } // Gambar mulai dimuat ketika 10% terlihat
  );

  if (imgRef.value) {
    observer.observe(imgRef.value);
  }
});

</script>

<template>
  <Navbar />
  <!-- <main>
    <TheWelcome />

  </main> -->
  <div class="bg-slate-100 text-black mt-20">
    <div class="">
      <!-- <img loading="lazy" src="https://readymadeui.com/cardImg.webp" alt="Background Image" class="w-full" /> -->
      <!-- <img loading="lazy" src="https://readymadeui.com/cardImg.webp" alt="Background Image"
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

      <!-- <Captcha /> -->
    </div>
    <!-- SECTION -1  -->
    <!-- <div class="bg-blue-100"> -->
    <section id="home">
      <div class="bg-slate-800 h-screen">
        <div style="overflow-y: hidden;">
          <div>
            <div style="position: relative;">
              <div class="banner-container-home">
                <div class="banner-container-custom">
                  <div class="banner-home-overlay"></div>
                  <div class="carousel-custom carousel slide carousel-fade">
                    <div class="carousel-inner">
                      <div class="carousel-item">
                        <div
                          style="background-image: url(&quot;/static/media/banner1.83d2afa6.jpg&quot;); background-repeat: no-repeat; background-size: cover; background-position: center center; width: 100%; height: 100%; position: relative;">
                          <div class="banner-home-overlay"></div>
                        </div>
                      </div>
                      <div class="carousel-item">
                        <div
                          style="background-image: url(&quot;/static/media/banner4.a119c62f.jpg&quot;); background-repeat: no-repeat; background-size: cover; background-position: center center; width: 100%; height: 100%; position: relative;">
                          <div class="banner-home-overlay"></div>
                        </div>
                      </div>
                      <div class="active carousel-item">
                        <div
                          style="background-image: url(&quot;/static/media/banner1.83d2afa6.jpg&quot;); background-repeat: no-repeat; background-size: cover; background-position: center center; width: 100%; height: 100%; position: relative;">
                          <div class="banner-home-overlay"></div>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
              <div class="banner-content">
                <div class="container">
                  <div class="banner-title">Dinas Pendidikan Jawa Barat</div>
                  <h3 class="banner-subtitle">Wilujeng Enjing #WargiDisdikJabar<span class="span-current-text"
                      aria-hidden="true"></span></h3>
                  <h3 class="banner-subtitle">#TERDIDIKTERBAIK</h3>
                  <div class="banner-search-content">
                    <div class="">Temukan informasi publik terkini dari Dinas Pendidikan Jawa Barat</div>
                    <div class="banner-search-input">
                      <div class="container" style="width: 100%;">
                        <form>
                          <div class="input-with-button d-flex align-items-center" style="width: 100%;">
                            <div class="input-container" style="flex: 1 1 0%;"><input class="form-control input-custom"
                                type="text" name="query_search" placeholder="Cari informasi" autocomplete="off"></div>
                            <button type="submit" class="btn btn-search-home-custom"><svg aria-hidden="true"
                                focusable="false" data-prefix="fas" data-icon="search"
                                class="svg-inline--fa fa-search fa-w-16 " role="img" xmlns="http://www.w3.org/2000/svg"
                                viewBox="0 0 512 512">
                                <path fill="currentColor"
                                  d="M505 442.7L405.3 343c-4.5-4.5-10.6-7-17-7H372c27.6-35.3 44-79.7 44-128C416 93.1 322.9 0 208 0S0 93.1 0 208s93.1 208 208 208c48.3 0 92.7-16.4 128-44v16.3c0 6.4 2.5 12.5 7 17l99.7 99.7c9.4 9.4 24.6 9.4 33.9 0l28.3-28.3c9.4-9.4 9.4-24.6.1-34zM208 336c-70.7 0-128-57.2-128-128 0-70.7 57.2-128 128-128 70.7 0 128 57.2 128 128 0 70.7-57.2 128-128 128z">
                                </path>
                              </svg></button>
                          </div>
                        </form>
                      </div>
                    </div>
                    <div class="banner-search-popular">Pencarian Terpopuler</div>
                    <div class="banner-search-most-popular">
                      <div class="slider-content">
                        <div class="slider-content-custom">
                          <div class="slider-container-custom"><button class="btn btn-most-popular">PPDB<svg
                                aria-hidden="true" focusable="false" data-prefix="fas" data-icon="external-link-alt"
                                class="svg-inline--fa fa-external-link-alt fa-w-16 icon" role="img"
                                xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">
                                <path fill="currentColor"
                                  d="M432,320H400a16,16,0,0,0-16,16V448H64V128H208a16,16,0,0,0,16-16V80a16,16,0,0,0-16-16H48A48,48,0,0,0,0,112V464a48,48,0,0,0,48,48H400a48,48,0,0,0,48-48V336A16,16,0,0,0,432,320ZM488,0h-128c-21.37,0-32.05,25.91-17,41l35.73,35.73L135,320.37a24,24,0,0,0,0,34L157.67,377a24,24,0,0,0,34,0L435.28,133.32,471,169c15,15,41,4.5,41-17V24A24,24,0,0,0,488,0Z">
                                </path>
                              </svg></button><button class="btn btn-most-popular">Beasiswa<svg aria-hidden="true"
                                focusable="false" data-prefix="fas" data-icon="external-link-alt"
                                class="svg-inline--fa fa-external-link-alt fa-w-16 icon" role="img"
                                xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">
                                <path fill="currentColor"
                                  d="M432,320H400a16,16,0,0,0-16,16V448H64V128H208a16,16,0,0,0,16-16V80a16,16,0,0,0-16-16H48A48,48,0,0,0,0,112V464a48,48,0,0,0,48,48H400a48,48,0,0,0,48-48V336A16,16,0,0,0,432,320ZM488,0h-128c-21.37,0-32.05,25.91-17,41l35.73,35.73L135,320.37a24,24,0,0,0,0,34L157.67,377a24,24,0,0,0,34,0L435.28,133.32,471,169c15,15,41,4.5,41-17V24A24,24,0,0,0,488,0Z">
                                </path>
                              </svg></button></div>
                        </div>
                      </div><br>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="container">
            <div class="card card-news-home">
              <div class="card-body-news-home">
                <div class="row">
                  <div class="col-lg-6 col-md-6 col-sm-12 d-flex flex-row justify-content-start align-items-center">
                    <div class="title-news-home">Berita Terkini</div>
                  </div>
                  <div class="col-lg-6 col-md col-sm-12 d-flex flex-row justify-content-end align-items-center"><button
                      class="btn btn-selengkapnya">Lihat Semua Berita <span class="icon" style="margin-left: 5px;"><svg
                          aria-hidden="true" focusable="false" data-prefix="fas" data-icon="external-link-alt"
                          class="svg-inline--fa fa-external-link-alt fa-w-16 " role="img"
                          xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">
                          <path fill="currentColor"
                            d="M432,320H400a16,16,0,0,0-16,16V448H64V128H208a16,16,0,0,0,16-16V80a16,16,0,0,0-16-16H48A48,48,0,0,0,0,112V464a48,48,0,0,0,48,48H400a48,48,0,0,0,48-48V336A16,16,0,0,0,432,320ZM488,0h-128c-21.37,0-32.05,25.91-17,41l35.73,35.73L135,320.37a24,24,0,0,0,0,34L157.67,377a24,24,0,0,0,34,0L435.28,133.32,471,169c15,15,41,4.5,41-17V24A24,24,0,0,0,488,0Z">
                          </path>
                        </svg></span></button></div>
                  <div class="col-lg-8 col-md-8 col-sm-12 mt-3">
                    <div class="carousel slide carousel-fade">
                      <ol class="carousel-indicators">
                        <li></li>
                        <li></li>
                        <li></li>
                        <li></li>
                        <li class="active"></li>
                        <li></li>
                      </ol>
                      <div class="carousel-inner">
                        <div class="carousel-item">
                          <div class="c-body-berita-terkini"
                            style="background-image: url(&quot;https://api-webdisdik.jabarprov.go.id/storage/disdik_web/assets/news/img/attachment-1741685344746.jpg&quot;);">
                            <div class="c-head-berita-terkini">
                              <h4>Sengketa Lahan SMAN 1 Bandung, Wamendikdasmen: Semoga Putusan Sesuai Fakta</h4>
                              <div class="text__details">
                                <p class="text__detail"><span class="icons icons-color"><svg aria-hidden="true"
                                      focusable="false" data-prefix="fas" data-icon="calendar-day"
                                      class="svg-inline--fa fa-calendar-day fa-w-14 " role="img"
                                      xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512">
                                      <path fill="currentColor"
                                        d="M0 464c0 26.5 21.5 48 48 48h352c26.5 0 48-21.5 48-48V192H0v272zm64-192c0-8.8 7.2-16 16-16h96c8.8 0 16 7.2 16 16v96c0 8.8-7.2 16-16 16H80c-8.8 0-16-7.2-16-16v-96zM400 64h-48V16c0-8.8-7.2-16-16-16h-32c-8.8 0-16 7.2-16 16v48H160V16c0-8.8-7.2-16-16-16h-32c-8.8 0-16 7.2-16 16v48H48C21.5 64 0 85.5 0 112v48h448v-48c0-26.5-21.5-48-48-48z">
                                      </path>
                                    </svg><span class="icons icons-red"></span> </span>Mar 11, 2025</p>
                                <p class="text__detail__tag"><span class="icons icons-color"><svg aria-hidden="true"
                                      focusable="false" data-prefix="fas" data-icon="clock"
                                      class="svg-inline--fa fa-clock fa-w-16 " role="img"
                                      xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">
                                      <path fill="currentColor"
                                        d="M256,8C119,8,8,119,8,256S119,504,256,504,504,393,504,256,393,8,256,8Zm92.49,313h0l-20,25a16,16,0,0,1-22.49,2.5h0l-67-49.72a40,40,0,0,1-15-31.23V112a16,16,0,0,1,16-16h32a16,16,0,0,1,16,16V256l58,42.5A16,16,0,0,1,348.49,321Z">
                                      </path>
                                    </svg> </span>16:29:04</p>
                                <p class="text__detail"><span class="icons icons-color"><svg aria-hidden="true"
                                      focusable="false" data-prefix="fas" data-icon="pencil-alt"
                                      class="svg-inline--fa fa-pencil-alt fa-w-16 " role="img"
                                      xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">
                                      <path fill="currentColor"
                                        d="M497.9 142.1l-46.1 46.1c-4.7 4.7-12.3 4.7-17 0l-111-111c-4.7-4.7-4.7-12.3 0-17l46.1-46.1c18.7-18.7 49.1-18.7 67.9 0l60.1 60.1c18.8 18.7 18.8 49.1 0 67.9zM284.2 99.8L21.6 362.4.4 483.9c-2.9 16.4 11.4 30.6 27.8 27.8l121.5-21.3 262.6-262.6c4.7-4.7 4.7-12.3 0-17l-111-111c-4.8-4.7-12.4-4.7-17.1 0zM124.1 339.9c-5.5-5.5-5.5-14.3 0-19.8l154-154c5.5-5.5 14.3-5.5 19.8 0s5.5 14.3 0 19.8l-154 154c-5.5 5.5-14.3 5.5-19.8 0zM88 424h48v36.3l-64.5 11.3-31.1-31.1L51.7 376H88v48z">
                                      </path>
                                    </svg> </span>Nizar</p>
                              </div>
                              <div class="mb-2 title-container"><button class="btn btn-custom-home btn-md">Lihat
                                  Selengkapnya<span class="icon" style="margin-left: 5px;"><svg aria-hidden="true"
                                      focusable="false" data-prefix="fas" data-icon="external-link-alt"
                                      class="svg-inline--fa fa-external-link-alt fa-w-16 " role="img"
                                      xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">
                                      <path fill="currentColor"
                                        d="M432,320H400a16,16,0,0,0-16,16V448H64V128H208a16,16,0,0,0,16-16V80a16,16,0,0,0-16-16H48A48,48,0,0,0,0,112V464a48,48,0,0,0,48,48H400a48,48,0,0,0,48-48V336A16,16,0,0,0,432,320ZM488,0h-128c-21.37,0-32.05,25.91-17,41l35.73,35.73L135,320.37a24,24,0,0,0,0,34L157.67,377a24,24,0,0,0,34,0L435.28,133.32,471,169c15,15,41,4.5,41-17V24A24,24,0,0,0,488,0Z">
                                      </path>
                                    </svg></span></button></div>
                            </div>
                          </div>
                        </div>
                        <div class="carousel-item">
                          <div class="c-body-berita-terkini"
                            style="background-image: url(&quot;https://api-webdisdik.jabarprov.go.id/storage/disdik_web/assets/news/img/attachment-1741574418939.jpeg&quot;);">
                            <div class="c-head-berita-terkini">
                              <h4>Redesign Thinking untuk Peningkatan SDM Pendidikan Berdaya Saing Global (Bagi Guru,
                                Kepala Sekolah, dan Pengawas Sekolah)</h4>
                              <div class="text__details">
                                <p class="text__detail"><span class="icons icons-color"><svg aria-hidden="true"
                                      focusable="false" data-prefix="fas" data-icon="calendar-day"
                                      class="svg-inline--fa fa-calendar-day fa-w-14 " role="img"
                                      xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512">
                                      <path fill="currentColor"
                                        d="M0 464c0 26.5 21.5 48 48 48h352c26.5 0 48-21.5 48-48V192H0v272zm64-192c0-8.8 7.2-16 16-16h96c8.8 0 16 7.2 16 16v96c0 8.8-7.2 16-16 16H80c-8.8 0-16-7.2-16-16v-96zM400 64h-48V16c0-8.8-7.2-16-16-16h-32c-8.8 0-16 7.2-16 16v48H160V16c0-8.8-7.2-16-16-16h-32c-8.8 0-16 7.2-16 16v48H48C21.5 64 0 85.5 0 112v48h448v-48c0-26.5-21.5-48-48-48z">
                                      </path>
                                    </svg><span class="icons icons-red"></span> </span>Mar 10, 2025</p>
                                <p class="text__detail__tag"><span class="icons icons-color"><svg aria-hidden="true"
                                      focusable="false" data-prefix="fas" data-icon="clock"
                                      class="svg-inline--fa fa-clock fa-w-16 " role="img"
                                      xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">
                                      <path fill="currentColor"
                                        d="M256,8C119,8,8,119,8,256S119,504,256,504,504,393,504,256,393,8,256,8Zm92.49,313h0l-20,25a16,16,0,0,1-22.49,2.5h0l-67-49.72a40,40,0,0,1-15-31.23V112a16,16,0,0,1,16-16h32a16,16,0,0,1,16,16V256l58,42.5A16,16,0,0,1,348.49,321Z">
                                      </path>
                                    </svg> </span>09:40:18</p>
                                <p class="text__detail"><span class="icons icons-color"><svg aria-hidden="true"
                                      focusable="false" data-prefix="fas" data-icon="pencil-alt"
                                      class="svg-inline--fa fa-pencil-alt fa-w-16 " role="img"
                                      xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">
                                      <path fill="currentColor"
                                        d="M497.9 142.1l-46.1 46.1c-4.7 4.7-12.3 4.7-17 0l-111-111c-4.7-4.7-4.7-12.3 0-17l46.1-46.1c18.7-18.7 49.1-18.7 67.9 0l60.1 60.1c18.8 18.7 18.8 49.1 0 67.9zM284.2 99.8L21.6 362.4.4 483.9c-2.9 16.4 11.4 30.6 27.8 27.8l121.5-21.3 262.6-262.6c4.7-4.7 4.7-12.3 0-17l-111-111c-4.8-4.7-12.4-4.7-17.1 0zM124.1 339.9c-5.5-5.5-5.5-14.3 0-19.8l154-154c5.5-5.5 14.3-5.5 19.8 0s5.5 14.3 0 19.8l-154 154c-5.5 5.5-14.3 5.5-19.8 0zM88 424h48v36.3l-64.5 11.3-31.1-31.1L51.7 376H88v48z">
                                      </path>
                                    </svg> </span>Dina Martha Tiraswati, M.Pd. (Pengawas SMK Cabang Dinas Pendidikan
                                  Wilayah I Provinsi Jawa Barat)</p>
                              </div>
                              <div class="mb-2 title-container"><button class="btn btn-custom-home btn-md">Lihat
                                  Selengkapnya<span class="icon" style="margin-left: 5px;"><svg aria-hidden="true"
                                      focusable="false" data-prefix="fas" data-icon="external-link-alt"
                                      class="svg-inline--fa fa-external-link-alt fa-w-16 " role="img"
                                      xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">
                                      <path fill="currentColor"
                                        d="M432,320H400a16,16,0,0,0-16,16V448H64V128H208a16,16,0,0,0,16-16V80a16,16,0,0,0-16-16H48A48,48,0,0,0,0,112V464a48,48,0,0,0,48,48H400a48,48,0,0,0,48-48V336A16,16,0,0,0,432,320ZM488,0h-128c-21.37,0-32.05,25.91-17,41l35.73,35.73L135,320.37a24,24,0,0,0,0,34L157.67,377a24,24,0,0,0,34,0L435.28,133.32,471,169c15,15,41,4.5,41-17V24A24,24,0,0,0,488,0Z">
                                      </path>
                                    </svg></span></button></div>
                            </div>
                          </div>
                        </div>
                        <div class="carousel-item">
                          <div class="c-body-berita-terkini"
                            style="background-image: url(&quot;https://api-webdisdik.jabarprov.go.id/storage/disdik_web/assets/news/img/attachment-1741321186554.jpeg&quot;);">
                            <div class="c-head-berita-terkini">
                              <h4>Ragam Kegiatan SmartTren di SMAN 11 Bandung</h4>
                              <div class="text__details">
                                <p class="text__detail"><span class="icons icons-color"><svg aria-hidden="true"
                                      focusable="false" data-prefix="fas" data-icon="calendar-day"
                                      class="svg-inline--fa fa-calendar-day fa-w-14 " role="img"
                                      xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512">
                                      <path fill="currentColor"
                                        d="M0 464c0 26.5 21.5 48 48 48h352c26.5 0 48-21.5 48-48V192H0v272zm64-192c0-8.8 7.2-16 16-16h96c8.8 0 16 7.2 16 16v96c0 8.8-7.2 16-16 16H80c-8.8 0-16-7.2-16-16v-96zM400 64h-48V16c0-8.8-7.2-16-16-16h-32c-8.8 0-16 7.2-16 16v48H160V16c0-8.8-7.2-16-16-16h-32c-8.8 0-16 7.2-16 16v48H48C21.5 64 0 85.5 0 112v48h448v-48c0-26.5-21.5-48-48-48z">
                                      </path>
                                    </svg><span class="icons icons-red"></span> </span>Mar 07, 2025</p>
                                <p class="text__detail__tag"><span class="icons icons-color"><svg aria-hidden="true"
                                      focusable="false" data-prefix="fas" data-icon="clock"
                                      class="svg-inline--fa fa-clock fa-w-16 " role="img"
                                      xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">
                                      <path fill="currentColor"
                                        d="M256,8C119,8,8,119,8,256S119,504,256,504,504,393,504,256,393,8,256,8Zm92.49,313h0l-20,25a16,16,0,0,1-22.49,2.5h0l-67-49.72a40,40,0,0,1-15-31.23V112a16,16,0,0,1,16-16h32a16,16,0,0,1,16,16V256l58,42.5A16,16,0,0,1,348.49,321Z">
                                      </path>
                                    </svg> </span>11:19:46</p>
                                <p class="text__detail"><span class="icons icons-color"><svg aria-hidden="true"
                                      focusable="false" data-prefix="fas" data-icon="pencil-alt"
                                      class="svg-inline--fa fa-pencil-alt fa-w-16 " role="img"
                                      xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">
                                      <path fill="currentColor"
                                        d="M497.9 142.1l-46.1 46.1c-4.7 4.7-12.3 4.7-17 0l-111-111c-4.7-4.7-4.7-12.3 0-17l46.1-46.1c18.7-18.7 49.1-18.7 67.9 0l60.1 60.1c18.8 18.7 18.8 49.1 0 67.9zM284.2 99.8L21.6 362.4.4 483.9c-2.9 16.4 11.4 30.6 27.8 27.8l121.5-21.3 262.6-262.6c4.7-4.7 4.7-12.3 0-17l-111-111c-4.8-4.7-12.4-4.7-17.1 0zM124.1 339.9c-5.5-5.5-5.5-14.3 0-19.8l154-154c5.5-5.5 14.3-5.5 19.8 0s5.5 14.3 0 19.8l-154 154c-5.5 5.5-14.3 5.5-19.8 0zM88 424h48v36.3l-64.5 11.3-31.1-31.1L51.7 376H88v48z">
                                      </path>
                                    </svg> </span>Nizar</p>
                              </div>
                              <div class="mb-2 title-container"><button class="btn btn-custom-home btn-md">Lihat
                                  Selengkapnya<span class="icon" style="margin-left: 5px;"><svg aria-hidden="true"
                                      focusable="false" data-prefix="fas" data-icon="external-link-alt"
                                      class="svg-inline--fa fa-external-link-alt fa-w-16 " role="img"
                                      xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">
                                      <path fill="currentColor"
                                        d="M432,320H400a16,16,0,0,0-16,16V448H64V128H208a16,16,0,0,0,16-16V80a16,16,0,0,0-16-16H48A48,48,0,0,0,0,112V464a48,48,0,0,0,48,48H400a48,48,0,0,0,48-48V336A16,16,0,0,0,432,320ZM488,0h-128c-21.37,0-32.05,25.91-17,41l35.73,35.73L135,320.37a24,24,0,0,0,0,34L157.67,377a24,24,0,0,0,34,0L435.28,133.32,471,169c15,15,41,4.5,41-17V24A24,24,0,0,0,488,0Z">
                                      </path>
                                    </svg></span></button></div>
                            </div>
                          </div>
                        </div>
                        <div class="carousel-item">
                          <div class="c-body-berita-terkini"
                            style="background-image: url(&quot;https://api-webdisdik.jabarprov.go.id/storage/disdik_web/assets/news/img/attachment-1741178379732.jpg&quot;);">
                            <div class="c-head-berita-terkini">
                              <h4>29 Sekolah Terdampak Banjir di Bekasi, Tak Ada Kerusakan Bangunan</h4>
                              <div class="text__details">
                                <p class="text__detail"><span class="icons icons-color"><svg aria-hidden="true"
                                      focusable="false" data-prefix="fas" data-icon="calendar-day"
                                      class="svg-inline--fa fa-calendar-day fa-w-14 " role="img"
                                      xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512">
                                      <path fill="currentColor"
                                        d="M0 464c0 26.5 21.5 48 48 48h352c26.5 0 48-21.5 48-48V192H0v272zm64-192c0-8.8 7.2-16 16-16h96c8.8 0 16 7.2 16 16v96c0 8.8-7.2 16-16 16H80c-8.8 0-16-7.2-16-16v-96zM400 64h-48V16c0-8.8-7.2-16-16-16h-32c-8.8 0-16 7.2-16 16v48H160V16c0-8.8-7.2-16-16-16h-32c-8.8 0-16 7.2-16 16v48H48C21.5 64 0 85.5 0 112v48h448v-48c0-26.5-21.5-48-48-48z">
                                      </path>
                                    </svg><span class="icons icons-red"></span> </span>Mar 05, 2025</p>
                                <p class="text__detail__tag"><span class="icons icons-color"><svg aria-hidden="true"
                                      focusable="false" data-prefix="fas" data-icon="clock"
                                      class="svg-inline--fa fa-clock fa-w-16 " role="img"
                                      xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">
                                      <path fill="currentColor"
                                        d="M256,8C119,8,8,119,8,256S119,504,256,504,504,393,504,256,393,8,256,8Zm92.49,313h0l-20,25a16,16,0,0,1-22.49,2.5h0l-67-49.72a40,40,0,0,1-15-31.23V112a16,16,0,0,1,16-16h32a16,16,0,0,1,16,16V256l58,42.5A16,16,0,0,1,348.49,321Z">
                                      </path>
                                    </svg> </span>19:39:39</p>
                                <p class="text__detail"><span class="icons icons-color"><svg aria-hidden="true"
                                      focusable="false" data-prefix="fas" data-icon="pencil-alt"
                                      class="svg-inline--fa fa-pencil-alt fa-w-16 " role="img"
                                      xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">
                                      <path fill="currentColor"
                                        d="M497.9 142.1l-46.1 46.1c-4.7 4.7-12.3 4.7-17 0l-111-111c-4.7-4.7-4.7-12.3 0-17l46.1-46.1c18.7-18.7 49.1-18.7 67.9 0l60.1 60.1c18.8 18.7 18.8 49.1 0 67.9zM284.2 99.8L21.6 362.4.4 483.9c-2.9 16.4 11.4 30.6 27.8 27.8l121.5-21.3 262.6-262.6c4.7-4.7 4.7-12.3 0-17l-111-111c-4.8-4.7-12.4-4.7-17.1 0zM124.1 339.9c-5.5-5.5-5.5-14.3 0-19.8l154-154c5.5-5.5 14.3-5.5 19.8 0s5.5 14.3 0 19.8l-154 154c-5.5 5.5-14.3 5.5-19.8 0zM88 424h48v36.3l-64.5 11.3-31.1-31.1L51.7 376H88v48z">
                                      </path>
                                    </svg> </span>Nizar</p>
                              </div>
                              <div class="mb-2 title-container"><button class="btn btn-custom-home btn-md">Lihat
                                  Selengkapnya<span class="icon" style="margin-left: 5px;"><svg aria-hidden="true"
                                      focusable="false" data-prefix="fas" data-icon="external-link-alt"
                                      class="svg-inline--fa fa-external-link-alt fa-w-16 " role="img"
                                      xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">
                                      <path fill="currentColor"
                                        d="M432,320H400a16,16,0,0,0-16,16V448H64V128H208a16,16,0,0,0,16-16V80a16,16,0,0,0-16-16H48A48,48,0,0,0,0,112V464a48,48,0,0,0,48,48H400a48,48,0,0,0,48-48V336A16,16,0,0,0,432,320ZM488,0h-128c-21.37,0-32.05,25.91-17,41l35.73,35.73L135,320.37a24,24,0,0,0,0,34L157.67,377a24,24,0,0,0,34,0L435.28,133.32,471,169c15,15,41,4.5,41-17V24A24,24,0,0,0,488,0Z">
                                      </path>
                                    </svg></span></button></div>
                            </div>
                          </div>
                        </div>
                        <div class="active carousel-item">
                          <div class="c-body-berita-terkini"
                            style="background-image: url(&quot;https://api-webdisdik.jabarprov.go.id/storage/disdik_web/assets/news/img/attachment-1741133156582.jpeg&quot;);">
                            <div class="c-head-berita-terkini">
                              <h4>Komitmen Bersama, Sinergitas Pembangunan Jabar Istimewa untuk Bangun Pendidikan
                                Istimewa</h4>
                              <div class="text__details">
                                <p class="text__detail"><span class="icons icons-color"><svg aria-hidden="true"
                                      focusable="false" data-prefix="fas" data-icon="calendar-day"
                                      class="svg-inline--fa fa-calendar-day fa-w-14 " role="img"
                                      xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512">
                                      <path fill="currentColor"
                                        d="M0 464c0 26.5 21.5 48 48 48h352c26.5 0 48-21.5 48-48V192H0v272zm64-192c0-8.8 7.2-16 16-16h96c8.8 0 16 7.2 16 16v96c0 8.8-7.2 16-16 16H80c-8.8 0-16-7.2-16-16v-96zM400 64h-48V16c0-8.8-7.2-16-16-16h-32c-8.8 0-16 7.2-16 16v48H160V16c0-8.8-7.2-16-16-16h-32c-8.8 0-16 7.2-16 16v48H48C21.5 64 0 85.5 0 112v48h448v-48c0-26.5-21.5-48-48-48z">
                                      </path>
                                    </svg><span class="icons icons-red"></span> </span>Mar 05, 2025</p>
                                <p class="text__detail__tag"><span class="icons icons-color"><svg aria-hidden="true"
                                      focusable="false" data-prefix="fas" data-icon="clock"
                                      class="svg-inline--fa fa-clock fa-w-16 " role="img"
                                      xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">
                                      <path fill="currentColor"
                                        d="M256,8C119,8,8,119,8,256S119,504,256,504,504,393,504,256,393,8,256,8Zm92.49,313h0l-20,25a16,16,0,0,1-22.49,2.5h0l-67-49.72a40,40,0,0,1-15-31.23V112a16,16,0,0,1,16-16h32a16,16,0,0,1,16,16V256l58,42.5A16,16,0,0,1,348.49,321Z">
                                      </path>
                                    </svg> </span>07:05:56</p>
                                <p class="text__detail"><span class="icons icons-color"><svg aria-hidden="true"
                                      focusable="false" data-prefix="fas" data-icon="pencil-alt"
                                      class="svg-inline--fa fa-pencil-alt fa-w-16 " role="img"
                                      xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">
                                      <path fill="currentColor"
                                        d="M497.9 142.1l-46.1 46.1c-4.7 4.7-12.3 4.7-17 0l-111-111c-4.7-4.7-4.7-12.3 0-17l46.1-46.1c18.7-18.7 49.1-18.7 67.9 0l60.1 60.1c18.8 18.7 18.8 49.1 0 67.9zM284.2 99.8L21.6 362.4.4 483.9c-2.9 16.4 11.4 30.6 27.8 27.8l121.5-21.3 262.6-262.6c4.7-4.7 4.7-12.3 0-17l-111-111c-4.8-4.7-12.4-4.7-17.1 0zM124.1 339.9c-5.5-5.5-5.5-14.3 0-19.8l154-154c5.5-5.5 14.3-5.5 19.8 0s5.5 14.3 0 19.8l-154 154c-5.5 5.5-14.3 5.5-19.8 0zM88 424h48v36.3l-64.5 11.3-31.1-31.1L51.7 376H88v48z">
                                      </path>
                                    </svg> </span>Rury Yuliatri</p>
                              </div>
                              <div class="mb-2 title-container"><button class="btn btn-custom-home btn-md">Lihat
                                  Selengkapnya<span class="icon" style="margin-left: 5px;"><svg aria-hidden="true"
                                      focusable="false" data-prefix="fas" data-icon="external-link-alt"
                                      class="svg-inline--fa fa-external-link-alt fa-w-16 " role="img"
                                      xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">
                                      <path fill="currentColor"
                                        d="M432,320H400a16,16,0,0,0-16,16V448H64V128H208a16,16,0,0,0,16-16V80a16,16,0,0,0-16-16H48A48,48,0,0,0,0,112V464a48,48,0,0,0,48,48H400a48,48,0,0,0,48-48V336A16,16,0,0,0,432,320ZM488,0h-128c-21.37,0-32.05,25.91-17,41l35.73,35.73L135,320.37a24,24,0,0,0,0,34L157.67,377a24,24,0,0,0,34,0L435.28,133.32,471,169c15,15,41,4.5,41-17V24A24,24,0,0,0,488,0Z">
                                      </path>
                                    </svg></span></button></div>
                            </div>
                          </div>
                        </div>
                        <div class="carousel-item">
                          <div class="c-body-berita-terkini"
                            style="background-image: url(&quot;https://api-webdisdik.jabarprov.go.id/storage/disdik_web/assets/news/img/attachment-1740985994340.jpeg&quot;);">
                            <div class="c-head-berita-terkini">
                              <h4>Pembukaan SmartTren 2025, Kadisdik: Beri Teladan dan Ajarkan Kebaikan</h4>
                              <div class="text__details">
                                <p class="text__detail"><span class="icons icons-color"><svg aria-hidden="true"
                                      focusable="false" data-prefix="fas" data-icon="calendar-day"
                                      class="svg-inline--fa fa-calendar-day fa-w-14 " role="img"
                                      xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512">
                                      <path fill="currentColor"
                                        d="M0 464c0 26.5 21.5 48 48 48h352c26.5 0 48-21.5 48-48V192H0v272zm64-192c0-8.8 7.2-16 16-16h96c8.8 0 16 7.2 16 16v96c0 8.8-7.2 16-16 16H80c-8.8 0-16-7.2-16-16v-96zM400 64h-48V16c0-8.8-7.2-16-16-16h-32c-8.8 0-16 7.2-16 16v48H160V16c0-8.8-7.2-16-16-16h-32c-8.8 0-16 7.2-16 16v48H48C21.5 64 0 85.5 0 112v48h448v-48c0-26.5-21.5-48-48-48z">
                                      </path>
                                    </svg><span class="icons icons-red"></span> </span>Mar 03, 2025</p>
                                <p class="text__detail__tag"><span class="icons icons-color"><svg aria-hidden="true"
                                      focusable="false" data-prefix="fas" data-icon="clock"
                                      class="svg-inline--fa fa-clock fa-w-16 " role="img"
                                      xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">
                                      <path fill="currentColor"
                                        d="M256,8C119,8,8,119,8,256S119,504,256,504,504,393,504,256,393,8,256,8Zm92.49,313h0l-20,25a16,16,0,0,1-22.49,2.5h0l-67-49.72a40,40,0,0,1-15-31.23V112a16,16,0,0,1,16-16h32a16,16,0,0,1,16,16V256l58,42.5A16,16,0,0,1,348.49,321Z">
                                      </path>
                                    </svg> </span>14:13:14</p>
                                <p class="text__detail"><span class="icons icons-color"><svg aria-hidden="true"
                                      focusable="false" data-prefix="fas" data-icon="pencil-alt"
                                      class="svg-inline--fa fa-pencil-alt fa-w-16 " role="img"
                                      xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">
                                      <path fill="currentColor"
                                        d="M497.9 142.1l-46.1 46.1c-4.7 4.7-12.3 4.7-17 0l-111-111c-4.7-4.7-4.7-12.3 0-17l46.1-46.1c18.7-18.7 49.1-18.7 67.9 0l60.1 60.1c18.8 18.7 18.8 49.1 0 67.9zM284.2 99.8L21.6 362.4.4 483.9c-2.9 16.4 11.4 30.6 27.8 27.8l121.5-21.3 262.6-262.6c4.7-4.7 4.7-12.3 0-17l-111-111c-4.8-4.7-12.4-4.7-17.1 0zM124.1 339.9c-5.5-5.5-5.5-14.3 0-19.8l154-154c5.5-5.5 14.3-5.5 19.8 0s5.5 14.3 0 19.8l-154 154c-5.5 5.5-14.3 5.5-19.8 0zM88 424h48v36.3l-64.5 11.3-31.1-31.1L51.7 376H88v48z">
                                      </path>
                                    </svg> </span>Nizar</p>
                              </div>
                              <div class="mb-2 title-container"><button class="btn btn-custom-home btn-md">Lihat
                                  Selengkapnya<span class="icon" style="margin-left: 5px;"><svg aria-hidden="true"
                                      focusable="false" data-prefix="fas" data-icon="external-link-alt"
                                      class="svg-inline--fa fa-external-link-alt fa-w-16 " role="img"
                                      xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">
                                      <path fill="currentColor"
                                        d="M432,320H400a16,16,0,0,0-16,16V448H64V128H208a16,16,0,0,0,16-16V80a16,16,0,0,0-16-16H48A48,48,0,0,0,0,112V464a48,48,0,0,0,48,48H400a48,48,0,0,0,48-48V336A16,16,0,0,0,432,320ZM488,0h-128c-21.37,0-32.05,25.91-17,41l35.73,35.73L135,320.37a24,24,0,0,0,0,34L157.67,377a24,24,0,0,0,34,0L435.28,133.32,471,169c15,15,41,4.5,41-17V24A24,24,0,0,0,488,0Z">
                                      </path>
                                    </svg></span></button></div>
                            </div>
                          </div>
                        </div>
                      </div><a class="carousel-control-prev" role="button" href="#"><span aria-hidden="true"
                          class="carousel-control-prev-icon"></span><span class="sr-only">Previous</span></a><a
                        class="carousel-control-next" role="button" href="#"><span aria-hidden="true"
                          class="carousel-control-next-icon"></span><span class="sr-only">Next</span></a>
                    </div>
                  </div>
                  <div class="col-lg-4 col-md-4 col-sm-12 mt-3">
                    <div class="tabsInfo-container">
                      <div class="tabs-menu-custom">
                        <div class="tabInfo active"><a href="/sapa-disdik/agenda"> Agenda</a></div>
                      </div>
                      <div class="tabs-content-custom">
                        <div>
                          <div class="row">
                            <div class="col-lg-12 col-sm-12" style="max-height: 100px;">
                              <div class="d-flex justify-content-center align-items-center" style="width: 100%;">
                                <ul class="agenda-week">
                                  <li class="d-flex flex-column justify-content-center align-items-center "><span
                                      class="hari">Min</span><span class="tanggal">9</span></li>
                                  <li class="d-flex flex-column justify-content-center align-items-center "><span
                                      class="hari">Sen</span><span class="tanggal">10</span></li>
                                  <li class="d-flex flex-column justify-content-center align-items-center "><span
                                      class="hari">Sel</span><span class="tanggal">11</span></li>
                                  <li class="d-flex flex-column justify-content-center align-items-center "><span
                                      class="hari">Rab</span><span class="tanggal">12</span></li>
                                  <li class="d-flex flex-column justify-content-center align-items-center active"><span
                                      class="hari">Kam</span><span class="tanggal">13</span></li>
                                  <li class="d-flex flex-column justify-content-center align-items-center "><span
                                      class="hari">Jum</span><span class="tanggal">14</span></li>
                                  <li class="d-flex flex-column justify-content-center align-items-center "><span
                                      class="hari">Sab</span><span class="tanggal">15</span></li>
                                </ul>
                              </div>
                            </div>
                            <div class="col-lg-12 col-sm-12">
                              <div>
                                <div class="d-flex justify-content-center align-items-center" style="height: 100px;">
                                  <div>Tidak ada kegiatan di hari ini</div>
                                </div>
                              </div>
                            </div>
                            <div class="col col-lg-12"></div>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="mb-5">
            <div data-test="carousel" length="3" class="carousel carousel-fade z-depth-1 mt-4 black"
              aria-label="carousel" style="z-index: -99;">
              <div data-test="carousel-inner" class="carousel-inner">
                <div data-test="carousel-item" class="carousel-item active carousel-slide-item" style="left: 0px;">
                  <div class="view" data-test="view">
                    <div class="row justify-content-md-center" style="margin: 0.01%;">
                      <div class="col-12 col-lg-4 c-list-news" style="position: relative; padding: 0px;">
                        <div class="col-separator d-flex justify-content-center flex-column black-and-white"
                          style="background-image: url(&quot;https://api-webdisdik.jabarprov.go.id/storage/disdik_web/assets/news/img/attachment-1741685344746.jpg&quot;);">
                        </div>
                        <div class="overlay-col-separator-blue"
                          style="display: flex; position: absolute; z-index: 99; top: 0px; bottom: 0px;"></div>
                        <div
                          style="display: flex; justify-content: center; align-items: center; position: absolute; z-index: 100; top: 0px; bottom: 0px; height: 100%;">
                          <h4 class="col-separator-subtitle align-self-center">Sengketa Lahan SMAN 1 Bandung,
                            Wamendikdasmen: Semoga Putusan Sesuai Fakta</h4>
                        </div>
                      </div>
                      <div class="col-12 col-lg-4 c-list-news" style="position: relative; padding: 0px;">
                        <div class="col-separator d-flex justify-content-center flex-column black-and-white"
                          style="background-image: url(&quot;https://api-webdisdik.jabarprov.go.id/storage/disdik_web/assets/news/img/attachment-1741574418939.jpeg&quot;);">
                        </div>
                        <div class="overlay-col-separator-yellow"
                          style="display: flex; position: absolute; z-index: 99; top: 0px; bottom: 0px;"></div>
                        <div
                          style="display: flex; justify-content: center; align-items: center; position: absolute; z-index: 100; top: 0px; bottom: 0px; height: 100%;">
                          <h4 class="col-separator-subtitle align-self-center">Redesign Thinking untuk Peningkatan SDM
                            Pendidikan Berdaya Saing Global (Bagi Guru, Kepala Sekolah, dan Pengawas Sekolah)</h4>
                        </div>
                      </div>
                      <div class="col-12 col-lg-4 c-list-news" style="position: relative; padding: 0px;">
                        <div class="col-separator d-flex justify-content-center flex-column black-and-white"
                          style="background-image: url(&quot;https://api-webdisdik.jabarprov.go.id/storage/disdik_web/assets/news/img/attachment-1741321186554.jpeg&quot;);">
                        </div>
                        <div class="overlay-col-separator-green"
                          style="display: flex; position: absolute; z-index: 99; top: 0px; bottom: 0px;"></div>
                        <div
                          style="display: flex; justify-content: center; align-items: center; position: absolute; z-index: 100; top: 0px; bottom: 0px; height: 100%;">
                          <h4 class="col-separator-subtitle align-self-center">Ragam Kegiatan SmartTren di SMAN 11
                            Bandung</h4>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
                <div data-test="carousel-item" class="carousel-item active carousel-slide-item"
                  style="left: 100%; position: absolute;">
                  <div class="view" data-test="view">
                    <div class="row justify-content-md-center" style="margin: 0.01%;">
                      <div class="col-12 col-lg-4 c-list-news" style="position: relative; padding: 0px;">
                        <div class="col-separator d-flex justify-content-center flex-column black-and-white"
                          style="background-image: url(&quot;https://api-webdisdik.jabarprov.go.id/storage/disdik_web/assets/news/img/attachment-1741178379732.jpg&quot;);">
                        </div>
                        <div class="overlay-col-separator-blue"
                          style="display: flex; position: absolute; z-index: 99; top: 0px; bottom: 0px;"></div>
                        <div
                          style="display: flex; justify-content: center; align-items: center; position: absolute; z-index: 100; top: 0px; bottom: 0px; height: 100%;">
                          <h4 class="col-separator-subtitle align-self-center">29 Sekolah Terdampak Banjir di Bekasi,
                            Tak Ada Kerusakan Bangunan</h4>
                        </div>
                      </div>
                      <div class="col-12 col-lg-4 c-list-news" style="position: relative; padding: 0px;">
                        <div class="col-separator d-flex justify-content-center flex-column black-and-white"
                          style="background-image: url(&quot;https://api-webdisdik.jabarprov.go.id/storage/disdik_web/assets/news/img/attachment-1741133156582.jpeg&quot;);">
                        </div>
                        <div class="overlay-col-separator-yellow"
                          style="display: flex; position: absolute; z-index: 99; top: 0px; bottom: 0px;"></div>
                        <div
                          style="display: flex; justify-content: center; align-items: center; position: absolute; z-index: 100; top: 0px; bottom: 0px; height: 100%;">
                          <h4 class="col-separator-subtitle align-self-center">Komitmen Bersama, Sinergitas Pembangunan
                            Jabar Istimewa untuk Bangun Pendidikan Istimewa</h4>
                        </div>
                      </div>
                      <div class="col-12 col-lg-4 c-list-news" style="position: relative; padding: 0px;">
                        <div class="col-separator d-flex justify-content-center flex-column black-and-white"
                          style="background-image: url(&quot;https://api-webdisdik.jabarprov.go.id/storage/disdik_web/assets/news/img/attachment-1740985994340.jpeg&quot;);">
                        </div>
                        <div class="overlay-col-separator-green"
                          style="display: flex; position: absolute; z-index: 99; top: 0px; bottom: 0px;"></div>
                        <div
                          style="display: flex; justify-content: center; align-items: center; position: absolute; z-index: 100; top: 0px; bottom: 0px; height: 100%;">
                          <h4 class="col-separator-subtitle align-self-center">Pembukaan SmartTren 2025, Kadisdik: Beri
                            Teladan dan Ajarkan Kebaikan</h4>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
                <div data-test="carousel-item" class="carousel-item active carousel-slide-item"
                  style="left: 100%; position: absolute;">
                  <div class="view" data-test="view">
                    <div class="row justify-content-md-center" style="margin: 0.01%;">
                      <div class="col-12 col-lg-4 c-list-news" style="position: relative; padding: 0px;">
                        <div class="col-separator d-flex justify-content-center flex-column black-and-white"
                          style="background-image: url(&quot;https://api-webdisdik.jabarprov.go.id/storage/disdik_web/assets/news/img/attachment-1740570020114.jpeg&quot;);">
                        </div>
                        <div class="overlay-col-separator-blue"
                          style="display: flex; position: absolute; z-index: 99; top: 0px; bottom: 0px;"></div>
                        <div
                          style="display: flex; justify-content: center; align-items: center; position: absolute; z-index: 100; top: 0px; bottom: 0px; height: 100%;">
                          <h4 class="col-separator-subtitle align-self-center">FOPD Bidang Pendidikan 2025, Sekda Jabar:
                            Buat Kebijakan Pendidikan yang Konkret!</h4>
                        </div>
                      </div>
                      <div class="col-12 col-lg-4 c-list-news" style="position: relative; padding: 0px;">
                        <div class="col-separator d-flex justify-content-center flex-column black-and-white"
                          style="background-image: url(&quot;https://api-webdisdik.jabarprov.go.id/storage/disdik_web/assets/news/img/attachment-1740569258032.jpeg&quot;);">
                        </div>
                        <div class="overlay-col-separator-yellow"
                          style="display: flex; position: absolute; z-index: 99; top: 0px; bottom: 0px;"></div>
                        <div
                          style="display: flex; justify-content: center; align-items: center; position: absolute; z-index: 100; top: 0px; bottom: 0px; height: 100%;">
                          <h4 class="col-separator-subtitle align-self-center">Kolaborasi Disdik Jabar dan Lazada,
                            Dorong Transformasi Insan SMK &nbsp;Menjadi Wirausaha Tangguh</h4>
                        </div>
                      </div>
                      <div class="col-12 col-lg-4 c-list-news" style="position: relative; padding: 0px;">
                        <div class="col-separator d-flex justify-content-center flex-column black-and-white"
                          style="background-image: url(&quot;https://api-webdisdik.jabarprov.go.id/storage/disdik_web/assets/news/img/attachment-1739437193910.jpeg&quot;);">
                        </div>
                        <div class="overlay-col-separator-green"
                          style="display: flex; position: absolute; z-index: 99; top: 0px; bottom: 0px;"></div>
                        <div
                          style="display: flex; justify-content: center; align-items: center; position: absolute; z-index: 100; top: 0px; bottom: 0px; height: 100%;">
                          <h4 class="col-separator-subtitle align-self-center">Sertifikasi TOEIC Bagi Siswa SMK,
                            Ciptakan Lulusan yang Berdaya Saing</h4>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div class="container mt-5 ">
              <div class="d-flex justify-content-start align-items-start padding-null">
                <h3 class="align-self-end mb-3">
                  <div class="title-news-home">Informasi</div>
                </h3>
              </div>
            </div>
            <div class="container" style="z-index: -99;">
              <div class="about-box">
                <div class="slider-container-custom" style="width: 100%;">
                  <div style="padding: 10px;">
                    <div class="slick-slider slick-initialized" dir="ltr">
                      <div class="slick-arrow slick-prev slick-disabled"
                        style="display: flex; justify-content: center; align-items: center; width: 50px; height: 50px; background: rgb(175, 211, 244); border-radius: 50%; margin-left: 10px; z-index: 99;">
                      </div>
                      <div class="slick-list">
                        <div class="slick-track"
                          style="width: 3640px; opacity: 1; transform: translate3d(0px, 0px, 0px);">
                          <div data-index="0" class="slick-slide slick-active slick-current" tabindex="-1"
                            aria-hidden="false" style="outline: none; width: 364px;">
                            <div>
                              <div tabindex="-1" style="width: 100%; display: inline-block;">
                                <div class="card card-akses-cepat" style="min-height: 380px; margin: 10px;">
                                  <div class="card-body d-flex flex-column justify-content-between">
                                    <div><span
                                        style="display: flex; width: 100%; justify-content: center; align-items: center; background-color: aliceblue; border-radius: 12px; padding: 15px; margin-bottom: 10px;"><svg
                                          aria-hidden="true" focusable="false" data-prefix="fas"
                                          data-icon="graduation-cap"
                                          class="svg-inline--fa fa-graduation-cap fa-w-20 fa-3x mb-2" role="img"
                                          xmlns="http://www.w3.org/2000/svg" viewBox="0 0 640 512"
                                          style="color: rgb(52, 141, 223);">
                                          <path fill="currentColor"
                                            d="M622.34 153.2L343.4 67.5c-15.2-4.67-31.6-4.67-46.79 0L17.66 153.2c-23.54 7.23-23.54 38.36 0 45.59l48.63 14.94c-10.67 13.19-17.23 29.28-17.88 46.9C38.78 266.15 32 276.11 32 288c0 10.78 5.68 19.85 13.86 25.65L20.33 428.53C18.11 438.52 25.71 448 35.94 448h56.11c10.24 0 17.84-9.48 15.62-19.47L82.14 313.65C90.32 307.85 96 298.78 96 288c0-11.57-6.47-21.25-15.66-26.87.76-15.02 8.44-28.3 20.69-36.72L296.6 284.5c9.06 2.78 26.44 6.25 46.79 0l278.95-85.7c23.55-7.24 23.55-38.36 0-45.6zM352.79 315.09c-28.53 8.76-52.84 3.92-65.59 0l-145.02-44.55L128 384c0 35.35 85.96 64 192 64s192-28.65 192-64l-14.18-113.47-145.03 44.56z">
                                          </path>
                                        </svg></span>
                                      <h5 style="text-transform: capitalize;">PPDB Dinas Pendidikan</h5>
                                      <p class="mt-2"> Pendaftaran dan Pengumuman hasil seleksi secara terbuka, hingga
                                        daftar ulang....</p>
                                    </div><a href="https://ppdb.jabarprov.go.id/" target="_blank"
                                      rel="noopener noreferrer" class="btn btn-custom-home btn-sm mt-2">Lihat
                                      Selengkapnya</a>
                                  </div>
                                </div>
                              </div>
                            </div>
                          </div>
                          <div data-index="1" class="slick-slide slick-active" tabindex="-1" aria-hidden="false"
                            style="outline: none; width: 364px;">
                            <div>
                              <div tabindex="-1" style="width: 100%; display: inline-block;">
                                <div class="card card-akses-cepat" style="min-height: 380px; margin: 10px;">
                                  <div class="card-body d-flex flex-column justify-content-between">
                                    <div><span
                                        style="display: flex; width: 100%; justify-content: center; align-items: center; background-color: aliceblue; border-radius: 12px; padding: 15px; margin-bottom: 10px;"><svg
                                          aria-hidden="true" focusable="false" data-prefix="fas"
                                          data-icon="graduation-cap"
                                          class="svg-inline--fa fa-graduation-cap fa-w-20 fa-3x mb-2" role="img"
                                          xmlns="http://www.w3.org/2000/svg" viewBox="0 0 640 512"
                                          style="color: rgb(52, 141, 223);">
                                          <path fill="currentColor"
                                            d="M622.34 153.2L343.4 67.5c-15.2-4.67-31.6-4.67-46.79 0L17.66 153.2c-23.54 7.23-23.54 38.36 0 45.59l48.63 14.94c-10.67 13.19-17.23 29.28-17.88 46.9C38.78 266.15 32 276.11 32 288c0 10.78 5.68 19.85 13.86 25.65L20.33 428.53C18.11 438.52 25.71 448 35.94 448h56.11c10.24 0 17.84-9.48 15.62-19.47L82.14 313.65C90.32 307.85 96 298.78 96 288c0-11.57-6.47-21.25-15.66-26.87.76-15.02 8.44-28.3 20.69-36.72L296.6 284.5c9.06 2.78 26.44 6.25 46.79 0l278.95-85.7c23.55-7.24 23.55-38.36 0-45.6zM352.79 315.09c-28.53 8.76-52.84 3.92-65.59 0l-145.02-44.55L128 384c0 35.35 85.96 64 192 64s192-28.65 192-64l-14.18-113.47-145.03 44.56z">
                                          </path>
                                        </svg></span>
                                      <h5 style="text-transform: capitalize;">Dokumen Pendukung PPDB 2024</h5>
                                      <p class="mt-2">Dokumen Pendukung PPDB 2024....</p>
                                    </div><a href="https://ppdb.jabarprov.go.id/dokumen" target="_blank"
                                      rel="noopener noreferrer" class="btn btn-custom-home btn-sm mt-2">Lihat
                                      Selengkapnya</a>
                                  </div>
                                </div>
                              </div>
                            </div>
                          </div>
                          <div data-index="2" class="slick-slide slick-active" tabindex="-1" aria-hidden="false"
                            style="outline: none; width: 364px;">
                            <div>
                              <div tabindex="-1" style="width: 100%; display: inline-block;">
                                <div class="card card-akses-cepat" style="min-height: 380px; margin: 10px;">
                                  <div class="card-body d-flex flex-column justify-content-between">
                                    <div><span
                                        style="display: flex; width: 100%; justify-content: center; align-items: center; background-color: aliceblue; border-radius: 12px; padding: 15px; margin-bottom: 10px;"><svg
                                          aria-hidden="true" focusable="false" data-prefix="fas"
                                          data-icon="graduation-cap"
                                          class="svg-inline--fa fa-graduation-cap fa-w-20 fa-3x mb-2" role="img"
                                          xmlns="http://www.w3.org/2000/svg" viewBox="0 0 640 512"
                                          style="color: rgb(52, 141, 223);">
                                          <path fill="currentColor"
                                            d="M622.34 153.2L343.4 67.5c-15.2-4.67-31.6-4.67-46.79 0L17.66 153.2c-23.54 7.23-23.54 38.36 0 45.59l48.63 14.94c-10.67 13.19-17.23 29.28-17.88 46.9C38.78 266.15 32 276.11 32 288c0 10.78 5.68 19.85 13.86 25.65L20.33 428.53C18.11 438.52 25.71 448 35.94 448h56.11c10.24 0 17.84-9.48 15.62-19.47L82.14 313.65C90.32 307.85 96 298.78 96 288c0-11.57-6.47-21.25-15.66-26.87.76-15.02 8.44-28.3 20.69-36.72L296.6 284.5c9.06 2.78 26.44 6.25 46.79 0l278.95-85.7c23.55-7.24 23.55-38.36 0-45.6zM352.79 315.09c-28.53 8.76-52.84 3.92-65.59 0l-145.02-44.55L128 384c0 35.35 85.96 64 192 64s192-28.65 192-64l-14.18-113.47-145.03 44.56z">
                                          </path>
                                        </svg></span>
                                      <h5 style="text-transform: capitalize;">Survey Lulusan SMA</h5>
                                      <p class="mt-2">Kuisioner ini dapat memberikan wawasan mendalam tentang kualitas
                                        pendidikan yang disediakan oleh sekolah-sekolah di bawa...</p>
                                    </div><a href="https://disdik.jabarprov.go.id/survey-lulusan-sma" target="_blank"
                                      rel="noopener noreferrer" class="btn btn-custom-home btn-sm mt-2">Lihat
                                      Selengkapnya</a>
                                  </div>
                                </div>
                              </div>
                            </div>
                          </div>
                          <div data-index="3" class="slick-slide" tabindex="-1" aria-hidden="true"
                            style="outline: none; width: 364px;">
                            <div>
                              <div tabindex="-1" style="width: 100%; display: inline-block;">
                                <div class="card card-akses-cepat" style="min-height: 380px; margin: 10px;">
                                  <div class="card-body d-flex flex-column justify-content-between">
                                    <div><span
                                        style="display: flex; width: 100%; justify-content: center; align-items: center; background-color: aliceblue; border-radius: 12px; padding: 15px; margin-bottom: 10px;"><svg
                                          aria-hidden="true" focusable="false" data-prefix="fas"
                                          data-icon="graduation-cap"
                                          class="svg-inline--fa fa-graduation-cap fa-w-20 fa-3x mb-2" role="img"
                                          xmlns="http://www.w3.org/2000/svg" viewBox="0 0 640 512"
                                          style="color: rgb(52, 141, 223);">
                                          <path fill="currentColor"
                                            d="M622.34 153.2L343.4 67.5c-15.2-4.67-31.6-4.67-46.79 0L17.66 153.2c-23.54 7.23-23.54 38.36 0 45.59l48.63 14.94c-10.67 13.19-17.23 29.28-17.88 46.9C38.78 266.15 32 276.11 32 288c0 10.78 5.68 19.85 13.86 25.65L20.33 428.53C18.11 438.52 25.71 448 35.94 448h56.11c10.24 0 17.84-9.48 15.62-19.47L82.14 313.65C90.32 307.85 96 298.78 96 288c0-11.57-6.47-21.25-15.66-26.87.76-15.02 8.44-28.3 20.69-36.72L296.6 284.5c9.06 2.78 26.44 6.25 46.79 0l278.95-85.7c23.55-7.24 23.55-38.36 0-45.6zM352.79 315.09c-28.53 8.76-52.84 3.92-65.59 0l-145.02-44.55L128 384c0 35.35 85.96 64 192 64s192-28.65 192-64l-14.18-113.47-145.03 44.56z">
                                          </path>
                                        </svg></span>
                                      <h5 style="text-transform: capitalize;">Seleksi Calon Pengawas Sekolah (CPS)</h5>
                                      <p class="mt-2">Seleksi Calon Pengawas Sekolah (CPS) untuk satuan pendidikan
                                        jenjang SMA, SMK dan SLB....</p>
                                    </div><a
                                      href="https://api-webdisdik.jabarprov.go.id/storage/disdik_web/assets/home/pdf/Surat_Pemberitahuan_Pelaksanaan_Seleksi_CPS_tahun_2023_kcd_03042023_154818_signed.pdf"
                                      target="_blank" rel="noopener noreferrer"
                                      class="btn btn-custom-home btn-sm mt-2">Lihat Selengkapnya</a>
                                  </div>
                                </div>
                              </div>
                            </div>
                          </div>
                          <div data-index="4" class="slick-slide" tabindex="-1" aria-hidden="true"
                            style="outline: none; width: 364px;">
                            <div>
                              <div tabindex="-1" style="width: 100%; display: inline-block;">
                                <div class="card card-akses-cepat" style="min-height: 380px; margin: 10px;">
                                  <div class="card-body d-flex flex-column justify-content-between">
                                    <div><span
                                        style="display: flex; width: 100%; justify-content: center; align-items: center; background-color: aliceblue; border-radius: 12px; padding: 15px; margin-bottom: 10px;"><svg
                                          aria-hidden="true" focusable="false" data-prefix="fas"
                                          data-icon="graduation-cap"
                                          class="svg-inline--fa fa-graduation-cap fa-w-20 fa-3x mb-2" role="img"
                                          xmlns="http://www.w3.org/2000/svg" viewBox="0 0 640 512"
                                          style="color: rgb(52, 141, 223);">
                                          <path fill="currentColor"
                                            d="M622.34 153.2L343.4 67.5c-15.2-4.67-31.6-4.67-46.79 0L17.66 153.2c-23.54 7.23-23.54 38.36 0 45.59l48.63 14.94c-10.67 13.19-17.23 29.28-17.88 46.9C38.78 266.15 32 276.11 32 288c0 10.78 5.68 19.85 13.86 25.65L20.33 428.53C18.11 438.52 25.71 448 35.94 448h56.11c10.24 0 17.84-9.48 15.62-19.47L82.14 313.65C90.32 307.85 96 298.78 96 288c0-11.57-6.47-21.25-15.66-26.87.76-15.02 8.44-28.3 20.69-36.72L296.6 284.5c9.06 2.78 26.44 6.25 46.79 0l278.95-85.7c23.55-7.24 23.55-38.36 0-45.6zM352.79 315.09c-28.53 8.76-52.84 3.92-65.59 0l-145.02-44.55L128 384c0 35.35 85.96 64 192 64s192-28.65 192-64l-14.18-113.47-145.03 44.56z">
                                          </path>
                                        </svg></span>
                                      <h5 style="text-transform: capitalize;">Sistem Terintegrasi olah pengaduan dan
                                        perundungan (STOPPER)</h5>
                                      <p class="mt-2">Stopper adalah Sistem Terintegrasi olah pengaduan dan
                                        perundungan...</p>
                                    </div><a href="http://wa.me/6282126030038" target="_blank" rel="noopener noreferrer"
                                      class="btn btn-custom-home btn-sm mt-2">Lihat Selengkapnya</a>
                                  </div>
                                </div>
                              </div>
                            </div>
                          </div>
                          <div data-index="5" class="slick-slide" tabindex="-1" aria-hidden="true"
                            style="outline: none; width: 364px;">
                            <div>
                              <div tabindex="-1" style="width: 100%; display: inline-block;">
                                <div class="card card-akses-cepat" style="min-height: 380px; margin: 10px;">
                                  <div class="card-body d-flex flex-column justify-content-between">
                                    <div><span
                                        style="display: flex; width: 100%; justify-content: center; align-items: center; background-color: aliceblue; border-radius: 12px; padding: 15px; margin-bottom: 10px;"><svg
                                          aria-hidden="true" focusable="false" data-prefix="fas"
                                          data-icon="graduation-cap"
                                          class="svg-inline--fa fa-graduation-cap fa-w-20 fa-3x mb-2" role="img"
                                          xmlns="http://www.w3.org/2000/svg" viewBox="0 0 640 512"
                                          style="color: rgb(52, 141, 223);">
                                          <path fill="currentColor"
                                            d="M622.34 153.2L343.4 67.5c-15.2-4.67-31.6-4.67-46.79 0L17.66 153.2c-23.54 7.23-23.54 38.36 0 45.59l48.63 14.94c-10.67 13.19-17.23 29.28-17.88 46.9C38.78 266.15 32 276.11 32 288c0 10.78 5.68 19.85 13.86 25.65L20.33 428.53C18.11 438.52 25.71 448 35.94 448h56.11c10.24 0 17.84-9.48 15.62-19.47L82.14 313.65C90.32 307.85 96 298.78 96 288c0-11.57-6.47-21.25-15.66-26.87.76-15.02 8.44-28.3 20.69-36.72L296.6 284.5c9.06 2.78 26.44 6.25 46.79 0l278.95-85.7c23.55-7.24 23.55-38.36 0-45.6zM352.79 315.09c-28.53 8.76-52.84 3.92-65.59 0l-145.02-44.55L128 384c0 35.35 85.96 64 192 64s192-28.65 192-64l-14.18-113.47-145.03 44.56z">
                                          </path>
                                        </svg></span>
                                      <h5 style="text-transform: capitalize;">Data pokok pendidikan</h5>
                                      <p class="mt-2">Data Pokok Pendidikan...</p>
                                    </div><a href="https://dapo.kemdikbud.go.id/" target="_blank"
                                      rel="noopener noreferrer" class="btn btn-custom-home btn-sm mt-2">Lihat
                                      Selengkapnya</a>
                                  </div>
                                </div>
                              </div>
                            </div>
                          </div>
                          <div data-index="6" class="slick-slide" tabindex="-1" aria-hidden="true"
                            style="outline: none; width: 364px;">
                            <div>
                              <div tabindex="-1" style="width: 100%; display: inline-block;">
                                <div class="card card-akses-cepat" style="min-height: 380px; margin: 10px;">
                                  <div class="card-body d-flex flex-column justify-content-between">
                                    <div><span
                                        style="display: flex; width: 100%; justify-content: center; align-items: center; background-color: aliceblue; border-radius: 12px; padding: 15px; margin-bottom: 10px;"><svg
                                          aria-hidden="true" focusable="false" data-prefix="fas"
                                          data-icon="graduation-cap"
                                          class="svg-inline--fa fa-graduation-cap fa-w-20 fa-3x mb-2" role="img"
                                          xmlns="http://www.w3.org/2000/svg" viewBox="0 0 640 512"
                                          style="color: rgb(52, 141, 223);">
                                          <path fill="currentColor"
                                            d="M622.34 153.2L343.4 67.5c-15.2-4.67-31.6-4.67-46.79 0L17.66 153.2c-23.54 7.23-23.54 38.36 0 45.59l48.63 14.94c-10.67 13.19-17.23 29.28-17.88 46.9C38.78 266.15 32 276.11 32 288c0 10.78 5.68 19.85 13.86 25.65L20.33 428.53C18.11 438.52 25.71 448 35.94 448h56.11c10.24 0 17.84-9.48 15.62-19.47L82.14 313.65C90.32 307.85 96 298.78 96 288c0-11.57-6.47-21.25-15.66-26.87.76-15.02 8.44-28.3 20.69-36.72L296.6 284.5c9.06 2.78 26.44 6.25 46.79 0l278.95-85.7c23.55-7.24 23.55-38.36 0-45.6zM352.79 315.09c-28.53 8.76-52.84 3.92-65.59 0l-145.02-44.55L128 384c0 35.35 85.96 64 192 64s192-28.65 192-64l-14.18-113.47-145.03 44.56z">
                                          </path>
                                        </svg></span>
                                      <h5 style="text-transform: capitalize;">Layanan Fasilitas Magang</h5>
                                      <p class="mt-2">Sarana melaksanakan program Magang dan PKL bagi Mahasiswa dan
                                        Siswa SMK...</p>
                                    </div><a href="https://bit.ly/Layanan_Fasilitas_Magang" target="_blank"
                                      rel="noopener noreferrer" class="btn btn-custom-home btn-sm mt-2">Lihat
                                      Selengkapnya</a>
                                  </div>
                                </div>
                              </div>
                            </div>
                          </div>
                          <div data-index="7" class="slick-slide" tabindex="-1" aria-hidden="true"
                            style="outline: none; width: 364px;">
                            <div>
                              <div tabindex="-1" style="width: 100%; display: inline-block;">
                                <div class="card card-akses-cepat" style="min-height: 380px; margin: 10px;">
                                  <div class="card-body d-flex flex-column justify-content-between">
                                    <div><span
                                        style="display: flex; width: 100%; justify-content: center; align-items: center; background-color: aliceblue; border-radius: 12px; padding: 15px; margin-bottom: 10px;"><svg
                                          aria-hidden="true" focusable="false" data-prefix="fas"
                                          data-icon="graduation-cap"
                                          class="svg-inline--fa fa-graduation-cap fa-w-20 fa-3x mb-2" role="img"
                                          xmlns="http://www.w3.org/2000/svg" viewBox="0 0 640 512"
                                          style="color: rgb(52, 141, 223);">
                                          <path fill="currentColor"
                                            d="M622.34 153.2L343.4 67.5c-15.2-4.67-31.6-4.67-46.79 0L17.66 153.2c-23.54 7.23-23.54 38.36 0 45.59l48.63 14.94c-10.67 13.19-17.23 29.28-17.88 46.9C38.78 266.15 32 276.11 32 288c0 10.78 5.68 19.85 13.86 25.65L20.33 428.53C18.11 438.52 25.71 448 35.94 448h56.11c10.24 0 17.84-9.48 15.62-19.47L82.14 313.65C90.32 307.85 96 298.78 96 288c0-11.57-6.47-21.25-15.66-26.87.76-15.02 8.44-28.3 20.69-36.72L296.6 284.5c9.06 2.78 26.44 6.25 46.79 0l278.95-85.7c23.55-7.24 23.55-38.36 0-45.6zM352.79 315.09c-28.53 8.76-52.84 3.92-65.59 0l-145.02-44.55L128 384c0 35.35 85.96 64 192 64s192-28.65 192-64l-14.18-113.47-145.03 44.56z">
                                          </path>
                                        </svg></span>
                                      <h5 style="text-transform: capitalize;">Mutasi Peserta Didik jenjang SMA, SMK, SLB
                                      </h5>
                                      <p class="mt-2">Layanan proses perpindahan peserta didik pada jenjang pendidikan
                                        SMA, SMK, SLB dari satu sekolah ke sekolah lain baik da...</p>
                                    </div><a href="https://bit.ly/Pelayanan_Mutasi_SMA_SMK_SLB" target="_blank"
                                      rel="noopener noreferrer" class="btn btn-custom-home btn-sm mt-2">Lihat
                                      Selengkapnya</a>
                                  </div>
                                </div>
                              </div>
                            </div>
                          </div>
                          <div data-index="8" class="slick-slide" tabindex="-1" aria-hidden="true"
                            style="outline: none; width: 364px;">
                            <div>
                              <div tabindex="-1" style="width: 100%; display: inline-block;">
                                <div class="card card-akses-cepat" style="min-height: 380px; margin: 10px;">
                                  <div class="card-body d-flex flex-column justify-content-between">
                                    <div><span
                                        style="display: flex; width: 100%; justify-content: center; align-items: center; background-color: aliceblue; border-radius: 12px; padding: 15px; margin-bottom: 10px;"><svg
                                          aria-hidden="true" focusable="false" data-prefix="fas"
                                          data-icon="graduation-cap"
                                          class="svg-inline--fa fa-graduation-cap fa-w-20 fa-3x mb-2" role="img"
                                          xmlns="http://www.w3.org/2000/svg" viewBox="0 0 640 512"
                                          style="color: rgb(52, 141, 223);">
                                          <path fill="currentColor"
                                            d="M622.34 153.2L343.4 67.5c-15.2-4.67-31.6-4.67-46.79 0L17.66 153.2c-23.54 7.23-23.54 38.36 0 45.59l48.63 14.94c-10.67 13.19-17.23 29.28-17.88 46.9C38.78 266.15 32 276.11 32 288c0 10.78 5.68 19.85 13.86 25.65L20.33 428.53C18.11 438.52 25.71 448 35.94 448h56.11c10.24 0 17.84-9.48 15.62-19.47L82.14 313.65C90.32 307.85 96 298.78 96 288c0-11.57-6.47-21.25-15.66-26.87.76-15.02 8.44-28.3 20.69-36.72L296.6 284.5c9.06 2.78 26.44 6.25 46.79 0l278.95-85.7c23.55-7.24 23.55-38.36 0-45.6zM352.79 315.09c-28.53 8.76-52.84 3.92-65.59 0l-145.02-44.55L128 384c0 35.35 85.96 64 192 64s192-28.65 192-64l-14.18-113.47-145.03 44.56z">
                                          </path>
                                        </svg></span>
                                      <h5 style="text-transform: capitalize;">Legalisir Ijazah SMA, SMK, SLB</h5>
                                      <p class="mt-2">Pengesahan secara resmi kebenaran atau keabsahan fotokopi
                                        Ijazah/Surat Tanda Tamat Belajar/Surat Keterangan Pengganti Ij...</p>
                                    </div><a href="https://bit.ly/Layanan_Pengurusan_Ijazah_SMA_SMK_SLB" target="_blank"
                                      rel="noopener noreferrer" class="btn btn-custom-home btn-sm mt-2">Lihat
                                      Selengkapnya</a>
                                  </div>
                                </div>
                              </div>
                            </div>
                          </div>
                          <div data-index="9" class="slick-slide" tabindex="-1" aria-hidden="true"
                            style="outline: none; width: 364px;">
                            <div>
                              <div tabindex="-1" style="width: 100%; display: inline-block;">
                                <div class="card card-akses-cepat" style="min-height: 380px; margin: 10px;">
                                  <div class="card-body d-flex flex-column justify-content-between">
                                    <div><span
                                        style="display: flex; width: 100%; justify-content: center; align-items: center; background-color: aliceblue; border-radius: 12px; padding: 15px; margin-bottom: 10px;"><svg
                                          aria-hidden="true" focusable="false" data-prefix="fas"
                                          data-icon="graduation-cap"
                                          class="svg-inline--fa fa-graduation-cap fa-w-20 fa-3x mb-2" role="img"
                                          xmlns="http://www.w3.org/2000/svg" viewBox="0 0 640 512"
                                          style="color: rgb(52, 141, 223);">
                                          <path fill="currentColor"
                                            d="M622.34 153.2L343.4 67.5c-15.2-4.67-31.6-4.67-46.79 0L17.66 153.2c-23.54 7.23-23.54 38.36 0 45.59l48.63 14.94c-10.67 13.19-17.23 29.28-17.88 46.9C38.78 266.15 32 276.11 32 288c0 10.78 5.68 19.85 13.86 25.65L20.33 428.53C18.11 438.52 25.71 448 35.94 448h56.11c10.24 0 17.84-9.48 15.62-19.47L82.14 313.65C90.32 307.85 96 298.78 96 288c0-11.57-6.47-21.25-15.66-26.87.76-15.02 8.44-28.3 20.69-36.72L296.6 284.5c9.06 2.78 26.44 6.25 46.79 0l278.95-85.7c23.55-7.24 23.55-38.36 0-45.6zM352.79 315.09c-28.53 8.76-52.84 3.92-65.59 0l-145.02-44.55L128 384c0 35.35 85.96 64 192 64s192-28.65 192-64l-14.18-113.47-145.03 44.56z">
                                          </path>
                                        </svg></span>
                                      <h5 style="text-transform: capitalize;">Alur Pengajuan E-Fungsional</h5>
                                      <p class="mt-2">Alur proses pengajuan usulan uji kompetensi jabatan fungsional
                                        melalui E-Fungsional...</p>
                                    </div><a href="https://bit.ly/Alur_Pengajuan_E_Fungsional" target="_blank"
                                      rel="noopener noreferrer" class="btn btn-custom-home btn-sm mt-2">Lihat
                                      Selengkapnya</a>
                                  </div>
                                </div>
                              </div>
                            </div>
                          </div>
                        </div>
                      </div>
                      <div class="slick-arrow slick-next"
                        style="display: flex; justify-content: center; align-items: center; width: 50px; height: 50px; background: rgb(175, 211, 244); border-radius: 50%; margin-right: 8px;">
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div class="row mt-5 " style="margin: 0.01%; z-index: -99;">
              <div class="col-12 col-lg-10 d-flex justify-content-end align-items-end padding-null">
                <h3 class="align-self-end mb-3"><b class="pr-2"
                    style="font-family: Montserrat-Bold; text-transform: uppercase; color: rgb(13, 71, 161);">Selamat
                    Ibadah Puasa 1446 H</b></h3>
              </div>
            </div>
            <div class="row" style="margin: 0.01%; z-index: -99;">
              <div class="col-12 col-lg-2 d-flex justify-content-end padding-null gratifikasi" style="z-index: -99;">
                <div class="my-about-div">
                  <h3 class="d-none d-lg-block" title="zzzzzzzz">Ibadah Puasa 1446 H</h3>
                </div>
              </div>
              <div class="col-12 col-lg-3 padding-null d-flex align-items-center"
                style="background-color: rgb(255, 255, 255);"><img class="img-fluid" alt="Third slide"
                  src="https://api-webdisdik.jabarprov.go.id/storage/disdik_web/assets/home/image/attachment-file-1741051609550.jpg">
              </div>
              <div class="col-12 col-lg-5 about-box">
                <div class="" style="margin: 3rem;">
                  <div>
                    <p>Selamat menunaikan ibadah puasa #WargiDisdikJabar!😇<br>.<br>Bulan penuh berkah dan magfirah
                      telah tiba. Mari sambut Ramadan dengan hati yang bersih dan penuh keikhlasan.<br>.<br>Semoga bulan
                      suci ini membawa keberkahan, kedamaian, dan ampunan bagi kita semua.✨<br>.<br>Marhaban ya
                      Ramadan🌙</p>
                  </div><br><a href="https://disdik.jabarprov.go.id/" target="_blank" rel="noreferrer"><button
                      class="btn btn-selengkapnya">Lihat Selengkapnya <span class="icon" style="margin-left: 5px;"><svg
                          aria-hidden="true" focusable="false" data-prefix="fas" data-icon="external-link-alt"
                          class="svg-inline--fa fa-external-link-alt fa-w-16 " role="img"
                          xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">
                          <path fill="currentColor"
                            d="M432,320H400a16,16,0,0,0-16,16V448H64V128H208a16,16,0,0,0,16-16V80a16,16,0,0,0-16-16H48A48,48,0,0,0,0,112V464a48,48,0,0,0,48,48H400a48,48,0,0,0,48-48V336A16,16,0,0,0,432,320ZM488,0h-128c-21.37,0-32.05,25.91-17,41l35.73,35.73L135,320.37a24,24,0,0,0,0,34L157.67,377a24,24,0,0,0,34,0L435.28,133.32,471,169c15,15,41,4.5,41-17V24A24,24,0,0,0,488,0Z">
                          </path>
                        </svg></span></button></a>
                </div>
              </div>
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
        <img loading="lazy" alt="A serene night landscape with a tree reflected in water" class="rounded-lg w-full" height="180"
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
      <div class="container" v-if="isVisible">
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
      <div v-else class="placeholder">Loading</div>
    </section>
    <!-- End of Section 2 -->

    <section class="mb-8 container">
      <h2 class="text-xl font-bold mb-4">
        Platform yang Digunakan
      </h2>
      <div class="grid grid-cols-1 md:grid-cols-3 gap-4" v-if="isVisible">
        <div class="bg-white p-4 rounded-lg shadow-md">
          <img alt="Ethereum logo with blockchain elements" class="w-full h-auto mb-4 rounded-lg" height="200"
            src="https://storage.googleapis.com/a1aa/image/FJjJhxyfdhymaCz6sBLvGJKSVWtXxv6f7bfIy29W4DVBOcfPB.jpg"
            width="300" loading="lazy">
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
            width="300" loading="lazy">
          <h3 class="text-lg font-bold mb-2">
            Hyperledger Fabric
          </h3>
          <p>
            Hyperledger Fabric adalah kerangka kerja blockchain yang diizinkan, yang dirancang untuk aplikasi
            perusahaan. Ini menawarkan modularitas dan fleksibilitas yang unik untuk berbagai kasus penggunaan industri.
          </p>
        </div>
        <div class="bg-white p-4 rounded-lg shadow-md">
          <img loading="lazy" alt="Quorum logo with secure transaction elements" class="w-full h-auto mb-4 rounded-lg"
            height="200"
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
      <div v-else class="placeholder">Loading</div>

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
      <div class="grid grid-cols-1 md:grid-cols-3 gap-4" v-if="isVisible">
        <div class="bg-white p-4 rounded-lg shadow-md">
          <img loading="lazy" alt="News image showing a blockchain conference" class="w-full h-auto mb-4 rounded-lg"
            height="200"
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
          <img loading="lazy" alt="News image showing a new blockchain partnership"
            class="w-full h-auto mb-4 rounded-lg" height="200"
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
          <img loading="lazy" alt="News image showing a blockchain workshop" class="w-full h-auto mb-4 rounded-lg"
            height="200"
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
      <div v-else class="placeholder">Loading</div>

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
          <img loading="lazy" src="../assets/eth.jpg" alt="" srcset="">
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
                  <img loading="lazy" :src="'https://primefaces.org/cdn/primevue/images/product/' + slotProps.data.image"
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
          <img loading="lazy" src="https://readymadeui.com/readymadeui-light.svg" alt="logo" class="w-52" />
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
<style>
.image-container {
  width: 400px;
  height: 300px;
  background: #f0f0f0;
  display: flex;
  align-items: center;
  justify-content: center;
}

.placeholder {
  font-size: 18px;
  color: gray;
}
</style>