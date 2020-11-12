package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
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
			db.QueryRow("SELECT COUNT(*) FROM Vote WHERE token = $1 AND votingId = $2", t.Token, t.VotingId).Scan(&currentVotes)
			db.QueryRow("SELECT visible FROM Voting WHERE id = $1", t.VotingId).Scan(&allowedVotes)
			if currentVotes < allowedVotes {
				_, err := db.Exec("INSERT INTO Vote(id, votingId, candidateId, token) VALUES($1, $2, $3, $4)", t.Id, t.VotingId, t.CandidateId, t.Token)
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
			votes, err := db.Query("SELECT id, votingid, candidateid, token FROM Vote ORDER BY hidden_id")
			defer votes.Close()
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
				_, err := db.Exec("DELETE FROM Vote WHERE id = $1", strings.Join(target, ""))
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
