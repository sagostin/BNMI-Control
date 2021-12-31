package db

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"os"
)

// connections must be closed manually.

func ConnectPool() (*pgxpool.Pool, error) {
	poolConfig, err := pgxpool.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Unable to parse DATABASE_URL\n----- Error -----\n", err)
		return nil, err
	}

	db, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		log.Fatal("Unable to create connection pool\n----- Error -----\n", err)
		return nil, err
	}
	return db, nil
}

func InitDatabase(pool *pgxpool.Pool) error {
	err := pool.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS TEAMS (TEAMNO INTEGER NOT NULL PRIMARY KEY,EmployeeNO    INTEGER NOT NULL,\n> DIVISION    CHAR(6) NOT NULL);", table).Scan(&counter)
	if err != nil {
		return 0, err
	}
}

func GetTableCount(pool *pgxpool.Pool, table string) (int, error) {
	var counter int
	err := pool.QueryRow(context.Background(), "SELECT count(*) FROM $1", table).Scan(&counter)
	if err != nil {
		return 0, err
	}
	return counter, nil
}
