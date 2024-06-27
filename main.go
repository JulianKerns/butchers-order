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
	//Loading the .env file into our main function
	errLoadEnv := godotenv.Load(".env")
	if errLoadEnv != nil {
		log.Printf("Error: %s\n", errLoadEnv)
		return
	}
	// extrcating the value of the PORT key in the .env file
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("Could not retrieve the environment variables")
		return
	}
	// Spinning up a new Multiplexer that handles the requests and directs the traffic to the different handlers
	mux := http.NewServeMux()

	// adding the handlers for certain endpoints to the multiplexer
	mux.HandleFunc("GET /readiness", handler.ReadinessHandler)
	mux.HandleFunc("GET /errors", handler.ErrorHandler)

	//Creating th server Struct with the port as adress and the multiplexer as our handler
	server := &http.Server{
		Addr:              ":" + port,
		Handler:           mux,
		ReadHeaderTimeout: time.Second * 30,
	}

	//Staring up the Server
	log.Printf("Serving on Port: %s", port)
	log.Fatal(server.ListenAndServe())

}
