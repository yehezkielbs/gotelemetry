package gotelemetry

type Batch map[string]*Flow

func (b Batch) Flow(name string) (*Flow, bool) {
	r, ok := b[name]

	return r, ok
}

func (b Batch) SetFlow(f *Flow) {
	b[f.Tag] = f
}

func (b Batch) DeleteFlow(name string) {
	delete(b, name)
}

func (b Batch) Publish(credentials Credentials) error {
	r, err := buildRequest(
		"POST",
		credentials,
		"/data",
		map[string]interface{}{
			"values": b,
		},
	)

	if err != nil {
		return err
	}

	return sendJSONRequest(r, nil)
}
