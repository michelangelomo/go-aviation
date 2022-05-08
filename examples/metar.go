package main

import (
	"fmt"

	"github.com/michelangelomo/go-aviation/aviation"
)

func main() {
	// Init client
	client := aviation.NewClient(nil)
	// Send request for metar
	opts := aviation.MetarOptions{
		Stations: "LIBP",
		HoursBeforeNow: "1",
	}
	metar, _, err := client.Metar.Get(opts)
	if err != nil {
		fmt.Println(err)
	}

	for _, m := range metar.Data.METAR {
		fmt.Println(m.RawText)
	}
}