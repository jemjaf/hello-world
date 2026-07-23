# Multi-stage build ultra-optimizado para tamaño mínimo
FROM --platform=$BUILDPLATFORM golang:1.21-alpine AS builder

# Argumento para el idioma (por defecto español)
ARG LANG=es

# Argumento proporcionado por Docker Buildx para la arquitectura destino
ARG TARGETOS
ARG TARGETARCH

# Instalar dependencias mínimas para compilar
RUN apk add --no-cache git ca-certificates

# Crear directorio de trabajo
WORKDIR /app

# Copiar archivos fuente y assets
COPY main.go .
COPY assets/ ./assets/

# Compilar la aplicación Go estáticamente para imagen scratch usando TARGETOS y TARGETARCH
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -a -installsuffix cgo -ldflags '-w -s -X main.Language='$LANG -o hello-world main.go

# Imagen final ultra-liviana usando scratch (imagen vacía)
FROM scratch

# Copiar certificados SSL desde Alpine
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copiar el binario compilado y assets
COPY --from=builder /app/hello-world /hello-world
COPY --from=builder /app/assets /assets

# Exponer puerto
EXPOSE 80

# Definir el punto de entrada
ENTRYPOINT ["/hello-world"]
