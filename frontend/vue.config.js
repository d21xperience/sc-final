module.exports = {
  devServer: {
    proxy: {
      "/auth": {
        target: "http://nginx",
        changeOrigin: true,
      },
      "/sc": {
        target: "http://nginx",
        changeOrigin: true,
      },
      "/sekolah": {
        target: "http://nginx",
        changeOrigin: true,
      },
    },
  },
};
