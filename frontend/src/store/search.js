import axios from "axios";

const state = {
  searchCache: {},
  results: [],
  cacheTimestamps: [],
};

const mutations = {
  SET_DAPO(state, { query, results }) {
    state.searchCache[query] = results;
    state.cacheTimestamps[query] = Date.now(); // Simpan timestamp
  },
  // SET_RESULTS(state, { query, results }) {
  //   state.searchCache[query] = results;
  //   state.results = results;
  // },
};

const actions = {
  async fetchResults({ state, commit }, query) {
    const now = Date.now();
    const cacheTTL = 300000; // TTL: 5 menit

    if (
      state.searchCache[query] &&
      now - state.cacheTimestamps[query] < cacheTTL
    ) {
      console.log("Data from cache..");
      return state.searchCache[query];
    }

    try {
      console.log("Ambil data dari server");
      const response = await axios.get(`/dapo/api/getHasilPencarian`, {
        params: { keyword: query },
      });
      const results = response.data;
      // console.log(results);
      commit("SET_DAPO", { query, results });
      return results;
    } catch (error) {
      console.error("Error fetching results:", error);
      throw error; // Opsional: Lempar error agar komponen tahu ada masalah
    }
    // const response = await axios.get(`/dapo/api/getHasilPencarian`, {
    //   params: { keyword: query },
    // });

    // const results = response.data;
    // commit("SET_RESULTS", { query, results });
    // return results;
  },
  async registerAdmin({ commit }, adminData) {
    console.log("send to server...");
    // try {
    //   const response = await axios.post(
    //     "http://localhost:8080/api/admin/register",
    //     adminData
    //   );
    //   return response.data;
    // } catch (error) {
    //   console.error("Registration failed:", error.response.data);
    //   throw error.response.data;
    // }
  },
};

const getters = {
  results: (state) => state.results,
};

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters,
};
