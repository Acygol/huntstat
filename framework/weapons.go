package framework

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
)

//
// Enum-like declaration for weapon categories used in the Weapon struct
type WeaponCategory int

//
// Used to generate primary, secondary, and sidearms
type Weapon struct {
	Name     string
	Category WeaponCategory
}

var Weapons []Weapon

//
// LoadWeapons reads in the JSON array containing
// all theHunter weapon names into Weapons,
// returning an error indicating success or failure
//
func LoadWeapons() error {
	body, err := ioutil.ReadFile(filepath.FromSlash("data/weapons.json"))
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &Weapons)

	for i, weapon := range Weapons {
		log.Printf("i: %d\tname: %s\tcategory: %+v\n", i, weapon.Name, weapon.Category)
	}
	return err
}
