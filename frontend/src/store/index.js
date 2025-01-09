import { createStore } from "vuex";
import searchModule from "./search";
import authModule from "./authService";
const store = createStore({
  modules: {
    search: searchModule,
    authService: authModule,
  },
});

export default store;
