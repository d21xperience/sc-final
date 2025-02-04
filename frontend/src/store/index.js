import { createStore } from "vuex";
import searchModule from "./search";
import authModule from "./authService";
import sekolahModule from "./sekolahService";
import scService from "./scService";
const store = createStore({
  modules: {
    search: searchModule,
    authService: authModule,
    sekolahService: sekolahModule,
    scService: scService,
  },
});

export default store;
// import { createStore } from "vuex";
// import searchModule from "./search";
// import authModule from "./authService";
// import sekolahModule from "./sekolahService";
// import scService from "./scService";

// export const store = createStore({
//   modules: {
//     search: searchModule,
//     authService: authModule,
//     sekolahService: sekolahModule,
//     scService: scService, // ✅ Pastikan konsisten
//   },
// });
