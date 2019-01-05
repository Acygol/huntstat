package framework

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
	err := LoadFromJson("data/weapons.json", &Weapons)
	return err
}
