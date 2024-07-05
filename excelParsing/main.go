package main

import (
	"fmt"
)

func main() {

	meat, err := ParsingExcelFile()

	if err != nil {
		return
	}
	fmt.Println(meat)
}
