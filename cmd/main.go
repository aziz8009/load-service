package main

import (
	"log"

	"github.com/aziz8009/load-service/cmd/http"
)

func main() {
	// Setup router (inisialisasi semua dependensi di dalamnya)
	r := http.SetupRouter()

	// Jalankan server di port 8080 (bisa disesuaikan)
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
