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

const CONNECTION_STRING string = "postgres://vaalit:vaalit@localhost:4531/vaalit?sslmode=disable"

var db *sql.DB

func main() {
	switch os.Args[1] {
	case "start":
		fmt.Println("starting server...")
	case "prepare":
		InitializeDb()
		return
	case "tokens":
		GenerateTokens(1000)
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
	http.HandleFunc("/drop", drop)
	http.HandleFunc("/api", HandleApiQuery)
	http.ListenAndServe(":8281", nil)
}

func drop(w http.ResponseWriter, r *http.Request) {
	InitializeDb()
	fmt.Fprintf(w, "done")
}

type Token struct {
	Value string
	Valid int
}
type Tokens = []Token

func HandleApiQuery(w http.ResponseWriter, r *http.Request) {
	fmt.Println("????")
	params := r.URL.Query()
	action, actionExists := params["action"]
	fmt.Println("at main function", action[0])
	if actionExists {
		switch action[0] {
		case "token":
			HandleTokenApiQuery(w, r)
		case "candidate":
			HandleCandidateApiQuery(w, r)
		case "voting":
			HandleVotingApiQuery(w, r)
		case "vote":
			HandleVoteApiQuery(w, r)
		case "login":
			HandleLoginApiQuery(w, r)
		case "availability":
			HandleAvailabilityApiQuery(w, r)
		case "voter":
			fmt.Println("going to voter api query")
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
