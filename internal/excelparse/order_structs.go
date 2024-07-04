package excelparse

import (
	"fmt"
)

type Meat interface {
	Print()
}

type TypeOfMeat struct {
	Name       string
	PricePerKg float64
}

type Rindfleisch struct {
	Meats  map[string]TypeOfMeat
	Animal string
}

type Schwein struct {
	Meats  map[string]TypeOfMeat
	Animal string
}

type SurSchwein struct {
	Meats  map[string]TypeOfMeat
	Animal string
}

func (r Rindfleisch) Print() {
	fmt.Printf("Printing Struct: %v\n", r)
}
func (s Schwein) Print() {
	fmt.Printf("Printing Struct: %v\n", s)
}

func (sur SurSchwein) Print() {
	fmt.Printf("Printing Struct: %v\n", sur)
}

type Beef struct {
	Schnitzel struct {
		pricePerKg float64
	}
	Gulasch struct {
		pricePerKg float64
	}
	Wadschinken struct {
		pricePerKg float64
	}
	Suppenfleisch struct {
		pricePerKg float64
	}
	RindfleischZumKochen struct {
		pricePerKg float64
	}
	Hüferschwanz struct {
		pricePerKg float64
	}
	Schultermeisel struct {
		pricePerKg float64
	}
	Schulterscherzel struct {
		pricePerKg float64
	}
	WeißesScherzel struct {
		pricePerKg float64
	}
	Beiried struct {
		pricePerKg float64
	}
	Rostbraten struct {
		pricePerKg float64
	}
	Tafelspitz struct {
		pricePerKg float64
	}
	Lungenbraten struct {
		pricePerKg float64
	}
	Paket5kg struct {
		pricePerKg float64
	}
	Paket10kg struct {
		pricePerKg float64
	}
	Hundefutter struct {
		pricePerKg float64
	}
	Rinderbraten struct {
		pricePerKg float64
	}
	Rinderknochen struct {
		pricePerKg float64
	}
}

type Pork struct {
	Schnitzel struct {
		pricePerKg float64
	}
	Karree struct {
		pricePerKg float64
	}
	Schopfbraten struct {
		pricePerKg float64
	}
	Schulterfleisch struct {
		price float64
	}
	DurchzogenesBauchfleisch struct {
		pricePerKg float64
	}
	Lungenbraten struct {
		pricePerKg float64
	}
	Schweineleber struct {
		pricePerKg float64
	}
	Schmalz struct {
		pricePerKg float64
	}
	Schweinehälfte struct {
		pricePerKg float64
	}
	Schweineknochen struct {
		pricePerKg float64
	}
	Faschiertes struct {
		pricePerKg float64
	}
	FaschiertesGemischt struct {
		pricePerKg float64
	}
}

type SurPork struct {
	Surstelze struct {
		pricePerKg float64
	}
	SurfleischMager struct {
		pricePerKg float64
	}
	SurfleischDurchzogen struct {
		pricePerKg float64
	}
}
