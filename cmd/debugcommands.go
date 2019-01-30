package cmd

import (
	"log"

	"github.com/acygol/huntstat/framework"
)

func UsersInMemory(ctx framework.Context) {
	for _, user := range framework.Users {
		log.Printf("User: %v", user)
	}
	ctx.Reply("Successfully logged")
}
