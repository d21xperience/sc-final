// // composables/useExcelUpload.js
// import * as XLSX from "xlsx";
// import { ref } from "vue";
// import store from "@/store";

// const schemaMap = {
//   guru: [
//     "ptk_id",
//     "nama",
//     "nip",
//     "jenis_ptk_id",
//     "jenis_kelamin",
//     "tempat_lahir",
//     "tanggal_lahir",
//     "nuptk",
//     "alamat_jalan",
//     "status_keaktifan_id",
//   ],
//   siswa: [
//     "peserta_didik_id",
//     "nis",
//     "nisn",
//     "nm_siswa",
//     "tempat_lahir",
//     "tanggal_lahir",
//     "jenis_kelamin",
//     "agama",
//     "alamat_siswa",
//     "telepon_siswa",
//     "diterima_tanggal",
//     "nm_ayah",
//     "nm_ibu",
//     "pekerjaan_ayah",
//     "pekerjaan_ibu",
//     "nm_wali",
//     "pekerjaan_wali",
//   ],
//   nilaiIjazah: [
//     /* ... */
//   ],
// };

// const mapExcelToObjects = (jsonData, schema) => {
//   const rows = jsonData.slice(1); // baris ke-1 adalah header
//   return rows.map((row) => {
//     const obj = {};
//     schema.forEach((key, index) => {
//       obj[key] = row[index] ?? ""; // fallback ke "" kalau null/undefined
//     });
//     return obj;
//   });
// };

// export function useExcelUpload({ maxSizeMB = 2, maxRows = 2000 } = {}) {
//   const excelData = ref([]);
//   const loading = ref(false);
//   const error = ref(null);

//   const handleFileUpload = (event) => {
//     error.value = null;
//     excelData.value = [];

//     const file = event.target.files[0];
//     if (!file) return;

//     // Ukuran maksimal (dalam MB)
//     const fileSizeMB = file.size / 1024 / 1024;
//     if (fileSizeMB > maxSizeMB) {
//       error.value = `Ukuran file melebihi batas ${maxSizeMB}MB. Harap upload langsung ke backend.`;
//       return;
//     }

//     loading.value = true;
//     const reader = new FileReader();

//     reader.onload = (e) => {
//       try {
//         const data = new Uint8Array(e.target.result);
//         const workbook = XLSX.read(data, { type: "array" });
//         const sheetName = workbook.SheetNames[0];
//         const worksheet = workbook.Sheets[sheetName];

//         const jsonData = XLSX.utils.sheet_to_json(worksheet, { header: 1 });

//         // Batas maksimum baris
//         if (jsonData.length > maxRows) {
//           error.value = `Jumlah baris lebih dari ${maxRows}. Silakan upload ke backend.`;
//           loading.value = false;
//           return;
//         }

//         excelData.value = jsonData;
//       } catch (err) {
//         error.value = "Gagal memproses file. Pastikan format valid.";
//       } finally {
//         loading.value = false;
//       }
//     };

//     reader.onerror = () => {
//       error.value = "Gagal membaca file.";
//       loading.value = false;
//     };

//     reader.readAsArrayBuffer(file);
//   };

//   const uploadToBackend = async (uri, uploadType) => {
//     if (!schemaMap[uploadType]) {
//       error.value = `Type upload harus diisi.`;
//       return;
//     }

//     if (excelData.value.length === 0) {
//       error.value = "Tidak ada data untuk diupload.";
//       return;
//     }

//     try {
//       loading.value = true;
//       error.value = null;
//       const mappedData = mapExcelToObjects(
//         excelData.value,
//         schemaMap[uploadType]
//       );
//       // Kirim data ke backend
//       const response = await store.dispatch(uri, mappedData);
//       // const response = await fetch("http://localhost:3000/api/upload", {
//       //   method: "POST",
//       //   headers: {
//       //     "Content-Type": "application/json",
//       //   },
//       //   body: JSON.stringify({ data: excelData.value }),
//       // });

//       const result = await response.json();
//       if (!response.ok) {
//         throw new Error(result.message || "Upload gagal.");
//       }

//       console.log("Upload berhasil:", result);
//     } catch (err) {
//       error.value = err.message || "Terjadi kesalahan saat upload.";
//     } finally {
//       loading.value = false;
//     }
//   };

//   return {
//     excelData,
//     loading,
//     error,
//     handleFileUpload,
//     uploadToBackend,
//   };
// }
