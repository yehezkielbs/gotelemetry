package gotelemetry

type Board struct {
	credentials      Credentials `json:"-"`
	Id               string      `json:"id,omitempty"`
	Name             string      `json:"name,omitempty"`
	Theme            string      `json:"theme,omitempty"`
	DisplayName      bool        `json:"display_board_name,omitempty"`
	AspectRatio      string      `json:"aspect_ratio,omitempty"`
	FontFamily       string      `json:"font_family",omitempty`
	FontSize         string      `json:"font_size",omitempty`
	WidgetBackground string      `json:"widget_background",omitempty`
	WidgetMargins    int64       `json:"widget_margins",omitempty`
	WidgetPadding    int64       `json:"widget_padding",omitempty`
	Widgets          []*Widget   `json:"widgets,omitempty"`
	ChannelIds       []string    `json:"channel_ids",omitempty"`
}

func NewBoard(credentials Credentials, name, theme string, displayName bool, aspectRatio string) (*Board, error) {
	result := &Board{
		credentials: credentials,
		Name:        name,
		Theme:       theme,
		DisplayName: displayName,
		AspectRatio: aspectRatio,
	}

	result.Save()

	return result, nil
}

func GetBoard(credentials Credentials, id string) (*Board, error) {
	request, err := buildRequest("GET", credentials, "/boards/"+id, nil)

	if err != nil {
		return nil, err
	}

	b := &Board{}

	err = sendJSONRequestInterface(request, b)

	if err != nil {
		return nil, err
	}

	b.credentials = credentials

	return b, err
}

func (b *Board) Save() error {
	request, err := buildRequest("POST", b.credentials, "/boards", b)

	if err != nil {
		return err
	}

	err = sendJSONRequestInterface(request, b)

	return err
}

func (b *Board) Delete() error {
	request, err := buildRequest("DELETE", b.credentials, "/boards/"+b.Id, nil)

	if err != nil {
		return err
	}

	_, err = sendJSONRequest(request)
	return err
}
