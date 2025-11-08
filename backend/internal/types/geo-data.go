package types

type GeoData struct {
	City      string  `json:"city"`
	Region    string  `json:"region"`
	Country   string  `json:"country_name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
