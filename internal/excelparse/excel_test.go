package excelparse

import (
	"fmt"
	"testing"
)

func TestStripNameWhitespaces(t *testing.T) {
	cases := []struct {
		have string
		want string
	}{
		{
			have: "Faschiert Gemischt",
			want: "FaschiertGemischt",
		},
		{
			have: "Schnitzel",
			want: "Schnitzel",
		},
		{
			have: "Packung 10 kg",
			want: "Packung10kg",
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Testing case %v", i), func(t *testing.T) {
			got := stripNameWhitespaces(c.have)
			if got != c.want {
				t.Errorf("expected the values to be equal")
			}

		})
	}

}

func TestStripPrice(t *testing.T) {
	cases := []struct {
		have string
		want float64
	}{
		{
			have: "18,5 €/kg",
			want: 18.5,
		},
		{
			have: "120 €",
			want: 120,
		},
		{
			have: "19,9 €/kg",
			want: 19.9,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Testing case %v", i), func(t *testing.T) {
			got, err := stripPrice(c.have)
			if err != nil {
				t.Errorf("Could not convert the string to a float64")
			}
			if c.want != got {
				t.Errorf("expected values to be equal")
			}
		})
	}
}
