import axios from "axios";
// const baseURL = "http://localhost:8080/api/v1";
const api = axios.create({
  baseURL: "http://localhost:8080/api/v1", // Pastikan menggunakan protokol HTTPS
  withCredentials: true, // Untuk mengirim cookie atau credensial
  headers: {
    "Content-Type": "application/json",
    "Content-Type": "Authorization",
  },
});

const state = {
  token: localStorage.getItem("token") || null,
  userRole: localStorage.getItem("userRole") || null,
  user: localStorage.getItem("user")
    ? JSON.parse(localStorage.getItem("user"))
    : null,
  userProfile: JSON.parse(localStorage.getItem("userProfile")) || null, // Ambil dari localStorage
  sekolah: JSON.parse(localStorage.getItem("sekolah")) || null, // Ambil dari localStorage
  refreshToken: null,
  loading: false,
  error: null,
};

const mutations = {
  setUser(state, user) {
    state.user = user;
    localStorage.setItem("user", JSON.stringify(user)); // Simpan user ke localStorage
  },
  setUserRole(state, userRole) {
    state.userRole = userRole;
    localStorage.setItem("userRole", userRole);
  },
  setUserProfile(state, userProfile) {
    state.userProfile = userProfile;
    localStorage.setItem("userProfile", JSON.stringify(userProfile));
  },

  setToken(state, token) {
    state.token = token;
    localStorage.setItem("token", token);
  },
  setRefreshToken(state, refreshToken) {
    state.refreshToken = refreshToken;
  },
  clearAuthData(state) {
    state.token = null;
    state.refreshToken = null;
    state.user = null;
    localStorage.removeItem("token"); // Hapus token saat logout
    localStorage.removeItem("userRole"); // Hapus userRole saat logout
    localStorage.removeItem("userProfile"); // Hapus userProfile saat logout
    localStorage.removeItem("user");
    localStorage.removeItem("sekolah");
    localStorage.removeItem("tabelTenant");
  },
  SET_SEKOLAH(state, sekolah) {
    state.sekolah = sekolah;
    localStorage.setItem("sekolah", JSON.stringify(sekolah));
  },
  // setSekolah(state, sekolah) {
  //   state.sekolah = sekolah;
  //   localStorage.setItem("sekolah", JSON.stringify(sekolah));
  // },
  SET_LOADING(state, isLoading) {
    state.loading = isLoading;
  },
  SET_ERROR(state, error) {
    state.error = error;
  },
};

const actions = {
  async login({ commit }, credentials) {
    try {
      commit("SET_LOADING", true);
      commit("SET_ERROR", null);
      const response = await api.post("/auth/login", credentials);
      if (response.data.ok) {
        commit("setToken", response.data.token);
        commit("setUser", response.data.user);
        commit("setUserRole", response.data.user.role);
        return true;
      }
    } catch (error) {
      commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
      console.error("Login failed:", error.response ? error.response.data : error.message);
      return error; // Menghindari mengembalikan error langsung
      // return false; // Menghindari mengembalikan error langsung
    } finally {
      commit("SET_LOADING", false);
    }
  },
  async logout({ commit }) {
    try {
      console.log("clear auth data");
      // await axios.post("/auth/logout"); // Opsional: panggil endpoint logout server
      commit("clearAuthData");
    } finally {
      commit("clearAuthData");
    }
  },
  async registerAdmin({ commit }, payload) {
    try {
      const response = await api.post("/auth/register", payload);
      // console.log(response);

      if (response.data.ok) {
        commit("setToken", response.data.token);
        // Simpan informasi pengguna setelah login
        commit("setUser", response.data.user);
        commit("setUserRole", response.data.user.role);
      }
      return response.data;
    } catch (error) {
      console.error("Registration failed:", error.response.data);
      throw error.response.data;
    }
  },
  // Fitur baru ceknpsn
  async ceknpsn({ commit }, npsn) {
    commit("SET_LOADING", true);
    commit("SET_ERROR", null);
    try {
      const response = await api.get(`/sekolah`, {
        params: {
          npsn: npsn,
        },
      });
      commit("SET_SEKOLAH", response.data);
      return response.data; // Mengembalikan data sekolah
    } catch (error) {
      commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
      console.error("Gagal mengambil data NPSN:", error);
      return null;
    } finally {
      commit("SET_LOADING", false);
    }
  },
  async getSekolahByID({ commit }, sekolahId) {
    commit("SET_LOADING", true);
    commit("SET_ERROR", null);
    try {
      const response = await api.get(`/sekolah`, {
        params: {
          sekolah_id: sekolahId,
        },
      });
      // console.log(response);
      commit("SET_SEKOLAH", response.data);
      // return response.data; // Mengembalikan data sekolah
    } catch (error) {
      commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
      console.error("Gagal mengambil data ID Sekola:", error);
      return null;
    } finally {
      commit("SET_LOADING", false);
    }
  },
  //  Ambil Profil Pengguna
  async getUserProfile({ commit }, userID) {
    try {
      const response = await api.get(`/user/${userID}/profile`); // Pastikan endpoint benar
      commit("setUserProfile", response.data.userProfile);
      return response.data;
    } catch (error) {
      console.error("Gagal mengambil profil pengguna:", error);
      return null;
    }
  },

  // Update Profil Pengguna
  async updateUserProfile({ commit }, updatedProfile) {
    console.log("Mengirim data ke server:", updatedProfile);
    try {
      const response = await api.put(`/user/${updatedProfile.userId}/profile`, {
        user_id: updatedProfile.userId, // Sesuai dengan .proto
        user_profile: {
          // Harus dikirim dalam objek "user_profile"
          nama: updatedProfile.nama,
          jk: updatedProfile.jk,
          phone: updatedProfile.phone,
          tpt_lahir: updatedProfile.tptLahir,
          alamat_jalan: updatedProfile.alamatJalan,
          kota_kab: updatedProfile.kotaKab,
          prov: updatedProfile.prov,
          kode_pos: updatedProfile.kodePos,
          nama_ayah: updatedProfile.namaAyah,
          nama_ibu: updatedProfile.namaIbu,
        },
      });

      console.log("Response dari server:", response.data);

      if (response.data.status === "success") {
        commit("setUserProfile", response.data.user_profile);
        return response.data;
      } else {
        console.error("Gagal memperbarui profil:", response.data);
        return null;
      }
    } catch (error) {
      console.error(
        "Gagal memperbarui profil pengguna:",
        error.response?.data || error.message
      );
      return null;
    }
  },
  // Upload Foto Profil Pengguna
  async uploadUserPhotoProfile({ commit }, file) {
    try {
      const formData = new FormData();
      formData.append("photo", file);

      const response = await api.post("/user/profile/photo", formData, {
        headers: {
          "Content-Type": "multipart/form-data",
        },
      });

      commit("setUser", response.data);
      return response.data;
    } catch (error) {
      console.error("Gagal mengunggah foto profil:", error);
      return null;
    }
  },
};

const getters = {
  isAuthenticated(state) {
    return !!state.token;
  },
  userRole(state) {
    return state.userRole;
  },
  getSekolah(state) {
    return state.sekolah;
  },
  getUserProfile(state) {
    const userData = { ...state.user, ...state.userProfile };
    return userData;
  },
  isLoading: (state) => state.loading,
  getError: (state) => state.error,
};

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters,
};
