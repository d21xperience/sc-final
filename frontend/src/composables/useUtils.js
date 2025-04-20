// src/composables/useOptions.js
import { computed } from "vue";

export function useUtils(dataRombel) {
  //   const tingkatPendidikanOptions = computed(() => {
  //     const map = new Map();
  //     for (const item of dataRombel.value) {
  //       const tp = item.tingkatPendidikan;
  //       if (tp && !map.has(tp.tingkatPendidikanId)) {
  //         map.set(tp.tingkatPendidikanId, {
  //           label: tp.nama,
  //           value: tp.tingkatPendidikanId,
  //         });
  //       }
  //     }
  //     return Array.from(map.values());
  //   });

  //   const jurusanOptions = computed(() => {
  //     const map = new Map();
  //     for (const item of dataRombel.value) {
  //       const jur = item.jurusan;
  //       if (jur && !map.has(jur.jurusanId)) {
  //         map.set(jur.jurusanId, {
  //           label: jur.namaJurusan,
  //           value: jur.jurusanId,
  //         });
  //       }
  //     }
  //     return Array.from(map.values());
  //   });

  const formatterDateID = (tanggalRaw) => {
    if (!tanggalRaw) return "-";

    // Normalize format: Ganti spasi ke 'T' jika diperlukan
    let normalizedDate = tanggalRaw.replace(" ", "T");

    // Buat objek Date
    const tanggal = new Date(normalizedDate);

    // Cek validitas
    if (isNaN(tanggal)) return "-";

    const formatter = new Intl.DateTimeFormat("id-ID", {
      day: "2-digit",
      month: "long",
      year: "numeric",
    });

    return formatter.format(tanggal);
  };

  return {
    // tingkatPendidikanOptions,
    // jurusanOptions,
    formatterDateID,
  };
}
