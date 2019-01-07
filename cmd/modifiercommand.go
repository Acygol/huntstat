package cmd

import (
	"math/rand"
	"time"

	"github.com/acygol/huntstat/framework"
)

//
// ModifierCommand is executed when someone calls 's!modifier(s)'
//
func ModifierCommand(ctx framework.Context) {
	rand.Seed(time.Now().UnixNano())
	modifiers := []string{
		"Everything Goes", "Shotguns", "Rifles", "No Scope",
		"Silent", "Classic ", "Pistol only",
	}
	ctx.Reply(modifiers[rand.Intn(len(modifiers))])
}
