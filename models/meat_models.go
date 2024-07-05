package models

import (
	"github.com/JulianKerns/butchers-order/internal/database"
)

type Meatprice struct {
	ID         string  `json:"id"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
	MeatSource string  `json:"meatsource"`
	Meatcut    string  `json:"meatcut"`
	Price      float64 `json:"price"`
}

func DatabaseMeatPricetoMeatPrice(meatprice database.Meatprice) Meatprice {
	return Meatprice{
		ID:         meatprice.ID,
		CreatedAt:  meatprice.CreatedAt,
		UpdatedAt:  meatprice.UpdatedAt,
		MeatSource: meatprice.MeatSource,
		Meatcut:    meatprice.Meatcut,
		Price:      meatprice.Price,
	}

}

func ListDatabaseMeatPricetoMeatPrice(databaseSaltedPork []database.Meatprice) []Meatprice {
	var saltedPorkList []Meatprice = make([]Meatprice, len(databaseSaltedPork))

	for i, item := range databaseSaltedPork {
		changedStruct := DatabaseMeatPricetoMeatPrice(item)
		saltedPorkList[i] = changedStruct
	}
	return saltedPorkList
}
