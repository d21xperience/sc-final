import axios from "axios";
// const baseURL = "http://localhost:8080/api/v1";
const api = axios.create({
  baseURL: "http://localhost:8082/api/v1", // Pastikan menggunakan protokol HTTPS
  withCredentials: true, // Untuk mengirim cookie atau credensial
  headers: {
    "Content-Type": "application/json",
    "Content-Type": "Authorization",
  },
});

const state = {
  loading: false,
  error: null,
  BCPlatformActivate: {},
  BCAccountActivate: {},
  BCNETWORK: JSON.parse(localStorage.getItem("BCNETWORK")) || null,
  BCACCOUNT: JSON.parse(localStorage.getItem("BCACCOUNT")) || null,
};

const mutations = {
  SET_LOADING(state, value) {
    state.loading = value;
  },
  SET_ERROR(state, value) {
    state.error = value;
  },
  SET_BCNETWORK(state, value) {
    state.BCNETWORK = value;
    localStorage.setItem("BCNETWORK", JSON.stringify(value));
  },
  SET_BCACCOUNT(state, value) {
    state.BCACCOUNT = value;
    localStorage.setItem("BCACCOUNT", JSON.stringify(value));
  },
  setBCPlatformActivate(state, value) {
    state.BCPlatformActivate = value;
  },
  setBCAccountActivate(state, value) {
    state.BCAccountActivate = value;
  },
};

const actions = {
  // Fitur baru ceknpsn
  async fetchBlockchainNetworks({ commit }) {
    // console.log(sekolahId);
    commit("SET_LOADING", true);
    commit("SET_ERROR", null);
    try {
      const response = await api.get(`/blockchain/list`);
      commit("SET_BCNETWORK", response.data.network);
      return response.data; // Mengembalikan data sekolah
    } catch (error) {
      commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
      console.error("Gagal memuat network:", error);
      return null;
    } finally {
      commit("SET_LOADING", false);
    }
  },

  // ==========================akun=========================
  // ===================================================
  async fetchBCAccount({ commit }, payload) {
    commit("SET_LOADING", true);
    commit("SET_ERROR", null);
    // return;
    try {
      const response = await api.get("/blockchainaccount/list", {
        params: {
          schemaname: payload.schemaname,
          user_id: payload.user_id,
          network_id: payload.network_id,
        },
      });
      // console.log(response.data);
      commit("SET_BCACCOUNT", response.data.blockchainaccounts);
      return response.data; // Mengembalikan data sekolah
    } catch (error) {
      commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
      // console.error("Gagal mengambil akun blockchain:", error);
      return error.response?.data;
    } finally {
      commit("SET_LOADING", false);
    }
  },
  async createBCAccount({ commit }, payload) {
    let network_name = payload.network_name || 0;
    let user_id = payload.user_id || 0;
    let password = payload.password || 0;
    let schemaname = payload.schemaname || null;
    commit("SET_LOADING", true);
    commit("SET_ERROR", null);
    try {
      const response = await api.post(`/blockchainaccount/create`, {
        params: {
          schemaname: schemaname,
          password: password,
          user_id: user_id,
          network_name: network_name,
        },
      });
      // commit("SET_BCNETWORK", response.data);
      return response.data; // Mengembalikan data sekolah
    } catch (error) {
      commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
      console.error("Gagal membuat akun:", error);
      return null;
    } finally {
      commit("SET_LOADING", false);
    }
  },
  async importBCAccount({ commit }, payload) {
    commit("SET_LOADING", true);
    commit("SET_ERROR", null);
    console.log("Payload yang dikirim:", JSON.stringify(payload, null, 2));

    try {
      const response = await api.post(
        `/blockchainaccount/import`,
        JSON.stringify(payload, null, 2),
        {
          headers: {
            "Content-Type": "application/json",
          },
        }
      );
      // commit("SET_BCNETWORK", response.data);
      return response.data; // Mengembalikan data sekolah
    } catch (error) {
      commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
      console.error("Gagal membuat akun:", error);
      return null;
    } finally {
      commit("SET_LOADING", false);
    }
  },
  async updateBCAccount({ commit }, payload) {
    let network_name = payload.network_name || 0;
    let user_id = payload.user_id || 0;
    let password = payload.password || 0;
    let schemaname = payload.schemaname || null;
    commit("SET_LOADING", true);
    commit("SET_ERROR", null);
    try {
      const response = await api.get(`/blockchainaccount/create`, {
        params: {
          schemaname: schemaname,
          password: password,
          user_id: user_id,
          network_name: network_name,
        },
      });
      // commit("SET_BCNETWORK", response.data);
      return response.data; // Mengembalikan data sekolah
    } catch (error) {
      commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
      console.error("Gagal membuat akun:", error);
      return null;
    } finally {
      commit("SET_LOADING", false);
    }
  },

  // ================================
  // ======== SMART CONTRACT =======
  async deployIjazahContract({ commit }, payload) {
    commit("SET_LOADING", true);
    commit("SET_ERROR", null);
    console.log("Payload yang dikirim:", JSON.stringify(payload, null, 2));

    try {
      const response = await api.post(
        `/contract/deploy`,
        JSON.stringify(payload, null, 2),
        {
          headers: {
            "Content-Type": "application/json",
          },
        }
      );
      // commit("SET_BCNETWORK", response.data);
      return response.data; // Mengembalikan data sekolah
    } catch (error) {
      commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
      console.error("Gagal membuat akun:", error);
      return null;
    } finally {
      commit("SET_LOADING", false);
    }
  },
  // ================================

  // ================================
  async updateBCPlatformActivate({ commit }, value) {
    commit("setBCPlatformActivate", value);
  },
  async updateBCAccountActivate({ commit }, value) {
    commit("setBCAccountActivate", value);
  },
};

// ==========================================
// ===============GETTERS=================
const getters = {
  isLoading: (state) => state.loading,
  getError: (state) => state.error,
  getBCNETWORK: (state) => state.BCNETWORK,
  getBCPlatformActivate: (state) => state.BCPlatformActivate,
  getBCAccount: (state) => state.BCACCOUNT,
  getBCAccountActivate: (state) => state.BCAccountActivate,
};

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters,
};
