package main

import (
	//"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	//"net/url"
	"os"
	"strings"
	//"strconv"
)

const DB_NAME string = "vaalit.db"

func main() {
	fmt.Println("arguments")
	fmt.Println(os.Args[1])
	switch os.Args[1] {
	case "start":
		fmt.Println("starting server...")
	case "prepare":
		InitializeDb()
	case "tokens":
		GenerateTokens(1000)
	}
	fmt.Println("opening database")
	//database, _ := sql.Open("sqlite3", "./vaalit.db")
	//statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	//statement.Exec()
	//statement, _ = database.Prepare("INSERT INTO people (firstname, lastname) VALUES (?, ?)")
	//statement.Exec("Rob", "Gronkowski")
	//rows, _ := database.Query("SELECT id, firstname, lastname FROM people")
	//var id int
	//var firstname string
	//var lastname string
	//for rows.Next() {
	//rows.Scan(&id, &firstname, &lastname)
	//fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
	//}
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/api", HandleTokenEvent)
	http.ListenAndServe(":8281", nil)
}

func HandleTokenEvent(w http.ResponseWriter, r *http.Request) {
	keys, _ := r.URL.Query()["action"]
	fmt.Println(keys[0] == "generate")
	fmt.Fprintf(w, "asdf, %s!", r.URL.Path[1:])
}
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func toJSON(m interface{}) string {
	js, err := json.Marshal(m)
	if err != nil {
		//log.Fatal(err)
		//Panic
	}
	return strings.ReplaceAll(string(js), ",", ", ")
}
