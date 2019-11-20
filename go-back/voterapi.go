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

type VoterViewableVoting struct {
	Name          string
	Id            string
	Description   string
	Open          int
	Ended         int
	VotesLeft     int
	VotesPerToken int
	Candidates    []VoterViewableCandidate
}

type VoterViewableCandidate struct {
	Name        string
	Id          string
	Description string
	Voted       bool
}

func constructVoterViewableVoting(votingid string, token string, db sql.DB) VoterViewableVoting {
	candidates, _ := db.Query("SELECT id, name, description FROM Candidate, Availability WHERE Candidate.id = Availability.candidateid AND Availability.votingid = $1", votingid)
	fmt.Println("constructVoterViewableVoting")
	candidatesStruct := []VoterViewableCandidate{}
	votesUsed := 0
	for candidates.Next() {
		var candidateid string
		var candidatename string
		var candidatedescription string
		candidates.Scan(&candidateid, &candidatename, &candidatedescription)
		var votes int
		voted := false
		db.QueryRow("SELECT COUNT(*) FROM Vote WHERE votingid = $1 AND candidateid = $2 AND token = $3", votingid, candidateid, token).Scan(&votes)
		if votes != 0 {
			voted = true
			votesUsed += 1
		}
		candidatesStruct = append(candidatesStruct, VoterViewableCandidate{Name: candidatename, Id: candidateid, Description: candidatedescription, Voted: voted})
	}
	var name string
	var id string
	var description string
	var open int
	var ended int
	var votespertoken int
	db.QueryRow("SELECT name, id, description, open, ended, votespertoken FROM Voting WHERE id = $1", votingid).Scan(&name, &id, &description, &open, &ended, &votespertoken)
	votesleft := votespertoken - votesUsed
	return VoterViewableVoting{Name: name, Id: id, Description: description, Open: open, Ended: ended, VotesLeft: votesleft, VotesPerToken: votespertoken, Candidates: candidatesStruct}
}
func HandleVoterApiQuery(w http.ResponseWriter, r *http.Request) {
	fmt.Println("at voter api")
	params := r.URL.Query()
	token := r.Header.Get("Authorization")
	isVoter := 0
	db.QueryRow("SELECT COUNT(*) description FROM Token WHERE value = $1 AND valid = 1", token).Scan(&isVoter)
	if isVoter != 1 {
		fmt.Fprintf(w, "denied")
		return
	}
	action, actionExists := params["a"]
	if actionExists {
		switch strings.Join(action, "") {
		case "show":
			var id string
			var votingsStruct []VoterViewableVoting
			votings, err := db.Query("SELECT id FROM Voting")
			if err == nil {
				for votings.Next() {
					votings.Scan(&id)
					votingsStruct = append(votingsStruct, constructVoterViewableVoting(id, token, *db))
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
		case "vote":
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
			db.QueryRow("SELECT votesPerToken FROM Voting WHERE id = $1", t.VotingId).Scan(&allowedVotes)
			if currentVotes < allowedVotes {
				_, err := db.Exec("INSERT INTO Vote(id, votingId, candidateId, token) VALUES($1, $2, $3, $4)", t.Id, t.VotingId, t.CandidateId, t.Token)
				if err != nil {
					fmt.Println(err)
					fmt.Fprintf(w, "nonexisting candidate/voting/token")
					break
				}
				fmt.Fprint(w, "ok i guess")
			} else {
				fmt.Fprintf(w, "already voted")
			}

		default:
			fmt.Fprint(w, "unknown action")
		}
	}
}
