package cmd

import (
	"github.com/acygol/huntstat/framework"
	"math/rand"
	"time"
)

func AnimalsCommand(ctx framework.Context) {
	rand.Seed(time.Now().UnixNano())
	animals := []string {
		"Alpine Ibex", "Arctic Fox", "Banteng", "Bighorn Sheep",
		"Bison", "Black Bear", "Blacktail Deer", "Bobcat",
		"Brown Bear", "Canada Goose", "Cottontail Rabbit", "Coyote",
		"Dall Sheep", "Eurasian Lynx", "European Rabbit", "Feral Goat",
		"Feral Hog", "Grey Wolf", "Grizzly Bear", "Magpie Goose",
		"Duck", "Moose", "Mule Deer", "Pheasant",
		"Ptarmigan", "Polar Bear", "Red Deer", "Red Fox",
		"Red Kangaroo", "Reindeer", "Rocky Mountain Elk", "Roe Deer",
		"Roosevelt Elk", "Rusa Deer", "Sambar Deer", "Sitka Deer",
		"Snowshoe Hare", "Turkey", "Water Buffalo", "Whitetail Deer",
		"Wild Boar",
	}
	ctx.Reply(animals[rand.Intn(len(animals))])
}
