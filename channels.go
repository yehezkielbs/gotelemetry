package gotelemetry

import (
	"encoding/json"
)

type ExportedWidget struct {
	Variant    string `json:"variant"`
	Tag        string `json:"tag"`
	Column     int    `json:"column"`
	Row        int    `json:"row"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
	BoardIndex int    `json:"in_board_index"`
	Background string `json:"background"`
}

type ExportedBoard struct {
	Name             string            `json:"name,omitempty"`
	Theme            string            `json:"theme,omitempty"`
	DisplayName      bool              `json:"display_board_name,omitempty"`
	AspectRatio      string            `json:"aspect_ratio,omitempty"`
	FontFamily       string            `json:"font_family",omitempty`
	FontSize         string            `json:"font_size",omitempty`
	WidgetBackground string            `json:"widget_background",omitempty`
	WidgetMargins    int64             `json:"widget_margins",omitempty`
	WidgetPadding    int64             `json:"widget_padding",omitempty`
	Widgets          []*ExportedWidget `json:"widgets"`
}

func ImportBoard(credentials Credentials, name string, prefix string, board *ExportedBoard) (*Board, error) {
	// First, make sure the board doesn't already exist

	result, err := GetBoardByName(credentials, name)

	if err == nil && result != nil {
		//TODO: Make sure that the board retrieved from the API matches
		// the makeup of the board template we're trying to import.

		return result, nil
	}

	result = &Board{credentials: credentials}

	encoded, err := json.Marshal(board)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(encoded, result)

	if err != nil {
		return nil, err
	}

	result.Widgets = []*Widget{}

	result.Name = name

	err = result.Save()

	if err != nil {
		return nil, err
	}

	for _, exportedWidget := range board.Widgets {
		flow, err := NewFlowWithLayout(
			credentials,
			prefix+exportedWidget.Tag,
			exportedWidget.Variant,
			"",
			"",
			"",
		)

		if err != nil {
			result.Delete()
			return nil, err
		}

		widget := &Widget{credentials: credentials}

		encoded, err = json.Marshal(exportedWidget)

		if err != nil {
			flow.Delete()
			result.Delete()
			return nil, err
		}

		err = json.Unmarshal(encoded, &widget)

		if err != nil {
			flow.Delete()
			result.Delete()
			return nil, err
		}

		widget.BoardId = result.Id
		err = widget.Save()

		if err != nil {
			flow.Delete()
			result.Delete()
			return nil, err
		}

		result.Widgets = append(result.Widgets, widget)
	}

	return result, nil
}

func (b *Board) Export() (*ExportedBoard, error) {
	result := &ExportedBoard{}

	encoded, err := json.Marshal(b)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(encoded, &result)

	if err != nil {
		return nil, err
	}

	for index, widget := range b.Widgets {
		flow, err := GetFlowLayout(b.credentials, widget.FlowIds[0])

		if err != nil {
			return nil, err
		}

		result.Widgets[index].Tag = flow.Tag
	}

	return result, nil
}