import axios from "axios";
// const baseURL = "http://localhost:8080/api/v1";
const api = axios.create({
  baseURL: "http://localhost:8183/api/v1", // Pastikan menggunakan protokol HTTPS
  withCredentials: true, // Untuk mengirim cookie atau credensial
  headers: {
    "Content-Type": "application/json",
    "Content-Type": "Authorization",
  },
});

const state = {
  loading: false,
  error: null,
  tabelTenant: JSON.parse(localStorage.getItem("tabelTenant")) || null,
  // tabelSemester: JSON.parse(localStorage.getItem("tabelSemester")) || null,
  tabelSiswa: JSON.parse(localStorage.getItem("tabelSiswa")) || null,
  tabelGuru: JSON.parse(localStorage.getItem("tabelGuru")) || null,
  selectedSemester:
    JSON.parse(localStorage.getItem("selectedSemester")) || null,
};

const mutations = {
  SET_LOADING(state, isLoading) {
    state.loading = isLoading;
  },
  SET_ERROR(state, error) {
    state.error = error;
  },
  SET_TABELTENANT(state, tabelTenant) {
    state.tabelTenant = tabelTenant;
    localStorage.setItem("tabelTenant", JSON.stringify(tabelTenant));
  },
  SET_TABELSEMESTER(state, tabelSemester) {
    state.tabelSemester = tabelSemester;
    localStorage.setItem("tabelSemester", JSON.stringify(tabelSemester));
  },
  SET_TABELSISWA(state, tabelSiswa) {
    state.tabelSiswa = tabelSiswa;
    localStorage.setItem("tabelSiswa", JSON.stringify(tabelSiswa));
  },
  SET_TABELGURU(state, tabelGuru) {
    state.tabelGuru = tabelGuru;
    localStorage.setItem("tabelGuru", JSON.stringify(tabelGuru));
  },
  SET_SELECTEDSEMESTER(state, value) {
    state.selectedSemester = value;
    localStorage.setItem("selectedSemester", JSON.stringify(value));
  },
};

const actions = {
  // Fitur baru ceknpsn
  async fetchTabeltenant({ commit }, sekolahId) {
    try {
      const response = await api.get(
        `/sekolah/sekolah-terdaftar?sekolah_id=${sekolahId}`
      );
      commit("SET_TABELTENANT", response.data);
      // console.log(response.data);
      return response.data; // Mengembalikan data sekolah
    } catch (error) {
      commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
      console.error("Gagal mengambil data tabel tenant:", error);
      return null;
    } finally {
      commit("SET_LOADING", false);
    }
  },

  async createTabeltenant({ commit }, sekolah) {
    const payload = {
      sekolah: sekolah.sekolahData,
    };

    console.log(payload);

    commit("SET_LOADING", true);
    commit("SET_ERROR", null);
    // return;
    try {
      const response = await api.post("/sekolah/registrasi-sekolah", payload);
      console.log(response.data);
      commit("SET_TABELTENANT", response.data);
      return response.data; // Mengembalikan data sekolah
    } catch (error) {
      commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
      console.error("Gagal membuat tabel tenant:", error);
      return null;
    } finally {
      commit("SET_LOADING", false);
    }
  },
  async fetchSemester({ commit }, semester_id) {
    commit("SET_LOADING", true);
    commit("SET_ERROR", null);
    try {
      const response = await api.get(`/ss/semester`, {
        params: {
          semester_id: semester_id,
        },
      });
      // console.log(response.data.semester);
      commit("SET_TABELSEMESTER", response.data.semester);
      return true; // Mengembalikan data sekolah
    } catch (error) {
      commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
      console.error("Gagal membuat semester:", error);
      return null;
    } finally {
      commit("SET_LOADING", false);
    }
  },

  async fetchSelectedSemester({ commit }, payload) {
    commit("SET_SELECTEDSEMESTER", payload);
  },

  async fetchRombel({ commit }, payload) {
    commit("SET_LOADING", true);
    commit("SET_ERROR", null);
    try {
      const response = await api.get(`/ss/${payload.schema_name}/kelas`, {
        params: {
          semester_id: payload.semester_id,
        },
      });
      // console.log(response.data);
      // commit("SET_TABELSEMESTER", response.data.semester);
      return response.data.kelas; // Mengembalikan data sekolah
    } catch (error) {
      commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
      console.error("Gagal membuat semester:", error);
      return null;
    } finally {
      commit("SET_LOADING", false);
    }
  },
  async fetchPTK({ commit }, payload) {
    console.log("guru");

    commit("SET_LOADING", true);
    commit("SET_ERROR", null);
    try {
      const response = await api.get("/ss/ptk-terdaftar", {
        params: {
          schema_name: payload.schemaname,
          tahun_ajaran_id: payload.tahunAjaranId,
        },
      });
      console.log(response.data);
      commit("SET_TABELGURU", response.data.ptkTerdaftar);
      return response.data.ptkTerdaftar; // Mengembalikan data sekolah
    } catch (error) {
      commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
      console.error("Gagal membuat semester:", error);
      return null;
    } finally {
      commit("SET_LOADING", false);
    }
  },
  async getTemplate({ commit }, payload) {
    commit("SET_LOADING", true);
    commit("SET_ERROR", null);
    try {
      const response = await api.get(`/ss/download/template`, {
        params: {
          "template-type": "siswa",
        },
      });
      // console.log(response.data.semester);
      // commit("SET_TABELSEMESTER", response.data.semester);
      return response; // Mengembalikan data sekolah
    } catch (error) {
      commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
      console.error("Gagal membuat semester:", error);
      return null;
    } finally {
      commit("SET_LOADING", false);
    }
  },
  async fetchSekolah({ commit }, payload) {
    try {
      const response = await api.get(`/ss/${payload.schemaName}/sekolah`);
      // console.log(response.data.semester);
      // commit("SET_TABELSEMESTER", response.data.semester);
      // console.log(response.data.sekolah);
      return response.data.sekolah; // Mengembalikan data sekolah
    } catch (error) {
      commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
      console.error("Gagal membuat semester:", error);
      return null;
    } finally {
      commit("SET_LOADING", false);
    }
  },
  async fetchSiswa({ commit }, payload) {
    try {
      const response = await api.get(
        `/ss/${payload.schemaName}/anggota-kelas`,
        {
          params: {
            semester_id: payload.semesterId,
            schemaname: payload.schemaName,
          },
        }
      );
      // console.log(response.data.anggotaKelas);
      commit("SET_TABELSISWA", response.data.anggotaKelas);
      return response.data; // Mengembalikan data sekolah
    } catch (error) {
      commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
      console.error("Gagal membuat semester:", error);
      return null;
    } finally {
      commit("SET_LOADING", false);
    }
  },
  async fetchSiswaCount({ commit }, payload) {
    try {
      const response = await api.get(`/ss/dashboard/countsiswa`, {
        params: {
          schemaname: payload.schemaname,
          semester_id: payload.semesterId,
        },
      });
      // commit("SET_TABELSISWA", response.data.anggotaKelas);
      return response.data; // Mengembalikan data sekolah
    } catch (error) {
      // commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
      console.error("Gagal membuat semester:", error);
      return null;
    } finally {
      // commit("SET_LOADING", false);
    }
  },
};

const getters = {
  isLoading: (state) => state.loading,
  getError: (state) => state.error,
  getTabeltenant: (state) => state.tabelTenant,
  getSemester: (state) => state.tabelSemester,
  getSiswa: (state) => state.tabelSiswa,
  getGuru: (state) => state.tabelGuru,
  getSelectedSemester: (state) => state.selectedSemester,
};

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters,
};
