package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
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
	db, _ := sql.Open("sqlite3", DB_NAME)
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
			_, err = db.Exec("INSERT INTO Voting(name, id, description, open, ended, votespertoken) VALUES(?, ?, ?, ?, ?, ?)", t.Name, t.Id, t.Description, t.Open, t.Ended, t.VotesPerToken)
			if err != nil {
				// TODO - actually check what the error is, now it is assumed to be error with Unique constraint Voting.id
				_, err = db.Exec("UPDATE Voting SET name = ?, description = ?, open = ?, ended = ?, votespertoken = ? WHERE id = ?", t.Name, t.Description, t.Open, t.Ended, t.VotesPerToken, t.Id)
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
				votings, err := db.Query("SELECT name, id, description, open, ended, votespertoken FROM Voting")
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
				}
				w.Write(votingsJson)
			} else {
				var name string
				var id string
				var description string
				var open int
				var ended int
				var votespertoken int
				row := db.QueryRow("SELECT name, id, description FROM Voting WHERE id = ?", strings.Join(target, ""))
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
				_, err := db.Exec("DELETE FROM Voting WHERE id = ?", strings.Join(target, ""))
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
