package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"time"

	"github.com/JulianKerns/butchers-order/internal/handler"
)

// type config struct {
//
// }

func main() {
	errLoadEnv := godotenv.Load(".env")
	if errLoadEnv != nil {
		log.Printf("Error: %s\n", errLoadEnv)
		return
	}
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("Could not retrieve the environment variables")
		return
	}
	mux := http.NewServeMux()

	mux.HandleFunc("GET /readiness", handler.ReadinessHandler)
	mux.HandleFunc("GET /errors", handler.ErrorHandler)

	server := &http.Server{
		Addr:              ":" + port,
		Handler:           mux,
		ReadHeaderTimeout: time.Minute,
	}

	log.Printf("Serving on Port: %s", port)
	log.Fatal(server.ListenAndServe())

}
