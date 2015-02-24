package gotelemetry

type Notification struct {
	Title    string `json:"title,omitempty"`
	Message  string `json:"message,omitempty"`
	Icon     string `json:"icon,omitempty"`
	Duration int    `json:"duration,omitempty"`
	SoundURL string `json:"sound_url,omitempty"`
}

type Channel struct {
	Tag string
}

func NewChannel(tag string) *Channel {
	return &Channel{Tag: tag}
}

func (c *Channel) SendNotification(credentials Credentials, notification Notification) error {
	if credentials.DebugChannel != nil {
		*credentials.DebugChannel <- NewDebugError("Sending notification %#v to channel %s", notification, c.Tag)
	}

	req, err := buildRequest(
		"POST",
		credentials,
		"channels/"+c.Tag+"/notifications",
		notification,
	)

	if err != nil {
		return err
	}

	_, err = sendJSONRequest(req)

	return err
}
