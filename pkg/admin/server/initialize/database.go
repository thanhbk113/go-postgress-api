package initialize

import (
	"database/sql"
	"fmt"
	"log"
	dbConn "thanhbk113/db/sqlc"
	"thanhbk113/internal/config"

	_ "github.com/lib/pq"
	"github.com/logrusorgru/aurora"
)

var (
	db    *dbConn.Queries
	sqlDB *sql.DB
)

func database() {

	conn, err := sql.Open(config.GetConfig().PostgreDriver, config.GetConfig().PostgresSource)
	if err != nil {
		log.Fatalf("could not connect to postgres database: %v", err)
	}

	db = dbConn.New(conn)
	sqlDB = conn
	fmt.Println(aurora.Green("***  PostgreSQL connected successfully:" + " ***"))
}

// GetDB
func GetDB() *dbConn.Queries {
	return db
}

// GetSQLDB
func GetSQLDB() *sql.DB {
	return sqlDB
}
