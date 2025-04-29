import axios from "axios";
// const baseURL = "http://localhost:8080/api/v1";
const api = axios.create({
  baseURL: "http://localhost:8183/api/v1", // Pastikan menggunakan protokol HTTPS
  withCredentials: true, // Untuk mengirim cookie atau credensial
  headers: {
    "Content-Type": "application/json",
    "Content-Type": "Authorization",
  },
  timeout: 20000, // 20 detik
});

const state = {
  // tabelKelas: JSON.parse(localStorage.getItem("tabelKelas")) || null,
  tabelKelas: (() => {
    try {
      const data = localStorage.getItem("tabelKelas");
      return data ? JSON.parse(data) : [];
    } catch (e) {
      console.error("Gagal parse tabelKelas dari localStorage", e);
      return [];
    }
  })(),
  loading: false,
  error: null,
  tabelTenant: JSON.parse(localStorage.getItem("tabelTenant")) || null,
  tabelSemester: JSON.parse(localStorage.getItem("tabelSemester")) || null,
  tabelSekolah: JSON.parse(localStorage.getItem("tabelSekolah")) || null,
  tabelSiswaAktif: JSON.parse(localStorage.getItem("tabelSiswaAktif")) || null,
  tabelGuru: JSON.parse(localStorage.getItem("tabelGuru")) || null,
  tabelPTKTerdaftar:
    JSON.parse(localStorage.getItem("tabelPTKTerdaftar")) || null,
  tabelTingkatPendidikan:
    JSON.parse(localStorage.getItem("tabelTingkatPendidikan")) || null,
  tabelKurikulum: JSON.parse(localStorage.getItem("tabelKurikulum")) || null,
  selectedSemester:
    JSON.parse(localStorage.getItem("selectedSemester")) || null,
  tabelAnggotaKelas:
    JSON.parse(localStorage.getItem("tabelAnggotaKelas")) || [],
  tabelMapel: JSON.parse(localStorage.getItem("tabelMapel")) || null,

  tabelNilaiakhir: JSON.parse(localStorage.getItem("tabelNilaiakhir")) || null,
  tabelTahunAjaran: JSON.parse(localStorage.getItem("tabelTahunAjaran")) || [],
  selectedTahunAjaran:
    JSON.parse(localStorage.getItem("selectedTahunAjaran")) || [],
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
  SET_TABELSEMESTER(state, value) {
    state.tabelSemester = value;
    localStorage.setItem("tabelSemester", JSON.stringify(value));
  },
  SET_TABELSEKOLAH(state, tabelSekolah) {
    state.tabelSekolah = tabelSekolah;
    localStorage.setItem("tabelSekolah", JSON.stringify(tabelSekolah));
  },
  SET_TABELSISWAAKTIF(state, tabelSiswaAktif) {
    state.tabelSiswaAktif = tabelSiswaAktif;
    localStorage.setItem("tabelSiswaAktif", JSON.stringify(tabelSiswaAktif));
  },
  SET_TABELGURU(state, tabelGuru) {
    state.tabelGuru = tabelGuru;
    localStorage.setItem("tabelGuru", JSON.stringify(tabelGuru));
  },
  SET_TABELPTKTERDAFTAR(state, value) {
    state.tabelPTKTerdaftar = value;
    localStorage.setItem("tabelPTKTerdaftar", JSON.stringify(value));
  },
  SET_TABELTINGKATPENDIDIKAN(state, value) {
    state.tabelTingkatPendidikan = value;
    localStorage.setItem("tabelTingkatPendidikan", JSON.stringify(value));
  },
  SET_TABELKURIKULUM(state, tabelKurikulum) {
    state.tabelKurikulum = tabelKurikulum;
    localStorage.setItem("tabelKurikulum", JSON.stringify(tabelKurikulum));
  },
  SET_SELECTEDSEMESTER(state, value) {
    state.selectedSemester = value;
    localStorage.setItem("selectedSemester", JSON.stringify(value));
  },
  SET_TABELKELAS(state, value) {
    state.tabelMapel = value;
    localStorage.setItem("tabelKelas", JSON.stringify(value));
  },
  SET_TABELANGGOTAKELAS(state, value) {
    state.selectedSemester = value;
    localStorage.setItem("tabelAnggotaKelas", JSON.stringify(value));
  },
  SET_TABELMAPEL(state, value) {
    state.tabelMapel = value;
    localStorage.setItem("tabelMapel", JSON.stringify(value));
  },
  SET_TABELNILAIAKHIR(state, value) {
    state.tabelNilaiakhir = value;
    localStorage.setItem("tabelNilaiakhir", JSON.stringify(value));
  },
  SET_TABELTAHUNAJARAN(state, value) {
    state.tabelTahunAjaran = value;
    localStorage.setItem("tabelTahunAjaran", JSON.stringify(value));
  },
  SET_SELECTEDTAHUNAJARAN(state, value) {
    state.selectedTahunAjaran = value;
    localStorage.setItem("selectedTahunAjaran", JSON.stringify(value));
  },
};

const actions = {
  // ================================================
  // Semester
  // ================================================
  async fetchSemester({ commit }, semester_id) {
    // commit("SET_LOADING", true);
    // commit("SET_ERROR", null);
    try {
      const response = await api.get(`/ss/semester`, {
        params: {
          semester_id: semester_id,
        },
      });
      // console.log(response.data.semester);
      commit("SET_TABELSEMESTER", response.data.semester);
      return true; // Mengembalikan data sekolah
    } catch (error) {
      commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
      console.error("Gagal membuat semester:", error);
      return null;
    } finally {
      commit("SET_LOADING", false);
    }
  },

  async fetchSelectedSemester({ commit }, payload) {
    commit("SET_SELECTEDSEMESTER", payload);
  },
  // ================================================

  async fetchRombel({ commit }, payload) {
    try {
      const response = await api.get(`/ss/${payload.schemaname}/kelas`, {
        params: {
          semester_id: payload.semester_id,
          kelas_id: payload.kelas_id,
          tingkat_pendidikan_id: payload.tingkat_pendidikan_id,
        },
      });
      const data = {
        semesterId: payload.semester_id,
        kelas: response.data.kelas,
      };
      commit("SET_TABELKELAS", data.kelas);
      return data; // Mengembalikan data sekolah
    } catch (error) {
      console.error("Gagal mengambil data rombel:", error);
      return null;
    } finally {
      commit("SET_LOADING", false);
    }
  },

  // ================================================
  // Service Guru
  // ================================================
  async fetchGuru({ commit }, payload) {
    try {
      const response = await api.get("/ss/ptk", {
        params: {
          schemaname: payload.schemaname,
          ptk_id: payload.ptk_id,
        },
      });
      // console.log(response.data);
      // commit("SET_TABELGURU", response.data.PTK);
      return response.data.PTK; // Mengembalikan data sekolah
    } catch (error) {
      // commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
      console.error("Gagal membuat guru:", error);
      return null;
    } finally {
      // commit("SET_LOADING", false);
    }
  },
  async fetchPTKTerdaftar({ commit }, payload) {
    // console.log("guru");
    try {
      const response = await api.get("/ss/ptk-terdaftar", {
        params: {
          schemaname: payload.schemaname,
          tahun_ajaran_id: payload.tahunAjaranId,
        },
      });
      // console.log(response.data);
      commit("SET_TABELPTKTERDAFTAR", response.data.ptkTerdaftar);
      return response.data.ptkTerdaftar; // Mengembalikan data sekolah
    } catch (error) {
      // commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
      // console.error("Gagal membuat semester:", error);
      return null;
    } finally {
      commit("SET_LOADING", false);
    }
  },
  async fetchGuruCount({ commit }, payload) {
    const response = await api.get("/ss/dashboard/countguru", {
      params: {
        schemaname: payload.schemaname,
        tahun_ajaran_id: payload.tahun_ajaran_id,
      },
    });
    return response.data.countGuru;
  },
  async saveGuru({ commit }, payload) {
    try {
      const response = await api.post("/ss/ptk/create", payload);
      return response.data.status;
    } catch (error) {
      console.log(error);
    }
  },
  async savePTKTerdaftar({ commit }, payload) {
    try {
      const response = await api.post(
        `/ss/${payload.schemaname}/ptk-terdaftar/create`,
        payload
      );
      // console.log(response);
      return response.data.status;
    } catch (error) {
      console.log(error);
    }
  },

  // =============================================

  async getTemplate({ commit }, payload) {
    commit("SET_LOADING", true);
    commit("SET_ERROR", null);
    try {
      const response = await api.get(`/ss/download/template`, {
        params: {
          "template-type": "siswa",
        },
      });
      // console.log(response.data.semester);
      // commit("SET_TABELSEMESTER", response.data.semester);
      return response; // Mengembalikan data sekolah
    } catch (error) {
      commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
      console.error("Gagal membuat semester:", error);
      return null;
    } finally {
      commit("SET_LOADING", false);
    }
  },
  async fetchSekolah({ commit }, payload) {
    try {
      const response = await api.get(`/ss/${payload.schemaname}/sekolah`);
      commit("SET_TABELSEKOLAH", response.data.sekolah);
      return response.data.sekolah; // Mengembalikan data sekolah
    } catch (error) {
      commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
      console.error("Gagal membuat semester:", error);
      return null;
    } finally {
      commit("SET_LOADING", false);
    }
  },
  async updateSekolah({ commit }, payload) {
    try {
      // console.log(payload);
      const response = await api.post(
        `/ss/${payload.schemaname}/update`,
        payload
      );
      // console.log(response.data);
      // commit("SET_TABELSEMESTER", response.data.semester);
      return response.data; // Mengembalikan data sekolah
    } catch (error) {
      commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
      console.error("Gagal membuat semester:", error);
      return null;
    } finally {
      commit("SET_LOADING", false);
    }
  },

  // ==================================
  // SISWA
  // ==================================
  async createBanyakSiswa({ commit }, payload) {
    try {
      const response = await api.post(
        `/ss/${payload.schemaname}/siswa/create-banyak`,
        payload
      );
      console.log(response.data);
      return response.data; // Mengembalikan data sekolah
    } catch (error) {
      commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
      console.error("Gagal mendaftarkan siswa baru:", error);
      return null;
    } finally {
      commit("SET_LOADING", false);
    }
  },
  async fetchSiswa({ commit }, payload) {
    try {
      const response = await api.get(`/ss/${payload.schemaname}/siswa`, {
        params: {
          page: payload.page,
          perpage: payload.perpage,
          // schemaname: schemaname,
        },
      });
      // console.log(response);

      return response.data; // Mengembalikan data sekolah
    } catch (error) {
      // commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
      console.error("Gagal mengambil data siswa:", error);
      return null;
    } finally {
      // commit("SET_LOADING", false);
    }
  },
  async fetchSiswaAktif({ commit }, payload) {
    try {
      const response = await api.get(
        `/ss/${payload.schemaname}/anggota-kelas`,
        {
          params: {
            semester_id: payload.semesterId,
            schemaname: payload.schemaname,
          },
        }
      );
      // console.log(response.data.anggotaKelas);
      commit("SET_TABELSISWAAKTIF", response.data.anggotaKelas);
      return response.data.anggotaKelas;
    } catch (error) {
      commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
      console.error("Gagal membuat semester:", error);
      return null;
    } finally {
      commit("SET_LOADING", false);
    }
  },
  async fetchSiswaCount({ commit }, payload) {
    try {
      const response = await api.get(`/ss/dashboard/countsiswa`, {
        params: {
          schemaname: payload.schemaname,
          semester_id: payload.semesterId,
        },
      });
      // commit("SET_TABELSISWA", response.data.anggotaKelas);
      return response.data; // Mengembalikan data sekolah
    } catch (error) {
      // commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
      console.error("Gagal membuat semester:", error);
      return null;
    } finally {
      // commit("SET_LOADING", false);
    }
  },

  async searchSiswa({ commit }, payload) {
    try {
      const response = await api.get(`/ss/${payload.schemaname}/siswa/search`, {
        params: {
          nm_siswa: payload.nm_siswa,
        },
      });
      // console.log("results",response.data);
      return response.data;
    } catch (error) {
      commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
      console.error("Gagal mengambil data siswa:", error);
      return null;
    } finally {
      commit("SET_LOADING", false);
    }
  },

  // =======================================
  // ANGGOTA KELAS
  // =======================================
  async searchAnggotaKelas({ commit }, payload) {
    try {
      // console.log(payload);
      const response = await api.get(
        `/ss/${payload.schemaname}/anggota-kelas/search`,
        {
          params: {
            semester_id: payload.semester_id,
            peserta_didik_id: payload.peserta_didik_id,
          },
        }
      );
      return response.data.anggotaKelas; // Mengembalikan data sekolah
    } catch (error) {
      if (error.code === "ECONNABORTED") {
        console.error(
          "Permintaan terlalu lama, server lambat atau tidak merespons."
        );
      } else {
        console.error("Terjadi kesalahan:", error.message);
      }
      return null;
    } finally {
      commit("SET_LOADING", false);
    }
  },

  // =======================================
  // REFERENSI TABEL
  // =======================================
  async fetchTabeltenant({ commit }, sekolahId) {
    try {
      const response = await api.get(
        `/sekolah/sekolah-terdaftar?sekolah_id=${sekolahId}`
      );
      commit("SET_TABELTENANT", response.data);
      // console.log(response.data);
      return response.data; // Mengembalikan data sekolah
    } catch (error) {
      commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
      console.error("Gagal mengambil data tabel tenant:", error);
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
  async fetchBentukPendidikan({ commit }) {
    try {
      const response = await api.get(`/ss/ref/bentuk-pendidikan`);
      return response.data.bentukPendidikan;
    } catch (error) {
      // commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
      console.error("Gagal membuat bentuk pendidikan:", error);
      return null;
    } finally {
      // commit("SET_LOADING", false);
    }
  },
  async fetchJenjangPendidikan({ commit }) {
    try {
      const response = await api.get(`/ss/ref/jenjang`);
      return response.data.jenjang;
    } catch (error) {
      // commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
      console.error("Gagal membuat jenjang pendidikan:", error);
      return null;
    } finally {
      // commit("SET_LOADING", false);
    }
  },
  async fetchTingkatPendidikan({ commit }, payload) {
    try {
      const response = await api.get(`/ss/ref/tingkat-pendidikan`, {
        params: {
          jenjang_pendidikan_id: payload.jenjang_pendidikan_id,
        },
      });
      commit("SET_TABELTINGKATPENDIDIKAN", response.data.tingkatPendidikan);
      return response.data.tingkatPendidikan;
    } catch (error) {
      // commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
      console.error("Gagal membuat tingkat pendidikan:", error);
      return null;
    } finally {
      // commit("SET_LOADING", false);
    }
  },
  async fetchKurikulum({ commit }, payload) {
    try {
      const response = await api.get(`/ss/ref/kurikulum`, {
        params: {
          jurusan_id: payload.jurusan_id,
        },
      });
      // console.log(response.data)
      // commit("SET_TABELKURIKULUM", response.data.kurikulum);

      return response.data.kurikulum;
    } catch (error) {
      // commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
      console.error("Gagal membuat kurikulum:", error);
      return null;
    } finally {
      // commit("SET_LOADING", false);
    }
  },
  async fetchJurusan({ commit }, payload) {
    try {
      const response = await api.get(`/ss/ref/jurusan`, {
        params: {
          jenjang_pendidikan_id: payload.jenjang_pendidikan_id,
          param: payload.param,
        },
      });
      return response.data.jurusan;
    } catch (error) {
      // commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
      console.error("Gagal membuat kurikulum:", error);
      return null;
    } finally {
      // commit("SET_LOADING", false);
    }
  },
  async fetchMapel({ commit }, payload) {
    try {
      const response = await api.get(`/ss/ref/mapel`, {
        params: {
          mapel_id: payload.mapel_id,
        },
      });
      console.log("response.data.mapel");
      commit("SET_TABELMAPEL", response.data.mapel);
      return response.data.mapel;
    } catch (error) {
      // commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
      console.error("Gagal memuat mata pelajaran:", error);
      return null;
    } finally {
      // commit("SET_LOADING", false);
    }
  },
  // =======================================
  // KELAS
  // =======================================

  async createKelas({ commit }, payload) {
    try {
      const response = await api.post(
        `/ss/${payload.schemaname}/tambah-kelas`,
        payload
      );
      return response.data; // Mengembalikan data sekolah
    } catch (error) {
      // commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
      console.error("Gagal membuat kelas:", error);
      return null;
    }
  },
  async editKelas({ commit }, payload) {
    try {
      console.log("Payload di edit kelas:", payload);
      const response = await api.put(
        `/ss/${payload.schemaname}/kelas`,
        payload
      );
      // console.log(response.data);
      // commit("SET_TABELSEMESTER", response.data.semester);
      return response.data; // Mengembalikan data sekolah
    } catch (error) {
      // commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
      console.error("Gagal update kelas:", error);
      return null;
    }
  },
  async deleteKelas({ commit }, payload) {
    try {
      console.log("Payload:", payload);
      const response = await api.delete(
        `/ss/${payload.schemaname}/kelas`,
        payload
      );
      console.log(response.data);
      // commit("SET_TABELSEMESTER", response.data.semester);
      return response.data; // Mengembalikan data sekolah
    } catch (error) {
      // commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
      console.error("Gagal menghapus kelas:", error);
      return null;
    }
  },

  // =========================
  // DASHBOARD
  // =========================

  // =======================
  // ANGGOTA KELAS
  // =======================
  async fetchAnggotaKelas({ commit }, payload) {
    try {
      const response = await api.get(
        `/ss/${payload.schemaname}/anggota-kelas`,
        {
          params: {
            semester_id: payload.semester_id,
            rombongan_belajar_id: payload.rombongan_belajar_id,
          },
        }
      );
      return response.data.anggotaKelas;
    } catch (error) {
      console.error("Gagal mendapatkan anggota kelas:", error);
      return null;
    } finally {
    }
  },
  async deleteAnggotaKelas({ commit }, payload) {
    try {
      const response = await api.delete(
        `/ss/${payload.schemaname}/anggota-kelas/delete`,
        {
          params: {
            anggota_rombel_id: payload.anggota_rombel_id,
          },
        }
      );
      // console.log(response.data);
      // commit("SET_TABELSEMESTER", response.data.semester);
      return response.data; // Mengembalikan data sekolah
    } catch (error) {
      // commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
      console.error("Gagal menghapus anggota kelas:", error);
      return null;
    }
  },

  // ==================================
  // PEMBELAJARAN SERVICE
  // ==================================
  async createPembelajaran({ commit }, payload) {
    try {
      const response = await api.post(`ss/pembelajaran/create`, payload);
      console.log(response.data);
      return response.data; // Mengembalikan data sekolah
    } catch (error) {
      commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
      console.error("Gagal mendaftarkan siswa baru:", error);
      return null;
    } finally {
      commit("SET_LOADING", false);
    }
  },

  // ==================================
  // IJAZAH SERVICE
  // ==================================
  async createProsesIjazah({ commit }, payload) {
    try {
      const response = await api.post(`ss/ijazah/create`, payload);
      return response.data; // Mengembalikan data sekolah
    } catch (error) {
      // console.error("Gagal mendaftarkan siswa baru:", error);
      throw error;
    } finally {
      commit("SET_LOADING", false);
    }
  },
  async fetchProsesIjazah({ commit }, payload) {
    try {
      const response = await api.get(`ss/proses-ijazah`, {
        params: {
          schemaname: payload.schemaname,
          semester_id: payload.tahun_ajaran_id,
          ijazah_id: payload.ijazah_id,
        },
      });
      return response.data.anggotaKelas; // Mengembalikan data sekolah
    } catch (error) {
      throw error;
    } finally {
      commit("SET_LOADING", false);
    }
  },
  async fetchNilaiSiswa({ commit }, payload) {
    try {
      const response = await api.get(`ss/${payload.schemaname}/nilai-akhir`, {
        params: {
          semester_id: payload.semesterId,
          peserta_didik_id: payload.semester_id,
        },
      });
      commit("SET_TABELNILAIAKHIR", response.data?.nilaiSiswa);
      return response.data.nilaiSiswa; // Mengembalikan nilai siswa
    } catch (error) {
      // commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
      console.error("Gagal nilai siswa:", error);
      return null;
    } finally {
      // commit("SET_LOADING", false);
    }
  },
  async createKenaikan({ commit }, payload) {
    try {
      const response = await api.post(`ss/kenaikan/create`, payload);
      // commit("SET_TABELNILAIAKHIR", response.data?.nilaiSiswa);

      console.log(response.data);
    } catch (error) {
      // commit("SET_ERROR", error.response?.data || "Terjadi kesalahan");
      console.error("Gagal nilai siswa:", error);
      return null;
    } finally {
      // commit("SET_LOADING", false);
    }
  },
};

const getters = {
  getKelas: (state) => (semesterId) => {
    return state.tabelKelas.filter((kelas) => kelas.semesterId === semesterId);
  },
  isLoading: (state) => state.loading,
  getError: (state) => state.error,
  getTabeltenant: (state) => state.tabelTenant,
  getSemester: (state) => state.tabelSemester,
  getSelectedSemester: (state) => state.selectedSemester,
  getSekolah: (state) => state.tabelSekolah,
  getKurikulum: (state) => state.tabelKurikulum,
  getGuru: (state) => state.tabelGuru,
  getPTKTerdaftar: (state) => state.tabelPTKTerdaftar,
  getTingkatPendidikan: (state) => state.tabelTingkatPendidikan,
  getTabelAnggotaKelas: (state) => state.tabelAnggotaKelas,
  getSiswaAktif: (state) => state.tabelSiswaAktif,
  getMapel: (state) => state.tabelMapel,
  getNilaiSiswa: (state) => state.tabelNilaiakhir,
  getTahunAjaran: (state) => state.tabelTahunAjaran,
  getSelectedTahunAjaran: (state) => state.selectedTahunAjaran,
};

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters,
};
