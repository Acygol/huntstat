package cmd

import (
	"math/rand"
	"time"

	"github.com/acygol/huntstat/framework"
)

func ReservesCommand(ctx framework.Context) {
	rand.Seed(time.Now().UnixNano())

	ctx.Reply(framework.Reserves[rand.Intn(len(framework.Reserves))])
}
