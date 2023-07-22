package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var (
	DB   *sql.DB
	once sync.Once
)

func createTables() {
	createTablesSQL := `
		CREATE TABLE IF NOT EXISTS categories (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL
		);

		CREATE TABLE IF NOT EXISTS tasks (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT,
			due_at TEXT,
			is_complete BOOLEAN DEFAULT FALSE,
			category_id INTEGER,
			FOREIGN KEY (category_id) REFERENCES categories (id)
		);
	`

	if _, err := DB.Exec(createTablesSQL); err != nil {
		log.Fatal(err)
	}
}

func createInstance() {
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		log.Fatal(err)
	}
	DB = db
}

func InitDB() {
	fmt.Println("Initializing database...")
	once.Do(func() {
		createInstance()
		createTables()
	})
	fmt.Println("Database initialized.")
}

func SeedDB() {
	fmt.Println("Seeding database...")
	_, err := DB.Exec(`
		INSERT INTO categories (name)
		VALUES
			('Work'),
			('Personal');
	`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = DB.Exec(`
		INSERT INTO tasks (name, description, due_at, is_complete, category_id)
		VALUES
			('Create a task manager', 'Use Go and Vue.js', '2024-01-01 10:00:00', 0, 1),
			('Test task manager', NULL, '2023-10-01 12:00:00', 0, 1),
			('Clean the house', 'Need to wash dishes', '2023-10-22 18:00:00', 0, 2),
			('Pick up groceries', 'Tomatoes, Wine', NULL, 0, 2),
			('Get coffee', 'a lot', NULL, 0, NULL),
			('NVIM config', 'Fix plugins', '2023-11-01 14:00:00', 0, NULL);
	`)
	fmt.Println("Database seeded.")
}
