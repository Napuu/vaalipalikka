package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	//"log"
	//"math/rand"
	//"strconv"
	//"time"
	"fmt"
	"net/http"
)

func HandleLoginApiQuery(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token != "" {
		isVoter := 0
		isAdmin := 0
		db, _ := sql.Open("sqlite3", "./vaalit.db")
		db.QueryRow("SELECT COUNT(*) description FROM Token WHERE value = ? AND valid = 1", token).Scan(&isVoter)
		db.QueryRow("SELECT COUNT(*) description FROM Mastertoken WHERE value = ?", token).Scan(&isAdmin)
		if isVoter == 1 {
			fmt.Fprintf(w, "voter")
		} else if isAdmin == 1 {
			fmt.Fprintf(w, "admin")
		} else {
			fmt.Fprintf(w, "denied")
		}
	}
	//params := r.URL.Query()
	//action, actionExists := params["action"]
}
