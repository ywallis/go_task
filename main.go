package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/pressly/goose/v3"
	"github.com/ywallis/go_task/database"
	_ "modernc.org/sqlite"
)

// TODO
// Persist data

type task struct {
	id        int
	name      string
}

type appConfig struct {
	db *database.Queries
}

func main() {
	db, err := sql.Open("sqlite", "db.db")
	if err != nil {
		log.Fatal("Could not open DB")
	}
	defer db.Close()

	queries := database.New(db)
	cfg := appConfig{db: queries}

	goose.SetDialect("sqlite")
	if err := goose.Up(db, "sql/schemas"); err != nil {
		log.Fatalf("Goose up: %v", err)
	}
	tasks := cfg.getAllTasks()
	fmt.Println(tasks)
	task := cfg.createTask("Test!")
	fmt.Println(task)

}
