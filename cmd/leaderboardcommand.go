package cmd

import (
	"fmt"
	"github.com/acygol/huntstat/framework"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type Record struct {
	animal		string
	holder		string
	score		float64
	scoresheet	string
}

func NewRecord(animal string) *Record {
	rec := new(Record)
	rec.animal = animal
	rec.holder = "<nobody>"
	rec.score = 0.0
	rec.scoresheet = "0"
	return rec
}

// slice of Records for each animal
var records [len(ANIMALS)]*Record

func LeaderboardCommand(ctx framework.Context) {
	for i := 0; i < len(records); i++ {
		records[i] = NewRecord(ANIMALS[i])
	}

	/*
	// matches:
	// 		<!-- 1654.51 11 65487546 -->
	// which are the comments in the widget's HTML code describing the animal field:
	// 		1. floating point number indicating the animal's score
	//		2. integer indicating the # of this animal the user has harvested
	//		3. integer indicating the scoresheet id used in the URL
	*/
	regex := regexp.MustCompile(`(<!--.(?P<score>[0-9.]+).(?P<harvested>[0-9]+).(?P<scoresheet>[0-9]+).-->)`)

	/*
	// Go through all users and for each, register each animal's score
	// if one of the scores beats the record, then the Record holder's name
	// is adjusted.
	//
	// The reason why this algorithm doesn't go through all animals and for each
	// go through the user's widget is because the HTML is scraped too often. In
	// this version, the HTML of a user is scraped exactly once per user. While
	// in the latter version, the HTML is scraped 'len(ANIMALS)' amount of times
	*/
	rows, err := ctx.Conf.DbHandle.Query("SELECT hunter_name FROM users WHERE guild_id = ?", ctx.Guild.ID)
	if err != nil {
		ctx.Reply("Unable to retrieve from database, contact the maintainer of this bot for more information")
		fmt.Println("error retrieving from database,", err)
		return
	}
	defer rows.Close()
	for huntername := ""; rows.Next(); {
		err := rows.Scan(&huntername)
		if err != nil {
			fmt.Println("Error attempting to scan the next row,", err)
			break
		}

		/*
		// Load HTML
		*/
		resp, _ := http.Get(GetWidget(huntername))
		bdy, _ := ioutil.ReadAll(resp.Body)
		body := string(bdy) // stringify body

		/*
		// Allows submatches to be referenced with their name
		// as defined in the regexp:
		//		(?P<score>...)
		// e.g.,
		//		result["score"]
		// yields the score part of a match
		*/
		match := regex.FindAllString(body, -1)
		result := make([]map[string]string, len(ANIMALS), len(ANIMALS)+1)
		j := 0

		for _, comment := range match {
			submatch := regex.FindStringSubmatch(comment)
			tmpmap := make(map[string]string)
			for i, subname := range regex.SubexpNames() {
				if i != 0 && subname != "" {
					tmpmap[subname] = submatch[i]
				}
			}
			result[j] = tmpmap
			j++
		}

		for i := 0; i < len(result); i++ {
			score, _ := strconv.ParseFloat(result[i]["score"], 64)
			if score > records[i].score {
				records[i].score = score
				records[i].holder = huntername
				records[i].scoresheet = result[i]["scoresheet"]
			}
		}
	}

	var reply strings.Builder
	if len(ctx.Args) > 0 {
		// a specific animal was given
		animal := strings.Join(ctx.Args, " ")

		if !isValidAnimalName(animal) {
			ctx.Reply("Invalid animal name")
			return
		}
		idx := getAnimalRecordIndex(animal)
		fmt.Fprintf(&reply, "\n\n%s (%s):\n%+v [<https://www.thehunter.com/#profile/%s/score/%s>]", records[idx].animal, records[idx].holder, records[idx].score, records[idx].holder, records[idx].scoresheet)
	} else {
		for i, record := range records {
			if i % 15 == 0 && i != 0 {
				ctx.Reply(reply.String())
				reply.Reset()
			}
			fmt.Fprintf(&reply, "\n\n%s (%s):\n%+v [<https://www.thehunter.com/#profile/%s/score/%s>]", record.animal, record.holder, record.score, record.holder, record.scoresheet)
		}
	}
	ctx.Reply(reply.String())
}

func isValidAnimalName(name string) bool {
	for _, n := range ANIMALS {
		if strings.EqualFold(name, n) {
			return true
		}
	}
	return false
}

func getAnimalRecordIndex(name string) int {
	for i, record := range records {
		if strings.EqualFold(record.animal, name) {
			return i
		}
	}
	return -1
}
