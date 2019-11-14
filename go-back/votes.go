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

type Vote struct {
	Id          string
	VotingId    string
	CandidateId string
	Token       string
}
type Votes = []Vote

func HandleVoteApiQuery(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	db, _ := sql.Open("sqlite3", DB_NAME)
	db.Exec("PRAGMA foreign_keys = ON")
	action, actionExists := params["a"]
	if actionExists {
		switch strings.Join(action, "") {
		case "add":
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Fprint(w, "err")
				break
			}
			var t Vote
			err = json.Unmarshal(body, &t)
			if err != nil || t.Id == "" || t.VotingId == "" || t.Token == "" {
				fmt.Fprint(w, "malformed json")
				break
			}
			var currentVotes int
			var allowedVotes int
			db.QueryRow("SELECT COUNT(*) FROM Vote WHERE token = ? AND votingId = ?", t.Token, t.VotingId).Scan(&currentVotes)
			db.QueryRow("SELECT votesPerToken FROM Voting WHERE id = ?", t.VotingId).Scan(&allowedVotes)
			if currentVotes < allowedVotes {
				_, err := db.Exec("INSERT INTO Vote(id, votingId, candidateId, token) VALUES(?, ?, ?, ?)", t.Id, t.VotingId, t.CandidateId, t.Token)
				if err != nil {
					fmt.Fprintf(w, "nonexisting candidate/voting/token")
					break
				}
				fmt.Fprint(w, "ok i guess")
			} else {
				fmt.Fprintf(w, "already voted")
			}
		case "show":
			var id string
			var votingid string
			var candidateid string
			var token string
			var votesStruct = Votes{}
			votes, err := db.Query("SELECT id, votingid, candidateid, token FROM Vote")
			if err == nil {
				for votes.Next() {
					votes.Scan(&id, &votingid, &candidateid, &token)
					votesStruct = append(votesStruct, Vote{id, votingid, candidateid, token})
				}
			} else {
				log.Fatal(err)
			}
			w.Header().Set("Content-Type", "application/json")
			votesJson, err := json.Marshal(votesStruct)
			if err != nil {
				log.Fatal(err)
			}
			w.Write(votesJson)
		case "del":
			target, targetExists := params["t"]
			if targetExists {
				_, err := db.Exec("DELETE FROM Vote WHERE id = ?", strings.Join(target, ""))
				if err == nil {
					fmt.Fprint(w, "deleted")
				} else {
					fmt.Println(err)
					fmt.Fprint(w, "something went wrong while deleting vote")
				}
			}
		default:
			fmt.Fprint(w, "unknown action")
		}
	}
}
