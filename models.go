package main

import (
	"github.com/JulianKerns/butchers-order/internal/database"
)

type Beef struct {
	ID        string  `json:"id"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	Meatcut   string  `json:"meatcut"`
	Price     float64 `json:"price"`
	Quantitiy float64 `json:"quantity"`
}

type Pork struct {
	ID        string  `json:"id"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	Meatcut   string  `json:"meatcut"`
	Price     float64 `json:"price"`
	Quantitiy float64 `json:"quantity"`
}

type Saltedpork struct {
	ID        string  `json:"id"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	Meatcut   string  `json:"meatcut"`
	Price     float64 `json:"price"`
	Quantitiy float64 `json:"quantity"`
}

func databaseBeeftoBeef(beef database.Beef) Beef {
	return Beef{
		ID:        beef.ID,
		CreatedAt: beef.CreatedAt,
		UpdatedAt: beef.UpdatedAt,
		Meatcut:   beef.Meatcut,
		Price:     beef.Price,
		Quantitiy: beef.Quantitiy.Float64,
	}

}

func listDatabaseBeeftoBeef(databaseBeef []database.Beef) []Beef {
	var beefList []Beef = make([]Beef, len(databaseBeef))

	for i, item := range databaseBeef {
		changedStruct := databaseBeeftoBeef(item)
		beefList[i] = changedStruct
	}
	return beefList
}
func databasePorktoPork(pork database.Pork) Pork {
	return Pork{
		ID:        pork.ID,
		CreatedAt: pork.CreatedAt,
		UpdatedAt: pork.UpdatedAt,
		Meatcut:   pork.Meatcut,
		Price:     pork.Price,
		Quantitiy: pork.Quantitiy.Float64,
	}

}

func listDatabasePorktoPork(databasePork []database.Pork) []Pork {
	var porkList []Pork = make([]Pork, len(databasePork))

	for i, item := range databasePork {
		changedStruct := databasePorktoPork(item)
		porkList[i] = changedStruct
	}
	return porkList
}

func databaseSaltedPorktoSaltedPork(saltedPork database.Saltedpork) Saltedpork {
	return Saltedpork{
		ID:        saltedPork.ID,
		CreatedAt: saltedPork.CreatedAt,
		UpdatedAt: saltedPork.UpdatedAt,
		Meatcut:   saltedPork.Meatcut,
		Price:     saltedPork.Price,
		Quantitiy: saltedPork.Quantitiy.Float64,
	}

}

func listDatabaseSaltedPorktoSaltedPork(databaseSaltedPork []database.Saltedpork) []Saltedpork {
	var saltedPorkList []Saltedpork = make([]Saltedpork, len(databaseSaltedPork))

	for i, item := range databaseSaltedPork {
		changedStruct := databaseSaltedPorktoSaltedPork(item)
		saltedPorkList[i] = changedStruct
	}
	return saltedPorkList
}
