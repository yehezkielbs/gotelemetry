package gotelemetry

import (
	"net/http"
)

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

	result.Save(credentials)

	return result, nil
}

func GetBoard(credentials Credentials, id string) (*Board, error) {
	request, err := buildRequest("GET", credentials, "/boards/"+id, nil)

	if err != nil {
		return nil, err
	}

	b := &Board{}

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		return nil, err
	}

	err = readJSONResponseBody(res, b)

	return b, err
}

func (b *Board) Save(credentials Credentials) error {
	request, err := buildRequest("POST", credentials, "/boards", b)

	if err != nil {
		return err
	}

	var responseBody interface{}
	responseBody, err = sendJSONRequest(request)

	if err != nil {
		return err
	}

	b.Id = responseBody.(map[string]interface{})["id"].(string) //FIXME: Find a better way to copy all attrbs from the return to the b Board

	return nil
}

func (b *Board) DeleteBoard(credentials Credentials) error {
	request, err := buildRequest("DELETE", credentials, "/boards/"+b.Id, nil)

	if err != nil {
		return err
	}

	_, err = sendJSONRequest(request)
	return err
}
