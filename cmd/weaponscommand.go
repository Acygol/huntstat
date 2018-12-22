package cmd

import (
	"github.com/acygol/huntstat/framework"
	"math/rand"
	"time"
)

const (
	PRIMARY 	framework.WeaponCategory = 0
	SIDEARM 	framework.WeaponCategory = 1
)

func WeaponsCommand(ctx framework.Context) {
	weapons := []framework.Weapon {
		{ ".454 Revolver", SIDEARM }, { "Cable-backed Bow", PRIMARY }, { "16 GA Side By Side Shotgun", PRIMARY }, { "303 British Bolt Action Rifle", PRIMARY },
		{ "30-30 Lever Action Rifle", PRIMARY }, { "6.5x55 Blaser R8 Bolt Action Rifle", PRIMARY }, { "270 Bolt Action Rifle", PRIMARY }, { "223 Bolt Action Rifle", PRIMARY },
		{ "8x57 IS Anschütz 1780 D FL Bolt Action", PRIMARY }, { "Longbow", PRIMARY }, { "Rifle10 GA Lever Action Shotgun", PRIMARY }, { ".30-06 Bolt Action Rifle", PRIMARY },
		{ ".357 Revolver", SIDEARM }, { "50 Inline Muzzleloading Pistol", SIDEARM }, { "Heavy Recurve Bow", PRIMARY }, { "12GA Single Shotgun", PRIMARY },
		{ ".17 HMR Lever Action Rifle", PRIMARY }, { ".405 Lever Action Rifle", PRIMARY }, { "12 GA Blaser F3", PRIMARY }, { "10mm Semi-Automatic Pistol", SIDEARM },
		{ "Compound Bow 'Parker Python'", PRIMARY }, { "Reverse Draw Crossbow", PRIMARY }, { "20 GA Semi-Automatic Shotgun", PRIMARY }, { "Recurve Bow", PRIMARY },
		{ ".50 Cap Lock Muzzleloader", PRIMARY }, { "8x57 IS K98k Bolt Action Rifle", PRIMARY }, { ".30 R O/U Break Action Rifle", PRIMARY }, { ".30-06 Lever Action Rifle", PRIMARY },
		{ "7mm Magnum Bullpup Rifle", PRIMARY }, { ".243 Bolt Action Rifle", PRIMARY }, { "Compound Bow “Pulsar”", PRIMARY }, { ".22 Air Rifle", PRIMARY },
		{ "12 GA Side by Side Shotgun", PRIMARY }, { ".44 Revolver", SIDEARM }, { ".22 Plinkington", PRIMARY }, { ".300 Bolt Action Rifle", PRIMARY },
		{ ".308 Bolt Action Rifle / Anschütz", PRIMARY }, { ".340 Weatherby Magnum Bolt Action Rifle", PRIMARY }, { "6.5x55 Bolt Action Rifle Panther", PRIMARY }, { "7x64mm Bolt Action Rifle", PRIMARY },
		{ "9.3x62 Anschütz 1780 D FL Bolt Action Rifle", PRIMARY }, { ".45-70 Government", PRIMARY }, { "7mm Magnum Break Action Rifle", PRIMARY }, { "9.3x74R O/U Break Action Rifle", PRIMARY },
		{ ".30-06 Stutzen Bolt Action Rifle", PRIMARY }, { ".223 Semi-Automatic Rifle", PRIMARY }, { "7.62x54R Bolt Action Rifle", PRIMARY }, { ".45-70 Buffalo Rifle", PRIMARY },
		{ "12 GA Pump Action Shotgun", PRIMARY }, { "16GA/9.3x74R Drilling", PRIMARY }, { ".22 Pistol Grasshopper", SIDEARM }, { ".45 Long Colt Revolver", SIDEARM },
		{ ".308 Handgun", SIDEARM }, { ".50 Inline Muzzleloader", PRIMARY }, { "Compound Bow 'Snakebite'", PRIMARY }, { "Crossbow tenpoint", PRIMARY },
	}
	index := generateRandomFromSlice(weapons, PRIMARY)
	ctx.Reply("Primary: " + weapons[index].Name)
	removeIndex(weapons, index)

	// To make sure the first weapon isn't chosen again, I remove it from the slice
	index = generateRandomFromSlice(weapons, PRIMARY)
	ctx.Reply("Secondary: " + weapons[index].Name)
	removeIndex(weapons, index)

	index = generateRandomFromSlice(weapons, SIDEARM)
	ctx.Reply("Sidearm (optional): " + weapons[index].Name)
	removeIndex(weapons, index)
}

// Takes as input an array of strings and returns from it a random index
// if 'remove' is true, it will also remove the generated item from the input array
func generateRandomFromSlice(input []framework.Weapon, weaptype framework.WeaponCategory) int {
	rand.Seed(time.Now().UnixNano())

	//
	// the for header does all that it is intended to do. It keeps
	// looping while the randomly generated element is not of the right type
	index := 0
	for index = rand.Intn(len(input)); input[index].Category != weaptype; index = rand.Intn(len(input)) {
	}
	return index
}

func removeIndex(input []framework.Weapon, index int) {
	input = append(input[:index], input[index+1:]...)
}
