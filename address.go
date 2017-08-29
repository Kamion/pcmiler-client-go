package pcmiler

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
