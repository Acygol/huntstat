package cmd

import (
	"github.com/acygol/huntstat/framework"
	"math/rand"
	"time"
)

func ModifierCommand(ctx framework.Context) {
	rand.Seed(time.Now().UnixNano())
	modifiers := []string {
		"Everything Goes", "Shotguns", "Rifles", "No Scope",
		"Silent", "Classic ", "Pistol only",
	}
	ctx.Reply(modifiers[rand.Intn(len(modifiers))])
}
