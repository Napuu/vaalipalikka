package main

import (
	//"database/sql"
	_ "github.com/lib/pq"
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
		//db, _ := sql.Open("postgres", CONNECTION_STRING)
		db.QueryRow("SELECT COUNT(*) description FROM Token WHERE value = $1 AND valid = 1", token).Scan(&isVoter)
		db.QueryRow("SELECT COUNT(*) description FROM Mastertoken WHERE value = $1", token).Scan(&isAdmin)
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
