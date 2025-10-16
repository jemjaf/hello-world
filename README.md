# 🐳 Hola Mundo / Hello World - Imágenes Docker

[![Docker Pulls](https://img.shields.io/docker/pulls/jemjaf/hola-mundo)](https://hub.docker.com/r/jemjaf/hola-mundo)
[![Image Size](https://img.shields.io/docker/image-size/jemjaf/hola-mundo/latest)](https://hub.docker.com/r/jemjaf/hola-mundo)

Imágenes Docker personalizadas inspiradas en el hello-world de Rancher, con soporte multiidioma, interfaz moderna e integración con redes sociales.

## ✨ Características

- 🌍 **Soporte Multiidioma**: Español (`hola-mundo`) e Inglés (`hello-world`)
- 🎨 **Interfaz Moderna**: Diseño limpio y responsivo con estilo inspirado en Rancher
- 📱 **Integración Social**: Enlaces directos a GitHub, TikTok, Instagram, LinkedIn, YouTube
- 🚀 **Ultra-liviana**: Solo 7.12MB de tamaño
- 🔧 **Endpoints de Salud**: Health checks y monitoreo integrados
- 📊 **Detalles de Solicitud**: Visualización interactiva de información de requests
- 🎯 **Marca Personal**: Personalizable con tus propios enlaces e información

## 🚀 Inicio Rápido

### Versión en Español
```bash
# Ejecutar la versión en español
docker run -p 80:80 jemjaf/hola-mundo

# O con versión específica
docker run -p 80:80 jemjaf/hola-mundo:v0.0.1
```

### Versión en Inglés
```bash
# Ejecutar la versión en inglés
docker run -p 80:80 jemjaf/hello-world

# O con versión específica
docker run -p 80:80 jemjaf/hello-world:v0.0.1
```

## 📋 Imágenes Disponibles

| Imagen | Idioma | Descripción | Tamaño |
|--------|--------|-------------|--------|
| `jemjaf/hola-mundo` | 🇪🇸 Español | "Hola Mundo!" con interfaz en español | 7.12MB |
| `jemjaf/hello-world` | 🇺🇸 Inglés | "Hello world!" con interfaz en inglés | 7.12MB |

## 🌐 Endpoints

Una vez ejecutándose, los siguientes endpoints están disponibles:

- **Página Principal**: `http://localhost/` - Interfaz web interactiva
- **Health Check**: `http://localhost/health` - Estado de salud del servicio
- **Información del Servicio**: `http://localhost/info` - Información y metadatos del servicio
- **Logs**: `http://localhost/logs` - Logs de la aplicación en formato Base64
- **Assets**: `http://localhost/assets/icon-*.svg` - Iconos de redes sociales

## 🛠️ Construcción desde Código Fuente

### Prerrequisitos
- Docker
- Go 1.21+ (para desarrollo local)

### Comandos de Construcción

```bash
# Clonar el repositorio
git clone https://github.com/jemjaf/hello-world.git
cd hello-world

# Construir versión en español (por defecto)
docker build -t hola-mundo-personalizado .

# Construir versión en inglés
docker build --build-arg LANG=en -t hello-world-personalizado .

# O usar el script de construcción
chmod +x build.sh
./build.sh
```

### Argumentos de Construcción Personalizados

```bash
# Construir con idioma personalizado
docker build --build-arg LANG=en -t mi-hello-world .
docker build --build-arg LANG=es -t mi-hola-mundo .
```

## 🎨 Personalización

### Enlaces de Redes Sociales
Edita `main.go` para personalizar los enlaces de redes sociales:

```go
// Actualiza estas URLs con tus propios perfiles
<a href="https://github.com/jemjaf" class="social-icon github">
<a href="https://tiktok.com/@jemjaf" class="social-icon rancher">
<a href="https://instagram.com/jemjaf" class="social-icon twitter">
<a href="https://linkedin.com/in/jemjaf" class="social-icon facebook">
<a href="https://youtube.com/@jemjaf" class="social-icon linkedin">
```

### Marca Personal
Actualiza el enlace del header en `main.go`:

```go
<a href="https://jemjaf.com" class="brand">Jean Reyes (@jemjaf)</a>
```

### Agregar Nuevos Idiomas
1. Agrega un nuevo caso de idioma en la función `getTexts()`
2. Actualiza la documentación del ARG en el Dockerfile
3. Construye con el nuevo argumento de idioma

## 📊 Detalles Técnicos

### Arquitectura
- **Imagen Base**: `scratch` (ultra-minimal)
- **Proceso de Construcción**: Build multi-etapa con compilación Go
- **Runtime**: Binario estático con assets embebidos
- **Puerto**: 80 (configurable vía variable de entorno PORT)

### Rendimiento
- **Tiempo de Inicio**: < 1 segundo
- **Uso de Memoria**: ~2MB RAM
- **Uso de CPU**: Mínimo
- **Red**: Listo para HTTP/HTTPS

### Seguridad
- **Sin root**: Ejecuta como usuario no privilegiado
- **Superficie de Ataque Mínima**: Solo componentes esenciales incluidos
- **Sin Shell**: No incluye shell ni gestor de paquetes
- **Binario Estático**: Sin dependencias dinámicas

## 🔧 Variables de Entorno

| Variable | Por Defecto | Descripción |
|----------|-------------|-------------|
| `PORT` | `80` | Puerto en el que escuchar |

## 📝 Ejemplos

### Docker Compose
```yaml
services:
  hola-mundo-es:
    image: jemjaf/hola-mundo
    ports:
      - "8080:80"
    environment: # Opcional para cambiar el puerto
      - PORT=80

  hello-world-en:
    image: jemjaf/hello-world
    ports:
      - "8081:80"
```

### Despliegue en Kubernetes
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: hola-mundo
spec:
  replicas: 3
  selector:
    matchLabels:
      app: hola-mundo
  template:
    metadata:
      labels:
        app: hola-mundo
    spec:
      containers:
      - name: hola-mundo
        image: jemjaf/hola-mundo
        ports:
        - containerPort: 80
```

### Integración de Health Check
```bash
# Health check de Docker
docker run -d \
  --name hola-mundo \
  --health-cmd="curl -f http://localhost/health || exit 1" \
  --health-interval=30s \
  jemjaf/hola-mundo
```

## 👨‍💻 Autor

**Jean Reyes (@jemjaf)**
- Sitio Web: [https://jemjaf.com](https://jemjaf.com)
- GitHub: [@jemjaf](https://github.com/jemjaf)
- LinkedIn: [jemjaf](https://linkedin.com/in/jemjaf)

---

⭐ **¡Dale estrella a este repositorio si te resulta útil!**
