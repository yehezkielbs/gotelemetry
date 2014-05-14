package gotelemetry

type Flow struct {
	Tag  string
	Data interface{}
}

func NewFlow(tag string, data interface{}) *Flow {
	return &Flow{tag, data}
}

func (f *Flow) Publish(credentials Credentials) error {
	r, err := buildRequest(
		"PUT",
		credentials,
		"/flows/"+f.Tag+"/data",
		f.Data,
	)

	if err != nil {
		return err
	}

	return sendJSONRequest(r, nil)
}
