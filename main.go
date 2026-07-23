package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Variable para el idioma (se establece en tiempo de compilación)
var Language string

// Estructura para los textos multiidioma
type Texts struct {
	Title       string
	Hostname    string
	ShowDetails string
	HideDetails string
	RequestInfo string
	Service     string
	Version     string
	Author      string
	Social      string
	Started     string
	Port        string
	RequestFrom string
	UserAgent   string
	Processed   string
}

// Función para obtener los textos según el idioma
func getTexts() Texts {
	if Language == "en" {
		return Texts{
			Title:       "Hello world!",
			Hostname:    "My hostname is",
			ShowDetails: "Show request details",
			HideDetails: "Hide request details",
			RequestInfo: "Request info",
			Service:     "hello-world",
			Version:     "1.0.0",
			Author:      "Author: Jean Reyes (@jemjaf)",
			Social:      "Social: GitHub, TikTok, Instagram, LinkedIn, YouTube, Skool",
			Started:     "Hello World Service started",
			Port:        "Port:",
			RequestFrom: "Request from:",
			UserAgent:   "User-Agent:",
			Processed:   "Request processed successfully",
		}
	}
	// Por defecto español
	return Texts{
		Title:       "Hola Mundo!",
		Hostname:    "Mi hostname es",
		ShowDetails: "Mostrar detalles de la solicitud",
		HideDetails: "Ocultar detalles de la solicitud",
		RequestInfo: "Información de la solicitud",
		Service:     "hola-mundo",
		Version:     "1.0.0",
		Author:      "Autor: Jean Reyes (@jemjaf)",
		Social:      "Social: GitHub, TikTok, Instagram, LinkedIn, YouTube, Skool",
		Started:     "Servicio Hola Mundo iniciado",
		Port:        "Puerto:",
		RequestFrom: "Solicitud desde:",
		UserAgent:   "User-Agent:",
		Processed:   "Solicitud procesada exitosamente",
	}
}

func getHeaderValue(r *http.Request, header string) string {
	value := r.Header.Get(header)
	if value == "" {
		return ""
	}
	return value
}

func main() {
	// Obtener textos según el idioma
	texts := getTexts()
	
	// Obtener puerto del entorno o usar 80 por defecto
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	// Handler para la ruta principal
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		
		// Generar HTML que replica el diseño de Rancher
		response := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <title>Hello world!</title>
    <style>
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
            margin: 0;
            padding: 0;
            background-color: #f8f9fa;
            color: #333;
        }
        .header {
            background-color: #007cbb;
            padding: 30px 40px;
            text-align: center;
        }
        .brand {
            font-size: 28px;
            font-weight: 600;
            color: white;
            text-decoration: none;
            transition: all 0.3s ease;
            display: inline-block;
            padding: 5px 10px;
            border-radius: 4px;
        }
        .brand:hover {
            color: #e8f4fd;
            background-color: rgba(255,255,255,0.1);
            transform: scale(1.05);
        }
        .container {
            max-width: 600px;
            margin: 40px auto;
            padding: 0 20px;
        }
        .main-content {
            background-color: white;
            padding: 50px;
            border-radius: 12px;
            box-shadow: 0 4px 20px rgba(0,0,0,0.1);
            text-align: center;
        }
        h1 {
            font-size: 32px;
            font-weight: 700;
            margin: 0 0 20px 0;
            color: #333;
        }
        .hostname {
            font-size: 18px;
            color: #666;
            margin-bottom: 30px;
        }
        .social-icons {
            display: flex;
            gap: 15px;
            margin-bottom: 30px;
            justify-content: center;
        }
        .social-icon {
            width: 32px;
            height: 32px;
            border-radius: 50%%;
            display: flex;
            align-items: center;
            justify-content: center;
            text-decoration: none;
            font-size: 16px;
            transition: all 0.3s ease;
            cursor: pointer;
            box-shadow: 0 2px 4px rgba(0,0,0,0.1);
        }
        .social-icon:hover {
            transform: scale(1.15) translateY(-2px);
            box-shadow: 0 4px 12px rgba(0,0,0,0.2);
            background-color: #f8f9fa;
        }
        .social-icon img {
            /* Sin filtros para mantener colores originales */
        }
        .github { background-color: white; color: #333; }
        .rancher { background-color: white; color: #333; }
        .twitter { background-color: white; color: #333; }
        .facebook { background-color: white; color: #333; }
        .linkedin { background-color: white; color: #333; }
        .show-details-btn {
            background-color: #007cbb;
            color: white;
            border: none;
            padding: 12px 24px;
            border-radius: 6px;
            font-size: 16px;
            cursor: pointer;
            margin-bottom: 30px;
            transition: all 0.3s ease;
            box-shadow: 0 2px 4px rgba(0,124,187,0.3);
        }
        .show-details-btn:hover {
            background-color: #005a8b;
            transform: translateY(-2px);
            box-shadow: 0 4px 12px rgba(0,124,187,0.4);
        }
        .show-details-btn:active {
            transform: translateY(0);
            box-shadow: 0 2px 4px rgba(0,124,187,0.3);
        }
        .request-info {
            background-color: #f8f9fa;
            padding: 20px;
            border-radius: 6px;
            border-left: 4px solid #007cbb;
        }
        .request-info h3 {
            margin: 0 0 15px 0;
            font-size: 18px;
            color: #333;
        }
        .header-list {
            font-family: 'Courier New', monospace;
            font-size: 14px;
            line-height: 1.6;
        }
        .header-item {
            margin-bottom: 8px;
        }
        .header-name {
            font-weight: bold;
            color: #333;
        }
        .header-value {
            color: #666;
            word-break: break-all;
            overflow-wrap: break-word;
        }
    </style>
</head>
<body>
    <div class="header">
        <a href="https://jemjaf.com" class="brand">Jean Reyes (@jemjaf)</a>
    </div>
    
    <div class="container">
        <div class="main-content">
            <h1>%s</h1>
            <div class="hostname">%s %s</div>
            
            <div class="social-icons">
                <a href="https://tiktok.com/@jemjaf" class="social-icon rancher" title="TikTok">
                    <img src="/assets/icon-tiktok.svg" alt="TikTok" width="24" height="24">
                </a>
                <a href="https://instagram.com/jemjaf" class="social-icon twitter" title="Instagram">
                    <img src="/assets/icon-instagram.svg" alt="Instagram" width="24" height="24">
                </a>
                <a href="https://youtube.com/@jemjaf" class="social-icon linkedin" title="YouTube">
                    <img src="/assets/icon-youtube.svg" alt="YouTube" width="24" height="24">
                </a>
                <a href="https://linkedin.com/in/jemjaf" class="social-icon facebook" title="LinkedIn">
                    <img src="/assets/icon-linkedin.svg" alt="LinkedIn" width="24" height="24">
                </a>
                <a href="https://github.com/jemjaf/hello-world" class="social-icon github" title="GitHub">
                    <img src="/assets/icon-github.svg" alt="GitHub" width="24" height="24">
                </a>
                <a href="https://skool.com/jemjaf" class="social-icon skool" title="Comunidad Skool" style="background-color: transparent;">
                    <img src="/assets/icon-skool.jpg" alt="Skool" width="32" height="32" style="border-radius: 50%;">
                </a>
            </div>
            
            <button class="show-details-btn" onclick="toggleDetails()">%s</button>
            
            <div class="request-info" id="requestDetails" style="display: none;">
                <h3>%s</h3>
                <div class="header-list">
                    <div class="header-item">
                        <span class="header-name">Host:</span> <span class="header-value">[%s]</span>
                    </div>
                    <div class="header-item">
                        <span class="header-name">Pod:</span> <span class="header-value">[%s]</span>
                    </div>
                    <div class="header-item">
                        <span class="header-name">Accept:</span> <span class="header-value">[%s]</span>
                    </div>
                    <div class="header-item">
                        <span class="header-name">Accept-Encoding:</span> <span class="header-value">[%s]</span>
                    </div>
                    <div class="header-item">
                        <span class="header-name">Accept-Language:</span> <span class="header-value">[%s]</span>
                    </div>
                    <div class="header-item">
                        <span class="header-name">Connection:</span> <span class="header-value">[%s]</span>
                    </div>
                    <div class="header-item">
                        <span class="header-name">Cookie:</span> <span class="header-value">[%s]</span>
                    </div>
                    <div class="header-item">
                        <span class="header-name">Sec-Ch-Ua:</span> <span class="header-value">[%s]</span>
                    </div>
                    <div class="header-item">
                        <span class="header-name">Sec-Ch-Ua-Mobile:</span> <span class="header-value">[%s]</span>
                    </div>
                    <div class="header-item">
                        <span class="header-name">Sec-Ch-Ua-Platform:</span> <span class="header-value">[%s]</span>
                    </div>
                    <div class="header-item">
                        <span class="header-name">Sec-Fetch-Dest:</span> <span class="header-value">[%s]</span>
                    </div>
                    <div class="header-item">
                        <span class="header-name">Sec-Fetch-Mode:</span> <span class="header-value">[%s]</span>
                    </div>
                    <div class="header-item">
                        <span class="header-name">Sec-Fetch-Site:</span> <span class="header-value">[%s]</span>
                    </div>
                    <div class="header-item">
                        <span class="header-name">Sec-Fetch-User:</span> <span class="header-value">[%s]</span>
                    </div>
                    <div class="header-item">
                        <span class="header-name">Sec-Gpc:</span> <span class="header-value">[%s]</span>
                    </div>
                    <div class="header-item">
                        <span class="header-name">Upgrade-Insecure-Requests:</span> <span class="header-value">[%s]</span>
                    </div>
                    <div class="header-item">
                        <span class="header-name">User-Agent:</span> <span class="header-value">[%s]</span>
                    </div>
                </div>
            </div>
        </div>
    </div>
    
    <script>
        function toggleDetails() {
            const details = document.getElementById('requestDetails');
            const button = document.querySelector('.show-details-btn');
            if (details.style.display === 'none') {
                details.style.display = 'block';
                button.textContent = '%s';
            } else {
                details.style.display = 'none';
                button.textContent = '%s';
            }
        }
    </script>
</body>
</html>`, 
			texts.Title,
			texts.Hostname,
			hostname,
			texts.ShowDetails,
			texts.RequestInfo,
			r.Host,
			hostname,
			getHeaderValue(r, "Accept"),
			getHeaderValue(r, "Accept-Encoding"),
			getHeaderValue(r, "Accept-Language"),
			getHeaderValue(r, "Connection"),
			getHeaderValue(r, "Cookie"),
			getHeaderValue(r, "Sec-Ch-Ua"),
			getHeaderValue(r, "Sec-Ch-Ua-Mobile"),
			getHeaderValue(r, "Sec-Ch-Ua-Platform"),
			getHeaderValue(r, "Sec-Fetch-Dest"),
			getHeaderValue(r, "Sec-Fetch-Mode"),
			getHeaderValue(r, "Sec-Fetch-Site"),
			getHeaderValue(r, "Sec-Fetch-User"),
			getHeaderValue(r, "Sec-Gpc"),
			getHeaderValue(r, "Upgrade-Insecure-Requests"),
			getHeaderValue(r, "User-Agent"),
			texts.HideDetails,
			texts.ShowDetails)
		
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, response)
	})

	// Handler para endpoint de salud
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"status":"ok","timestamp":"`+time.Now().Format(time.RFC3339)+`"}`)
	})

	// Handler para endpoint de información
	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{
			"service": "hello-world-personalizado",
			"version": "1.0.0",
			"hostname": "%s",
			"timestamp": "%s",
			"urls": {
				"main": "/",
				"health": "/health",
				"info": "/info",
				"logs": "/logs"
			}
		}`, hostname, time.Now().Format(time.RFC3339))
	})

	// Handler para servir iconos SVG
	http.HandleFunc("/assets/", func(w http.ResponseWriter, r *http.Request) {
		// Servir archivos SVG desde el directorio assets
		http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))).ServeHTTP(w, r)
	})

	// Handler para logs en base64
	http.HandleFunc("/logs", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		
		// Crear log de ejemplo
		logContent := fmt.Sprintf(`[%s] INFO: Hello World Service started
[%s] INFO: Hostname: %s
[%s] INFO: Port: %s
[%s] INFO: Request from: %s
[%s] INFO: User-Agent: %s
[%s] INFO: Service: hello-world-personalizado
[%s] INFO: Version: 1.0.0
[%s] INFO: Author: Jean Reyes (@jemjaf)
[%s] INFO: Social: GitHub, TikTok, Instagram, LinkedIn, YouTube
[%s] INFO: Request processed successfully`, 
			currentTime, currentTime, hostname, currentTime, port,
			currentTime, r.RemoteAddr, currentTime, r.UserAgent(),
			currentTime, currentTime, currentTime, currentTime, currentTime)
		
		// Codificar en base64
		encodedLogs := base64.StdEncoding.EncodeToString([]byte(logContent))
		
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{
			"timestamp": "%s",
			"hostname": "%s",
			"logs_base64": "%s",
			"logs_decoded": "%s"
		}`, time.Now().Format(time.RFC3339), hostname, encodedLogs, logContent)
	})

	log.Printf("Servidor iniciado en puerto %s", port)
	log.Printf("URLs disponibles:")
	log.Printf("  - Principal: http://localhost:%s/", port)
	log.Printf("  - Salud: http://localhost:%s/health", port)
	log.Printf("  - Info: http://localhost:%s/info", port)
	log.Printf("  - Logs: http://localhost:%s/logs", port)
	
	log.Fatal(http.ListenAndServe(":"+port, nil))
}