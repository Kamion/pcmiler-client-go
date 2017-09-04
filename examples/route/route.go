package main

import (
	"fmt"
	"log"
	"os"

	pcmiler "github.com/Kamion/pcmiler-client-go"
)

func main() {
	client := pcmiler.New(os.Getenv("PCMILER_API_KEY"))

	stops := []pcmiler.Coordinates{}
	stops = append(stops, pcmiler.Coordinates{
		Lat: "39.942892",
		Lon: "-75.173297",
	})

	stops = append(stops, pcmiler.Coordinates{
		Lat: "39.61703",
		Lon: "-74.83153",
	})

	stops = append(stops, pcmiler.Coordinates{
		Lat: "39.362469",
		Lon: "-74.438942",
	})

	request := pcmiler.RouteRequest{
		Stops: stops,
	}

	response, err := client.RoutePath(request)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v", response)
}
