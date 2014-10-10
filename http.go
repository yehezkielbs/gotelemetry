package gotelemetry

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"path"
)

func buildRequest(method string, credentials Credentials, fragment string, body interface{}, parameters ...map[string]string) (*http.Request, error) {
	url := *credentials.ServerURL

	url.Path = path.Join(url.Path, fragment)

	if len(parameters) > 0 {
		for index, value := range parameters[0] {
			url.Query().Add(index, value)
		}
	}

	var b []byte
	var err error

	if body == nil {
		b = []byte{}
	} else {
		b, err = json.Marshal(body)

		if err != nil {
			return nil, err
		}
	}

	r, err := http.NewRequest(method, url.String(), bytes.NewReader(b))

	if err != nil {
		return nil, err
	}

	r.Header.Set("content-type", "application/json")
	r.SetBasicAuth(credentials.APIKey, "")

	return r, nil
}

func readJSONResponseBody(r *http.Response, target interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(target); err != nil && err != io.EOF {
		return NewError(400, "Invalid JSON response.")
	}

	return nil
}

func sendJSONRequest(request *http.Request) (interface{}, error) {
	r, err := http.DefaultClient.Do(request)

	if err != nil {
		return nil, err
	}

	if r.Body != nil {
		defer r.Body.Close()
	}

	var body interface{}

	err = readJSONResponseBody(r, &body)

	if err != nil {
		return nil, err
	}

	return body, NewErrorWithData(r.StatusCode, r.Status, body)
}
