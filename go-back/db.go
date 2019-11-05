package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os/exec"
)

func InitializeDb() {
	exec.Command("rm", DB_NAME).Output()
	fmt.Println("old database deleted (if it existed)")
	fmt.Println("creating new tables")
	database, _ := sql.Open("sqlite3", DB_NAME)
	query := "CREATE TABLE Token(value TEXT, valid INTEGER, used INTEGER);"
	query += "CREATE TABLE Voting(title TEXT, id TEXT, description TEXT, open INTEGER, ended INTEGER, votesPerToken INTEGER);"
	query += "CREATE TABLE Candidate(name TEXT, id TEXT, description TEXT);"
	query += "CREATE TABLE Vote(id TEXT, votingId TEXT, candidateId TEXT, tokenId TEXT);"
	statement, _ := database.Prepare(query)
	statement.Exec()
}
