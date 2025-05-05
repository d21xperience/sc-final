import { fileURLToPath, URL } from "node:url";

import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import tailwindcss from "@tailwindcss/vite";
import path from "path";
// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue(), tailwindcss()],
  server: {
    proxy: {
      // "/dapo": {
      //   target: "https://dapo.kemdikbud.go.id",
      //   changeOrigin: true,
      //   rewrite: (path) => path.replace(/^\/dapo/, ""),
      //   // secure: false,
      // },
      // Pastikan path `/dapo/api` sesuai dengan yang kamu gunakan di frontend
      // "/dapo/api": {
      //   target: "https://dapo.dikdasmen.go.id",
      //   changeOrigin: true,
      //   rewrite: (path) => path.replace(/^\/dapo\/api/, "/api"),
      // },
      // "/api": {
      //   target: "https://dapo.kemdikbud.go.id",
      //   changeOrigin: true,
      //   rewrite: (path) => {
      //     console.log(`Proxying request: ${path}`);
      //     console.log(`Proxying request: ${path.replace(/^\/api/, "")}`);
      //     return path.replace(/^\/api/, "");
      //   },
      //   // tambahan
      //   secure: false,
      //   ws: true,
      // },
      // "/dapo": {
      //   target: "http://localhost:5774",
      //   changeOrigin: true,
      //   rewrite: (path) => {
      //     console.log(`Proxying request: ${path}`);
      //     console.log(`Proxying request: ${path.replace(/^\/api/, "")}`);
      //     return path.replace(/^\/dapo/, "");
      //   },
      // },
    },
  },
  resolve: {
    alias: [{ find: "@", replacement: path.resolve(__dirname, "./src") }],
  },
});
