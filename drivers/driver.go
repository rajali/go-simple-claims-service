package drivers

import (
	"database/sql"
	"log"
	"os"

	"github.com/lib/pq"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectDB() *sql.DB {
	pgUrl, err := pq.ParseURL(os.Getenv("POSTGRESSQL_URL"))
	logFatal(err)

	db, err := sql.Open("postgres", pgUrl)
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	return db
}
