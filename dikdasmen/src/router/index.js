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
            {
              path: "setting",
              name: "setingBlockchain",
              meta: {
                title: "Blockhain setting",
                requiresAuth: true,
                role: "admin",
              },
              component: () =>
                import("../views/sc_ijazah/BlockchainSettings.vue"),
              children: [
                {
                  path: "send-krypto",
                  name: "sendKrypto",
                  component: () => import("../views/sc_ijazah/SendTrx.vue"),
                },
              ],
            },
            {
              path: "list-bcnetwork",
              name: "listBCNetwork",
              meta: {
                title: "Daftar Blockhain",
                requiresAuth: true,
                role: "admin",
              },
              component: () => import("../views/sc_ijazah/ListBCNetwork.vue"),
            },
            {
              path: "add-bcnetworks",
              name: "addBCNetworks",
              component: () => import("../views/sc_ijazah/AddBCNetwork.vue"),
            },
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
          path: "ipfs-network",
          name: "ipfs",
          component: () => import("../views/ipfs_ijazah/Main.vue"),
          children: [
            {
              path: "setting",
              name: "ipfsNetwork",
              meta: {
                title: "IPFS setting",
                requiresAuth: true,
                role: "admin",
              },
              component: () =>
                import("../views/sc_ijazah/BlockchainSettings.vue"),
              children: [
                {
                  path: "send-krypto",
                  name: "sendKrypto",
                  component: () => import("../views/sc_ijazah/SendTrx.vue"),
                },
              ],
            },
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
            //   path: "add-ipfsnetworks",
            //   name: "addIPFSNetworks",
            //   component: () => import("../views/sc_ijazah/AddBCNetwork.vue"),
            // },
            // {
            //   path: "ipfs-ijazah",
            //   name: "ipfsIjazah",
            //   component: () => import("../views/ipfs_ijazah/IPFSIjazah.vue"),
            // },
            // {
            //   path: "daftar-trx",
            //   name: "daftarTrx",
            //   component: () => import("../views/sc_ijazah/ListTrx.vue"),
            // },
          ],
        },

        {
          path: "input-ijazah",
          name: "inputIjazah",
          component: () => import("../views/dapodik/DataSiswa.vue"),
          meta: { title: "Data Ijazah", requiresAuth: true, role: "admin" },
        },

        // Data DAPODIK
        {
          path: "seting-dapodik",
          name: "syncDapodik",
          component: () => import("../views/dapodik/SetingDapodik.vue"),
          meta: { title: "Seting Dapodik", requiresAuth: true, role: "admin" },
        },
        {
          path: "data-sekolah",
          name: "dapodikSekolah",
          component: () => import("../views/dapodik/DataSekolah.vue"),
          meta: { title: "Data Sekolah", requiresAuth: true, role: "admin" },
        },
        {
          path: "data-guru",
          name: "dapodikGuru",
          component: () => import("../views/dapodik/DataGuru.vue"),
          meta: { title: "Data Guru", requiresAuth: true, role: "admin" },
        },
        {
          path: "data-siswa",
          name: "dapodikSiswa",
          component: () => import("../views/dapodik/DataSiswa.vue"),
          meta: { title: "Data Siswa", requiresAuth: true, role: "admin" },
        },
        {
          path: "data-nilai",
          name: "dapodikNilaiSiswa",
          component: () => import("../views/dapodik/DataNilai.vue"),
          meta: { title: "Data Nilai", requiresAuth: true, role: "admin" },
        },
        {
          path: "data-kelas",
          name: "dapodikKelas",
          component: () => import("../views/dapodik/DataKelas.vue"),
          meta: { title: "Data Kelas", requiresAuth: true, role: "admin" },
          // children:[
          //   {
          //     path:""
          //   }
          // ]
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
