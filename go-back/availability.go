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

type Availability struct {
	CandidateId string
	VotingId    string
}
type Availabilities = []Availability

func HandleAvailabilityApiQuery(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
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
		_, err = db.Exec("INSERT INTO Availability(candidateid, votingid) VALUES($1, $2)", t.CandidateId, t.VotingId)
		if err != nil {
			_, err = db.Exec("UPDATE Availability SET candidateid = $1, votingid = $2 WHERE candidateid = $3 AND votingid = $4", t.CandidateId, t.VotingId, t.CandidateId, t.VotingId)
			fmt.Println("here???2")
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
		err = nil
		err = json.Unmarshal(body, &t)
		if err != nil {
			fmt.Println("???ejson")
			fmt.Fprint(w, "malformed json")
			break
		}
		if len(t) > 0 {
			fmt.Println("now trying to execute delete")
			_, err := db.Exec("DELETE FROM Availability WHERE votingid = $1", t[0].VotingId)
			//fmt.Println(err)
			//fmt.Println("error from exec", err.Error())
			for _, element := range t {
				err = nil
				fmt.Println("adding ", element.CandidateId, element.VotingId)
				_, err = db.Exec("INSERT INTO Availability(CandidateId, VotingId) VALUES($1, $2)", element.CandidateId, element.VotingId)
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
