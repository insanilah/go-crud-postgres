# Gunakan image Go resmi berbasis Alpine untuk build
FROM golang:1.23-alpine AS build

# Setel direktori kerja
WORKDIR /app

# Copy semua file Go ke dalam container
COPY . .

# Install dependencies untuk build
RUN go mod tidy

# Build aplikasi Go
RUN go build -o go-crud-postgres .

# Gunakan image ringan untuk menjalankan aplikasi
FROM alpine:latest

# ARG untuk environment
ARG ENV=production
ENV APP_ENV=${ENV}

# Install dependencies yang dibutuhkan menggunakan apk (untuk Alpine)
RUN apk update && apk add --no-cache \
    ca-certificates \
    libpq \
    curl \
    && rm -rf /var/cache/apk/*

# Tambahkan metadata
LABEL maintainer="satria.gitu4@gmail.com"
LABEL version="1.0.0"
LABEL description="Aplikasi CRUD Go dengan PostgreSQL"

# Copy aplikasi yang sudah dibuild dari tahap build
COPY --from=build /app/go-crud-postgres /usr/local/bin/

# Expose port untuk aplikasi
EXPOSE 8080

# Healthcheck untuk memastikan container aktif
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
  CMD curl --fail http://localhost:8080/health || exit 1

# Jalankan aplikasi
CMD ["go-crud-postgres"]
