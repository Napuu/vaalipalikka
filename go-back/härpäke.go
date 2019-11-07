package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	//"net/url"
	"os"
	"strconv"
	"strings"
)

const DB_NAME string = "vaalit.db?_foreign_keys=on"

func main() {
	fmt.Println("arguments")
	fmt.Println(os.Args[1])
	switch os.Args[1] {
	case "start":
		fmt.Println("starting server...")
	case "prepare":
		InitializeDb()
	case "tokens":
		GenerateTokens(1000)
	}
	fmt.Println("opening database")
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/api", HandleApiQuery)
	http.ListenAndServe(":8281", nil)
}

type Token struct {
	Value string
	Valid int
	Used  int
}
type Tokens = []Token

func HandleApiQuery(w http.ResponseWriter, r *http.Request) {
	db, _ := sql.Open("sqlite3", "./vaalit.db")
	params := r.URL.Query()
	action, actionExists := params["action"]
	if actionExists {
		switch action[0] {
		case "generatetokens":
			_n, nExists := params["n"]
			var n int
			if nExists {
				n, _ = strconv.Atoi(strings.Join(_n, ""))
			} else {
				n = 100
			}
			newTokens := GenerateTokens(n)
			InsertNewTokens(newTokens, db)
			fmt.Fprintf(w, "tokens generated")
		case "showtokens":
			tokens, ok := db.Query("SELECT value, valid, used FROM Token")
			var value string
			var valid int
			var used int
			var tokensStruct = Tokens{}
			if ok == nil {
				for tokens.Next() {
					tokens.Scan(&value, &valid, &used)
					tokensStruct = append(tokensStruct, Token{value, valid, used})
				}
			}
			w.Header().Set("Content-Type", "application/json")
			tokensJson, err := json.Marshal(tokensStruct)
			if err != nil {
				log.Fatal(err)
			}
			w.Write(tokensJson)
		case "candidate":
			HandleCandidateApiQuery(w, r)
		case "voting":
			HandleVotingApiQuery(w, r)
		case "vote":
			HandleVoteApiQuery(w, r)
		case "login":
			HandleLoginApiQuery(w, r)
		default:
			fmt.Println(action[0])
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
