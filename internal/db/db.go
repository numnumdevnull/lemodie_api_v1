package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect(dsn string) *sql.DB {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("db.Connect: не вдалось відкрити з'єднання: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("db.Connect: не вдалось пінгонути БД: %v", err)
	}

	fmt.Println("✅ MySQL підключено")
	return db
}
