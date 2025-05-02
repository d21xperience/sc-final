import axios from "axios";
// const baseURL = "http://localhost:8080/api/v1";
const api = axios.create({
  baseURL: "http://localhost:8184/api/v1", // Pastikan menggunakan protokol HTTPS
  withCredentials: true, // Untuk mengirim cookie atau credensial
  headers: {
    "Content-Type": "application/json",
    "Content-Type": "Authorization",
  },
});

const state = {
  loading: false,
  error: null,
  BCPlatformSelected:
    JSON.parse(localStorage.getItem("BCPlatformSelected")) || null,
  BCAccountActivate: null,

  // BCNETWORK: JSON.parse(localStorage.getItem("BCNETWORK")) || null,
  // BCACCOUNT: JSON.parse(localStorage.getItem("BCACCOUNT")) || null,
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
  setBCPlatformSelected(state, value) {
    state.BCPlatformSelected = value;
  },
  setBCAccountActivate(state, value) {
    state.BCAccountActivate = value;
  },
};

const actions = {
  async fetchBCPlatform({ commit }, payload) {
    try {
      // console.log(payload);
      const response = await api.get(`/sc/platform`, {
        params: {
          schemaname: payload.schemaname,
        },
      });
      // console.log(response.data.bcPlatform);
      response.data.bcPlatform.find((platform) => {
        if (platform.active) {
          commit("setBCPlatformSelected", platform);
        }
      });
      return response.data; // Mengembalikan data sekolah
    } catch (error) {
      commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
      console.error("Gagal memuat network:", error);
      return null;
    } finally {
      commit("SET_LOADING", false);
    }
  },

  async setBCPlatform({ commit }, payload) {
    // console.log("in vuex: ", payload);
    const py = {
      bc_platform: {
        id: payload.bc_platform.id,
        name: payload.bc_platform.name,
        active: payload.bc_platform.active,
      },
      schemaname: payload.schemaname,
    };

    // console.log(py);
    try {
      const response = await api.put("/sc/platform", py);
      return response.data;
    } catch (error) {
      console.error("Gagal set BC platform:", error);
      return null;
    }
    // console.log(response);
  },
  // ================================
  async updateBCPlatformSelected({ commit }, value) {
    commit("setBCPlatformSelected", value);
  },

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
  async createIjazahBC({ commit }, payload) {
    try {
      const response = await api.post(
        `/sc/ijazah-bc/create`,
        JSON.stringify(payload)
      );
      return response.data;
    } catch (error) {
      throw error;
    }
  },
  async fetchIjazahBC({ commit }, payload) {
    try {
      const response = await api.get(`/sc/ijazah-bc`, {
        params: {
          sekolah_id: payload.sekolah_id,
          tahun_ajaran_id: payload.tahun_ajaran_id,
        },
      });
      return response.data;
    } catch (error) {
      throw error;
    }
  },
  async searchIjazahBC({ commit }, payload) {
    try {
      const response = await api.get(`/sc/ijazah-bc/search`, {
        params: {
          nisn: payload.nisn,
        },
      });
      return response.data.ijazahBc;
    } catch (error) {
      throw error;
    }
  },
  // ================================

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
  getBCPlatformSelected: (state) => state.BCPlatformSelected,
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
