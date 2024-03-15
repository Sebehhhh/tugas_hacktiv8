// internal/handler/handler.go

package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"tugas3-hacktiv8/internal/service"
	"tugas3-hacktiv8/internal/utils"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	status := service.GenerateRandomStatus()

	// Simpan status ke file JSON
	err := utils.WriteStatusToFile(status, "status.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Mengirim status sebagai respons HTTP
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func StartAutoReload() {
	ticker := time.NewTicker(15 * time.Second)
	go func() {
		for range ticker.C {
			status := service.GenerateRandomStatus()
			_ = utils.WriteStatusToFile(status, "status.json")
		}
	}()
}
