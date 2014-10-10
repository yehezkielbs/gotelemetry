package gotelemetry

type Board struct {
	Id                 string `json:"id,omitempty"`
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

	var responseBody interface{}
	responseBody, err = sendJSONRequest(request)

	b.Id = responseBody.(map[string]interface{})["id"].(string) //FIXME: Find a better way to copy all attrbs from the return to the b Board

	return err
}

func (b *Board) DeleteBoard(credentials Credentials) error {
	request, err := buildRequest("DELETE", credentials, "/boards/"+b.Id, nil)

	if err != nil {
		return err
	}

	_, err = sendJSONRequest(request)
	return err
}
