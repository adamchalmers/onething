package main

import (
	"database/sql"
	"fmt"
	// "github.com/adamchalmers/onething/handlers"
	. "github.com/adamchalmers/onething/util"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
)

const (
	dbFile = "./db/file.db"
	port   = ":8000"
)

func main() {

	// Set up the DB
	db, err := sql.Open("sqlite3", dbFile)
	DieIfErr(err)
	defer db.Close()

	// Set up router and handling, start server
	r := mux.NewRouter()
	for k, v := range MakeHandlers(db) {
		r.HandleFunc(k, v)
	}

	fmt.Println("Listening on port " + port)
	http.ListenAndServe(port, r)
}
