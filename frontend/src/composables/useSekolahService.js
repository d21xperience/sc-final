import { ref } from "vue";
import { useStore } from "vuex";

export function useSekolahService(schemaname, selectedSemester) {
  const store = useStore();
  const guruList = ref([]);
  const guruTerdaftarList = ref([]);
  const siswaList = ref([]);
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

  const fetchKelas = async (kelasId = null) => {
    try {
      const payload = {
        schemaname: schemaname.value,
        semester_id: selectedSemester.value?.semesterId,
      };

      if (kelasId) {
        payload.kelas_id = kelasId;
      }
      const response = await store.dispatch(
        "sekolahService/fetchRombel",
        payload
      );
      kelasList.value = response;
    } catch (error) {
      console.error("Gagal mengambil data kelas:", error);
    }
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
  return {
    fetchGuru,
    guruList,
    fetchGuruTerdaftar,
    guruTerdaftarList,
    fetchKelas,
    kelasList,
    fetchSemester,
  };
}
