package teams

import (
	"cleur-barcelona-2018/cli/ciscosparkapi/constants"
	"cleur-barcelona-2018/cli/ciscosparkapi/general"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type CreateResponse struct {
	Created string `json:"created"`
	ID      string `json:"id"`
	Name    string `json:"name"`
}

func Create(name, token string) (*CreateResponse, error) {
	// create the spark payload
	payload := map[string]interface{}{
		"name": name,
	}

	res, err := general.Post(constants.Teams, token, payload)
	if err != nil {
		return nil, fmt.Errorf("http error: " + err.Error())
	}

	// we use defer to close the http handles
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("http error: " + err.Error())
	}

	createResponse := &CreateResponse{}
	err = json.Unmarshal(body, createResponse)
	if err != nil {
		return nil, fmt.Errorf("http error: " + err.Error())
	}
	return createResponse, nil
}

func Delete(id, token string) error {
	_, err := general.Delete(constants.Teams, token, id)
	return err
}
