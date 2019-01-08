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
	if !ctx.CmdHandler.MustGet(ctx.CmdName).ValidateArgs(ctx) {
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

	weapon, _ := framework.GenerateRandomWeaponOnce(weapons, framework.Primary, reserve, inventoryCap)
	switch inventoryCap {
	case 10:
		// generate one primary and one sidearm
		fmt.Fprintf(&reply, "Primary: %s\n", weapon.Name)
		inventoryCap -= weapon.Weight

		weapon, _ = framework.GenerateRandomWeaponOnce(weapons, framework.Sidearm, reserve, inventoryCap)
		fmt.Fprintf(&reply, "Sidearm: %s\n", weapon.Name)
	case 20:
		fmt.Fprintf(&reply, "Primary: %s\n", weapon.Name)
		inventoryCap -= weapon.Weight

		weapon, _ = framework.GenerateRandomWeaponOnce(weapons, framework.Primary, reserve, inventoryCap)
		fmt.Fprintf(&reply, "Secondary: %s\n", weapon.Name)
		inventoryCap -= weapon.Weight

		weapon, _ = framework.GenerateRandomWeaponOnce(weapons, framework.Sidearm, reserve, inventoryCap)
		fmt.Fprintf(&reply, "Sidearm: %s\n", weapon.Name)
	}
	ctx.Reply(reply.String())
}
