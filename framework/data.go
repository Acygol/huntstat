package framework

import (
	"log"
	"math/rand"
	"strings"
	"time"
)

//
// Weapon holds information regarding a	weapon:
//		name: 	the full name of the weapon
//				as seen in the store
//
//		type: 	either "primary" or "sidearm"
//				used to generate a random loadout
//				in cmd/weaponscommand.go
//
type Weapon struct {
	Name string
	Type string
}

var (
	//
	// Animals holds all existing animals
	// currently in theHunter
	//
	Animals []string

	//
	// Weapons holds all existing weapons
	// currently in theHunter
	//
	Weapons []Weapon

	//
	// Reserves holds all existing reserves
	// currently in theHunter
	//
	Reserves []string
)

//
// LoadGameData reads from JSON files containing
// all theHunter data such as animals, guns, and
// reserves
//
func LoadGameData() {
	err := LoadFromJSON("data/weapons.json", &Weapons)
	if err != nil {
		log.Fatal("error loading weapons.json,", err)
		return
	}
	if err = LoadFromJSON("data/animals.json", &Animals); err != nil {
		log.Fatal("error loading animals.json,", err)
	}
	if err = LoadFromJSON("data/reserves.json", &Reserves); err != nil {
		log.Fatal("error loading reserves.json,", err)
	}
}

//
// GenerateRandomWeaponOnce takes as input a slice
// which should be a copy of framework.Weapons and
// generates a random weapon, deletes the index
// generated, and returns the name of weapon
//
func GenerateRandomWeaponOnce(weapons []Weapon, weaptype string) string {
	rand.Seed(time.Now().UnixNano())

	index := 0
	for index = rand.Intn(len(weapons)); !strings.EqualFold(weapons[index].Type, weaptype); index = rand.Intn(len(weapons)) {
		// empty
	}

	// delete the element at [index] from the input slice
	// but only after it has returned the name
	defer func() {
		weapons = append(weapons[:index], weapons[index+1:]...)
	}()

	return weapons[index].Name
}

//
// GenerateRandomWeapon returns the index of a random
// weapon within the Weapons slice if it matches the
// requested weapon type ("primary" or "secondary")
//
/*
func GenerateRandomWeapon(weaptype string) int {
	rand.Seed(time.Now().UnixNano())

	index := 0
	for index = rand.Intn(len(Weapons)); Weapons[index].Type != weaptype; index = rand.Intn(len(Weapons)) {
		// empty
	}
	return index
}
*/

//
// IsValidAnimalName takes as input an animal name
// and checks if it is valid by sequentially
// searching through Animals. It is sufficient to
// check if the input name is a substring of an
// actual animal name so that users don't have to
// be extremely literal when using animal names
// given that some are quite lengthy:
//		e.g., "White-tailed Ptarmigan"
//
func IsValidAnimalName(name string) bool {
	loweredName := strings.ToLower(name)
	for _, n := range Animals {
		if strings.Contains(strings.ToLower(n), loweredName) {
			return true
		}
	}
	return false
}
