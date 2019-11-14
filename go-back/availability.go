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

type Availability struct {
	CandidateId string
	VotingId    string
}
type Availabilities = []Availability

func HandleAvailabilityApiQuery(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	db, _ := sql.Open("sqlite3", DB_NAME)
	db.Exec("PRAGMA foreign_keys = ON")
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
			_, err = db.Exec("UPDATE Availability SET candidateid = $cid, votingid = $vid WHERE candidateid = $cid AND votingid = $vid", t.CandidateId, t.VotingId, t.CandidateId, t.VotingId)
			log.Panic(err)
		}
		fmt.Fprint(w, "ok i guess")
	case "show":
		availabilities, ok := db.Query("SELECT candidateid, votingid FROM Availability")
		var candidateid string
		var votingid string
		var availabilitiesStruct = Availabilities{}
		if ok == nil {
			for availabilities.Next() {
				availabilities.Scan(&candidateid, &votingid)
				availabilitiesStruct = append(availabilitiesStruct, Availability{candidateid, votingid})
			}
		} else {
			log.Fatal(ok)
		}
		w.Header().Set("Content-Type", "application/json")
		availabilitiesJson, err := json.Marshal(availabilitiesStruct)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(availabilitiesJson)
		// case "del":

	case "clearadd":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprint(w, "err")
			break
		}
		var t []Availability
		//var availability Availability
		fmt.Println("clearadd")
		err = json.Unmarshal(body, &t)
		if len(t) > 0 {
			db.Exec("DELETE FROM Availability WHERE votingid = ?", t[0].VotingId)
			for _, element := range t {
				err = nil
				fmt.Println("adding ", element.CandidateId, element.VotingId)
				_, err = db.Exec("INSERT INTO Availability(CandidateId, VotingId) VALUES(?, ?)", element.CandidateId, element.VotingId)
				fmt.Println(err)
			}
		}

		//_, err = db.Exec("INSERT INTO Availability(candidateid, votingid) VALUES(?, ?)", t.CandidateId, t.VotingId)
		//if err != nil {
		//_, err = db.Exec("UPDATE Availability SET candidateid = $cid, votingid = $vid WHERE candidateid = $cid AND votingid = $vid", t.CandidateId, t.VotingId, t.CandidateId, t.VotingId)
		//log.Panic(err)
		//}
		fmt.Fprint(w, "ok i guess")
	}
}
