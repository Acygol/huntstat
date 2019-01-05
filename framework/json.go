package framework

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

func LoadFromJson(fileName string, data interface{}) error {
	body, err := ioutil.ReadFile(filepath.FromSlash(fileName))
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &data)
	return err
}
