package main

import (
	"bytes"
	"image"
	_ "image/png"
	"log"
	"os"

	pcmiler "github.com/Kamion/pcmiler-client-go"
)

func main() {
	client := pcmiler.New(os.Getenv("PCMILER_API_KEY"))

	// single map
	imageRaw, err := client.SingleMap(pcmiler.SingleMapRequest{
		Point1: pcmiler.Coordinates{Lat: "-74.654726", Lon: "40.38825"},
		Point2: pcmiler.Coordinates{Lat: "-76.12345", Lon: "42.12345"},
		Width:  600,
		Height: 600,
	})

	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(bytes.NewReader(imageRaw))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v", img.Bounds())

	// map route
	pins := []pcmiler.Pin{}
	pins = append(pins, pcmiler.Pin{
		Point: pcmiler.Coordinates{
			Lat: "41.63411",
			Lon: "-87.96074",
		},
		Image: "ltruck_r",
	})
	pins = append(pins, pcmiler.Pin{
		Point: pcmiler.Coordinates{
			Lat: "25.75312",
			Lon: "-80.29229",
		},
		Image: "lbldg_bl",
	})

	request := pcmiler.MapRouteRequest{}
	request.Map = pcmiler.Map{
		Width:     1024,
		Height:    768,
		Drawers:   []int{8, 2, 7, 17, 15},
		PinDrawer: pcmiler.PinDrawer{Pins: pins},
	}
	request.Map.LegendDrawer = append(request.Map.LegendDrawer, pcmiler.Legend{
		DrawOnMap: true,
		Type:      0,
	})

	stops := []pcmiler.StopLocation{}
	stops = append(stops, pcmiler.StopLocation{
		Address: pcmiler.Address{
			City:  "Chicago",
			State: "IL",
		},
	})
	stops = append(stops, pcmiler.StopLocation{
		Address: pcmiler.Address{
			City:  "Miami",
			State: "FL",
		},
	})

	request.Routes = append(request.Routes, pcmiler.Route{
		Stops: stops,
		Options: pcmiler.Options{
			HighwayOnly: true,
		},
	})

	routeRaw, err := client.MapRoute(request)
	if err != nil {
		log.Fatal(err)
	}

	imgRoute, _, err := image.Decode(bytes.NewReader(routeRaw))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v", imgRoute.Bounds())

}
