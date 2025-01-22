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
  async getSemester({ commit }, payload) {
    let limit = payload.limit || 10;
    let offset = payload.offset || 0;
    let semester_id = payload.semester_id || null;
    commit("SET_LOADING", true);
    commit("SET_ERROR", null);
    try {
      const response = await api.get(`/semester`, {
        params: {
          semester_id: semester_id,
          limit: limit,
          offset: offset,
        },
      });
      // commit("SET_TABELTENANT", response.data);
      return response.data; // Mengembalikan data sekolah
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
};

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters,
};
