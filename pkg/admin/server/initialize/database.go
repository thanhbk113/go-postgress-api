package initialize

import (
	"database/sql"
	"fmt"
	"log"
	dbConn "thanhbk113/db/sqlc"
	"thanhbk113/internal/config"
)

var (
	db *dbConn.Queries
)

func database() {

	conn, err := sql.Open(config.GetConfig().PostgreDriver, config.GetConfig().PostgresSource)
	if err != nil {
		log.Fatalf("could not connect to postgres database: %v", err)
	}

	db = dbConn.New(conn)
	fmt.Println("PostgreSQL connected successfully...")
}

// GetDB
func GetDB() *dbConn.Queries {
	return db
}
