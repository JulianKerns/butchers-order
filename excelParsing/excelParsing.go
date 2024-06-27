package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

func ParsingExcelFile() {

	file, err := excelize.OpenFile("order_list.xlsx", excelize.Options{})
	if err != nil {
		fmt.Printf("could not open the Excel file: %v", err)
	}
	defer file.Close()
	//fmt.Printf("%v", file)

	rows, errRows := file.GetRows("Tabelle1")
	if errRows != nil {
		fmt.Printf("could not read the rows: %v", err)
	}
	//fmt.Printf("%v\n", stripPrice(rows[2][1]))
	orders := &Rindfleisch{}

	for i := 2; i < 5; i++ {
		specificOrder := TypeOfMeat{
			name:       rows[i][0],
			pricePerKg: stripPrice((rows[i][1])),

			//	fmt.Printf("%v\n", stripPrice(rows[i][1]))

		}
		orders.Meats = append(orders.Meats, specificOrder)

	}
	fmt.Printf("%v\n", *orders)
}

func stripPrice(price string) float64 {
	var strippedString string

	if !strings.Contains(price, "kg") {
		stripped, _ := strings.CutSuffix(price, " €")

		if stripped == price {
			fmt.Println("Could not cut the string")
			return 0.0
		}
		strippedString = stripped

	} else {
		strippedStringKg, _ := strings.CutSuffix(price, " €/kg")
		if strippedStringKg == "" {
			fmt.Println("Could not cut the string")
			return 0.0
		}
		strippedString = strippedStringKg
	}

	sliceStrings := strings.Split(strippedString, "")

	for i, char := range sliceStrings {
		if char == "," {
			sliceStrings[i] = "."
		}
	}
	joinedStrings := strings.Join(sliceStrings, "")

	float, err := strconv.ParseFloat((joinedStrings), 64)
	if err != nil {
		fmt.Printf("Error parsing string: %v\n", err)
		return 0.0
	}
	return float

}
