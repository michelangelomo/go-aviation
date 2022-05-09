package aviation

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMetarService_Get(t *testing.T) {
	client, mux := setupClient()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `<response><data num_results="1"><METAR><raw_text>LIBP 082020Z AUTO 24004KT 9999 SCT083/// OVC110/// 15/14 Q1020</raw_text><station_id>LIBP</station_id></METAR></data></response>`)
	})

	opts := Options{}
	opts.SetStations("LIBP")
	opts.SetMostRecent(true)
	opts.SetHoursBeforeNow(1)
	metar, _, err := client.Metar.Get(opts)
	if err != nil {
		t.Error(err)
	}

	want := &Metar{
		Data: Data{
			NumResults: "1",
			METAR: []METAR{
				{
					RawText:   "LIBP 082020Z AUTO 24004KT 9999 SCT083/// OVC110/// 15/14 Q1020",
					StationID: "LIBP",
				},
			},
		},
	}

	if !cmp.Equal(metar, want) {
		t.Errorf("mismatch %+v, want %+v", metar, want)
	}
}
