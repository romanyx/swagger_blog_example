package main

import (
	"database/sql"
	"flag"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/romanyx/swagger_blog_example/internal/swagger"
)

var (
	port = flag.Int("port", 3332, "server port")
	host = flag.String("host", "localhost", "server host")
	dsn  = flag.String("mysql", "blog:blog@tcp(localhost:3306)/blog?parseTime=true", "mysql connection string")
)

func main() {
	flag.Parse()

	db, err := sql.Open("mysql", *dsn)
	if err != nil {
		log.Fatalf("open sql connection failed: %v\n", err)
	}

	s := swagger.NewServer(db, *host, *port)

	if err := s.Serve(); err != nil {
		log.Fatalf("server serve: %s", err)
	}
}
