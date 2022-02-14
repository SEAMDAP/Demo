package utils

// Struct for reading custom SenML messages
// This is a custom SenML version for SEAMDAP


type SenMLPos struct {
	TimeRecord string                 `json:"t"`
	Name       string                 `json:"n"`
	Data       map[string]interface{} `json:"v"`
}

type Custom struct {
	Record []SenMLPos `json:"senml"`
}


