package pcmiler

type GeocodeResponse struct {
	Address          Address        `json:"Address"`
	Coords           Coordinates    `json:"Coords"`
	Region           int            `json:"Region"`
	Label            string         `json:"Label"`
	PlaceName        string         `json:"PlaceName"`
	Errors           []Error        `json:"Errors"`
	ConfidenceLevel  string         `json:"ConfidenceLevel"`
	DistanceFromRoad float64        `json:"DistanceFromRoad"`
	SpeedLimitInfo   SpeedLimitInfo `json:"SpeedLimitInfo"`
}

type GeocodeRequest struct {
	// Geocoding
	Street           string `url:"street,omitempty"`
	City             string `url:"city,omitempty"`
	State            string `url:"state,omitempty"`
	Postcode         string `url:"postcode,omitempty"`
	SearchString     string `url:"searchString,omitempty"`
	PostcodeFilter   string `url:"postcodeFilter,omitempty"`
	CitySearchFilter string `url:"citySearchFilter,omitempty"`
	Splc             string `url:"splc,omitempty"`
	List             int    `url:"list,omitempty"`

	// Reverse geocoding
	Coords                  Coordinates `url:"-"`
	MatchNamedRoadsOnly     bool        `url:"matchNamedRoadsOnly,omitempty"`
	MaxCleanupMiles         float64     `url:"maxCleanupMiles,omitempty"`
	IncludePostedSpeedLimit bool        `url:"includePostedSpeedLimit,omitempty"`
	Heading                 float32     `url:"heading,omitempty"`
	IncludeLinkInfo         bool        `url:"includeLinkInfo,omitempty"`
	Timestamp               string      `url:"timestamp,omitempty"`

	Region  string `url:"region,omitempty"`
	Dataset string `url:"dataset,omitempty"`
}