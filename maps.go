package pcmiler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Kamion/go-querystring/query"
)

type SingleMapRequest struct {
	Point1  Coordinates `url:"-"`
	Point2  Coordinates `url:"-"`
	Style   string      `url:"style,omitempty"`
	Width   int         `url:"width,omitempty"`
	Height  int         `url:"height,omitempty"`
	SRS     string      `url:"srs,omitempty"`
	Region  string      `url:"region,omitempty"`
	Dataset string      `json:"dataset,omitempty"`
	Imgsrc  string      `json:"imgsrc,omitempty"`
}

type MapRouteRequest struct {
	Map     Map     `json:"Map"`
	Routes  []Route `json:"Routes"`
	Dataset string  `json:"-"`
}

// SingleMap generates a map based on the request
// and returns the an image as the byte array
func (c *Client) SingleMap(request SingleMapRequest) ([]byte, error) {
	v, err := query.Values(request)
	if err != nil {
		return nil, err
	}

	url := "map?" + v.Encode()

	url += fmt.Sprintf("&pt1=%s,%s&pt2=%s,%s", request.Point1.Lat, request.Point1.Lon,
		request.Point2.Lat, request.Point2.Lon)

	return c.request(http.MethodGet, url, nil)
}

func (c *Client) MapRoute(request MapRouteRequest) ([]byte, error) {
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	url := "mapRoutes"
	if request.Dataset != "" {
		url += "?dataset=" + request.Dataset
	}

	return c.request(http.MethodPost, url, bytes.NewReader(body))
}
