package framework

var Reserves []string

//
// LoadReserves reads in the JSON array containing
// all theHunter reserve names into Reserves,
// returning an error indicating success or failure
//
func LoadReserves() error {
	err := LoadFromJson("data/reserves.json", &Reserves)
	return err
}
