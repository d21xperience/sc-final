// src/composables/useOptions.js
import { computed } from "vue";

export function useRombelOptions(dataRombel) {
  const tingkatPendidikanOptions = computed(() => {
    const map = new Map();
    for (const item of dataRombel.value) {
      const tp = item.tingkatPendidikan;
      if (tp && !map.has(tp.tingkatPendidikanId)) {
        map.set(tp.tingkatPendidikanId, {
          label: tp.nama,
          value: tp.tingkatPendidikanId,
        });
      }
    }
    return Array.from(map.values());
  });

  const jurusanOptions = computed(() => {
    console.log(dataRombel)
    const map = new Map();
    for (const item of dataRombel.value) {
      const jur = item.jurusan;
      if (jur && !map.has(jur.jurusanId)) {
        map.set(jur.jurusanId, {
          label: jur.namaJurusan,
          value: jur.jurusanId,
        });
      }
    }
    return Array.from(map.values());
  });

  const formatterDateID = (tanggalRaw) => {
    // const tanggalRaw = "2006-11-28 00:00:00 +0000 UTC";
    const tanggal = new Date(tanggalRaw);

    const formatter = new Intl.DateTimeFormat("id-ID", {
      day: "2-digit",
      month: "long",
      year: "numeric",
    });
    return formatter.format(tanggal);
  };
  return {
    tingkatPendidikanOptions,
    jurusanOptions,
    formatterDateID,
  };
}
