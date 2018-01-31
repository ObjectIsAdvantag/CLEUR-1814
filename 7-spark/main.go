package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

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

func main() {
	sparkToken := ""
	response, err := Post("https://api.ciscospark.com/v1/rooms", sparkToken, map[string]interface{}{
		"title": "CLEUR Demo",
	})
	if err != nil {
		panic(err)
	}
	// The client must close the response body when finished with it:
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(fmt.Errorf("http error: " + err.Error()))
	}
	createSparkRoomResponse := map[string]interface{}{}
	err = json.Unmarshal(body, &createSparkRoomResponse)
	if err != nil {
		panic(fmt.Errorf("http error: " + err.Error()))
	}

	_, err = Post("https://api.ciscospark.com/v1/messages", sparkToken, map[string]interface{}{
		"roomId": createSparkRoomResponse["id"],
		"text":   "Hello from Barcelona!",
	})

	if err != nil {
		panic(fmt.Errorf("http error: " + err.Error()))
	}

}
