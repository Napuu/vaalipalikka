package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"math/rand"
	"strconv"
	"time"
)

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
		tx.Exec("INSERT INTO Token(value, valid, used) VALUES(?, 0)", token)
	}
	tx.Commit()
}
