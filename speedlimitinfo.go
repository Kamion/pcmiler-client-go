package pcmiler

type SpeedLimitInfo struct {
	Speed int `json:"Speed"`
	SpeedType string `json:"SpeedType"`
	LinkIds int `json:"LinkIds"`
	RoadClass string `json:"RoadClass"`
}