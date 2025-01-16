# Gunakan image Node.js terbaru sebagai base image
FROM node:20-alpine AS build-stage

# Set direktori kerja di dalam container
WORKDIR /app

# Salin file package.json dan package-lock.json (jika ada)
COPY package.json ./

# Install dependensi
RUN npm install

# Salin seluruh proyek ke dalam container
COPY . .

# Build proyek menggunakan Vite
RUN npm run build

# Gunakan image Nginx sebagai production server
FROM nginx:1.25.3-alpine AS production-stage

# Salin file build hasil dari tahap sebelumnya ke direktori Nginx
COPY --from=build-stage /app/dist /usr/share/nginx/html

# Salin konfigurasi default Nginx
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Expose port 80
EXPOSE 80

# Jalankan Nginx
CMD ["nginx", "-g", "daemon off;"]
