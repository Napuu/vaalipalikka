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

func Permissions(token string) int {
	if token != "" {
		isVoter := 0
		isAdmin := 0
		db.QueryRow("SELECT COUNT(*) FROM Token WHERE value = $1 AND valid = 1", token).Scan(&isVoter)
		db.QueryRow("SELECT COUNT(*) FROM Mastertoken WHERE value = $1", token).Scan(&isAdmin)
		if isVoter == 1 {
			return 1
		} else if isAdmin == 1 {
			return 2
		}
	}
	return 0;
}

func CheckQueryPermissions(w http.ResponseWriter, r *http.Request, expected int) bool {
	token := r.Header.Get("Authorization")
	permissions := Permissions(token)
	if permissions != expected {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("401 - Unauthorized"))
		return false
	}
	return true
}

func HandleLoginApiQuery(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	permissions := Permissions(token)
	if permissions == 1 {
		fmt.Fprintf(w, "voter")
	} else if permissions == 2 {
		fmt.Fprintf(w, "admin")
	} else {
		fmt.Fprintf(w, "denied")
	}
}
