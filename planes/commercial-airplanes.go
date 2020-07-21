package planes

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gocolly/colly"
)

type commercialAirplanes struct {
	Name         string `json:"name"`
	Manufactured string `json:"manufactured"`
	Country      string `json:"country"`
	Performance  string `json:"performance"`
	Weight       string `json:"weight"`
	Dimensions   string `json:"dimensions"`
}

func scrapePrivateJetPlane(w http.ResponseWriter, r *http.Request) {
	c := colly.NewCollector()

	c.OnHTML("body > div.site > div:nth-child(1) > div.row.facetwp-template > b", func(e *colly.HTMLElement) {

		// set the struct properties
		commercial := new(commercialAirplanes)
		commercial.Name = ""
		commercial.Manufactured = ""
		commercial.Country = ""
		commercial.Performance = ""
		commercial.Weight = ""
		commercial.Dimensions = ""

		// serialzie the struct to JSON
		bf := bytes.NewBuffer([]byte{})
		jsonEncoder := json.NewEncoder(bf)
		jsonEncoder.SetEscapeHTML(false)
		jsonEncoder.Encode(commercial)

		// return JSON response with data given
		w.Header().Set("Content-Type", "application/json")
		w.Write(bf.Bytes())
	})
	c.Visit("COMMERCIAL-AIRPLANES")
}
