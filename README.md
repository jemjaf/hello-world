# 🐳 Hola Mundo / Hello World

[![Docker Pulls](https://img.shields.io/docker/pulls/jemjaf/hola-mundo)](https://hub.docker.com/r/jemjaf/hola-mundo)
[![Image Size](https://img.shields.io/docker/image-size/jemjaf/hola-mundo/latest)](https://hub.docker.com/r/jemjaf/hola-mundo)

Imagen Docker súper liviana (~7MB) y multiarquitectura (`amd64` y `arm64`) perfecta para pruebas, tutoriales y health checks. 
Disponible en Español (`hola-mundo`) e Inglés (`hello-world`).

## 🚀 Uso Rápido (Solo usarla)

```bash
# Versión en Español
docker run -p 80:80 jemjaf/hola-mundo

# Versión en Inglés
docker run -p 80:80 jemjaf/hello-world
```

¡Y listo! Entrá a [http://localhost](http://localhost) en tu navegador.

## 🛠️ Cómo construirla (Para modificarla)

Si querés clonar el repo y hacer tu propia versión:

```bash
git clone https://github.com/jemjaf/hello-world.git
cd hello-world

# Compilar para tu arquitectura local (español por defecto)
docker build -t mi-hola-mundo .

# Compilar en inglés
docker build --build-arg LANG=en -t mi-hello-world .

# 🚀 PRO TIP: Compilar para MÚLTIPLES arquitecturas (amd64 y arm64)
docker buildx build --platform linux/amd64,linux/arm64 -t tu-usuario/hola-mundo:latest --push .
```

## 🌐 Endpoints Disponibles

- `http://localhost/` : Interfaz visual
- `http://localhost/health` : Health check JSON (ideal para Kubernetes/Docker Swarm)
- `http://localhost/info` : Metadatos del pod/contenedor
- `http://localhost/logs` : Ejemplo de logs en Base64

## 👨‍💻 Comunidad y Redes

- 🎓 **Comunidad de Aprendizaje:** [skool.com/jemjaf](https://skool.com/jemjaf) (¡Sumate!)
- 💻 **GitHub:** [@jemjaf](https://github.com/jemjaf)
- 📱 **TikTok:** [@jemjaf](https://tiktok.com/@jemjaf)
- 📸 **Instagram:** [@jemjaf](https://instagram.com/jemjaf)
- 💼 **LinkedIn:** [jemjaf](https://linkedin.com/in/jemjaf)
- 🎥 **YouTube:** [@jemjaf](https://youtube.com/@jemjaf)
- 🌐 **Web:** [jemjaf.com](https://jemjaf.com)

---

⭐ **¡Dale estrella a este repositorio si te resulta útil!**
