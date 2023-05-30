package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	postgres_host     = "dpg-chqr8lu7avjb90mv1j5g-a.singapore-postgres.render.com"
	postgres_port     = 5432
	postgres_user     = "postgres_admin"
	postgres_password = "WecY2cM3z9wedEedYhpWPmaLqtXLHWmW"
	postgres_dbname   = "my_db_xvdh"
)

var Db *sql.DB

func init() {
	db_info := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s ", postgres_host, postgres_port, postgres_user, postgres_password, postgres_dbname)

	var err error
	Db, err = sql.Open("postgres", db_info)

	if err != nil {
		panic(err)
	} else {
		log.Println("Database successfully configured")
	}
}
