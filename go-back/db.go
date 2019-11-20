package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func InitializeDb() {
	database, _ := sql.Open("postgres", CONNECTION_STRING)
	database.Exec("PRAGMA foreign_keys = ON")
	database.Exec("DROP TABLE Vote")
	database.Exec("DROP TABLE Token")
	database.Exec("DROP TABLE Availability")
	database.Exec("DROP TABLE Candidate")
	database.Exec("DROP TABLE Voting")
	database.Exec("DROP TABLE MasterToken")
	fmt.Println("old database deleted (if it existed)")
	fmt.Println("creating new tables")
	database.Exec("CREATE TABLE Voting(name TEXT, id TEXT PRIMARY KEY, description TEXT, open INTEGER, ended INTEGER, votesPerToken INTEGER);")
	database.Exec("CREATE TABLE Token(value TEXT PRIMARY KEY, valid INTEGER, hidden_id SERIAL);")
	database.Exec("CREATE TABLE Candidate(name TEXT, id TEXT PRIMARY KEY, description TEXT);")
	database.Exec("CREATE TABLE Availability(candidateId TEXT REFERENCES Candidate(id) ON DELETE CASCADE ON UPDATE CASCADE, votingId TEXT REFERENCES Voting(id) ON DELETE CASCADE ON UPDATE CASCADE, PRIMARY KEY (candidateId, votingId));")
	database.Exec("CREATE TABLE Vote(id TEXT PRIMARY KEY, votingId TEXT REFERENCES Voting(id) ON DELETE CASCADE ON UPDATE CASCADE, candidateId TEXT REFERENCES Candidate(id) ON DELETE CASCADE ON UPDATE CASCADE, token TEXT REFERENCES Token(value) ON DELETE CASCADE ON UPDATE CASCADE);")
	database.Exec("CREATE TABLE Mastertoken(value TEXT PRIMARY KEY);")
}
