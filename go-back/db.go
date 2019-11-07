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
	//query := "CREATE TABLE Token(value TEXT, valid INTEGER, used INTEGER);"
	database.Exec("CREATE TABLE Voting(name TEXT, id TEXT PRIMARY KEY, description TEXT, open INTEGER, ended INTEGER, votesPerToken INTEGER);")
	database.Exec("CREATE TABLE Token(value TEXT PRIMARY KEY, valid INTEGER);")
	database.Exec("CREATE TABLE Candidate(name TEXT, id TEXT PRIMARY KEY, description TEXT);")
	database.Exec("CREATE TABLE Vote(id TEXT PRIMARY KEY, votingId TEXT, candidateId TEXT, tokenId TEXT, FOREIGN KEY (candidateId) REFERENCES Candidate(id), FOREIGN KEY (votingId) REFERENCES Voting(id), FOREIGN KEY (tokenId) REFERENCES Token(value));")
	database.Exec("CREATE TABLE Mastertoken(value TEXT PRIMARY KEY);")
}
