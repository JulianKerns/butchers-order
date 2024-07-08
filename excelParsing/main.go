package main

import (
	//"fmt"
	"time"

	"github.com/JulianKerns/butchers-order/internal/auth"
)

func main() {

	// meat, err := ParsingExcelFile()
	//
	//	if err != nil {
	//		return
	//	}
	//
	// fmt.Println(meat)
	authConfig := auth.GetAuthConfig()
	authConfig.GenerateJWTToken(time.Minute, "1")
}
