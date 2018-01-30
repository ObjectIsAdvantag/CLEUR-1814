package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

type Result interface {
	GetStatus() int
	JSON() []byte
}

type Success struct {
	Status int         `json:"status"`
	Result interface{} `json:"result"`
}

type Err struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (s *Success) GetStatus() int {
	return s.Status
}

func (s *Success) JSON() []byte {
	return marshal(s)
}

func (e *Err) GetStatus() int {
	return e.Status
}

func (e *Err) JSON() []byte {
	return e.marshal()
}

func (e *Err) marshal() []byte {
	bytes, err := json.Marshal(e)
	if err != nil {
		e.Status = 500
		return []byte(err.Error())
	}
	return bytes
}

func Marshal(v interface{}) []byte {
	return marshal(v)
}

func marshal(v interface{}) []byte {
	bytes, err := json.Marshal(v)
	if err != nil {
		return []byte(err.Error())
	}
	return bytes
}

func BadParams(e error) Result {
	return &Err{400, e.Error()}
}

func ServerError() Result {
	return &Err{500, "server error"}
}

func UnAuthorized() Result {
	return &Err{401, "unauthorized"}

}

func NotFound() Result {
	return &Err{404, "not found"}
}

func NoContent() Result {
	return &Success{204, "no content"}
}

func Error(e error) Result {
	if e == nil {
		return &Err{500, "server error"}
	}
	if strings.Contains(e.Error(), "400: ") {
		return &Err{400, strings.Replace(e.Error(), "400: ", "", 1)}
	}
	if strings.Contains(e.Error(), "401: ") {
		return &Err{401, strings.Replace(e.Error(), "401: ", "", 1)}
	}
	return &Err{500, e.Error()}
}

func CustomError(s int, msg string) Result {
	return &Err{s, msg}
}

func Ok(v interface{}) Result {
	return &Success{Status: 200, Result: v}
}

func LoadYAML(relativePath string) error {

	path, err := filepath.Abs(relativePath)
	if err != nil {
		panic(err)
		return err
	}

	fmt.Println("loading yaml: " + path)

	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
		return err
	}

	var envMap map[string]string

	err = yaml.Unmarshal(data, &envMap)
	if err != nil {
		panic(err)
		return err
	}

	for k, v := range envMap {
		err = os.Setenv(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}
