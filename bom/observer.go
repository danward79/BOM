package bom

import (
	"fmt"
	"log"
	"time"
)

// Observer gathers and manages latest observations
// Latest Observations are avilable from:
// http://www.bom.gov.au/vic/observations/vicall.shtml
type Observer struct {
	CheckFrequency     time.Duration
	LastCheck          time.Time
	URL                string
	LatestObservations map[string]Observation
	ticker             *time.Ticker
}

//String prints Observer details
func (o *Observer) String() string {
	return fmt.Sprintf("Observer - Check Freq: %s, Last Check : %s, URL: %s", o.CheckFrequency, o.LastCheck, o.URL)
}

// New Observer starts a new observation routine
func New(t time.Duration, u string) Observer {
	return Observer{
		CheckFrequency:     t * time.Minute,
		URL:                u,
		LatestObservations: make(map[string]Observation),
	}
}

// Start check ticker, returns a <-chan time.Time,
// which should be used for any other downstream updates
func (o *Observer) Start() <-chan time.Time {
	o.ticker = time.NewTicker(o.CheckFrequency)

	return o.ticker.C
}

// Stop the check ticker
func (o *Observer) Stop() {
	o.ticker.Stop()
	log.Println("Check ticker stopped")
}
