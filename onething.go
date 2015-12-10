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

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	for k, v := range MakeHandlers(db) {
		r.HandleFunc(k, v)
	}

	// Bind to a port and pass our router in
	fmt.Println("Listening on port " + port)
	http.ListenAndServe(port, r)
}

func dieIfErr(e error) {
	if e != nil {
		panic(e)
	}
}

func MakeHandlers(db *sql.DB) map[string]func(http.ResponseWriter, *http.Request) {
	handlers := make(map[string]func(http.ResponseWriter, *http.Request))
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

//
type Item struct {
	Url   string
	Title string
	When  string
}
