package pcmiler

type TruckConfig struct {
	Axles  int  `json:"Axles"`
	Height int  `json:"Height"`
	LCV    bool `json:"LCV"`
	Length int  `json:"Length"`
	Units  int  `json:"Units"`
	Weight int  `json:"Weight"`
	Width  int  `json:"Width"`
}

type FuelOptions struct {
	UserID   string  `json:"UserId"`
	Password string  `json:"Password"`
	Account  string  `json:"Account"`
	FuelCap  float64 `json:"FuelCap"`
	Level    float64 `json:"Level"`
	MPG      float64 `json:"MPG"`
}
