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
	var animalName string
	rand.Seed(time.Now().UnixNano())
	if len(ctx.Args) > 0 {
		reserve, err := framework.GetReserveFromName(ctx.Args[0])
		if err != nil {
			ctx.Reply("Invalid reserve")
			return
		}
		animals := framework.GetAnimalsOnReserve(reserve)
		animalName = animals[rand.Intn(len(animals))].Name
	} else {
		animalName = framework.Animals[rand.Intn(len(framework.Animals))].Name
	}
	ctx.Reply(animalName)
}
