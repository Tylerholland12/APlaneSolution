package planes

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gocolly/colly"
)

type PropellorPlaneInfo struct {
	Name         string `json:"name"`
	Manufactured string `json:"manufactured"`
	Country      string `json:"country"`
	Performance  string `json:"performance"`
	Weight       string `json:"weight"`
	Dimensions   string `json:"dimensions"`
}

func scrapePropellorPlane(w http.ResponseWriter, r *http.Request) {
	c := colly.NewCollector()

	c.OnHTML("body > div.site > div:nth-child(1) > div.row.facetwp-template > b", func(e *colly.HTMLElement) {

		// set the struct properties
		propellorPlane := new(PropellorPlaneInfo)
		propellorPlane.Name = ""
		propellorPlane.Manufactured = ""
		propellorPlane.Country = ""
		propellorPlane.Performance = ""
		propellorPlane.Weight = ""
		propellorPlane.Dimensions = ""

		// serialzie the struct to JSON
		bf := bytes.NewBuffer([]byte{})
		jsonEncoder := json.NewEncoder(bf)
		jsonEncoder.SetEscapeHTML(false)
		jsonEncoder.Encode(propellorPlane)

		// return JSON response with data given
		w.Header().Set("Content-Type", "application/json")
		w.Write(bf.Bytes())
	})
	c.Visit("PRIVATE-PROPELLOR-PLANES")
}
