package gotelemetry

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

func buildRequest(method string, credentials Credentials, fragment string, body interface{}, parameters ...map[string]string) (*http.Request, error) {
	URL := *credentials.ServerURL

	URL.Path = path.Join(URL.Path, fragment)

	if len(parameters) > 0 {
		p := url.Values{}

		for index, value := range parameters[0] {
			p.Add(index, value)
		}

		URL.RawQuery = p.Encode()
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

	r, err := http.NewRequest(method, URL.String(), bytes.NewReader(b))

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

func sendJSONRequestInterface(request *http.Request, target interface{}) error {
	r, err := http.DefaultClient.Do(request)

	if err != nil {
		return err
	}

	if r.Body != nil {
		defer r.Body.Close()
	}

	if r.StatusCode > 399 {
		v, _ := ioutil.ReadAll(r.Body)

		return NewErrorWithData(r.StatusCode, r.Status, v)
	}

	return readJSONResponseBody(r, target)
}

func sendJSONRequest(request *http.Request) (interface{}, error) {
	var body interface{}

	err := sendJSONRequestInterface(request, &body)

	return body, err
}
