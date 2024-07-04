package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/JulianKerns/butchers-order/internal/database"
	"github.com/JulianKerns/butchers-order/models"
	"github.com/google/uuid"
)

func (cfg *HandlerConfig) HandlerPOSTOrder(w http.ResponseWriter, r *http.Request) {
	type Parameter struct {
		UserID     string  `json:"user_id"`
		MeatSource string  `json:"meat_source"`
		Meatcut    string  `json:"meatcut"`
		Price      float64 `json:"price"`
		Quantitiy  float64 `json:"quantity"`
	}
	params := Parameter{}
	decoder := json.NewDecoder(r.Body)
	errDecode := decoder.Decode(&params)
	if errDecode != nil {
		log.Printf("Error %v", errDecode)
		respondWithError(w, http.StatusBadRequest, "Bad Request sent")
		return
	}

	//write the order to the database

	now := time.Now().Format(time.DateTime)
	order, err := cfg.DB.MakeOrder(r.Context(), database.MakeOrderParams{
		ID:         uuid.NewString(),
		CreatedAt:  now,
		UpdatedAt:  now,
		UserID:     params.UserID,
		MeatSource: params.MeatSource,
		Meatcut:    params.Meatcut,
		Price:      params.Price,
		Quantitiy:  params.Quantitiy,
	})

	if err != nil {
		log.Printf("Error: %v", err)
		respondWithError(w, http.StatusInternalServerError, "could not save the order")
		return
	}

	respondWithJSON(w, http.StatusCreated, models.DatabaseOrdertoOrder(order))

}
