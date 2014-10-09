package gotelemetry

// Type Batch describes a collection of flows that can be submitted simultaneously to the Telemetry servers.
//
// Note the underlying data structure of the batch is a map, and therefore batches are not thread safe
// by default. If you require thread safety, you must mediate access to the batch through some kind
// of synchronization mechansism, like a mutex.
type Batch map[string]*Flow

// Flow() retrieves a flow from the batch, given its tag.
func (b Batch) Flow(tag string) (*Flow, bool) {
	r, ok := b[tag]

	return r, ok
}

// SetFlow() adds or overwrites a flow to the batch
func (b Batch) SetFlow(f *Flow) {
	b[f.Tag] = f
}

// DeleteFlow() deletes a flow from the batch
func (b Batch) DeleteFlow(tag string) {
	delete(b, tag)
}

// Publish() submits a batch to the Telemetry API servers, and returns either an instance
// of gotelemetry.Error if a REST error occurs, or errors.Error if any other error occurs.
func (b Batch) Publish(credentials Credentials) error {
	data := map[string]interface{}{}

	for key, flow := range b {
		data[key] = flow.Data
	}

	r, err := buildRequest(
		"POST",
		credentials,
		"/data",
		map[string]interface{}{
			"data": data,
		},
	)

	if err != nil {
		return err
	}

	_, err = sendJSONRequest(r)

	return err
}
