package planes

import (
	"fmt"

	"github.com/gocolly/colly"
)

func crawl(url string, path string) {
	c := colly.NewCollector()

	c.OnHTML("div[class='row']", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Println("Links found:", "-->", link)
	})

	c.Visit("https://www.aircraftcompare.com/")
}
