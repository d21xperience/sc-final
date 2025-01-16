import { createStore } from "vuex";
import searchModule from "./search";
import authModule from "./authService";
import sekolahModule from "./sekolahService";
const store = createStore({
  modules: {
    search: searchModule,
    authService: authModule,
    sekolahService: sekolahModule,
  },
});

export default store;
