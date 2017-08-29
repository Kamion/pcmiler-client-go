package pcmiler

import (
	"fmt"
)

type Address struct {
	StreetAddress       string `json:"StreetAddress"`
	City                string `json:"City"`
	State               string `json:"State"`
	Zip                 string `json:"Zip"`
	County              string `json:"County"`
	Country             string `json:"Country"`
	SPLC                string `json:"SPLC"`
	CountryPostalFilter int    `json:"CountryPostalFilter"`
	AbbreviationFormat  int    `json:"AbbreviationFormat"`
}

func (a Address) String() string {
	return fmt.Sprintf("%s, %s, %s, %s", a.StreetAddress,
		a.City, a.State, a.Country)
}
