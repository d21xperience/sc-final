import axios from "axios";
// const baseURL = "http://localhost:8080/api/v1";
const api = axios.create({
  baseURL: "http://localhost:8081/api/v1/ss", // Pastikan menggunakan protokol HTTPS
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
  tabelSemester: JSON.parse(localStorage.getItem("tabelSemester")) || null,
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
};

const actions = {
  // Fitur baru ceknpsn
  async getTabeltenant({ commit }, sekolahId) {
    console.log(sekolahId);
    commit("SET_LOADING", true);
    commit("SET_ERROR", null);
    try {
      const response = await api.get(`/sekolah/sekolah-terdaftar`, {
        params: {
          sekolah_id: sekolahId,
        },
      });
      commit("SET_TABELTENANT", response.data);
      console.log(response.data);
      return true; // Mengembalikan data sekolah
    } catch (error) {
      commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
      console.error("Gagal membuat tabel tenant:", error);
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
      const response = await api.get(`/semester`, {
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
  async fetchRombel({ commit }, payload) {
    commit("SET_LOADING", true);
    commit("SET_ERROR", null);
    try {
      const response = await api.get(`/semester`, {
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
};

const getters = {
  isLoading: (state) => state.loading,
  getError: (state) => state.error,
  getTabeltenant: (state) => state.tabelTenant,
  getSemester: (state) => state.tabelSemester,
};

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters,
};
