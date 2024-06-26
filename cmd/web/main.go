package main

// cmd will contain the app-specific code for the execs
// in the project. we'll have 1 exec so far, will live here

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	// prefixed with an _ because we need its init function
	// we don't use anything else from this package
)

// we could also make these global variables but this is best practice
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// command line flag named addr, default is ":4000"
	// and some short text to explain what it controls
	// having these values hardcoded isn't ideal
	addr := flag.String("addr", ":4000", "HTTP network address")
	// we use String because if the val can't be an str
	// it will log an error and exit
	dsn := flag.String("dsn", "web:2467@/snippetbox?parseTime=true", "MySQL data source name")
	// mysql dsn string command line flag
	flag.Parse()
	// we need to call Parse before we use the vars

	// creating a logger for info messages
	// 3 parameters: where to write logs (os.Stdout)
	// a string prefix (INFO and tab) and flags for more info
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	// logger for error messages, using Stderr as output
	// and Lshortfile to include the file name and line number
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	// log.New loggers are safe from race conditions

	// creating a connection pool for the MySQL database
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	// we need to close the connection pool before the main() func exits
	defer db.Close()
	// this isn't really necessary but it's good practice
	// for when we add a graceful shutdown feature

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// we need to change http.Server's defaults to use our error logger
	// instead of the default one which it uses
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("starting server on %s", *addr)
	err = srv.ListenAndServe()
	// because we already declared err above, we use = instead of :=
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	// opening a connection to the MySQL database
	db, err := sql.Open("mysql", dsn)
	// sql.Open doesn't actually create a connection
	// it just initializes the database object

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		// to check for errors we need to use Ping
		return nil, err
	}
	return db, nil
}
