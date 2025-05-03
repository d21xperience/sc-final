import { createRouter, createWebHistory } from "vue-router";
import store from "@/store";
// import HomeView from "../views/HomeView.vue";
// // import Login from "../views/auth/Login.vue";
const router = createRouter({
  linkExactActiveClass: "bg-blue-100",
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/tes",
      name: "tes",
      component: () => import("../Product.vue"),
      meta: { title: "Products" },
    },
    {
      path: "/:pathMatch(.*)*", // Catch-all route
      name: "not-found",
      component: () => import("@/views/NotFoundView.vue"),
      meta: { title: "Page Not Found" },
    },
    {
      path: "/",
      name: "home",
      component: () => import("../views/HomeView.vue"),
      meta: { title: "Home" },
    },
    {
      path: "/auth",
      // component: () => import("../views/auth/Login.vue"),
      children: [
        {
          path: "login",
          name: "login",
          component: () => import("../views/auth/Login.vue"),
          meta: { title: "Login" },
        },
        {
          path: "register",
          name: "register",
          component: () => import("../views/auth/Register.vue"),
          meta: { title: "Sign Up - Register" },
        },
      ],
    },
    {
      path: "/about",
      name: "about",
      // route level code-splitting
      // this generates a separate chunk (About.[hash].js) for this route
      // which is lazy-loaded when the route is visited.
      component: () => import("../views/AboutView.vue"),
    },
    {
      // path: "/main",
      path: "/admin",
      component: () => import("../views/admin/Main.vue"),
      props: true,
      // meta: { requiresAuth: true },
      children: [
        {
          path: "",
          component: () => import("../views/admin/Dashboard.vue"),
          name: "admin",
          meta: { title: "Dashboard", requiresAuth: true },
        },
        {
          path: "profile",
          name: "profile",
          component: () => import("../views/admin/Profile.vue"),
          meta: { title: "Profile", requiresAuth: true },
        },
        {
          path: "data-users",
          name: "dataUsers",
          component: () => import("../views/admin/DataUsers.vue"),
          meta: { title: "Data User", requiresAuth: true, role: "admin" },
          children: [
            {
              path: "admin",
              name: "usersAdmin",
              component: () => import("../views/admin/UserAdminRole.vue"),
              meta: { title: "Admin role", requiresAuth: true, role: "admin" },
            },
            {
              path: "siswa",
              name: "usersSiswa",
              component: () => import("../views/admin/UserSiswaRole.vue"),
              meta: { title: "Admin role", requiresAuth: true, role: "admin" },
            },
          ],
        },

        // ----------------------------------------------------------
        // BLOCKHAIN
        // ----------------------------------------------------------
        {
          path: "blockchain",
          name: "blockchain",
          component: () => import("../views/sc_ijazah/Main.vue"),
          children: [
            // {
            //   path: "setting",
            //   name: "setingBlockchain",
            //   meta: {
            //     title: "Blockhain setting",
            //     requiresAuth: true,
            //     role: "admin",
            //   },
            //   component: () =>
            //     import("../views/sc_ijazah/BlockchainSettings.vue"),
            //   children: [
            //     {
            //       path: "send-krypto",
            //       name: "sendKrypto",
            //       component: () => import("../views/sc_ijazah/SendTrx.vue"),
            //     },
            //   ],
            // },
            // {
            //   path: "list-bcnetwork",
            //   name: "listBCNetwork",
            //   meta: {
            //     title: "Daftar Blockhain",
            //     requiresAuth: true,
            //     role: "admin",
            //   },
            //   component: () => import("../views/sc_ijazah/ListBCNetwork.vue"),
            // },
            // {
            //   path: "add-bcnetworks",
            //   name: "addBCNetworks",
            //   component: () => import("../views/sc_ijazah/AddBCNetwork.vue"),
            // },
            {
              path: "sc-ijazah",
              name: "scIjazah",
              meta: {
                title: "Input Ijazah",
                requiresAuth: true,
                role: "admin",
              },
              component: () => import("../views/sc_ijazah/SCIjazah.vue"),
            },

            {
              path: "daftar-trx",
              name: "daftarTrx",
              component: () => import("../views/sc_ijazah/ListTrx.vue"),
            },
          ],
        },
        {
          path: "settings",
          name: "settings",
          component: () => import("../views/sc_ijazah/settings/Main.vue"),
          children: [
            {
              path: "ipfs",
              name: "ipfs",
              meta: {
                title: "IPFS setting",
                requiresAuth: true,
                role: "admin",
              },
              component: () =>
                import("../views/sc_ijazah/settings/Client_IPFS.vue"),
            },
            {
              path: "ethereum",
              name: "ethereum",
              meta: {
                title: "Setting Ethereum",
                requiresAuth: true,
                role: "admin",
              },
              component: () =>
                import("../views/sc_ijazah/settings/Client_Ethereum.vue"),
            },
            {
              path: "quorum",
              name: "quorum",
              meta: {
                title: "Setting Quorum",
                requiresAuth: true,
                role: "admin",
              },
              component: () =>
                import("../views/sc_ijazah/settings/Client_Quorum.vue"),
            },
            {
              path: "hyperldeger",
              name: "hyperledger",
              meta: {
                title: "Setting Hyperledger",
                requiresAuth: true,
                role: "admin",
              },
              component: () =>
                import("../views/sc_ijazah/settings/Client_Hyperledger.vue"),
            },
            {
              path: "ijazah",
              name: "ijazah",
              meta: {
                title: "Setting ijazah",
                requiresAuth: true,
                role: "admin",
              },
              component: () =>
                import("../views/sc_ijazah/settings/Ijazah_Setting.vue"),
            },
          ],
        },

        {
          path: "input-ijazah",
          name: "inputIjazah",
          component: () => import("../views/dapodik/DataSiswa.vue"),
          meta: { title: "Data Ijazah", requiresAuth: true, role: "admin" },
        },
        {
          path: "data-dapodik",
          name: "dataDapodik",
          component: () => import("../views/dapodik/Main.vue"),
          meta: { requiresAuth: true, role: "admin" },
          children: [
            {
              path: "sekolah",
              name: "dapodikSekolah",
              component: () => import("../views/dapodik/DataSekolah.vue"),
              meta: {
                title: "Data Sekolah",
                requiresAuth: true,
                role: "admin",
              },
            },
            {
              path: "guru",
              name: "dapodikGuru",
              // component: () => import("../views/dapodik/DataGuru.vue"),
              meta: { title: "Data Guru", requiresAuth: true, role: "admin" },
              children: [
                {
                  path: "",
                  name: "readGuru",
                  component: () =>
                    import("../views/dapodik/data_guru/ReadGuru.vue"),
                },
                {
                  path: "input-guru",
                  name: "inputGuru",
                  meta: { disableSelect: true, title: "Tambah Guru" },
                  component: () =>
                    import("../views/dapodik/data_guru/AddGuru.vue"),
                },
                {
                  path: "edit-guru",
                  name: "editGuru",
                  meta: { disableSelect: true, title: "Edit Guru" },
                  component: () =>
                    import("../views/dapodik/data_guru/AddGuru.vue"),
                  // props: true,
                },
              ],
            },
            {
              path: "mapel",
              name: "dapodikMapel",
              // component: () => import("../views/dapodik/DataGuru.vue"),
              meta: {
                title: "Data Mata pelajaran",
                requiresAuth: true,
                role: "admin",
              },
              children: [
                {
                  path: "",
                  name: "readMapel",
                  component: () =>
                    import("../views/dapodik/data_matapelajaran/ReadMapel.vue"),
                },
                {
                  path: "input-matapelajaran",
                  name: "inputMapel",
                  component: () =>
                    import("../views/dapodik/data_matapelajaran/AddMapel.vue"),
                },
              ],
            },
            {
              path: "siswa",
              name: "dapodikSiswa",
              meta: { title: "Data Siswa", requiresAuth: true, role: "admin" },
              children: [
                {
                  path: "",
                  name: "readSiswa",
                  component: () =>
                    import("../views/dapodik/data_siswa/ReadSiswa.vue"),
                },
                {
                  path: "input",
                  name: "inputSiswa",
                  meta: { disableSelect: true },
                  component: () =>
                    import("../views/dapodik/data_siswa/AddSiswa.vue"),
                },
                {
                  path: "edit-siswa/:id",
                  name: "editSiswa",
                  meta: { disableSelect: true },
                  component: () =>
                    import("../views/dapodik/data_siswa/AddSiswa.vue"),
                },
              ],
            },
            {
              path: "data-nilai",
              name: "dapodikNilaiSiswa",
              component: () => import("../views/dapodik/DataNilai.vue"),
              meta: {
                breadcrumb: "Parent Page",
                title: "Data Nilai",
                requiresAuth: true,
                role: "admin",
              },
            },
            {
              path: "data-kelas",
              name: "dapodikKelas",
              meta: { title: "Data Kelas", requiresAuth: true, role: "admin" },
              children: [
                {
                  path: "",
                  name: "readKelas",
                  meta: { title: "Data Kelas" },
                  component: () =>
                    import("../views/dapodik/data_kelas/ReadKelas.vue"),
                },
                {
                  path: "tambah-kelas",
                  name: "addKelas",
                  meta: { disableSelect: true, title: "Tambah Kelas" },
                  component: () =>
                    import("../views/dapodik/data_kelas/AddKelas.vue"),
                },
                {
                  path: "edit-kelas",
                  name: "editKelas",
                  meta: { disableSelect: true, title: "Edit Kelas" },
                  component: () =>
                    import("../views/dapodik/data_kelas/AddKelas.vue"),
                  // props: true,
                },
              ],
            },
          ],
        },

        // Data akademik siswa
        {
          path: "ketuntasan-rapor",
          name: "ketuntasanRapor",
          component: () => import("../views/data_akademik/KetuntasanRapor.vue"),
          meta: {
            title: "Ketuntasan Rapor",
            requiresAuth: true,
            role: "siswa",
          },
        },
        {
          path: "data-ijazah",
          name: "dataIjazah",
          component: () => import("../views/data_akademik/DataIjazah.vue"),
          meta: { title: "Data Ijazah", requiresAuth: true, role: "siswa" },
        },
      ],
    },
  ],
});

// router.beforeEach((to, from, next) => {
//   document.title = to.meta.title || "Default Title";
//   const isAuthenticated = store.getters["authService/isAuthenticated"];
//   if (to.meta.requiresAuth && !isAuthenticated) {
//     // console.log(to);
//     next({ name: "login" }); // Redirect ke Login jika tidak login
//   }
//   // Halaman yang tidak boleh diakses oleh user yang sudah login (seperti login dan register)
//   else if ((to.name === "login" || to.name === "register") && isAuthenticated) {
//     next({ name: "home" }); // Redirect ke dashboard jika sudah login
//   } else {
//     next();
//   }
// });

router.beforeEach(async (to, from, next) => {
  document.title = to.meta.title || "Verifikasi Ijazah App";

  // Ambil informasi autentikasi dan peran pengguna dari store
  const isAuthenticated = await store.getters["authService/isAuthenticated"];
  const userRole = await store.getters["authService/userRole"]; // 'admin' atau 'siswa'
  // console.log(userRole);
  if (to.meta.requiresAuth && !isAuthenticated) {
    // Redirect ke Login jika tidak login
    next({ name: "login" });
  } else if (
    (to.name === "login" || to.name === "register") &&
    isAuthenticated
  ) {
    // Redirect ke dashboard jika sudah login
    next({ name: "home" });
  } else {
    // Periksa peran pengguna untuk akses khusus
    if (to.meta.role && to.meta.role !== userRole) {
      // Jika pengguna tidak memiliki akses ke rute tertentu
      if (userRole === "admin") {
        next();
        // next({ name: "adminDashboard" }); // Redirect ke dashboard admin
      } else {
        // next({ name: "studentDashboard" }); // Redirect ke dashboard siswa
      }
    } else {
      next();
    }
  }
});
export default router;
