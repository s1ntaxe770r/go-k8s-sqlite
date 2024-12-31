package squeal

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Write saves the level-order traversal result to a SQLite database
func Write(db *sql.DB, levelOrder []int) error {
	// Create the table if it doesn't exist
	createTableSQL := `
    CREATE TABLE IF NOT EXISTS level_order (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        value INTEGER NOT NULL
    );`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		return fmt.Errorf("failed to create table: %v", err)
	}

	// Insert values into the table
	insertSQL := `INSERT INTO level_order (value) VALUES (?);`
	for _, value := range levelOrder {
		_, err := db.Exec(insertSQL, value)
		if err != nil {
			return fmt.Errorf("failed to insert value %d: %v", value, err)
		}
	}

	return nil
}
