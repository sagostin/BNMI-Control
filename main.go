package main

import (
	db2 "BNMI-Control/db"
	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
	"log"
)

var dbPool *pgxpool.Pool

func main() {

	var err error
	dbPool, err = db2.ConnectPool()
	if err != nil {
		log.Fatalln(err)
	}

	db2.GetTableCount(dbPool, "users")

	//println(db.Stats().InUse)

	//age := 21
	//rows, err := db.Query("SELECT name FROM users WHERE age = $1", age)

	/*http.HandleFunc("/", urlHandler)

	log.Info("Starting URL shortener on localhost:8080")
	err = http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Crit("Unable to start web server", "error", err)
		os.Exit(1)
	}*/
}
