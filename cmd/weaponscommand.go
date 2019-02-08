package cmd

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/acygol/huntstat/framework"
)

//
// WeaponsCommand is executed when someone calls 's!weapon(s)'
//
func WeaponsCommand(ctx framework.Context) {
	if !ctx.Cmd.ValidateArgs(len(ctx.Args)) {
		return
	}

	reserve, err := framework.GetReserveFromName(ctx.Args[0])
	if err != nil {
		ctx.Reply("Invalid reserve name")
		return
	}
	inventoryCap, err := strconv.ParseFloat(ctx.Args[1], 64)
	if err != nil {
		log.Println("error parsing ctx.Args[1] as float64,", err)
		return
	}
	if inventoryCap != 10.0 && inventoryCap != 20 {
		ctx.Reply("Inventory capacity is either 10 or 20.")
		return
	}
	weapons := framework.Weapons
	var reply strings.Builder

	weapName, err := getRandomWeapon(weapons, framework.Primary, reserve, &inventoryCap)
	if err != nil {
		ctx.Reply(fmt.Sprintf("error while retrieving primary weapon: %v", err))
		return
	}
	fmt.Fprintf(&reply, "Primary: %v\n", weapName)

	// when the inventoryCap is 20, generate second primary
	if inventoryCap > 10.0 {
		weapName, err = getRandomWeapon(weapons, framework.Primary, reserve, &inventoryCap)
		if err != nil {
			ctx.Reply(fmt.Sprintf("error while retrieving second primary weapon: %v", err))
			return
		}
		fmt.Fprintf(&reply, "Second primary: %v\n", weapName)
	}
	weapName, err = getRandomWeapon(weapons, framework.Sidearm, reserve, &inventoryCap)
	if err != nil {
		ctx.Reply(fmt.Sprintf("error while retrieving sidearm weapon: %v", err))
		return
	}
	fmt.Fprintf(&reply, "Sidearm: %v", weapName)
	ctx.Reply(reply.String())
}

//
// getRandomWeapon acts as a local wrapper for framework.GenerateRandomWeaponOnce
//
func getRandomWeapon(weapons []framework.Weapon, weaptype string, reserve framework.Reserve, inventoryCap *float64) (string, error) {
	weapon, err := framework.GenerateRandomWeaponOnce(weapons, weaptype, reserve, *inventoryCap)
	if err != nil {
		return "", err
	}
	*inventoryCap -= weapon.Weight
	return weapon.Name, nil
}
