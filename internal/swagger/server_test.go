package swagger

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/khaiql/dbcleaner"
	"github.com/khaiql/dbcleaner/engine"
)

var db *sql.DB
var dbCleaner dbcleaner.DbCleaner

func init() {
	dsn := "blog:blog@tcp(localhost:3306)/blog?parseTime=true"

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("open mysql connection: %s", err)
	}

	dbCleaner = dbcleaner.New()
	dbCleaner.SetEngine(engine.NewMySQLEngine(dsn))
}

func testHandler() (handler http.Handler) {
	server := NewServer(db, "", 80)
	return server.GetHandler()
}
