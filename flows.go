package gotelemetry

type Flow struct {
	Tag     string
	Payload interface{}
}

func NewFlow(tag string, payload interface{}) *Flow {
	return &Flow{tag, payload}
}
