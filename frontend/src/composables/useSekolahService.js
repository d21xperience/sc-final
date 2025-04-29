import { ref } from "vue";
import { useStore } from "vuex";

export function useSekolahService(schemaname, selectedSemester) {
  const store = useStore();
  const guruList = ref([]);
  const guruTerdaftarList = ref([]);
  // const siswaList = ref([]);
  const kelasList = ref([]);

  // console.log("schemaname di composable:", schemaname.value);
  // console.log("selectedSemester di composable:", selectedSemester.value);
  const fetchTabeltenant = async () => {};
  const fetchGuruTerdaftar = async (ptkId = null) => {
    try {
      const payload = {
        tahunAjaranId: selectedSemester.value?.tahunAjaranId,
        schemaname: schemaname.value,
      };

      if (ptkId) {
        payload.ptk_id = ptkId;
      }
      // console.log(payload)
      const response = await store.dispatch(
        "sekolahService/fetchPTKTerdaftar",
        payload
      );
      guruTerdaftarList.value = response;
    } catch (error) {
      console.error("Gagal mengambil data guru:", error);
    }
  };
  const fetchGuru = async (ptkId = null) => {
    try {
      const payload = {
        schemaname: schemaname.value,
      };

      if (ptkId) {
        payload.ptk_id = ptkId;
      }
      const response = await store.dispatch(
        "sekolahService/fetchGuru",
        payload
      );
      guruList.value = response;
    } catch (error) {
      console.error("Gagal mengambil data guru:", error);
    }
  };

  const fetchKelas = async (kelasId = null, tingkatPendidikanId = null) => {
    try {
      const payload = {
        schemaname: schemaname.value,
        semester_id: selectedSemester.value?.semesterId,
      };

      if (kelasId) {
        payload.kelas_id = kelasId;
      }
      if (tingkatPendidikanId) {
        payload.tingkat_pendidikan_id = tingkatPendidikanId;
      }
      const response = await store.dispatch(
        "sekolahService/fetchRombel",
        payload
      );
      // console.log(response);
      kelasList.value = response.kelas;
      return response.kelas;
    } catch (error) {
      console.error("Gagal mengambil data kelas:", error);
    }
  };
  const fetchAnggotaKelas = async (anggotaRombelId = null) => {
    try {
      const payload = {
        schemaname: schemaname.value,
        semester_id: selectedSemester.value?.semesterId,
      };

      if (anggotaRombelId) {
        payload.anggota_rombel_id = anggotaRombelId;
      }
      const response = await store.dispatch(
        "sekolahService/fetchSiswaAktif",
        payload
      );
      return response;
    } catch (error) {
      console.error("Gagal mengambil data kelas:", error);
    }
  };
  const fetchSiswaAktif = async () => {
    const payload = {
      // page: 1,
      semesterId: selectedSemester.value.semesterId,
      schemaname: schemaname.value,
    };
    // console.log(payload)
    const results = await store.dispatch(
      "sekolahService/fetchSiswaAktif",
      payload
    );
    // console.log(results)
    return results;
    // siswaList.value = results;
    // results.forEach(item => {
    //     siswa.value.push(item)
    // });
  };
  const fetchSemester = async () => {
    try {
      const results = await store.dispatch("sekolahService/fetchSemester");
      if (results) {
        semester.value = store.getters["sekolahService/getSemester"];
        // Cek apakah di vuex ada nilai
        selectedSemester.value = await store.getters[
          "sekolahService/getSelectedSemester"
        ];
        if (selectedSemester.value == null) {
          // jika tidak ada, ambil semester terbaru berdasarkan ID terbesar
          selectedSemester.value = semester.value.reduce((latest, current) =>
            current.semesterId > latest.semesterId ? current : latest
          );
        }
        // tetapkan semester yang dipilih
        store.commit(
          "sekolahService/SET_SELECTEDSEMESTER",
          selectedSemester.value
        );
      }
    } catch (error) {}
  };

  const fetchNilaiSiswa = async (pesertaDidikId = null) => {
    const payload = {
      // page: 1,
      semesterId: selectedSemester.value.semesterId,
      schemaname: schemaname.value,
    };
    // console.log(payload);
    if (pesertaDidikId) {
      payload.peserta_didik_id = pesertaDidikId;
    }
    const results = await store.dispatch(
      "sekolahService/fetchNilaiSiswa",
      payload
    );
    // console.log(results)
    return results;
    // siswaList.value = results;
    // results.forEach(item => {
    //     siswa.value.push(item)
    // });
  };

  const fetchTingkat = async () => {
    try {
      let response = await store.getters["sekolahService/getRefTingkat"];
      if (!response) {
        const payload = {
          jenjang_pendidikan_id: await store.getters[
            "sekolahService/getSekolah"
          ]?.jenjangPendidikanId, //sekolah.value?.jenjangPendidikanId
        };
        response = await store.dispatch(
          "sekolahService/fetchTingkatPendidikan",
          payload
        );
        return response;
      }
    } catch (error) {
      console.log(error);
    }

    // console.log(response);
  };
  return {
    fetchGuru,
    guruList,
    fetchGuruTerdaftar,
    guruTerdaftarList,
    fetchKelas,
    // fetchSiswa,
    kelasList,
    fetchSemester,
    fetchSiswaAktif,
    fetchNilaiSiswa,
    fetchTingkat,
  };
}
