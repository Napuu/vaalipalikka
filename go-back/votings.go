package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Voting struct {
	Name          string
	Id            string
	Description   string
	Open          int
	Ended         int
	VotesPerToken int
}
type Votings = []Voting

func HandleVotingApiQuery(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	action, actionExists := params["a"]
	if actionExists {
		switch strings.Join(action, "") {
		case "add":
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Fprint(w, "err")
				break
			}
			var t Voting
			err = json.Unmarshal(body, &t)
			if err != nil || t.Name == "" || t.Id == "" {
				fmt.Println(err)
				fmt.Fprint(w, "malformed json")
				break
			}
			fmt.Println("trying")
			_, err = db.Exec("INSERT INTO Voting(name, id, description, open, ended, votespertoken) VALUES($1, $2, $3, $4, $5, $6)", t.Name, t.Id, t.Description, t.Open, t.Ended, t.VotesPerToken)
			if err != nil {
				_, err2 := db.Exec("UPDATE Voting SET name = $1, description = $2, open = $3, ended = $4, votespertoken = $5 WHERE id = $6", t.Name, t.Description, t.Open, t.Ended, t.VotesPerToken, t.Id)
				fmt.Println(err2)
				fmt.Fprint(w, "replaced")
				break
			}
			fmt.Fprint(w, "ok i guess")
		case "show":
			target, targetExists := params["t"]
			if !targetExists {
				var name string
				var id string
				var description string
				var open int
				var ended int
				var votespertoken int
				var votingsStruct = Votings{}
				votings, err := db.Query("SELECT name, id, description, open, ended, votespertoken FROM Voting ORDER BY hidden_id")
				defer votings.Close()
				if err == nil {
					for votings.Next() {
						votings.Scan(&name, &id, &description, &open, &ended, &votespertoken)
						votingsStruct = append(votingsStruct, Voting{name, id, description, open, ended, votespertoken})
					}
				} else {
					log.Fatal(err)
				}
				w.Header().Set("Content-Type", "application/json")
				votingsJson, err := json.Marshal(votingsStruct)
				if err != nil {
					log.Fatal(err)
					fmt.Fprintf(w, "malformed json")
					break
				}
				w.Write(votingsJson)
			} else {
				var name string
				var id string
				var description string
				var open int
				var ended int
				var votespertoken int
				row := db.QueryRow("SELECT name, id, description FROM Voting WHERE id = $1", strings.Join(target, ""))
				switch err := row.Scan(&name, &id, &description); err {
				case sql.ErrNoRows:
					fmt.Fprint(w, "no votings matching that id")
				case nil:
					w.Header().Set("Content-Type", "application/json")
					votingJson, _ := json.Marshal(Voting{name, id, description, open, ended, votespertoken})
					w.Write(votingJson)
				default:
					log.Panic(err)
				}
			}
		case "del":
			target, targetExists := params["t"]
			if targetExists {
				_, err := db.Exec("DELETE FROM Voting WHERE id = $1", strings.Join(target, ""))
				if err == nil {
					fmt.Fprint(w, "deleted")
				} else {
					fmt.Fprint(w, "something went wrong while deleting voting")
				}
			}
		default:
			fmt.Fprint(w, "unknown action")
		}
	}
}
