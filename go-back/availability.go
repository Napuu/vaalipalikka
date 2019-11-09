
package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
	"encoding/json"
	"log"
	"net/http"
	"io/ioutil"
	"strings"
)

type Availability struct {
	CandidateId string
	VotingId            string
}

func HandleAvailabilityApiQuery(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	db, _ := sql.Open("sqlite3", DB_NAME)
	action, _ := params["a"]
	switch strings.Join(action, "") {
	case "add":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprint(w, "err")
			break
		}
		var t Availability
		err = json.Unmarshal(body, &t)
		if err != nil || t.CandidateId == "" || t.CandidateId == "" {
			fmt.Println(err)
			fmt.Fprint(w, "malformed json")
			break
		}
		_, err = db.Exec("INSERT INTO Availability(candidateid, votingid) VALUES(?, ?)", t.CandidateId, t.VotingId)
		if err != nil {
			log.Panic(err)
		}
		fmt.Fprint(w, "ok i guess")
	// case "del":

	}
}

