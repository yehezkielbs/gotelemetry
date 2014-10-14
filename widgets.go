package gotelemetry

type Widget struct {
	credentials Credentials `json:"-"`
	Id          string      `json:"id"`
	Variant     string      `json:"variant"`
	BoardId     string      `json:"board_id"`
	FlowIds     []string    `json:"flow_ids,omitempty"`
	Column      int         `json:"column"`
	Row         int         `json:"row"`
	Width       int         `json:"width"`
	Height      int         `json:"height"`
	BoardIndex  int         `json:"in_board_index"`
	Background  string      `json:"background"`
}

func NewWidget(credentials Credentials, board *Board, variant string, column, row, width, height, boardIndex int, background string) (*Widget, error) {
	w := &Widget{
		credentials: credentials,
		Variant:     variant,
		BoardId:     board.Id,
		Column:      column,
		Row:         row,
		Width:       width,
		Height:      height,
		BoardIndex:  boardIndex,
		Background:  background,
	}

	err := w.Save()

	if err != nil {
		return nil, err
	}

	return w, nil
}

func GetWidget(credentials Credentials, id string) (*Widget, error) {
	request, err := buildRequest("GET", credentials, "/widgets/"+id, nil)

	if err != nil {
		return nil, err
	}

	w := &Widget{}

	err = sendJSONRequestInterface(request, &w)

	if err != nil {
		return nil, err
	}

	w.credentials = credentials

	return w, nil
}

func (w *Widget) Delete() error {
	request, err := buildRequest("DELETE", w.credentials, "/widgets/"+w.Id, nil)

	if err != nil {
		return err
	}

	_, err = sendJSONRequest(request)

	return err
}

func (w *Widget) Save() error {
	request, err := buildRequest("POST", w.credentials, "/widgets", w)

	if err != nil {
		return err
	}

	return sendJSONRequestInterface(request, &w)
}
