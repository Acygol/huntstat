package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/acygol/huntstat/framework"
)

//
// LeaderboardCommand is executed when someone calls 's!leaderboard(s)'
//
func LeaderboardCommand(ctx framework.Context) {
	//
	// Each index in records has a corresponding animal
	// and thus initialization is necessary to establish
	// this relationship
	//
	var records []*framework.Record
	for i := 0; i < len(framework.Animals); i++ {
		records = append(records, framework.NewRecord(framework.Animals[i]))
	}

	//
	// The widget's HTML page has comments for each animal
	// tile with the following format:
	//		<!-- 1654.51 11 65487546 -->
	// Where
	//		1654.51 	is the animal's score
	//		11 			is the amount of this animal the user has harvested
	//		65487546	is the scoresheet id used in the scoresheet's URL
	// Extracting these comments from the HTML code is done
	// with a regular expression, which is used further
	// down the code but is compiled here:
	//
	regex := regexp.MustCompile(`(<!--.(?P<score>[0-9.]+).(?P<harvested>[0-9]+).(?P<scoresheet>[0-9]+).-->)`)

	//
	// Query the database to retrieve the hunter names of
	// all registered community members within the server
	// in which the command was executed
	//
	rows, err := ctx.Conf.Database.Handle.Query("SELECT hunter_name FROM users WHERE guild_id = ?", ctx.Guild.ID)
	if err != nil {
		ctx.Reply("Unable to retrieve from database, contact the maintainer of this bot for more information.")
		fmt.Println("erro retrieving from database,", err)
		return
	}
	defer rows.Close()

	//
	// For each member in the community, process their widget
	// and compare each score against that of its corresponding
	// record. The HTML is scraped exactly $x$ amount of times
	// where $x$ is the amount of members the community has.
	// Doing it the other way around: "for each animal, process
	// widgets of all community members" results in the widget
	// being scraped more than it should; $x * len(ANIMALS)$
	// times as opposed to just $x$ times.
	//
	for huntername := ""; rows.Next(); {
		err := rows.Scan(&huntername)
		if err != nil {
			fmt.Println("error while scanning the next row,", err)
			break
		}

		//
		// Get the widget's page, read its HTML
		// and then stringify it such that it becomes
		// easier to use in finding all regex matches
		//
		resp, _ := http.Get(GetURL(huntername, WidgetURL))
		retbody, _ := ioutil.ReadAll(resp.Body)
		body := string(retbody)

		//
		// The regex used to extract necessary data from
		// the HTML page is composed of 3 named groups:
		//		(?P<score>), (?P<harvested>) and (?P<scoresheet>)
		// The following extracts said groups and puts
		// them in an associative array (= map) so that it
		// becomes easier to call upon specific submatches:
		//		result["score"]
		// yields the submatch related to the animal score,
		// as defined by the regular expression used
		//
		matches := regex.FindAllString(body, -1)

		//
		// result is a slice of maps. Each index in the slice
		// is an associative array for the submatches of each
		// animal in the widget.
		//		result[0]["score"]
		// yields the score of the first animal tile, etc.
		//
		result := make([]map[string]string, len(framework.Animals))

		//
		//
		//
		j := 0
		for _, match := range matches {
			submatch := regex.FindStringSubmatch(match)

			//
			// result is already initializes as a slice
			// of maps, but the maps themselves are not
			// initialized as of yet
			//
			tmpmap := make(map[string]string)
			for i, submatchname := range regex.SubexpNames() {
				if i != 0 && submatchname != "" {
					tmpmap[submatchname] = submatch[i]
				}
			}
			result[j] = tmpmap
			j++
		}

		for i := range result {

			//
			// The score part of a record is stored as an float
			// (because it is a float), but since the score
			// from the HTML comment is extracted as a string,
			// the conversion is neccesary
			//
			score, _ := strconv.ParseFloat(result[i]["score"], 64)

			//
			// The record must only be updated when the
			// recently retrieved score exceeds that of
			// the current record
			//
			if score > records[i].Score {
				records[i].Score = score
				records[i].Holder = huntername
				records[i].Scoresheet = result[i]["scoresheet"]
			}
		}
	}

	//
	// All members are processed and all records have been
	// populated with data
	//
	var reply strings.Builder

	//
	// ctx.Args for 's!leaderboard(s)' is not 0 when
	// an animal name has been passed along with the
	// command call. In that case, it's sufficient
	// to only show that animal's record. The core
	// logic of this command remains the same
	//
	if len(ctx.Args) > 0 {
		//
		// Some animal names have spaces in them:
		//		"Alpine Ibex"
		// However, the command processor uses the
		// space character as an argument delimiter:
		//		ctx.Args[0] = "Alpine"
		//		ctx.Args[1] = "Ibex"
		// In this case, the delimiter can be ignored
		// and all arguments joined together equals
		// the animal name
		//
		animal := strings.Join(ctx.Args, " ")

		if !framework.IsValidAnimalName(animal) {
			ctx.Reply("Invalid animal name")
			return
		}

		//
		// I can assume that idx is never going to be -1
		// in this case, because if it ever could, the above
		// if statement would fail before this is ever called
		//
		idx := framework.GetRecordIndexByAnimal(records, animal)
		rec := records[idx]
		fmt.Fprintf(&reply, "%s (%s):\n%+v [<https://www.thehunter.com/#profile/%s/score/%s>]", rec.Animal, rec.Holder, rec.Score, rec.Holder, rec.Scoresheet)
	} else {
		for i, rec := range records {
			//
			// Sending each record as a separate message results in the
			// last ones being sent with an enormous delay for causes
			// unknown to me. To prevent this, I gather them into a
			// single strings.Builder and send them all together at once.
			// However, Discord only allows messages up to 2'000 characters
			// and so, I must divide them further up into smaller chunks
			//
			if i%15 == 0 && i != 0 {
				ctx.Reply(reply.String())
				reply.Reset()
			}
			fmt.Fprintf(&reply, "\n\n%s (%s):\n%+v [<https://www.thehunter.com/#profile/%s/score/%s>]", rec.Animal, rec.Holder, rec.Score, rec.Holder, rec.Scoresheet)
		}
	}
	ctx.Reply(reply.String())
}
