package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"log"
	"net/http"

	//"net/url"
	//"os"
	//"strconv"
	"strings"
)

type Candidate struct {
	Name        string
	Id          string
	Description string
}
type Candidates = []Candidate

func HandleCandidateApiQuery(w http.ResponseWriter, r *http.Request) {
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
			var t Candidate
			err = json.Unmarshal(body, &t)
			if err != nil || t.Description == "" || t.Name == "" || t.Id == "" {
				fmt.Fprint(w, "malformed json")
				break
			}
			_, err = db.Exec("INSERT INTO Candidate(name, id, description) VALUES(?, ?, ?)", t.Name, t.Id, t.Description)
			if err != nil {
				if err.Error() == "UNIQUE constraint failed: Candidate.id" {
					fmt.Fprint(w, "candidate already exists")
					break
				} else {
					log.Fatal(err)
				}
			}
			fmt.Fprint(w, "ok i guess")
		case "show":
			target, targetExists := params["t"]
			if !targetExists {
				var name string
				var id string
				var description string
				var candidatesStruct = Candidates{}
				candidates, err := db.Query("SELECT name, id, description FROM Candidate")
				if err == nil {
					for candidates.Next() {
						candidates.Scan(&name, &id, &description)
						candidatesStruct = append(candidatesStruct, Candidate{name, id, description})
					}
				} else {
					log.Fatal(err)
				}
				w.Header().Set("Content-Type", "application/json")
				candidatesJson, err := json.Marshal(candidatesStruct)
				if err != nil {
					log.Fatal(err)
				}
				w.Write(candidatesJson)
			} else {
				var name string
				var id string
				var description string
				row := db.QueryRow("SELECT name, id, description FROM Candidate WHERE id = ?", strings.Join(target, ""))
				switch err := row.Scan(&name, &id, &description); err {
				case sql.ErrNoRows:
					fmt.Fprint(w, "no candidates matching that id")
				case nil:
					w.Header().Set("Content-Type", "application/json")
					candidateJson, _ := json.Marshal(Candidate{name, id, description})
					w.Write(candidateJson)
				default:
					log.Panic(err)
				}
			}
		case "del":
			target, targetExists := params["t"]
			if targetExists {
				_, err := db.Exec("DELETE FROM Candidate WHERE id = ?", strings.Join(target, ""))
				if err != nil {
					fmt.Fprint(w, "deleted")
				} else {
					fmt.Fprint(w, "something went wrong while deleting candidate")
				}
			}
		default:
			fmt.Fprint(w, "unknown action")
		}
	}
}
