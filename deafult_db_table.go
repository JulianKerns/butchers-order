package main

import (
	"log"
	"time"

	"github.com/JulianKerns/butchers-order/internal/excelparse"

	"context"

	"github.com/JulianKerns/butchers-order/internal/database"
	"github.com/google/uuid"
)

const currentExcelFilePath string = "internal/excelparse/order_list.xlsx"

func GetDefaultMeatPrices() excelparse.Meats {
	// Getting the data from the excelfile
	meats, err := excelparse.ParsingExcelFile(currentExcelFilePath)
	if err != nil {
		log.Println("could not parse the excel file")
		return excelparse.Meats{}
	}
	return meats
}

func (cfg *excelConfig) DeleteAllMeatprices() error {

	errDelete := cfg.DB.DeleteAllEntries(context.Background())
	if errDelete != nil {
		log.Printf("Error: %v", errDelete)
		return errDelete
	}
	return nil
}

func (cfg *excelConfig) SettingDefaulttablesvalues() {
	// Getting the data from the excelfile
	meats := GetDefaultMeatPrices()

	now := time.Now().Format(time.DateTime)

	// Setting the beef table
	for _, value := range meats.Beef.Meats {

		_, err := cfg.DB.AddingDefaultMeatPrices(context.Background(), database.AddingDefaultMeatPricesParams{
			ID:         uuid.NewString(),
			CreatedAt:  now,
			UpdatedAt:  now,
			MeatSource: meats.Beef.Animal,
			Meatcut:    value.Name,
			Price:      value.PricePerKg,
		})
		if err != nil {
			log.Printf("Database error: %v\n", err)
			return
		}
	}
	// Setting the pork table
	for _, value := range meats.Pork.Meats {
		now := time.Now().Format(time.DateTime)
		_, err := cfg.DB.AddingDefaultMeatPrices(context.Background(), database.AddingDefaultMeatPricesParams{
			ID:         uuid.NewString(),
			CreatedAt:  now,
			UpdatedAt:  now,
			MeatSource: meats.Pork.Animal,
			Meatcut:    value.Name,
			Price:      value.PricePerKg,
		})
		if err != nil {
			log.Printf("Database error: %v\n", err)
			return
		}
	}
	// Setting the saltedpork table
	for _, value := range meats.SaltedPork.Meats {
		now := time.Now().Format(time.DateTime)
		_, err := cfg.DB.AddingDefaultMeatPrices(context.Background(), database.AddingDefaultMeatPricesParams{
			ID:         uuid.NewString(),
			CreatedAt:  now,
			UpdatedAt:  now,
			MeatSource: meats.SaltedPork.Animal,
			Meatcut:    value.Name,
			Price:      value.PricePerKg,
		})
		if err != nil {
			log.Printf("Database error: %v\n", err)
			return
		}
	}
}
