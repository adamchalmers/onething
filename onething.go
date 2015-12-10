package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
	"net/http"
	"time"
)

const (
	dbFile = "./db/file.db"
	port   = ":8000"
)

func main() {

	// Set up the DB
	db, err := sql.Open("sqlite3", dbFile)
	dieIfErr(err)
	defer db.Close()

	// Set up router and handling, start server
	r := mux.NewRouter()
	for k, v := range MakeHandlers(db) {
		r.HandleFunc(k, v)
	}

	fmt.Println("Listening on port " + port)
	http.ListenAndServe(port, r)
}

func dieIfErr(e error) {
	if e != nil {
		panic(e)
	}
}

// Returns a map of url patterns to the handlers that should handle them.
func MakeHandlers(db *sql.DB) map[string]func(http.ResponseWriter, *http.Request) {
	handlers := make(map[string]func(http.ResponseWriter, *http.Request))
	handlers["/"] = func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("tmpl/index.html")
		t.Execute(w, struct{}{})
	}

	handlers["/user/{username}"] = func(w http.ResponseWriter, r *http.Request) {
		rows := DoQuery(db, "SELECT * FROM items WHERE username = ? ORDER BY postWhen", mux.Vars(r)["username"])
		items := make([]Item, 0)
		for rows.Next() {
			var username string
			var link string
			var title string
			var when time.Time
			rows.Scan(&username, &link, &title, &when)
			items = append(items, Item{link, title, when.Format(time.UnixDate)})
		}

		t, _ := template.ParseFiles("tmpl/user.html")
		t.Execute(w, struct {
			Username string
			Items    []Item
		}{
			mux.Vars(r)["username"],
			items,
		})

	}
	return handlers
}

// DoQuery executes a template query with given parameters and returns the results.
func DoQuery(db *sql.DB, statement string, args ...interface{}) *sql.Rows {
	stmt, err := db.Prepare(statement)
	dieIfErr(err)
	defer stmt.Close()
	rows, err := stmt.Query(args...)
	return rows
}

// put Prints a string to a writer.
func put(w http.ResponseWriter, s string) {
	w.Write([]byte(s))
}

// Represents an item posted by a user.
type Item struct {
	Url   string
	Title string
	When  string
}
