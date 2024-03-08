package data

type CaptureLocationResponseModel struct {
	Status  int                       `json:"status"`
	Message string                    `json:"message"`
	Data    *CaptureLocationDataModel `json:"Data"`
}

type CaptureLocationDataModel struct {
	Location    string `json:"location"`
	Address     string `json:"address"`
	Information string `json:"information"`
}
