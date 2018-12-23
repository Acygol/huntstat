package cmd

import (
	"github.com/acygol/huntstat/framework"
	"math/rand"
	"time"
)

// Also used in leaderboardcommand.go
var ANIMALS = [...]string {
	"Alpine Ibex", "American Black Duck", "Arctic Fox", "Banteng",
	"Bighorn Sheep", "Bison", "Black Bear", "Blacktail Deer",
	"Bobcat", "Brown Bear", "Canada Goose", "Cottontail Rabbit",
	"Coyote", "Dallsheep", "Eurasian Lynx", "European Rabbit",
	"Feral Goat", "Feral Hog", "Gadwall", "Grey Wolf",
	"Grizzly Bear", "Magpie Goose", "Mallard", "Moose",
	"Mule Deer", "Northern Pintail", "Pheasant", "Polar Bear",
	"Red Deer", "Red Fox", "Red Kangaroo", "Reindeer",
	"Rock Ptarmigan", "Rocky Mountain Elk", "Roe Deer", "Roosevelt Elk",
	"Rusa deer", "Sambar Deer", "Sitka Deer", "Snowshoe Hare",
	"Turkey", "Water Buffalo", "White-tailed Ptarmigan", "Whitetail",
	"Wild Boar", "Willow Ptarmigan",
}

func AnimalsCommand(ctx framework.Context) {
	rand.Seed(time.Now().UnixNano())
	ctx.Reply(ANIMALS[rand.Intn(len(ANIMALS))])
}
