package planes

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gocolly/colly"
)

type privateJetsInfo struct {
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
		privateJet := new(privateJetsInfo)
		privateJet.Name = ""
		privateJet.Manufactured = ""
		privateJet.Country = ""
		privateJet.Performance = ""
		privateJet.Weight = ""
		privateJet.Dimensions = ""

		// serialzie the struct to JSON
		bf := bytes.NewBuffer([]byte{})
		jsonEncoder := json.NewEncoder(bf)
		jsonEncoder.SetEscapeHTML(false)
		jsonEncoder.Encode(privateJet)

		// return JSON response with data given
		w.Header().Set("Content-Type", "application/json")
		w.Write(bf.Bytes())
	})
	c.Visit("PRIVATE-JETS")
}
