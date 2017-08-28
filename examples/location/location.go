package main

import (
	"log"
	"os"

	pcmiler "github.com/Kamion/pcmiler-client-go"
)

func main() {
	client := pcmiler.New(os.Getenv("PCMILER_API_KEY"))

	response, err := client.Geocode(pcmiler.GeocodeRequest{
		Coords: pcmiler.Coordinates{
			"-75.163244", "40.958188",
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v", response)
}
