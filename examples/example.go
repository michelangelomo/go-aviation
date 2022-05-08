package main

import (
	"fmt"

	"github.com/michelangelomo/go-aviation/aviation"
)

func main() {
	// Init client
	client := aviation.NewClient(nil)
	// Send request for metar
	metarOpts := aviation.MetarOptions{
		Stations:       "LIBP",
		HoursBeforeNow: "1",
	}
	metar, _, err := client.Metar.Get(metarOpts)
	if err != nil {
		fmt.Println(err)
	}

	for _, m := range metar.Data.METAR {
		fmt.Println(m.RawText)
	}

	tafOpts := aviation.TafOptions{
		Stations:       "LIBP",
		HoursBeforeNow: "1",
	}
	taf, _, err := client.Taf.Get(tafOpts)
	if err != nil {
		fmt.Println(err)
	}

	for _, t := range taf.Data.TAF {
		fmt.Println(t.RawText)
	}
}
