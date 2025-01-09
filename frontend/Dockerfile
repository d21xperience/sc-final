# Gunakan image Node.js untuk build
FROM node:18 AS builder

# Set working directory
WORKDIR /app

# Copy package.json dan install dependencies
COPY package*.json ./
RUN npm install

# Copy seluruh source code
COPY . .

# Build aplikasi
RUN npm run build

# Image runtime
FROM nginx:alpine

# Copy build output ke NGINX
COPY --from=builder /app/dist /usr/share/nginx/html

# Expose port default NGINX
EXPOSE 80

# Jalankan NGINX
CMD ["nginx", "-g", "daemon off;"]
