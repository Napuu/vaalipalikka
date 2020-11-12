package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	//"net/url"
	"database/sql"
	"os"
	"strings"
	"time"
)

const CONNECTION_STRING string = "postgres://vaalit:vaalit@postgres:5432/vaalit?sslmode=disable"

var db *sql.DB

func main() {
	switch os.Args[1] {
	case "start":
		fmt.Println("starting server...")
	case "tokens":
		GenerateTokens(100)
	}
	fmt.Println("opening database")
	var err error
	db, err = sql.Open("postgres", CONNECTION_STRING)
	db.SetMaxOpenConns(500)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Second)
	if err != nil {
		log.Fatal("error connecting to postgres")
	}
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/api", HandleApiQuery)
	fmt.Println("listening...")
	http.ListenAndServe(":8281", nil)
}

type Token struct {
	Value string
	Valid int
}
type Tokens = []Token

func HandleApiQuery(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	action, actionExists := params["action"]
	if (actionExists) {
		fmt.Println("at main function", action[0])
	} else {
		fmt.Println("no action parameter provided")
		return
	}
	if actionExists {
		switch action[0] {
		case "token":
			if !CheckQueryPermissions(w, r, 2) {
				return
			}
			HandleTokenApiQuery(w, r)
		case "candidate":
			if !CheckQueryPermissions(w, r, 2) {
				return
			}
			HandleCandidateApiQuery(w, r)
		case "voting":
			if !CheckQueryPermissions(w, r, 2) {
				return
			}
			HandleVotingApiQuery(w, r)
		case "vote":
			if !CheckQueryPermissions(w, r, 2) {
				return
			}
			HandleVoteApiQuery(w, r)
		// this is not an actual login just checking header, jwt or something is needed if this had "actual" use
		case "login":
			HandleLoginApiQuery(w, r)
		case "availability":
			if !CheckQueryPermissions(w, r, 2) {
				return
			}
			HandleAvailabilityApiQuery(w, r)
		case "voter":
			fmt.Println("going to voter api query")
			if !CheckQueryPermissions(w, r, 1) {
				return
			}
			HandleVoterApiQuery(w, r)
		default:
			fmt.Println("from default case", action[0])
			fmt.Fprintf(w, "asdf, %s!", r.URL.Path[1:])
		}
	}
}
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func toJSON(m interface{}) string {
	js, err := json.Marshal(m)
	if err != nil {
		//log.Fatal(err)
		//Panic
	}
	return strings.ReplaceAll(string(js), ",", ", ")
}
