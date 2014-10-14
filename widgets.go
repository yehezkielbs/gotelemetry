package gotelemetry

type Widget struct {
	Variant    string `json:"variant"`
	BoardId    string `json:"board_id"`
	Column     int    `json:"column"`
	Row        int    `json:"row"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
	BoardIndex int    `json:"in_board_index"`
	Background string `json:"background"`
}

func NewWidget(credentials Credentials, board *Board, variant string, column, row, width, height, boardIndex int, background string) (*Widget, error) {
	w := &Widget{
		Variant:    variant,
		BoardId:    board.Id,
		Column:     column,
		Row:        row,
		Width:      width,
		Height:     height,
		BoardIndex: boardIndex,
		Background: background,
	}

	err := w.Save(credentials)

	if err != nil {
		return nil, err
	}

	return w, nil
}

func (w *Widget) Save(credentials Credentials) error {
	request, err := buildRequest("POST", credentials, "/widgets", w)

	if err != nil {
		return err
	}

	_, err = sendJSONRequest(request)

	return err
}
