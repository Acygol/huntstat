package framework

import (
	"errors"
	"log"
	"math/rand"
	"strings"
	"time"
)

//
// Ammo holds basic information regarding an ammotype
// in TheHunter
//
type Ammo struct {
	Name   string
	Weight float64
}

//
// Animal holds basic information regarding an animal
// in TheHunter
//
type Animal struct {
	Name string
	Ammo []Ammo
}

//
// Weapon holds all information regarding a weapon
// in TheHunter
//
type Weapon struct {
	Name     string
	Type     string
	Editions []string
	Ammo     []Ammo
	Weight   float64
}

//
// Reserve holds all information regarding a reserve
// in TheHunter
//
type Reserve struct {
	Name    string
	Animals []Animal
}

const (
	Primary = "primary"
	Sidearm = "sidearm"
)

var (
	//
	// Ammunitions holds all ammo types currently
	// in TheHunter
	//
	Ammunitions []Ammo

	//
	// Animals holds all existing animals currently
	// in TheHunter
	//
	Animals []Animal

	//
	// Weapons holds all existing weapons currently
	// in TheHunter
	//
	Weapons []Weapon

	//
	// Reserves holds all existing reserves currently
	// in TheHunter
	//
	Reserves []Reserve
)

//
// LoadGameData reads from JSON files containing
// all theHunter data such as animals, guns, and
// reserves
//
func LoadGameData() {
	err := LoadFromJSON("data/ammo.json", &Ammunitions)
	if err != nil {
		log.Fatal("error loading ammo.json", err)
		return
	}
	loadWeapons()
	loadAnimals()
	loadReserves()
}

func loadWeapons() {
	//
	// Ammo types are stored as a string in data/ammo.json
	// and not as their object representation. For this reason
	// there is a need to define a temporary alternative
	// struct for the Weapon struct which can hold ammo types
	// by their names
	//
	type tmpWeap struct {
		Name     string
		Editions []string
		Type     string
		Ammo     []string
		Weight   float64
	}
	var weap []tmpWeap
	if err := LoadFromJSON("data/weapons.json", &weap); err != nil {
		log.Fatal("error loading weapons.json,", err)
	}

	//
	// Ammo types loaded from the JSON must be converted
	// to their datatype representation
	//
	var ammotypes []Ammo
	for _, tmpw := range weap {
		for _, ammoTypeName := range tmpw.Ammo {
			for _, ammoType := range Ammunitions {
				if strings.Compare(ammoType.Name, ammoTypeName) == 0 {
					ammotypes = append(ammotypes, ammoType)
				}
			}
		}
		//
		// Construct, and append to the global Weapons slice,
		// an instance of the Weapon type with its ammo
		// field value as the correct data type
		//
		Weapons = append(Weapons, Weapon{Name: tmpw.Name, Type: tmpw.Type, Editions: tmpw.Editions, Ammo: ammotypes, Weight: tmpw.Weight})
	}
}

func loadAnimals() {
	//
	// Ammo types are stored as a string in data/ammo.json
	// and not as their object representation. For this reason
	// there is a need to define a temporary alternative
	// struct for the Animal struct which can hold ammo types
	// by their names
	//
	type tmpAnimal struct {
		Name string
		Ammo []string
	}
	var tmpanimals []tmpAnimal
	if err := LoadFromJSON("data/animals.json", &tmpanimals); err != nil {
		log.Fatal("error loading animals.json,", err)
	}

	//
	// Ammo types loaded from the JSON must be converted
	// to their datatype representation
	//
	var ammotypes []Ammo
	for _, tmpa := range tmpanimals {
		for _, ammoTypeName := range tmpa.Ammo {
			for _, ammoType := range Ammunitions {
				if strings.Compare(ammoType.Name, ammoTypeName) == 0 {
					ammotypes = append(ammotypes, ammoType)
				}
			}
		}
		//
		// Construct, and append to the global Animals slice,
		// an instance of the Animal type with its ammo
		// field value as the correct data type
		//
		Animals = append(Animals, Animal{Name: tmpa.Name, Ammo: ammotypes})
	}
}

func loadReserves() {
	//
	// Animals are stored as a string in data/reserves.json
	// and not as their object representation. For this reason
	// there is a need to define a temporary alternative
	// struct for the Reserve struct which can hold animals by
	// their names
	//
	type tmpReserve struct {
		Name    string
		Animals []string
	}
	var tmpReserves []tmpReserve

	if err := LoadFromJSON("data/reserves.json", &tmpReserves); err != nil {
		log.Fatal("error loading reserves.json,", err)
	}

	//
	// Animals loaded from the JSON must be converted
	// to their datatype representation
	//
	var tmpAnimals []Animal
	for _, tmpr := range tmpReserves {
		//
		// Go over all animal names in tmpr.Animals
		//
		for _, animalName := range tmpr.Animals {
			//
			// Go over all animal objects in the global Animals
			// slice and check if its name equals animalName
			//
			for _, animal := range Animals {
				if strings.EqualFold(animal.Name, animalName) {
					tmpAnimals = append(tmpAnimals, animal)
				}
			}
		}
		//
		// Construct, and append to the global Reserves slice,
		// an instance of the Reserve type with its animals
		// field value as the correct data type
		//
		Reserves = append(Reserves, Reserve{Name: tmpr.Name, Animals: tmpAnimals})
	}
}

//
// GenerateRandomWeaponOnce takes as input a slice
// which should be a copy of framework.Weapons and
// generates a random weapon, deletes the index
// generated, and returns the name of weapon
//
func GenerateRandomWeaponOnce(weapons []Weapon, weaptype string, reserve Reserve, inventoryCap float64) (Weapon, error) {
	rand.Seed(time.Now().UnixNano())
	index := -1

	//
	// A primary weapon cannot be generated when the inventoryCap
	// is less than 4 because the minimum weight of a primary weapon
	// is 4. Equally, the minimum weight of a sidearm is 1.
	// Otherwise the for loop in the else statement will go on forever.
	//
	if strings.Compare(weaptype, Primary) == 0 && inventoryCap < 4 {
		return Weapon{}, errors.New("cannot generate primary weapon, inventoryCap too low")
	} else if strings.Compare(weaptype, Sidearm) == 0 && inventoryCap < 1 {
		return Weapon{}, errors.New("cannot generate sidearm weapon, inventoryCap too low")
	} else {
		for {

			//
			// Get a random weapon index
			//
			index = rand.Intn(len(weapons))

			//
			// If it doesn't match the requested weapontype, continue
			//
			if !strings.EqualFold(weapons[index].Type, weaptype) {
				continue
			}

			//
			// If it doesn't fit the reported inventory capacity, continue
			//
			if (inventoryCap - weapons[index].Weight) < 0 {
				continue
			}

			//
			// Go over all animals in the provided reserve and check
			// if any of the ammo of the current random weapon is
			// permitted to shoot an animal on the reserve. If these
			// checks all evaluate to true, then the weapon is a good
			// pick
			//
			for _, animal := range reserve.Animals {
				for _, weapAmmo := range weapons[index].Ammo {
					if weapAmmo.IsPermittedAmmo(animal) {
						// delete the element at [index] from the input slice
						// but only after it has returned the weapon
						defer func() {
							if index != -1 {
								weapons = append(weapons[:index], weapons[index+1:]...)
							}
						}()
						return weapons[index], nil
					}
				}
			}
		}
	}
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
// IsPermittedAmmo checks if the given ammo type can be
// used to hunt on a particular animal
//
func (ammoType *Ammo) IsPermittedAmmo(animal Animal) bool {
	for _, ammo := range animal.Ammo {
		if strings.Compare(ammo.Name, ammoType.Name) == 0 {
			return true
		}
	}
	return false
}

//
// GetReserveFromName returns a Reserve when the name
// matches that of an existing reserve. If not, it returns an
// error.
//
func GetReserveFromName(reservename string) (Reserve, error) {
	reservename = strings.ToLower(reservename)
	for _, reserve := range Reserves {
		resName := strings.ToLower(reserve.Name)
		if strings.Contains(resName, reservename) {
			return reserve, nil
		}
	}
	return Reserve{}, errors.New("reserve not found")
}

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
		if strings.Contains(strings.ToLower(n.Name), loweredName) {
			return true
		}
	}
	return false
}
