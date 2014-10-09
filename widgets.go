package gotelemetry

type Widget struct {
	Variant        string `json:"variant"`
	Board_id       string `json:"board_id"`
	Column         int    `json:"column"`
	Row            int    `json:"row"`
	Width          int    `json:"width"`
	Height         int    `json:"height"`
	In_board_index int    `json:"in_board_index"`
	Background     string `json:"background"`
}

func (w *Widget) CreateWidget(credentials Credentials) error {
	request, err := buildRequest("POST", credentials, "/widgets", w)

	if err != nil {
		return err
	}

	_, err = sendJSONRequest(request)

	return err
}
