package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
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
	db, _ := sql.Open("postgres", CONNECTION_STRING)
	//db.Exec("PRAGMA foreign_keys = ON")
	action, actionExists := params["a"]
	fmt.Println("adding candidate ???")
	if actionExists {
		fmt.Println("adding candidate ??")
		switch strings.Join(action, "") {
		case "add":
			fmt.Println("adding candidate")
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Println("err")
				fmt.Fprint(w, "err")
				break
			}
			var t Candidate
			err = json.Unmarshal(body, &t)
			if err != nil || t.Name == "" || t.Id == "" {
				fmt.Println("json")
				fmt.Println(err)
				fmt.Fprint(w, "malformed json")
				break
			}
			fmt.Println("now inserting")
			_, err = db.Exec("INSERT INTO Candidate(name, id, description) VALUES($1, $2, $3)", t.Name, t.Id, t.Description)
			if err != nil {
				if err.Error() == "UNIQUE constraint failed: Candidate.id" {
					// fmt.Fprint(w, "candidate already exists")
					_, err2 := db.Exec("UPDATE Candidate SET name = $1, description = $2 WHERE id = $3", t.Name, t.Description, t.Id)
					fmt.Println(err2)
					fmt.Fprintf(w, "replaced")
					break
				} else {
					fmt.Println("fatality")
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
				row := db.QueryRow("SELECT name, id, description FROM Candidate WHERE id = $1", strings.Join(target, ""))
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
				_, err := db.Exec("DELETE FROM Candidate WHERE id = $2", strings.Join(target, ""))
				if err == nil {
					fmt.Fprint(w, "deleted")
				} else {
					fmt.Fprint(w, "something went wrong while deleting candidate")
					fmt.Println(err)
				}
			}
		default:
			fmt.Fprint(w, "unknown action")
		}
	}
}
