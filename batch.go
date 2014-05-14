package gotelemetry

type Batch map[string]Flow

func (b *Batch) Flow(name string) (*Flow, bool) {
	return b[name]
}

func (b *Batch) SetFlow(f *Flow) {
	b[f.Name] = f
}

func (b *Batch) DeleteFlow(name string) {
	delete(b, name)
}

func (b *Batch) Publish(credentials Credentials) error {
	r, err := buildRequest(
		"PUT",
		credentials,
		"/flows/"+f.Tag+"/data",
		map[string]interface{}{
			"values": b,
		},
	)

	if err != nil {
		return err
	}

	return sendJSONRequest(r, nil)
}
