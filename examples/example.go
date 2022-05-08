package main

import (
	"fmt"
	"time"

	"github.com/michelangelomo/go-aviation/aviation"
)

func main() {
	// Init client
	client := aviation.NewClient(nil)
	// Send request for metar
	opts := aviation.Options{}
	opts.SetStations("LIBP")
	opts.SetMostRecent(true)
	opts.SetHoursBeforeNow(1)
	metar, _, err := client.Metar.Get(opts)
	if err != nil {
		fmt.Println(err)
	}

	for _, m := range metar.Data.METAR {
		fmt.Println(m.RawText)
	}

	// Send request for taf
	opts = aviation.Options{}
	opts.SetStations("LIBP")
	opts.SetStartTime(time.Now().Add(-time.Hour * 12))
	opts.SetEndTime(time.Now())
	opts.SetMostRecent(true)
	taf, _, err := client.Taf.Get(opts)
	if err != nil {
		fmt.Println(err)
	}

	for _, t := range taf.Data.TAF {
		fmt.Println(t.RawText)
	}
}
