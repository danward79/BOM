package bom

import (
	"log"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

// Observation details observation details
// TODO: Should readings below be converted to appropriate types?
type Observation struct {
	Name          string
	LastUpdate    string
	Temp          string
	AppTemp       string
	DewPoint      string
	RelHumidity   string
	DeltaTemp     string
	Dir           string
	Speedkph      string
	Gustkph       string
	Speedkts      string
	Gustkts       string
	PresshPa      string
	RainSince0900 string
	LowTemp       string
	HighTemp      string
	HighDir       string
	Highkph       string
	Highkts       string
}

// GetLatestObservations for all the stations on the URL provided.
func (o *Observer) GetLatestObservations() {

	c := colly.NewCollector(
		colly.AllowedDomains("www.bom.gov.au", "bom.gov.au"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML("[class=rowleftcolumn]", func(e *colly.HTMLElement) {
		var obs Observation

		obs.Name = e.ChildText("th")

		if obs.Name != "" {

			e.ForEach("td", func(i int, elem *colly.HTMLElement) {
				hdr := elem.Attr("headers")

				switch {
				case strings.Contains(hdr, "-datetime"):
					obs.LastUpdate = elem.Text
				case strings.Contains(hdr, "-tmp"):
					obs.Temp = elem.Text
				case strings.Contains(hdr, "-apptmp"):
					obs.AppTemp = elem.Text
				case strings.Contains(hdr, "-dewpoint"):
					obs.DewPoint = elem.Text
				case strings.Contains(hdr, "-relhum"):
					obs.RelHumidity = elem.Text
				case strings.Contains(hdr, "-delta-t"):
					obs.DeltaTemp = elem.Text
				case strings.Contains(hdr, "-wind-dir"):
					obs.Dir = elem.Text
				case strings.Contains(hdr, "-wind-dir"):
					obs.Dir = elem.Text
				case strings.Contains(hdr, "-wind-spd-kmh"):
					obs.Speedkph = elem.Text
				case strings.Contains(hdr, "-wind-gust-kmh"):
					obs.Gustkph = elem.Text
				case strings.Contains(hdr, "-wind-spd-kts"):
					obs.Speedkts = elem.Text
				case strings.Contains(hdr, "-wind-gust-kts"):
					obs.Gustkts = elem.Text
				case strings.Contains(hdr, "-press"):
					obs.PresshPa = elem.Text
				case strings.Contains(hdr, "-rainsince9am"):
					obs.RainSince0900 = elem.Text
				case strings.Contains(hdr, "-lowtmp"):
					obs.LowTemp = elem.Text
				case strings.Contains(hdr, "-hightmp"):
					obs.HighTemp = elem.Text
				case strings.Contains(hdr, "-highwind-dir"):
					obs.HighDir = elem.Text
				case strings.Contains(hdr, "-highwind-gust-kmh"):
					obs.Highkph = elem.Text
				case strings.Contains(hdr, "-highwind-gust-kts"):
					obs.Highkts = elem.Text
				default:
					log.Println("Unknown Attr:", hdr)
				}
			})
			o.LatestObservations[obs.Name] = obs
		}
		o.LastCheck = time.Now()
	})

	c.Visit(o.URL)
}
