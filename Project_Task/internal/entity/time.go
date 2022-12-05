package entity

type CoordinatesArr struct {
	Results []Coordinates `json:"results"`
}

type Coordinates struct {
	Longitude float32 `json:"longitude"`
	Latitude  float32 `json:"latitude"`
	Timezone  string  `json:"timezone"`
}

type Time struct {
	Time     string `json:"time"`
	TimeZone string `json:"timeZone"`
}
