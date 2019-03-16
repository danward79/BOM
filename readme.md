# BOM Observations

Bureau of Meteorology(BOM) provides interfaces and webpages which allow data to be grabbed from various places around Australia. They also provide the latest observations every 10 minutes from weather stations around Australia.

At the time of writting the more frequent data available upto every 10 minutes, does not seem to be available for free in JSON format. So this work provides a method to obtain these readings.

## Details about the observations
Details about the observations can be found at ["About Latest Weather Observations"](http://www.bom.gov.au/catalogue/observations/about-weather-observations.shtml)

There appears to be several types of data.

### Fawkner Beacon - Example Data
Available at 30 minute intervals [in readable format](http://www.bom.gov.au/products/IDV60901/IDV60901.95872.shtml) and in [JSON](http://www.bom.gov.au/fwo/IDV60901/IDV60901.95872.json) or more frequently every 10 minutes or so, [here](http://www.bom.gov.au/vic/observations/vicall.shtml).

## How to use this code
```go
func main() {
	url := "http://www.bom.gov.au/vic/observations/vicall.shtml"

	observer := bom.New(10, url)

	observer.GetLatestObservations()
	c := observer.Start()

	for t := range c {
		log.Println("Check at", t)
		fmt.Println(observer.LatestObservations["St Kilda Harbour RMYS"], observer.LastCheck)
		fmt.Println(observer.LatestObservations["Fawkner Beacon"], observer.LastCheck)
		fmt.Println(observer.LatestObservations["South Channel Island"], observer.LastCheck)
	}

}
```
