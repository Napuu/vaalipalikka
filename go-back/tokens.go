package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func HandleTokenApiQuery(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	db, _ := sql.Open("sqlite3", DB_NAME)
	db.Exec("PRAGMA foreign_keys = ON")
	action, _ := params["a"]
	switch strings.Join(action, "") {
	case "generate":
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
	case "show":
		tokens, ok := db.Query("SELECT value, valid FROM Token")
		var value string
		var valid int
		var tokensStruct = Tokens{}
		if ok == nil {
			for tokens.Next() {
				tokens.Scan(&value, &valid)
				tokensStruct = append(tokensStruct, Token{value, valid})
			}
		} else {
			log.Fatal(ok)
		}
		w.Header().Set("Content-Type", "application/json")
		tokensJson, err := json.Marshal(tokensStruct)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(tokensJson)
	case "toggle":
		//_t, tExists := params["t"]
		//_v, vExists := params["v"]
		_t, _ := params["t"]
		_v, _ := params["v"]
		v, err := strconv.Atoi(strings.Join(_v, ""))
		t := strings.Join(_t, "")
		_, err = db.Exec("UPDATE Token SET valid = ? WHERE value = ?", v, t)
		if err != nil {
			fmt.Println(err)
			fmt.Fprintf(w, "something went wrong while updating token status")
			break
		}
		fmt.Fprintf(w, "ok i guess")
	}
}

func GenerateTokens(amount int) map[string]struct{} {
	const TOKENLENGTH = 6
	var tokens = make(map[string]struct{})
	var exists = struct{}{}
	upper := 1
	for i := 0; i < TOKENLENGTH; i++ {
		upper = upper * 10
	}
	a := make([]int, upper)
	for i := 0; i < upper; i++ {
		a[i] = i
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	for i := 0; i < amount; i++ {
		tokens[pad(a[i], TOKENLENGTH)] = exists
	}
	return tokens
}

func pad(n int, tl int) string {
	asStr := strconv.Itoa(n)
	missing := tl - len(asStr)
	padding := ""
	for i := 0; i < missing; i++ {
		padding = "0" + padding
	}
	return padding + asStr
}

func InsertNewTokens(tokens map[string]struct{}, db *sql.DB) {
	_, err := db.Exec("DELETE FROM Token")
	if err != nil {
		log.Fatal(err)
	}
	tx, _ := db.Begin()
	for token, _ := range tokens {
		tx.Exec("INSERT INTO Token(value, valid) VALUES(?, 0)", token)
	}
	tx.Commit()
}
