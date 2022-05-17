package aviation

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestTafService_Get(t *testing.T) {
	client, mux := setupClient()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `<response><time_taken_ms>5</time_taken_ms><data num_results="1"><TAF><raw_text>TAF LIBP 171700Z 1718/1818 VRB05KT CAVOK PROB40 TEMPO 1718/1722 4000 TSRA BECMG 1806/1808 04010KT</raw_text><station_id>LIBP</station_id><forecast><wind_speed_kt>5</wind_speed_kt></forecast></TAF></data></response>`)
	})

	opts := Options{}
	opts.SetStations("LIBP")
	opts.SetMostRecent(true)
	opts.SetHoursBeforeNow(1)
	taf, _, err := client.Taf.Get(opts)
	if err != nil {
		t.Error(err)
	}

	want := &Taf{
		AWG: AWG{
			TimeTakenMs: "5",
		},
		Data: TafData{
			NumResults: "1",
			TAF: []TAF{
				{
					RawText:   "TAF LIBP 171700Z 1718/1818 VRB05KT CAVOK PROB40 TEMPO 1718/1722 4000 TSRA BECMG 1806/1808 04010KT",
					StationID: "LIBP",
					Forecast: []Forecast{
						{
							WindSpeedKt: "5",
						},
					},
				},
			},
		},
	}

	if !cmp.Equal(taf, want) {
		t.Errorf("mismatch %+v, want %+v", taf, want)
	}
}
