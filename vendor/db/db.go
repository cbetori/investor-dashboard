package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

//Database setup
var Db *sql.DB

func Connection() {
	db_uri := os.Getenv("db_uri")
	db, err := sql.Open("postgres", db_uri)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	Db = db
	fmt.Println("Successfully connected!")
}

func Lite() {
	os.Remove("sqlite-database.db")

	log.Println("Creating sqlite-database.db...")
	file, err := os.Create("sqlite-database.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("sqlite-database.db created")

	sqliteDatabase, err := sql.Open("sqlite3", "./sqlite-database.db")
	if err != nil {
		log.Fatal(err.Error())
	}

	createTables(sqliteDatabase)
	Db = sqliteDatabase
	fmt.Println("Successfully connected!")
}
