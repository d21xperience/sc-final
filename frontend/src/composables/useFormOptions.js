// composables/useFormOptions.js
import { ref } from "vue";

export function useFormOptions() {
  const selectedJenisKelamin = ref();
  const jenisKelaminOptions = ref([
    { label: "Laki-Laki", value: "L" },
    { label: "Perempuan", value: "P" },
  ]);

  const selectedAgama = ref();
  const agamaOptions = ref([
    { label: "Islam", value: "Islam" },
    { label: "Kristen", value: "Kristen" },
    { label: "Katolik", value: "Katolik" },
    { label: "Hindu", value: "Hindu" },
    { label: "Buddha", value: "Buddha" },
    { label: "Konghucu", value: "Konghucu" },
  ]);

  return {
    selectedJenisKelamin,
    jenisKelaminOptions,
    selectedAgama,
    agamaOptions,
  };
}
