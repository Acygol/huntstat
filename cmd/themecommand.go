package cmd

import (
	"github.com/acygol/huntstat/framework"
	"math/rand"
	"time"
)

func ThemeCommand(ctx framework.Context) {
	rand.Seed(time.Now().UnixNano())
	themes := []string {
		"What's up, Doc! Any Doc Monsignor/English suit,16GA or 12GA side by Side and .270",
		"Free to Play Experience: Only standard Clothes, .243 rifle and Single Shotgun, Basic Bino and Bleat Caller",
		"Sniper Elite: Any Ghillie suit, atleast one rifle with full scope (x12/Long Range) No Kills below 100m",
		"Almost Heroes: Trapper outfit, Muzzle Loader (classic) Longbow/cableback, Only 'natural?Wooden' Callers",
		"Silent But Deadly: Only Bow or Crossbow, (airgun allowed depending on team)",
		"High Noone! No 3D Camo, Only .45/70, 30-30, .30-06 or non-compound bows as main weapons (no semi auto sidearms)",
		"Boone and Crocket outfit, .300 Rifle, 12 Pump Action shotgun",
		"Say What? No Callers at all!",
		"All-In! No restrictions",
		"Hunt the one! Hunt only one species of Animal*",
	}
	ctx.Reply(themes[rand.Intn(len(themes))])
}
