package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type VoterViewableVoting struct {
	Name          string
	Id            string
	Description   string
	Open          int
	Ended         int
	VotesLeft int
	VotesPerToken int
	Candidates []VoterViewableCandidates
}
type VoterViewableVotings = []VoterViewableVoting

type VoterViewableCandidate struct {
	Name        string
	Id          string
	Description string
	Voted bool
}
type VoterViewableCandidates = []VoterViewableCandidate

type VoterViewableState struct {
	Votings []VoterViewableVoting
}

func constructVoterViewableVoting(votingid string, token string, db: sql.DB) {
	candidates, err := db.Query("SELECT id, FROM Candidate, Availability WHERE Candidate.id = Availability.candidateid AND Availability.votingid = ?, ", votingid)
	
}
func HandleVoterApiQuery(w http.ResponseWriter, r *http.Request) {
	fmt.Println("at voter api")
	params := r.URL.Query()
	db, _ := sql.Open("sqlite3", DB_NAME)
	token := r.Header.Get("Authorization")
	isVoter := 0
	db.QueryRow("SELECT COUNT(*) description FROM Token WHERE value = ?", token).Scan(&isVoter)
	if isVoter != 1 {
		fmt.Fprintf(w, "denied")
		return
	}
	action, actionExists := params["a"]
	if actionExists {
		switch strings.Join(action, "") {
		case "show":
			var name string
			var id string
			var description string
			var open int
			var ended int
			var votesleft int
			var votespertoken int
			var votingsStruct = VoterViewableVotings{}
			//var candidates = Candidates{}
			votings, err := db.Query("SELECT id FROM Voting WHERE NOT open = 0")
			if err == nil {
				for votings.Next() {
					votings.Scan(&id)
					constructVoterViewableVoting(id, token)
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
		default:
			fmt.Fprint(w, "unknown action")
		}
	}
}
