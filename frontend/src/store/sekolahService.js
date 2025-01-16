import axios from "axios";
// const baseURL = "http://localhost:8080/api/v1";
const api = axios.create({
  baseURL: "http://localhost:8081/api/v1", // Pastikan menggunakan protokol HTTPS
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
  //   setUser(state, user) {
  //     state.user = user;
  //     localStorage.setItem("user", JSON.stringify(user)); // Simpan user ke localStorage
  //   },
  //   setUserRole(state, userRole) {
  //     state.userRole = userRole;
  //     localStorage.setItem("userRole", userRole);
  //   },
  //   setUserProfile(state, userProfile) {
  //     state.userProfile = userProfile;
  //     localStorage.setItem("userProfile", JSON.stringify(userProfile));
  //   },

  //   setToken(state, token) {
  //     state.token = token;
  //     localStorage.setItem("token", token);
  //   },
  //   setRefreshToken(state, refreshToken) {
  //     state.refreshToken = refreshToken;
  //   },
  //   clearAuthData(state) {
  //     state.token = null;
  //     state.refreshToken = null;
  //     state.user = null;
  //     localStorage.removeItem("token"); // Hapus token saat logout
  //     localStorage.removeItem("userRole"); // Hapus userRole saat logout
  //     localStorage.removeItem("userProfile"); // Hapus userProfile saat logout
  //     localStorage.removeItem("user");
  //     localStorage.removeItem("sekolah");
  //   },
  //   SET_SEKOLAH(state, sekolah) {
  //     state.sekolah = sekolah;
  //     localStorage.setItem("sekolah", JSON.stringify(sekolah));
  //   },
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
