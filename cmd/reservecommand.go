package cmd

import (
	"github.com/acygol/huntstat/framework"
	"math/rand"
	"time"
)

func ReservesCommand(ctx framework.Context) {
	rand.Seed(time.Now().UnixNano())
	reserves := []string {
		"Whitehart Island", "Logger's Point", "Settler Creeks", "Redfeather Falls",
		"Hirschfelden", "Hemmeldal", "Rougarou Bayou", "Val-des-Bois",
		"Bushrangers Run", "Whiterime Ridge", "Timbergold Trails", "Piccabeen Bay",
	}

	ctx.Reply(reserves[rand.Intn(len(reserves))])
}
