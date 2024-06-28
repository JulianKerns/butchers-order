package main

func main() {

	rind, schwein, surschwein, err := ParsingExcelFile()

	if err != nil {
		return
	}
	rind.Print()
	schwein.Print()
	surschwein.Print()
}
