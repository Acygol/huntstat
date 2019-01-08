package cmd

import (
	"math/rand"
	"time"

	"github.com/acygol/huntstat/framework"
)

//
// AnimalsCommand is executed when someone calls 's!animal(s)'
//
func AnimalsCommand(ctx framework.Context) {
	rand.Seed(time.Now().UnixNano())
	ctx.Reply(framework.Animals[rand.Intn(len(framework.Animals))].Name)
}
