// internal/service/status_service.go

package service

import (
	"math/rand"
	"time"

	"tugas3-hacktiv8/internal/model"
)

func GenerateRandomStatus() model.Status {
	rand.Seed(time.Now().UnixNano())
	water := rand.Intn(100) + 1
	wind := rand.Intn(100) + 1

	statusWater := ""
	statusWind := ""

	// Logika untuk menentukan status air
	if water < 5 {
		statusWater = "Aman"
	} else if water >= 6 && water <= 8 {
		statusWater = "Siaga"
	} else {
		statusWater = "Bahaya"
	}

	// Logika untuk menentukan status angin
	if wind < 6 {
		statusWind = "Aman"
	} else if wind >= 7 && wind <= 15 {
		statusWind = "Siaga"
	} else {
		statusWind = "Bahaya"
	}

	return model.Status{
		Water:       water,
		Wind:        wind,
		StatusWater: statusWater,
		StatusWind:  statusWind,
	}
}
