package gotelemetry

import (
	"time"
)

type BatchStream struct {
	C              chan *Flow
	errorChannel   *chan error
	credentials    Credentials
	control        chan bool
	updates        []*Flow
	updateInterval time.Duration
}

func NewBatchStream(credentials Credentials, submissionInterval time.Duration, errorChannel *chan error) (*BatchStream, error) {
	if submissionInterval < time.Second {
		return nil, NewError(500, "Invalid submission interval (must be >= 1s)")
	}

	result := &BatchStream{
		C:              make(chan *Flow, 10),
		errorChannel:   errorChannel,
		credentials:    credentials,
		control:        make(chan bool, 0),
		updates:        []*Flow{},
		updateInterval: submissionInterval,
	}

	go result.handle()

	return result, nil
}

func (b *BatchStream) Send(f *Flow) {
	b.C <- f
}

func (b *BatchStream) Stop() {
	b.control <- true
}

func (b *BatchStream) handle() {
	t := time.After(b.updateInterval)

	for {
		select {
		case flow := <-b.C:

			if flow == nil {
				continue
			}

			b.updates = append(b.updates, flow)

		case <-b.control:

			b.sendUpdates()
			return

			return

		case <-t:

			b.sendUpdates()
			t = time.After(b.updateInterval)
		}

	}
}

func (b *BatchStream) sendUpdates() {
	if len(b.updates) == 0 {
		return
	}

	batch := Batch{}

	for _, flow := range b.updates {
		batch.SetFlow(flow)
	}

	err := batch.Publish(b.credentials)

	if err != nil && b.errorChannel != nil {
		*b.errorChannel <- err
	}
}
