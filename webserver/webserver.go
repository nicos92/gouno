package webserver

import (
	"log"
	"net/http"
)

const PORT = ":3003"

func MiWebServer() {
	http.HandleFunc("/", home)
	log.Printf("\nServidor Web iniciado en http://localhost%s", PORT)
	http.ListenAndServe(PORT, nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}
