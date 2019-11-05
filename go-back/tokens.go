package main

import (
	_ "github.com/mattn/go-sqlite3"
	"math/rand"
	"strconv"
	"time"
)

func generateToken() string {
	const tokenLength = 6
	offset := 10
	for i := 0; i < tokenLength+2; i++ {
		offset = offset * 10
	}
	rand.Seed(time.Now().UTC().UnixNano())
	base := rand.Intn(1<<32) + offset
	str := strconv.Itoa(base)
	return str[len(str)-tokenLength : len(str)]
}
func GenerateTokens(amount int) map[string]struct{} {
	var tokens = make(map[string]struct{})
	var exists = struct{}{}
	addedTokens := 0
	for addedTokens < amount {
		var newToken = generateToken()
		if _, isUsed := tokens[newToken]; !isUsed {
			tokens[newToken] = exists
			addedTokens += 1
		}
	}
	return tokens
}
