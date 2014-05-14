package gotelemetry

type BarchartBar struct {
	Color string `json:"color,omitempty"`
	Label string `json:"label,omitempty"`
	Value string `json:"value,omitempty"`
}

type Barchart struct {
	ExpiresAt int           `json:"expires_at,omitempty"`
	Title     string        `json:"title,omitempty"`
	Bars      []BarchartBar `json:"bars"`
}
