FROM nginx:latest

# Copy konfigurasi NGINX
COPY nginx.conf /etc/nginx/nginx.conf

# Copy sertifikat SSL/TLS
COPY certs /etc/nginx/certs

# Expose ports for HTTP and HTTPS
EXPOSE 80
EXPOSE 443
