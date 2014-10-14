package gotelemetry

type Board struct {
	Id          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Theme       string `json:"theme,omitempty"`
	DisplayName bool   `json:"display_board_name,omitempty"`
	AspectRatio string `json:"aspect_ratio,omitempty"`
}

func NewBoard(credentials Credentials, name, theme string, displayName bool, aspectRatio string) (*Board, error) {
	result := &Board{
		Name:        name,
		Theme:       theme,
		DisplayName: displayName,
		AspectRatio: aspectRatio,
	}

	request, err := buildRequest("POST", credentials, "/boards", result)

	if err != nil {
		return nil, err
	}

	var responseBody interface{}
	responseBody, err = sendJSONRequest(request)

	if err != nil {
		return nil, err
	}

	result.Id = responseBody.(map[string]interface{})["id"].(string) //FIXME: Find a better way to copy all attrbs from the return to the b Board

	return result, nil
}

func (b *Board) DeleteBoard(credentials Credentials) error {
	request, err := buildRequest("DELETE", credentials, "/boards/"+b.Id, nil)

	if err != nil {
		return err
	}

	_, err = sendJSONRequest(request)
	return err
}
