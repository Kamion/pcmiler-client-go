package pcmiler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Kamion/go-querystring/query"
)

func (c *Client) Geocode(request GeocodeRequest) ([]GeocodeResponse, error) {
	v, err := query.Values(request)
	if err != nil {
		return []GeocodeResponse{}, err
	}

	url := "locations?" + v.Encode()

	if request.Coords.Lat != "" && request.Coords.Lon != "" {
		url += fmt.Sprintf("&coords=%s,%s", request.Coords.Lat, request.Coords.Lon)
	}

	body, err := c.request(http.MethodGet, url, nil)
	if err != nil {
		return []GeocodeResponse{}, err
	}

	var response []GeocodeResponse
	err = json.Unmarshal(body, &response)

	return response, err
}
