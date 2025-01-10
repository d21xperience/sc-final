import axios from "axios";
const baseURL = "http://localhost:8080/api/v1/auth";
const state = {
  token: true,//localStorage.getItem("token") || null,
  admin: null,
  refreshToken: null,
};

const mutations = {
  setToken(state, token) {
    state.token = token;
    localStorage.setItem("token", token);
  },
  // setLoggedIn(state, loggedIn) {
  //   state.loggedIn = loggedIn;
  // },
  setAdmin(state, admin) {
    state.admin = admin;
  },
  // logout(state) {
  //   state.token = null;
  //   state.admin = null;
  //   localStorage.removeItem("token");
  // },
  setRefreshToken(state, refreshToken) {
    state.refreshToken = refreshToken;
  },
  clearAuthData(state) {
    state.token = null;
    state.refreshToken = null;
    state.user = null;
  },
};

const actions = {
  async login({ commit }, credentials) {
    try {
      const response = await axios.post(`${baseURL}/login`, credentials);
      commit("setToken", response.data.token);
      // commit("setLoggedIn", true);
      return true;
    } catch (error) {
      // console.error("Login failed:", error);
      return false;
    }
  },
  async logout({ commit }) {
    try {
      await axios.post(`${baseURL}/logout`); // Opsional: panggil endpoint logout server
    } finally {
      commit("clearAuthData");
    }
  },
  async registerAdmin({ commit }, adminData) {
    try {
      const response = await axios.post(`${baseURL}/register`, adminData);
      return response.data;
    } catch (error) {
      console.error("Registration failed:", error.response.data);
      throw error.response.data;
    }
  },
  async fetchAdmin({ commit, state }) {
    if (!state.token) return;
    try {
      const response = await axios.get("http://localhost:8080/api/admin/me", {
        headers: { Authorization: state.token },
      });
      commit("setAdmin", response.data);
    } catch (error) {
      console.error("Failed to fetch admin data:", error);
    }
  },
};

const getters = {
  isAuthenticated(state) {
    return !!state.token;
  },
  admin(state) {
    return state.admin;
  },
};

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters,
};
