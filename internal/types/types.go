package types

// JSON metadata for the postUpload handler
type Metadata struct {
	Dir    string         `json:"dir"`
	Size   int            `json:"size"`
	Scales map[string]int `json:"scales"`
}
