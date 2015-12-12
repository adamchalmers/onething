package main

import (
	"database/sql"
	. "github.com/adamchalmers/onething/util"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"time"
)

// Returns a map of url patterns to the handlers that should handle them.
func MakeHandlers(db *sql.DB) map[string]func(http.ResponseWriter, *http.Request) {
	handlers := make(map[string]func(http.ResponseWriter, *http.Request))

	// Index
	handlers["/"] = func(w http.ResponseWriter, r *http.Request) {
		me := "adam_chal"
		items := itemsFromRows(DoQuery(db,
			`SELECT i.username, i.link, i.title, i.postWhen 
     FROM items i INNER JOIN follows f ON f.following = i.username 
     WHERE f.username = ? ORDER BY postWhen DESC`, me))
		t, _ := template.ParseFiles("tmpl/index.html")

		t.Execute(w, struct {
			Me    string
			Items []Item
		}{
			me,
			items,
		})
	}

	// User profile
	handlers["/user/{username}"] = func(w http.ResponseWriter, r *http.Request) {
		items := itemsFromRows(DoQuery(db,
			`SELECT * FROM items 
       WHERE username = ? 
       ORDER BY postWhen`, mux.Vars(r)["username"]))

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

// itemFromRows returns the Items that an SQL query result represents.
func itemsFromRows(rows *sql.Rows) []Item {
	items := make([]Item, 0)
	for rows.Next() {
		var username string
		var link string
		var title string
		var when time.Time
		rows.Scan(&username, &link, &title, &when)
		items = append(items, Item{username, link, title, when.Format(time.UnixDate)})
	}
	return items
}

// DoQuery executes a template query with given parameters and returns the results.
func DoQuery(db *sql.DB, statement string, args ...interface{}) *sql.Rows {
	stmt, err := db.Prepare(statement)
	DieIfErr(err)
	defer stmt.Close()
	rows, err := stmt.Query(args...)
	return rows
}

// Represents an item posted by a user.
type Item struct {
	Username string
	Url      string
	Title    string
	When     string
}
