package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func InitializeDb() {
	database, _ := sql.Open("sqlite3", DB_NAME)
	database.Exec("DROP TABLE Vote")
	database.Exec("DROP TABLE Token")
	database.Exec("DROP TABLE Availability")
	database.Exec("DROP TABLE Candidate")
	database.Exec("DROP TABLE Voting")
	database.Exec("DROP TABLE MasterToken")
	fmt.Println("old database deleted (if it existed)")
	fmt.Println("creating new tables")
	//query := "CREATE TABLE Token(value TEXT, valid INTEGER, used INTEGER);"
	database.Exec("CREATE TABLE Voting(name TEXT, id TEXT PRIMARY KEY, description TEXT, open INTEGER, ended INTEGER, votesPerToken INTEGER);")
	database.Exec("CREATE TABLE Token(value TEXT PRIMARY KEY, valid INTEGER);")
	database.Exec("CREATE TABLE Candidate(name TEXT, id TEXT PRIMARY KEY, description TEXT);")
	database.Exec("CREATE TABLE Availability(candidateId TEXT, votingId TEXT, FOREIGN KEY (candidateId) REFERENCES Candidate(id), FOREIGN KEY(votingId) REFERENCES Voting(id), PRIMARY KEY (candidateId, votingId));")
	database.Exec("CREATE TABLE Vote(id TEXT PRIMARY KEY, votingId TEXT, candidateId TEXT, tokenId TEXT, FOREIGN KEY (candidateId) REFERENCES Candidate(id), FOREIGN KEY (votingId) REFERENCES Voting(id), FOREIGN KEY (tokenId) REFERENCES Token(value));")
	database.Exec("CREATE TABLE Mastertoken(value TEXT PRIMARY KEY);")
}
