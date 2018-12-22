package framework

//
// Enum-like declaration for weapon categories used in the Weapon struct
type WeaponCategory int

//
// Used to generate primary, secondary, and sidearms
type Weapon struct {
	Name 		string
	Category 	WeaponCategory
}
