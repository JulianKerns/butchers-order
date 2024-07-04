package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"time"

	"github.com/JulianKerns/butchers-order/internal/database"
	"github.com/JulianKerns/butchers-order/internal/excelparse"
	"github.com/JulianKerns/butchers-order/internal/handler"

	"github.com/google/uuid"

	_ "github.com/lib/pq"
)

type Config struct {
	DB *database.Queries
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

	config := Config{
		DB: dbQueries,
	}

	// Using the Configs DBconnection to insert default values in the table
	if newDefaultDB == "true" {
		config.SettingDefaulttablesvalues()
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

func (cfg *Config) SettingDefaulttablesvalues() {
	// Getting the data from the excelfile
	beef, pork, saltedpork, err := excelparse.ParsingExcelFile("internal/excelparse/order_list.xlsx")
	if err != nil {
		log.Println("could not parse the excel file")
		return
	}
	now := time.Now().Format(time.DateTime)

	// Setting the beef table
	for _, value := range beef.Meats {

		_, err := cfg.DB.AddingDefaultMeatPrices(context.Background(), database.AddingDefaultMeatPricesParams{
			ID:         uuid.NewString(),
			CreatedAt:  now,
			UpdatedAt:  now,
			MeatSource: beef.Animal,
			Meatcut:    value.Name,
			Price:      value.PricePerKg,
		})
		if err != nil {
			log.Printf("Database error: %v\n", err)
			return
		}
	}
	// Setting the pork table
	for _, value := range pork.Meats {
		now := time.Now().Format(time.DateTime)
		_, err := cfg.DB.AddingDefaultMeatPrices(context.Background(), database.AddingDefaultMeatPricesParams{
			ID:         uuid.NewString(),
			CreatedAt:  now,
			UpdatedAt:  now,
			MeatSource: pork.Animal,
			Meatcut:    value.Name,
			Price:      value.PricePerKg,
		})
		if err != nil {
			log.Printf("Database error: %v\n", err)
			return
		}
	}
	// Setting the saltedpork table
	for _, value := range saltedpork.Meats {
		now := time.Now().Format(time.DateTime)
		_, err := cfg.DB.AddingDefaultMeatPrices(context.Background(), database.AddingDefaultMeatPricesParams{
			ID:         uuid.NewString(),
			CreatedAt:  now,
			UpdatedAt:  now,
			MeatSource: saltedpork.Animal,
			Meatcut:    value.Name,
			Price:      value.PricePerKg,
		})
		if err != nil {
			log.Printf("Database error: %v\n", err)
			return
		}
	}

}
