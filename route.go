package pcmiler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Kamion/go-querystring/query"
)

type RouteRequest struct {
	Stops            []Coordinates `url:"-"`
	AvoidTolls       bool          `url:"avoidTolls,omitempty"`
	HubRouting       bool          `url:"hubRouting,omitempty"`
	VehHeight        string        `url:"vehHeight,omitempty"`
	VehLength        string        `url:"vehLength,omitempty"`
	StopsAsViaPoints bool          `url:"stopsAsViaPoints,omitempty"`
	AvoidFavors      bool          `url:"avoidFavors,omitempty"`
	AfSetIDs         string        `url:"afSetIDs,omitempty"`
	VehWeight        string        `url:"vehWeight,omitempty"`
	RouteOpt         string        `url:"routeOpt,omitempty"`
	RouteType        string        `url:"routeType,omitempty"`
	VehType          string        `url:"vehType,omitempty"`
	OverrideClass    string        `url:"overrideClass,omitempty"`
	Axles            int           `url:"axles,omitempty"`
	VehDimUnits      string        `url:"vehDimUnits,omitempty"`
	OpenBorders      bool          `url:"openBorders,omitempty"`
	LCV              bool          `url:"LCV,omitempty"`
	HighwayOnly      bool          `url:"hwyOnly,omitempty"`
	DistUnits        string        `url:"distUnits,omitempty"`
	OverrideRestrict bool          `url:"overrideRestrict,omitempty"`
	HazMat           string        `url:"hazMat,omitempty"`
	VehWidth         string        `url:"vehWidth,omitempty"`
	Region           string        `url:"region,omitempty"`
	Dataset          string        `url:"dataset,omitempty"`
}

type RouteResponse struct {
	Type      string   `json:"type"`
	Geometry  Geometry `json:"geometry"`
	TMinutes  int      `json:"TMinutes"`
	TDistance float64  `json:"TDistance"`
}

type Geometry struct {
	Type        string        `json:"type"`
	Coordinates [][][]float64 `json:"coordinates"`
}

func (c *Client) RoutePath(request RouteRequest) (RouteResponse, error) {
	result := RouteResponse{}

	v, err := query.Values(request)
	if err != nil {
		return result, err
	}

	url := "route/routepath?" + v.Encode()
	url += "&stops="

	for i, stop := range request.Stops {
		url += fmt.Sprintf("%s,%s", stop.Lon, stop.Lat)

		if i < len(request.Stops)-1 {
			url += ";"
		}
	}

	response, err := c.request(http.MethodGet, url, nil)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(response, &result)
	if err != nil {
		log.Println(string(response))
		return result, err
	}

	return result, nil
}
