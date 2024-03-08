package data

type CaptureLocationRequestModel struct {
	Picture     string `json:"picture"`
	Location    string `json:"location"`
	Address     string `json:"address"`
	Information string `json:"information"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
}
