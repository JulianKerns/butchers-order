package excelparse

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

func ParsingExcelFile() (Rindfleisch, Schwein, SurSchwein, error) {
	// opening the excel file
	file, err := excelize.OpenFile("order_list.xlsx", excelize.Options{})
	if err != nil {
		fmt.Printf("could not open the Excel file: %v", err)
	}
	defer file.Close()

	// Reading from the excel file
	rows, errRows := file.GetRows("Tabelle1")
	if errRows != nil {
		fmt.Printf("could not read the rows: %v", err)
	}
	// initializing Meat structs
	tableRind := &Rindfleisch{Meats: make(map[string]TypeOfMeat)}
	tableSchwein := &Schwein{Meats: make(map[string]TypeOfMeat)}
	tableSurSchwein := &SurSchwein{Meats: make(map[string]TypeOfMeat)}

	//setting the rindfleisch table
	for i := 2; i < 20; i++ {
		strippedPrice, err := stripPrice((rows[i][1]))
		if err != nil {
			return Rindfleisch{}, Schwein{}, SurSchwein{}, err
		}
		specificOrder := TypeOfMeat{
			Name:       stripNameWhitespaces(rows[i][0]),
			PricePerKg: strippedPrice,
		}
		tableRind.Meats[stripNameWhitespaces(rows[i][0])] = specificOrder
	}

	//setting the Schwein table
	for i := 2; i < 14; i++ {
		strippedPrice, err := stripPrice((rows[i][5]))
		if err != nil {
			return Rindfleisch{}, Schwein{}, SurSchwein{}, err
		}
		specificOrder := TypeOfMeat{
			Name:       stripNameWhitespaces(rows[i][4]),
			PricePerKg: strippedPrice,
		}
		tableSchwein.Meats[stripNameWhitespaces(rows[i][4])] = specificOrder

	}

	//setting the SurSchwein table
	for i := 16; i < 19; i++ {
		strippedPrice, err := stripPrice((rows[i][5]))
		if err != nil {
			return Rindfleisch{}, Schwein{}, SurSchwein{}, err
		}
		specificOrder := TypeOfMeat{
			Name:       stripNameWhitespaces(rows[i][4]),
			PricePerKg: strippedPrice,
		}
		tableSurSchwein.Meats[stripNameWhitespaces(rows[i][4])] = specificOrder

	}
	return *tableRind, *tableSchwein, *tableSurSchwein, nil
}

func stripPrice(price string) (float64, error) {
	// setting the variable that get different value depending on the price string content
	var strippedString string

	if !strings.Contains(price, "kg") {
		stripped, _ := strings.CutSuffix(price, " €")

		if stripped == price {
			fmt.Println("Could not cut the string")
			return 0.0, errors.New("could not cut the string")
		}
		strippedString = stripped

	} else {
		strippedStringKg, _ := strings.CutSuffix(price, " €/kg")
		if strippedStringKg == "" {
			fmt.Println("Could not cut the string")
			return 0.0, errors.New("could not cut the string")
		}
		strippedString = strippedStringKg
	}
	// Splitting the string so we can change the comma to a full stop
	sliceStrings := strings.Split(strippedString, "")

	for i, char := range sliceStrings {
		if char == "," {
			sliceStrings[i] = "."
		}
	}
	//Joining the string back together
	joinedStrings := strings.Join(sliceStrings, "")
	//Converting the string into a float64
	float, err := strconv.ParseFloat((joinedStrings), 64)
	if err != nil {
		fmt.Printf("Error parsing string: %v\n", err)
		return 0.0, errors.New("could not parse the string")
	}
	return float, nil

}

// if the name of the meatcut has whitespaces in them we want to remove them
func stripNameWhitespaces(name string) string {
	if !strings.Contains(name, " ") {
		return name
	} else {
		newString := strings.ReplaceAll(name, " ", "")
		return newString
	}

}
