package handler

import (
	"log"
	"net/http"

	"github.com/JulianKerns/butchers-order/models"
)

func (cfg *HandlerConfig) GetAllDefaultTables(w http.ResponseWriter, r *http.Request) {
	// initializing the ResponseStruct
	type Response struct {
		BeefTable       []models.Beef       `json:"Rindfleisch"`
		PorkTable       []models.Pork       `json:"Schwein"`
		SaltedPorkTable []models.Saltedpork `json:"Surschwein"`
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

	// Setting the ResponseStruct

	response := Response{
		BeefTable:       models.ListDatabaseBeeftoBeef(beefData),
		PorkTable:       models.ListDatabasePorktoPork(porkData),
		SaltedPorkTable: models.ListDatabaseSaltedPorktoSaltedPork(saltedPorkData),
	}

	respondWithJSON(w, http.StatusOK, response)

}
