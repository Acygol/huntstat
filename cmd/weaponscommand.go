package cmd

import (
	"fmt"
	"strings"

	"github.com/acygol/huntstat/framework"
)

const (
	primary = "primary"
	sidearm = "sidearm"
)

//
// WeaponsCommand is executed when someone calls 's!weapon(s)'
//
func WeaponsCommand(ctx framework.Context) {
	weapons := framework.Weapons
	var reply strings.Builder

	name := framework.GenerateRandomWeaponOnce(weapons, primary)
	fmt.Fprintf(&reply, "Primary: %s\n", name)

	// To make sure the first weapon isn't chosen again, I remove it from the slice
	name = framework.GenerateRandomWeaponOnce(weapons, primary)
	fmt.Fprintf(&reply, "Secondary: %s\n", name)

	name = framework.GenerateRandomWeaponOnce(weapons, sidearm)
	fmt.Fprintf(&reply, "Sidearm (optional): %s\n", name)

	ctx.Reply(reply.String())
}
