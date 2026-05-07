package main

import (
	"log"
	"net/http"
	"os"
)

func healthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9302"
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", healthz)
	addr := "0.0.0.0:" + port
	log.Printf("pulse-svc listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
