package pcmiler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Kamion/go-querystring/query"
)

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

func (c *Client) Geocode(request GeocodeRequest) ([]GeocodeResponse, error) {
	v, err := query.Values(request)
	if err != nil {
		return []GeocodeResponse{}, err
	}

	url := "locations?" + v.Encode()

	if request.Coords.Lat != "" && request.Coords.Lon != "" {
		url += fmt.Sprintf("&coords=%s,%s", request.Coords.Lon, request.Coords.Lat)
	}

	body, err := c.request(http.MethodGet, url, nil)
	if err != nil {
		return []GeocodeResponse{}, err
	}

	var response []GeocodeResponse
	err = json.Unmarshal(body, &response)

	return response, err
}
