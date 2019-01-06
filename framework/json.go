package framework

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

//
// LoadFromJSON reads from a JSON file at pathtofile and
// unmarshals it into 'data'
//
func LoadFromJSON(pathtofile string, data interface{}) error {
	body, err := ioutil.ReadFile(filepath.FromSlash(pathtofile))
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &data)
	return err
}
