package framework

import (
	"strings"
)

var Animals []string

//
// LoadAnimals reads in the JSON array containing
// all theHunter animal names into Animals,
// returning an error indicating success or failure
//
func LoadAnimals() error {
	err := LoadFromJson("data/animals.json", &Animals)
	return err
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
	for _, n := range Animals {
		if strings.Contains(strings.ToLower(n), strings.ToLower(name)) {
			return true
		}
	}
	return false
}
