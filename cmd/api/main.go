package main

import (
	"net/http"
	"os"
	"log"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("api is working"))
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	srv := &http.Server{
		Addr:    ":"+port,	
		Handler: mux,
	}
	log.Printf("server running on http://localhost:%s", port)
	if err := srv.ListenAndServe(); err != nil {
    panic(err)
	}
}