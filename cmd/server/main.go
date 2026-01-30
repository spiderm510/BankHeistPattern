package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"case.cubi.bankheist/internal/handler"
	"case.cubi.bankheist/internal/service"
	"case.cubi.bankheist/internal/storage"
)

func main() {
	log.Println("Monopoly Go - Bank Heist Pattern Recognition Service starting...")

	path := dataPath("patterns.json")
	storage := storage.NewFileStorage(path)
	pm := service.NewPatternManager(storage)

	log.Println("Starting API server on :8080")
	http.Handle("/api/predict", handler.NewPredictHandler(pm))
	http.Handle("/api/update", handler.NewUpateHandler(pm))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func dataPath(filename string) string {
	exe, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	base := filepath.Dir(exe)
	return filepath.Join(base, "data", filename)
}
