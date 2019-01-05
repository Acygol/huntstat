package cmd

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/acygol/huntstat/framework"
)

const (
	PRIMARY framework.WeaponCategory = 0
	SIDEARM framework.WeaponCategory = 1
)

func WeaponsCommand(ctx framework.Context) {
	weapons := framework.Weapons
	var reply strings.Builder

	index := generateRandomFromSlice(weapons, PRIMARY)
	fmt.Fprintf(&reply, "Primary: %s\n", weapons[index].Name)
	removeIndex(weapons, index)

	// To make sure the first weapon isn't chosen again, I remove it from the slice
	index = generateRandomFromSlice(weapons, PRIMARY)
	fmt.Fprintf(&reply, "Secondary: %s\n", weapons[index].Name)
	removeIndex(weapons, index)

	index = generateRandomFromSlice(weapons, SIDEARM)
	fmt.Fprintf(&reply, "Sidearm (optional): %s\n", weapons[index].Name)
	removeIndex(weapons, index)

	ctx.Reply(reply.String())
}

//
// Takes as input an array of Weapons and returns from it a random index of the requested
// category
func generateRandomFromSlice(input []framework.Weapon, weaptype framework.WeaponCategory) int {
	rand.Seed(time.Now().UnixNano())

	//
	// the for header does all that it is intended to be done. It keeps
	// looping while the randomly generated element is not of the right type
	//
	index := 0
	for index = rand.Intn(len(input)); input[index].Category != weaptype; index = rand.Intn(len(input)) {
		// empty
	}
	return index
}

func removeIndex(input []framework.Weapon, index int) {
	input = append(input[:index], input[index+1:]...)
}
