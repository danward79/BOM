package bom

// This file details the JSON station observations available from the BOM at, for example:
// http://www.bom.gov.au/fwo/IDV60801/IDV60801.95872.json
//
// This is not implemented at this time. As only live data was required.

// StationData ...
type StationData struct {
	Observations Observations `json:"observations"`
}

// Notice ...
type Notice struct {
	Copyright     string `json:"copyright"`
	CopyrightURL  string `json:"copyright_url"`
	DisclaimerURL string `json:"disclaimer_url"`
	FeedbackURL   string `json:"feedback_url"`
}

// Header ...
type Header struct {
	RefreshMessage string `json:"refresh_message"`
	ID             string `json:"ID"`
	MainID         string `json:"main_ID"`
	Name           string `json:"name"`
	StateTimeZone  string `json:"state_time_zone"`
	TimeZone       string `json:"time_zone"`
	ProductName    string `json:"product_name"`
	State          string `json:"state"`
}

// Data ...
type Data struct {
	SortOrder         int         `json:"sort_order"`
	Wmo               int         `json:"wmo"`
	Name              string      `json:"name"`
	HistoryProduct    string      `json:"history_product"`
	LocalDateTime     string      `json:"local_date_time"`
	LocalDateTimeFull string      `json:"local_date_time_full"`
	AifstimeUtc       string      `json:"aifstime_utc"`
	Lat               float64     `json:"lat"`
	Lon               float64     `json:"lon"`
	ApparentT         interface{} `json:"apparent_t"`
	Cloud             string      `json:"cloud"`
	CloudBaseM        interface{} `json:"cloud_base_m"`
	CloudOktas        interface{} `json:"cloud_oktas"`
	CloudTypeID       interface{} `json:"cloud_type_id"`
	CloudType         string      `json:"cloud_type"`
	DeltaT            interface{} `json:"delta_t"`
	GustKmh           int         `json:"gust_kmh"`
	GustKt            int         `json:"gust_kt"`
	AirTemp           interface{} `json:"air_temp"`
	Dewpt             interface{} `json:"dewpt"`
	Press             interface{} `json:"press"`
	PressQnh          interface{} `json:"press_qnh"`
	PressMsl          interface{} `json:"press_msl"`
	PressTend         string      `json:"press_tend"`
	RainTrace         string      `json:"rain_trace"`
	RelHum            interface{} `json:"rel_hum"`
	SeaState          string      `json:"sea_state"`
	SwellDirWorded    string      `json:"swell_dir_worded"`
	SwellHeight       interface{} `json:"swell_height"`
	SwellPeriod       interface{} `json:"swell_period"`
	VisKm             string      `json:"vis_km"`
	Weather           string      `json:"weather"`
	WindDir           string      `json:"wind_dir"`
	WindSpdKmh        int         `json:"wind_spd_kmh"`
	WindSpdKt         int         `json:"wind_spd_kt"`
}

// Observations ...
type Observations struct {
	Notice []Notice `json:"notice"`
	Header []Header `json:"header"`
	Data   []Data   `json:"data"`
}
