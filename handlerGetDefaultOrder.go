package main

import (
	"log"
	"net/http"
)

func (cfg *Config) GetAllDefaultTables(w http.ResponseWriter, r *http.Request) {
	// initializing the ResponseStruct
	type Response struct {
		BeefTable       []Beef       `json:"Rindfleisch"`
		PorkTable       []Pork       `json:"Schwein"`
		SaltedPorkTable []Saltedpork `json:"Surschwein"`
	}

	// Getting the Database data
	beefData, errBeef := cfg.DB.GetAllDefaultBeef(r.Context())
	if errBeef != nil {
		log.Printf("Error: %v", errBeef)
		respondWithError(w, http.StatusInternalServerError, "Error retrieving database defaulkt data")
		return
	}

	porkData, errPork := cfg.DB.GetAllDefaultPork(r.Context())
	if errPork != nil {
		log.Printf("Error: %v", errPork)
		respondWithError(w, http.StatusInternalServerError, "Error retrieving database defaulkt data")
		return
	}

	saltedPorkData, errSaltedPork := cfg.DB.GetAllDefaultSaltedPork(r.Context())
	if errSaltedPork != nil {
		log.Printf("Error: %v", errSaltedPork)
		respondWithError(w, http.StatusInternalServerError, "Error retrieving database defaulkt data")
		return
	}
	log.Printf("%v", saltedPorkData)
	// Setting the ResponseStruct

	response := Response{
		BeefTable:       listDatabaseBeeftoBeef(beefData),
		PorkTable:       listDatabasePorktoPork(porkData),
		SaltedPorkTable: listDatabaseSaltedPorktoSaltedPork(saltedPorkData),
	}

	respondWithJSON(w, http.StatusOK, response)

}
