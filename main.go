package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"time"

	"github.com/JulianKerns/butchers-order/internal/database"
	"github.com/JulianKerns/butchers-order/internal/excelparse"
	"github.com/JulianKerns/butchers-order/internal/handler"

	_ "github.com/lib/pq"
)

type excelConfig struct {
	DB           *database.Queries
	CurrentMeats *excelparse.Meats
}

func main() {
	//Loading the .env file into our main function
	errLoadEnv := godotenv.Load(".env")
	if errLoadEnv != nil {
		log.Printf("Error: %s\n", errLoadEnv)
		return
	}
	// extrcating the value of the PORT and DB_Connetion key in the .env file
	port := os.Getenv("PORT")
	if port == "" {
		log.Println("Could not retrieve the environment variables")
		return
	}
	databaseConnection := os.Getenv("CONNECTION_STRING")
	if databaseConnection == "" {
		log.Println("Could not retrieve the environment variables")
		return
	}

	newDefaultDB := os.Getenv("NEW_DEFAULT")
	if newDefaultDB == "" {
		log.Println("Could not retrieve the environment variables")
		return
	}

	//Setting up the database connection
	db, err := sql.Open("postgres", databaseConnection)
	if err != nil {
		log.Println("Could not establish database connection")
		return
	}
	//setting up the queries for sqlc
	dbQueries := database.New(db)

	//Setting the GeneralConfig and the handler config to use the same db queries pointer
	handlerConfig := handler.GetHandlerConfig()
	handlerConfig.DB = dbQueries

	// getting the current Meatprices we use from the excelfile
	meats := GetDefaultMeatPrices()
	excelconfig := excelConfig{
		DB:           dbQueries,
		CurrentMeats: &meats,
	}

	// Using the Configs DBconnection to insert default values in the table
	if newDefaultDB == "true" {
		err := excelconfig.DeleteAllMeatprices()
		if err != nil {
			log.Fatalf("DatabaseError: %v", err)
		}
		excelconfig.SettingDefaulttablesvalues()
		log.Println("Setting new default table values")
	}

	// Spinning up a new Multiplexer that handles the requests and directs the traffic to the different handlers
	mux := http.NewServeMux()

	// adding the handlers for certain endpoints to the multiplexer
	mux.HandleFunc("GET /readiness", handler.ReadinessHandler)
	mux.HandleFunc("GET /errors", handler.ErrorHandler)

	//Getting the default table values
	mux.HandleFunc("GET /default", handlerConfig.GetDefaultTables)

	// Creating a User
	mux.HandleFunc("POST /users", handlerConfig.HandlerCreateUser)

	// Making an Order
	mux.HandleFunc("POST /orders", handlerConfig.HandlerPOSTOrder)

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
