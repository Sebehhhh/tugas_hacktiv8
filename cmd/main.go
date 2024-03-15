// cmd/main.go

package main

import (
	"net/http"

	"tugas3-hacktiv8/internal/handler"
)

func main() {
	// Mulai auto-reload status setiap 15 detik
	handler.StartAutoReload()

	// Routing HTTP
	http.HandleFunc("/status", handler.StatusHandler)
	http.HandleFunc("/", handler.IndexHandler)

	// Jalankan server HTTP
	http.ListenAndServe(":8080", nil)
}
