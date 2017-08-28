package pcmiler

type Address struct {
	StreetAddress       string `json:"StreetAddress"`
	City                string `json:"City"`
	State               string `json:"state"`
	Zip                 string `json:"zip"`
	County              string `json:"County"`
	Country             string `json:"Country"`
	SPLC                string `json:"SPLC"`
	CountryPostalFilter int    `json:"CountryPostalFilter"`
	AbbreviationFormat  int    `json:"AbbreviationFormat"`
}
