package gotelemetry

type Widget struct {
	Variant        string
	Board_id       string
	Column         int
	Row            int
	Width          int
	Height         int
	In_board_index int
	Background     string
}

func CreateWidget(widget Widget, credentials Credentials) error {
	request, err := buildRequest("POST", credentials, "/widgets", widget)

	if err != nil {
		return err
	}

	return sendJSONRequest(request, nil)
}
