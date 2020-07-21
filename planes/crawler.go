package planes

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func crawlWebsite() {
	c := colly.NewCollector(
		colly.AllowedDomains("aircraftcompare.com"),
	)

	c.OnHTML("article", func(e *colly.HTMLElement) {
		isAirplanePage := false

		// extract the meta tags from the page
		e.DOM.ParentsUntil("~").Find("meta").Each(func(_ int, s *goquery.Selection) {
			// search for og:type meta tags 
			if property, _ := s.Attr("property"); strings.EqualFold(property, "og:type") {
				content, _ := s.Attr("content")

				isAirplanePage = strings.EqualFold(content, "article")
			}

		})

		e.DOM.ParentsUntil("~").Find("link").Each(func(_ int, s *goquery.Selection) {
			// search for link
			if rel, _ := s.Attr("rel"); strings.EqualFold(rel, "stylesheet") {
				id, _ := s.Attr("id")
				href, _ := s.Attr("href")
			}
		})

		e.DOM.ParentsUntil("~").Find("link").Each(func(_, int, s *goquery.Selection) {
			// search for link
			if rel, _ := s.Attr("rel"); strings.EqualFold(rel, "dns-prefetch") {
				href, _ := s.Attr("href")
			}
		})

		e.DOM.ParentsUntil("~").Find("link").Each(func(_, int, s *goquery.Selection){
			// search for link
			if rel, _ := s.Attr("rel"); strings.EqualFold(rel, "alternate") {
				type, _ := s.Attr("type")
			}
		})

		if isAirplanePage {
			// find the titles
			fmt.Println("Type of airplane:", e.Dom.Find("row.row-small-gutter").Text())
			// find href images
			fmt.Println("Links:", e.DOM.Find("body > main > div.container > div.row > div.col-md-8.content-area > div.site-main > selection.row.row-small-gutter > img.src").Text())
		}
}
