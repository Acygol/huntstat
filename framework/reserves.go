package framework

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

var Reserves []string

//
// LoadReserves reads in the JSON array containing
// all theHunter reserve names into Reserves,
// returning an error indicating success or failure
//
func LoadReserves() error {
	body, err := ioutil.ReadFile(filepath.FromSlash("data/reserves.json"))
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &Reserves)
	return err
}
