package handler

import (
	"log"
	"net/http"

	"github.com/JulianKerns/butchers-order/models"
)

func (cfg *HandlerConfig) GetDefaultTables(w http.ResponseWriter, r *http.Request) {
	// initializing the ResponseStruct
	type Response struct {
		MeatPriceTable []models.Meatprice `json:"Meatprices"`
	}

	// Getting the Database data
	meatPriceData, errMeatprice := cfg.DB.GetAllDefaultMeatPrices(r.Context())
	if errMeatprice != nil {
		log.Printf("Error: %v", errMeatprice)
		respondWithError(w, http.StatusInternalServerError, "Error retrieving database defaulkt data")
		return
	}
	// Setting the ResponseStruct

	response := Response{
		MeatPriceTable: models.ListDatabaseMeatPricetoMeatPrice(meatPriceData),
	}

	respondWithJSON(w, http.StatusOK, response)

}
