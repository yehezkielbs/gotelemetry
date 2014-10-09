package gotelemetry

type Board struct {
	Name               string `json:"name,omitempty"`
	Theme              string `json:"theme,omitempty"`
	Display_board_name bool   `json:"display_board_name,omitempty"`
	Aspect_ratio       string `json:"aspect_ratio,omitempty"`
}

func (b *Board) CreateBoard(credentials Credentials) error {
	request, err := buildRequest("POST", credentials, "/boards", b)

	if err != nil {
		return err
	}
	return sendJSONRequest(request, nil)
}
