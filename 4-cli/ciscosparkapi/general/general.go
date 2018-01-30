package general

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type CreateResponse struct {
	Created string `json:"created"`
	ID      string `json:"id"`
	Name    string `json:"name"`
}

func Post(uri, token string, payload map[string]interface{}) (*http.Response, error) {

	data, _ := json.Marshal(payload)
	buffer := bytes.NewBuffer(data)

	req, err := http.NewRequest(http.MethodPost, uri, buffer)
	if err != nil {
		return nil, fmt.Errorf("http error: " + err.Error())
	}

	headers := http.Header{}
	headers.Add("Content-Type", "application/json")
	headers.Add("Accept", "application/json")
	headers.Add("Authorization", "Bearer "+token)
	req.Header = headers

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http error: " + err.Error())
	}
	if res.StatusCode == 401 {
		return nil, fmt.Errorf("401: %v", errors.New("invalid authorization"))
	}
	return res, nil
}

func Delete(uri, token, id string) (*http.Response, error) {

	uri = uri + "/" + id

	req, err := http.NewRequest(http.MethodDelete, uri, nil)
	if err != nil {
		return nil, fmt.Errorf("http error: " + err.Error())
	}

	headers := http.Header{}
	headers.Add("Authorization", "Bearer "+token)
	req.Header = headers

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http error: " + err.Error())
	}
	if res.StatusCode == 401 {
		return nil, fmt.Errorf("401: %v", errors.New("invalid authorization"))
	}

	if res.StatusCode == 404 {
		return nil, fmt.Errorf("404: %v", errors.New("not found"))
	}
	return res, nil

}
