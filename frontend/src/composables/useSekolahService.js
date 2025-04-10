import { ref } from "vue";
import { useStore } from "vuex";

export function useSekolahService(schemaName, selectedSemester) {
  const store = useStore();
  const guruList = ref([]);

  const fetchGuru = async () => {
    try {
      const existing = await store.getters["sekolahService/getGuru"];
      if (existing == null) {
        const payload = {
          tahunAjaranId: selectedSemester.value?.tahunAjaranId,
          schemaname: schemaName.value,
        };
        const response = await store.dispatch(
          "sekolahService/fetchGuru",
          payload
        );
        guruList.value = response;
      } else {
        guruList.value = existing;
      }
    } catch (error) {
      console.error("Gagal mengambil data guru:", error);
    }
  };

  const fetchKelas = async (kelasId = null) => {
    try {
      const payload = {
        schema_name: schemaName.value,
        semester_id: selectedSemester.value?.semesterId,
      };

      if (kelasId) {
        payload.kelas_id = kelasId;
      }

      const response = await store.dispatch(
        "sekolahService/fetchRombel",
        payload
      );
      return response;
    } catch (error) {
      console.error("Gagal mengambil data kelas:", error);
    }
  };

  return {
    guruList,
    fetchGuru,
    fetchKelas,
  };
}
